#!/usr/bin/env bash
#yq '.components.schemas.InvalidParameters.items = { "$ref": "#/components/schemas/InvalidParameterChoiceItem" }' openapi.yaml > fixed.yaml
yq 'del(.components.schemas.DocumentTree.properties.children)' openapi.yaml > fixed.yaml
mv fixed.yaml openapi.yaml