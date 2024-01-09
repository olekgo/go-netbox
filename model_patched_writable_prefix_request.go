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

// checks if the PatchedWritablePrefixRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritablePrefixRequest{}

// PatchedWritablePrefixRequest Adds support for custom fields and tags.
type PatchedWritablePrefixRequest struct {
	Prefix *string                             `json:"prefix,omitempty"`
	Site   NullableInt32                       `json:"site,omitempty"`
	Vrf    NullableInt32                       `json:"vrf,omitempty"`
	Tenant NullableInt32                       `json:"tenant,omitempty"`
	Vlan   NullableInt32                       `json:"vlan,omitempty"`
	Status *PatchedWritablePrefixRequestStatus `json:"status,omitempty"`
	// The primary function of this prefix
	Role NullableInt32 `json:"role,omitempty"`
	// All IP addresses within this prefix are considered usable
	IsPool *bool `json:"is_pool,omitempty"`
	// Treat as 100% utilized
	MarkUtilized         *bool                  `json:"mark_utilized,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritablePrefixRequest PatchedWritablePrefixRequest

// NewPatchedWritablePrefixRequest instantiates a new PatchedWritablePrefixRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritablePrefixRequest() *PatchedWritablePrefixRequest {
	this := PatchedWritablePrefixRequest{}
	return &this
}

// NewPatchedWritablePrefixRequestWithDefaults instantiates a new PatchedWritablePrefixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritablePrefixRequestWithDefaults() *PatchedWritablePrefixRequest {
	this := PatchedWritablePrefixRequest{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *PatchedWritablePrefixRequest) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePrefixRequest) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *PatchedWritablePrefixRequest) SetPrefix(v string) {
	o.Prefix = &v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePrefixRequest) GetSite() int32 {
	if o == nil || IsNil(o.Site.Get()) {
		var ret int32
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePrefixRequest) GetSiteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableInt32 and assigns it to the Site field.
func (o *PatchedWritablePrefixRequest) SetSite(v int32) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *PatchedWritablePrefixRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *PatchedWritablePrefixRequest) UnsetSite() {
	o.Site.Unset()
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePrefixRequest) GetVrf() int32 {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret int32
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePrefixRequest) GetVrfOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableInt32 and assigns it to the Vrf field.
func (o *PatchedWritablePrefixRequest) SetVrf(v int32) {
	o.Vrf.Set(&v)
}

// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *PatchedWritablePrefixRequest) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *PatchedWritablePrefixRequest) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePrefixRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePrefixRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *PatchedWritablePrefixRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritablePrefixRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritablePrefixRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetVlan returns the Vlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePrefixRequest) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan.Get()) {
		var ret int32
		return ret
	}
	return *o.Vlan.Get()
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePrefixRequest) GetVlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vlan.Get(), o.Vlan.IsSet()
}

// HasVlan returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasVlan() bool {
	if o != nil && o.Vlan.IsSet() {
		return true
	}

	return false
}

// SetVlan gets a reference to the given NullableInt32 and assigns it to the Vlan field.
func (o *PatchedWritablePrefixRequest) SetVlan(v int32) {
	o.Vlan.Set(&v)
}

// SetVlanNil sets the value for Vlan to be an explicit nil
func (o *PatchedWritablePrefixRequest) SetVlanNil() {
	o.Vlan.Set(nil)
}

// UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
func (o *PatchedWritablePrefixRequest) UnsetVlan() {
	o.Vlan.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritablePrefixRequest) GetStatus() PatchedWritablePrefixRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritablePrefixRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePrefixRequest) GetStatusOk() (*PatchedWritablePrefixRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritablePrefixRequestStatus and assigns it to the Status field.
func (o *PatchedWritablePrefixRequest) SetStatus(v PatchedWritablePrefixRequestStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePrefixRequest) GetRole() int32 {
	if o == nil || IsNil(o.Role.Get()) {
		var ret int32
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePrefixRequest) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableInt32 and assigns it to the Role field.
func (o *PatchedWritablePrefixRequest) SetRole(v int32) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *PatchedWritablePrefixRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PatchedWritablePrefixRequest) UnsetRole() {
	o.Role.Unset()
}

// GetIsPool returns the IsPool field value if set, zero value otherwise.
func (o *PatchedWritablePrefixRequest) GetIsPool() bool {
	if o == nil || IsNil(o.IsPool) {
		var ret bool
		return ret
	}
	return *o.IsPool
}

// GetIsPoolOk returns a tuple with the IsPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePrefixRequest) GetIsPoolOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPool) {
		return nil, false
	}
	return o.IsPool, true
}

// HasIsPool returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasIsPool() bool {
	if o != nil && !IsNil(o.IsPool) {
		return true
	}

	return false
}

// SetIsPool gets a reference to the given bool and assigns it to the IsPool field.
func (o *PatchedWritablePrefixRequest) SetIsPool(v bool) {
	o.IsPool = &v
}

// GetMarkUtilized returns the MarkUtilized field value if set, zero value otherwise.
func (o *PatchedWritablePrefixRequest) GetMarkUtilized() bool {
	if o == nil || IsNil(o.MarkUtilized) {
		var ret bool
		return ret
	}
	return *o.MarkUtilized
}

// GetMarkUtilizedOk returns a tuple with the MarkUtilized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePrefixRequest) GetMarkUtilizedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkUtilized) {
		return nil, false
	}
	return o.MarkUtilized, true
}

// HasMarkUtilized returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasMarkUtilized() bool {
	if o != nil && !IsNil(o.MarkUtilized) {
		return true
	}

	return false
}

// SetMarkUtilized gets a reference to the given bool and assigns it to the MarkUtilized field.
func (o *PatchedWritablePrefixRequest) SetMarkUtilized(v bool) {
	o.MarkUtilized = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritablePrefixRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePrefixRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritablePrefixRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritablePrefixRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePrefixRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritablePrefixRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritablePrefixRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePrefixRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritablePrefixRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritablePrefixRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePrefixRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritablePrefixRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritablePrefixRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritablePrefixRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritablePrefixRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Vlan.IsSet() {
		toSerialize["vlan"] = o.Vlan.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if !IsNil(o.IsPool) {
		toSerialize["is_pool"] = o.IsPool
	}
	if !IsNil(o.MarkUtilized) {
		toSerialize["mark_utilized"] = o.MarkUtilized
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

func (o *PatchedWritablePrefixRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritablePrefixRequest := _PatchedWritablePrefixRequest{}

	err = json.Unmarshal(data, &varPatchedWritablePrefixRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritablePrefixRequest(varPatchedWritablePrefixRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "site")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "vlan")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "is_pool")
		delete(additionalProperties, "mark_utilized")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritablePrefixRequest struct {
	value *PatchedWritablePrefixRequest
	isSet bool
}

func (v NullablePatchedWritablePrefixRequest) Get() *PatchedWritablePrefixRequest {
	return v.value
}

func (v *NullablePatchedWritablePrefixRequest) Set(val *PatchedWritablePrefixRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePrefixRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePrefixRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePrefixRequest(val *PatchedWritablePrefixRequest) *NullablePatchedWritablePrefixRequest {
	return &NullablePatchedWritablePrefixRequest{value: val, isSet: true}
}

func (v NullablePatchedWritablePrefixRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePrefixRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
