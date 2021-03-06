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

// IdentityUserUpdateDto struct for IdentityUserUpdateDto
type IdentityUserUpdateDto struct {
	ExtraProperties map[string]interface{} `json:"extraProperties,omitempty"`
	UserName string `json:"userName"`
	Name NullableString `json:"name,omitempty"`
	Surname NullableString `json:"surname,omitempty"`
	Email string `json:"email"`
	PhoneNumber NullableString `json:"phoneNumber,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	LockoutEnabled *bool `json:"lockoutEnabled,omitempty"`
	RoleNames []string `json:"roleNames,omitempty"`
	Password NullableString `json:"password,omitempty"`
	ConcurrencyStamp NullableString `json:"concurrencyStamp,omitempty"`
}

// NewIdentityUserUpdateDto instantiates a new IdentityUserUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityUserUpdateDto(userName string, email string) *IdentityUserUpdateDto {
	this := IdentityUserUpdateDto{}
	this.UserName = userName
	this.Email = email
	return &this
}

// NewIdentityUserUpdateDtoWithDefaults instantiates a new IdentityUserUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityUserUpdateDtoWithDefaults() *IdentityUserUpdateDto {
	this := IdentityUserUpdateDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserUpdateDto) GetExtraProperties() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserUpdateDto) GetExtraPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExtraProperties == nil {
		return nil, false
	}
	return &o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasExtraProperties() bool {
	if o != nil && o.ExtraProperties != nil {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]interface{} and assigns it to the ExtraProperties field.
func (o *IdentityUserUpdateDto) SetExtraProperties(v map[string]interface{}) {
	o.ExtraProperties = v
}

// GetUserName returns the UserName field value
func (o *IdentityUserUpdateDto) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *IdentityUserUpdateDto) SetUserName(v string) {
	o.UserName = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserUpdateDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserUpdateDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IdentityUserUpdateDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IdentityUserUpdateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IdentityUserUpdateDto) UnsetName() {
	o.Name.Unset()
}

// GetSurname returns the Surname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserUpdateDto) GetSurname() string {
	if o == nil || o.Surname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Surname.Get()
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserUpdateDto) GetSurnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Surname.Get(), o.Surname.IsSet()
}

// HasSurname returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasSurname() bool {
	if o != nil && o.Surname.IsSet() {
		return true
	}

	return false
}

// SetSurname gets a reference to the given NullableString and assigns it to the Surname field.
func (o *IdentityUserUpdateDto) SetSurname(v string) {
	o.Surname.Set(&v)
}
// SetSurnameNil sets the value for Surname to be an explicit nil
func (o *IdentityUserUpdateDto) SetSurnameNil() {
	o.Surname.Set(nil)
}

// UnsetSurname ensures that no value is present for Surname, not even an explicit nil
func (o *IdentityUserUpdateDto) UnsetSurname() {
	o.Surname.Unset()
}

// GetEmail returns the Email field value
func (o *IdentityUserUpdateDto) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *IdentityUserUpdateDto) SetEmail(v string) {
	o.Email = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserUpdateDto) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserUpdateDto) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *IdentityUserUpdateDto) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *IdentityUserUpdateDto) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *IdentityUserUpdateDto) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *IdentityUserUpdateDto) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLockoutEnabled returns the LockoutEnabled field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetLockoutEnabled() bool {
	if o == nil || o.LockoutEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LockoutEnabled
}

// GetLockoutEnabledOk returns a tuple with the LockoutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetLockoutEnabledOk() (*bool, bool) {
	if o == nil || o.LockoutEnabled == nil {
		return nil, false
	}
	return o.LockoutEnabled, true
}

// HasLockoutEnabled returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasLockoutEnabled() bool {
	if o != nil && o.LockoutEnabled != nil {
		return true
	}

	return false
}

// SetLockoutEnabled gets a reference to the given bool and assigns it to the LockoutEnabled field.
func (o *IdentityUserUpdateDto) SetLockoutEnabled(v bool) {
	o.LockoutEnabled = &v
}

// GetRoleNames returns the RoleNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserUpdateDto) GetRoleNames() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.RoleNames
}

// GetRoleNamesOk returns a tuple with the RoleNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserUpdateDto) GetRoleNamesOk() (*[]string, bool) {
	if o == nil || o.RoleNames == nil {
		return nil, false
	}
	return &o.RoleNames, true
}

// HasRoleNames returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasRoleNames() bool {
	if o != nil && o.RoleNames != nil {
		return true
	}

	return false
}

// SetRoleNames gets a reference to the given []string and assigns it to the RoleNames field.
func (o *IdentityUserUpdateDto) SetRoleNames(v []string) {
	o.RoleNames = v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserUpdateDto) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserUpdateDto) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *IdentityUserUpdateDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *IdentityUserUpdateDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *IdentityUserUpdateDto) UnsetPassword() {
	o.Password.Unset()
}

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserUpdateDto) GetConcurrencyStamp() string {
	if o == nil || o.ConcurrencyStamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp.Get()
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserUpdateDto) GetConcurrencyStampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConcurrencyStamp.Get(), o.ConcurrencyStamp.IsSet()
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasConcurrencyStamp() bool {
	if o != nil && o.ConcurrencyStamp.IsSet() {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given NullableString and assigns it to the ConcurrencyStamp field.
func (o *IdentityUserUpdateDto) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp.Set(&v)
}
// SetConcurrencyStampNil sets the value for ConcurrencyStamp to be an explicit nil
func (o *IdentityUserUpdateDto) SetConcurrencyStampNil() {
	o.ConcurrencyStamp.Set(nil)
}

// UnsetConcurrencyStamp ensures that no value is present for ConcurrencyStamp, not even an explicit nil
func (o *IdentityUserUpdateDto) UnsetConcurrencyStamp() {
	o.ConcurrencyStamp.Unset()
}

func (o IdentityUserUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtraProperties != nil {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	if true {
		toSerialize["userName"] = o.UserName
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Surname.IsSet() {
		toSerialize["surname"] = o.Surname.Get()
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phoneNumber"] = o.PhoneNumber.Get()
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	if o.LockoutEnabled != nil {
		toSerialize["lockoutEnabled"] = o.LockoutEnabled
	}
	if o.RoleNames != nil {
		toSerialize["roleNames"] = o.RoleNames
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.ConcurrencyStamp.IsSet() {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityUserUpdateDto struct {
	value *IdentityUserUpdateDto
	isSet bool
}

func (v NullableIdentityUserUpdateDto) Get() *IdentityUserUpdateDto {
	return v.value
}

func (v *NullableIdentityUserUpdateDto) Set(val *IdentityUserUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUserUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUserUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUserUpdateDto(val *IdentityUserUpdateDto) *NullableIdentityUserUpdateDto {
	return &NullableIdentityUserUpdateDto{value: val, isSet: true}
}

func (v NullableIdentityUserUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUserUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


