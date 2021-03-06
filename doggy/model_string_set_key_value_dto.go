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

// StringSetKeyValueDto struct for StringSetKeyValueDto
type StringSetKeyValueDto struct {
	Value NullableString `json:"value,omitempty"`
	DurationSeconds NullableFloat64 `json:"durationSeconds,omitempty"`
}

// NewStringSetKeyValueDto instantiates a new StringSetKeyValueDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringSetKeyValueDto() *StringSetKeyValueDto {
	this := StringSetKeyValueDto{}
	return &this
}

// NewStringSetKeyValueDtoWithDefaults instantiates a new StringSetKeyValueDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringSetKeyValueDtoWithDefaults() *StringSetKeyValueDto {
	this := StringSetKeyValueDto{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringSetKeyValueDto) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringSetKeyValueDto) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *StringSetKeyValueDto) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *StringSetKeyValueDto) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *StringSetKeyValueDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *StringSetKeyValueDto) UnsetValue() {
	o.Value.Unset()
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringSetKeyValueDto) GetDurationSeconds() float64 {
	if o == nil || o.DurationSeconds.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DurationSeconds.Get()
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringSetKeyValueDto) GetDurationSecondsOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DurationSeconds.Get(), o.DurationSeconds.IsSet()
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *StringSetKeyValueDto) HasDurationSeconds() bool {
	if o != nil && o.DurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given NullableFloat64 and assigns it to the DurationSeconds field.
func (o *StringSetKeyValueDto) SetDurationSeconds(v float64) {
	o.DurationSeconds.Set(&v)
}
// SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil
func (o *StringSetKeyValueDto) SetDurationSecondsNil() {
	o.DurationSeconds.Set(nil)
}

// UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
func (o *StringSetKeyValueDto) UnsetDurationSeconds() {
	o.DurationSeconds.Unset()
}

func (o StringSetKeyValueDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.DurationSeconds.IsSet() {
		toSerialize["durationSeconds"] = o.DurationSeconds.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableStringSetKeyValueDto struct {
	value *StringSetKeyValueDto
	isSet bool
}

func (v NullableStringSetKeyValueDto) Get() *StringSetKeyValueDto {
	return v.value
}

func (v *NullableStringSetKeyValueDto) Set(val *StringSetKeyValueDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStringSetKeyValueDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStringSetKeyValueDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringSetKeyValueDto(val *StringSetKeyValueDto) *NullableStringSetKeyValueDto {
	return &NullableStringSetKeyValueDto{value: val, isSet: true}
}

func (v NullableStringSetKeyValueDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringSetKeyValueDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


