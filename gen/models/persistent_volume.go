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

// PersistentVolume persistent volume
// swagger:model persistentVolume
type PersistentVolume struct {

	// The desired access mode for the volume
	// Required: true
	// Min Length: 1
	// Enum: [ReadWriteOnce ReadOnlyMany ReadWriteMany]
	AccessMode string `json:"accessMode"`

	// The desired size of the volume in GB
	// Required: true
	Capacity int64 `json:"capacity"`

	// The name of the volume
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// The desired storage class for the volume
	// Required: true
	// Min Length: 1
	StorageClassName string `json:"storageClassName"`
}

// Validate validates this persistent volume
func (m *PersistentVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageClassName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var persistentVolumeTypeAccessModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ReadWriteOnce","ReadOnlyMany","ReadWriteMany"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		persistentVolumeTypeAccessModePropEnum = append(persistentVolumeTypeAccessModePropEnum, v)
	}
}

const (

	// PersistentVolumeAccessModeReadWriteOnce captures enum value "ReadWriteOnce"
	PersistentVolumeAccessModeReadWriteOnce string = "ReadWriteOnce"

	// PersistentVolumeAccessModeReadOnlyMany captures enum value "ReadOnlyMany"
	PersistentVolumeAccessModeReadOnlyMany string = "ReadOnlyMany"

	// PersistentVolumeAccessModeReadWriteMany captures enum value "ReadWriteMany"
	PersistentVolumeAccessModeReadWriteMany string = "ReadWriteMany"
)

// prop value enum
func (m *PersistentVolume) validateAccessModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, persistentVolumeTypeAccessModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PersistentVolume) validateAccessMode(formats strfmt.Registry) error {

	if err := validate.RequiredString("accessMode", "body", string(m.AccessMode)); err != nil {
		return err
	}

	if err := validate.MinLength("accessMode", "body", string(m.AccessMode), 1); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccessModeEnum("accessMode", "body", m.AccessMode); err != nil {
		return err
	}

	return nil
}

func (m *PersistentVolume) validateCapacity(formats strfmt.Registry) error {

	if err := validate.Required("capacity", "body", int64(m.Capacity)); err != nil {
		return err
	}

	return nil
}

func (m *PersistentVolume) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *PersistentVolume) validateStorageClassName(formats strfmt.Registry) error {

	if err := validate.RequiredString("storageClassName", "body", string(m.StorageClassName)); err != nil {
		return err
	}

	if err := validate.MinLength("storageClassName", "body", string(m.StorageClassName), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PersistentVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PersistentVolume) UnmarshalBinary(b []byte) error {
	var res PersistentVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
