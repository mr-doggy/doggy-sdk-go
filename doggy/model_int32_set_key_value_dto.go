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

// Int32SetKeyValueDto struct for Int32SetKeyValueDto
type Int32SetKeyValueDto struct {
	Value *int32 `json:"value,omitempty"`
	DurationSeconds NullableFloat64 `json:"durationSeconds,omitempty"`
}

// NewInt32SetKeyValueDto instantiates a new Int32SetKeyValueDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInt32SetKeyValueDto() *Int32SetKeyValueDto {
	this := Int32SetKeyValueDto{}
	return &this
}

// NewInt32SetKeyValueDtoWithDefaults instantiates a new Int32SetKeyValueDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInt32SetKeyValueDtoWithDefaults() *Int32SetKeyValueDto {
	this := Int32SetKeyValueDto{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Int32SetKeyValueDto) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Int32SetKeyValueDto) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Int32SetKeyValueDto) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *Int32SetKeyValueDto) SetValue(v int32) {
	o.Value = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Int32SetKeyValueDto) GetDurationSeconds() float64 {
	if o == nil || o.DurationSeconds.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DurationSeconds.Get()
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Int32SetKeyValueDto) GetDurationSecondsOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DurationSeconds.Get(), o.DurationSeconds.IsSet()
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *Int32SetKeyValueDto) HasDurationSeconds() bool {
	if o != nil && o.DurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given NullableFloat64 and assigns it to the DurationSeconds field.
func (o *Int32SetKeyValueDto) SetDurationSeconds(v float64) {
	o.DurationSeconds.Set(&v)
}
// SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil
func (o *Int32SetKeyValueDto) SetDurationSecondsNil() {
	o.DurationSeconds.Set(nil)
}

// UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
func (o *Int32SetKeyValueDto) UnsetDurationSeconds() {
	o.DurationSeconds.Unset()
}

func (o Int32SetKeyValueDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.DurationSeconds.IsSet() {
		toSerialize["durationSeconds"] = o.DurationSeconds.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInt32SetKeyValueDto struct {
	value *Int32SetKeyValueDto
	isSet bool
}

func (v NullableInt32SetKeyValueDto) Get() *Int32SetKeyValueDto {
	return v.value
}

func (v *NullableInt32SetKeyValueDto) Set(val *Int32SetKeyValueDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInt32SetKeyValueDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInt32SetKeyValueDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt32SetKeyValueDto(val *Int32SetKeyValueDto) *NullableInt32SetKeyValueDto {
	return &NullableInt32SetKeyValueDto{value: val, isSet: true}
}

func (v NullableInt32SetKeyValueDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt32SetKeyValueDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


