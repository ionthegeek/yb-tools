// Code generated by go-swagger; DO NOT EDIT.

package session_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCustomerCountParams creates a new CustomerCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerCountParams() *CustomerCountParams {
	return &CustomerCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerCountParamsWithTimeout creates a new CustomerCountParams object
// with the ability to set a timeout on a request.
func NewCustomerCountParamsWithTimeout(timeout time.Duration) *CustomerCountParams {
	return &CustomerCountParams{
		timeout: timeout,
	}
}

// NewCustomerCountParamsWithContext creates a new CustomerCountParams object
// with the ability to set a context for a request.
func NewCustomerCountParamsWithContext(ctx context.Context) *CustomerCountParams {
	return &CustomerCountParams{
		Context: ctx,
	}
}

// NewCustomerCountParamsWithHTTPClient creates a new CustomerCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerCountParamsWithHTTPClient(client *http.Client) *CustomerCountParams {
	return &CustomerCountParams{
		HTTPClient: client,
	}
}

/* CustomerCountParams contains all the parameters to send to the API endpoint
   for the customer count operation.

   Typically these are written to a http.Request.
*/
type CustomerCountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerCountParams) WithDefaults() *CustomerCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer count params
func (o *CustomerCountParams) WithTimeout(timeout time.Duration) *CustomerCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer count params
func (o *CustomerCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer count params
func (o *CustomerCountParams) WithContext(ctx context.Context) *CustomerCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer count params
func (o *CustomerCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer count params
func (o *CustomerCountParams) WithHTTPClient(client *http.Client) *CustomerCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer count params
func (o *CustomerCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}