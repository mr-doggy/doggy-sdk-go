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

// File struct for File
type File struct {
	Id *string `json:"id,omitempty"`
	ExtraProperties map[string]interface{} `json:"extraProperties,omitempty"`
	ConcurrencyStamp NullableString `json:"concurrencyStamp,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId NullableString `json:"creatorId,omitempty"`
	LastModificationTime NullableTime `json:"lastModificationTime,omitempty"`
	LastModifierId NullableString `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId NullableString `json:"deleterId,omitempty"`
	DeletionTime NullableTime `json:"deletionTime,omitempty"`
	Key NullableString `json:"key,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Md5 NullableString `json:"md5,omitempty"`
	ContentType NullableString `json:"contentType,omitempty"`
	Type *ItemType `json:"type,omitempty"`
	Extension NullableString `json:"extension,omitempty"`
	StorageClass NullableString `json:"storageClass,omitempty"`
	Items []Item `json:"items,omitempty"`
}

// NewFile instantiates a new File object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFile() *File {
	this := File{}
	return &this
}

// NewFileWithDefaults instantiates a new File object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileWithDefaults() *File {
	this := File{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *File) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *File) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *File) SetId(v string) {
	o.Id = &v
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetExtraProperties() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetExtraPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExtraProperties == nil {
		return nil, false
	}
	return &o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *File) HasExtraProperties() bool {
	if o != nil && o.ExtraProperties != nil {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]interface{} and assigns it to the ExtraProperties field.
func (o *File) SetExtraProperties(v map[string]interface{}) {
	o.ExtraProperties = v
}

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetConcurrencyStamp() string {
	if o == nil || o.ConcurrencyStamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp.Get()
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetConcurrencyStampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConcurrencyStamp.Get(), o.ConcurrencyStamp.IsSet()
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *File) HasConcurrencyStamp() bool {
	if o != nil && o.ConcurrencyStamp.IsSet() {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given NullableString and assigns it to the ConcurrencyStamp field.
func (o *File) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp.Set(&v)
}
// SetConcurrencyStampNil sets the value for ConcurrencyStamp to be an explicit nil
func (o *File) SetConcurrencyStampNil() {
	o.ConcurrencyStamp.Set(nil)
}

// UnsetConcurrencyStamp ensures that no value is present for ConcurrencyStamp, not even an explicit nil
func (o *File) UnsetConcurrencyStamp() {
	o.ConcurrencyStamp.Unset()
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *File) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *File) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *File) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetCreatorId() string {
	if o == nil || o.CreatorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetCreatorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *File) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given NullableString and assigns it to the CreatorId field.
func (o *File) SetCreatorId(v string) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *File) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *File) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetLastModificationTime() time.Time {
	if o == nil || o.LastModificationTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime.Get()
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModificationTime.Get(), o.LastModificationTime.IsSet()
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *File) HasLastModificationTime() bool {
	if o != nil && o.LastModificationTime.IsSet() {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given NullableTime and assigns it to the LastModificationTime field.
func (o *File) SetLastModificationTime(v time.Time) {
	o.LastModificationTime.Set(&v)
}
// SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil
func (o *File) SetLastModificationTimeNil() {
	o.LastModificationTime.Set(nil)
}

// UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
func (o *File) UnsetLastModificationTime() {
	o.LastModificationTime.Unset()
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetLastModifierId() string {
	if o == nil || o.LastModifierId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastModifierId.Get()
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetLastModifierIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifierId.Get(), o.LastModifierId.IsSet()
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *File) HasLastModifierId() bool {
	if o != nil && o.LastModifierId.IsSet() {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given NullableString and assigns it to the LastModifierId field.
func (o *File) SetLastModifierId(v string) {
	o.LastModifierId.Set(&v)
}
// SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil
func (o *File) SetLastModifierIdNil() {
	o.LastModifierId.Set(nil)
}

// UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
func (o *File) UnsetLastModifierId() {
	o.LastModifierId.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *File) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *File) HasIsDeleted() bool {
	if o != nil && o.IsDeleted != nil {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *File) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetDeleterId() string {
	if o == nil || o.DeleterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeleterId.Get()
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetDeleterIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeleterId.Get(), o.DeleterId.IsSet()
}

// HasDeleterId returns a boolean if a field has been set.
func (o *File) HasDeleterId() bool {
	if o != nil && o.DeleterId.IsSet() {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given NullableString and assigns it to the DeleterId field.
func (o *File) SetDeleterId(v string) {
	o.DeleterId.Set(&v)
}
// SetDeleterIdNil sets the value for DeleterId to be an explicit nil
func (o *File) SetDeleterIdNil() {
	o.DeleterId.Set(nil)
}

// UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
func (o *File) UnsetDeleterId() {
	o.DeleterId.Unset()
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetDeletionTime() time.Time {
	if o == nil || o.DeletionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime.Get()
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletionTime.Get(), o.DeletionTime.IsSet()
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *File) HasDeletionTime() bool {
	if o != nil && o.DeletionTime.IsSet() {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given NullableTime and assigns it to the DeletionTime field.
func (o *File) SetDeletionTime(v time.Time) {
	o.DeletionTime.Set(&v)
}
// SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil
func (o *File) SetDeletionTimeNil() {
	o.DeletionTime.Set(nil)
}

// UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
func (o *File) UnsetDeletionTime() {
	o.DeletionTime.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *File) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *File) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *File) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *File) UnsetKey() {
	o.Key.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *File) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *File) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *File) SetSize(v int64) {
	o.Size = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetMd5() string {
	if o == nil || o.Md5.Get() == nil {
		var ret string
		return ret
	}
	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetMd5Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// HasMd5 returns a boolean if a field has been set.
func (o *File) HasMd5() bool {
	if o != nil && o.Md5.IsSet() {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given NullableString and assigns it to the Md5 field.
func (o *File) SetMd5(v string) {
	o.Md5.Set(&v)
}
// SetMd5Nil sets the value for Md5 to be an explicit nil
func (o *File) SetMd5Nil() {
	o.Md5.Set(nil)
}

// UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
func (o *File) UnsetMd5() {
	o.Md5.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *File) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *File) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *File) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *File) UnsetContentType() {
	o.ContentType.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *File) GetType() ItemType {
	if o == nil || o.Type == nil {
		var ret ItemType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetTypeOk() (*ItemType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *File) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ItemType and assigns it to the Type field.
func (o *File) SetType(v ItemType) {
	o.Type = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetExtension() string {
	if o == nil || o.Extension.Get() == nil {
		var ret string
		return ret
	}
	return *o.Extension.Get()
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetExtensionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Extension.Get(), o.Extension.IsSet()
}

// HasExtension returns a boolean if a field has been set.
func (o *File) HasExtension() bool {
	if o != nil && o.Extension.IsSet() {
		return true
	}

	return false
}

// SetExtension gets a reference to the given NullableString and assigns it to the Extension field.
func (o *File) SetExtension(v string) {
	o.Extension.Set(&v)
}
// SetExtensionNil sets the value for Extension to be an explicit nil
func (o *File) SetExtensionNil() {
	o.Extension.Set(nil)
}

// UnsetExtension ensures that no value is present for Extension, not even an explicit nil
func (o *File) UnsetExtension() {
	o.Extension.Unset()
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// HasStorageClass returns a boolean if a field has been set.
func (o *File) HasStorageClass() bool {
	if o != nil && o.StorageClass.IsSet() {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given NullableString and assigns it to the StorageClass field.
func (o *File) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}
// SetStorageClassNil sets the value for StorageClass to be an explicit nil
func (o *File) SetStorageClassNil() {
	o.StorageClass.Set(nil)
}

// UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
func (o *File) UnsetStorageClass() {
	o.StorageClass.Unset()
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *File) GetItems() []Item {
	if o == nil  {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *File) GetItemsOk() (*[]Item, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *File) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *File) SetItems(v []Item) {
	o.Items = v
}

func (o File) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ExtraProperties != nil {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	if o.ConcurrencyStamp.IsSet() {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp.Get()
	}
	if o.CreationTime != nil {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.CreatorId.IsSet() {
		toSerialize["creatorId"] = o.CreatorId.Get()
	}
	if o.LastModificationTime.IsSet() {
		toSerialize["lastModificationTime"] = o.LastModificationTime.Get()
	}
	if o.LastModifierId.IsSet() {
		toSerialize["lastModifierId"] = o.LastModifierId.Get()
	}
	if o.IsDeleted != nil {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if o.DeleterId.IsSet() {
		toSerialize["deleterId"] = o.DeleterId.Get()
	}
	if o.DeletionTime.IsSet() {
		toSerialize["deletionTime"] = o.DeletionTime.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Md5.IsSet() {
		toSerialize["md5"] = o.Md5.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["contentType"] = o.ContentType.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Extension.IsSet() {
		toSerialize["extension"] = o.Extension.Get()
	}
	if o.StorageClass.IsSet() {
		toSerialize["storageClass"] = o.StorageClass.Get()
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableFile struct {
	value *File
	isSet bool
}

func (v NullableFile) Get() *File {
	return v.value
}

func (v *NullableFile) Set(val *File) {
	v.value = val
	v.isSet = true
}

func (v NullableFile) IsSet() bool {
	return v.isSet
}

func (v *NullableFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFile(val *File) *NullableFile {
	return &NullableFile{value: val, isSet: true}
}

func (v NullableFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

