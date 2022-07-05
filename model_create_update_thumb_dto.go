/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
)

// CreateUpdateThumbDto struct for CreateUpdateThumbDto
type CreateUpdateThumbDto struct {
	Key NullableString `json:"key,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Md5 NullableString `json:"md5,omitempty"`
	ContentType NullableString `json:"contentType,omitempty"`
	Extension NullableString `json:"extension,omitempty"`
	StorageClass NullableString `json:"storageClass,omitempty"`
}

// NewCreateUpdateThumbDto instantiates a new CreateUpdateThumbDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpdateThumbDto() *CreateUpdateThumbDto {
	this := CreateUpdateThumbDto{}
	return &this
}

// NewCreateUpdateThumbDtoWithDefaults instantiates a new CreateUpdateThumbDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpdateThumbDtoWithDefaults() *CreateUpdateThumbDto {
	this := CreateUpdateThumbDto{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpdateThumbDto) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpdateThumbDto) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *CreateUpdateThumbDto) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *CreateUpdateThumbDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *CreateUpdateThumbDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *CreateUpdateThumbDto) UnsetKey() {
	o.Key.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CreateUpdateThumbDto) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateThumbDto) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CreateUpdateThumbDto) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *CreateUpdateThumbDto) SetSize(v int32) {
	o.Size = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpdateThumbDto) GetMd5() string {
	if o == nil || o.Md5.Get() == nil {
		var ret string
		return ret
	}
	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpdateThumbDto) GetMd5Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// HasMd5 returns a boolean if a field has been set.
func (o *CreateUpdateThumbDto) HasMd5() bool {
	if o != nil && o.Md5.IsSet() {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given NullableString and assigns it to the Md5 field.
func (o *CreateUpdateThumbDto) SetMd5(v string) {
	o.Md5.Set(&v)
}
// SetMd5Nil sets the value for Md5 to be an explicit nil
func (o *CreateUpdateThumbDto) SetMd5Nil() {
	o.Md5.Set(nil)
}

// UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
func (o *CreateUpdateThumbDto) UnsetMd5() {
	o.Md5.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpdateThumbDto) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpdateThumbDto) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *CreateUpdateThumbDto) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *CreateUpdateThumbDto) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *CreateUpdateThumbDto) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *CreateUpdateThumbDto) UnsetContentType() {
	o.ContentType.Unset()
}

// GetExtension returns the Extension field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpdateThumbDto) GetExtension() string {
	if o == nil || o.Extension.Get() == nil {
		var ret string
		return ret
	}
	return *o.Extension.Get()
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpdateThumbDto) GetExtensionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Extension.Get(), o.Extension.IsSet()
}

// HasExtension returns a boolean if a field has been set.
func (o *CreateUpdateThumbDto) HasExtension() bool {
	if o != nil && o.Extension.IsSet() {
		return true
	}

	return false
}

// SetExtension gets a reference to the given NullableString and assigns it to the Extension field.
func (o *CreateUpdateThumbDto) SetExtension(v string) {
	o.Extension.Set(&v)
}
// SetExtensionNil sets the value for Extension to be an explicit nil
func (o *CreateUpdateThumbDto) SetExtensionNil() {
	o.Extension.Set(nil)
}

// UnsetExtension ensures that no value is present for Extension, not even an explicit nil
func (o *CreateUpdateThumbDto) UnsetExtension() {
	o.Extension.Unset()
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpdateThumbDto) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpdateThumbDto) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// HasStorageClass returns a boolean if a field has been set.
func (o *CreateUpdateThumbDto) HasStorageClass() bool {
	if o != nil && o.StorageClass.IsSet() {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given NullableString and assigns it to the StorageClass field.
func (o *CreateUpdateThumbDto) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}
// SetStorageClassNil sets the value for StorageClass to be an explicit nil
func (o *CreateUpdateThumbDto) SetStorageClassNil() {
	o.StorageClass.Set(nil)
}

// UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
func (o *CreateUpdateThumbDto) UnsetStorageClass() {
	o.StorageClass.Unset()
}

func (o CreateUpdateThumbDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Extension.IsSet() {
		toSerialize["extension"] = o.Extension.Get()
	}
	if o.StorageClass.IsSet() {
		toSerialize["storageClass"] = o.StorageClass.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateUpdateThumbDto struct {
	value *CreateUpdateThumbDto
	isSet bool
}

func (v NullableCreateUpdateThumbDto) Get() *CreateUpdateThumbDto {
	return v.value
}

func (v *NullableCreateUpdateThumbDto) Set(val *CreateUpdateThumbDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpdateThumbDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpdateThumbDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpdateThumbDto(val *CreateUpdateThumbDto) *NullableCreateUpdateThumbDto {
	return &NullableCreateUpdateThumbDto{value: val, isSet: true}
}

func (v NullableCreateUpdateThumbDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpdateThumbDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

