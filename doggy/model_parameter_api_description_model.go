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

// ParameterApiDescriptionModel struct for ParameterApiDescriptionModel
type ParameterApiDescriptionModel struct {
	NameOnMethod NullableString `json:"nameOnMethod,omitempty"`
	Name NullableString `json:"name,omitempty"`
	JsonName NullableString `json:"jsonName,omitempty"`
	Type NullableString `json:"type,omitempty"`
	TypeSimple NullableString `json:"typeSimple,omitempty"`
	IsOptional *bool `json:"isOptional,omitempty"`
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	ConstraintTypes []string `json:"constraintTypes,omitempty"`
	BindingSourceId NullableString `json:"bindingSourceId,omitempty"`
	DescriptorName NullableString `json:"descriptorName,omitempty"`
}

// NewParameterApiDescriptionModel instantiates a new ParameterApiDescriptionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterApiDescriptionModel() *ParameterApiDescriptionModel {
	this := ParameterApiDescriptionModel{}
	return &this
}

// NewParameterApiDescriptionModelWithDefaults instantiates a new ParameterApiDescriptionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterApiDescriptionModelWithDefaults() *ParameterApiDescriptionModel {
	this := ParameterApiDescriptionModel{}
	return &this
}

// GetNameOnMethod returns the NameOnMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterApiDescriptionModel) GetNameOnMethod() string {
	if o == nil || o.NameOnMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.NameOnMethod.Get()
}

// GetNameOnMethodOk returns a tuple with the NameOnMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterApiDescriptionModel) GetNameOnMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NameOnMethod.Get(), o.NameOnMethod.IsSet()
}

// HasNameOnMethod returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasNameOnMethod() bool {
	if o != nil && o.NameOnMethod.IsSet() {
		return true
	}

	return false
}

// SetNameOnMethod gets a reference to the given NullableString and assigns it to the NameOnMethod field.
func (o *ParameterApiDescriptionModel) SetNameOnMethod(v string) {
	o.NameOnMethod.Set(&v)
}
// SetNameOnMethodNil sets the value for NameOnMethod to be an explicit nil
func (o *ParameterApiDescriptionModel) SetNameOnMethodNil() {
	o.NameOnMethod.Set(nil)
}

// UnsetNameOnMethod ensures that no value is present for NameOnMethod, not even an explicit nil
func (o *ParameterApiDescriptionModel) UnsetNameOnMethod() {
	o.NameOnMethod.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterApiDescriptionModel) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterApiDescriptionModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ParameterApiDescriptionModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ParameterApiDescriptionModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ParameterApiDescriptionModel) UnsetName() {
	o.Name.Unset()
}

// GetJsonName returns the JsonName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterApiDescriptionModel) GetJsonName() string {
	if o == nil || o.JsonName.Get() == nil {
		var ret string
		return ret
	}
	return *o.JsonName.Get()
}

// GetJsonNameOk returns a tuple with the JsonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterApiDescriptionModel) GetJsonNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JsonName.Get(), o.JsonName.IsSet()
}

// HasJsonName returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasJsonName() bool {
	if o != nil && o.JsonName.IsSet() {
		return true
	}

	return false
}

// SetJsonName gets a reference to the given NullableString and assigns it to the JsonName field.
func (o *ParameterApiDescriptionModel) SetJsonName(v string) {
	o.JsonName.Set(&v)
}
// SetJsonNameNil sets the value for JsonName to be an explicit nil
func (o *ParameterApiDescriptionModel) SetJsonNameNil() {
	o.JsonName.Set(nil)
}

// UnsetJsonName ensures that no value is present for JsonName, not even an explicit nil
func (o *ParameterApiDescriptionModel) UnsetJsonName() {
	o.JsonName.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterApiDescriptionModel) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterApiDescriptionModel) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ParameterApiDescriptionModel) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ParameterApiDescriptionModel) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ParameterApiDescriptionModel) UnsetType() {
	o.Type.Unset()
}

// GetTypeSimple returns the TypeSimple field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterApiDescriptionModel) GetTypeSimple() string {
	if o == nil || o.TypeSimple.Get() == nil {
		var ret string
		return ret
	}
	return *o.TypeSimple.Get()
}

// GetTypeSimpleOk returns a tuple with the TypeSimple field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterApiDescriptionModel) GetTypeSimpleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TypeSimple.Get(), o.TypeSimple.IsSet()
}

// HasTypeSimple returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasTypeSimple() bool {
	if o != nil && o.TypeSimple.IsSet() {
		return true
	}

	return false
}

// SetTypeSimple gets a reference to the given NullableString and assigns it to the TypeSimple field.
func (o *ParameterApiDescriptionModel) SetTypeSimple(v string) {
	o.TypeSimple.Set(&v)
}
// SetTypeSimpleNil sets the value for TypeSimple to be an explicit nil
func (o *ParameterApiDescriptionModel) SetTypeSimpleNil() {
	o.TypeSimple.Set(nil)
}

// UnsetTypeSimple ensures that no value is present for TypeSimple, not even an explicit nil
func (o *ParameterApiDescriptionModel) UnsetTypeSimple() {
	o.TypeSimple.Unset()
}

// GetIsOptional returns the IsOptional field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetIsOptional() bool {
	if o == nil || o.IsOptional == nil {
		var ret bool
		return ret
	}
	return *o.IsOptional
}

