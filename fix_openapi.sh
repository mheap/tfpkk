#!/usr/bin/env bash

# Remove recursive types
yq -i '.components.schemas.InvalidParameters.items = { "$ref": "#/components/schemas/InvalidParameterChoiceItem" }' openapi.yaml
yq -i 'del(.components.schemas.DocumentTree.properties.children)' openapi.yaml
yq -i 'del(.components.schemas.BlockNode.allOf[1])' openapi.yaml

# Make API Products gateway_service only support control_plane payloads
yq -i '.components.schemas.CreateAPIProductVersionDTO.properties.gateway_service = { "$ref": "#/components/schemas/GatewayServicePayload" }' openapi.yaml
yq -i '.components.schemas.UpdateAPIProductVersionDTO.properties.gateway_service = { "$ref": "#/components/schemas/GatewayServicePayload" }' openapi.yaml