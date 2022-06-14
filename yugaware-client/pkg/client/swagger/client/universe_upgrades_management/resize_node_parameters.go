// Code generated by go-swagger; DO NOT EDIT.

package universe_upgrades_management

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

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// NewResizeNodeParams creates a new ResizeNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResizeNodeParams() *ResizeNodeParams {
	return &ResizeNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResizeNodeParamsWithTimeout creates a new ResizeNodeParams object
// with the ability to set a timeout on a request.
func NewResizeNodeParamsWithTimeout(timeout time.Duration) *ResizeNodeParams {
	return &ResizeNodeParams{
		timeout: timeout,
	}
}

// NewResizeNodeParamsWithContext creates a new ResizeNodeParams object
// with the ability to set a context for a request.
func NewResizeNodeParamsWithContext(ctx context.Context) *ResizeNodeParams {
	return &ResizeNodeParams{
		Context: ctx,
	}
}

// NewResizeNodeParamsWithHTTPClient creates a new ResizeNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewResizeNodeParamsWithHTTPClient(client *http.Client) *ResizeNodeParams {
	return &ResizeNodeParams{
		HTTPClient: client,
	}
}

/* ResizeNodeParams contains all the parameters to send to the API endpoint
   for the resize node operation.

   Typically these are written to a http.Request.
*/
type ResizeNodeParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	/* ResizeNodeParams.

	   Resize Node Params
	*/
	ResizeNodeParams *models.ResizeNodeParams

	// UniUUID.
	//
	// Format: uuid
	UniUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resize node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResizeNodeParams) WithDefaults() *ResizeNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resize node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResizeNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resize node params
func (o *ResizeNodeParams) WithTimeout(timeout time.Duration) *ResizeNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resize node params
func (o *ResizeNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resize node params
func (o *ResizeNodeParams) WithContext(ctx context.Context) *ResizeNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resize node params
func (o *ResizeNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resize node params
func (o *ResizeNodeParams) WithHTTPClient(client *http.Client) *ResizeNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resize node params
func (o *ResizeNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the resize node params
func (o *ResizeNodeParams) WithCUUID(cUUID strfmt.UUID) *ResizeNodeParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the resize node params
func (o *ResizeNodeParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithResizeNodeParams adds the resizeNodeParams to the resize node params
func (o *ResizeNodeParams) WithResizeNodeParams(resizeNodeParams *models.ResizeNodeParams) *ResizeNodeParams {
	o.SetResizeNodeParams(resizeNodeParams)
	return o
}

// SetResizeNodeParams adds the resizeNodeParams to the resize node params
func (o *ResizeNodeParams) SetResizeNodeParams(resizeNodeParams *models.ResizeNodeParams) {
	o.ResizeNodeParams = resizeNodeParams
}

// WithUniUUID adds the uniUUID to the resize node params
func (o *ResizeNodeParams) WithUniUUID(uniUUID strfmt.UUID) *ResizeNodeParams {
	o.SetUniUUID(uniUUID)
	return o
}

// SetUniUUID adds the uniUuid to the resize node params
func (o *ResizeNodeParams) SetUniUUID(uniUUID strfmt.UUID) {
	o.UniUUID = uniUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ResizeNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}
	if o.ResizeNodeParams != nil {
		if err := r.SetBodyParam(o.ResizeNodeParams); err != nil {
			return err
		}
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