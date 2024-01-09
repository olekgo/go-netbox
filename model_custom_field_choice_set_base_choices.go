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

// checks if the CustomFieldChoiceSetBaseChoices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldChoiceSetBaseChoices{}

// CustomFieldChoiceSetBaseChoices struct for CustomFieldChoiceSetBaseChoices
type CustomFieldChoiceSetBaseChoices struct {
	Value                *CustomFieldChoiceSetBaseChoicesValue `json:"value,omitempty"`
	Label                *CustomFieldChoiceSetBaseChoicesLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomFieldChoiceSetBaseChoices CustomFieldChoiceSetBaseChoices

// NewCustomFieldChoiceSetBaseChoices instantiates a new CustomFieldChoiceSetBaseChoices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldChoiceSetBaseChoices() *CustomFieldChoiceSetBaseChoices {
	this := CustomFieldChoiceSetBaseChoices{}
	return &this
}

// NewCustomFieldChoiceSetBaseChoicesWithDefaults instantiates a new CustomFieldChoiceSetBaseChoices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldChoiceSetBaseChoicesWithDefaults() *CustomFieldChoiceSetBaseChoices {
	this := CustomFieldChoiceSetBaseChoices{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CustomFieldChoiceSetBaseChoices) GetValue() CustomFieldChoiceSetBaseChoicesValue {
	if o == nil || IsNil(o.Value) {
		var ret CustomFieldChoiceSetBaseChoicesValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSetBaseChoices) GetValueOk() (*CustomFieldChoiceSetBaseChoicesValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CustomFieldChoiceSetBaseChoices) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CustomFieldChoiceSetBaseChoicesValue and assigns it to the Value field.
func (o *CustomFieldChoiceSetBaseChoices) SetValue(v CustomFieldChoiceSetBaseChoicesValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CustomFieldChoiceSetBaseChoices) GetLabel() CustomFieldChoiceSetBaseChoicesLabel {
	if o == nil || IsNil(o.Label) {
		var ret CustomFieldChoiceSetBaseChoicesLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSetBaseChoices) GetLabelOk() (*CustomFieldChoiceSetBaseChoicesLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CustomFieldChoiceSetBaseChoices) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given CustomFieldChoiceSetBaseChoicesLabel and assigns it to the Label field.
func (o *CustomFieldChoiceSetBaseChoices) SetLabel(v CustomFieldChoiceSetBaseChoicesLabel) {
	o.Label = &v
}

func (o CustomFieldChoiceSetBaseChoices) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldChoiceSetBaseChoices) ToMap() (map[string]interface{}, error) {
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

func (o *CustomFieldChoiceSetBaseChoices) UnmarshalJSON(data []byte) (err error) {
	varCustomFieldChoiceSetBaseChoices := _CustomFieldChoiceSetBaseChoices{}

	err = json.Unmarshal(data, &varCustomFieldChoiceSetBaseChoices)

	if err != nil {
		return err
	}

	*o = CustomFieldChoiceSetBaseChoices(varCustomFieldChoiceSetBaseChoices)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomFieldChoiceSetBaseChoices struct {
	value *CustomFieldChoiceSetBaseChoices
	isSet bool
}

func (v NullableCustomFieldChoiceSetBaseChoices) Get() *CustomFieldChoiceSetBaseChoices {
	return v.value
}

func (v *NullableCustomFieldChoiceSetBaseChoices) Set(val *CustomFieldChoiceSetBaseChoices) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldChoiceSetBaseChoices) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldChoiceSetBaseChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldChoiceSetBaseChoices(val *CustomFieldChoiceSetBaseChoices) *NullableCustomFieldChoiceSetBaseChoices {
	return &NullableCustomFieldChoiceSetBaseChoices{value: val, isSet: true}
}

func (v NullableCustomFieldChoiceSetBaseChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldChoiceSetBaseChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
