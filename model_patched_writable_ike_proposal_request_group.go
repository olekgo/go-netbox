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

// PatchedWritableIKEProposalRequestGroup Diffie-Hellman group ID  * `1` - Group 1 * `2` - Group 2 * `5` - Group 5 * `14` - Group 14 * `16` - Group 16 * `17` - Group 17 * `18` - Group 18 * `19` - Group 19 * `20` - Group 20 * `21` - Group 21 * `22` - Group 22 * `23` - Group 23 * `24` - Group 24 * `25` - Group 25 * `26` - Group 26 * `27` - Group 27 * `28` - Group 28 * `29` - Group 29 * `30` - Group 30 * `31` - Group 31 * `32` - Group 32 * `33` - Group 33 * `34` - Group 34
type PatchedWritableIKEProposalRequestGroup int32

// List of PatchedWritableIKEProposalRequest_group
const (
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__1  PatchedWritableIKEProposalRequestGroup = 1
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__2  PatchedWritableIKEProposalRequestGroup = 2
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__5  PatchedWritableIKEProposalRequestGroup = 5
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__14 PatchedWritableIKEProposalRequestGroup = 14
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__16 PatchedWritableIKEProposalRequestGroup = 16
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__17 PatchedWritableIKEProposalRequestGroup = 17
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__18 PatchedWritableIKEProposalRequestGroup = 18
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__19 PatchedWritableIKEProposalRequestGroup = 19
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__20 PatchedWritableIKEProposalRequestGroup = 20
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__21 PatchedWritableIKEProposalRequestGroup = 21
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__22 PatchedWritableIKEProposalRequestGroup = 22
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__23 PatchedWritableIKEProposalRequestGroup = 23
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__24 PatchedWritableIKEProposalRequestGroup = 24
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__25 PatchedWritableIKEProposalRequestGroup = 25
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__26 PatchedWritableIKEProposalRequestGroup = 26
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__27 PatchedWritableIKEProposalRequestGroup = 27
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__28 PatchedWritableIKEProposalRequestGroup = 28
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__29 PatchedWritableIKEProposalRequestGroup = 29
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__30 PatchedWritableIKEProposalRequestGroup = 30
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__31 PatchedWritableIKEProposalRequestGroup = 31
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__32 PatchedWritableIKEProposalRequestGroup = 32
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__33 PatchedWritableIKEProposalRequestGroup = 33
	PATCHEDWRITABLEIKEPROPOSALREQUESTGROUP__34 PatchedWritableIKEProposalRequestGroup = 34
)

// All allowed values of PatchedWritableIKEProposalRequestGroup enum
var AllowedPatchedWritableIKEProposalRequestGroupEnumValues = []PatchedWritableIKEProposalRequestGroup{
	1,
	2,
	5,
	14,
	16,
	17,
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	25,
	26,
	27,
	28,
	29,
	30,
	31,
	32,
	33,
	34,
}

func (v *PatchedWritableIKEProposalRequestGroup) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableIKEProposalRequestGroup(value)
	for _, existing := range AllowedPatchedWritableIKEProposalRequestGroupEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableIKEProposalRequestGroup", value)
}

// NewPatchedWritableIKEProposalRequestGroupFromValue returns a pointer to a valid PatchedWritableIKEProposalRequestGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableIKEProposalRequestGroupFromValue(v int32) (*PatchedWritableIKEProposalRequestGroup, error) {
	ev := PatchedWritableIKEProposalRequestGroup(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableIKEProposalRequestGroup: valid values are %v", v, AllowedPatchedWritableIKEProposalRequestGroupEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableIKEProposalRequestGroup) IsValid() bool {
	for _, existing := range AllowedPatchedWritableIKEProposalRequestGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableIKEProposalRequest_group value
func (v PatchedWritableIKEProposalRequestGroup) Ptr() *PatchedWritableIKEProposalRequestGroup {
	return &v
}

type NullablePatchedWritableIKEProposalRequestGroup struct {
	value *PatchedWritableIKEProposalRequestGroup
	isSet bool
}

func (v NullablePatchedWritableIKEProposalRequestGroup) Get() *PatchedWritableIKEProposalRequestGroup {
	return v.value
}

func (v *NullablePatchedWritableIKEProposalRequestGroup) Set(val *PatchedWritableIKEProposalRequestGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableIKEProposalRequestGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableIKEProposalRequestGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableIKEProposalRequestGroup(val *PatchedWritableIKEProposalRequestGroup) *NullablePatchedWritableIKEProposalRequestGroup {
	return &NullablePatchedWritableIKEProposalRequestGroup{value: val, isSet: true}
}

func (v NullablePatchedWritableIKEProposalRequestGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableIKEProposalRequestGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
