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

// AuthenticationType1 * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise
type AuthenticationType1 string

// List of Authentication_type_1
const (
	AUTHENTICATIONTYPE1_OPEN           AuthenticationType1 = "open"
	AUTHENTICATIONTYPE1_WEP            AuthenticationType1 = "wep"
	AUTHENTICATIONTYPE1_WPA_PERSONAL   AuthenticationType1 = "wpa-personal"
	AUTHENTICATIONTYPE1_WPA_ENTERPRISE AuthenticationType1 = "wpa-enterprise"
	AUTHENTICATIONTYPE1_EMPTY          AuthenticationType1 = ""
)

// All allowed values of AuthenticationType1 enum
var AllowedAuthenticationType1EnumValues = []AuthenticationType1{
	"open",
	"wep",
	"wpa-personal",
	"wpa-enterprise",
	"",
}

func (v *AuthenticationType1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticationType1(value)
	for _, existing := range AllowedAuthenticationType1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticationType1", value)
}

// NewAuthenticationType1FromValue returns a pointer to a valid AuthenticationType1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticationType1FromValue(v string) (*AuthenticationType1, error) {
	ev := AuthenticationType1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthenticationType1: valid values are %v", v, AllowedAuthenticationType1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthenticationType1) IsValid() bool {
	for _, existing := range AllowedAuthenticationType1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Authentication_type_1 value
func (v AuthenticationType1) Ptr() *AuthenticationType1 {
	return &v
}

type NullableAuthenticationType1 struct {
	value *AuthenticationType1
	isSet bool
}

func (v NullableAuthenticationType1) Get() *AuthenticationType1 {
	return v.value
}

func (v *NullableAuthenticationType1) Set(val *AuthenticationType1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationType1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationType1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationType1(val *AuthenticationType1) *NullableAuthenticationType1 {
	return &NullableAuthenticationType1{value: val, isSet: true}
}

func (v NullableAuthenticationType1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationType1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
