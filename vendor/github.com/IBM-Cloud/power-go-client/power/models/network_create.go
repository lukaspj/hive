// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkCreate network create
// Example: {"accessConfig":"internal-only","cidr":"192.168.1.0/24","gateway":"192.168.1.1","ipAddressRanges":[{"endingIPAddress":"192.168.1.254","startingIPAddress":"192.168.1.2"}],"mtu":1450,"name":"sample-network","type":"vlan"}
//
// swagger:model NetworkCreate
type NetworkCreate struct {

	// access config
	AccessConfig AccessConfig `json:"accessConfig,omitempty"`

	// Network in CIDR notation (192.168.0.0/24)
	Cidr string `json:"cidr,omitempty"`

	// DNS Servers. If not specified, default is 127.0.0.1 for 'vlan' (private network) and 9.9.9.9 for 'pub-vlan' (public network)
	DNSServers []string `json:"dnsServers"`

	// Gateway IP Address
	Gateway string `json:"gateway,omitempty"`

	// IP Address Ranges
	IPAddressRanges []*IPAddressRange `json:"ipAddressRanges"`

	// Enable MTU Jumbo Network (for multi-zone locations only)
	Jumbo bool `json:"jumbo,omitempty"`

	// Maximum transmission unit (for satellite locations only)
	// Maximum: 9000
	// Minimum: 1450
	Mtu *int64 `json:"mtu,omitempty"`

	// Network Name
	// Max Length: 255
	Name string `json:"name,omitempty"`

	// Type of Network - 'vlan' (private network) 'pub-vlan' (public network) 'dhcp-vlan' (for satellite locations only)
	// Required: true
	// Enum: [vlan pub-vlan dhcp-vlan]
	Type *string `json:"type"`
}

// Validate validates this network create
func (m *NetworkCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddressRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkCreate) validateAccessConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessConfig) { // not required
		return nil
	}

	if err := m.AccessConfig.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accessConfig")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accessConfig")
		}
		return err
	}

	return nil
}

func (m *NetworkCreate) validateIPAddressRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAddressRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.IPAddressRanges); i++ {
		if swag.IsZero(m.IPAddressRanges[i]) { // not required
			continue
		}

		if m.IPAddressRanges[i] != nil {
			if err := m.IPAddressRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkCreate) validateMtu(formats strfmt.Registry) error {
	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", *m.Mtu, 1450, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", *m.Mtu, 9000, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkCreate) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 255); err != nil {
		return err
	}

	return nil
}

var networkCreateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vlan","pub-vlan","dhcp-vlan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkCreateTypeTypePropEnum = append(networkCreateTypeTypePropEnum, v)
	}
}

const (

	// NetworkCreateTypeVlan captures enum value "vlan"
	NetworkCreateTypeVlan string = "vlan"

	// NetworkCreateTypePubDashVlan captures enum value "pub-vlan"
	NetworkCreateTypePubDashVlan string = "pub-vlan"

	// NetworkCreateTypeDhcpDashVlan captures enum value "dhcp-vlan"
	NetworkCreateTypeDhcpDashVlan string = "dhcp-vlan"
)

// prop value enum
func (m *NetworkCreate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkCreateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkCreate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this network create based on the context it is used
func (m *NetworkCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPAddressRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkCreate) contextValidateAccessConfig(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AccessConfig) { // not required
		return nil
	}

	if err := m.AccessConfig.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accessConfig")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accessConfig")
		}
		return err
	}

	return nil
}

func (m *NetworkCreate) contextValidateIPAddressRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAddressRanges); i++ {

		if m.IPAddressRanges[i] != nil {

			if swag.IsZero(m.IPAddressRanges[i]) { // not required
				return nil
			}

			if err := m.IPAddressRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCreate) UnmarshalBinary(b []byte) error {
	var res NetworkCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}