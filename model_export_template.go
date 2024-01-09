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
	"time"
)

// checks if the ExportTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportTemplate{}

// ExportTemplate Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ExportTemplate struct {
	Id           int32    `json:"id"`
	Url          string   `json:"url"`
	Display      string   `json:"display"`
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
	AsAttachment *bool             `json:"as_attachment,omitempty"`
	DataSource   *NestedDataSource `json:"data_source,omitempty"`
	// Path to remote file (relative to data source root)
	DataPath             string         `json:"data_path"`
	DataFile             NestedDataFile `json:"data_file"`
	DataSynced           NullableTime   `json:"data_synced"`
	Created              NullableTime   `json:"created"`
	LastUpdated          NullableTime   `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _ExportTemplate ExportTemplate

// NewExportTemplate instantiates a new ExportTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportTemplate(id int32, url string, display string, contentTypes []string, name string, templateCode string, dataPath string, dataFile NestedDataFile, dataSynced NullableTime, created NullableTime, lastUpdated NullableTime) *ExportTemplate {
	this := ExportTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.ContentTypes = contentTypes
	this.Name = name
	this.TemplateCode = templateCode
	this.DataPath = dataPath
	this.DataFile = dataFile
	this.DataSynced = dataSynced
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewExportTemplateWithDefaults instantiates a new ExportTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportTemplateWithDefaults() *ExportTemplate {
	this := ExportTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *ExportTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExportTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ExportTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ExportTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *ExportTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ExportTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetContentTypes returns the ContentTypes field value
func (o *ExportTemplate) GetContentTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetContentTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// SetContentTypes sets field value
func (o *ExportTemplate) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value
func (o *ExportTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExportTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExportTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExportTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExportTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetTemplateCode returns the TemplateCode field value
func (o *ExportTemplate) GetTemplateCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateCode
}

// GetTemplateCodeOk returns a tuple with the TemplateCode field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetTemplateCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateCode, true
}

// SetTemplateCode sets field value
func (o *ExportTemplate) SetTemplateCode(v string) {
	o.TemplateCode = v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *ExportTemplate) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *ExportTemplate) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *ExportTemplate) SetMimeType(v string) {
	o.MimeType = &v
}

// GetFileExtension returns the FileExtension field value if set, zero value otherwise.
func (o *ExportTemplate) GetFileExtension() string {
	if o == nil || IsNil(o.FileExtension) {
		var ret string
		return ret
	}
	return *o.FileExtension
}

// GetFileExtensionOk returns a tuple with the FileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetFileExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.FileExtension) {
		return nil, false
	}
	return o.FileExtension, true
}

// HasFileExtension returns a boolean if a field has been set.
func (o *ExportTemplate) HasFileExtension() bool {
	if o != nil && !IsNil(o.FileExtension) {
		return true
	}

	return false
}

// SetFileExtension gets a reference to the given string and assigns it to the FileExtension field.
func (o *ExportTemplate) SetFileExtension(v string) {
	o.FileExtension = &v
}

// GetAsAttachment returns the AsAttachment field value if set, zero value otherwise.
func (o *ExportTemplate) GetAsAttachment() bool {
	if o == nil || IsNil(o.AsAttachment) {
		var ret bool
		return ret
	}
	return *o.AsAttachment
}

// GetAsAttachmentOk returns a tuple with the AsAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetAsAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.AsAttachment) {
		return nil, false
	}
	return o.AsAttachment, true
}

// HasAsAttachment returns a boolean if a field has been set.
func (o *ExportTemplate) HasAsAttachment() bool {
	if o != nil && !IsNil(o.AsAttachment) {
		return true
	}

	return false
}

// SetAsAttachment gets a reference to the given bool and assigns it to the AsAttachment field.
func (o *ExportTemplate) SetAsAttachment(v bool) {
	o.AsAttachment = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ExportTemplate) GetDataSource() NestedDataSource {
	if o == nil || IsNil(o.DataSource) {
		var ret NestedDataSource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetDataSourceOk() (*NestedDataSource, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ExportTemplate) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given NestedDataSource and assigns it to the DataSource field.
func (o *ExportTemplate) SetDataSource(v NestedDataSource) {
	o.DataSource = &v
}

// GetDataPath returns the DataPath field value
func (o *ExportTemplate) GetDataPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataPath
}

// GetDataPathOk returns a tuple with the DataPath field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetDataPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataPath, true
}

// SetDataPath sets field value
func (o *ExportTemplate) SetDataPath(v string) {
	o.DataPath = v
}

// GetDataFile returns the DataFile field value
func (o *ExportTemplate) GetDataFile() NestedDataFile {
	if o == nil {
		var ret NestedDataFile
		return ret
	}

	return o.DataFile
}

// GetDataFileOk returns a tuple with the DataFile field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetDataFileOk() (*NestedDataFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataFile, true
}

// SetDataFile sets field value
func (o *ExportTemplate) SetDataFile(v NestedDataFile) {
	o.DataFile = v
}

// GetDataSynced returns the DataSynced field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ExportTemplate) GetDataSynced() time.Time {
	if o == nil || o.DataSynced.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DataSynced.Get()
}

// GetDataSyncedOk returns a tuple with the DataSynced field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportTemplate) GetDataSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSynced.Get(), o.DataSynced.IsSet()
}

// SetDataSynced sets field value
func (o *ExportTemplate) SetDataSynced(v time.Time) {
	o.DataSynced.Set(&v)
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ExportTemplate) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportTemplate) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ExportTemplate) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ExportTemplate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportTemplate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ExportTemplate) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o ExportTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
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
	toSerialize["data_path"] = o.DataPath
	toSerialize["data_file"] = o.DataFile
	toSerialize["data_synced"] = o.DataSynced.Get()
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExportTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"content_types",
		"name",
		"template_code",
		"data_path",
		"data_file",
		"data_synced",
		"created",
		"last_updated",
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

	varExportTemplate := _ExportTemplate{}

	err = json.Unmarshal(data, &varExportTemplate)

	if err != nil {
		return err
	}

	*o = ExportTemplate(varExportTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "template_code")
		delete(additionalProperties, "mime_type")
		delete(additionalProperties, "file_extension")
		delete(additionalProperties, "as_attachment")
		delete(additionalProperties, "data_source")
		delete(additionalProperties, "data_path")
		delete(additionalProperties, "data_file")
		delete(additionalProperties, "data_synced")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportTemplate struct {
	value *ExportTemplate
	isSet bool
}

func (v NullableExportTemplate) Get() *ExportTemplate {
	return v.value
}

func (v *NullableExportTemplate) Set(val *ExportTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableExportTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableExportTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportTemplate(val *ExportTemplate) *NullableExportTemplate {
	return &NullableExportTemplate{value: val, isSet: true}
}

func (v NullableExportTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
