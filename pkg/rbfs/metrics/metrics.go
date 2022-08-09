// Package metrics contains a client to query RBFS metrics
package metrics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rtbrick/go-rbfs-client/pkg/rbfs/commons"
	"github.com/rtbrick/go-rbfs-client/pkg/rbfs/state"
)

type (
	// LabeledValue contains a metric value with this assigned labels.
	LabeledValue struct {
		// Value holds the metric value.
		Value float64
		// Labels contains the metric value.
		Labels map[string]string
	}

	// Metric describes a single switch metric.
	Metric struct {
		// MetricName holds the metric name.
		MetricName string
		// Values holds all metric values.
		Values []LabeledValue
	}

	// Metrics providess access to the switch metrics.
	Metrics interface {
		// QueryMetric queries a single metric.
		QueryMetric(ctx commons.RbfsContext, metric string) (*Metric, error)
	}

	service struct {
		client *http.Client
	}
)

// NewClient creates a new client to query switch metrics.
func NewClient(client *http.Client) Metrics {
	return &service{client}
}

func (s *service) QueryMetric(ctx commons.RbfsContext, metric string) (*Metric, error) {

	endpoint, err := ctx.GetServiceEndpoint(commons.PrometheusServiceName)
	if err != nil {
		return nil, err
	}

	// Compose the metric query
	queryURL := fmt.Sprintf("%s/api/v1/query?query=%s", endpoint, url.QueryEscape(metric))

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, queryURL, nil)

	if err != nil {
		return nil, err
	}

	if accessToken, ok := ctx.Value(state.ContextAccessToken).(string); ok {
		request.Header.Add("Authorization", "Bearer "+accessToken)
	}

	response, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}

	var responseJSON map[string]interface{}

	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)

	decoder.Decode(&responseJSON)

	if response.StatusCode == http.StatusOK {
		if data, ok := responseJSON["data"].(map[string]interface{}); ok {
			if result, ok := data["result"].([]interface{}); ok {
				var labeledValues []LabeledValue
				for _, i := range result {
					item := i.(map[string]interface{})
					var metricValue float64
					var metricLabels = make(map[string]string)
					if metric, ok := item["metric"].(map[string]interface{}); ok {
						for k, v := range metric {
							metricLabels[k] = v.(string)
						}
					}
					if value, ok := item["value"].([]interface{}); ok && len(value) == 2 {
						metricValue, _ = strconv.ParseFloat(value[1].(string), 64)
					}
					labeledValue := LabeledValue{Value: metricValue, Labels: metricLabels}
					labeledValues = append(labeledValues, labeledValue)
				}
				if len(labeledValues) > 0 {
					m := &Metric{
						MetricName: labeledValues[0].Labels["__name__"],
						Values:     labeledValues,
					}
					return m, nil
				}
			}
		}
		return nil, fmt.Errorf("no values for %s found", metric)
	}
	return nil, fmt.Errorf("cannot resolve %s metric. Status: %d", metric, response.StatusCode)
}
