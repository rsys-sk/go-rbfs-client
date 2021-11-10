#! /bin/bash

#
# Copyright (C) 2021, RtBrick, Inc.
# SPDX-License-Identifier: BSD-3-Clause
#

wget https://documents.rtbrick.com/21_8_1/api/api/swagger_opsd.yaml -O opsd-openapi.yaml

# Generate code from scratch to avoid preserving previsouly generated code the recent API specification does not generate anymore
rm -r ./pkg/rbfs/state

# Generate OPSD client from OPSD OpenAPI specification
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3 generate \
    -l go -c "/local/opsd-config.json" -o "/local/pkg/rbfs/state" -i "/local/opsd-openapi.yaml"

# Remove unneeded resources
rm ./pkg/rbfs/state/.travis.yml
rm -r ./pkg/rbfs/state/.swagger-codegen
rm -r ./pkg/rbfs/state/.gitignore
rm ./pkg/rbfs/state/.swagger-codegen-ignore
rm ./pkg/rbfs/state/git_push.sh
rm ./pkg/rbfs/state/README.md
rm -rf ./pkg/rbfs/state/docs
rm -rf ./pkg/rbfs/state/api
rm opsd-openapi.yaml

# Format generated go code
gofmt -w .