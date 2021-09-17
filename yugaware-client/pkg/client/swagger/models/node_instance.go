// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeInstance A single node instance, attached to a provider and availability zone
//
// swagger:model NodeInstance
type NodeInstance struct {

	// details
	// Required: true
	Details *NodeInstanceData `json:"details"`

	// Node details (as a JSON object)
	// Example: {\"ip\":\"1.1.1.1\",\"sshUser\":\"centos\"}
	DetailsJSON string `json:"detailsJson,omitempty"`

	// True if the node is in use
	InUse bool `json:"inUse,omitempty"`

	// The node instance's name
	// Example: Mumbai instance
	InstanceName string `json:"instanceName,omitempty"`

	// The node's type code
	// Example: c5large
	InstanceTypeCode string `json:"instanceTypeCode,omitempty"`

	// The node's name
	// Example: India node
	NodeName string `json:"nodeName,omitempty"`

	// The node's UUID
	// Read Only: true
	// Format: uuid
	NodeUUID strfmt.UUID `json:"nodeUuid,omitempty"`

	// The availability zone's UUID
	// Format: uuid
	ZoneUUID strfmt.UUID `json:"zoneUuid,omitempty"`
}

// Validate validates this node instance
func (m *NodeInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeInstance) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("details", "body", m.Details); err != nil {
		return err
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *NodeInstance) validateNodeUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("nodeUuid", "body", "uuid", m.NodeUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NodeInstance) validateZoneUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.ZoneUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("zoneUuid", "body", "uuid", m.ZoneUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this node instance based on the context it is used
func (m *NodeInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeInstance) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.Details != nil {
		if err := m.Details.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *NodeInstance) contextValidateNodeUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nodeUuid", "body", strfmt.UUID(m.NodeUUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeInstance) UnmarshalBinary(b []byte) error {
	var res NodeInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}