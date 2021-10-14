// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetControllersOKCode is the HTTP code returned for type GetControllersOK
const GetControllersOKCode int = 200

/*GetControllersOK Successful pull of controllers info

swagger:response getControllersOK
*/
type GetControllersOK struct {

	/*
	  In: Body
	*/
	Payload *GetControllersOKBody `json:"body,omitempty"`
}

// NewGetControllersOK creates GetControllersOK with default headers values
func NewGetControllersOK() *GetControllersOK {

	return &GetControllersOK{}
}

// WithPayload adds the payload to the get controllers o k response
func (o *GetControllersOK) WithPayload(payload *GetControllersOKBody) *GetControllersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get controllers o k response
func (o *GetControllersOK) SetPayload(payload *GetControllersOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetControllersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetControllersInternalServerErrorCode is the HTTP code returned for type GetControllersInternalServerError
const GetControllersInternalServerErrorCode int = 500

/*GetControllersInternalServerError Server error

swagger:response getControllersInternalServerError
*/
type GetControllersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetControllersInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetControllersInternalServerError creates GetControllersInternalServerError with default headers values
func NewGetControllersInternalServerError() *GetControllersInternalServerError {

	return &GetControllersInternalServerError{}
}

// WithPayload adds the payload to the get controllers internal server error response
func (o *GetControllersInternalServerError) WithPayload(payload *GetControllersInternalServerErrorBody) *GetControllersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get controllers internal server error response
func (o *GetControllersInternalServerError) SetPayload(payload *GetControllersInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetControllersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
