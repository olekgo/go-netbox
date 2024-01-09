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

// PowerFeedSupplyLabel the model 'PowerFeedSupplyLabel'
type PowerFeedSupplyLabel string

// List of PowerFeed_supply_label
const (
	POWERFEEDSUPPLYLABEL_AC PowerFeedSupplyLabel = "AC"
	POWERFEEDSUPPLYLABEL_DC PowerFeedSupplyLabel = "DC"
)

// All allowed values of PowerFeedSupplyLabel enum
var AllowedPowerFeedSupplyLabelEnumValues = []PowerFeedSupplyLabel{
	"AC",
	"DC",
}

func (v *PowerFeedSupplyLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerFeedSupplyLabel(value)
	for _, existing := range AllowedPowerFeedSupplyLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerFeedSupplyLabel", value)
}

// NewPowerFeedSupplyLabelFromValue returns a pointer to a valid PowerFeedSupplyLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerFeedSupplyLabelFromValue(v string) (*PowerFeedSupplyLabel, error) {
	ev := PowerFeedSupplyLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerFeedSupplyLabel: valid values are %v", v, AllowedPowerFeedSupplyLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerFeedSupplyLabel) IsValid() bool {
	for _, existing := range AllowedPowerFeedSupplyLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerFeed_supply_label value
func (v PowerFeedSupplyLabel) Ptr() *PowerFeedSupplyLabel {
	return &v
}

type NullablePowerFeedSupplyLabel struct {
	value *PowerFeedSupplyLabel
	isSet bool
}

func (v NullablePowerFeedSupplyLabel) Get() *PowerFeedSupplyLabel {
	return v.value
}

func (v *NullablePowerFeedSupplyLabel) Set(val *PowerFeedSupplyLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedSupplyLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedSupplyLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedSupplyLabel(val *PowerFeedSupplyLabel) *NullablePowerFeedSupplyLabel {
	return &NullablePowerFeedSupplyLabel{value: val, isSet: true}
}

func (v NullablePowerFeedSupplyLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedSupplyLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
