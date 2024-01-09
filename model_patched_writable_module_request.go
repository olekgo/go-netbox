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

// checks if the PatchedWritableModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableModuleRequest{}

// PatchedWritableModuleRequest Adds support for custom fields and tags.
type PatchedWritableModuleRequest struct {
	Device     *int32             `json:"device,omitempty"`
	ModuleBay  *int32             `json:"module_bay,omitempty"`
	ModuleType *int32             `json:"module_type,omitempty"`
	Status     *ModuleStatusValue `json:"status,omitempty"`
	Serial     *string            `json:"serial,omitempty"`
	// A unique tag used to identify this device
	AssetTag             NullableString         `json:"asset_tag,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableModuleRequest PatchedWritableModuleRequest

// NewPatchedWritableModuleRequest instantiates a new PatchedWritableModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableModuleRequest() *PatchedWritableModuleRequest {
	this := PatchedWritableModuleRequest{}
	return &this
}

// NewPatchedWritableModuleRequestWithDefaults instantiates a new PatchedWritableModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableModuleRequestWithDefaults() *PatchedWritableModuleRequest {
	this := PatchedWritableModuleRequest{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedWritableModuleRequest) GetDevice() int32 {
	if o == nil || IsNil(o.Device) {
		var ret int32
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleRequest) GetDeviceOk() (*int32, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedWritableModuleRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given int32 and assigns it to the Device field.
func (o *PatchedWritableModuleRequest) SetDevice(v int32) {
	o.Device = &v
}

// GetModuleBay returns the ModuleBay field value if set, zero value otherwise.
func (o *PatchedWritableModuleRequest) GetModuleBay() int32 {
	if o == nil || IsNil(o.ModuleBay) {
		var ret int32
		return ret
	}
	return *o.ModuleBay
}

// GetModuleBayOk returns a tuple with the ModuleBay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleRequest) GetModuleBayOk() (*int32, bool) {
	if o == nil || IsNil(o.ModuleBay) {
		return nil, false
	}
	return o.ModuleBay, true
}

// HasModuleBay returns a boolean if a field has been set.
func (o *PatchedWritableModuleRequest) HasModuleBay() bool {
	if o != nil && !IsNil(o.ModuleBay) {
		return true
	}

	return false
}

// SetModuleBay gets a reference to the given int32 and assigns it to the ModuleBay field.
func (o *PatchedWritableModuleRequest) SetModuleBay(v int32) {
	o.ModuleBay = &v
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise.
func (o *PatchedWritableModuleRequest) GetModuleType() int32 {
	if o == nil || IsNil(o.ModuleType) {
		var ret int32
		return ret
	}
	return *o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleRequest) GetModuleTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.ModuleType) {
		return nil, false
	}
	return o.ModuleType, true
}

// HasModuleType returns a boolean if a field has been set.
func (o *PatchedWritableModuleRequest) HasModuleType() bool {
	if o != nil && !IsNil(o.ModuleType) {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given int32 and assigns it to the ModuleType field.
func (o *PatchedWritableModuleRequest) SetModuleType(v int32) {
	o.ModuleType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableModuleRequest) GetStatus() ModuleStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret ModuleStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleRequest) GetStatusOk() (*ModuleStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableModuleRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ModuleStatusValue and assigns it to the Status field.
func (o *PatchedWritableModuleRequest) SetStatus(v ModuleStatusValue) {
	o.Status = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *PatchedWritableModuleRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *PatchedWritableModuleRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *PatchedWritableModuleRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableModuleRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableModuleRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *PatchedWritableModuleRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *PatchedWritableModuleRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}

// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *PatchedWritableModuleRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *PatchedWritableModuleRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableModuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableModuleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableModuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableModuleRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableModuleRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableModuleRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableModuleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableModuleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableModuleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableModuleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableModuleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableModuleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.ModuleBay) {
		toSerialize["module_bay"] = o.ModuleBay
	}
	if !IsNil(o.ModuleType) {
		toSerialize["module_type"] = o.ModuleType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *PatchedWritableModuleRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableModuleRequest := _PatchedWritableModuleRequest{}

	err = json.Unmarshal(data, &varPatchedWritableModuleRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableModuleRequest(varPatchedWritableModuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module_bay")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableModuleRequest struct {
	value *PatchedWritableModuleRequest
	isSet bool
}

func (v NullablePatchedWritableModuleRequest) Get() *PatchedWritableModuleRequest {
	return v.value
}

func (v *NullablePatchedWritableModuleRequest) Set(val *PatchedWritableModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableModuleRequest(val *PatchedWritableModuleRequest) *NullablePatchedWritableModuleRequest {
	return &NullablePatchedWritableModuleRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
