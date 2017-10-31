// Package azurereport implements the Azure ARM Azurereport service API version 1.0.0.
//
// Test Infrastructure for AutoRest
package azurereport

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

const (
	// DefaultBaseURI is the default URI used for the service Azurereport
	DefaultBaseURI = "http://localhost"
)

// ManagementClient is the base client for Azurereport.
type ManagementClient struct {
	autorest.Client
	BaseURI string
}

// New creates an instance of the ManagementClient client.
func New() ManagementClient {
	return NewWithBaseURI(DefaultBaseURI)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string) ManagementClient {
	return ManagementClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}

// GetReport get test coverage report
func (client ManagementClient) GetReport(ctx context.Context) (result SetInt32, err error) {
	req, err := client.GetReportPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurereport.ManagementClient", "GetReport", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetReportSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurereport.ManagementClient", "GetReport", resp, "Failure sending request")
		return
	}

	result, err = client.GetReportResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurereport.ManagementClient", "GetReport", resp, "Failure responding to request")
	}

	return
}

// GetReportPreparer prepares the GetReport request.
func (client ManagementClient) GetReportPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/report/azure"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetReportSender sends the GetReport request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetReportSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetReportResponder handles the response to the GetReport request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetReportResponder(resp *http.Response) (result SetInt32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
