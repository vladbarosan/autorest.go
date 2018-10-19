package complexgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DictionaryClient is the test Infrastructure for AutoRest
type DictionaryClient struct {
	BaseClient
}

// NewDictionaryClient creates an instance of the DictionaryClient client.
func NewDictionaryClient() DictionaryClient {
	return NewDictionaryClientWithBaseURI(DefaultBaseURI)
}

// NewDictionaryClientWithBaseURI creates an instance of the DictionaryClient client.
func NewDictionaryClientWithBaseURI(baseURI string) DictionaryClient {
	return DictionaryClient{NewWithBaseURI(baseURI)}
}

// GetEmpty get complex types with dictionary property which is empty
func (client DictionaryClient) GetEmpty(ctx context.Context) (result DictionaryWrapper, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/DictionaryClient.GetEmpty")
	defer func() {
		sc := -1
		if result.Response.Response != nil {
			sc = result.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.GetEmptyPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEmptySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.GetEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetEmpty", resp, "Failure responding to request")
	}

	return
}

// GetEmptyPreparer prepares the GetEmpty request.
func (client DictionaryClient) GetEmptyPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/dictionary/typed/empty"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEmptySender sends the GetEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client DictionaryClient) GetEmptySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetEmptyResponder handles the response to the GetEmpty request. The method always
// closes the http.Response Body.
func (client DictionaryClient) GetEmptyResponder(resp *http.Response) (result DictionaryWrapper, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNotProvided get complex types with dictionary property while server doesn't provide a response payload
func (client DictionaryClient) GetNotProvided(ctx context.Context) (result DictionaryWrapper, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/DictionaryClient.GetNotProvided")
	defer func() {
		sc := -1
		if result.Response.Response != nil {
			sc = result.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.GetNotProvidedPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetNotProvided", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNotProvidedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetNotProvided", resp, "Failure sending request")
		return
	}

	result, err = client.GetNotProvidedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetNotProvided", resp, "Failure responding to request")
	}

	return
}

// GetNotProvidedPreparer prepares the GetNotProvided request.
func (client DictionaryClient) GetNotProvidedPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/dictionary/typed/notprovided"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNotProvidedSender sends the GetNotProvided request. The method will close the
// http.Response Body if it receives an error.
func (client DictionaryClient) GetNotProvidedSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNotProvidedResponder handles the response to the GetNotProvided request. The method always
// closes the http.Response Body.
func (client DictionaryClient) GetNotProvidedResponder(resp *http.Response) (result DictionaryWrapper, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get complex types with dictionary property which is null
func (client DictionaryClient) GetNull(ctx context.Context) (result DictionaryWrapper, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/DictionaryClient.GetNull")
	defer func() {
		sc := -1
		if result.Response.Response != nil {
			sc = result.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.GetNullPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client DictionaryClient) GetNullPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/dictionary/typed/null"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client DictionaryClient) GetNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client DictionaryClient) GetNullResponder(resp *http.Response) (result DictionaryWrapper, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetValid get complex types with dictionary property
func (client DictionaryClient) GetValid(ctx context.Context) (result DictionaryWrapper, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/DictionaryClient.GetValid")
	defer func() {
		sc := -1
		if result.Response.Response != nil {
			sc = result.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.GetValidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetValidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetValid", resp, "Failure sending request")
		return
	}

	result, err = client.GetValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "GetValid", resp, "Failure responding to request")
	}

	return
}

// GetValidPreparer prepares the GetValid request.
func (client DictionaryClient) GetValidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/dictionary/typed/valid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetValidSender sends the GetValid request. The method will close the
// http.Response Body if it receives an error.
func (client DictionaryClient) GetValidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetValidResponder handles the response to the GetValid request. The method always
// closes the http.Response Body.
func (client DictionaryClient) GetValidResponder(resp *http.Response) (result DictionaryWrapper, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutEmpty put complex types with dictionary property which is empty
// Parameters:
// complexBody - please put an empty dictionary
func (client DictionaryClient) PutEmpty(ctx context.Context, complexBody DictionaryWrapper) (result autorest.Response, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/DictionaryClient.PutEmpty")
	defer func() {
		sc := -1
		if result.Response != nil {
			sc = result.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.PutEmptyPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "PutEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutEmptySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "PutEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.PutEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "PutEmpty", resp, "Failure responding to request")
	}

	return
}

// PutEmptyPreparer prepares the PutEmpty request.
func (client DictionaryClient) PutEmptyPreparer(ctx context.Context, complexBody DictionaryWrapper) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/dictionary/typed/empty"),
		autorest.WithJSON(complexBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutEmptySender sends the PutEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client DictionaryClient) PutEmptySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutEmptyResponder handles the response to the PutEmpty request. The method always
// closes the http.Response Body.
func (client DictionaryClient) PutEmptyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutValid put complex types with dictionary property
// Parameters:
// complexBody - please put a dictionary with 5 key-value pairs: "txt":"notepad", "bmp":"mspaint",
// "xls":"excel", "exe":"", "":null
func (client DictionaryClient) PutValid(ctx context.Context, complexBody DictionaryWrapper) (result autorest.Response, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/DictionaryClient.PutValid")
	defer func() {
		sc := -1
		if result.Response != nil {
			sc = result.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.PutValidPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "PutValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutValidSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "PutValid", resp, "Failure sending request")
		return
	}

	result, err = client.PutValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.DictionaryClient", "PutValid", resp, "Failure responding to request")
	}

	return
}

// PutValidPreparer prepares the PutValid request.
func (client DictionaryClient) PutValidPreparer(ctx context.Context, complexBody DictionaryWrapper) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/dictionary/typed/valid"),
		autorest.WithJSON(complexBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutValidSender sends the PutValid request. The method will close the
// http.Response Body if it receives an error.
func (client DictionaryClient) PutValidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutValidResponder handles the response to the PutValid request. The method always
// closes the http.Response Body.
func (client DictionaryClient) PutValidResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
