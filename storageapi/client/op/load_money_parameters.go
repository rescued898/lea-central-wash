// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLoadMoneyParams creates a new LoadMoneyParams object
// with the default values initialized.
func NewLoadMoneyParams() *LoadMoneyParams {
	var ()
	return &LoadMoneyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadMoneyParamsWithTimeout creates a new LoadMoneyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadMoneyParamsWithTimeout(timeout time.Duration) *LoadMoneyParams {
	var ()
	return &LoadMoneyParams{

		timeout: timeout,
	}
}

// NewLoadMoneyParamsWithContext creates a new LoadMoneyParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadMoneyParamsWithContext(ctx context.Context) *LoadMoneyParams {
	var ()
	return &LoadMoneyParams{

		Context: ctx,
	}
}

// NewLoadMoneyParamsWithHTTPClient creates a new LoadMoneyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadMoneyParamsWithHTTPClient(client *http.Client) *LoadMoneyParams {
	var ()
	return &LoadMoneyParams{
		HTTPClient: client,
	}
}

/*LoadMoneyParams contains all the parameters to send to the API endpoint
for the load money operation typically these are written to a http.Request
*/
type LoadMoneyParams struct {

	/*Args*/
	Args LoadMoneyBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load money params
func (o *LoadMoneyParams) WithTimeout(timeout time.Duration) *LoadMoneyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load money params
func (o *LoadMoneyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load money params
func (o *LoadMoneyParams) WithContext(ctx context.Context) *LoadMoneyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load money params
func (o *LoadMoneyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load money params
func (o *LoadMoneyParams) WithHTTPClient(client *http.Client) *LoadMoneyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load money params
func (o *LoadMoneyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the load money params
func (o *LoadMoneyParams) WithArgs(args LoadMoneyBody) *LoadMoneyParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the load money params
func (o *LoadMoneyParams) SetArgs(args LoadMoneyBody) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *LoadMoneyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Args); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}