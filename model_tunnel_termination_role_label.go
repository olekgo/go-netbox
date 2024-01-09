/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.0 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// TunnelTerminationRoleLabel the model 'TunnelTerminationRoleLabel'
type TunnelTerminationRoleLabel string

// List of TunnelTermination_role_label
const (
	TUNNELTERMINATIONROLELABEL_PEER  TunnelTerminationRoleLabel = "Peer"
	TUNNELTERMINATIONROLELABEL_HUB   TunnelTerminationRoleLabel = "Hub"
	TUNNELTERMINATIONROLELABEL_SPOKE TunnelTerminationRoleLabel = "Spoke"
)

// All allowed values of TunnelTerminationRoleLabel enum
var AllowedTunnelTerminationRoleLabelEnumValues = []TunnelTerminationRoleLabel{
	"Peer",
	"Hub",
	"Spoke",
}

func (v *TunnelTerminationRoleLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TunnelTerminationRoleLabel(value)
	for _, existing := range AllowedTunnelTerminationRoleLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TunnelTerminationRoleLabel", value)
}

// NewTunnelTerminationRoleLabelFromValue returns a pointer to a valid TunnelTerminationRoleLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTunnelTerminationRoleLabelFromValue(v string) (*TunnelTerminationRoleLabel, error) {
	ev := TunnelTerminationRoleLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TunnelTerminationRoleLabel: valid values are %v", v, AllowedTunnelTerminationRoleLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TunnelTerminationRoleLabel) IsValid() bool {
	for _, existing := range AllowedTunnelTerminationRoleLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TunnelTermination_role_label value
func (v TunnelTerminationRoleLabel) Ptr() *TunnelTerminationRoleLabel {
	return &v
}

type NullableTunnelTerminationRoleLabel struct {
	value *TunnelTerminationRoleLabel
	isSet bool
}

func (v NullableTunnelTerminationRoleLabel) Get() *TunnelTerminationRoleLabel {
	return v.value
}

func (v *NullableTunnelTerminationRoleLabel) Set(val *TunnelTerminationRoleLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelTerminationRoleLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelTerminationRoleLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelTerminationRoleLabel(val *TunnelTerminationRoleLabel) *NullableTunnelTerminationRoleLabel {
	return &NullableTunnelTerminationRoleLabel{value: val, isSet: true}
}

func (v NullableTunnelTerminationRoleLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelTerminationRoleLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
