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

// FileDto struct for FileDto
type FileDto struct {
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId NullableString `json:"creatorId,omitempty"`
	LastModificationTime NullableTime `json:"lastModificationTime,omitempty"`
	LastModifierId NullableString `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId NullableString `json:"deleterId,omitempty"`
	DeletionTime NullableTime `json:"deletionTime,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Key NullableString `json:"key,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Md5 NullableString `json:"md5,omitempty"`
	SliceMd5 NullableString `json:"sliceMd5,omitempty"`
	RapidCode NullableString `json:"rapidCode,omitempty"`
	ContentType NullableString `json:"contentType,omitempty"`
	Extension NullableString `json:"extension,omitempty"`
	StorageClass NullableString `json:"storageClass,omitempty"`
	FileCreatedAt NullableTime `json:"fileCreatedAt,omitempty"`
	FileUpdatedAt NullableTime `json:"fileUpdatedAt,omitempty"`
	SyncVersion *int64 `json:"syncVersion,omitempty"`
}

// NewFileDto instantiates a new FileDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDto() *FileDto {
	this := FileDto{}
	return &this
}

// NewFileDtoWithDefaults instantiates a new FileDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDtoWithDefaults() *FileDto {
	this := FileDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FileDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *FileDto) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *FileDto) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *FileDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetCreatorId() string {
	if o == nil || o.CreatorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetCreatorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *FileDto) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given NullableString and assigns it to the CreatorId field.
