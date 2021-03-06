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

// IdentityRoleDto struct for IdentityRoleDto
type IdentityRoleDto struct {
	ExtraProperties map[string]interface{} `json:"extraProperties,omitempty"`
	Id *string `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsStatic *bool `json:"isStatic,omitempty"`
	IsPublic *bool `json:"isPublic,omitempty"`
	ConcurrencyStamp NullableString `json:"concurrencyStamp,omitempty"`
}

// NewIdentityRoleDto instantiates a new IdentityRoleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityRoleDto() *IdentityRoleDto {
	this := IdentityRoleDto{}
	return &this
}

// NewIdentityRoleDtoWithDefaults instantiates a new IdentityRoleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityRoleDtoWithDefaults() *IdentityRoleDto {
	this := IdentityRoleDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityRoleDto) GetExtraProperties() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityRoleDto) GetExtraPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExtraProperties == nil {
		return nil, false
	}
	return &o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasExtraProperties() bool {
	if o != nil && o.ExtraProperties != nil {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]interface{} and assigns it to the ExtraProperties field.
func (o *IdentityRoleDto) SetExtraProperties(v map[string]interface{}) {
	o.ExtraProperties = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityRoleDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityRoleDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityRoleDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IdentityRoleDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IdentityRoleDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IdentityRoleDto) UnsetName() {
	o.Name.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *IdentityRoleDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsStatic returns the IsStatic field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetIsStatic() bool {
	if o == nil || o.IsStatic == nil {
		var ret bool
		return ret
	}
	return *o.IsStatic
}

// GetIsStaticOk returns a tuple with the IsStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetIsStaticOk() (*bool, bool) {
	if o == nil || o.IsStatic == nil {
		return nil, false
	}
	return o.IsStatic, true
}

// HasIsStatic returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasIsStatic() bool {
	if o != nil && o.IsStatic != nil {
		return true
	}

	return false
}

// SetIsStatic gets a reference to the given bool and assigns it to the IsStatic field.
func (o *IdentityRoleDto) SetIsStatic(v bool) {
	o.IsStatic = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *IdentityRoleDto) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityRoleDto) GetConcurrencyStamp() string {
	if o == nil || o.ConcurrencyStamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp.Get()
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityRoleDto) GetConcurrencyStampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConcurrencyStamp.Get(), o.ConcurrencyStamp.IsSet()
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasConcurrencyStamp() bool {
	if o != nil && o.ConcurrencyStamp.IsSet() {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given NullableString and assigns it to the ConcurrencyStamp field.
func (o *IdentityRoleDto) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp.Set(&v)
}
// SetConcurrencyStampNil sets the value for ConcurrencyStamp to be an explicit nil
func (o *IdentityRoleDto) SetConcurrencyStampNil() {
	o.ConcurrencyStamp.Set(nil)
}

// UnsetConcurrencyStamp ensures that no value is present for ConcurrencyStamp, not even an explicit nil
func (o *IdentityRoleDto) UnsetConcurrencyStamp() {
	o.ConcurrencyStamp.Unset()
}

func (o IdentityRoleDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtraProperties != nil {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.IsStatic != nil {
		toSerialize["isStatic"] = o.IsStatic
	}
	if o.IsPublic != nil {
		toSerialize["isPublic"] = o.IsPublic
	}
	if o.ConcurrencyStamp.IsSet() {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityRoleDto struct {
	value *IdentityRoleDto
	isSet bool
}

func (v NullableIdentityRoleDto) Get() *IdentityRoleDto {
	return v.value
}

func (v *NullableIdentityRoleDto) Set(val *IdentityRoleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRoleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRoleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRoleDto(val *IdentityRoleDto) *NullableIdentityRoleDto {
	return &NullableIdentityRoleDto{value: val, isSet: true}
}

func (v NullableIdentityRoleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRoleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


