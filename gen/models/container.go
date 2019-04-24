// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Container container
// swagger:model container
type Container struct {

	// The command to run for the docker image's entrypoint.
	Command *string `json:"command,omitempty"`

	// The docker image name for the container
	// Min Length: 1
	Image string `json:"image,omitempty"`

	// Image pull policy. One of Always or IfNotPresent.
	// Min Length: 1
	// Enum: [Always IfNotPresent]
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	// The docker image tag for the container
	// Min Length: 1
	ImageTag string `json:"imageTag,omitempty"`

	// The name of this container within the service
	// Min Length: 1
	Name string `json:"name,omitempty"`
}

// Validate validates this container
func (m *Container) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImagePullPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Container) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if err := validate.MinLength("image", "body", string(m.Image), 1); err != nil {
		return err
	}

	return nil
}

var containerTypeImagePullPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Always","IfNotPresent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerTypeImagePullPolicyPropEnum = append(containerTypeImagePullPolicyPropEnum, v)
	}
}

const (

	// ContainerImagePullPolicyAlways captures enum value "Always"
	ContainerImagePullPolicyAlways string = "Always"

	// ContainerImagePullPolicyIfNotPresent captures enum value "IfNotPresent"
	ContainerImagePullPolicyIfNotPresent string = "IfNotPresent"
)

// prop value enum
func (m *Container) validateImagePullPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, containerTypeImagePullPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Container) validateImagePullPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ImagePullPolicy) { // not required
		return nil
	}

	if err := validate.MinLength("imagePullPolicy", "body", string(m.ImagePullPolicy), 1); err != nil {
		return err
	}

	// value enum
	if err := m.validateImagePullPolicyEnum("imagePullPolicy", "body", m.ImagePullPolicy); err != nil {
		return err
	}

	return nil
}

func (m *Container) validateImageTag(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageTag) { // not required
		return nil
	}

	if err := validate.MinLength("imageTag", "body", string(m.ImageTag), 1); err != nil {
		return err
	}

	return nil
}

func (m *Container) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Container) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Container) UnmarshalBinary(b []byte) error {
	var res Container
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}