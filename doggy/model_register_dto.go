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

// RegisterDto struct for RegisterDto
type RegisterDto struct {
	ExtraProperties map[string]interface{} `json:"extraProperties,omitempty"`
	UserName string `json:"userName"`
	EmailAddress string `json:"emailAddress"`
	Password string `json:"password"`
	AppName string `json:"appName"`
}

// NewRegisterDto instantiates a new RegisterDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterDto(userName string, emailAddress string, password string, appName string) *RegisterDto {
	this := RegisterDto{}
	this.UserName = userName
	this.EmailAddress = emailAddress
	this.Password = password
	this.AppName = appName
	return &this
}

// NewRegisterDtoWithDefaults instantiates a new RegisterDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterDtoWithDefaults() *RegisterDto {
	this := RegisterDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterDto) GetExtraProperties() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterDto) GetExtraPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExtraProperties == nil {
		return nil, false
	}
	return &o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *RegisterDto) HasExtraProperties() bool {
	if o != nil && o.ExtraProperties != nil {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]interface{} and assigns it to the ExtraProperties field.
func (o *RegisterDto) SetExtraProperties(v map[string]interface{}) {
	o.ExtraProperties = v
}

// GetUserName returns the UserName field value
func (o *RegisterDto) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *RegisterDto) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *RegisterDto) SetUserName(v string) {
	o.UserName = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *RegisterDto) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *RegisterDto) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *RegisterDto) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetPassword returns the Password field value
func (o *RegisterDto) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *RegisterDto) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *RegisterDto) SetPassword(v string) {
	o.Password = v
}

// GetAppName returns the AppName field value
func (o *RegisterDto) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *RegisterDto) GetAppNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *RegisterDto) SetAppName(v string) {
	o.AppName = v
}

func (o RegisterDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtraProperties != nil {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	if true {
		toSerialize["userName"] = o.UserName
	}
	if true {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["appName"] = o.AppName
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterDto struct {
	value *RegisterDto
	isSet bool
}

func (v NullableRegisterDto) Get() *RegisterDto {
	return v.value
}

func (v *NullableRegisterDto) Set(val *RegisterDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterDto(val *RegisterDto) *NullableRegisterDto {
	return &NullableRegisterDto{value: val, isSet: true}
}

func (v NullableRegisterDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


