// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostControllerHandlerFunc turns a function with the right signature into a post controller handler
type PostControllerHandlerFunc func(PostControllerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostControllerHandlerFunc) Handle(params PostControllerParams) middleware.Responder {
	return fn(params)
}

// PostControllerHandler interface for that can handle valid post controller params
type PostControllerHandler interface {
	Handle(PostControllerParams) middleware.Responder
}

// NewPostController creates a new http.Handler for the post controller operation
func NewPostController(ctx *middleware.Context, handler PostControllerHandler) *PostController {
	return &PostController{Context: ctx, Handler: handler}
}

/* PostController swagger:route POST /controller postController

Add controller

Add controller in DB

*/
type PostController struct {
	Context *middleware.Context
	Handler PostControllerHandler
}

func (o *PostController) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostControllerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostControllerBadRequestBody post controller bad request body
//
// swagger:model PostControllerBadRequestBody
type PostControllerBadRequestBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this post controller bad request body
func (o *PostControllerBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post controller bad request body based on context it is used
func (o *PostControllerBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostControllerBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostControllerBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostControllerBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostControllerBody post controller body
//
// swagger:model PostControllerBody
type PostControllerBody struct {

	// serial
	Serial string `json:"serial,omitempty"`
}

// Validate validates this post controller body
func (o *PostControllerBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post controller body based on context it is used
func (o *PostControllerBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostControllerBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostControllerBody) UnmarshalBinary(b []byte) error {
	var res PostControllerBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostControllerInternalServerErrorBody post controller internal server error body
//
// swagger:model PostControllerInternalServerErrorBody
type PostControllerInternalServerErrorBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this post controller internal server error body
func (o *PostControllerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post controller internal server error body based on context it is used
func (o *PostControllerInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostControllerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostControllerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostControllerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostControllerOKBody post controller o k body
//
// swagger:model PostControllerOKBody
type PostControllerOKBody struct {

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this post controller o k body
func (o *PostControllerOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post controller o k body based on context it is used
func (o *PostControllerOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostControllerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostControllerOKBody) UnmarshalBinary(b []byte) error {
	var res PostControllerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
