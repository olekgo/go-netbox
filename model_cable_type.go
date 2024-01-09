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

// CableType * `cat3` - CAT3 * `cat5` - CAT5 * `cat5e` - CAT5e * `cat6` - CAT6 * `cat6a` - CAT6a * `cat7` - CAT7 * `cat7a` - CAT7a * `cat8` - CAT8 * `dac-active` - Direct Attach Copper (Active) * `dac-passive` - Direct Attach Copper (Passive) * `mrj21-trunk` - MRJ21 Trunk * `coaxial` - Coaxial * `mmf` - Multimode Fiber * `mmf-om1` - Multimode Fiber (OM1) * `mmf-om2` - Multimode Fiber (OM2) * `mmf-om3` - Multimode Fiber (OM3) * `mmf-om4` - Multimode Fiber (OM4) * `mmf-om5` - Multimode Fiber (OM5) * `smf` - Singlemode Fiber * `smf-os1` - Singlemode Fiber (OS1) * `smf-os2` - Singlemode Fiber (OS2) * `aoc` - Active Optical Cabling (AOC) * `power` - Power
type CableType string

// List of Cable_type
const (
	CABLETYPE_CAT3        CableType = "cat3"
	CABLETYPE_CAT5        CableType = "cat5"
	CABLETYPE_CAT5E       CableType = "cat5e"
	CABLETYPE_CAT6        CableType = "cat6"
	CABLETYPE_CAT6A       CableType = "cat6a"
	CABLETYPE_CAT7        CableType = "cat7"
	CABLETYPE_CAT7A       CableType = "cat7a"
	CABLETYPE_CAT8        CableType = "cat8"
	CABLETYPE_DAC_ACTIVE  CableType = "dac-active"
	CABLETYPE_DAC_PASSIVE CableType = "dac-passive"
	CABLETYPE_MRJ21_TRUNK CableType = "mrj21-trunk"
	CABLETYPE_COAXIAL     CableType = "coaxial"
	CABLETYPE_MMF         CableType = "mmf"
	CABLETYPE_MMF_OM1     CableType = "mmf-om1"
	CABLETYPE_MMF_OM2     CableType = "mmf-om2"
	CABLETYPE_MMF_OM3     CableType = "mmf-om3"
	CABLETYPE_MMF_OM4     CableType = "mmf-om4"
	CABLETYPE_MMF_OM5     CableType = "mmf-om5"
	CABLETYPE_SMF         CableType = "smf"
	CABLETYPE_SMF_OS1     CableType = "smf-os1"
	CABLETYPE_SMF_OS2     CableType = "smf-os2"
	CABLETYPE_AOC         CableType = "aoc"
	CABLETYPE_POWER       CableType = "power"
	CABLETYPE_EMPTY       CableType = ""
)

// All allowed values of CableType enum
var AllowedCableTypeEnumValues = []CableType{
	"cat3",
	"cat5",
	"cat5e",
	"cat6",
	"cat6a",
	"cat7",
	"cat7a",
	"cat8",
	"dac-active",
	"dac-passive",
	"mrj21-trunk",
	"coaxial",
	"mmf",
	"mmf-om1",
	"mmf-om2",
	"mmf-om3",
	"mmf-om4",
	"mmf-om5",
	"smf",
	"smf-os1",
	"smf-os2",
	"aoc",
	"power",
	"",
}

func (v *CableType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CableType(value)
	for _, existing := range AllowedCableTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CableType", value)
}

// NewCableTypeFromValue returns a pointer to a valid CableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCableTypeFromValue(v string) (*CableType, error) {
	ev := CableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CableType: valid values are %v", v, AllowedCableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CableType) IsValid() bool {
	for _, existing := range AllowedCableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Cable_type value
func (v CableType) Ptr() *CableType {
	return &v
}

type NullableCableType struct {
	value *CableType
	isSet bool
}

func (v NullableCableType) Get() *CableType {
	return v.value
}

func (v *NullableCableType) Set(val *CableType) {
	v.value = val
	v.isSet = true
}

func (v NullableCableType) IsSet() bool {
	return v.isSet
}

func (v *NullableCableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableType(val *CableType) *NullableCableType {
	return &NullableCableType{value: val, isSet: true}
}

func (v NullableCableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
