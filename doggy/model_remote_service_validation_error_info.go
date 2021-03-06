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

// RemoteServiceValidationErrorInfo struct for RemoteServiceValidationErrorInfo
type RemoteServiceValidationErrorInfo struct {
	Message NullableString `json:"message,omitempty"`
	Members []string `json:"members,omitempty"`
}

// NewRemoteServiceValidationErrorInfo instantiates a new RemoteServiceValidationErrorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteServiceValidationErrorInfo() *RemoteServiceValidationErrorInfo {
	this := RemoteServiceValidationErrorInfo{}
	return &this
}

// NewRemoteServiceValidationErrorInfoWithDefaults instantiates a new RemoteServiceValidationErrorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteServiceValidationErrorInfoWithDefaults() *RemoteServiceValidationErrorInfo {
	this := RemoteServiceValidationErrorInfo{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteServiceValidationErrorInfo) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteServiceValidationErrorInfo) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *RemoteServiceValidationErrorInfo) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *RemoteServiceValidationErrorInfo) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *RemoteServiceValidationErrorInfo) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *RemoteServiceValidationErrorInfo) UnsetMessage() {
	o.Message.Unset()
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteServiceValidationErrorInfo) GetMembers() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteServiceValidationErrorInfo) GetMembersOk() (*[]string, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return &o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *RemoteServiceValidationErrorInfo) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *RemoteServiceValidationErrorInfo) SetMembers(v []string) {
	o.Members = v
}

func (o RemoteServiceValidationErrorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteServiceValidationErrorInfo struct {
	value *RemoteServiceValidationErrorInfo
	isSet bool
}

func (v NullableRemoteServiceValidationErrorInfo) Get() *RemoteServiceValidationErrorInfo {
	return v.value
}

func (v *NullableRemoteServiceValidationErrorInfo) Set(val *RemoteServiceValidationErrorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteServiceValidationErrorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteServiceValidationErrorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteServiceValidationErrorInfo(val *RemoteServiceValidationErrorInfo) *NullableRemoteServiceValidationErrorInfo {
	return &NullableRemoteServiceValidationErrorInfo{value: val, isSet: true}
}

func (v NullableRemoteServiceValidationErrorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteServiceValidationErrorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


