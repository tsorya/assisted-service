// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OpenshiftVersion Version of the OpenShift cluster.
//
// swagger:model openshift-version
type OpenshiftVersion string

const (

	// OpenshiftVersionNr45 captures enum value "4.5"
	OpenshiftVersionNr45 OpenshiftVersion = "4.5"

	// OpenshiftVersionNr46 captures enum value "4.6"
	OpenshiftVersionNr46 OpenshiftVersion = "4.6"
)

// for schema
var openshiftVersionEnum []interface{}

func init() {
	var res []OpenshiftVersion
	if err := json.Unmarshal([]byte(`["4.5","4.6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openshiftVersionEnum = append(openshiftVersionEnum, v)
	}
}

func (m OpenshiftVersion) validateOpenshiftVersionEnum(path, location string, value OpenshiftVersion) error {
	if err := validate.EnumCase(path, location, value, openshiftVersionEnum, false); err != nil {
		return err
	}
	return nil
}

// Validate validates this openshift version
func (m OpenshiftVersion) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenshiftVersionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
