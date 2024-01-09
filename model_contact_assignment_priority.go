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

// checks if the ContactAssignmentPriority type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactAssignmentPriority{}

// ContactAssignmentPriority struct for ContactAssignmentPriority
type ContactAssignmentPriority struct {
	Value                *ContactAssignmentPriorityValue `json:"value,omitempty"`
	Label                *ContactAssignmentPriorityLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContactAssignmentPriority ContactAssignmentPriority

// NewContactAssignmentPriority instantiates a new ContactAssignmentPriority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactAssignmentPriority() *ContactAssignmentPriority {
	this := ContactAssignmentPriority{}
	return &this
}

// NewContactAssignmentPriorityWithDefaults instantiates a new ContactAssignmentPriority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactAssignmentPriorityWithDefaults() *ContactAssignmentPriority {
	this := ContactAssignmentPriority{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ContactAssignmentPriority) GetValue() ContactAssignmentPriorityValue {
	if o == nil || IsNil(o.Value) {
		var ret ContactAssignmentPriorityValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAssignmentPriority) GetValueOk() (*ContactAssignmentPriorityValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ContactAssignmentPriority) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ContactAssignmentPriorityValue and assigns it to the Value field.
func (o *ContactAssignmentPriority) SetValue(v ContactAssignmentPriorityValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ContactAssignmentPriority) GetLabel() ContactAssignmentPriorityLabel {
	if o == nil || IsNil(o.Label) {
		var ret ContactAssignmentPriorityLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAssignmentPriority) GetLabelOk() (*ContactAssignmentPriorityLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ContactAssignmentPriority) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given ContactAssignmentPriorityLabel and assigns it to the Label field.
func (o *ContactAssignmentPriority) SetLabel(v ContactAssignmentPriorityLabel) {
	o.Label = &v
}

func (o ContactAssignmentPriority) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactAssignmentPriority) ToMap() (map[string]interface{}, error) {
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

func (o *ContactAssignmentPriority) UnmarshalJSON(data []byte) (err error) {
	varContactAssignmentPriority := _ContactAssignmentPriority{}

	err = json.Unmarshal(data, &varContactAssignmentPriority)

	if err != nil {
		return err
	}

	*o = ContactAssignmentPriority(varContactAssignmentPriority)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContactAssignmentPriority struct {
	value *ContactAssignmentPriority
	isSet bool
}

func (v NullableContactAssignmentPriority) Get() *ContactAssignmentPriority {
	return v.value
}

func (v *NullableContactAssignmentPriority) Set(val *ContactAssignmentPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableContactAssignmentPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableContactAssignmentPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactAssignmentPriority(val *ContactAssignmentPriority) *NullableContactAssignmentPriority {
	return &NullableContactAssignmentPriority{value: val, isSet: true}
}

func (v NullableContactAssignmentPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactAssignmentPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
