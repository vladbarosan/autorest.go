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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PolymorphismClient is the test Infrastructure for AutoRest
type PolymorphismClient struct {
	BaseClient
}

// NewPolymorphismClient creates an instance of the PolymorphismClient client.
func NewPolymorphismClient() PolymorphismClient {
	return NewPolymorphismClientWithBaseURI(DefaultBaseURI)
}

// NewPolymorphismClientWithBaseURI creates an instance of the PolymorphismClient client.
func NewPolymorphismClientWithBaseURI(baseURI string) PolymorphismClient {
	return PolymorphismClient{NewWithBaseURI(baseURI)}
}

// GetComplicated get complex types that are polymorphic, but not at the root of the hierarchy; also have additional
// properties
func (client PolymorphismClient) GetComplicated(ctx context.Context) (result SalmonModel, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/PolymorphismClient.GetComplicated")
	defer func() {
		sc := -1
		if result.Response.Response != nil {
			sc = result.Response.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.GetComplicatedPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "GetComplicated", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetComplicatedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "GetComplicated", resp, "Failure sending request")
		return
	}

	result, err = client.GetComplicatedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "GetComplicated", resp, "Failure responding to request")
	}

	return
}

// GetComplicatedPreparer prepares the GetComplicated request.
func (client PolymorphismClient) GetComplicatedPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/polymorphism/complicated"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetComplicatedSender sends the GetComplicated request. The method will close the
// http.Response Body if it receives an error.
func (client PolymorphismClient) GetComplicatedSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetComplicatedResponder handles the response to the GetComplicated request. The method always
// closes the http.Response Body.
func (client PolymorphismClient) GetComplicatedResponder(resp *http.Response) (result SalmonModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetValid get complex types that are polymorphic
func (client PolymorphismClient) GetValid(ctx context.Context) (result FishModel, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/PolymorphismClient.GetValid")
	defer func() {
		sc := -1
		if result.Response.Response != nil {
			sc = result.Response.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.GetValidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "GetValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetValidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "GetValid", resp, "Failure sending request")
		return
	}

	result, err = client.GetValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "GetValid", resp, "Failure responding to request")
	}

	return
}

// GetValidPreparer prepares the GetValid request.
func (client PolymorphismClient) GetValidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/polymorphism/valid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetValidSender sends the GetValid request. The method will close the
// http.Response Body if it receives an error.
func (client PolymorphismClient) GetValidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetValidResponder handles the response to the GetValid request. The method always
// closes the http.Response Body.
func (client PolymorphismClient) GetValidResponder(resp *http.Response) (result FishModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutComplicated put complex types that are polymorphic, but not at the root of the hierarchy; also have additional
// properties
func (client PolymorphismClient) PutComplicated(ctx context.Context, complexBody BasicSalmon) (result autorest.Response, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/PolymorphismClient.PutComplicated")
	defer func() {
		sc := -1
		if result.Response != nil {
			sc = result.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.PutComplicatedPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutComplicated", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutComplicatedSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutComplicated", resp, "Failure sending request")
		return
	}

	result, err = client.PutComplicatedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutComplicated", resp, "Failure responding to request")
	}

	return
}

// PutComplicatedPreparer prepares the PutComplicated request.
func (client PolymorphismClient) PutComplicatedPreparer(ctx context.Context, complexBody BasicSalmon) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/polymorphism/complicated"),
		autorest.WithJSON(complexBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutComplicatedSender sends the PutComplicated request. The method will close the
// http.Response Body if it receives an error.
func (client PolymorphismClient) PutComplicatedSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutComplicatedResponder handles the response to the PutComplicated request. The method always
// closes the http.Response Body.
func (client PolymorphismClient) PutComplicatedResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutMissingDiscriminator put complex types that are polymorphic, omitting the discriminator
func (client PolymorphismClient) PutMissingDiscriminator(ctx context.Context, complexBody BasicSalmon) (result SalmonModel, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/PolymorphismClient.PutMissingDiscriminator")
	defer func() {
		sc := -1
		if result.Response.Response != nil {
			sc = result.Response.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	req, err := client.PutMissingDiscriminatorPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutMissingDiscriminator", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMissingDiscriminatorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutMissingDiscriminator", resp, "Failure sending request")
		return
	}

	result, err = client.PutMissingDiscriminatorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutMissingDiscriminator", resp, "Failure responding to request")
	}

	return
}

