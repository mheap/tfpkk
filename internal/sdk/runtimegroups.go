// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"konnect/internal/sdk/pkg/models/operations"
	"konnect/internal/sdk/pkg/models/shared"
	"konnect/internal/sdk/pkg/utils"
	"net/http"
	"strings"
)

type runtimeGroups struct {
	sdkConfiguration sdkConfiguration
}

func newRuntimeGroups(sdkConfig sdkConfiguration) *runtimeGroups {
	return &runtimeGroups{
		sdkConfiguration: sdkConfig,
	}
}

// CreateRuntimeGroup - Create Runtime Group
// Create a runtime group in the Konnect Organization.
func (s *runtimeGroups) CreateRuntimeGroup(ctx context.Context, request shared.CreateRuntimeGroupRequest) (*operations.CreateRuntimeGroupResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/runtime-groups"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

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
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateRuntimeGroupResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RuntimeGroup
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RuntimeGroup = out
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.BadRequestError = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.UnauthorizedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UnauthorizedError = out
		}
	case httpRes.StatusCode == 403:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.ForbiddenError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ForbiddenError = out
		}
	case httpRes.StatusCode == 409:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.ConflictError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ConflictError = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *operations.CreateRuntimeGroupInternalServerError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.InternalServerError = out
		}
	case httpRes.StatusCode == 503:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *operations.CreateRuntimeGroupServiceUnavailableError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ServiceUnavailableError = out
		}
	}

	return res, nil
}

// DeleteRuntimeGroup - Delete Runtime Group
// Delete an individual runtime group.
func (s *runtimeGroups) DeleteRuntimeGroup(ctx context.Context, request operations.DeleteRuntimeGroupRequest) (*operations.DeleteRuntimeGroupResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/runtime-groups/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/problem+json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.DeleteRuntimeGroupResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.BadRequestError = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.UnauthorizedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UnauthorizedError = out
		}
	case httpRes.StatusCode == 403:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.ForbiddenError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ForbiddenError = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.NotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.NotFoundError = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *operations.DeleteRuntimeGroupInternalServerError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.InternalServerError = out
		}
	case httpRes.StatusCode == 503:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *operations.DeleteRuntimeGroupServiceUnavailableError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ServiceUnavailableError = out
		}
	}

	return res, nil
}

// GetRuntimeGroup - Fetch Runtime Group
// Returns information about a team from a given team ID.
func (s *runtimeGroups) GetRuntimeGroup(ctx context.Context, request operations.GetRuntimeGroupRequest) (*operations.GetRuntimeGroupResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/runtime-groups/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.GetRuntimeGroupResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RuntimeGroup
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RuntimeGroup = out
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.BadRequestError = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.UnauthorizedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UnauthorizedError = out
		}
	case httpRes.StatusCode == 403:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.ForbiddenError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ForbiddenError = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.NotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.NotFoundError = out
		}
	case httpRes.StatusCode == 503:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *operations.GetRuntimeGroupServiceUnavailableError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ServiceUnavailableError = out
		}
	}

	return res, nil
}

// UpdateRuntimeGroup - Update Runtime Group
// Update an individual runtime group.
func (s *runtimeGroups) UpdateRuntimeGroup(ctx context.Context, request operations.UpdateRuntimeGroupRequest) (*operations.UpdateRuntimeGroupResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/runtime-groups/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "UpdateRuntimeGroupRequest", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

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
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateRuntimeGroupResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RuntimeGroup
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RuntimeGroup = out
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.BadRequestError = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.UnauthorizedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UnauthorizedError = out
		}
	case httpRes.StatusCode == 403:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.ForbiddenError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ForbiddenError = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *shared.NotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.NotFoundError = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *operations.UpdateRuntimeGroupInternalServerError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.InternalServerError = out
		}
	case httpRes.StatusCode == 503:
		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out *operations.UpdateRuntimeGroupServiceUnavailableError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ServiceUnavailableError = out
		}
	}

	return res, nil
}
