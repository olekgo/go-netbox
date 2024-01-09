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

// checks if the NestedFHRPGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedFHRPGroup{}

// NestedFHRPGroup Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedFHRPGroup struct {
	Id                   int32             `json:"id"`
	Url                  string            `json:"url"`
	Display              string            `json:"display"`
	Protocol             FHRPGroupProtocol `json:"protocol"`
	GroupId              int32             `json:"group_id"`
	AdditionalProperties map[string]interface{}
}

type _NestedFHRPGroup NestedFHRPGroup

// NewNestedFHRPGroup instantiates a new NestedFHRPGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedFHRPGroup(id int32, url string, display string, protocol FHRPGroupProtocol, groupId int32) *NestedFHRPGroup {
	this := NestedFHRPGroup{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Protocol = protocol
	this.GroupId = groupId
	return &this
}

// NewNestedFHRPGroupWithDefaults instantiates a new NestedFHRPGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedFHRPGroupWithDefaults() *NestedFHRPGroup {
	this := NestedFHRPGroup{}
	return &this
}

// GetId returns the Id field value
func (o *NestedFHRPGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedFHRPGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedFHRPGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedFHRPGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedFHRPGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedFHRPGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedFHRPGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedFHRPGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedFHRPGroup) SetDisplay(v string) {
	o.Display = v
}

// GetProtocol returns the Protocol field value
func (o *NestedFHRPGroup) GetProtocol() FHRPGroupProtocol {
	if o == nil {
		var ret FHRPGroupProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *NestedFHRPGroup) GetProtocolOk() (*FHRPGroupProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *NestedFHRPGroup) SetProtocol(v FHRPGroupProtocol) {
	o.Protocol = v
}

// GetGroupId returns the GroupId field value
func (o *NestedFHRPGroup) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *NestedFHRPGroup) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *NestedFHRPGroup) SetGroupId(v int32) {
	o.GroupId = v
}

func (o NestedFHRPGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedFHRPGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["protocol"] = o.Protocol
	toSerialize["group_id"] = o.GroupId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedFHRPGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"protocol",
		"group_id",
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

	varNestedFHRPGroup := _NestedFHRPGroup{}

	err = json.Unmarshal(data, &varNestedFHRPGroup)

	if err != nil {
		return err
	}

	*o = NestedFHRPGroup(varNestedFHRPGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "group_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedFHRPGroup struct {
	value *NestedFHRPGroup
	isSet bool
}

func (v NullableNestedFHRPGroup) Get() *NestedFHRPGroup {
	return v.value
}

func (v *NullableNestedFHRPGroup) Set(val *NestedFHRPGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedFHRPGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedFHRPGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedFHRPGroup(val *NestedFHRPGroup) *NullableNestedFHRPGroup {
	return &NullableNestedFHRPGroup{value: val, isSet: true}
}

func (v NullableNestedFHRPGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedFHRPGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
