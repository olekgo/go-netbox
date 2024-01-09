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

// checks if the LocationStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationStatus{}

// LocationStatus struct for LocationStatus
type LocationStatus struct {
	Value                *LocationStatusValue `json:"value,omitempty"`
	Label                *LocationStatusLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LocationStatus LocationStatus

// NewLocationStatus instantiates a new LocationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationStatus() *LocationStatus {
	this := LocationStatus{}
	return &this
}

// NewLocationStatusWithDefaults instantiates a new LocationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationStatusWithDefaults() *LocationStatus {
	this := LocationStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LocationStatus) GetValue() LocationStatusValue {
	if o == nil || IsNil(o.Value) {
		var ret LocationStatusValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationStatus) GetValueOk() (*LocationStatusValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LocationStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given LocationStatusValue and assigns it to the Value field.
func (o *LocationStatus) SetValue(v LocationStatusValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *LocationStatus) GetLabel() LocationStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret LocationStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationStatus) GetLabelOk() (*LocationStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *LocationStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given LocationStatusLabel and assigns it to the Label field.
func (o *LocationStatus) SetLabel(v LocationStatusLabel) {
	o.Label = &v
}

func (o LocationStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationStatus) ToMap() (map[string]interface{}, error) {
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

func (o *LocationStatus) UnmarshalJSON(data []byte) (err error) {
	varLocationStatus := _LocationStatus{}

	err = json.Unmarshal(data, &varLocationStatus)

	if err != nil {
		return err
	}

	*o = LocationStatus(varLocationStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLocationStatus struct {
	value *LocationStatus
	isSet bool
}

func (v NullableLocationStatus) Get() *LocationStatus {
	return v.value
}

func (v *NullableLocationStatus) Set(val *LocationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationStatus(val *LocationStatus) *NullableLocationStatus {
	return &NullableLocationStatus{value: val, isSet: true}
}

func (v NullableLocationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
