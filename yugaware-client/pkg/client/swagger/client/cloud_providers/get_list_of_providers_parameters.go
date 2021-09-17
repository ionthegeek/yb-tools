// Code generated by go-swagger; DO NOT EDIT.

package cloud_providers

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

// NewGetListOfProvidersParams creates a new GetListOfProvidersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetListOfProvidersParams() *GetListOfProvidersParams {
	return &GetListOfProvidersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetListOfProvidersParamsWithTimeout creates a new GetListOfProvidersParams object
// with the ability to set a timeout on a request.
func NewGetListOfProvidersParamsWithTimeout(timeout time.Duration) *GetListOfProvidersParams {
	return &GetListOfProvidersParams{
		timeout: timeout,
	}
}

// NewGetListOfProvidersParamsWithContext creates a new GetListOfProvidersParams object
// with the ability to set a context for a request.
func NewGetListOfProvidersParamsWithContext(ctx context.Context) *GetListOfProvidersParams {
	return &GetListOfProvidersParams{
		Context: ctx,
	}
}

// NewGetListOfProvidersParamsWithHTTPClient creates a new GetListOfProvidersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetListOfProvidersParamsWithHTTPClient(client *http.Client) *GetListOfProvidersParams {
	return &GetListOfProvidersParams{
		HTTPClient: client,
	}
}

/* GetListOfProvidersParams contains all the parameters to send to the API endpoint
   for the get list of providers operation.

   Typically these are written to a http.Request.
*/
type GetListOfProvidersParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get list of providers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetListOfProvidersParams) WithDefaults() *GetListOfProvidersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get list of providers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetListOfProvidersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get list of providers params
func (o *GetListOfProvidersParams) WithTimeout(timeout time.Duration) *GetListOfProvidersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get list of providers params
func (o *GetListOfProvidersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get list of providers params
func (o *GetListOfProvidersParams) WithContext(ctx context.Context) *GetListOfProvidersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get list of providers params
func (o *GetListOfProvidersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get list of providers params
func (o *GetListOfProvidersParams) WithHTTPClient(client *http.Client) *GetListOfProvidersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get list of providers params
func (o *GetListOfProvidersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the get list of providers params
func (o *GetListOfProvidersParams) WithCUUID(cUUID strfmt.UUID) *GetListOfProvidersParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the get list of providers params
func (o *GetListOfProvidersParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetListOfProvidersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}