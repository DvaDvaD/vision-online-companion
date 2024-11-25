#!/bin/bash

set -e

licenses=("license" "kotg")

for license in "${licenses[@]}"; do
  docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/docs/openapi.$license.yml \
    -g k6 \
    -o /local/test/$license &
done

wait