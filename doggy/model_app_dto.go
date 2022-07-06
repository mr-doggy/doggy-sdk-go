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

// AppDto struct for AppDto
type AppDto struct {
	Name NullableString `json:"name,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	Fromework NullableString `json:"fromework,omitempty"`
	AppType NullableString `json:"appType,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	GitRepository NullableString `json:"gitRepository,omitempty"`
	GitRepositoryType NullableString `json:"gitRepositoryType,omitempty"`
	LatestRelease *AppReleaseDto `json:"latestRelease,omitempty"`
}

// NewAppDto instantiates a new AppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDto() *AppDto {
	this := AppDto{}
	return &this
}

// NewAppDtoWithDefaults instantiates a new AppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDtoWithDefaults() *AppDto {
	this := AppDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AppDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AppDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AppDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AppDto) UnsetName() {
	o.Name.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppDto) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppDto) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AppDto) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *AppDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *AppDto) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *AppDto) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetFromework returns the Fromework field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppDto) GetFromework() string {
	if o == nil || o.Fromework.Get() == nil {
		var ret string
		return ret
	}
	return *o.Fromework.Get()
}

// GetFromeworkOk returns a tuple with the Fromework field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppDto) GetFromeworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Fromework.Get(), o.Fromework.IsSet()
}

// HasFromework returns a boolean if a field has been set.
func (o *AppDto) HasFromework() bool {
	if o != nil && o.Fromework.IsSet() {
		return true
	}

	return false
}

// SetFromework gets a reference to the given NullableString and assigns it to the Fromework field.
func (o *AppDto) SetFromework(v string) {
	o.Fromework.Set(&v)
}
// SetFromeworkNil sets the value for Fromework to be an explicit nil
func (o *AppDto) SetFromeworkNil() {
	o.Fromework.Set(nil)
}

// UnsetFromework ensures that no value is present for Fromework, not even an explicit nil
func (o *AppDto) UnsetFromework() {
	o.Fromework.Unset()
}

// GetAppType returns the AppType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppDto) GetAppType() string {
	if o == nil || o.AppType.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppType.Get()
}

// GetAppTypeOk returns a tuple with the AppType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppDto) GetAppTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppType.Get(), o.AppType.IsSet()
}

// HasAppType returns a boolean if a field has been set.
func (o *AppDto) HasAppType() bool {
	if o != nil && o.AppType.IsSet() {
		return true
	}

	return false
}

// SetAppType gets a reference to the given NullableString and assigns it to the AppType field.
func (o *AppDto) SetAppType(v string) {
	o.AppType.Set(&v)
}
// SetAppTypeNil sets the value for AppType to be an explicit nil
func (o *AppDto) SetAppTypeNil() {
	o.AppType.Set(nil)
}

// UnsetAppType ensures that no value is present for AppType, not even an explicit nil
func (o *AppDto) UnsetAppType() {
	o.AppType.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppDto) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AppDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AppDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AppDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AppDto) UnsetDescription() {
	o.Description.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppDto) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppDto) GetIconOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *AppDto) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *AppDto) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *AppDto) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *AppDto) UnsetIcon() {
	o.Icon.Unset()
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppDto) GetGitRepository() string {
	if o == nil || o.GitRepository.Get() == nil {
		var ret string
		return ret
	}
	return *o.GitRepository.Get()
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppDto) GetGitRepositoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GitRepository.Get(), o.GitRepository.IsSet()
}

// HasGitRepository returns a boolean if a field has been set.
func (o *AppDto) HasGitRepository() bool {
	if o != nil && o.GitRepository.IsSet() {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given NullableString and assigns it to the GitRepository field.
func (o *AppDto) SetGitRepository(v string) {
	o.GitRepository.Set(&v)
}
// SetGitRepositoryNil sets the value for GitRepository to be an explicit nil
func (o *AppDto) SetGitRepositoryNil() {
	o.GitRepository.Set(nil)
}

// UnsetGitRepository ensures that no value is present for GitRepository, not even an explicit nil
func (o *AppDto) UnsetGitRepository() {
	o.GitRepository.Unset()
}

// GetGitRepositoryType returns the GitRepositoryType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppDto) GetGitRepositoryType() string {
	if o == nil || o.GitRepositoryType.Get() == nil {
		var ret string
		return ret
	}
	return *o.GitRepositoryType.Get()
}

// GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppDto) GetGitRepositoryTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GitRepositoryType.Get(), o.GitRepositoryType.IsSet()
}

// HasGitRepositoryType returns a boolean if a field has been set.
func (o *AppDto) HasGitRepositoryType() bool {
	if o != nil && o.GitRepositoryType.IsSet() {
		return true
	}

	return false
}

// SetGitRepositoryType gets a reference to the given NullableString and assigns it to the GitRepositoryType field.
func (o *AppDto) SetGitRepositoryType(v string) {
	o.GitRepositoryType.Set(&v)
}
// SetGitRepositoryTypeNil sets the value for GitRepositoryType to be an explicit nil
func (o *AppDto) SetGitRepositoryTypeNil() {
	o.GitRepositoryType.Set(nil)
}

// UnsetGitRepositoryType ensures that no value is present for GitRepositoryType, not even an explicit nil
func (o *AppDto) UnsetGitRepositoryType() {
	o.GitRepositoryType.Unset()
}

// GetLatestRelease returns the LatestRelease field value if set, zero value otherwise.
func (o *AppDto) GetLatestRelease() AppReleaseDto {
	if o == nil || o.LatestRelease == nil {
		var ret AppReleaseDto
		return ret
	}
	return *o.LatestRelease
}

// GetLatestReleaseOk returns a tuple with the LatestRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetLatestReleaseOk() (*AppReleaseDto, bool) {
	if o == nil || o.LatestRelease == nil {
		return nil, false
	}
	return o.LatestRelease, true
}

// HasLatestRelease returns a boolean if a field has been set.
func (o *AppDto) HasLatestRelease() bool {
	if o != nil && o.LatestRelease != nil {
		return true
	}

	return false
}

// SetLatestRelease gets a reference to the given AppReleaseDto and assigns it to the LatestRelease field.
func (o *AppDto) SetLatestRelease(v AppReleaseDto) {
	o.LatestRelease = &v
}

func (o AppDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Fromework.IsSet() {
		toSerialize["fromework"] = o.Fromework.Get()
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
	if o.LatestRelease != nil {
		toSerialize["latestRelease"] = o.LatestRelease
	}
	return json.Marshal(toSerialize)
}

type NullableAppDto struct {
	value *AppDto
	isSet bool
}

func (v NullableAppDto) Get() *AppDto {
	return v.value
}

func (v *NullableAppDto) Set(val *AppDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDto(val *AppDto) *NullableAppDto {
	return &NullableAppDto{value: val, isSet: true}
}

func (v NullableAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

