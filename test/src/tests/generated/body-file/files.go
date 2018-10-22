package filegroup

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
	"net/http"
)

// FilesClient is the test Infrastructure for AutoRest Swagger BAT
type FilesClient struct {
	BaseClient
}

// NewFilesClient creates an instance of the FilesClient client.
func NewFilesClient() FilesClient {
	return NewFilesClientWithBaseURI(DefaultBaseURI)
}

// NewFilesClientWithBaseURI creates an instance of the FilesClient client.
func NewFilesClientWithBaseURI(baseURI string) FilesClient {
	return FilesClient{NewWithBaseURI(baseURI)}
}

// GetEmptyFile get empty file
func (client FilesClient) GetEmptyFile(ctx context.Context) (result ReadCloser, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fmt.Sprintf("%s/FilesClient.GetEmptyFile", fqdn))
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetEmptyFilePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "filegroup.FilesClient", "GetEmptyFile", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEmptyFileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "filegroup.FilesClient", "GetEmptyFile", resp, "Failure sending request")
		return
	}

	result, err = client.GetEmptyFileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "filegroup.FilesClient", "GetEmptyFile", resp, "Failure responding to request")
	}

	return
}

// GetEmptyFilePreparer prepares the GetEmptyFile request.
func (client FilesClient) GetEmptyFilePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/files/stream/empty"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEmptyFileSender sends the GetEmptyFile request. The method will close the
// http.Response Body if it receives an error.
func (client FilesClient) GetEmptyFileSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetEmptyFileResponder handles the response to the GetEmptyFile request. The method always
// closes the http.Response Body.
func (client FilesClient) GetEmptyFileResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// GetFile get file
func (client FilesClient) GetFile(ctx context.Context) (result ReadCloser, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fmt.Sprintf("%s/FilesClient.GetFile", fqdn))
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetFilePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "filegroup.FilesClient", "GetFile", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "filegroup.FilesClient", "GetFile", resp, "Failure sending request")
		return
	}

	result, err = client.GetFileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "filegroup.FilesClient", "GetFile", resp, "Failure responding to request")
	}

	return
}

// GetFilePreparer prepares the GetFile request.
func (client FilesClient) GetFilePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/files/stream/nonempty"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFileSender sends the GetFile request. The method will close the
// http.Response Body if it receives an error.
func (client FilesClient) GetFileSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetFileResponder handles the response to the GetFile request. The method always
// closes the http.Response Body.
func (client FilesClient) GetFileResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// GetFileLarge get a large file
func (client FilesClient) GetFileLarge(ctx context.Context) (result ReadCloser, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fmt.Sprintf("%s/FilesClient.GetFileLarge", fqdn))
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetFileLargePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "filegroup.FilesClient", "GetFileLarge", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFileLargeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "filegroup.FilesClient", "GetFileLarge", resp, "Failure sending request")
		return
	}

	result, err = client.GetFileLargeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "filegroup.FilesClient", "GetFileLarge", resp, "Failure responding to request")
	}

	return
}

// GetFileLargePreparer prepares the GetFileLarge request.
func (client FilesClient) GetFileLargePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/files/stream/verylarge"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFileLargeSender sends the GetFileLarge request. The method will close the
// http.Response Body if it receives an error.
func (client FilesClient) GetFileLargeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetFileLargeResponder handles the response to the GetFileLarge request. The method always
// closes the http.Response Body.
func (client FilesClient) GetFileLargeResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}
