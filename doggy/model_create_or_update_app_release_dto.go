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

// CreateOrUpdateAppReleaseDto struct for CreateOrUpdateAppReleaseDto
type CreateOrUpdateAppReleaseDto struct {
	Version NullableString `json:"version,omitempty"`
	Notes NullableString `json:"notes,omitempty"`
	Platform NullableString `json:"platform,omitempty"`
	Key NullableString `json:"key,omitempty"`
	RapidCode NullableString `json:"rapidCode,omitempty"`
	Size NullableInt64 `json:"size,omitempty"`
	Md5 NullableString `json:"md5,omitempty"`
	SliceMd5 NullableString `json:"sliceMd5,omitempty"`
	ProductType NullableString `json:"productType,omitempty"`
	IsForceUpdate *bool `json:"isForceUpdate,omitempty"`
	AppId *string `json:"appId,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	Channel NullableString `json:"channel,omitempty"`
	Environment NullableString `json:"environment,omitempty"`
}

// NewCreateOrUpdateAppReleaseDto instantiates a new CreateOrUpdateAppReleaseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateAppReleaseDto() *CreateOrUpdateAppReleaseDto {
	this := CreateOrUpdateAppReleaseDto{}
	return &this
}

// NewCreateOrUpdateAppReleaseDtoWithDefaults instantiates a new CreateOrUpdateAppReleaseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateAppReleaseDtoWithDefaults() *CreateOrUpdateAppReleaseDto {
	this := CreateOrUpdateAppReleaseDto{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *CreateOrUpdateAppReleaseDto) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetVersion() {
	o.Version.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetNotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *CreateOrUpdateAppReleaseDto) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetNotes() {
	o.Notes.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetPlatform() string {
	if o == nil || o.Platform.Get() == nil {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetPlatformOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *CreateOrUpdateAppReleaseDto) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetPlatform() {
	o.Platform.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *CreateOrUpdateAppReleaseDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetKey() {
	o.Key.Unset()
}

// GetRapidCode returns the RapidCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetRapidCode() string {
	if o == nil || o.RapidCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.RapidCode.Get()
}

// GetRapidCodeOk returns a tuple with the RapidCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetRapidCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RapidCode.Get(), o.RapidCode.IsSet()
}

// HasRapidCode returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasRapidCode() bool {
	if o != nil && o.RapidCode.IsSet() {
		return true
	}

	return false
}

// SetRapidCode gets a reference to the given NullableString and assigns it to the RapidCode field.
func (o *CreateOrUpdateAppReleaseDto) SetRapidCode(v string) {
	o.RapidCode.Set(&v)
}
// SetRapidCodeNil sets the value for RapidCode to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetRapidCodeNil() {
	o.RapidCode.Set(nil)
}

// UnsetRapidCode ensures that no value is present for RapidCode, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetRapidCode() {
	o.RapidCode.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetSize() int64 {
	if o == nil || o.Size.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt64 and assigns it to the Size field.
func (o *CreateOrUpdateAppReleaseDto) SetSize(v int64) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetSize() {
	o.Size.Unset()
}

// GetMd5 returns the Md5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetMd5() string {
	if o == nil || o.Md5.Get() == nil {
		var ret string
		return ret
	}
	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetMd5Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// HasMd5 returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasMd5() bool {
	if o != nil && o.Md5.IsSet() {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given NullableString and assigns it to the Md5 field.
func (o *CreateOrUpdateAppReleaseDto) SetMd5(v string) {
	o.Md5.Set(&v)
}
// SetMd5Nil sets the value for Md5 to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetMd5Nil() {
	o.Md5.Set(nil)
}

// UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetMd5() {
	o.Md5.Unset()
}

// GetSliceMd5 returns the SliceMd5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetSliceMd5() string {
	if o == nil || o.SliceMd5.Get() == nil {
		var ret string
		return ret
	}
	return *o.SliceMd5.Get()
}

// GetSliceMd5Ok returns a tuple with the SliceMd5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetSliceMd5Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SliceMd5.Get(), o.SliceMd5.IsSet()
}

// HasSliceMd5 returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasSliceMd5() bool {
	if o != nil && o.SliceMd5.IsSet() {
		return true
	}

	return false
}

// SetSliceMd5 gets a reference to the given NullableString and assigns it to the SliceMd5 field.
func (o *CreateOrUpdateAppReleaseDto) SetSliceMd5(v string) {
	o.SliceMd5.Set(&v)
}
// SetSliceMd5Nil sets the value for SliceMd5 to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetSliceMd5Nil() {
	o.SliceMd5.Set(nil)
}

// UnsetSliceMd5 ensures that no value is present for SliceMd5, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetSliceMd5() {
	o.SliceMd5.Unset()
}

// GetProductType returns the ProductType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetProductType() string {
	if o == nil || o.ProductType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProductType.Get()
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetProductTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProductType.Get(), o.ProductType.IsSet()
}

// HasProductType returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasProductType() bool {
	if o != nil && o.ProductType.IsSet() {
		return true
	}

	return false
}

// SetProductType gets a reference to the given NullableString and assigns it to the ProductType field.
func (o *CreateOrUpdateAppReleaseDto) SetProductType(v string) {
	o.ProductType.Set(&v)
}
// SetProductTypeNil sets the value for ProductType to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetProductTypeNil() {
	o.ProductType.Set(nil)
}

// UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetProductType() {
	o.ProductType.Unset()
}

// GetIsForceUpdate returns the IsForceUpdate field value if set, zero value otherwise.
func (o *CreateOrUpdateAppReleaseDto) GetIsForceUpdate() bool {
	if o == nil || o.IsForceUpdate == nil {
		var ret bool
		return ret
	}
	return *o.IsForceUpdate
}

// GetIsForceUpdateOk returns a tuple with the IsForceUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppReleaseDto) GetIsForceUpdateOk() (*bool, bool) {
	if o == nil || o.IsForceUpdate == nil {
		return nil, false
	}
	return o.IsForceUpdate, true
}

// HasIsForceUpdate returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasIsForceUpdate() bool {
	if o != nil && o.IsForceUpdate != nil {
		return true
	}

	return false
}

// SetIsForceUpdate gets a reference to the given bool and assigns it to the IsForceUpdate field.
func (o *CreateOrUpdateAppReleaseDto) SetIsForceUpdate(v bool) {
	o.IsForceUpdate = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *CreateOrUpdateAppReleaseDto) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppReleaseDto) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *CreateOrUpdateAppReleaseDto) SetAppId(v string) {
	o.AppId = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *CreateOrUpdateAppReleaseDto) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppReleaseDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *CreateOrUpdateAppReleaseDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetChannel() string {
	if o == nil || o.Channel.Get() == nil {
		var ret string
		return ret
	}
	return *o.Channel.Get()
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetChannelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Channel.Get(), o.Channel.IsSet()
}

// HasChannel returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasChannel() bool {
	if o != nil && o.Channel.IsSet() {
		return true
	}

	return false
}

// SetChannel gets a reference to the given NullableString and assigns it to the Channel field.
func (o *CreateOrUpdateAppReleaseDto) SetChannel(v string) {
	o.Channel.Set(&v)
}
// SetChannelNil sets the value for Channel to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetChannelNil() {
	o.Channel.Set(nil)
}

// UnsetChannel ensures that no value is present for Channel, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetChannel() {
	o.Channel.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppReleaseDto) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppReleaseDto) GetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CreateOrUpdateAppReleaseDto) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *CreateOrUpdateAppReleaseDto) SetEnvironment(v string) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *CreateOrUpdateAppReleaseDto) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *CreateOrUpdateAppReleaseDto) UnsetEnvironment() {
	o.Environment.Unset()
}

func (o CreateOrUpdateAppReleaseDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.RapidCode.IsSet() {
		toSerialize["rapidCode"] = o.RapidCode.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.Md5.IsSet() {
		toSerialize["md5"] = o.Md5.Get()
	}
	if o.SliceMd5.IsSet() {
		toSerialize["sliceMd5"] = o.SliceMd5.Get()
	}
	if o.ProductType.IsSet() {
		toSerialize["productType"] = o.ProductType.Get()
	}
	if o.IsForceUpdate != nil {
		toSerialize["isForceUpdate"] = o.IsForceUpdate
	}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.IsEnabled != nil {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.Channel.IsSet() {
		toSerialize["channel"] = o.Channel.Get()
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrUpdateAppReleaseDto struct {
	value *CreateOrUpdateAppReleaseDto
	isSet bool
}

func (v NullableCreateOrUpdateAppReleaseDto) Get() *CreateOrUpdateAppReleaseDto {
	return v.value
}

func (v *NullableCreateOrUpdateAppReleaseDto) Set(val *CreateOrUpdateAppReleaseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateAppReleaseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateAppReleaseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateAppReleaseDto(val *CreateOrUpdateAppReleaseDto) *NullableCreateOrUpdateAppReleaseDto {
	return &NullableCreateOrUpdateAppReleaseDto{value: val, isSet: true}
}

func (v NullableCreateOrUpdateAppReleaseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateAppReleaseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


