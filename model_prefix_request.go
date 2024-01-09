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

// checks if the PrefixRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrefixRequest{}

// PrefixRequest Adds support for custom fields and tags.
type PrefixRequest struct {
	Prefix string                      `json:"prefix"`
	Site   NullableNestedSiteRequest   `json:"site,omitempty"`
	Vrf    NullableNestedVRFRequest    `json:"vrf,omitempty"`
	Tenant NullableNestedTenantRequest `json:"tenant,omitempty"`
	Vlan   NullableNestedVLANRequest   `json:"vlan,omitempty"`
	Status *PrefixStatusValue          `json:"status,omitempty"`
	Role   NullableNestedRoleRequest   `json:"role,omitempty"`
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

type _PrefixRequest PrefixRequest

// NewPrefixRequest instantiates a new PrefixRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrefixRequest(prefix string) *PrefixRequest {
	this := PrefixRequest{}
	this.Prefix = prefix
	return &this
}

// NewPrefixRequestWithDefaults instantiates a new PrefixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrefixRequestWithDefaults() *PrefixRequest {
	this := PrefixRequest{}
	return &this
}

// GetPrefix returns the Prefix field value
func (o *PrefixRequest) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *PrefixRequest) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *PrefixRequest) SetPrefix(v string) {
	o.Prefix = v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrefixRequest) GetSite() NestedSiteRequest {
	if o == nil || IsNil(o.Site.Get()) {
		var ret NestedSiteRequest
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrefixRequest) GetSiteOk() (*NestedSiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *PrefixRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableNestedSiteRequest and assigns it to the Site field.
func (o *PrefixRequest) SetSite(v NestedSiteRequest) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *PrefixRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *PrefixRequest) UnsetSite() {
	o.Site.Unset()
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrefixRequest) GetVrf() NestedVRFRequest {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret NestedVRFRequest
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrefixRequest) GetVrfOk() (*NestedVRFRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *PrefixRequest) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableNestedVRFRequest and assigns it to the Vrf field.
func (o *PrefixRequest) SetVrf(v NestedVRFRequest) {
	o.Vrf.Set(&v)
}

// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *PrefixRequest) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *PrefixRequest) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrefixRequest) GetTenant() NestedTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret NestedTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrefixRequest) GetTenantOk() (*NestedTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PrefixRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableNestedTenantRequest and assigns it to the Tenant field.
func (o *PrefixRequest) SetTenant(v NestedTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PrefixRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PrefixRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetVlan returns the Vlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrefixRequest) GetVlan() NestedVLANRequest {
	if o == nil || IsNil(o.Vlan.Get()) {
		var ret NestedVLANRequest
		return ret
	}
	return *o.Vlan.Get()
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrefixRequest) GetVlanOk() (*NestedVLANRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vlan.Get(), o.Vlan.IsSet()
}

// HasVlan returns a boolean if a field has been set.
func (o *PrefixRequest) HasVlan() bool {
	if o != nil && o.Vlan.IsSet() {
		return true
	}

	return false
}

// SetVlan gets a reference to the given NullableNestedVLANRequest and assigns it to the Vlan field.
func (o *PrefixRequest) SetVlan(v NestedVLANRequest) {
	o.Vlan.Set(&v)
}

// SetVlanNil sets the value for Vlan to be an explicit nil
func (o *PrefixRequest) SetVlanNil() {
	o.Vlan.Set(nil)
}

// UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
func (o *PrefixRequest) UnsetVlan() {
	o.Vlan.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PrefixRequest) GetStatus() PrefixStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret PrefixStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixRequest) GetStatusOk() (*PrefixStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PrefixRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PrefixStatusValue and assigns it to the Status field.
func (o *PrefixRequest) SetStatus(v PrefixStatusValue) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrefixRequest) GetRole() NestedRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret NestedRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrefixRequest) GetRoleOk() (*NestedRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PrefixRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableNestedRoleRequest and assigns it to the Role field.
func (o *PrefixRequest) SetRole(v NestedRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *PrefixRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PrefixRequest) UnsetRole() {
	o.Role.Unset()
}

// GetIsPool returns the IsPool field value if set, zero value otherwise.
func (o *PrefixRequest) GetIsPool() bool {
	if o == nil || IsNil(o.IsPool) {
		var ret bool
		return ret
	}
	return *o.IsPool
}

// GetIsPoolOk returns a tuple with the IsPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixRequest) GetIsPoolOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPool) {
		return nil, false
	}
	return o.IsPool, true
}

// HasIsPool returns a boolean if a field has been set.
func (o *PrefixRequest) HasIsPool() bool {
	if o != nil && !IsNil(o.IsPool) {
		return true
	}

	return false
}

// SetIsPool gets a reference to the given bool and assigns it to the IsPool field.
func (o *PrefixRequest) SetIsPool(v bool) {
	o.IsPool = &v
}

// GetMarkUtilized returns the MarkUtilized field value if set, zero value otherwise.
func (o *PrefixRequest) GetMarkUtilized() bool {
	if o == nil || IsNil(o.MarkUtilized) {
		var ret bool
		return ret
	}
	return *o.MarkUtilized
}

// GetMarkUtilizedOk returns a tuple with the MarkUtilized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixRequest) GetMarkUtilizedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkUtilized) {
		return nil, false
	}
	return o.MarkUtilized, true
}

// HasMarkUtilized returns a boolean if a field has been set.
func (o *PrefixRequest) HasMarkUtilized() bool {
	if o != nil && !IsNil(o.MarkUtilized) {
		return true
	}

	return false
}

// SetMarkUtilized gets a reference to the given bool and assigns it to the MarkUtilized field.
func (o *PrefixRequest) SetMarkUtilized(v bool) {
	o.MarkUtilized = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PrefixRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PrefixRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PrefixRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PrefixRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PrefixRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PrefixRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PrefixRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PrefixRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PrefixRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PrefixRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PrefixRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PrefixRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PrefixRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrefixRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prefix"] = o.Prefix
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

func (o *PrefixRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prefix",
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

	varPrefixRequest := _PrefixRequest{}

	err = json.Unmarshal(data, &varPrefixRequest)

	if err != nil {
		return err
	}

	*o = PrefixRequest(varPrefixRequest)

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

type NullablePrefixRequest struct {
	value *PrefixRequest
	isSet bool
}

func (v NullablePrefixRequest) Get() *PrefixRequest {
	return v.value
}

func (v *NullablePrefixRequest) Set(val *PrefixRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePrefixRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePrefixRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrefixRequest(val *PrefixRequest) *NullablePrefixRequest {
	return &NullablePrefixRequest{value: val, isSet: true}
}

func (v NullablePrefixRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrefixRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
