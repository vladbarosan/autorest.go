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

// FlattencomplexClient is the test Infrastructure for AutoRest
type FlattencomplexClient struct {
	BaseClient
}

// NewFlattencomplexClient creates an instance of the FlattencomplexClient client.
func NewFlattencomplexClient() FlattencomplexClient {
	return NewFlattencomplexClientWithBaseURI(DefaultBaseURI)
}

// NewFlattencomplexClientWithBaseURI creates an instance of the FlattencomplexClient client.
func NewFlattencomplexClientWithBaseURI(baseURI string) FlattencomplexClient {
	return FlattencomplexClient{NewWithBaseURI(baseURI)}
}

// GetValid sends the get valid request.
func (client FlattencomplexClient) GetValid(ctx context.Context) (result MyBaseTypeModel, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/FlattencomplexClient.GetValid")
	defer func() {
		sc := result.StatusCode
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.GetValidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.FlattencomplexClient", "GetValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetValidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.FlattencomplexClient", "GetValid", resp, "Failure sending request")
		return
	}

	result, err = client.GetValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.FlattencomplexClient", "GetValid", resp, "Failure responding to request")
	}

	return
}

// GetValidPreparer prepares the GetValid request.
func (client FlattencomplexClient) GetValidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/flatten/valid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetValidSender sends the GetValid request. The method will close the
// http.Response Body if it receives an error.
func (client FlattencomplexClient) GetValidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetValidResponder handles the response to the GetValid request. The method always
// closes the http.Response Body.
func (client FlattencomplexClient) GetValidResponder(resp *http.Response) (result MyBaseTypeModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