func (o *FileDto) SetCreatorId(v string) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *FileDto) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *FileDto) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetLastModificationTime() time.Time {
	if o == nil || o.LastModificationTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime.Get()
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModificationTime.Get(), o.LastModificationTime.IsSet()
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *FileDto) HasLastModificationTime() bool {
	if o != nil && o.LastModificationTime.IsSet() {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given NullableTime and assigns it to the LastModificationTime field.
func (o *FileDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime.Set(&v)
}
// SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil
func (o *FileDto) SetLastModificationTimeNil() {
	o.LastModificationTime.Set(nil)
}

// UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
func (o *FileDto) UnsetLastModificationTime() {
	o.LastModificationTime.Unset()
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetLastModifierId() string {
	if o == nil || o.LastModifierId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastModifierId.Get()
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetLastModifierIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifierId.Get(), o.LastModifierId.IsSet()
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *FileDto) HasLastModifierId() bool {
	if o != nil && o.LastModifierId.IsSet() {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given NullableString and assigns it to the LastModifierId field.
func (o *FileDto) SetLastModifierId(v string) {
	o.LastModifierId.Set(&v)
}
// SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil
func (o *FileDto) SetLastModifierIdNil() {
	o.LastModifierId.Set(nil)
}

// UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
func (o *FileDto) UnsetLastModifierId() {
	o.LastModifierId.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *FileDto) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *FileDto) HasIsDeleted() bool {
	if o != nil && o.IsDeleted != nil {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *FileDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetDeleterId() string {
	if o == nil || o.DeleterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeleterId.Get()
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetDeleterIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeleterId.Get(), o.DeleterId.IsSet()
}

// HasDeleterId returns a boolean if a field has been set.
func (o *FileDto) HasDeleterId() bool {
	if o != nil && o.DeleterId.IsSet() {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given NullableString and assigns it to the DeleterId field.
func (o *FileDto) SetDeleterId(v string) {
	o.DeleterId.Set(&v)
}
// SetDeleterIdNil sets the value for DeleterId to be an explicit nil
func (o *FileDto) SetDeleterIdNil() {
	o.DeleterId.Set(nil)
}

// UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
func (o *FileDto) UnsetDeleterId() {
	o.DeleterId.Unset()
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetDeletionTime() time.Time {
	if o == nil || o.DeletionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime.Get()
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletionTime.Get(), o.DeletionTime.IsSet()
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *FileDto) HasDeletionTime() bool {
	if o != nil && o.DeletionTime.IsSet() {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given NullableTime and assigns it to the DeletionTime field.
func (o *FileDto) SetDeletionTime(v time.Time) {
	o.DeletionTime.Set(&v)
}
// SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil
func (o *FileDto) SetDeletionTimeNil() {
	o.DeletionTime.Set(nil)
}

// UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
func (o *FileDto) UnsetDeletionTime() {
	o.DeletionTime.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FileDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FileDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *FileDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FileDto) UnsetName() {
	o.Name.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *FileDto) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *FileDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *FileDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *FileDto) UnsetKey() {
	o.Key.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *FileDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *FileDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *FileDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *FileDto) UnsetUrl() {
	o.Url.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FileDto) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDto) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FileDto) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *FileDto) SetSize(v int32) {
	o.Size = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetMd5() string {
	if o == nil || o.Md5.Get() == nil {
		var ret string
		return ret
	}
	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetMd5Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// HasMd5 returns a boolean if a field has been set.
func (o *FileDto) HasMd5() bool {
	if o != nil && o.Md5.IsSet() {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given NullableString and assigns it to the Md5 field.
func (o *FileDto) SetMd5(v string) {
	o.Md5.Set(&v)
}
// SetMd5Nil sets the value for Md5 to be an explicit nil
func (o *FileDto) SetMd5Nil() {
	o.Md5.Set(nil)
}

// UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
func (o *FileDto) UnsetMd5() {
	o.Md5.Unset()
}

// GetSliceMd5 returns the SliceMd5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetSliceMd5() string {
	if o == nil || o.SliceMd5.Get() == nil {
		var ret string
		return ret
	}
	return *o.SliceMd5.Get()
}

// GetSliceMd5Ok returns a tuple with the SliceMd5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetSliceMd5Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SliceMd5.Get(), o.SliceMd5.IsSet()
}

// HasSliceMd5 returns a boolean if a field has been set.
func (o *FileDto) HasSliceMd5() bool {
	if o != nil && o.SliceMd5.IsSet() {
		return true
	}

	return false
}

// SetSliceMd5 gets a reference to the given NullableString and assigns it to the SliceMd5 field.
func (o *FileDto) SetSliceMd5(v string) {
	o.SliceMd5.Set(&v)
}
// SetSliceMd5Nil sets the value for SliceMd5 to be an explicit nil
func (o *FileDto) SetSliceMd5Nil() {
	o.SliceMd5.Set(nil)
}

// UnsetSliceMd5 ensures that no value is present for SliceMd5, not even an explicit nil
func (o *FileDto) UnsetSliceMd5() {
	o.SliceMd5.Unset()
}

// GetRapidCode returns the RapidCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetRapidCode() string {
	if o == nil || o.RapidCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.RapidCode.Get()
}

// GetRapidCodeOk returns a tuple with the RapidCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetRapidCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RapidCode.Get(), o.RapidCode.IsSet()
}

// HasRapidCode returns a boolean if a field has been set.
func (o *FileDto) HasRapidCode() bool {
	if o != nil && o.RapidCode.IsSet() {
		return true
	}

	return false
}

// SetRapidCode gets a reference to the given NullableString and assigns it to the RapidCode field.
func (o *FileDto) SetRapidCode(v string) {
	o.RapidCode.Set(&v)
}
// SetRapidCodeNil sets the value for RapidCode to be an explicit nil
func (o *FileDto) SetRapidCodeNil() {
	o.RapidCode.Set(nil)
}

// UnsetRapidCode ensures that no value is present for RapidCode, not even an explicit nil
func (o *FileDto) UnsetRapidCode() {
	o.RapidCode.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *FileDto) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *FileDto) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *FileDto) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *FileDto) UnsetContentType() {
	o.ContentType.Unset()
}

// GetExtension returns the Extension field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetExtension() string {
	if o == nil || o.Extension.Get() == nil {
		var ret string
		return ret
	}
	return *o.Extension.Get()
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetExtensionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Extension.Get(), o.Extension.IsSet()
}

// HasExtension returns a boolean if a field has been set.
func (o *FileDto) HasExtension() bool {
	if o != nil && o.Extension.IsSet() {
		return true
	}

	return false
}

// SetExtension gets a reference to the given NullableString and assigns it to the Extension field.
func (o *FileDto) SetExtension(v string) {
	o.Extension.Set(&v)
}
// SetExtensionNil sets the value for Extension to be an explicit nil
func (o *FileDto) SetExtensionNil() {
	o.Extension.Set(nil)
}

// UnsetExtension ensures that no value is present for Extension, not even an explicit nil
func (o *FileDto) UnsetExtension() {
	o.Extension.Unset()
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// HasStorageClass returns a boolean if a field has been set.
func (o *FileDto) HasStorageClass() bool {
	if o != nil && o.StorageClass.IsSet() {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given NullableString and assigns it to the StorageClass field.
func (o *FileDto) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}
// SetStorageClassNil sets the value for StorageClass to be an explicit nil
func (o *FileDto) SetStorageClassNil() {
	o.StorageClass.Set(nil)
}

// UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
func (o *FileDto) UnsetStorageClass() {
	o.StorageClass.Unset()
}

// GetFileCreatedAt returns the FileCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetFileCreatedAt() time.Time {
	if o == nil || o.FileCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FileCreatedAt.Get()
}

// GetFileCreatedAtOk returns a tuple with the FileCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetFileCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileCreatedAt.Get(), o.FileCreatedAt.IsSet()
}

// HasFileCreatedAt returns a boolean if a field has been set.
func (o *FileDto) HasFileCreatedAt() bool {
	if o != nil && o.FileCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetFileCreatedAt gets a reference to the given NullableTime and assigns it to the FileCreatedAt field.
func (o *FileDto) SetFileCreatedAt(v time.Time) {
	o.FileCreatedAt.Set(&v)
}
// SetFileCreatedAtNil sets the value for FileCreatedAt to be an explicit nil
func (o *FileDto) SetFileCreatedAtNil() {
	o.FileCreatedAt.Set(nil)
}

// UnsetFileCreatedAt ensures that no value is present for FileCreatedAt, not even an explicit nil
func (o *FileDto) UnsetFileCreatedAt() {
	o.FileCreatedAt.Unset()
}

// GetFileUpdatedAt returns the FileUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDto) GetFileUpdatedAt() time.Time {
	if o == nil || o.FileUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FileUpdatedAt.Get()
}

// GetFileUpdatedAtOk returns a tuple with the FileUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDto) GetFileUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileUpdatedAt.Get(), o.FileUpdatedAt.IsSet()
}

// HasFileUpdatedAt returns a boolean if a field has been set.
func (o *FileDto) HasFileUpdatedAt() bool {
	if o != nil && o.FileUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetFileUpdatedAt gets a reference to the given NullableTime and assigns it to the FileUpdatedAt field.
func (o *FileDto) SetFileUpdatedAt(v time.Time) {
	o.FileUpdatedAt.Set(&v)
}
// SetFileUpdatedAtNil sets the value for FileUpdatedAt to be an explicit nil
func (o *FileDto) SetFileUpdatedAtNil() {
	o.FileUpdatedAt.Set(nil)
}

// UnsetFileUpdatedAt ensures that no value is present for FileUpdatedAt, not even an explicit nil
func (o *FileDto) UnsetFileUpdatedAt() {
	o.FileUpdatedAt.Unset()
}

// GetSyncVersion returns the SyncVersion field value if set, zero value otherwise.
func (o *FileDto) GetSyncVersion() int64 {
	if o == nil || o.SyncVersion == nil {
		var ret int64
		return ret
	}
	return *o.SyncVersion
}

// GetSyncVersionOk returns a tuple with the SyncVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDto) GetSyncVersionOk() (*int64, bool) {
	if o == nil || o.SyncVersion == nil {
		return nil, false
	}
	return o.SyncVersion, true
}

// HasSyncVersion returns a boolean if a field has been set.
func (o *FileDto) HasSyncVersion() bool {
	if o != nil && o.SyncVersion != nil {
		return true
	}

	return false
}

// SetSyncVersion gets a reference to the given int64 and assigns it to the SyncVersion field.
func (o *FileDto) SetSyncVersion(v int64) {
	o.SyncVersion = &v
}

func (o FileDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Md5.IsSet() {
		toSerialize["md5"] = o.Md5.Get()
	}
	if o.SliceMd5.IsSet() {
		toSerialize["sliceMd5"] = o.SliceMd5.Get()
	}
	if o.RapidCode.IsSet() {
		toSerialize["rapidCode"] = o.RapidCode.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["contentType"] = o.ContentType.Get()
	}
	if o.Extension.IsSet() {
		toSerialize["extension"] = o.Extension.Get()
	}
	if o.StorageClass.IsSet() {
		toSerialize["storageClass"] = o.StorageClass.Get()
	}
	if o.FileCreatedAt.IsSet() {
		toSerialize["fileCreatedAt"] = o.FileCreatedAt.Get()
	}
	if o.FileUpdatedAt.IsSet() {
		toSerialize["fileUpdatedAt"] = o.FileUpdatedAt.Get()
	}
	if o.SyncVersion != nil {
		toSerialize["syncVersion"] = o.SyncVersion
	}
	return json.Marshal(toSerialize)
}

type NullableFileDto struct {
	value *FileDto
	isSet bool
}

func (v NullableFileDto) Get() *FileDto {
	return v.value
}

func (v *NullableFileDto) Set(val *FileDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDto(val *FileDto) *NullableFileDto {
	return &NullableFileDto{value: val, isSet: true}
}

func (v NullableFileDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


