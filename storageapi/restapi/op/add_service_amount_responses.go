// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
)

// AddServiceAmountNoContentCode is the HTTP code returned for type AddServiceAmountNoContent
const AddServiceAmountNoContentCode int = 204

/*AddServiceAmountNoContent OK

swagger:response addServiceAmountNoContent
*/
type AddServiceAmountNoContent struct {
}

// NewAddServiceAmountNoContent creates AddServiceAmountNoContent with default headers values
func NewAddServiceAmountNoContent() *AddServiceAmountNoContent {

	return &AddServiceAmountNoContent{}
}

// WriteResponse to the client
func (o *AddServiceAmountNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *AddServiceAmountNoContent) AddServiceAmountResponder() {}

// AddServiceAmountNotFoundCode is the HTTP code returned for type AddServiceAmountNotFound
const AddServiceAmountNotFoundCode int = 404

/*AddServiceAmountNotFound not found

swagger:response addServiceAmountNotFound
*/
type AddServiceAmountNotFound struct {
}

// NewAddServiceAmountNotFound creates AddServiceAmountNotFound with default headers values
func NewAddServiceAmountNotFound() *AddServiceAmountNotFound {

	return &AddServiceAmountNotFound{}
}

// WriteResponse to the client
func (o *AddServiceAmountNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *AddServiceAmountNotFound) AddServiceAmountResponder() {}

// AddServiceAmountInternalServerErrorCode is the HTTP code returned for type AddServiceAmountInternalServerError
const AddServiceAmountInternalServerErrorCode int = 500

/*AddServiceAmountInternalServerError internal error

swagger:response addServiceAmountInternalServerError
*/
type AddServiceAmountInternalServerError struct {
}

// NewAddServiceAmountInternalServerError creates AddServiceAmountInternalServerError with default headers values
func NewAddServiceAmountInternalServerError() *AddServiceAmountInternalServerError {

	return &AddServiceAmountInternalServerError{}
}

// WriteResponse to the client
func (o *AddServiceAmountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *AddServiceAmountInternalServerError) AddServiceAmountResponder() {}

type AddServiceAmountNotImplementedResponder struct {
	middleware.Responder
}

func (*AddServiceAmountNotImplementedResponder) AddServiceAmountResponder() {}

func AddServiceAmountNotImplemented() AddServiceAmountResponder {
	return &AddServiceAmountNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.AddServiceAmount has not yet been implemented",
		),
	}
}

type AddServiceAmountResponder interface {
	middleware.Responder
	AddServiceAmountResponder()
}
