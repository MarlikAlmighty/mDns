// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DNSEntry dns entry
//
// swagger:model dns_entry
type DNSEntry struct {

	// acme
	Acme []string `json:"acme"`

	// dkim private key
	DkimPrivateKey string `json:"dkim_private_key,omitempty"`

	// dkim public key
	DkimPublicKey string `json:"dkim_public_key,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// ipv4s
	Ipv4s []string `json:"ipv4s"`

	// ipv6s
	Ipv6s []string `json:"ipv6s"`
}

// Validate validates this dns entry
func (m *DNSEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dns entry based on context it is used
func (m *DNSEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSEntry) UnmarshalBinary(b []byte) error {
	var res DNSEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
