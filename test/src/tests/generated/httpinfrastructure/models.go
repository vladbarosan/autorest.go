package httpinfrastructuregroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"net/http"
	"reflect"
	"strings"
)

// concatenates a slice of const values with the specified separator between each item
func joinConst(s interface{}, sep string) string {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic("s wasn't a slice or array")
	}
	ss := make([]string, 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		ss = append(ss, v.Index(i).String())
	}
	return strings.Join(ss, sep)
}

// A ...
type A struct {
	rawResponse *http.Response
	StatusCode  *string `json:"statusCode,omitempty"`
}

// Response returns the raw HTTP response object.
func (a A) Response() *http.Response {
	return a.rawResponse
}

// HTTPStatusCode returns the HTTP status code of the response, e.g. 200.
func (a A) HTTPStatusCode() int {
	return a.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (a A) Status() string {
	return a.rawResponse.Status
}

// B ...
type B struct {
	StatusCode     *string `json:"statusCode,omitempty"`
	TextStatusCode *string `json:"textStatusCode,omitempty"`
}

// C ...
type C struct {
	HTTPCode *string `json:"httpCode,omitempty"`
}

// D ...
type D struct {
	HTTPStatusCode *string `json:"httpStatusCode,omitempty"`
}

// Error ...
type Error struct {
	rawResponse *http.Response
	Status      *int32  `json:"status,omitempty"`
	Message     *string `json:"message,omitempty"`
}

// Response returns the raw HTTP response object.
func (eVar Error) Response() *http.Response {
	return eVar.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (eVar Error) StatusCode() int {
	return eVar.rawResponse.StatusCode
}

// HTTPStatus returns the HTTP status message of the response, e.g. "200 OK".
func (eVar Error) HTTPStatus() string {
	return eVar.rawResponse.Status
}

// Get200ModelA201ModelC404ModelDDefaultError200ValidResponse ...
type Get200ModelA201ModelC404ModelDDefaultError200ValidResponse struct {
	rawResponse *http.Response
	Value       map[string]interface{} `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (g2mamcmdde2vr Get200ModelA201ModelC404ModelDDefaultError200ValidResponse) Response() *http.Response {
	return g2mamcmdde2vr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (g2mamcmdde2vr Get200ModelA201ModelC404ModelDDefaultError200ValidResponse) StatusCode() int {
	return g2mamcmdde2vr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (g2mamcmdde2vr Get200ModelA201ModelC404ModelDDefaultError200ValidResponse) Status() string {
	return g2mamcmdde2vr.rawResponse.Status
}

// Get200ModelA201ModelC404ModelDDefaultError201ValidResponse ...
type Get200ModelA201ModelC404ModelDDefaultError201ValidResponse struct {
	rawResponse *http.Response
	Value       map[string]interface{} `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (g2mamcmdde2vr Get200ModelA201ModelC404ModelDDefaultError201ValidResponse) Response() *http.Response {
	return g2mamcmdde2vr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (g2mamcmdde2vr Get200ModelA201ModelC404ModelDDefaultError201ValidResponse) StatusCode() int {
	return g2mamcmdde2vr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (g2mamcmdde2vr Get200ModelA201ModelC404ModelDDefaultError201ValidResponse) Status() string {
	return g2mamcmdde2vr.rawResponse.Status
}

// Get200ModelA201ModelC404ModelDDefaultError400ValidResponse ...
type Get200ModelA201ModelC404ModelDDefaultError400ValidResponse struct {
	rawResponse *http.Response
	Value       map[string]interface{} `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (g2mamcmdde4vr Get200ModelA201ModelC404ModelDDefaultError400ValidResponse) Response() *http.Response {
	return g2mamcmdde4vr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (g2mamcmdde4vr Get200ModelA201ModelC404ModelDDefaultError400ValidResponse) StatusCode() int {
	return g2mamcmdde4vr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (g2mamcmdde4vr Get200ModelA201ModelC404ModelDDefaultError400ValidResponse) Status() string {
	return g2mamcmdde4vr.rawResponse.Status
}

// Get200ModelA201ModelC404ModelDDefaultError404ValidResponse ...
type Get200ModelA201ModelC404ModelDDefaultError404ValidResponse struct {
	rawResponse *http.Response
	Value       map[string]interface{} `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (g2mamcmdde4vr Get200ModelA201ModelC404ModelDDefaultError404ValidResponse) Response() *http.Response {
	return g2mamcmdde4vr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (g2mamcmdde4vr Get200ModelA201ModelC404ModelDDefaultError404ValidResponse) StatusCode() int {
	return g2mamcmdde4vr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (g2mamcmdde4vr Get200ModelA201ModelC404ModelDDefaultError404ValidResponse) Status() string {
	return g2mamcmdde4vr.rawResponse.Status
}

// Get200Response ...
type Get200Response struct {
	rawResponse *http.Response
	Value       *bool `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (g2r Get200Response) Response() *http.Response {
	return g2r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (g2r Get200Response) StatusCode() int {
	return g2r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (g2r Get200Response) Status() string {
	return g2r.rawResponse.Status
}

// Get300Response ...
type Get300Response struct {
	rawResponse *http.Response
	Value       []string `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (g3r Get300Response) Response() *http.Response {
	return g3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (g3r Get300Response) StatusCode() int {
	return g3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (g3r Get300Response) Status() string {
	return g3r.rawResponse.Status
}

// Location returns the value for header Location.
func (g3r Get300Response) Location() string {
	return string(g3r.rawResponse.Header.Get("Location"))
}

// GetEmptyErrorResponse ...
type GetEmptyErrorResponse struct {
	rawResponse *http.Response
	Value       *bool `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (geer GetEmptyErrorResponse) Response() *http.Response {
	return geer.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (geer GetEmptyErrorResponse) StatusCode() int {
	return geer.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (geer GetEmptyErrorResponse) Status() string {
	return geer.rawResponse.Status
}

// GetNoModelEmptyResponse ...
type GetNoModelEmptyResponse struct {
	rawResponse *http.Response
	Value       *bool `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (gnmer GetNoModelEmptyResponse) Response() *http.Response {
	return gnmer.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (gnmer GetNoModelEmptyResponse) StatusCode() int {
	return gnmer.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (gnmer GetNoModelEmptyResponse) Status() string {
	return gnmer.rawResponse.Status
}

// GetNoModelErrorResponse ...
type GetNoModelErrorResponse struct {
	rawResponse *http.Response
	Value       *bool `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (gnmer GetNoModelErrorResponse) Response() *http.Response {
	return gnmer.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (gnmer GetNoModelErrorResponse) StatusCode() int {
	return gnmer.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (gnmer GetNoModelErrorResponse) Status() string {
	return gnmer.rawResponse.Status
}

// HTTPRedirectsDelete307Response ...
type HTTPRedirectsDelete307Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrd3r HTTPRedirectsDelete307Response) Response() *http.Response {
	return hrd3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrd3r HTTPRedirectsDelete307Response) StatusCode() int {
	return hrd3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrd3r HTTPRedirectsDelete307Response) Status() string {
	return hrd3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrd3r HTTPRedirectsDelete307Response) Location() string {
	return string(hrd3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsGet301Response ...
type HTTPRedirectsGet301Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrg3r HTTPRedirectsGet301Response) Response() *http.Response {
	return hrg3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrg3r HTTPRedirectsGet301Response) StatusCode() int {
	return hrg3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrg3r HTTPRedirectsGet301Response) Status() string {
	return hrg3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrg3r HTTPRedirectsGet301Response) Location() string {
	return string(hrg3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsGet302Response ...
type HTTPRedirectsGet302Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrg3r HTTPRedirectsGet302Response) Response() *http.Response {
	return hrg3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrg3r HTTPRedirectsGet302Response) StatusCode() int {
	return hrg3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrg3r HTTPRedirectsGet302Response) Status() string {
	return hrg3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrg3r HTTPRedirectsGet302Response) Location() string {
	return string(hrg3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsGet307Response ...
type HTTPRedirectsGet307Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrg3r HTTPRedirectsGet307Response) Response() *http.Response {
	return hrg3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrg3r HTTPRedirectsGet307Response) StatusCode() int {
	return hrg3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrg3r HTTPRedirectsGet307Response) Status() string {
	return hrg3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrg3r HTTPRedirectsGet307Response) Location() string {
	return string(hrg3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsHead300Response ...
type HTTPRedirectsHead300Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrh3r HTTPRedirectsHead300Response) Response() *http.Response {
	return hrh3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrh3r HTTPRedirectsHead300Response) StatusCode() int {
	return hrh3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrh3r HTTPRedirectsHead300Response) Status() string {
	return hrh3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrh3r HTTPRedirectsHead300Response) Location() string {
	return string(hrh3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsHead301Response ...
type HTTPRedirectsHead301Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrh3r HTTPRedirectsHead301Response) Response() *http.Response {
	return hrh3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrh3r HTTPRedirectsHead301Response) StatusCode() int {
	return hrh3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrh3r HTTPRedirectsHead301Response) Status() string {
	return hrh3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrh3r HTTPRedirectsHead301Response) Location() string {
	return string(hrh3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsHead302Response ...
type HTTPRedirectsHead302Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrh3r HTTPRedirectsHead302Response) Response() *http.Response {
	return hrh3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrh3r HTTPRedirectsHead302Response) StatusCode() int {
	return hrh3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrh3r HTTPRedirectsHead302Response) Status() string {
	return hrh3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrh3r HTTPRedirectsHead302Response) Location() string {
	return string(hrh3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsHead307Response ...
type HTTPRedirectsHead307Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrh3r HTTPRedirectsHead307Response) Response() *http.Response {
	return hrh3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrh3r HTTPRedirectsHead307Response) StatusCode() int {
	return hrh3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrh3r HTTPRedirectsHead307Response) Status() string {
	return hrh3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrh3r HTTPRedirectsHead307Response) Location() string {
	return string(hrh3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsPatch302Response ...
type HTTPRedirectsPatch302Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrp3r HTTPRedirectsPatch302Response) Response() *http.Response {
	return hrp3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrp3r HTTPRedirectsPatch302Response) StatusCode() int {
	return hrp3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrp3r HTTPRedirectsPatch302Response) Status() string {
	return hrp3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrp3r HTTPRedirectsPatch302Response) Location() string {
	return string(hrp3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsPatch307Response ...
type HTTPRedirectsPatch307Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrp3r HTTPRedirectsPatch307Response) Response() *http.Response {
	return hrp3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrp3r HTTPRedirectsPatch307Response) StatusCode() int {
	return hrp3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrp3r HTTPRedirectsPatch307Response) Status() string {
	return hrp3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrp3r HTTPRedirectsPatch307Response) Location() string {
	return string(hrp3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsPost303Response ...
type HTTPRedirectsPost303Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrp3r HTTPRedirectsPost303Response) Response() *http.Response {
	return hrp3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrp3r HTTPRedirectsPost303Response) StatusCode() int {
	return hrp3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrp3r HTTPRedirectsPost303Response) Status() string {
	return hrp3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrp3r HTTPRedirectsPost303Response) Location() string {
	return string(hrp3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsPost307Response ...
type HTTPRedirectsPost307Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrp3r HTTPRedirectsPost307Response) Response() *http.Response {
	return hrp3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrp3r HTTPRedirectsPost307Response) StatusCode() int {
	return hrp3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrp3r HTTPRedirectsPost307Response) Status() string {
	return hrp3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrp3r HTTPRedirectsPost307Response) Location() string {
	return string(hrp3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsPut301Response ...
type HTTPRedirectsPut301Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrp3r HTTPRedirectsPut301Response) Response() *http.Response {
	return hrp3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrp3r HTTPRedirectsPut301Response) StatusCode() int {
	return hrp3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrp3r HTTPRedirectsPut301Response) Status() string {
	return hrp3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrp3r HTTPRedirectsPut301Response) Location() string {
	return string(hrp3r.rawResponse.Header.Get("Location"))
}

// HTTPRedirectsPut307Response ...
type HTTPRedirectsPut307Response struct {
	rawResponse *http.Response
}

// Response returns the raw HTTP response object.
func (hrp3r HTTPRedirectsPut307Response) Response() *http.Response {
	return hrp3r.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (hrp3r HTTPRedirectsPut307Response) StatusCode() int {
	return hrp3r.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (hrp3r HTTPRedirectsPut307Response) Status() string {
	return hrp3r.rawResponse.Status
}

// Location returns the value for header Location.
func (hrp3r HTTPRedirectsPut307Response) Location() string {
	return string(hrp3r.rawResponse.Header.Get("Location"))
}
