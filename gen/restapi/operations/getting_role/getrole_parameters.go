// Code generated by go-swagger; DO NOT EDIT.

package getting_role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetroleParams creates a new GetroleParams object
//
// There are no default values defined in the spec.
func NewGetroleParams() GetroleParams {

	return GetroleParams{}
}

// GetroleParams contains all the bound params for the getrole operation
// typically these are obtained from a http.Request
//
// swagger:parameters getrole
type GetroleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Enter apiName.
	  Required: true
	  In: path
	*/
	APIName string
	/*Enter desination.
	  Required: true
	  In: path
	*/
	Designation string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetroleParams() beforehand.
func (o *GetroleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAPIName, rhkAPIName, _ := route.Params.GetOK("apiName")
	if err := o.bindAPIName(rAPIName, rhkAPIName, route.Formats); err != nil {
		res = append(res, err)
	}

	rDesignation, rhkDesignation, _ := route.Params.GetOK("designation")
	if err := o.bindDesignation(rDesignation, rhkDesignation, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAPIName binds and validates parameter APIName from path.
func (o *GetroleParams) bindAPIName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.APIName = raw

	return nil
}

// bindDesignation binds and validates parameter Designation from path.
func (o *GetroleParams) bindDesignation(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Designation = raw

	return nil
}
