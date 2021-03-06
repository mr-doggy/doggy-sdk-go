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

// ExtensionEnumFieldDto struct for ExtensionEnumFieldDto
type ExtensionEnumFieldDto struct {
	Name NullableString `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// NewExtensionEnumFieldDto instantiates a new ExtensionEnumFieldDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionEnumFieldDto() *ExtensionEnumFieldDto {
	this := ExtensionEnumFieldDto{}
	return &this
}

// NewExtensionEnumFieldDtoWithDefaults instantiates a new ExtensionEnumFieldDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionEnumFieldDtoWithDefaults() *ExtensionEnumFieldDto {
	this := ExtensionEnumFieldDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionEnumFieldDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionEnumFieldDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ExtensionEnumFieldDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ExtensionEnumFieldDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ExtensionEnumFieldDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ExtensionEnumFieldDto) UnsetName() {
	o.Name.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionEnumFieldDto) GetValue() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionEnumFieldDto) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ExtensionEnumFieldDto) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *ExtensionEnumFieldDto) SetValue(v interface{}) {
	o.Value = v
}

func (o ExtensionEnumFieldDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableExtensionEnumFieldDto struct {
	value *ExtensionEnumFieldDto
	isSet bool
}

func (v NullableExtensionEnumFieldDto) Get() *ExtensionEnumFieldDto {
	return v.value
}

func (v *NullableExtensionEnumFieldDto) Set(val *ExtensionEnumFieldDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionEnumFieldDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionEnumFieldDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionEnumFieldDto(val *ExtensionEnumFieldDto) *NullableExtensionEnumFieldDto {
	return &NullableExtensionEnumFieldDto{value: val, isSet: true}
}

func (v NullableExtensionEnumFieldDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionEnumFieldDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


