package formdatagroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"fmt"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"io"
	"net/http"
)

// FormdataClient is the test Infrastructure for AutoRest Swagger BAT
type FormdataClient struct {
	BaseClient
}

// NewFormdataClient creates an instance of the FormdataClient client.
func NewFormdataClient() FormdataClient {
	return NewFormdataClientWithBaseURI(DefaultBaseURI)
}

// NewFormdataClientWithBaseURI creates an instance of the FormdataClient client.
func NewFormdataClientWithBaseURI(baseURI string) FormdataClient {
	return FormdataClient{NewWithBaseURI(baseURI)}
}

// UploadFile upload file
// Parameters:
// fileContent - file to upload.
// fileName - file name to upload. Name has to be spelled exactly as written here.
func (client FormdataClient) UploadFile(ctx context.Context, fileContent io.ReadCloser, fileName string) (result ReadCloser, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fmt.Sprintf("%s/FormdataClient.UploadFile", fqdn))
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UploadFilePreparer(ctx, fileContent, fileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFile", nil, "Failure preparing request")
		return
	}

	resp, err := client.UploadFileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFile", resp, "Failure sending request")
		return
	}

	result, err = client.UploadFileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFile", resp, "Failure responding to request")
	}

	return
}

// UploadFilePreparer prepares the UploadFile request.
func (client FormdataClient) UploadFilePreparer(ctx context.Context, fileContent io.ReadCloser, fileName string) (*http.Request, error) {
	formDataParameters := map[string]interface{}{
		"fileContent": fileContent,
		"fileName":    fileName,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/formdata/stream/uploadfile"),
		autorest.WithMultiPartFormData(formDataParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UploadFileSender sends the UploadFile request. The method will close the
// http.Response Body if it receives an error.
func (client FormdataClient) UploadFileSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UploadFileResponder handles the response to the UploadFile request. The method always
// closes the http.Response Body.
func (client FormdataClient) UploadFileResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// UploadFileViaBody upload file
// Parameters:
// fileContent - file to upload.
func (client FormdataClient) UploadFileViaBody(ctx context.Context, fileContent io.ReadCloser) (result ReadCloser, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fmt.Sprintf("%s/FormdataClient.UploadFileViaBody", fqdn))
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UploadFileViaBodyPreparer(ctx, fileContent)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFileViaBody", nil, "Failure preparing request")
		return
	}

	resp, err := client.UploadFileViaBodySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFileViaBody", resp, "Failure sending request")
		return
	}

	result, err = client.UploadFileViaBodyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFileViaBody", resp, "Failure responding to request")
	}

	return
}

// UploadFileViaBodyPreparer prepares the UploadFileViaBody request.
func (client FormdataClient) UploadFileViaBodyPreparer(ctx context.Context, fileContent io.ReadCloser) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/octet-stream"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/formdata/stream/uploadfile"),
		autorest.WithFile(fileContent))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UploadFileViaBodySender sends the UploadFileViaBody request. The method will close the
// http.Response Body if it receives an error.
func (client FormdataClient) UploadFileViaBodySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UploadFileViaBodyResponder handles the response to the UploadFileViaBody request. The method always
// closes the http.Response Body.
func (client FormdataClient) UploadFileViaBodyResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}
