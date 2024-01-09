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

// checks if the PaginatedL2VPNTerminationList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedL2VPNTerminationList{}

// PaginatedL2VPNTerminationList struct for PaginatedL2VPNTerminationList
type PaginatedL2VPNTerminationList struct {
	Count                *int32             `json:"count,omitempty"`
	Next                 NullableString     `json:"next,omitempty"`
	Previous             NullableString     `json:"previous,omitempty"`
	Results              []L2VPNTermination `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedL2VPNTerminationList PaginatedL2VPNTerminationList

// NewPaginatedL2VPNTerminationList instantiates a new PaginatedL2VPNTerminationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedL2VPNTerminationList() *PaginatedL2VPNTerminationList {
	this := PaginatedL2VPNTerminationList{}
	return &this
}

// NewPaginatedL2VPNTerminationListWithDefaults instantiates a new PaginatedL2VPNTerminationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedL2VPNTerminationListWithDefaults() *PaginatedL2VPNTerminationList {
	this := PaginatedL2VPNTerminationList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedL2VPNTerminationList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedL2VPNTerminationList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedL2VPNTerminationList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedL2VPNTerminationList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedL2VPNTerminationList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedL2VPNTerminationList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedL2VPNTerminationList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedL2VPNTerminationList) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedL2VPNTerminationList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedL2VPNTerminationList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedL2VPNTerminationList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedL2VPNTerminationList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedL2VPNTerminationList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedL2VPNTerminationList) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedL2VPNTerminationList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedL2VPNTerminationList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedL2VPNTerminationList) GetResults() []L2VPNTermination {
	if o == nil || IsNil(o.Results) {
		var ret []L2VPNTermination
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedL2VPNTerminationList) GetResultsOk() ([]L2VPNTermination, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedL2VPNTerminationList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []L2VPNTermination and assigns it to the Results field.
func (o *PaginatedL2VPNTerminationList) SetResults(v []L2VPNTermination) {
	o.Results = v
}

func (o PaginatedL2VPNTerminationList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedL2VPNTerminationList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedL2VPNTerminationList) UnmarshalJSON(data []byte) (err error) {
	varPaginatedL2VPNTerminationList := _PaginatedL2VPNTerminationList{}

	err = json.Unmarshal(data, &varPaginatedL2VPNTerminationList)

	if err != nil {
		return err
	}

	*o = PaginatedL2VPNTerminationList(varPaginatedL2VPNTerminationList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedL2VPNTerminationList struct {
	value *PaginatedL2VPNTerminationList
	isSet bool
}

func (v NullablePaginatedL2VPNTerminationList) Get() *PaginatedL2VPNTerminationList {
	return v.value
}

func (v *NullablePaginatedL2VPNTerminationList) Set(val *PaginatedL2VPNTerminationList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedL2VPNTerminationList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedL2VPNTerminationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedL2VPNTerminationList(val *PaginatedL2VPNTerminationList) *NullablePaginatedL2VPNTerminationList {
	return &NullablePaginatedL2VPNTerminationList{value: val, isSet: true}
}

func (v NullablePaginatedL2VPNTerminationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedL2VPNTerminationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
