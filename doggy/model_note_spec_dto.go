/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
	"time"
)

// NoteSpecDto struct for NoteSpecDto
type NoteSpecDto struct {
	Id *string `json:"id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Content NullableString `json:"content,omitempty"`
	Version *int32 `json:"version,omitempty"`
	Author NullableString `json:"author,omitempty"`
	AuthorId *string `json:"authorId,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	LastModificationTime *time.Time `json:"lastModificationTime,omitempty"`
	Extension NullableString `json:"extension,omitempty"`
	ContentFormating NullableString `json:"contentFormating,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Website NullableString `json:"website,omitempty"`
}

// NewNoteSpecDto instantiates a new NoteSpecDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteSpecDto() *NoteSpecDto {
	this := NoteSpecDto{}
	return &this
}

// NewNoteSpecDtoWithDefaults instantiates a new NoteSpecDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteSpecDtoWithDefaults() *NoteSpecDto {
	this := NoteSpecDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NoteSpecDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteSpecDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NoteSpecDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NoteSpecDto) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NoteSpecDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NoteSpecDto) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *NoteSpecDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *NoteSpecDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *NoteSpecDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *NoteSpecDto) UnsetTitle() {
	o.Title.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NoteSpecDto) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NoteSpecDto) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *NoteSpecDto) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *NoteSpecDto) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *NoteSpecDto) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *NoteSpecDto) UnsetContent() {
	o.Content.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NoteSpecDto) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteSpecDto) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NoteSpecDto) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *NoteSpecDto) SetVersion(v int32) {
	o.Version = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NoteSpecDto) GetAuthor() string {
	if o == nil || o.Author.Get() == nil {
		var ret string
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NoteSpecDto) GetAuthorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a field has been set.
func (o *NoteSpecDto) HasAuthor() bool {
	if o != nil && o.Author.IsSet() {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given NullableString and assigns it to the Author field.
func (o *NoteSpecDto) SetAuthor(v string) {
	o.Author.Set(&v)
}
// SetAuthorNil sets the value for Author to be an explicit nil
func (o *NoteSpecDto) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnsetAuthor ensures that no value is present for Author, not even an explicit nil
func (o *NoteSpecDto) UnsetAuthor() {
	o.Author.Unset()
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise.
func (o *NoteSpecDto) GetAuthorId() string {
	if o == nil || o.AuthorId == nil {
		var ret string
		return ret
	}
	return *o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteSpecDto) GetAuthorIdOk() (*string, bool) {
	if o == nil || o.AuthorId == nil {
		return nil, false
	}
	return o.AuthorId, true
}

// HasAuthorId returns a boolean if a field has been set.
func (o *NoteSpecDto) HasAuthorId() bool {
	if o != nil && o.AuthorId != nil {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given string and assigns it to the AuthorId field.
func (o *NoteSpecDto) SetAuthorId(v string) {
	o.AuthorId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *NoteSpecDto) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteSpecDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *NoteSpecDto) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *NoteSpecDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise.
func (o *NoteSpecDto) GetLastModificationTime() time.Time {
	if o == nil || o.LastModificationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteSpecDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModificationTime == nil {
		return nil, false
	}
	return o.LastModificationTime, true
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *NoteSpecDto) HasLastModificationTime() bool {
	if o != nil && o.LastModificationTime != nil {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given time.Time and assigns it to the LastModificationTime field.
func (o *NoteSpecDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NoteSpecDto) GetExtension() string {
	if o == nil || o.Extension.Get() == nil {
		var ret string
		return ret
	}
	return *o.Extension.Get()
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NoteSpecDto) GetExtensionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Extension.Get(), o.Extension.IsSet()
}

// HasExtension returns a boolean if a field has been set.
func (o *NoteSpecDto) HasExtension() bool {
	if o != nil && o.Extension.IsSet() {
		return true
	}

	return false
}

// SetExtension gets a reference to the given NullableString and assigns it to the Extension field.
func (o *NoteSpecDto) SetExtension(v string) {
	o.Extension.Set(&v)
}
// SetExtensionNil sets the value for Extension to be an explicit nil
func (o *NoteSpecDto) SetExtensionNil() {
	o.Extension.Set(nil)
}

// UnsetExtension ensures that no value is present for Extension, not even an explicit nil
func (o *NoteSpecDto) UnsetExtension() {
	o.Extension.Unset()
}

// GetContentFormating returns the ContentFormating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NoteSpecDto) GetContentFormating() string {
	if o == nil || o.ContentFormating.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentFormating.Get()
}

// GetContentFormatingOk returns a tuple with the ContentFormating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NoteSpecDto) GetContentFormatingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentFormating.Get(), o.ContentFormating.IsSet()
}

// HasContentFormating returns a boolean if a field has been set.
func (o *NoteSpecDto) HasContentFormating() bool {
	if o != nil && o.ContentFormating.IsSet() {
		return true
	}

	return false
}

// SetContentFormating gets a reference to the given NullableString and assigns it to the ContentFormating field.
func (o *NoteSpecDto) SetContentFormating(v string) {
	o.ContentFormating.Set(&v)
}
// SetContentFormatingNil sets the value for ContentFormating to be an explicit nil
func (o *NoteSpecDto) SetContentFormatingNil() {
	o.ContentFormating.Set(nil)
}

// UnsetContentFormating ensures that no value is present for ContentFormating, not even an explicit nil
func (o *NoteSpecDto) UnsetContentFormating() {
	o.ContentFormating.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NoteSpecDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NoteSpecDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NoteSpecDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NoteSpecDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *NoteSpecDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NoteSpecDto) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NoteSpecDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NoteSpecDto) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *NoteSpecDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *NoteSpecDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *NoteSpecDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *NoteSpecDto) UnsetDescription() {
	o.Description.Unset()
}

// GetWebsite returns the Website field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NoteSpecDto) GetWebsite() string {
	if o == nil || o.Website.Get() == nil {
		var ret string
		return ret
	}
	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NoteSpecDto) GetWebsiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// HasWebsite returns a boolean if a field has been set.
func (o *NoteSpecDto) HasWebsite() bool {
	if o != nil && o.Website.IsSet() {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given NullableString and assigns it to the Website field.
func (o *NoteSpecDto) SetWebsite(v string) {
	o.Website.Set(&v)
}
// SetWebsiteNil sets the value for Website to be an explicit nil
func (o *NoteSpecDto) SetWebsiteNil() {
	o.Website.Set(nil)
}

// UnsetWebsite ensures that no value is present for Website, not even an explicit nil
func (o *NoteSpecDto) UnsetWebsite() {
	o.Website.Unset()
}

func (o NoteSpecDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Author.IsSet() {
		toSerialize["author"] = o.Author.Get()
	}
	if o.AuthorId != nil {
		toSerialize["authorId"] = o.AuthorId
	}
	if o.CreationTime != nil {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.LastModificationTime != nil {
		toSerialize["lastModificationTime"] = o.LastModificationTime
	}
	if o.Extension.IsSet() {
		toSerialize["extension"] = o.Extension.Get()
	}
	if o.ContentFormating.IsSet() {
		toSerialize["contentFormating"] = o.ContentFormating.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Website.IsSet() {
		toSerialize["website"] = o.Website.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNoteSpecDto struct {
	value *NoteSpecDto
	isSet bool
}

func (v NullableNoteSpecDto) Get() *NoteSpecDto {
	return v.value
}

func (v *NullableNoteSpecDto) Set(val *NoteSpecDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteSpecDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteSpecDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteSpecDto(val *NoteSpecDto) *NullableNoteSpecDto {
	return &NullableNoteSpecDto{value: val, isSet: true}
}

func (v NullableNoteSpecDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteSpecDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


