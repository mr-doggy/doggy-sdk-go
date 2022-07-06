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

// CreateOrUpdateAppDto struct for CreateOrUpdateAppDto
type CreateOrUpdateAppDto struct {
	Name NullableString `json:"name,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	Framework NullableString `json:"framework,omitempty"`
	AppType NullableString `json:"appType,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	GitRepository NullableString `json:"gitRepository,omitempty"`
	GitRepositoryType NullableString `json:"gitRepositoryType,omitempty"`
}

// NewCreateOrUpdateAppDto instantiates a new CreateOrUpdateAppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateAppDto() *CreateOrUpdateAppDto {
	this := CreateOrUpdateAppDto{}
	return &this
}

// NewCreateOrUpdateAppDtoWithDefaults instantiates a new CreateOrUpdateAppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateAppDtoWithDefaults() *CreateOrUpdateAppDto {
	this := CreateOrUpdateAppDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateOrUpdateAppDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateOrUpdateAppDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateOrUpdateAppDto) UnsetName() {
	o.Name.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppDto) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppDto) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *CreateOrUpdateAppDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *CreateOrUpdateAppDto) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *CreateOrUpdateAppDto) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetFramework returns the Framework field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppDto) GetFramework() string {
	if o == nil || o.Framework.Get() == nil {
		var ret string
		return ret
	}
	return *o.Framework.Get()
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppDto) GetFrameworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Framework.Get(), o.Framework.IsSet()
}

// HasFramework returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasFramework() bool {
	if o != nil && o.Framework.IsSet() {
		return true
	}

	return false
}

// SetFramework gets a reference to the given NullableString and assigns it to the Framework field.
func (o *CreateOrUpdateAppDto) SetFramework(v string) {
	o.Framework.Set(&v)
}
// SetFrameworkNil sets the value for Framework to be an explicit nil
func (o *CreateOrUpdateAppDto) SetFrameworkNil() {
	o.Framework.Set(nil)
}

// UnsetFramework ensures that no value is present for Framework, not even an explicit nil
func (o *CreateOrUpdateAppDto) UnsetFramework() {
	o.Framework.Unset()
}

// GetAppType returns the AppType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppDto) GetAppType() string {
	if o == nil || o.AppType.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppType.Get()
}

// GetAppTypeOk returns a tuple with the AppType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppDto) GetAppTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppType.Get(), o.AppType.IsSet()
}

// HasAppType returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasAppType() bool {
	if o != nil && o.AppType.IsSet() {
		return true
	}

	return false
}

// SetAppType gets a reference to the given NullableString and assigns it to the AppType field.
func (o *CreateOrUpdateAppDto) SetAppType(v string) {
	o.AppType.Set(&v)
}
// SetAppTypeNil sets the value for AppType to be an explicit nil
func (o *CreateOrUpdateAppDto) SetAppTypeNil() {
	o.AppType.Set(nil)
}

// UnsetAppType ensures that no value is present for AppType, not even an explicit nil
func (o *CreateOrUpdateAppDto) UnsetAppType() {
	o.AppType.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppDto) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateOrUpdateAppDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateOrUpdateAppDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateOrUpdateAppDto) UnsetDescription() {
	o.Description.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppDto) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppDto) GetIconOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *CreateOrUpdateAppDto) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *CreateOrUpdateAppDto) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *CreateOrUpdateAppDto) UnsetIcon() {
	o.Icon.Unset()
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppDto) GetGitRepository() string {
	if o == nil || o.GitRepository.Get() == nil {
		var ret string
		return ret
	}
	return *o.GitRepository.Get()
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppDto) GetGitRepositoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GitRepository.Get(), o.GitRepository.IsSet()
}

// HasGitRepository returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasGitRepository() bool {
	if o != nil && o.GitRepository.IsSet() {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given NullableString and assigns it to the GitRepository field.
func (o *CreateOrUpdateAppDto) SetGitRepository(v string) {
	o.GitRepository.Set(&v)
}
// SetGitRepositoryNil sets the value for GitRepository to be an explicit nil
func (o *CreateOrUpdateAppDto) SetGitRepositoryNil() {
	o.GitRepository.Set(nil)
}

// UnsetGitRepository ensures that no value is present for GitRepository, not even an explicit nil
func (o *CreateOrUpdateAppDto) UnsetGitRepository() {
	o.GitRepository.Unset()
}

// GetGitRepositoryType returns the GitRepositoryType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppDto) GetGitRepositoryType() string {
	if o == nil || o.GitRepositoryType.Get() == nil {
		var ret string
		return ret
	}
	return *o.GitRepositoryType.Get()
}

// GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppDto) GetGitRepositoryTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GitRepositoryType.Get(), o.GitRepositoryType.IsSet()
}

// HasGitRepositoryType returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasGitRepositoryType() bool {
	if o != nil && o.GitRepositoryType.IsSet() {
		return true
	}

	return false
}

// SetGitRepositoryType gets a reference to the given NullableString and assigns it to the GitRepositoryType field.
func (o *CreateOrUpdateAppDto) SetGitRepositoryType(v string) {
	o.GitRepositoryType.Set(&v)
}
// SetGitRepositoryTypeNil sets the value for GitRepositoryType to be an explicit nil
func (o *CreateOrUpdateAppDto) SetGitRepositoryTypeNil() {
	o.GitRepositoryType.Set(nil)
}

// UnsetGitRepositoryType ensures that no value is present for GitRepositoryType, not even an explicit nil
func (o *CreateOrUpdateAppDto) UnsetGitRepositoryType() {
	o.GitRepositoryType.Unset()
}

func (o CreateOrUpdateAppDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Framework.IsSet() {
		toSerialize["framework"] = o.Framework.Get()
	}
	if o.AppType.IsSet() {
		toSerialize["appType"] = o.AppType.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.GitRepository.IsSet() {
		toSerialize["gitRepository"] = o.GitRepository.Get()
	}
	if o.GitRepositoryType.IsSet() {
		toSerialize["gitRepositoryType"] = o.GitRepositoryType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrUpdateAppDto struct {
	value *CreateOrUpdateAppDto
	isSet bool
}

func (v NullableCreateOrUpdateAppDto) Get() *CreateOrUpdateAppDto {
	return v.value
}

func (v *NullableCreateOrUpdateAppDto) Set(val *CreateOrUpdateAppDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateAppDto(val *CreateOrUpdateAppDto) *NullableCreateOrUpdateAppDto {
	return &NullableCreateOrUpdateAppDto{value: val, isSet: true}
}

func (v NullableCreateOrUpdateAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

