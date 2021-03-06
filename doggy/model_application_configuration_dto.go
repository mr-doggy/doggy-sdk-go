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

// ApplicationConfigurationDto struct for ApplicationConfigurationDto
type ApplicationConfigurationDto struct {
	Localization *ApplicationLocalizationConfigurationDto `json:"localization,omitempty"`
	Auth *ApplicationAuthConfigurationDto `json:"auth,omitempty"`
	Setting *ApplicationSettingConfigurationDto `json:"setting,omitempty"`
	CurrentUser *CurrentUserDto `json:"currentUser,omitempty"`
	Features *ApplicationFeatureConfigurationDto `json:"features,omitempty"`
	MultiTenancy *MultiTenancyInfoDto `json:"multiTenancy,omitempty"`
	CurrentTenant *CurrentTenantDto `json:"currentTenant,omitempty"`
	Timing *TimingDto `json:"timing,omitempty"`
	Clock *ClockDto `json:"clock,omitempty"`
	ObjectExtensions *ObjectExtensionsDto `json:"objectExtensions,omitempty"`
}

// NewApplicationConfigurationDto instantiates a new ApplicationConfigurationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationConfigurationDto() *ApplicationConfigurationDto {
	this := ApplicationConfigurationDto{}
	return &this
}

// NewApplicationConfigurationDtoWithDefaults instantiates a new ApplicationConfigurationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationConfigurationDtoWithDefaults() *ApplicationConfigurationDto {
	this := ApplicationConfigurationDto{}
	return &this
}

// GetLocalization returns the Localization field value if set, zero value otherwise.
func (o *ApplicationConfigurationDto) GetLocalization() ApplicationLocalizationConfigurationDto {
	if o == nil || o.Localization == nil {
		var ret ApplicationLocalizationConfigurationDto
		return ret
	}
	return *o.Localization
}

// GetLocalizationOk returns a tuple with the Localization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationConfigurationDto) GetLocalizationOk() (*ApplicationLocalizationConfigurationDto, bool) {
	if o == nil || o.Localization == nil {
		return nil, false
	}
	return o.Localization, true
}

// HasLocalization returns a boolean if a field has been set.
func (o *ApplicationConfigurationDto) HasLocalization() bool {
	if o != nil && o.Localization != nil {
		return true
	}

	return false
}

// SetLocalization gets a reference to the given ApplicationLocalizationConfigurationDto and assigns it to the Localization field.
func (o *ApplicationConfigurationDto) SetLocalization(v ApplicationLocalizationConfigurationDto) {
	o.Localization = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ApplicationConfigurationDto) GetAuth() ApplicationAuthConfigurationDto {
	if o == nil || o.Auth == nil {
		var ret ApplicationAuthConfigurationDto
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationConfigurationDto) GetAuthOk() (*ApplicationAuthConfigurationDto, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ApplicationConfigurationDto) HasAuth() bool {
	if o != nil && o.Auth != nil {
		return true
	}

	return false
}

// SetAuth gets a reference to the given ApplicationAuthConfigurationDto and assigns it to the Auth field.
func (o *ApplicationConfigurationDto) SetAuth(v ApplicationAuthConfigurationDto) {
	o.Auth = &v
}

// GetSetting returns the Setting field value if set, zero value otherwise.
func (o *ApplicationConfigurationDto) GetSetting() ApplicationSettingConfigurationDto {
	if o == nil || o.Setting == nil {
		var ret ApplicationSettingConfigurationDto
		return ret
	}
	return *o.Setting
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationConfigurationDto) GetSettingOk() (*ApplicationSettingConfigurationDto, bool) {
	if o == nil || o.Setting == nil {
		return nil, false
	}
	return o.Setting, true
}

// HasSetting returns a boolean if a field has been set.
func (o *ApplicationConfigurationDto) HasSetting() bool {
	if o != nil && o.Setting != nil {
		return true
	}

	return false
}

// SetSetting gets a reference to the given ApplicationSettingConfigurationDto and assigns it to the Setting field.
func (o *ApplicationConfigurationDto) SetSetting(v ApplicationSettingConfigurationDto) {
	o.Setting = &v
}

// GetCurrentUser returns the CurrentUser field value if set, zero value otherwise.
func (o *ApplicationConfigurationDto) GetCurrentUser() CurrentUserDto {
	if o == nil || o.CurrentUser == nil {
		var ret CurrentUserDto
		return ret
	}
	return *o.CurrentUser
}

// GetCurrentUserOk returns a tuple with the CurrentUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationConfigurationDto) GetCurrentUserOk() (*CurrentUserDto, bool) {
	if o == nil || o.CurrentUser == nil {
		return nil, false
	}
	return o.CurrentUser, true
}

