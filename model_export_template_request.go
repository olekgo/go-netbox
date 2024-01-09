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

// checks if the ExportTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportTemplateRequest{}

// ExportTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ExportTemplateRequest struct {
	ContentTypes []string `json:"content_types"`
	Name         string   `json:"name"`
	Description  *string  `json:"description,omitempty"`
	// Jinja2 template code. The list of objects being exported is passed as a context variable named <code>queryset</code>.
	TemplateCode string `json:"template_code"`
	// Defaults to <code>text/plain; charset=utf-8</code>
	MimeType *string `json:"mime_type,omitempty"`
	// Extension to append to the rendered filename
	FileExtension *string `json:"file_extension,omitempty"`
	// Download file as attachment
	AsAttachment         *bool                    `json:"as_attachment,omitempty"`
	DataSource           *NestedDataSourceRequest `json:"data_source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExportTemplateRequest ExportTemplateRequest

// NewExportTemplateRequest instantiates a new ExportTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportTemplateRequest(contentTypes []string, name string, templateCode string) *ExportTemplateRequest {
	this := ExportTemplateRequest{}
	this.ContentTypes = contentTypes
	this.Name = name
	this.TemplateCode = templateCode
	return &this
}

// NewExportTemplateRequestWithDefaults instantiates a new ExportTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportTemplateRequestWithDefaults() *ExportTemplateRequest {
	this := ExportTemplateRequest{}
	return &this
}

// GetContentTypes returns the ContentTypes field value
func (o *ExportTemplateRequest) GetContentTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value
// and a boolean to check if the value has been set.
func (o *ExportTemplateRequest) GetContentTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// SetContentTypes sets field value
func (o *ExportTemplateRequest) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value
func (o *ExportTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExportTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExportTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExportTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExportTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExportTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTemplateCode returns the TemplateCode field value
func (o *ExportTemplateRequest) GetTemplateCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateCode
}

// GetTemplateCodeOk returns a tuple with the TemplateCode field value
// and a boolean to check if the value has been set.
func (o *ExportTemplateRequest) GetTemplateCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateCode, true
}

// SetTemplateCode sets field value
func (o *ExportTemplateRequest) SetTemplateCode(v string) {
	o.TemplateCode = v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *ExportTemplateRequest) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplateRequest) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *ExportTemplateRequest) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *ExportTemplateRequest) SetMimeType(v string) {
	o.MimeType = &v
}

// GetFileExtension returns the FileExtension field value if set, zero value otherwise.
func (o *ExportTemplateRequest) GetFileExtension() string {
	if o == nil || IsNil(o.FileExtension) {
		var ret string
		return ret
	}
	return *o.FileExtension
}

// GetFileExtensionOk returns a tuple with the FileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplateRequest) GetFileExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.FileExtension) {
		return nil, false
	}
	return o.FileExtension, true
}

// HasFileExtension returns a boolean if a field has been set.
func (o *ExportTemplateRequest) HasFileExtension() bool {
	if o != nil && !IsNil(o.FileExtension) {
		return true
	}

	return false
}

// SetFileExtension gets a reference to the given string and assigns it to the FileExtension field.
func (o *ExportTemplateRequest) SetFileExtension(v string) {
	o.FileExtension = &v
}

// GetAsAttachment returns the AsAttachment field value if set, zero value otherwise.
func (o *ExportTemplateRequest) GetAsAttachment() bool {
	if o == nil || IsNil(o.AsAttachment) {
		var ret bool
		return ret
	}
	return *o.AsAttachment
}

// GetAsAttachmentOk returns a tuple with the AsAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplateRequest) GetAsAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.AsAttachment) {
		return nil, false
	}
	return o.AsAttachment, true
}

// HasAsAttachment returns a boolean if a field has been set.
func (o *ExportTemplateRequest) HasAsAttachment() bool {
	if o != nil && !IsNil(o.AsAttachment) {
		return true
	}

	return false
}

// SetAsAttachment gets a reference to the given bool and assigns it to the AsAttachment field.
func (o *ExportTemplateRequest) SetAsAttachment(v bool) {
	o.AsAttachment = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ExportTemplateRequest) GetDataSource() NestedDataSourceRequest {
	if o == nil || IsNil(o.DataSource) {
		var ret NestedDataSourceRequest
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplateRequest) GetDataSourceOk() (*NestedDataSourceRequest, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ExportTemplateRequest) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given NestedDataSourceRequest and assigns it to the DataSource field.
func (o *ExportTemplateRequest) SetDataSource(v NestedDataSourceRequest) {
	o.DataSource = &v
}

func (o ExportTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content_types"] = o.ContentTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["template_code"] = o.TemplateCode
	if !IsNil(o.MimeType) {
		toSerialize["mime_type"] = o.MimeType
	}
	if !IsNil(o.FileExtension) {
		toSerialize["file_extension"] = o.FileExtension
	}
	if !IsNil(o.AsAttachment) {
		toSerialize["as_attachment"] = o.AsAttachment
	}
	if !IsNil(o.DataSource) {
		toSerialize["data_source"] = o.DataSource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExportTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content_types",
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

	varExportTemplateRequest := _ExportTemplateRequest{}

	err = json.Unmarshal(data, &varExportTemplateRequest)

	if err != nil {
		return err
	}

	*o = ExportTemplateRequest(varExportTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "template_code")
		delete(additionalProperties, "mime_type")
		delete(additionalProperties, "file_extension")
		delete(additionalProperties, "as_attachment")
		delete(additionalProperties, "data_source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportTemplateRequest struct {
	value *ExportTemplateRequest
	isSet bool
}

func (v NullableExportTemplateRequest) Get() *ExportTemplateRequest {
	return v.value
}

func (v *NullableExportTemplateRequest) Set(val *ExportTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExportTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExportTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportTemplateRequest(val *ExportTemplateRequest) *NullableExportTemplateRequest {
	return &NullableExportTemplateRequest{value: val, isSet: true}
}

func (v NullableExportTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
