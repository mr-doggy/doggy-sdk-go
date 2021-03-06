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

// RemoteServiceErrorResponse struct for RemoteServiceErrorResponse
type RemoteServiceErrorResponse struct {
	Error *RemoteServiceErrorInfo `json:"error,omitempty"`
}

// NewRemoteServiceErrorResponse instantiates a new RemoteServiceErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteServiceErrorResponse() *RemoteServiceErrorResponse {
	this := RemoteServiceErrorResponse{}
	return &this
}

// NewRemoteServiceErrorResponseWithDefaults instantiates a new RemoteServiceErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteServiceErrorResponseWithDefaults() *RemoteServiceErrorResponse {
	this := RemoteServiceErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RemoteServiceErrorResponse) GetError() RemoteServiceErrorInfo {
	if o == nil || o.Error == nil {
		var ret RemoteServiceErrorInfo
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteServiceErrorResponse) GetErrorOk() (*RemoteServiceErrorInfo, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RemoteServiceErrorResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given RemoteServiceErrorInfo and assigns it to the Error field.
func (o *RemoteServiceErrorResponse) SetError(v RemoteServiceErrorInfo) {
	o.Error = &v
}

func (o RemoteServiceErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteServiceErrorResponse struct {
	value *RemoteServiceErrorResponse
	isSet bool
}

func (v NullableRemoteServiceErrorResponse) Get() *RemoteServiceErrorResponse {
	return v.value
}

func (v *NullableRemoteServiceErrorResponse) Set(val *RemoteServiceErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteServiceErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteServiceErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteServiceErrorResponse(val *RemoteServiceErrorResponse) *NullableRemoteServiceErrorResponse {
	return &NullableRemoteServiceErrorResponse{value: val, isSet: true}
}

func (v NullableRemoteServiceErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteServiceErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


