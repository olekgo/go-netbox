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

// checks if the WritableIKEProposalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableIKEProposalRequest{}

// WritableIKEProposalRequest Adds support for custom fields and tags.
type WritableIKEProposalRequest struct {
	Name                    string                                                    `json:"name"`
	Description             *string                                                   `json:"description,omitempty"`
	AuthenticationMethod    IKEProposalAuthenticationMethodValue                      `json:"authentication_method"`
	EncryptionAlgorithm     IKEProposalEncryptionAlgorithmValue                       `json:"encryption_algorithm"`
	AuthenticationAlgorithm *PatchedWritableIKEProposalRequestAuthenticationAlgorithm `json:"authentication_algorithm,omitempty"`
	Group                   PatchedWritableIKEProposalRequestGroup                    `json:"group"`
	// Security association lifetime (in seconds)
	SaLifetime           NullableInt32          `json:"sa_lifetime,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableIKEProposalRequest WritableIKEProposalRequest

// NewWritableIKEProposalRequest instantiates a new WritableIKEProposalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableIKEProposalRequest(name string, authenticationMethod IKEProposalAuthenticationMethodValue, encryptionAlgorithm IKEProposalEncryptionAlgorithmValue, group PatchedWritableIKEProposalRequestGroup) *WritableIKEProposalRequest {
	this := WritableIKEProposalRequest{}
	this.Name = name
	this.AuthenticationMethod = authenticationMethod
	this.EncryptionAlgorithm = encryptionAlgorithm
	this.Group = group
	return &this
}

// NewWritableIKEProposalRequestWithDefaults instantiates a new WritableIKEProposalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableIKEProposalRequestWithDefaults() *WritableIKEProposalRequest {
	this := WritableIKEProposalRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableIKEProposalRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableIKEProposalRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableIKEProposalRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableIKEProposalRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIKEProposalRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableIKEProposalRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableIKEProposalRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value
func (o *WritableIKEProposalRequest) GetAuthenticationMethod() IKEProposalAuthenticationMethodValue {
	if o == nil {
		var ret IKEProposalAuthenticationMethodValue
		return ret
	}

	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *WritableIKEProposalRequest) GetAuthenticationMethodOk() (*IKEProposalAuthenticationMethodValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value
func (o *WritableIKEProposalRequest) SetAuthenticationMethod(v IKEProposalAuthenticationMethodValue) {
	o.AuthenticationMethod = v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value
func (o *WritableIKEProposalRequest) GetEncryptionAlgorithm() IKEProposalEncryptionAlgorithmValue {
	if o == nil {
		var ret IKEProposalEncryptionAlgorithmValue
		return ret
	}

	return o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value
// and a boolean to check if the value has been set.
func (o *WritableIKEProposalRequest) GetEncryptionAlgorithmOk() (*IKEProposalEncryptionAlgorithmValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionAlgorithm, true
}

// SetEncryptionAlgorithm sets field value
func (o *WritableIKEProposalRequest) SetEncryptionAlgorithm(v IKEProposalEncryptionAlgorithmValue) {
	o.EncryptionAlgorithm = v
}

// GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field value if set, zero value otherwise.
func (o *WritableIKEProposalRequest) GetAuthenticationAlgorithm() PatchedWritableIKEProposalRequestAuthenticationAlgorithm {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		var ret PatchedWritableIKEProposalRequestAuthenticationAlgorithm
		return ret
	}
	return *o.AuthenticationAlgorithm
}

// GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIKEProposalRequest) GetAuthenticationAlgorithmOk() (*PatchedWritableIKEProposalRequestAuthenticationAlgorithm, bool) {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		return nil, false
	}
	return o.AuthenticationAlgorithm, true
}

// HasAuthenticationAlgorithm returns a boolean if a field has been set.
func (o *WritableIKEProposalRequest) HasAuthenticationAlgorithm() bool {
	if o != nil && !IsNil(o.AuthenticationAlgorithm) {
		return true
	}

	return false
}

// SetAuthenticationAlgorithm gets a reference to the given PatchedWritableIKEProposalRequestAuthenticationAlgorithm and assigns it to the AuthenticationAlgorithm field.
func (o *WritableIKEProposalRequest) SetAuthenticationAlgorithm(v PatchedWritableIKEProposalRequestAuthenticationAlgorithm) {
	o.AuthenticationAlgorithm = &v
}

// GetGroup returns the Group field value
func (o *WritableIKEProposalRequest) GetGroup() PatchedWritableIKEProposalRequestGroup {
	if o == nil {
		var ret PatchedWritableIKEProposalRequestGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *WritableIKEProposalRequest) GetGroupOk() (*PatchedWritableIKEProposalRequestGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *WritableIKEProposalRequest) SetGroup(v PatchedWritableIKEProposalRequestGroup) {
	o.Group = v
}

// GetSaLifetime returns the SaLifetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIKEProposalRequest) GetSaLifetime() int32 {
	if o == nil || IsNil(o.SaLifetime.Get()) {
		var ret int32
		return ret
	}
	return *o.SaLifetime.Get()
}

// GetSaLifetimeOk returns a tuple with the SaLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIKEProposalRequest) GetSaLifetimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetime.Get(), o.SaLifetime.IsSet()
}

// HasSaLifetime returns a boolean if a field has been set.
func (o *WritableIKEProposalRequest) HasSaLifetime() bool {
	if o != nil && o.SaLifetime.IsSet() {
		return true
	}

	return false
}

// SetSaLifetime gets a reference to the given NullableInt32 and assigns it to the SaLifetime field.
func (o *WritableIKEProposalRequest) SetSaLifetime(v int32) {
	o.SaLifetime.Set(&v)
}

// SetSaLifetimeNil sets the value for SaLifetime to be an explicit nil
func (o *WritableIKEProposalRequest) SetSaLifetimeNil() {
	o.SaLifetime.Set(nil)
}

// UnsetSaLifetime ensures that no value is present for SaLifetime, not even an explicit nil
func (o *WritableIKEProposalRequest) UnsetSaLifetime() {
	o.SaLifetime.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableIKEProposalRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIKEProposalRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableIKEProposalRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableIKEProposalRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableIKEProposalRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIKEProposalRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableIKEProposalRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableIKEProposalRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableIKEProposalRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIKEProposalRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableIKEProposalRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableIKEProposalRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableIKEProposalRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableIKEProposalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["authentication_method"] = o.AuthenticationMethod
	toSerialize["encryption_algorithm"] = o.EncryptionAlgorithm
	if !IsNil(o.AuthenticationAlgorithm) {
		toSerialize["authentication_algorithm"] = o.AuthenticationAlgorithm
	}
	toSerialize["group"] = o.Group
	if o.SaLifetime.IsSet() {
		toSerialize["sa_lifetime"] = o.SaLifetime.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
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

func (o *WritableIKEProposalRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"authentication_method",
		"encryption_algorithm",
		"group",
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

	varWritableIKEProposalRequest := _WritableIKEProposalRequest{}

	err = json.Unmarshal(data, &varWritableIKEProposalRequest)

	if err != nil {
		return err
	}

	*o = WritableIKEProposalRequest(varWritableIKEProposalRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "authentication_method")
		delete(additionalProperties, "encryption_algorithm")
		delete(additionalProperties, "authentication_algorithm")
		delete(additionalProperties, "group")
		delete(additionalProperties, "sa_lifetime")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableIKEProposalRequest struct {
	value *WritableIKEProposalRequest
	isSet bool
}

func (v NullableWritableIKEProposalRequest) Get() *WritableIKEProposalRequest {
	return v.value
}

func (v *NullableWritableIKEProposalRequest) Set(val *WritableIKEProposalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableIKEProposalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableIKEProposalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableIKEProposalRequest(val *WritableIKEProposalRequest) *NullableWritableIKEProposalRequest {
	return &NullableWritableIKEProposalRequest{value: val, isSet: true}
}

func (v NullableWritableIKEProposalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableIKEProposalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
