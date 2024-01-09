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

// AggregateFamilyValue * `4` - IPv4 * `6` - IPv6
type AggregateFamilyValue int32

// List of Aggregate_family_value
const (
	AGGREGATEFAMILYVALUE__4 AggregateFamilyValue = 4
	AGGREGATEFAMILYVALUE__6 AggregateFamilyValue = 6
)

// All allowed values of AggregateFamilyValue enum
var AllowedAggregateFamilyValueEnumValues = []AggregateFamilyValue{
	4,
	6,
}

func (v *AggregateFamilyValue) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AggregateFamilyValue(value)
	for _, existing := range AllowedAggregateFamilyValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AggregateFamilyValue", value)
}

// NewAggregateFamilyValueFromValue returns a pointer to a valid AggregateFamilyValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAggregateFamilyValueFromValue(v int32) (*AggregateFamilyValue, error) {
	ev := AggregateFamilyValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AggregateFamilyValue: valid values are %v", v, AllowedAggregateFamilyValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AggregateFamilyValue) IsValid() bool {
	for _, existing := range AllowedAggregateFamilyValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Aggregate_family_value value
func (v AggregateFamilyValue) Ptr() *AggregateFamilyValue {
	return &v
}

type NullableAggregateFamilyValue struct {
	value *AggregateFamilyValue
	isSet bool
}

func (v NullableAggregateFamilyValue) Get() *AggregateFamilyValue {
	return v.value
}

func (v *NullableAggregateFamilyValue) Set(val *AggregateFamilyValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateFamilyValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateFamilyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateFamilyValue(val *AggregateFamilyValue) *NullableAggregateFamilyValue {
	return &NullableAggregateFamilyValue{value: val, isSet: true}
}

func (v NullableAggregateFamilyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateFamilyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