// HasCurrentUser returns a boolean if a field has been set.
func (o *ApplicationConfigurationDto) HasCurrentUser() bool {
	if o != nil && o.CurrentUser != nil {
		return true
	}

	return false
}

// SetCurrentUser gets a reference to the given CurrentUserDto and assigns it to the CurrentUser field.
func (o *ApplicationConfigurationDto) SetCurrentUser(v CurrentUserDto) {
	o.CurrentUser = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ApplicationConfigurationDto) GetFeatures() ApplicationFeatureConfigurationDto {
	if o == nil || o.Features == nil {
		var ret ApplicationFeatureConfigurationDto
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationConfigurationDto) GetFeaturesOk() (*ApplicationFeatureConfigurationDto, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ApplicationConfigurationDto) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given ApplicationFeatureConfigurationDto and assigns it to the Features field.
func (o *ApplicationConfigurationDto) SetFeatures(v ApplicationFeatureConfigurationDto) {
	o.Features = &v
}

// GetMultiTenancy returns the MultiTenancy field value if set, zero value otherwise.
func (o *ApplicationConfigurationDto) GetMultiTenancy() MultiTenancyInfoDto {
	if o == nil || o.MultiTenancy == nil {
		var ret MultiTenancyInfoDto
		return ret
	}
	return *o.MultiTenancy
}

// GetMultiTenancyOk returns a tuple with the MultiTenancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationConfigurationDto) GetMultiTenancyOk() (*MultiTenancyInfoDto, bool) {
	if o == nil || o.MultiTenancy == nil {
		return nil, false
	}
	return o.MultiTenancy, true
}

// HasMultiTenancy returns a boolean if a field has been set.
func (o *ApplicationConfigurationDto) HasMultiTenancy() bool {
	if o != nil && o.MultiTenancy != nil {
		return true
	}

	return false
}

// SetMultiTenancy gets a reference to the given MultiTenancyInfoDto and assigns it to the MultiTenancy field.
func (o *ApplicationConfigurationDto) SetMultiTenancy(v MultiTenancyInfoDto) {
	o.MultiTenancy = &v
}

// GetCurrentTenant returns the CurrentTenant field value if set, zero value otherwise.
func (o *ApplicationConfigurationDto) GetCurrentTenant() CurrentTenantDto {
	if o == nil || o.CurrentTenant == nil {
		var ret CurrentTenantDto
		return ret
	}
	return *o.CurrentTenant
}

// GetCurrentTenantOk returns a tuple with the CurrentTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationConfigurationDto) GetCurrentTenantOk() (*CurrentTenantDto, bool) {
	if o == nil || o.CurrentTenant == nil {
		return nil, false
	}
	return o.CurrentTenant, true
}

// HasCurrentTenant returns a boolean if a field has been set.
func (o *ApplicationConfigurationDto) HasCurrentTenant() bool {
	if o != nil && o.CurrentTenant != nil {
		return true
	}

	return false
}

// SetCurrentTenant gets a reference to the given CurrentTenantDto and assigns it to the CurrentTenant field.
func (o *ApplicationConfigurationDto) SetCurrentTenant(v CurrentTenantDto) {
	o.CurrentTenant = &v
}

// GetTiming returns the Timing field value if set, zero value otherwise.
func (o *ApplicationConfigurationDto) GetTiming() TimingDto {
	if o == nil || o.Timing == nil {
		var ret TimingDto
		return ret
	}
	return *o.Timing
}

