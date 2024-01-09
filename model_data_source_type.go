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

// checks if the DataSourceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSourceType{}

// DataSourceType struct for DataSourceType
type DataSourceType struct {
	Value                *DataSourceTypeValue `json:"value,omitempty"`
	Label                *DataSourceTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DataSourceType DataSourceType

// NewDataSourceType instantiates a new DataSourceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceType() *DataSourceType {
	this := DataSourceType{}
	return &this
}

// NewDataSourceTypeWithDefaults instantiates a new DataSourceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceTypeWithDefaults() *DataSourceType {
	this := DataSourceType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DataSourceType) GetValue() DataSourceTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret DataSourceTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceType) GetValueOk() (*DataSourceTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DataSourceType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given DataSourceTypeValue and assigns it to the Value field.
func (o *DataSourceType) SetValue(v DataSourceTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DataSourceType) GetLabel() DataSourceTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret DataSourceTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceType) GetLabelOk() (*DataSourceTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DataSourceType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given DataSourceTypeLabel and assigns it to the Label field.
func (o *DataSourceType) SetLabel(v DataSourceTypeLabel) {
	o.Label = &v
}

func (o DataSourceType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourceType) ToMap() (map[string]interface{}, error) {
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

func (o *DataSourceType) UnmarshalJSON(data []byte) (err error) {
	varDataSourceType := _DataSourceType{}

	err = json.Unmarshal(data, &varDataSourceType)

	if err != nil {
		return err
	}

	*o = DataSourceType(varDataSourceType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataSourceType struct {
	value *DataSourceType
	isSet bool
}

func (v NullableDataSourceType) Get() *DataSourceType {
	return v.value
}

func (v *NullableDataSourceType) Set(val *DataSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceType(val *DataSourceType) *NullableDataSourceType {
	return &NullableDataSourceType{value: val, isSet: true}
}

func (v NullableDataSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
