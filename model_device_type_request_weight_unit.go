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

// DeviceTypeRequestWeightUnit * `kg` - Kilograms * `g` - Grams * `lb` - Pounds * `oz` - Ounces
type DeviceTypeRequestWeightUnit string

// List of DeviceTypeRequest_weight_unit
const (
	DEVICETYPEREQUESTWEIGHTUNIT_KG    DeviceTypeRequestWeightUnit = "kg"
	DEVICETYPEREQUESTWEIGHTUNIT_G     DeviceTypeRequestWeightUnit = "g"
	DEVICETYPEREQUESTWEIGHTUNIT_LB    DeviceTypeRequestWeightUnit = "lb"
	DEVICETYPEREQUESTWEIGHTUNIT_OZ    DeviceTypeRequestWeightUnit = "oz"
	DEVICETYPEREQUESTWEIGHTUNIT_EMPTY DeviceTypeRequestWeightUnit = ""
)

// All allowed values of DeviceTypeRequestWeightUnit enum
var AllowedDeviceTypeRequestWeightUnitEnumValues = []DeviceTypeRequestWeightUnit{
	"kg",
	"g",
	"lb",
	"oz",
	"",
}

func (v *DeviceTypeRequestWeightUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceTypeRequestWeightUnit(value)
	for _, existing := range AllowedDeviceTypeRequestWeightUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceTypeRequestWeightUnit", value)
}

// NewDeviceTypeRequestWeightUnitFromValue returns a pointer to a valid DeviceTypeRequestWeightUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceTypeRequestWeightUnitFromValue(v string) (*DeviceTypeRequestWeightUnit, error) {
	ev := DeviceTypeRequestWeightUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceTypeRequestWeightUnit: valid values are %v", v, AllowedDeviceTypeRequestWeightUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceTypeRequestWeightUnit) IsValid() bool {
	for _, existing := range AllowedDeviceTypeRequestWeightUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceTypeRequest_weight_unit value
func (v DeviceTypeRequestWeightUnit) Ptr() *DeviceTypeRequestWeightUnit {
	return &v
}

type NullableDeviceTypeRequestWeightUnit struct {
	value *DeviceTypeRequestWeightUnit
	isSet bool
}

func (v NullableDeviceTypeRequestWeightUnit) Get() *DeviceTypeRequestWeightUnit {
	return v.value
}

func (v *NullableDeviceTypeRequestWeightUnit) Set(val *DeviceTypeRequestWeightUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTypeRequestWeightUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTypeRequestWeightUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTypeRequestWeightUnit(val *DeviceTypeRequestWeightUnit) *NullableDeviceTypeRequestWeightUnit {
	return &NullableDeviceTypeRequestWeightUnit{value: val, isSet: true}
}

func (v NullableDeviceTypeRequestWeightUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTypeRequestWeightUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
