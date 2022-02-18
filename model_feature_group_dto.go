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

// FeatureGroupDto struct for FeatureGroupDto
type FeatureGroupDto struct {
	Name NullableString `json:"name,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	Features []FeatureDto `json:"features,omitempty"`
}

// NewFeatureGroupDto instantiates a new FeatureGroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureGroupDto() *FeatureGroupDto {
	this := FeatureGroupDto{}
	return &this
}

// NewFeatureGroupDtoWithDefaults instantiates a new FeatureGroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureGroupDtoWithDefaults() *FeatureGroupDto {
	this := FeatureGroupDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureGroupDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureGroupDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FeatureGroupDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FeatureGroupDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *FeatureGroupDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FeatureGroupDto) UnsetName() {
	o.Name.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureGroupDto) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureGroupDto) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FeatureGroupDto) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *FeatureGroupDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *FeatureGroupDto) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *FeatureGroupDto) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetFeatures returns the Features field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureGroupDto) GetFeatures() []FeatureDto {
	if o == nil  {
		var ret []FeatureDto
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureGroupDto) GetFeaturesOk() (*[]FeatureDto, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return &o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FeatureGroupDto) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []FeatureDto and assigns it to the Features field.
func (o *FeatureGroupDto) SetFeatures(v []FeatureDto) {
	o.Features = v
}

func (o FeatureGroupDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableFeatureGroupDto struct {
	value *FeatureGroupDto
	isSet bool
}

func (v NullableFeatureGroupDto) Get() *FeatureGroupDto {
	return v.value
}

func (v *NullableFeatureGroupDto) Set(val *FeatureGroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureGroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureGroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureGroupDto(val *FeatureGroupDto) *NullableFeatureGroupDto {
	return &NullableFeatureGroupDto{value: val, isSet: true}
}

func (v NullableFeatureGroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureGroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

