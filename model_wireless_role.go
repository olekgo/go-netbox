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

// WirelessRole * `ap` - Access point * `station` - Station
type WirelessRole string

// List of Wireless_role
const (
	WIRELESSROLE_AP      WirelessRole = "ap"
	WIRELESSROLE_STATION WirelessRole = "station"
	WIRELESSROLE_EMPTY   WirelessRole = ""
)

// All allowed values of WirelessRole enum
var AllowedWirelessRoleEnumValues = []WirelessRole{
	"ap",
	"station",
	"",
}

func (v *WirelessRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WirelessRole(value)
	for _, existing := range AllowedWirelessRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WirelessRole", value)
}

// NewWirelessRoleFromValue returns a pointer to a valid WirelessRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWirelessRoleFromValue(v string) (*WirelessRole, error) {
	ev := WirelessRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WirelessRole: valid values are %v", v, AllowedWirelessRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WirelessRole) IsValid() bool {
	for _, existing := range AllowedWirelessRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Wireless_role value
func (v WirelessRole) Ptr() *WirelessRole {
	return &v
}

type NullableWirelessRole struct {
	value *WirelessRole
	isSet bool
}

func (v NullableWirelessRole) Get() *WirelessRole {
	return v.value
}

func (v *NullableWirelessRole) Set(val *WirelessRole) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessRole) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessRole(val *WirelessRole) *NullableWirelessRole {
	return &NullableWirelessRole{value: val, isSet: true}
}

func (v NullableWirelessRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
