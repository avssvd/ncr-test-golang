// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostControllerOKCode is the HTTP code returned for type PostControllerOK
const PostControllerOKCode int = 200

/*PostControllerOK Status

swagger:response postControllerOK
*/
type PostControllerOK struct {

	/*
	  In: Body
	*/
	Payload *PostControllerOKBody `json:"body,omitempty"`
}

// NewPostControllerOK creates PostControllerOK with default headers values
func NewPostControllerOK() *PostControllerOK {

	return &PostControllerOK{}
}

// WithPayload adds the payload to the post controller o k response
func (o *PostControllerOK) WithPayload(payload *PostControllerOKBody) *PostControllerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post controller o k response
func (o *PostControllerOK) SetPayload(payload *PostControllerOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostControllerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostControllerBadRequestCode is the HTTP code returned for type PostControllerBadRequest
const PostControllerBadRequestCode int = 400

/*PostControllerBadRequest Controller already exists

swagger:response postControllerBadRequest
*/
type PostControllerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostControllerBadRequestBody `json:"body,omitempty"`
}

// NewPostControllerBadRequest creates PostControllerBadRequest with default headers values
func NewPostControllerBadRequest() *PostControllerBadRequest {

	return &PostControllerBadRequest{}
}

// WithPayload adds the payload to the post controller bad request response
func (o *PostControllerBadRequest) WithPayload(payload *PostControllerBadRequestBody) *PostControllerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post controller bad request response
func (o *PostControllerBadRequest) SetPayload(payload *PostControllerBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostControllerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostControllerInternalServerErrorCode is the HTTP code returned for type PostControllerInternalServerError
const PostControllerInternalServerErrorCode int = 500

/*PostControllerInternalServerError Server error

swagger:response postControllerInternalServerError
*/
type PostControllerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostControllerInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostControllerInternalServerError creates PostControllerInternalServerError with default headers values
func NewPostControllerInternalServerError() *PostControllerInternalServerError {

	return &PostControllerInternalServerError{}
}

// WithPayload adds the payload to the post controller internal server error response
func (o *PostControllerInternalServerError) WithPayload(payload *PostControllerInternalServerErrorBody) *PostControllerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post controller internal server error response
func (o *PostControllerInternalServerError) SetPayload(payload *PostControllerInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostControllerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
