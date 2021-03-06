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

// SendSmsCodeDto struct for SendSmsCodeDto
type SendSmsCodeDto struct {
	PhoneNumber NullableString `json:"phoneNumber,omitempty"`
}

// NewSendSmsCodeDto instantiates a new SendSmsCodeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendSmsCodeDto() *SendSmsCodeDto {
	this := SendSmsCodeDto{}
	return &this
}

// NewSendSmsCodeDtoWithDefaults instantiates a new SendSmsCodeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendSmsCodeDtoWithDefaults() *SendSmsCodeDto {
	this := SendSmsCodeDto{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendSmsCodeDto) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendSmsCodeDto) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *SendSmsCodeDto) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *SendSmsCodeDto) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *SendSmsCodeDto) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *SendSmsCodeDto) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

func (o SendSmsCodeDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PhoneNumber.IsSet() {
		toSerialize["phoneNumber"] = o.PhoneNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSendSmsCodeDto struct {
	value *SendSmsCodeDto
	isSet bool
}

func (v NullableSendSmsCodeDto) Get() *SendSmsCodeDto {
	return v.value
}

func (v *NullableSendSmsCodeDto) Set(val *SendSmsCodeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSendSmsCodeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSendSmsCodeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendSmsCodeDto(val *SendSmsCodeDto) *NullableSendSmsCodeDto {
	return &NullableSendSmsCodeDto{value: val, isSet: true}
}

func (v NullableSendSmsCodeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendSmsCodeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


