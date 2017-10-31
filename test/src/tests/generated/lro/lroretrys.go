package lrogroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// LRORetrysClient is the long-running Operation for AutoRest
type LRORetrysClient struct {
	ManagementClient
}

// NewLRORetrysClient creates an instance of the LRORetrysClient client.
func NewLRORetrysClient() LRORetrysClient {
	return NewLRORetrysClientWithBaseURI(DefaultBaseURI)
}

// NewLRORetrysClientWithBaseURI creates an instance of the LRORetrysClient client.
func NewLRORetrysClientWithBaseURI(baseURI string) LRORetrysClient {
	return LRORetrysClient{NewWithBaseURI(baseURI)}
}

// Delete202Retry200 long running delete request, service returns a 500, then a 202 to the initial request. Polls
// return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client LRORetrysClient) Delete202Retry200(ctx context.Context) (result LRORetrysDelete202Retry200Future, err error) {
	req, err := client.Delete202Retry200Preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Delete202Retry200", nil, "Failure preparing request")
		return
	}

	result, err = client.Delete202Retry200Sender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Delete202Retry200", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// Delete202Retry200Preparer prepares the Delete202Retry200 request.
func (client LRORetrysClient) Delete202Retry200Preparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/delete/202/retry/200"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Delete202Retry200Sender sends the Delete202Retry200 request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) Delete202Retry200Sender(req *http.Request) (LRORetrysDelete202Retry200Future, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LRORetrysDelete202Retry200Future{Future: future}
	return f, err
}

// Delete202Retry200Responder handles the response to the Delete202Retry200 request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) Delete202Retry200Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAsyncRelativeRetrySucceeded long running delete request, service returns a 500, then a 202 to the initial
// request. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client LRORetrysClient) DeleteAsyncRelativeRetrySucceeded(ctx context.Context) (result LRORetrysDeleteAsyncRelativeRetrySucceededFuture, err error) {
	req, err := client.DeleteAsyncRelativeRetrySucceededPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "DeleteAsyncRelativeRetrySucceeded", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteAsyncRelativeRetrySucceededSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "DeleteAsyncRelativeRetrySucceeded", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// DeleteAsyncRelativeRetrySucceededPreparer prepares the DeleteAsyncRelativeRetrySucceeded request.
func (client LRORetrysClient) DeleteAsyncRelativeRetrySucceededPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/deleteasync/retry/succeeded"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteAsyncRelativeRetrySucceededSender sends the DeleteAsyncRelativeRetrySucceeded request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) DeleteAsyncRelativeRetrySucceededSender(req *http.Request) (LRORetrysDeleteAsyncRelativeRetrySucceededFuture, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LRORetrysDeleteAsyncRelativeRetrySucceededFuture{Future: future}
	return f, err
}

// DeleteAsyncRelativeRetrySucceededResponder handles the response to the DeleteAsyncRelativeRetrySucceeded request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) DeleteAsyncRelativeRetrySucceededResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteProvisioning202Accepted200Succeeded long running delete request, service returns a 500, then a  202 to the
// initial request, with an entity that contains ProvisioningState=’Accepted’.  Polls return this value until the last
// poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client LRORetrysClient) DeleteProvisioning202Accepted200Succeeded(ctx context.Context) (result LRORetrysDeleteProvisioning202Accepted200SucceededFuture, err error) {
	req, err := client.DeleteProvisioning202Accepted200SucceededPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "DeleteProvisioning202Accepted200Succeeded", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteProvisioning202Accepted200SucceededSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "DeleteProvisioning202Accepted200Succeeded", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// DeleteProvisioning202Accepted200SucceededPreparer prepares the DeleteProvisioning202Accepted200Succeeded request.
func (client LRORetrysClient) DeleteProvisioning202Accepted200SucceededPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/delete/provisioning/202/accepted/200/succeeded"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteProvisioning202Accepted200SucceededSender sends the DeleteProvisioning202Accepted200Succeeded request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) DeleteProvisioning202Accepted200SucceededSender(req *http.Request) (LRORetrysDeleteProvisioning202Accepted200SucceededFuture, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LRORetrysDeleteProvisioning202Accepted200SucceededFuture{Future: future}
	return f, err
}

