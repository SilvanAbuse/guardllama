/*
AdGuard Home

AdGuard Home REST-ish API.  Our admin web interface is built on top of this REST-ish API. 

API version: 0.107
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package adguard

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


type I18nApi interface {

	/*
	ChangeLanguage Change current language.  Argument must be an ISO 639-1 two-letter code. 

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return I18nApiChangeLanguageRequest
	*/
	ChangeLanguage(ctx context.Context) I18nApiChangeLanguageRequest

	// ChangeLanguageExecute executes the request
	ChangeLanguageExecute(r I18nApiChangeLanguageRequest) (*http.Response, error)

	/*
	CurrentLanguage Get currently set language.  Result is ISO 639-1 two-letter code.  Empty result means default language. 

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return I18nApiCurrentLanguageRequest
	*/
	CurrentLanguage(ctx context.Context) I18nApiCurrentLanguageRequest

	// CurrentLanguageExecute executes the request
	CurrentLanguageExecute(r I18nApiCurrentLanguageRequest) (*http.Response, error)
}

// I18nApiService I18nApi service
type I18nApiService service

type I18nApiChangeLanguageRequest struct {
	ctx context.Context
	ApiService I18nApi
	body *string
}

// New language.  It must be known to the server and must be an ISO 639-1 two-letter code. 
func (r I18nApiChangeLanguageRequest) Body(body string) I18nApiChangeLanguageRequest {
	r.body = &body
	return r
}

func (r I18nApiChangeLanguageRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChangeLanguageExecute(r)
}

/*
ChangeLanguage Change current language.  Argument must be an ISO 639-1 two-letter code. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return I18nApiChangeLanguageRequest
*/
func (a *I18nApiService) ChangeLanguage(ctx context.Context) I18nApiChangeLanguageRequest {
	return I18nApiChangeLanguageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *I18nApiService) ChangeLanguageExecute(r I18nApiChangeLanguageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "I18nApiService.ChangeLanguage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/i18n/change_language"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type I18nApiCurrentLanguageRequest struct {
	ctx context.Context
	ApiService I18nApi
}

func (r I18nApiCurrentLanguageRequest) Execute() (*http.Response, error) {
	return r.ApiService.CurrentLanguageExecute(r)
}

/*
CurrentLanguage Get currently set language.  Result is ISO 639-1 two-letter code.  Empty result means default language. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return I18nApiCurrentLanguageRequest
*/
func (a *I18nApiService) CurrentLanguage(ctx context.Context) I18nApiCurrentLanguageRequest {
	return I18nApiCurrentLanguageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *I18nApiService) CurrentLanguageExecute(r I18nApiCurrentLanguageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "I18nApiService.CurrentLanguage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/i18n/current_language"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
