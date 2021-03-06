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

// ExtensionEnumDto struct for ExtensionEnumDto
type ExtensionEnumDto struct {
	Fields []ExtensionEnumFieldDto `json:"fields,omitempty"`
	LocalizationResource NullableString `json:"localizationResource,omitempty"`
}

// NewExtensionEnumDto instantiates a new ExtensionEnumDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionEnumDto() *ExtensionEnumDto {
	this := ExtensionEnumDto{}
	return &this
}

// NewExtensionEnumDtoWithDefaults instantiates a new ExtensionEnumDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionEnumDtoWithDefaults() *ExtensionEnumDto {
	this := ExtensionEnumDto{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionEnumDto) GetFields() []ExtensionEnumFieldDto {
	if o == nil  {
		var ret []ExtensionEnumFieldDto
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionEnumDto) GetFieldsOk() (*[]ExtensionEnumFieldDto, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ExtensionEnumDto) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []ExtensionEnumFieldDto and assigns it to the Fields field.
func (o *ExtensionEnumDto) SetFields(v []ExtensionEnumFieldDto) {
	o.Fields = v
}

// GetLocalizationResource returns the LocalizationResource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionEnumDto) GetLocalizationResource() string {
	if o == nil || o.LocalizationResource.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalizationResource.Get()
}

// GetLocalizationResourceOk returns a tuple with the LocalizationResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionEnumDto) GetLocalizationResourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalizationResource.Get(), o.LocalizationResource.IsSet()
}

// HasLocalizationResource returns a boolean if a field has been set.
func (o *ExtensionEnumDto) HasLocalizationResource() bool {
	if o != nil && o.LocalizationResource.IsSet() {
		return true
	}

	return false
}

// SetLocalizationResource gets a reference to the given NullableString and assigns it to the LocalizationResource field.
func (o *ExtensionEnumDto) SetLocalizationResource(v string) {
	o.LocalizationResource.Set(&v)
}
// SetLocalizationResourceNil sets the value for LocalizationResource to be an explicit nil
func (o *ExtensionEnumDto) SetLocalizationResourceNil() {
	o.LocalizationResource.Set(nil)
}

// UnsetLocalizationResource ensures that no value is present for LocalizationResource, not even an explicit nil
func (o *ExtensionEnumDto) UnsetLocalizationResource() {
	o.LocalizationResource.Unset()
}

func (o ExtensionEnumDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.LocalizationResource.IsSet() {
		toSerialize["localizationResource"] = o.LocalizationResource.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableExtensionEnumDto struct {
	value *ExtensionEnumDto
	isSet bool
}

func (v NullableExtensionEnumDto) Get() *ExtensionEnumDto {
	return v.value
}

func (v *NullableExtensionEnumDto) Set(val *ExtensionEnumDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionEnumDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionEnumDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionEnumDto(val *ExtensionEnumDto) *NullableExtensionEnumDto {
	return &NullableExtensionEnumDto{value: val, isSet: true}
}

func (v NullableExtensionEnumDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionEnumDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


