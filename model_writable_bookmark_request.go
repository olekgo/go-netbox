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

// checks if the WritableBookmarkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableBookmarkRequest{}

// WritableBookmarkRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableBookmarkRequest struct {
	ObjectType           string `json:"object_type"`
	ObjectId             int64  `json:"object_id"`
	User                 int32  `json:"user"`
	AdditionalProperties map[string]interface{}
}

type _WritableBookmarkRequest WritableBookmarkRequest

// NewWritableBookmarkRequest instantiates a new WritableBookmarkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableBookmarkRequest(objectType string, objectId int64, user int32) *WritableBookmarkRequest {
	this := WritableBookmarkRequest{}
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.User = user
	return &this
}

// NewWritableBookmarkRequestWithDefaults instantiates a new WritableBookmarkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableBookmarkRequestWithDefaults() *WritableBookmarkRequest {
	this := WritableBookmarkRequest{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *WritableBookmarkRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WritableBookmarkRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WritableBookmarkRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *WritableBookmarkRequest) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *WritableBookmarkRequest) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *WritableBookmarkRequest) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetUser returns the User field value
func (o *WritableBookmarkRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *WritableBookmarkRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *WritableBookmarkRequest) SetUser(v int32) {
	o.User = v
}

func (o WritableBookmarkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableBookmarkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableBookmarkRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_type",
		"object_id",
		"user",
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

	varWritableBookmarkRequest := _WritableBookmarkRequest{}

	err = json.Unmarshal(data, &varWritableBookmarkRequest)

	if err != nil {
		return err
	}

	*o = WritableBookmarkRequest(varWritableBookmarkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableBookmarkRequest struct {
	value *WritableBookmarkRequest
	isSet bool
}

func (v NullableWritableBookmarkRequest) Get() *WritableBookmarkRequest {
	return v.value
}

func (v *NullableWritableBookmarkRequest) Set(val *WritableBookmarkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableBookmarkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableBookmarkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableBookmarkRequest(val *WritableBookmarkRequest) *NullableWritableBookmarkRequest {
	return &NullableWritableBookmarkRequest{value: val, isSet: true}
}

func (v NullableWritableBookmarkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableBookmarkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
