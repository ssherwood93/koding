package j_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJUserVerifyPasswordParams creates a new PostRemoteAPIJUserVerifyPasswordParams object
// with the default values initialized.
func NewPostRemoteAPIJUserVerifyPasswordParams() *PostRemoteAPIJUserVerifyPasswordParams {
	var ()
	return &PostRemoteAPIJUserVerifyPasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJUserVerifyPasswordParamsWithTimeout creates a new PostRemoteAPIJUserVerifyPasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJUserVerifyPasswordParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJUserVerifyPasswordParams {
	var ()
	return &PostRemoteAPIJUserVerifyPasswordParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJUserVerifyPasswordParamsWithContext creates a new PostRemoteAPIJUserVerifyPasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJUserVerifyPasswordParamsWithContext(ctx context.Context) *PostRemoteAPIJUserVerifyPasswordParams {
	var ()
	return &PostRemoteAPIJUserVerifyPasswordParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJUserVerifyPasswordParams contains all the parameters to send to the API endpoint
for the post remote API j user verify password operation typically these are written to a http.Request
*/
type PostRemoteAPIJUserVerifyPasswordParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j user verify password params
func (o *PostRemoteAPIJUserVerifyPasswordParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJUserVerifyPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j user verify password params
func (o *PostRemoteAPIJUserVerifyPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j user verify password params
func (o *PostRemoteAPIJUserVerifyPasswordParams) WithContext(ctx context.Context) *PostRemoteAPIJUserVerifyPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j user verify password params
func (o *PostRemoteAPIJUserVerifyPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j user verify password params
func (o *PostRemoteAPIJUserVerifyPasswordParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJUserVerifyPasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j user verify password params
func (o *PostRemoteAPIJUserVerifyPasswordParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJUserVerifyPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
