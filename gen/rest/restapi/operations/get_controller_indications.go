// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetControllerIndicationsHandlerFunc turns a function with the right signature into a get controller indications handler
type GetControllerIndicationsHandlerFunc func(GetControllerIndicationsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetControllerIndicationsHandlerFunc) Handle(params GetControllerIndicationsParams) middleware.Responder {
	return fn(params)
}

// GetControllerIndicationsHandler interface for that can handle valid get controller indications params
type GetControllerIndicationsHandler interface {
	Handle(GetControllerIndicationsParams) middleware.Responder
}

// NewGetControllerIndications creates a new http.Handler for the get controller indications operation
func NewGetControllerIndications(ctx *middleware.Context, handler GetControllerIndicationsHandler) *GetControllerIndications {
	return &GetControllerIndications{Context: ctx, Handler: handler}
}

/* GetControllerIndications swagger:route GET /controller/indications getControllerIndications

List of controller's indications

Get list of controller's indications

*/
type GetControllerIndications struct {
	Context *middleware.Context
	Handler GetControllerIndicationsHandler
}

func (o *GetControllerIndications) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetControllerIndicationsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetControllerIndicationsBadRequestBody get controller indications bad request body
//
// swagger:model GetControllerIndicationsBadRequestBody
type GetControllerIndicationsBadRequestBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this get controller indications bad request body
func (o *GetControllerIndicationsBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get controller indications bad request body based on context it is used
func (o *GetControllerIndicationsBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetControllerIndicationsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetControllerIndicationsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetControllerIndicationsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetControllerIndicationsBody get controller indications body
//
// swagger:model GetControllerIndicationsBody
type GetControllerIndicationsBody struct {

	// serial
	Serial string `json:"serial,omitempty"`
}

// Validate validates this get controller indications body
func (o *GetControllerIndicationsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get controller indications body based on context it is used
func (o *GetControllerIndicationsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetControllerIndicationsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetControllerIndicationsBody) UnmarshalBinary(b []byte) error {
	var res GetControllerIndicationsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetControllerIndicationsInternalServerErrorBody get controller indications internal server error body
//
// swagger:model GetControllerIndicationsInternalServerErrorBody
type GetControllerIndicationsInternalServerErrorBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this get controller indications internal server error body
func (o *GetControllerIndicationsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get controller indications internal server error body based on context it is used
func (o *GetControllerIndicationsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetControllerIndicationsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetControllerIndicationsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetControllerIndicationsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetControllerIndicationsOKBody get controller indications o k body
//
// swagger:model GetControllerIndicationsOKBody
type GetControllerIndicationsOKBody struct {

	// indications
	Indications []*GetControllerIndicationsOKBodyIndicationsItems0 `json:"indications"`
}

// Validate validates this get controller indications o k body
func (o *GetControllerIndicationsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIndications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetControllerIndicationsOKBody) validateIndications(formats strfmt.Registry) error {
	if swag.IsZero(o.Indications) { // not required
		return nil
	}

	for i := 0; i < len(o.Indications); i++ {
		if swag.IsZero(o.Indications[i]) { // not required
			continue
		}

		if o.Indications[i] != nil {
			if err := o.Indications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getControllerIndicationsOK" + "." + "indications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getControllerIndicationsOK" + "." + "indications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get controller indications o k body based on the context it is used
func (o *GetControllerIndicationsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIndications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetControllerIndicationsOKBody) contextValidateIndications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Indications); i++ {

		if o.Indications[i] != nil {
			if err := o.Indications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getControllerIndicationsOK" + "." + "indications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getControllerIndicationsOK" + "." + "indications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetControllerIndicationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetControllerIndicationsOKBody) UnmarshalBinary(b []byte) error {
	var res GetControllerIndicationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetControllerIndicationsOKBodyIndicationsItems0 get controller indications o k body indications items0
//
// swagger:model GetControllerIndicationsOKBodyIndicationsItems0
type GetControllerIndicationsOKBodyIndicationsItems0 struct {

	// indication
	Indication float32 `json:"indication,omitempty"`

	// sent at
	// Format: date-time
	SentAt strfmt.DateTime `json:"sent_at,omitempty"`
}

// Validate validates this get controller indications o k body indications items0
func (o *GetControllerIndicationsOKBodyIndicationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSentAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetControllerIndicationsOKBodyIndicationsItems0) validateSentAt(formats strfmt.Registry) error {
	if swag.IsZero(o.SentAt) { // not required
		return nil
	}

	if err := validate.FormatOf("sent_at", "body", "date-time", o.SentAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get controller indications o k body indications items0 based on context it is used
func (o *GetControllerIndicationsOKBodyIndicationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetControllerIndicationsOKBodyIndicationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetControllerIndicationsOKBodyIndicationsItems0) UnmarshalBinary(b []byte) error {
	var res GetControllerIndicationsOKBodyIndicationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
