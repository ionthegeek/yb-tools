// Code generated by go-swagger; DO NOT EDIT.

package schedule_external_script

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

// NewExternalScriptScheduleParams creates a new ExternalScriptScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExternalScriptScheduleParams() *ExternalScriptScheduleParams {
	return &ExternalScriptScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExternalScriptScheduleParamsWithTimeout creates a new ExternalScriptScheduleParams object
// with the ability to set a timeout on a request.
func NewExternalScriptScheduleParamsWithTimeout(timeout time.Duration) *ExternalScriptScheduleParams {
	return &ExternalScriptScheduleParams{
		timeout: timeout,
	}
}

// NewExternalScriptScheduleParamsWithContext creates a new ExternalScriptScheduleParams object
// with the ability to set a context for a request.
func NewExternalScriptScheduleParamsWithContext(ctx context.Context) *ExternalScriptScheduleParams {
	return &ExternalScriptScheduleParams{
		Context: ctx,
	}
}

// NewExternalScriptScheduleParamsWithHTTPClient creates a new ExternalScriptScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewExternalScriptScheduleParamsWithHTTPClient(client *http.Client) *ExternalScriptScheduleParams {
	return &ExternalScriptScheduleParams{
		HTTPClient: client,
	}
}

/* ExternalScriptScheduleParams contains all the parameters to send to the API endpoint
   for the external script schedule operation.

   Typically these are written to a http.Request.
*/
type ExternalScriptScheduleParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// UniUUID.
	//
	// Format: uuid
	UniUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the external script schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalScriptScheduleParams) WithDefaults() *ExternalScriptScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the external script schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalScriptScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the external script schedule params
func (o *ExternalScriptScheduleParams) WithTimeout(timeout time.Duration) *ExternalScriptScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the external script schedule params
func (o *ExternalScriptScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the external script schedule params
func (o *ExternalScriptScheduleParams) WithContext(ctx context.Context) *ExternalScriptScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the external script schedule params
func (o *ExternalScriptScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the external script schedule params
func (o *ExternalScriptScheduleParams) WithHTTPClient(client *http.Client) *ExternalScriptScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the external script schedule params
func (o *ExternalScriptScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the external script schedule params
func (o *ExternalScriptScheduleParams) WithCUUID(cUUID strfmt.UUID) *ExternalScriptScheduleParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the external script schedule params
func (o *ExternalScriptScheduleParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithUniUUID adds the uniUUID to the external script schedule params
func (o *ExternalScriptScheduleParams) WithUniUUID(uniUUID strfmt.UUID) *ExternalScriptScheduleParams {
	o.SetUniUUID(uniUUID)
	return o
}

// SetUniUUID adds the uniUuid to the external script schedule params
func (o *ExternalScriptScheduleParams) SetUniUUID(uniUUID strfmt.UUID) {
	o.UniUUID = uniUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ExternalScriptScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param uniUUID
	if err := r.SetPathParam("uniUUID", o.UniUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}