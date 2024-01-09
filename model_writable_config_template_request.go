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

// checks if the WritableConfigTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableConfigTemplateRequest{}

// WritableConfigTemplateRequest Introduces support for Tag assignment. Adds `tags` serialization, and handles tag assignment on create() and update().
type WritableConfigTemplateRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Any <a href=\"https://jinja.palletsprojects.com/en/3.1.x/api/#jinja2.Environment\">additional parameters</a> to pass when constructing the Jinja2 environment.
	EnvironmentParams interface{} `json:"environment_params,omitempty"`
	// Jinja2 template code.
	TemplateCode string `json:"template_code"`
	// Remote data source
	DataSource           NullableInt32      `json:"data_source,omitempty"`
	DataFile             NullableInt32      `json:"data_file,omitempty"`
	Tags                 []NestedTagRequest `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableConfigTemplateRequest WritableConfigTemplateRequest

// NewWritableConfigTemplateRequest instantiates a new WritableConfigTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableConfigTemplateRequest(name string, templateCode string) *WritableConfigTemplateRequest {
	this := WritableConfigTemplateRequest{}
	this.Name = name
	this.TemplateCode = templateCode
	return &this
}

// NewWritableConfigTemplateRequestWithDefaults instantiates a new WritableConfigTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableConfigTemplateRequestWithDefaults() *WritableConfigTemplateRequest {
	this := WritableConfigTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableConfigTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableConfigTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableConfigTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableConfigTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableConfigTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableConfigTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironmentParams returns the EnvironmentParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableConfigTemplateRequest) GetEnvironmentParams() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnvironmentParams
}

// GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableConfigTemplateRequest) GetEnvironmentParamsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnvironmentParams) {
		return nil, false
	}
	return &o.EnvironmentParams, true
}

// HasEnvironmentParams returns a boolean if a field has been set.
func (o *WritableConfigTemplateRequest) HasEnvironmentParams() bool {
	if o != nil && IsNil(o.EnvironmentParams) {
		return true
	}

	return false
}

// SetEnvironmentParams gets a reference to the given interface{} and assigns it to the EnvironmentParams field.
func (o *WritableConfigTemplateRequest) SetEnvironmentParams(v interface{}) {
	o.EnvironmentParams = v
}

// GetTemplateCode returns the TemplateCode field value
func (o *WritableConfigTemplateRequest) GetTemplateCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateCode
}

// GetTemplateCodeOk returns a tuple with the TemplateCode field value
// and a boolean to check if the value has been set.
func (o *WritableConfigTemplateRequest) GetTemplateCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateCode, true
}

// SetTemplateCode sets field value
func (o *WritableConfigTemplateRequest) SetTemplateCode(v string) {
	o.TemplateCode = v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableConfigTemplateRequest) GetDataSource() int32 {
	if o == nil || IsNil(o.DataSource.Get()) {
		var ret int32
		return ret
	}
	return *o.DataSource.Get()
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableConfigTemplateRequest) GetDataSourceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSource.Get(), o.DataSource.IsSet()
}

// HasDataSource returns a boolean if a field has been set.
func (o *WritableConfigTemplateRequest) HasDataSource() bool {
	if o != nil && o.DataSource.IsSet() {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given NullableInt32 and assigns it to the DataSource field.
func (o *WritableConfigTemplateRequest) SetDataSource(v int32) {
	o.DataSource.Set(&v)
}

// SetDataSourceNil sets the value for DataSource to be an explicit nil
func (o *WritableConfigTemplateRequest) SetDataSourceNil() {
	o.DataSource.Set(nil)
}

// UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
func (o *WritableConfigTemplateRequest) UnsetDataSource() {
	o.DataSource.Unset()
}

// GetDataFile returns the DataFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableConfigTemplateRequest) GetDataFile() int32 {
	if o == nil || IsNil(o.DataFile.Get()) {
		var ret int32
		return ret
	}
	return *o.DataFile.Get()
}

// GetDataFileOk returns a tuple with the DataFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableConfigTemplateRequest) GetDataFileOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataFile.Get(), o.DataFile.IsSet()
}

// HasDataFile returns a boolean if a field has been set.
func (o *WritableConfigTemplateRequest) HasDataFile() bool {
	if o != nil && o.DataFile.IsSet() {
		return true
	}

	return false
}

// SetDataFile gets a reference to the given NullableInt32 and assigns it to the DataFile field.
func (o *WritableConfigTemplateRequest) SetDataFile(v int32) {
	o.DataFile.Set(&v)
}

// SetDataFileNil sets the value for DataFile to be an explicit nil
func (o *WritableConfigTemplateRequest) SetDataFileNil() {
	o.DataFile.Set(nil)
}

// UnsetDataFile ensures that no value is present for DataFile, not even an explicit nil
func (o *WritableConfigTemplateRequest) UnsetDataFile() {
	o.DataFile.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableConfigTemplateRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConfigTemplateRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableConfigTemplateRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableConfigTemplateRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

func (o WritableConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableConfigTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.EnvironmentParams != nil {
		toSerialize["environment_params"] = o.EnvironmentParams
	}
	toSerialize["template_code"] = o.TemplateCode
	if o.DataSource.IsSet() {
		toSerialize["data_source"] = o.DataSource.Get()
	}
	if o.DataFile.IsSet() {
		toSerialize["data_file"] = o.DataFile.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableConfigTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"template_code",
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

	varWritableConfigTemplateRequest := _WritableConfigTemplateRequest{}

	err = json.Unmarshal(data, &varWritableConfigTemplateRequest)

	if err != nil {
		return err
	}

	*o = WritableConfigTemplateRequest(varWritableConfigTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "environment_params")
		delete(additionalProperties, "template_code")
		delete(additionalProperties, "data_source")
		delete(additionalProperties, "data_file")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableConfigTemplateRequest struct {
	value *WritableConfigTemplateRequest
	isSet bool
}

func (v NullableWritableConfigTemplateRequest) Get() *WritableConfigTemplateRequest {
	return v.value
}

func (v *NullableWritableConfigTemplateRequest) Set(val *WritableConfigTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableConfigTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableConfigTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableConfigTemplateRequest(val *WritableConfigTemplateRequest) *NullableWritableConfigTemplateRequest {
	return &NullableWritableConfigTemplateRequest{value: val, isSet: true}
}

func (v NullableWritableConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableConfigTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
