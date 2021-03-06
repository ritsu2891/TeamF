// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPostTagParams creates a new PostTagParams object
// no default values defined in spec.
func NewPostTagParams() PostTagParams {

	return PostTagParams{}
}

// PostTagParams contains all the bound params for the post tag operation
// typically these are obtained from a http.Request
//
// swagger:parameters postTag
type PostTagParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Tag PostTagBody
	/*
	  Required: true
	  In: header
	*/
	Token string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostTagParams() beforehand.
func (o *PostTagParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PostTagBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("tag", "body", ""))
			} else {
				res = append(res, errors.NewParseError("tag", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Tag = body
			}
		}
	} else {
		res = append(res, errors.Required("tag", "body", ""))
	}
	if err := o.bindToken(r.Header[http.CanonicalHeaderKey("token")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindToken binds and validates parameter Token from header.
func (o *PostTagParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("token", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("token", "header", raw); err != nil {
		return err
	}

	o.Token = raw

	return nil
}
