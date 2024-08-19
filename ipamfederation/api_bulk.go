/*
IPAM Federation API

The DDI IPAM Federation application enables a SaaS administrator to manage multiple IPAM systems from one central control point CSP.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipamfederation

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type BulkAPI interface {
	/*
		DelegationBulk Execute multiple operations on delegation objects.

		Use this method to create, update, or delete multiple __Delegation__ objects.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return BulkAPIDelegationBulkRequest
	*/
	DelegationBulk(ctx context.Context) BulkAPIDelegationBulkRequest

	// DelegationBulkExecute executes the request
	//  @return FederationDelegationBulkResponse
	DelegationBulkExecute(r BulkAPIDelegationBulkRequest) (*FederationDelegationBulkResponse, *http.Response, error)
}

// BulkAPIService BulkAPI service
type BulkAPIService internal.Service

type BulkAPIDelegationBulkRequest struct {
	ctx        context.Context
	ApiService BulkAPI
	body       *FederationDelegationBulkRequest
}

func (r BulkAPIDelegationBulkRequest) Body(body FederationDelegationBulkRequest) BulkAPIDelegationBulkRequest {
	r.body = &body
	return r
}

func (r BulkAPIDelegationBulkRequest) Execute() (*FederationDelegationBulkResponse, *http.Response, error) {
	return r.ApiService.DelegationBulkExecute(r)
}

/*
DelegationBulk Execute multiple operations on delegation objects.

Use this method to create, update, or delete multiple __Delegation__ objects.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BulkAPIDelegationBulkRequest
*/
func (a *BulkAPIService) DelegationBulk(ctx context.Context) BulkAPIDelegationBulkRequest {
	return BulkAPIDelegationBulkRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FederationDelegationBulkResponse
func (a *BulkAPIService) DelegationBulkExecute(r BulkAPIDelegationBulkRequest) (*FederationDelegationBulkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *FederationDelegationBulkResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "BulkAPIService.DelegationBulk")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/federation/delegation_bulk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}