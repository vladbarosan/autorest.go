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

// ArrayClient is the test Infrastructure for AutoRest
type ArrayClient struct {
	BaseClient
}

// NewArrayClient creates an instance of the ArrayClient client.
func NewArrayClient() ArrayClient {
	return NewArrayClientWithBaseURI(DefaultBaseURI)
}

// NewArrayClientWithBaseURI creates an instance of the ArrayClient client.
func NewArrayClientWithBaseURI(baseURI string) ArrayClient {
	return ArrayClient{NewWithBaseURI(baseURI)}
}

// GetEmpty get complex types with array property which is empty
func (client ArrayClient) GetEmpty(ctx context.Context) (result ArrayWrapper, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/ArrayClient.GetEmpty")
	defer func() {
		sc := -1
		if result.Response.Response != nil {
			sc = result.Response.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.GetEmptyPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "GetEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEmptySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "GetEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.GetEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "GetEmpty", resp, "Failure responding to request")
	}

	return
}

// GetEmptyPreparer prepares the GetEmpty request.
func (client ArrayClient) GetEmptyPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/array/empty"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEmptySender sends the GetEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client ArrayClient) GetEmptySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetEmptyResponder handles the response to the GetEmpty request. The method always
// closes the http.Response Body.
func (client ArrayClient) GetEmptyResponder(resp *http.Response) (result ArrayWrapper, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNotProvided get complex types with array property while server doesn't provide a response payload
func (client ArrayClient) GetNotProvided(ctx context.Context) (result ArrayWrapper, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/ArrayClient.GetNotProvided")
	defer func() {
		sc := -1
		if result.Response.Response != nil {
			sc = result.Response.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.GetNotProvidedPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "GetNotProvided", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNotProvidedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "GetNotProvided", resp, "Failure sending request")
		return
	}

	result, err = client.GetNotProvidedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "GetNotProvided", resp, "Failure responding to request")
	}

	return
}

// GetNotProvidedPreparer prepares the GetNotProvided request.
func (client ArrayClient) GetNotProvidedPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/array/notprovided"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNotProvidedSender sends the GetNotProvided request. The method will close the
// http.Response Body if it receives an error.
func (client ArrayClient) GetNotProvidedSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNotProvidedResponder handles the response to the GetNotProvided request. The method always
// closes the http.Response Body.
func (client ArrayClient) GetNotProvidedResponder(resp *http.Response) (result ArrayWrapper, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetValid get complex types with array property
func (client ArrayClient) GetValid(ctx context.Context) (result ArrayWrapper, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/ArrayClient.GetValid")
	defer func() {
		sc := -1
		if result.Response.Response != nil {
			sc = result.Response.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.GetValidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "GetValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetValidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "GetValid", resp, "Failure sending request")
		return
	}

	result, err = client.GetValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "GetValid", resp, "Failure responding to request")
	}

	return
}

// GetValidPreparer prepares the GetValid request.
func (client ArrayClient) GetValidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/array/valid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetValidSender sends the GetValid request. The method will close the
// http.Response Body if it receives an error.
func (client ArrayClient) GetValidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetValidResponder handles the response to the GetValid request. The method always
// closes the http.Response Body.
func (client ArrayClient) GetValidResponder(resp *http.Response) (result ArrayWrapper, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutEmpty put complex types with array property which is empty
// Parameters:
// complexBody - please put an empty array
func (client ArrayClient) PutEmpty(ctx context.Context, complexBody ArrayWrapper) (result autorest.Response, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/ArrayClient.PutEmpty")
	defer func() {
		sc := -1
		if result.Response != nil {
			sc = result.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.PutEmptyPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "PutEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutEmptySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "PutEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.PutEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "PutEmpty", resp, "Failure responding to request")
	}

	return
}

// PutEmptyPreparer prepares the PutEmpty request.
func (client ArrayClient) PutEmptyPreparer(ctx context.Context, complexBody ArrayWrapper) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/array/empty"),
		autorest.WithJSON(complexBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutEmptySender sends the PutEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client ArrayClient) PutEmptySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutEmptyResponder handles the response to the PutEmpty request. The method always
// closes the http.Response Body.
func (client ArrayClient) PutEmptyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutValid put complex types with array property
// Parameters:
// complexBody - please put an array with 4 items: "1, 2, 3, 4", "", null, "&S#$(*Y", "The quick brown fox
// jumps over the lazy dog"
func (client ArrayClient) PutValid(ctx context.Context, complexBody ArrayWrapper) (result autorest.Response, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/ArrayClient.PutValid")
	defer func() {
		sc := -1
		if result.Response != nil {
			sc = result.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.PutValidPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "PutValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutValidSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "PutValid", resp, "Failure sending request")
		return
	}

	result, err = client.PutValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.ArrayClient", "PutValid", resp, "Failure responding to request")
	}

	return
}

// PutValidPreparer prepares the PutValid request.
func (client ArrayClient) PutValidPreparer(ctx context.Context, complexBody ArrayWrapper) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/array/valid"),
		autorest.WithJSON(complexBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutValidSender sends the PutValid request. The method will close the
// http.Response Body if it receives an error.
func (client ArrayClient) PutValidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutValidResponder handles the response to the PutValid request. The method always
// closes the http.Response Body.
func (client ArrayClient) PutValidResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
