// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceInfo Device information
//
// swagger:model DeviceInfo
type DeviceInfo struct {

	// Desired IOPS for the volumes mounted on this instance
	DiskIops int32 `json:"diskIops,omitempty"`

	// Comma-separated list of mount points for the devices in each instance
	MountPoints string `json:"mountPoints,omitempty"`

	// Number of volumes to be mounted on this instance at the default path
	NumVolumes int32 `json:"numVolumes,omitempty"`

	// Name of the storage class
	StorageClass string `json:"storageClass,omitempty"`

	// Storage type used for this instance
	// Enum: [IO1 GP2 GP3 Scratch Persistent StandardSSD_LRS Premium_LRS UltraSSD_LRS]
	StorageType string `json:"storageType,omitempty"`

	// Desired throughput for the volumes mounted on this instance
	Throughput int32 `json:"throughput,omitempty"`

	// The size of each volume in each instance
	VolumeSize int32 `json:"volumeSize,omitempty"`
}

// Validate validates this device info
func (m *DeviceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deviceInfoTypeStorageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IO1","GP2","GP3","Scratch","Persistent","StandardSSD_LRS","Premium_LRS","UltraSSD_LRS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceInfoTypeStorageTypePropEnum = append(deviceInfoTypeStorageTypePropEnum, v)
	}
}

const (

	// DeviceInfoStorageTypeIO1 captures enum value "IO1"
	DeviceInfoStorageTypeIO1 string = "IO1"

	// DeviceInfoStorageTypeGP2 captures enum value "GP2"
	DeviceInfoStorageTypeGP2 string = "GP2"

	// DeviceInfoStorageTypeGP3 captures enum value "GP3"
	DeviceInfoStorageTypeGP3 string = "GP3"

	// DeviceInfoStorageTypeScratch captures enum value "Scratch"
	DeviceInfoStorageTypeScratch string = "Scratch"

	// DeviceInfoStorageTypePersistent captures enum value "Persistent"
	DeviceInfoStorageTypePersistent string = "Persistent"

	// DeviceInfoStorageTypeStandardSSDLRS captures enum value "StandardSSD_LRS"
	DeviceInfoStorageTypeStandardSSDLRS string = "StandardSSD_LRS"

	// DeviceInfoStorageTypePremiumLRS captures enum value "Premium_LRS"
	DeviceInfoStorageTypePremiumLRS string = "Premium_LRS"

	// DeviceInfoStorageTypeUltraSSDLRS captures enum value "UltraSSD_LRS"
	DeviceInfoStorageTypeUltraSSDLRS string = "UltraSSD_LRS"
)

// prop value enum
func (m *DeviceInfo) validateStorageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceInfoTypeStorageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceInfo) validateStorageType(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateStorageTypeEnum("storageType", "body", m.StorageType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device info based on context it is used
func (m *DeviceInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceInfo) UnmarshalBinary(b []byte) error {
	var res DeviceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}