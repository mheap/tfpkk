// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"Konnect/internal/sdk/pkg/models/operations"
	"Konnect/internal/sdk/pkg/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

// debug - Debug routes
type debug struct {
	sdkConfiguration sdkConfiguration
}

func newDebug(sdkConfig sdkConfiguration) *debug {
	return &debug{
		sdkConfiguration: sdkConfig,
	}
}

// GetDebugNodeLogLevel - Retrieve Node Log Level of A Node
// Retrieve the current log level of a node.
//
// See the [NGINX documentation](http://nginx.org/en/docs/ngx_core_module.html#error_log) for the list of possible return values.
func (s *debug) GetDebugNodeLogLevel(ctx context.Context, request operations.GetDebugNodeLogLevelRequest) (*operations.GetDebugNodeLogLevelResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/runtime-groups/{runtimeGroupId}/core-entities/debug/node/log-level", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetDebugNodeLogLevelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetDebugNodeLogLevel200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.GetDebugNodeLogLevel200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PutDebugClusterControlPlanesNodesLogLevelLogLevel - Set Node Log Level of All Control Plane Nodes
// Change the log level of all Control Plane nodes deployed in Hybrid (CP/DP) cluster.
//
// See the [NGINX docs](http://nginx.org/en/docs/ngx_core_module.html#error_log) for a list of accepted values.
//
// Care must be taken when changing the log level of a node to `debug` in a production environment because the disk could fill up quickly. As soon as the debug logging finishes, revert back to a higher level such as notice.
//
// It’s currently not possible to change the log level of DP and DB-less nodes.
func (s *debug) PutDebugClusterControlPlanesNodesLogLevelLogLevel(ctx context.Context, request operations.PutDebugClusterControlPlanesNodesLogLevelLogLevelRequest) (*operations.PutDebugClusterControlPlanesNodesLogLevelLogLevelResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/runtime-groups/{runtimeGroupId}/core-entities/debug/cluster/control-planes-nodes/log-level/{log_level}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutDebugClusterControlPlanesNodesLogLevelLogLevelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PutDebugClusterControlPlanesNodesLogLevelLogLevel200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PutDebugClusterControlPlanesNodesLogLevelLogLevel200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PutDebugClusterLogLevelLogLevel - Set Node Log Level of All Nodes
// Change the log level of all nodes in a cluster.
//
// See the [NGINX docs](http://nginx.org/en/docs/ngx_core_module.html#error_log) for a list of accepted values.
//
// It’s currently not possible to change the log level of DP and DB-less nodes.
//
// Currently, when a user dynamically changes the log level for the entire cluster, if a new node joins a cluster the new node will run at the previous log level, not at the log level that was previously set dynamically for the entire cluster.
func (s *debug) PutDebugClusterLogLevelLogLevel(ctx context.Context, request operations.PutDebugClusterLogLevelLogLevelRequest) (*operations.PutDebugClusterLogLevelLogLevelResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/runtime-groups/{runtimeGroupId}/core-entities/debug/cluster/log-level/{log_level}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutDebugClusterLogLevelLogLevelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PutDebugClusterLogLevelLogLevel200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PutDebugClusterLogLevelLogLevel200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PutDebugNodeLogLevelLogLevel - Set log level of a single node
// Change the log level of a node.
//
// See the [NGINX documentation](http://nginx.org/en/docs/ngx_core_module.html#error_log) for the list of possible return values.
func (s *debug) PutDebugNodeLogLevelLogLevel(ctx context.Context, request operations.PutDebugNodeLogLevelLogLevelRequest) (*operations.PutDebugNodeLogLevelLogLevelResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/runtime-groups/{runtimeGroupId}/core-entities/debug/node/log-level/{log_level}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutDebugNodeLogLevelLogLevelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PutDebugNodeLogLevelLogLevel200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PutDebugNodeLogLevelLogLevel200ApplicationJSONObject = out
		}
	}

	return res, nil
}