// DeleteProvisioning202Accepted200SucceededResponder handles the response to the DeleteProvisioning202Accepted200Succeeded request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) DeleteProvisioning202Accepted200SucceededResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Post202Retry200 long running post request, service returns a 500, then a 202 to the initial request, with 'Location'
// and 'Retry-After' headers, Polls return a 200 with a response body after success
//
// product is product to put
func (client LRORetrysClient) Post202Retry200(ctx context.Context, product *Product) (result LRORetrysPost202Retry200Future, err error) {
	req, err := client.Post202Retry200Preparer(ctx, product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Post202Retry200", nil, "Failure preparing request")
		return
	}

	result, err = client.Post202Retry200Sender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Post202Retry200", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// Post202Retry200Preparer prepares the Post202Retry200 request.
func (client LRORetrysClient) Post202Retry200Preparer(ctx context.Context, product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/post/202/retry/200"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Post202Retry200Sender sends the Post202Retry200 request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) Post202Retry200Sender(req *http.Request) (LRORetrysPost202Retry200Future, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LRORetrysPost202Retry200Future{Future: future}
	return f, err
}

// Post202Retry200Responder handles the response to the Post202Retry200 request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) Post202Retry200Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PostAsyncRelativeRetrySucceeded long running post request, service returns a 500, then a 202 to the initial request,
// with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
//
// product is product to put
func (client LRORetrysClient) PostAsyncRelativeRetrySucceeded(ctx context.Context, product *Product) (result LRORetrysPostAsyncRelativeRetrySucceededFuture, err error) {
	req, err := client.PostAsyncRelativeRetrySucceededPreparer(ctx, product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "PostAsyncRelativeRetrySucceeded", nil, "Failure preparing request")
		return
	}

	result, err = client.PostAsyncRelativeRetrySucceededSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "PostAsyncRelativeRetrySucceeded", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// PostAsyncRelativeRetrySucceededPreparer prepares the PostAsyncRelativeRetrySucceeded request.
func (client LRORetrysClient) PostAsyncRelativeRetrySucceededPreparer(ctx context.Context, product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/postasync/retry/succeeded"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostAsyncRelativeRetrySucceededSender sends the PostAsyncRelativeRetrySucceeded request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) PostAsyncRelativeRetrySucceededSender(req *http.Request) (LRORetrysPostAsyncRelativeRetrySucceededFuture, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LRORetrysPostAsyncRelativeRetrySucceededFuture{Future: future}
	return f, err
}

// PostAsyncRelativeRetrySucceededResponder handles the response to the PostAsyncRelativeRetrySucceeded request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) PostAsyncRelativeRetrySucceededResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Put201CreatingSucceeded200 long running put request, service returns a 500, then a 201 to the initial request, with
// an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’
// with ProvisioningState=’Succeeded’
//
// product is product to put
func (client LRORetrysClient) Put201CreatingSucceeded200(ctx context.Context, product *Product) (result LRORetrysPut201CreatingSucceeded200Future, err error) {
	req, err := client.Put201CreatingSucceeded200Preparer(ctx, product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Put201CreatingSucceeded200", nil, "Failure preparing request")
		return
	}

	result, err = client.Put201CreatingSucceeded200Sender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "Put201CreatingSucceeded200", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// Put201CreatingSucceeded200Preparer prepares the Put201CreatingSucceeded200 request.
func (client LRORetrysClient) Put201CreatingSucceeded200Preparer(ctx context.Context, product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/put/201/creating/succeeded/200"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Put201CreatingSucceeded200Sender sends the Put201CreatingSucceeded200 request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) Put201CreatingSucceeded200Sender(req *http.Request) (LRORetrysPut201CreatingSucceeded200Future, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LRORetrysPut201CreatingSucceeded200Future{Future: future}
	return f, err
}

// Put201CreatingSucceeded200Responder handles the response to the Put201CreatingSucceeded200 request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) Put201CreatingSucceeded200Responder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutAsyncRelativeRetrySucceeded long running put request, service returns a 500, then a 200 to the initial request,
// with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
//
// product is product to put
func (client LRORetrysClient) PutAsyncRelativeRetrySucceeded(ctx context.Context, product *Product) (result LRORetrysPutAsyncRelativeRetrySucceededFuture, err error) {
	req, err := client.PutAsyncRelativeRetrySucceededPreparer(ctx, product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "PutAsyncRelativeRetrySucceeded", nil, "Failure preparing request")
		return
	}

	result, err = client.PutAsyncRelativeRetrySucceededSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LRORetrysClient", "PutAsyncRelativeRetrySucceeded", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// PutAsyncRelativeRetrySucceededPreparer prepares the PutAsyncRelativeRetrySucceeded request.
func (client LRORetrysClient) PutAsyncRelativeRetrySucceededPreparer(ctx context.Context, product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/retryerror/putasync/retry/succeeded"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutAsyncRelativeRetrySucceededSender sends the PutAsyncRelativeRetrySucceeded request. The method will close the
// http.Response Body if it receives an error.
func (client LRORetrysClient) PutAsyncRelativeRetrySucceededSender(req *http.Request) (LRORetrysPutAsyncRelativeRetrySucceededFuture, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LRORetrysPutAsyncRelativeRetrySucceededFuture{Future: future}
	return f, err
}

// PutAsyncRelativeRetrySucceededResponder handles the response to the PutAsyncRelativeRetrySucceeded request. The method always
// closes the http.Response Body.
func (client LRORetrysClient) PutAsyncRelativeRetrySucceededResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
