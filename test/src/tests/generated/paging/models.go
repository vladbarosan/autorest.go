package paginggroup

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
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "test/src/tests/generated/paging"

// Status enumerates the values for status.
type Status string

const (
	// Accepted ...
	Accepted Status = "Accepted"
	// Canceled ...
	Canceled Status = "canceled"
	// Created ...
	Created Status = "Created"
	// Creating ...
	Creating Status = "Creating"
	// Deleted ...
	Deleted Status = "Deleted"
	// Deleting ...
	Deleting Status = "Deleting"
	// Failed ...
	Failed Status = "Failed"
	// OK ...
	OK Status = "OK"
	// Succeeded ...
	Succeeded Status = "Succeeded"
	// Updated ...
	Updated Status = "Updated"
	// Updating ...
	Updating Status = "Updating"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Accepted, Canceled, Created, Creating, Deleted, Deleting, Failed, OK, Succeeded, Updated, Updating}
}

// OdataProductResult ...
type OdataProductResult struct {
	autorest.Response `json:"-"`
	Values            *[]Product `json:"values,omitempty"`
	OdataNextLink     *string    `json:"odata.nextLink,omitempty"`
}

// OdataProductResultIterator provides access to a complete listing of Product values.
type OdataProductResultIterator struct {
	i    int
	page OdataProductResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OdataProductResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fmt.Sprintf("%s/OdataProductResultIterator.NextWithContext", fqdn))
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OdataProductResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OdataProductResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OdataProductResultIterator) Response() OdataProductResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OdataProductResultIterator) Value() Product {
	if !iter.page.NotDone() {
		return Product{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (opr OdataProductResult) IsEmpty() bool {
	return opr.Values == nil || len(*opr.Values) == 0
}

// odataProductResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (opr OdataProductResult) odataProductResultPreparer(ctx context.Context) (*http.Request, error) {
	if opr.OdataNextLink == nil || len(to.String(opr.OdataNextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(opr.OdataNextLink)))
}

// OdataProductResultPage contains a page of Product values.
type OdataProductResultPage struct {
	fn  func(context.Context, OdataProductResult) (OdataProductResult, error)
	opr OdataProductResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OdataProductResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fmt.Sprintf("%s/OdataProductResultPage.NextWithContext", fqdn))
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.opr)
	if err != nil {
		return err
	}
	page.opr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OdataProductResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OdataProductResultPage) NotDone() bool {
	return !page.opr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OdataProductResultPage) Response() OdataProductResult {
	return page.opr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OdataProductResultPage) Values() []Product {
	if page.opr.IsEmpty() {
		return nil
	}
	return *page.opr.Values
}

// OperationResult ...
type OperationResult struct {
	// Status - The status of the request. Possible values include: 'Succeeded', 'Failed', 'Canceled', 'Accepted', 'Creating', 'Created', 'Updating', 'Updated', 'Deleting', 'Deleted', 'OK'
	Status Status `json:"status,omitempty"`
}

// PagingGetMultiplePagesLROAllFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type PagingGetMultiplePagesLROAllFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *PagingGetMultiplePagesLROAllFuture) Result(client PagingClient) (prp ProductResultPage, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "paginggroup.PagingGetMultiplePagesLROAllFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("paginggroup.PagingGetMultiplePagesLROAllFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if prp.pr.Response.Response, err = future.GetResult(sender); err == nil && prp.pr.Response.Response.StatusCode != http.StatusNoContent {
		prp, err = client.GetMultiplePagesLROResponder(prp.pr.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "paginggroup.PagingGetMultiplePagesLROAllFuture", "Result", prp.pr.Response.Response, "Failure responding to request")
		}
	}
	return
}

// PagingGetMultiplePagesLROFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type PagingGetMultiplePagesLROFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *PagingGetMultiplePagesLROFuture) Result(client PagingClient) (prp ProductResultPage, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "paginggroup.PagingGetMultiplePagesLROFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("paginggroup.PagingGetMultiplePagesLROFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if prp.pr.Response.Response, err = future.GetResult(sender); err == nil && prp.pr.Response.Response.StatusCode != http.StatusNoContent {
		prp, err = client.GetMultiplePagesLROResponder(prp.pr.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "paginggroup.PagingGetMultiplePagesLROFuture", "Result", prp.pr.Response.Response, "Failure responding to request")
		}
	}
	return
}

// Product ...
type Product struct {
	Properties *ProductProperties `json:"properties,omitempty"`
}

// ProductProperties ...
type ProductProperties struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ProductResult ...
type ProductResult struct {
	autorest.Response `json:"-"`
	Values            *[]Product `json:"values,omitempty"`
	NextLink          *string    `json:"nextLink,omitempty"`
}

// ProductResultIterator provides access to a complete listing of Product values.
type ProductResultIterator struct {
	i    int
	page ProductResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ProductResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fmt.Sprintf("%s/ProductResultIterator.NextWithContext", fqdn))
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ProductResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ProductResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ProductResultIterator) Response() ProductResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ProductResultIterator) Value() Product {
	if !iter.page.NotDone() {
		return Product{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (pr ProductResult) IsEmpty() bool {
	return pr.Values == nil || len(*pr.Values) == 0
}

// productResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (pr ProductResult) productResultPreparer(ctx context.Context) (*http.Request, error) {
	if pr.NextLink == nil || len(to.String(pr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(pr.NextLink)))
}

// ProductResultPage contains a page of Product values.
type ProductResultPage struct {
	fn func(context.Context, ProductResult) (ProductResult, error)
	pr ProductResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ProductResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fmt.Sprintf("%s/ProductResultPage.NextWithContext", fqdn))
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.pr)
	if err != nil {
		return err
	}
	page.pr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ProductResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ProductResultPage) NotDone() bool {
	return !page.pr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ProductResultPage) Response() ProductResult {
	return page.pr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ProductResultPage) Values() []Product {
	if page.pr.IsEmpty() {
		return nil
	}
	return *page.pr.Values
}