// GetTimingOk returns a tuple with the Timing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationConfigurationDto) GetTimingOk() (*TimingDto, bool) {
	if o == nil || o.Timing == nil {
		return nil, false
	}
	return o.Timing, true
}

// HasTiming returns a boolean if a field has been set.
func (o *ApplicationConfigurationDto) HasTiming() bool {
	if o != nil && o.Timing != nil {
		return true
	}

	return false
}

// SetTiming gets a reference to the given TimingDto and assigns it to the Timing field.
func (o *ApplicationConfigurationDto) SetTiming(v TimingDto) {
	o.Timing = &v
}

// GetClock returns the Clock field value if set, zero value otherwise.
func (o *ApplicationConfigurationDto) GetClock() ClockDto {
	if o == nil || o.Clock == nil {
		var ret ClockDto
		return ret
	}
	return *o.Clock
}

// GetClockOk returns a tuple with the Clock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationConfigurationDto) GetClockOk() (*ClockDto, bool) {
	if o == nil || o.Clock == nil {
		return nil, false
	}
	return o.Clock, true
}

// HasClock returns a boolean if a field has been set.
func (o *ApplicationConfigurationDto) HasClock() bool {
	if o != nil && o.Clock != nil {
		return true
	}

	return false
}

// SetClock gets a reference to the given ClockDto and assigns it to the Clock field.
func (o *ApplicationConfigurationDto) SetClock(v ClockDto) {
	o.Clock = &v
}

// GetObjectExtensions returns the ObjectExtensions field value if set, zero value otherwise.
func (o *ApplicationConfigurationDto) GetObjectExtensions() ObjectExtensionsDto {
	if o == nil || o.ObjectExtensions == nil {
		var ret ObjectExtensionsDto
		return ret
	}
	return *o.ObjectExtensions
}

// GetObjectExtensionsOk returns a tuple with the ObjectExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationConfigurationDto) GetObjectExtensionsOk() (*ObjectExtensionsDto, bool) {
	if o == nil || o.ObjectExtensions == nil {
		return nil, false
	}
	return o.ObjectExtensions, true
}

// HasObjectExtensions returns a boolean if a field has been set.
func (o *ApplicationConfigurationDto) HasObjectExtensions() bool {
	if o != nil && o.ObjectExtensions != nil {
		return true
	}

	return false
}

// SetObjectExtensions gets a reference to the given ObjectExtensionsDto and assigns it to the ObjectExtensions field.
func (o *ApplicationConfigurationDto) SetObjectExtensions(v ObjectExtensionsDto) {
	o.ObjectExtensions = &v
}

func (o ApplicationConfigurationDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Localization != nil {
		toSerialize["localization"] = o.Localization
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if o.Setting != nil {
		toSerialize["setting"] = o.Setting
	}
	if o.CurrentUser != nil {
		toSerialize["currentUser"] = o.CurrentUser
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.MultiTenancy != nil {
		toSerialize["multiTenancy"] = o.MultiTenancy
	}
	if o.CurrentTenant != nil {
		toSerialize["currentTenant"] = o.CurrentTenant
	}
	if o.Timing != nil {
		toSerialize["timing"] = o.Timing
	}
	if o.Clock != nil {
		toSerialize["clock"] = o.Clock
	}
	if o.ObjectExtensions != nil {
		toSerialize["objectExtensions"] = o.ObjectExtensions
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationConfigurationDto struct {
	value *ApplicationConfigurationDto
	isSet bool
}

func (v NullableApplicationConfigurationDto) Get() *ApplicationConfigurationDto {
	return v.value
}

func (v *NullableApplicationConfigurationDto) Set(val *ApplicationConfigurationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationConfigurationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationConfigurationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationConfigurationDto(val *ApplicationConfigurationDto) *NullableApplicationConfigurationDto {
	return &NullableApplicationConfigurationDto{value: val, isSet: true}
}

func (v NullableApplicationConfigurationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationConfigurationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


