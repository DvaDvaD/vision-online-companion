#!/bin/bash

set -e

export IMPORT_GENERIC='cfg.generic.yaml:"github.com/portierglobal/vision-online-companion/api/internal/gen/generic"'
export EXCLUDE_GENERIC='Response200,Response201,Response400,Response401,Response403,Response404'

export OPENAPI_GENERIC='docs/openapi.generic.yml'
export OPENAPI_LICENSE='docs/openapi.license.yml'
export OPENAPI_KOTG='docs/openapi.kotg.yml'

mkdir -p internal/gen/generic
oapi-codegen -generate "types,skip-prune" -package api -exclude-schemas $EXCLUDE_GENERIC $OPENAPI_GENERIC > internal/gen/generic/models.gen.go
oapi-codegen -generate "spec,skip-prune" -package api $OPENAPI_GENERIC > internal/gen/generic/spec.gen.go

mkdir -p internal/gen/license
oapi-codegen -generate "types,skip-prune" -package api $OPENAPI_LICENSE > internal/gen/license/models.gen.go
oapi-codegen -generate "echo,skip-prune" -package api $OPENAPI_LICENSE > internal/gen/license/server.gen.go
oapi-codegen -generate "spec,skip-prune" -package api -import-mapping $IMPORT_GENERIC $OPENAPI_LICENSE > internal/gen/license/spec.gen.go

mkdir -p internal/gen/kotg
oapi-codegen -generate "types,skip-prune" -package api $OPENAPI_KOTG > internal/gen/kotg/models.gen.go
oapi-codegen -generate "echo,skip-prune" -package api $OPENAPI_KOTG > internal/gen/kotg/server.gen.go
oapi-codegen -generate "spec,skip-prune" -package api -import-mapping $IMPORT_GENERIC $OPENAPI_KOTG > internal/gen/kotg/spec.gen.go