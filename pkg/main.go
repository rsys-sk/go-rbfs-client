package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rtbrick/go-rbfs-client/pkg/rbfs"
	"github.com/rtbrick/go-rbfs-client/pkg/rbfs/alerts"
	"github.com/rtbrick/go-rbfs-client/pkg/rbfs/metrics"
)

var client http.Client

const endpointURL string = "http://10.200.134.28:19091"
const elementName string = "ufi08.q2c.u23.r4.nbg.rtbrick.net"
const accessToken string = ""

func main() {
	m := metrics.NewClient(&client)
	endpoint, _ := url.Parse(endpointURL)
	ctx, _ := rbfs.NewRbfsContext(context.Background(), endpoint, elementName, rbfs.RbfsAccessToken(accessToken))

	metric, err := m.QueryMetric(ctx, "chassis_temperature_millicelsius")
	fmt.Println(err)
	b, _ := json.MarshalIndent(metric, " ", " ")
	fmt.Println(string(b))

	a := alerts.NewClient(&client)
	alerts, err := a.QueryAlerts(ctx)
	b, _ = json.MarshalIndent(alerts, " ", " ")
	fmt.Println(string(b))
}
