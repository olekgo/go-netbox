/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.0 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PowerFeedPhase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerFeedPhase{}

// PowerFeedPhase struct for PowerFeedPhase
type PowerFeedPhase struct {
	Value                *PatchedWritablePowerFeedRequestPhase `json:"value,omitempty"`
	Label                *PowerFeedPhaseLabel                  `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerFeedPhase PowerFeedPhase

// NewPowerFeedPhase instantiates a new PowerFeedPhase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerFeedPhase() *PowerFeedPhase {
	this := PowerFeedPhase{}
	return &this
}

// NewPowerFeedPhaseWithDefaults instantiates a new PowerFeedPhase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerFeedPhaseWithDefaults() *PowerFeedPhase {
	this := PowerFeedPhase{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PowerFeedPhase) GetValue() PatchedWritablePowerFeedRequestPhase {
	if o == nil || IsNil(o.Value) {
		var ret PatchedWritablePowerFeedRequestPhase
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedPhase) GetValueOk() (*PatchedWritablePowerFeedRequestPhase, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PowerFeedPhase) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given PatchedWritablePowerFeedRequestPhase and assigns it to the Value field.
func (o *PowerFeedPhase) SetValue(v PatchedWritablePowerFeedRequestPhase) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerFeedPhase) GetLabel() PowerFeedPhaseLabel {
	if o == nil || IsNil(o.Label) {
		var ret PowerFeedPhaseLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedPhase) GetLabelOk() (*PowerFeedPhaseLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerFeedPhase) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given PowerFeedPhaseLabel and assigns it to the Label field.
func (o *PowerFeedPhase) SetLabel(v PowerFeedPhaseLabel) {
	o.Label = &v
}

func (o PowerFeedPhase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerFeedPhase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PowerFeedPhase) UnmarshalJSON(data []byte) (err error) {
	varPowerFeedPhase := _PowerFeedPhase{}

	err = json.Unmarshal(data, &varPowerFeedPhase)

	if err != nil {
		return err
	}

	*o = PowerFeedPhase(varPowerFeedPhase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerFeedPhase struct {
	value *PowerFeedPhase
	isSet bool
}

func (v NullablePowerFeedPhase) Get() *PowerFeedPhase {
	return v.value
}

func (v *NullablePowerFeedPhase) Set(val *PowerFeedPhase) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedPhase) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedPhase(val *PowerFeedPhase) *NullablePowerFeedPhase {
	return &NullablePowerFeedPhase{value: val, isSet: true}
}

func (v NullablePowerFeedPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
