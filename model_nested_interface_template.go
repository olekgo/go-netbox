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

// checks if the NestedInterfaceTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedInterfaceTemplate{}

// NestedInterfaceTemplate Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedInterfaceTemplate struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedInterfaceTemplate NestedInterfaceTemplate

// NewNestedInterfaceTemplate instantiates a new NestedInterfaceTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedInterfaceTemplate(id int32, url string, display string, name string) *NestedInterfaceTemplate {
	this := NestedInterfaceTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	return &this
}

// NewNestedInterfaceTemplateWithDefaults instantiates a new NestedInterfaceTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedInterfaceTemplateWithDefaults() *NestedInterfaceTemplate {
	this := NestedInterfaceTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *NestedInterfaceTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedInterfaceTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedInterfaceTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedInterfaceTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedInterfaceTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedInterfaceTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedInterfaceTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedInterfaceTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedInterfaceTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedInterfaceTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedInterfaceTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedInterfaceTemplate) SetName(v string) {
	o.Name = v
}

func (o NestedInterfaceTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedInterfaceTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedInterfaceTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNestedInterfaceTemplate := _NestedInterfaceTemplate{}

	err = json.Unmarshal(data, &varNestedInterfaceTemplate)

	if err != nil {
		return err
	}

	*o = NestedInterfaceTemplate(varNestedInterfaceTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedInterfaceTemplate struct {
	value *NestedInterfaceTemplate
	isSet bool
}

func (v NullableNestedInterfaceTemplate) Get() *NestedInterfaceTemplate {
	return v.value
}

func (v *NullableNestedInterfaceTemplate) Set(val *NestedInterfaceTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedInterfaceTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedInterfaceTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedInterfaceTemplate(val *NestedInterfaceTemplate) *NullableNestedInterfaceTemplate {
	return &NullableNestedInterfaceTemplate{value: val, isSet: true}
}

func (v NullableNestedInterfaceTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedInterfaceTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
