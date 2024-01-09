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

// ConsolePortRequestSpeed * `1200` - 1200 bps * `2400` - 2400 bps * `4800` - 4800 bps * `9600` - 9600 bps * `19200` - 19.2 kbps * `38400` - 38.4 kbps * `57600` - 57.6 kbps * `115200` - 115.2 kbps
type ConsolePortRequestSpeed int32

// List of ConsolePortRequest_speed
const (
	CONSOLEPORTREQUESTSPEED__1200   ConsolePortRequestSpeed = 1200
	CONSOLEPORTREQUESTSPEED__2400   ConsolePortRequestSpeed = 2400
	CONSOLEPORTREQUESTSPEED__4800   ConsolePortRequestSpeed = 4800
	CONSOLEPORTREQUESTSPEED__9600   ConsolePortRequestSpeed = 9600
	CONSOLEPORTREQUESTSPEED__19200  ConsolePortRequestSpeed = 19200
	CONSOLEPORTREQUESTSPEED__38400  ConsolePortRequestSpeed = 38400
	CONSOLEPORTREQUESTSPEED__57600  ConsolePortRequestSpeed = 57600
	CONSOLEPORTREQUESTSPEED__115200 ConsolePortRequestSpeed = 115200
)

// All allowed values of ConsolePortRequestSpeed enum
var AllowedConsolePortRequestSpeedEnumValues = []ConsolePortRequestSpeed{
	1200,
	2400,
	4800,
	9600,
	19200,
	38400,
	57600,
	115200,
}

func (v *ConsolePortRequestSpeed) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsolePortRequestSpeed(value)
	for _, existing := range AllowedConsolePortRequestSpeedEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsolePortRequestSpeed", value)
}

// NewConsolePortRequestSpeedFromValue returns a pointer to a valid ConsolePortRequestSpeed
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsolePortRequestSpeedFromValue(v int32) (*ConsolePortRequestSpeed, error) {
	ev := ConsolePortRequestSpeed(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsolePortRequestSpeed: valid values are %v", v, AllowedConsolePortRequestSpeedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsolePortRequestSpeed) IsValid() bool {
	for _, existing := range AllowedConsolePortRequestSpeedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsolePortRequest_speed value
func (v ConsolePortRequestSpeed) Ptr() *ConsolePortRequestSpeed {
	return &v
}

type NullableConsolePortRequestSpeed struct {
	value *ConsolePortRequestSpeed
	isSet bool
}

func (v NullableConsolePortRequestSpeed) Get() *ConsolePortRequestSpeed {
	return v.value
}

func (v *NullableConsolePortRequestSpeed) Set(val *ConsolePortRequestSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePortRequestSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePortRequestSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePortRequestSpeed(val *ConsolePortRequestSpeed) *NullableConsolePortRequestSpeed {
	return &NullableConsolePortRequestSpeed{value: val, isSet: true}
}

func (v NullableConsolePortRequestSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePortRequestSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
