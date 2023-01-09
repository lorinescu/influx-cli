/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	_context "context"
	_fmt "fmt"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type SecretsApi interface {

	/*
	 * DeleteOrgsIDSecretsID Delete a secret from an organization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param orgID The organization ID.
	 * @param secretID The secret ID.
	 * @return ApiDeleteOrgsIDSecretsIDRequest
	 */
	DeleteOrgsIDSecretsID(ctx _context.Context, orgID string, secretID string) ApiDeleteOrgsIDSecretsIDRequest

	/*
	 * DeleteOrgsIDSecretsIDExecute executes the request
	 */
	DeleteOrgsIDSecretsIDExecute(r ApiDeleteOrgsIDSecretsIDRequest) error

	/*
	 * DeleteOrgsIDSecretsIDExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 */
	DeleteOrgsIDSecretsIDExecuteWithHttpInfo(r ApiDeleteOrgsIDSecretsIDRequest) (*_nethttp.Response, error)

	/*
	 * GetOrgsIDSecrets List all secret keys for an organization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param orgID The organization ID.
	 * @return ApiGetOrgsIDSecretsRequest
	 */
	GetOrgsIDSecrets(ctx _context.Context, orgID string) ApiGetOrgsIDSecretsRequest

	/*
	 * GetOrgsIDSecretsExecute executes the request
	 * @return SecretKeysResponse
	 */
	GetOrgsIDSecretsExecute(r ApiGetOrgsIDSecretsRequest) (SecretKeysResponse, error)

	/*
	 * GetOrgsIDSecretsExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return SecretKeysResponse
	 */
	GetOrgsIDSecretsExecuteWithHttpInfo(r ApiGetOrgsIDSecretsRequest) (SecretKeysResponse, *_nethttp.Response, error)

	/*
	 * PatchOrgsIDSecrets Update secrets in an organization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param orgID The organization ID.
	 * @return ApiPatchOrgsIDSecretsRequest
	 */
	PatchOrgsIDSecrets(ctx _context.Context, orgID string) ApiPatchOrgsIDSecretsRequest

	/*
	 * PatchOrgsIDSecretsExecute executes the request
	 */
	PatchOrgsIDSecretsExecute(r ApiPatchOrgsIDSecretsRequest) error

	/*
	 * PatchOrgsIDSecretsExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 */
	PatchOrgsIDSecretsExecuteWithHttpInfo(r ApiPatchOrgsIDSecretsRequest) (*_nethttp.Response, error)

	/*
	 * PostOrgsIDSecrets Delete secrets from an organization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param orgID The organization ID.
	 * @return ApiPostOrgsIDSecretsRequest
	 */
	PostOrgsIDSecrets(ctx _context.Context, orgID string) ApiPostOrgsIDSecretsRequest

	/*
	 * PostOrgsIDSecretsExecute executes the request
	 */
	PostOrgsIDSecretsExecute(r ApiPostOrgsIDSecretsRequest) error

	/*
	 * PostOrgsIDSecretsExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 */
	PostOrgsIDSecretsExecuteWithHttpInfo(r ApiPostOrgsIDSecretsRequest) (*_nethttp.Response, error)
}

// SecretsApiService SecretsApi service
type SecretsApiService service

type ApiDeleteOrgsIDSecretsIDRequest struct {
	ctx          _context.Context
	ApiService   SecretsApi
	orgID        string
	secretID     string
	zapTraceSpan *string
}

func (r ApiDeleteOrgsIDSecretsIDRequest) OrgID(orgID string) ApiDeleteOrgsIDSecretsIDRequest {
	r.orgID = orgID
	return r
}
func (r ApiDeleteOrgsIDSecretsIDRequest) GetOrgID() string {
	return r.orgID
}

func (r ApiDeleteOrgsIDSecretsIDRequest) SecretID(secretID string) ApiDeleteOrgsIDSecretsIDRequest {
	r.secretID = secretID
	return r
}
func (r ApiDeleteOrgsIDSecretsIDRequest) GetSecretID() string {
	return r.secretID
}

func (r ApiDeleteOrgsIDSecretsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiDeleteOrgsIDSecretsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiDeleteOrgsIDSecretsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiDeleteOrgsIDSecretsIDRequest) Execute() error {
	return r.ApiService.DeleteOrgsIDSecretsIDExecute(r)
}

func (r ApiDeleteOrgsIDSecretsIDRequest) ExecuteWithHttpInfo() (*_nethttp.Response, error) {
	return r.ApiService.DeleteOrgsIDSecretsIDExecuteWithHttpInfo(r)
}

