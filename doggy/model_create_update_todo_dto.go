/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
	"time"
)

// CreateUpdateTodoDto struct for CreateUpdateTodoDto
type CreateUpdateTodoDto struct {
	Title NullableString `json:"title,omitempty"`
	Priority *Priority `json:"priority,omitempty"`
	TagIds []string `json:"tagIds,omitempty"`
	IsDone *bool `json:"isDone,omitempty"`
	ParentId NullableString `json:"parentId,omitempty"`
	EndAt NullableTime `json:"endAt,omitempty"`
}

// NewCreateUpdateTodoDto instantiates a new CreateUpdateTodoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpdateTodoDto() *CreateUpdateTodoDto {
	this := CreateUpdateTodoDto{}
	return &this
}

// NewCreateUpdateTodoDtoWithDefaults instantiates a new CreateUpdateTodoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpdateTodoDtoWithDefaults() *CreateUpdateTodoDto {
	this := CreateUpdateTodoDto{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpdateTodoDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpdateTodoDto) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateUpdateTodoDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *CreateUpdateTodoDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *CreateUpdateTodoDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *CreateUpdateTodoDto) UnsetTitle() {
	o.Title.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CreateUpdateTodoDto) GetPriority() Priority {
	if o == nil || o.Priority == nil {
		var ret Priority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateTodoDto) GetPriorityOk() (*Priority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CreateUpdateTodoDto) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given Priority and assigns it to the Priority field.
func (o *CreateUpdateTodoDto) SetPriority(v Priority) {
	o.Priority = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpdateTodoDto) GetTagIds() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpdateTodoDto) GetTagIdsOk() (*[]string, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return &o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *CreateUpdateTodoDto) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []string and assigns it to the TagIds field.
func (o *CreateUpdateTodoDto) SetTagIds(v []string) {
	o.TagIds = v
}

// GetIsDone returns the IsDone field value if set, zero value otherwise.
func (o *CreateUpdateTodoDto) GetIsDone() bool {
	if o == nil || o.IsDone == nil {
		var ret bool
		return ret
	}
	return *o.IsDone
}

// GetIsDoneOk returns a tuple with the IsDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateTodoDto) GetIsDoneOk() (*bool, bool) {
	if o == nil || o.IsDone == nil {
		return nil, false
	}
	return o.IsDone, true
}

// HasIsDone returns a boolean if a field has been set.
func (o *CreateUpdateTodoDto) HasIsDone() bool {
	if o != nil && o.IsDone != nil {
		return true
	}

	return false
}

// SetIsDone gets a reference to the given bool and assigns it to the IsDone field.
func (o *CreateUpdateTodoDto) SetIsDone(v bool) {
	o.IsDone = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpdateTodoDto) GetParentId() string {
	if o == nil || o.ParentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpdateTodoDto) GetParentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *CreateUpdateTodoDto) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *CreateUpdateTodoDto) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *CreateUpdateTodoDto) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *CreateUpdateTodoDto) UnsetParentId() {
	o.ParentId.Unset()
}

// GetEndAt returns the EndAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpdateTodoDto) GetEndAt() time.Time {
	if o == nil || o.EndAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndAt.Get()
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpdateTodoDto) GetEndAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndAt.Get(), o.EndAt.IsSet()
}

// HasEndAt returns a boolean if a field has been set.
func (o *CreateUpdateTodoDto) HasEndAt() bool {
	if o != nil && o.EndAt.IsSet() {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given NullableTime and assigns it to the EndAt field.
func (o *CreateUpdateTodoDto) SetEndAt(v time.Time) {
	o.EndAt.Set(&v)
}
// SetEndAtNil sets the value for EndAt to be an explicit nil
func (o *CreateUpdateTodoDto) SetEndAtNil() {
	o.EndAt.Set(nil)
}

// UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
func (o *CreateUpdateTodoDto) UnsetEndAt() {
	o.EndAt.Unset()
}

func (o CreateUpdateTodoDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.TagIds != nil {
		toSerialize["tagIds"] = o.TagIds
	}
	if o.IsDone != nil {
		toSerialize["isDone"] = o.IsDone
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if o.EndAt.IsSet() {
		toSerialize["endAt"] = o.EndAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateUpdateTodoDto struct {
	value *CreateUpdateTodoDto
	isSet bool
}

func (v NullableCreateUpdateTodoDto) Get() *CreateUpdateTodoDto {
	return v.value
}

func (v *NullableCreateUpdateTodoDto) Set(val *CreateUpdateTodoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpdateTodoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpdateTodoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpdateTodoDto(val *CreateUpdateTodoDto) *NullableCreateUpdateTodoDto {
	return &NullableCreateUpdateTodoDto{value: val, isSet: true}
}

func (v NullableCreateUpdateTodoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpdateTodoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

