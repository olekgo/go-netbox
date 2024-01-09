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

// checks if the WritableL2VPNTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableL2VPNTerminationRequest{}

// WritableL2VPNTerminationRequest Adds support for custom fields and tags.
type WritableL2VPNTerminationRequest struct {
	L2vpn                int32                  `json:"l2vpn"`
	AssignedObjectType   string                 `json:"assigned_object_type"`
	AssignedObjectId     int64                  `json:"assigned_object_id"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableL2VPNTerminationRequest WritableL2VPNTerminationRequest

// NewWritableL2VPNTerminationRequest instantiates a new WritableL2VPNTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableL2VPNTerminationRequest(l2vpn int32, assignedObjectType string, assignedObjectId int64) *WritableL2VPNTerminationRequest {
	this := WritableL2VPNTerminationRequest{}
	this.L2vpn = l2vpn
	this.AssignedObjectType = assignedObjectType
	this.AssignedObjectId = assignedObjectId
	return &this
}

// NewWritableL2VPNTerminationRequestWithDefaults instantiates a new WritableL2VPNTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableL2VPNTerminationRequestWithDefaults() *WritableL2VPNTerminationRequest {
	this := WritableL2VPNTerminationRequest{}
	return &this
}

// GetL2vpn returns the L2vpn field value
func (o *WritableL2VPNTerminationRequest) GetL2vpn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.L2vpn
}

// GetL2vpnOk returns a tuple with the L2vpn field value
// and a boolean to check if the value has been set.
func (o *WritableL2VPNTerminationRequest) GetL2vpnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L2vpn, true
}

// SetL2vpn sets field value
func (o *WritableL2VPNTerminationRequest) SetL2vpn(v int32) {
	o.L2vpn = v
}

// GetAssignedObjectType returns the AssignedObjectType field value
func (o *WritableL2VPNTerminationRequest) GetAssignedObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedObjectType
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value
// and a boolean to check if the value has been set.
func (o *WritableL2VPNTerminationRequest) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectType, true
}

// SetAssignedObjectType sets field value
func (o *WritableL2VPNTerminationRequest) SetAssignedObjectType(v string) {
	o.AssignedObjectType = v
}

// GetAssignedObjectId returns the AssignedObjectId field value
func (o *WritableL2VPNTerminationRequest) GetAssignedObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AssignedObjectId
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value
// and a boolean to check if the value has been set.
func (o *WritableL2VPNTerminationRequest) GetAssignedObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectId, true
}

// SetAssignedObjectId sets field value
func (o *WritableL2VPNTerminationRequest) SetAssignedObjectId(v int64) {
	o.AssignedObjectId = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableL2VPNTerminationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableL2VPNTerminationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableL2VPNTerminationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableL2VPNTerminationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableL2VPNTerminationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableL2VPNTerminationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableL2VPNTerminationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableL2VPNTerminationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableL2VPNTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableL2VPNTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["l2vpn"] = o.L2vpn
	toSerialize["assigned_object_type"] = o.AssignedObjectType
	toSerialize["assigned_object_id"] = o.AssignedObjectId
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableL2VPNTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"l2vpn",
		"assigned_object_type",
		"assigned_object_id",
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

	varWritableL2VPNTerminationRequest := _WritableL2VPNTerminationRequest{}

	err = json.Unmarshal(data, &varWritableL2VPNTerminationRequest)

	if err != nil {
		return err
	}

	*o = WritableL2VPNTerminationRequest(varWritableL2VPNTerminationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "l2vpn")
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableL2VPNTerminationRequest struct {
	value *WritableL2VPNTerminationRequest
	isSet bool
}

func (v NullableWritableL2VPNTerminationRequest) Get() *WritableL2VPNTerminationRequest {
	return v.value
}

func (v *NullableWritableL2VPNTerminationRequest) Set(val *WritableL2VPNTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableL2VPNTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableL2VPNTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableL2VPNTerminationRequest(val *WritableL2VPNTerminationRequest) *NullableWritableL2VPNTerminationRequest {
	return &NullableWritableL2VPNTerminationRequest{value: val, isSet: true}
}

func (v NullableWritableL2VPNTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableL2VPNTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