/*
 * DeleteOrgsIDSecretsID Delete a secret from an organization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgID The organization ID.
 * @param secretID The secret ID.
 * @return ApiDeleteOrgsIDSecretsIDRequest
 */
func (a *SecretsApiService) DeleteOrgsIDSecretsID(ctx _context.Context, orgID string, secretID string) ApiDeleteOrgsIDSecretsIDRequest {
	return ApiDeleteOrgsIDSecretsIDRequest{
		ApiService: a,
		ctx:        ctx,
		orgID:      orgID,
		secretID:   secretID,
	}
}

/*
 * Execute executes the request
 */
func (a *SecretsApiService) DeleteOrgsIDSecretsIDExecute(r ApiDeleteOrgsIDSecretsIDRequest) error {
	_, err := a.DeleteOrgsIDSecretsIDExecuteWithHttpInfo(r)
	return err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 */
func (a *SecretsApiService) DeleteOrgsIDSecretsIDExecuteWithHttpInfo(r ApiDeleteOrgsIDSecretsIDRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretsApiService.DeleteOrgsIDSecretsID")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{orgID}/secrets/{secretID}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgID"+"}", _neturl.PathEscape(parameterToString(r.orgID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretID"+"}", _neturl.PathEscape(parameterToString(r.secretID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarHTTPResponse, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetOrgsIDSecretsRequest struct {
	ctx          _context.Context
	ApiService   SecretsApi
	orgID        string
	zapTraceSpan *string
}

func (r ApiGetOrgsIDSecretsRequest) OrgID(orgID string) ApiGetOrgsIDSecretsRequest {
	r.orgID = orgID
	return r
}
func (r ApiGetOrgsIDSecretsRequest) GetOrgID() string {
	return r.orgID
}

func (r ApiGetOrgsIDSecretsRequest) ZapTraceSpan(zapTraceSpan string) ApiGetOrgsIDSecretsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetOrgsIDSecretsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetOrgsIDSecretsRequest) Execute() (SecretKeysResponse, error) {
	return r.ApiService.GetOrgsIDSecretsExecute(r)
}

func (r ApiGetOrgsIDSecretsRequest) ExecuteWithHttpInfo() (SecretKeysResponse, *_nethttp.Response, error) {
	return r.ApiService.GetOrgsIDSecretsExecuteWithHttpInfo(r)
}

/*
 * GetOrgsIDSecrets List all secret keys for an organization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgID The organization ID.
 * @return ApiGetOrgsIDSecretsRequest
 */
func (a *SecretsApiService) GetOrgsIDSecrets(ctx _context.Context, orgID string) ApiGetOrgsIDSecretsRequest {
	return ApiGetOrgsIDSecretsRequest{
		ApiService: a,
		ctx:        ctx,
		orgID:      orgID,
	}
}

/*
 * Execute executes the request
 * @return SecretKeysResponse
 */
func (a *SecretsApiService) GetOrgsIDSecretsExecute(r ApiGetOrgsIDSecretsRequest) (SecretKeysResponse, error) {
	returnVal, _, err := a.GetOrgsIDSecretsExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return SecretKeysResponse
 */
func (a *SecretsApiService) GetOrgsIDSecretsExecuteWithHttpInfo(r ApiGetOrgsIDSecretsRequest) (SecretKeysResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SecretKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretsApiService.GetOrgsIDSecrets")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{orgID}/secrets"
	localVarPath = strings.Replace(localVarPath, "{"+"orgID"+"}", _neturl.PathEscape(parameterToString(r.orgID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	body, err := GunzipIfNeeded(localVarHTTPResponse)
	if err != nil {
		body.Close()
		newErr.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	localVarBody, err := _io.ReadAll(body)
	body.Close()
	if err != nil {
		newErr.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	newErr.body = localVarBody
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchOrgsIDSecretsRequest struct {
	ctx          _context.Context
	ApiService   SecretsApi
	orgID        string
	requestBody  *map[string]string
	zapTraceSpan *string
}

func (r ApiPatchOrgsIDSecretsRequest) OrgID(orgID string) ApiPatchOrgsIDSecretsRequest {
	r.orgID = orgID
	return r
}
func (r ApiPatchOrgsIDSecretsRequest) GetOrgID() string {
	return r.orgID
}

func (r ApiPatchOrgsIDSecretsRequest) RequestBody(requestBody map[string]string) ApiPatchOrgsIDSecretsRequest {
	r.requestBody = &requestBody
	return r
}
func (r ApiPatchOrgsIDSecretsRequest) GetRequestBody() *map[string]string {
	return r.requestBody
}

func (r ApiPatchOrgsIDSecretsRequest) ZapTraceSpan(zapTraceSpan string) ApiPatchOrgsIDSecretsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPatchOrgsIDSecretsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPatchOrgsIDSecretsRequest) Execute() error {
	return r.ApiService.PatchOrgsIDSecretsExecute(r)
}

func (r ApiPatchOrgsIDSecretsRequest) ExecuteWithHttpInfo() (*_nethttp.Response, error) {
	return r.ApiService.PatchOrgsIDSecretsExecuteWithHttpInfo(r)
}

/*
 * PatchOrgsIDSecrets Update secrets in an organization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgID The organization ID.
 * @return ApiPatchOrgsIDSecretsRequest
 */
func (a *SecretsApiService) PatchOrgsIDSecrets(ctx _context.Context, orgID string) ApiPatchOrgsIDSecretsRequest {
	return ApiPatchOrgsIDSecretsRequest{
		ApiService: a,
		ctx:        ctx,
		orgID:      orgID,
	}
}

/*
 * Execute executes the request
 */
func (a *SecretsApiService) PatchOrgsIDSecretsExecute(r ApiPatchOrgsIDSecretsRequest) error {
	_, err := a.PatchOrgsIDSecretsExecuteWithHttpInfo(r)
	return err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 */
func (a *SecretsApiService) PatchOrgsIDSecretsExecuteWithHttpInfo(r ApiPatchOrgsIDSecretsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretsApiService.PatchOrgsIDSecrets")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{orgID}/secrets"
	localVarPath = strings.Replace(localVarPath, "{"+"orgID"+"}", _neturl.PathEscape(parameterToString(r.orgID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.requestBody == nil {
		return nil, reportError("requestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	// body params
	localVarPostBody = r.requestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarHTTPResponse, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostOrgsIDSecretsRequest struct {
	ctx          _context.Context
	ApiService   SecretsApi
	orgID        string
	secretKeys   *SecretKeys
	zapTraceSpan *string
}

func (r ApiPostOrgsIDSecretsRequest) OrgID(orgID string) ApiPostOrgsIDSecretsRequest {
	r.orgID = orgID
	return r
}
func (r ApiPostOrgsIDSecretsRequest) GetOrgID() string {
	return r.orgID
}

func (r ApiPostOrgsIDSecretsRequest) SecretKeys(secretKeys SecretKeys) ApiPostOrgsIDSecretsRequest {
	r.secretKeys = &secretKeys
	return r
}
func (r ApiPostOrgsIDSecretsRequest) GetSecretKeys() *SecretKeys {
	return r.secretKeys
}

func (r ApiPostOrgsIDSecretsRequest) ZapTraceSpan(zapTraceSpan string) ApiPostOrgsIDSecretsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostOrgsIDSecretsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostOrgsIDSecretsRequest) Execute() error {
	return r.ApiService.PostOrgsIDSecretsExecute(r)
}

func (r ApiPostOrgsIDSecretsRequest) ExecuteWithHttpInfo() (*_nethttp.Response, error) {
	return r.ApiService.PostOrgsIDSecretsExecuteWithHttpInfo(r)
}

/*
 * PostOrgsIDSecrets Delete secrets from an organization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgID The organization ID.
 * @return ApiPostOrgsIDSecretsRequest
 */
func (a *SecretsApiService) PostOrgsIDSecrets(ctx _context.Context, orgID string) ApiPostOrgsIDSecretsRequest {
	return ApiPostOrgsIDSecretsRequest{
		ApiService: a,
		ctx:        ctx,
		orgID:      orgID,
	}
}

/*
 * Execute executes the request
 */
func (a *SecretsApiService) PostOrgsIDSecretsExecute(r ApiPostOrgsIDSecretsRequest) error {
	_, err := a.PostOrgsIDSecretsExecuteWithHttpInfo(r)
	return err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 */
func (a *SecretsApiService) PostOrgsIDSecretsExecuteWithHttpInfo(r ApiPostOrgsIDSecretsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretsApiService.PostOrgsIDSecrets")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{orgID}/secrets/{delete}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgID"+"}", _neturl.PathEscape(parameterToString(r.orgID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.secretKeys == nil {
		return nil, reportError("secretKeys is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	// body params
	localVarPostBody = r.secretKeys
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarHTTPResponse, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