// GetIsOptionalOk returns a tuple with the IsOptional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetIsOptionalOk() (*bool, bool) {
	if o == nil || o.IsOptional == nil {
		return nil, false
	}
	return o.IsOptional, true
}

// HasIsOptional returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasIsOptional() bool {
	if o != nil && o.IsOptional != nil {
		return true
	}

	return false
}

// SetIsOptional gets a reference to the given bool and assigns it to the IsOptional field.
func (o *ParameterApiDescriptionModel) SetIsOptional(v bool) {
	o.IsOptional = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterApiDescriptionModel) GetDefaultValue() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterApiDescriptionModel) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given interface{} and assigns it to the DefaultValue field.
func (o *ParameterApiDescriptionModel) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetConstraintTypes returns the ConstraintTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterApiDescriptionModel) GetConstraintTypes() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.ConstraintTypes
}

// GetConstraintTypesOk returns a tuple with the ConstraintTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterApiDescriptionModel) GetConstraintTypesOk() (*[]string, bool) {
	if o == nil || o.ConstraintTypes == nil {
		return nil, false
	}
	return &o.ConstraintTypes, true
}

// HasConstraintTypes returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasConstraintTypes() bool {
	if o != nil && o.ConstraintTypes != nil {
		return true
	}

	return false
}

// SetConstraintTypes gets a reference to the given []string and assigns it to the ConstraintTypes field.
func (o *ParameterApiDescriptionModel) SetConstraintTypes(v []string) {
	o.ConstraintTypes = v
}

// GetBindingSourceId returns the BindingSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterApiDescriptionModel) GetBindingSourceId() string {
	if o == nil || o.BindingSourceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.BindingSourceId.Get()
}

// GetBindingSourceIdOk returns a tuple with the BindingSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterApiDescriptionModel) GetBindingSourceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BindingSourceId.Get(), o.BindingSourceId.IsSet()
}

// HasBindingSourceId returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasBindingSourceId() bool {
	if o != nil && o.BindingSourceId.IsSet() {
		return true
	}

	return false
}

// SetBindingSourceId gets a reference to the given NullableString and assigns it to the BindingSourceId field.
func (o *ParameterApiDescriptionModel) SetBindingSourceId(v string) {
	o.BindingSourceId.Set(&v)
}
// SetBindingSourceIdNil sets the value for BindingSourceId to be an explicit nil
func (o *ParameterApiDescriptionModel) SetBindingSourceIdNil() {
	o.BindingSourceId.Set(nil)
}

// UnsetBindingSourceId ensures that no value is present for BindingSourceId, not even an explicit nil
func (o *ParameterApiDescriptionModel) UnsetBindingSourceId() {
	o.BindingSourceId.Unset()
}

// GetDescriptorName returns the DescriptorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterApiDescriptionModel) GetDescriptorName() string {
	if o == nil || o.DescriptorName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DescriptorName.Get()
}

// GetDescriptorNameOk returns a tuple with the DescriptorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterApiDescriptionModel) GetDescriptorNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DescriptorName.Get(), o.DescriptorName.IsSet()
}

// HasDescriptorName returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasDescriptorName() bool {
	if o != nil && o.DescriptorName.IsSet() {
		return true
	}

	return false
}

// SetDescriptorName gets a reference to the given NullableString and assigns it to the DescriptorName field.
func (o *ParameterApiDescriptionModel) SetDescriptorName(v string) {
	o.DescriptorName.Set(&v)
}
// SetDescriptorNameNil sets the value for DescriptorName to be an explicit nil
func (o *ParameterApiDescriptionModel) SetDescriptorNameNil() {
	o.DescriptorName.Set(nil)
}

// UnsetDescriptorName ensures that no value is present for DescriptorName, not even an explicit nil
func (o *ParameterApiDescriptionModel) UnsetDescriptorName() {
	o.DescriptorName.Unset()
}

func (o ParameterApiDescriptionModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NameOnMethod.IsSet() {
		toSerialize["nameOnMethod"] = o.NameOnMethod.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.JsonName.IsSet() {
		toSerialize["jsonName"] = o.JsonName.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.TypeSimple.IsSet() {
		toSerialize["typeSimple"] = o.TypeSimple.Get()
	}
	if o.IsOptional != nil {
		toSerialize["isOptional"] = o.IsOptional
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.ConstraintTypes != nil {
		toSerialize["constraintTypes"] = o.ConstraintTypes
	}
	if o.BindingSourceId.IsSet() {
		toSerialize["bindingSourceId"] = o.BindingSourceId.Get()
	}
	if o.DescriptorName.IsSet() {
		toSerialize["descriptorName"] = o.DescriptorName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableParameterApiDescriptionModel struct {
	value *ParameterApiDescriptionModel
	isSet bool
}

func (v NullableParameterApiDescriptionModel) Get() *ParameterApiDescriptionModel {
	return v.value
}

func (v *NullableParameterApiDescriptionModel) Set(val *ParameterApiDescriptionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterApiDescriptionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterApiDescriptionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterApiDescriptionModel(val *ParameterApiDescriptionModel) *NullableParameterApiDescriptionModel {
	return &NullableParameterApiDescriptionModel{value: val, isSet: true}
}

func (v NullableParameterApiDescriptionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterApiDescriptionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