// PutMissingDiscriminatorPreparer prepares the PutMissingDiscriminator request.
func (client PolymorphismClient) PutMissingDiscriminatorPreparer(ctx context.Context, complexBody BasicSalmon) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/polymorphism/missingdiscriminator"),
		autorest.WithJSON(complexBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutMissingDiscriminatorSender sends the PutMissingDiscriminator request. The method will close the
// http.Response Body if it receives an error.
func (client PolymorphismClient) PutMissingDiscriminatorSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMissingDiscriminatorResponder handles the response to the PutMissingDiscriminator request. The method always
// closes the http.Response Body.
func (client PolymorphismClient) PutMissingDiscriminatorResponder(resp *http.Response) (result SalmonModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutValid put complex types that are polymorphic
// Parameters:
// complexBody - please put a salmon that looks like this:
// {
// 'fishtype':'Salmon',
// 'location':'alaska',
// 'iswild':true,
// 'species':'king',
// 'length':1.0,
// 'siblings':[
// {
// 'fishtype':'Shark',
// 'age':6,
// 'birthday': '2012-01-05T01:00:00Z',
// 'length':20.0,
// 'species':'predator',
// },
// {
// 'fishtype':'Sawshark',
// 'age':105,
// 'birthday': '1900-01-05T01:00:00Z',
// 'length':10.0,
// 'picture': new Buffer([255, 255, 255, 255, 254]).toString('base64'),
// 'species':'dangerous',
// },
// {
// 'fishtype': 'goblin',
// 'age': 1,
// 'birthday': '2015-08-08T00:00:00Z',
// 'length': 30.0,
// 'species': 'scary',
// 'jawsize': 5
// }
// ]
// };
func (client PolymorphismClient) PutValid(ctx context.Context, complexBody BasicFish) (result autorest.Response, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/PolymorphismClient.PutValid")
	defer func() {
		sc := -1
		if result.Response != nil {
			sc = result.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	if err := validation.Validate([]validation.Validation{
		{TargetValue: complexBody,
			Constraints: []validation.Constraint{{Target: "complexBody.Length", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("complexgroup.PolymorphismClient", "PutValid", err.Error())
	}

	req, err := client.PutValidPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutValidSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutValid", resp, "Failure sending request")
		return
	}

	result, err = client.PutValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutValid", resp, "Failure responding to request")
	}

	return
}

// PutValidPreparer prepares the PutValid request.
func (client PolymorphismClient) PutValidPreparer(ctx context.Context, complexBody BasicFish) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/polymorphism/valid"),
		autorest.WithJSON(complexBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutValidSender sends the PutValid request. The method will close the
// http.Response Body if it receives an error.
func (client PolymorphismClient) PutValidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutValidResponder handles the response to the PutValid request. The method always
// closes the http.Response Body.
func (client PolymorphismClient) PutValidResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutValidMissingRequired put complex types that are polymorphic, attempting to omit required 'birthday' field - the
// request should not be allowed from the client
// Parameters:
// complexBody - please attempt put a sawshark that looks like this, the client should not allow this data to
// be sent:
// {
// "fishtype": "sawshark",
// "species": "snaggle toothed",
// "length": 18.5,
// "age": 2,
// "birthday": "2013-06-01T01:00:00Z",
// "location": "alaska",
// "picture": base64(FF FF FF FF FE),
// "siblings": [
// {
// "fishtype": "shark",
// "species": "predator",
// "birthday": "2012-01-05T01:00:00Z",
// "length": 20,
// "age": 6
// },
// {
// "fishtype": "sawshark",
// "species": "dangerous",
// "picture": base64(FF FF FF FF FE),
// "length": 10,
// "age": 105
// }
// ]
// }
func (client PolymorphismClient) PutValidMissingRequired(ctx context.Context, complexBody BasicFish) (result autorest.Response, err error) {
	ctx = tracing.StartSpan(ctx, "generated/body-complex/PolymorphismClient.PutValidMissingRequired")
	defer func() {
		sc := -1
		if result.Response != nil {
			sc = result.Response.StatusCode
		}
		tracing.EndSpan(ctx, sc, err)
	}()
	if err := validation.Validate([]validation.Validation{
		{TargetValue: complexBody,
			Constraints: []validation.Constraint{{Target: "complexBody.Length", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("complexgroup.PolymorphismClient", "PutValidMissingRequired", err.Error())
	}

	req, err := client.PutValidMissingRequiredPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutValidMissingRequired", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutValidMissingRequiredSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutValidMissingRequired", resp, "Failure sending request")
		return
	}

	result, err = client.PutValidMissingRequiredResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphismClient", "PutValidMissingRequired", resp, "Failure responding to request")
	}

	return
}

// PutValidMissingRequiredPreparer prepares the PutValidMissingRequired request.
func (client PolymorphismClient) PutValidMissingRequiredPreparer(ctx context.Context, complexBody BasicFish) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/polymorphism/missingrequired/invalid"),
		autorest.WithJSON(complexBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutValidMissingRequiredSender sends the PutValidMissingRequired request. The method will close the
// http.Response Body if it receives an error.
func (client PolymorphismClient) PutValidMissingRequiredSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutValidMissingRequiredResponder handles the response to the PutValidMissingRequired request. The method always
// closes the http.Response Body.
func (client PolymorphismClient) PutValidMissingRequiredResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
