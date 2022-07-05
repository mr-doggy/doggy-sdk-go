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

// IdentityUserDto struct for IdentityUserDto
type IdentityUserDto struct {
	ExtraProperties map[string]interface{} `json:"extraProperties,omitempty"`
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId NullableString `json:"creatorId,omitempty"`
	LastModificationTime NullableTime `json:"lastModificationTime,omitempty"`
	LastModifierId NullableString `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId NullableString `json:"deleterId,omitempty"`
	DeletionTime NullableTime `json:"deletionTime,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	UserName NullableString `json:"userName,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Surname NullableString `json:"surname,omitempty"`
	Email NullableString `json:"email,omitempty"`
	EmailConfirmed *bool `json:"emailConfirmed,omitempty"`
	PhoneNumber NullableString `json:"phoneNumber,omitempty"`
	PhoneNumberConfirmed *bool `json:"phoneNumberConfirmed,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	LockoutEnabled *bool `json:"lockoutEnabled,omitempty"`
	LockoutEnd NullableTime `json:"lockoutEnd,omitempty"`
	ConcurrencyStamp NullableString `json:"concurrencyStamp,omitempty"`
}

// NewIdentityUserDto instantiates a new IdentityUserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityUserDto() *IdentityUserDto {
	this := IdentityUserDto{}
	return &this
}

// NewIdentityUserDtoWithDefaults instantiates a new IdentityUserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityUserDtoWithDefaults() *IdentityUserDto {
	this := IdentityUserDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetExtraProperties() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetExtraPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExtraProperties == nil {
		return nil, false
	}
	return &o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *IdentityUserDto) HasExtraProperties() bool {
	if o != nil && o.ExtraProperties != nil {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]interface{} and assigns it to the ExtraProperties field.
func (o *IdentityUserDto) SetExtraProperties(v map[string]interface{}) {
	o.ExtraProperties = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityUserDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityUserDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityUserDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *IdentityUserDto) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *IdentityUserDto) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *IdentityUserDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetCreatorId() string {
	if o == nil || o.CreatorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetCreatorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *IdentityUserDto) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given NullableString and assigns it to the CreatorId field.
func (o *IdentityUserDto) SetCreatorId(v string) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *IdentityUserDto) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *IdentityUserDto) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetLastModificationTime() time.Time {
	if o == nil || o.LastModificationTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime.Get()
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModificationTime.Get(), o.LastModificationTime.IsSet()
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *IdentityUserDto) HasLastModificationTime() bool {
	if o != nil && o.LastModificationTime.IsSet() {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given NullableTime and assigns it to the LastModificationTime field.
func (o *IdentityUserDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime.Set(&v)
}
// SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil
func (o *IdentityUserDto) SetLastModificationTimeNil() {
	o.LastModificationTime.Set(nil)
}

// UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
func (o *IdentityUserDto) UnsetLastModificationTime() {
	o.LastModificationTime.Unset()
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetLastModifierId() string {
	if o == nil || o.LastModifierId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastModifierId.Get()
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetLastModifierIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifierId.Get(), o.LastModifierId.IsSet()
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *IdentityUserDto) HasLastModifierId() bool {
	if o != nil && o.LastModifierId.IsSet() {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given NullableString and assigns it to the LastModifierId field.
func (o *IdentityUserDto) SetLastModifierId(v string) {
	o.LastModifierId.Set(&v)
}
// SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil
func (o *IdentityUserDto) SetLastModifierIdNil() {
	o.LastModifierId.Set(nil)
}

// UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
func (o *IdentityUserDto) UnsetLastModifierId() {
	o.LastModifierId.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *IdentityUserDto) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *IdentityUserDto) HasIsDeleted() bool {
	if o != nil && o.IsDeleted != nil {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *IdentityUserDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetDeleterId() string {
	if o == nil || o.DeleterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeleterId.Get()
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetDeleterIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeleterId.Get(), o.DeleterId.IsSet()
}

// HasDeleterId returns a boolean if a field has been set.
func (o *IdentityUserDto) HasDeleterId() bool {
	if o != nil && o.DeleterId.IsSet() {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given NullableString and assigns it to the DeleterId field.
func (o *IdentityUserDto) SetDeleterId(v string) {
	o.DeleterId.Set(&v)
}
// SetDeleterIdNil sets the value for DeleterId to be an explicit nil
func (o *IdentityUserDto) SetDeleterIdNil() {
	o.DeleterId.Set(nil)
}

// UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
func (o *IdentityUserDto) UnsetDeleterId() {
	o.DeleterId.Unset()
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetDeletionTime() time.Time {
	if o == nil || o.DeletionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime.Get()
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletionTime.Get(), o.DeletionTime.IsSet()
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *IdentityUserDto) HasDeletionTime() bool {
	if o != nil && o.DeletionTime.IsSet() {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given NullableTime and assigns it to the DeletionTime field.
func (o *IdentityUserDto) SetDeletionTime(v time.Time) {
	o.DeletionTime.Set(&v)
}
// SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil
func (o *IdentityUserDto) SetDeletionTimeNil() {
	o.DeletionTime.Set(nil)
}

// UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
func (o *IdentityUserDto) UnsetDeletionTime() {
	o.DeletionTime.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *IdentityUserDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *IdentityUserDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *IdentityUserDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *IdentityUserDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *IdentityUserDto) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *IdentityUserDto) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *IdentityUserDto) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *IdentityUserDto) UnsetUserName() {
	o.UserName.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IdentityUserDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IdentityUserDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IdentityUserDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IdentityUserDto) UnsetName() {
	o.Name.Unset()
}

// GetSurname returns the Surname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetSurname() string {
	if o == nil || o.Surname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Surname.Get()
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetSurnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Surname.Get(), o.Surname.IsSet()
}

// HasSurname returns a boolean if a field has been set.
func (o *IdentityUserDto) HasSurname() bool {
	if o != nil && o.Surname.IsSet() {
		return true
	}

	return false
}

// SetSurname gets a reference to the given NullableString and assigns it to the Surname field.
func (o *IdentityUserDto) SetSurname(v string) {
	o.Surname.Set(&v)
}
// SetSurnameNil sets the value for Surname to be an explicit nil
func (o *IdentityUserDto) SetSurnameNil() {
	o.Surname.Set(nil)
}

// UnsetSurname ensures that no value is present for Surname, not even an explicit nil
func (o *IdentityUserDto) UnsetSurname() {
	o.Surname.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *IdentityUserDto) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *IdentityUserDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *IdentityUserDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *IdentityUserDto) UnsetEmail() {
	o.Email.Unset()
}

// GetEmailConfirmed returns the EmailConfirmed field value if set, zero value otherwise.
func (o *IdentityUserDto) GetEmailConfirmed() bool {
	if o == nil || o.EmailConfirmed == nil {
		var ret bool
		return ret
	}
	return *o.EmailConfirmed
}

// GetEmailConfirmedOk returns a tuple with the EmailConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserDto) GetEmailConfirmedOk() (*bool, bool) {
	if o == nil || o.EmailConfirmed == nil {
		return nil, false
	}
	return o.EmailConfirmed, true
}

// HasEmailConfirmed returns a boolean if a field has been set.
func (o *IdentityUserDto) HasEmailConfirmed() bool {
	if o != nil && o.EmailConfirmed != nil {
		return true
	}

	return false
}

// SetEmailConfirmed gets a reference to the given bool and assigns it to the EmailConfirmed field.
func (o *IdentityUserDto) SetEmailConfirmed(v bool) {
	o.EmailConfirmed = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *IdentityUserDto) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *IdentityUserDto) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *IdentityUserDto) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *IdentityUserDto) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetPhoneNumberConfirmed returns the PhoneNumberConfirmed field value if set, zero value otherwise.
func (o *IdentityUserDto) GetPhoneNumberConfirmed() bool {
	if o == nil || o.PhoneNumberConfirmed == nil {
		var ret bool
		return ret
	}
	return *o.PhoneNumberConfirmed
}

// GetPhoneNumberConfirmedOk returns a tuple with the PhoneNumberConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserDto) GetPhoneNumberConfirmedOk() (*bool, bool) {
	if o == nil || o.PhoneNumberConfirmed == nil {
		return nil, false
	}
	return o.PhoneNumberConfirmed, true
}

// HasPhoneNumberConfirmed returns a boolean if a field has been set.
func (o *IdentityUserDto) HasPhoneNumberConfirmed() bool {
	if o != nil && o.PhoneNumberConfirmed != nil {
		return true
	}

	return false
}

// SetPhoneNumberConfirmed gets a reference to the given bool and assigns it to the PhoneNumberConfirmed field.
func (o *IdentityUserDto) SetPhoneNumberConfirmed(v bool) {
	o.PhoneNumberConfirmed = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *IdentityUserDto) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserDto) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *IdentityUserDto) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *IdentityUserDto) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLockoutEnabled returns the LockoutEnabled field value if set, zero value otherwise.
func (o *IdentityUserDto) GetLockoutEnabled() bool {
	if o == nil || o.LockoutEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LockoutEnabled
}

// GetLockoutEnabledOk returns a tuple with the LockoutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserDto) GetLockoutEnabledOk() (*bool, bool) {
	if o == nil || o.LockoutEnabled == nil {
		return nil, false
	}
	return o.LockoutEnabled, true
}

// HasLockoutEnabled returns a boolean if a field has been set.
func (o *IdentityUserDto) HasLockoutEnabled() bool {
	if o != nil && o.LockoutEnabled != nil {
		return true
	}

	return false
}

// SetLockoutEnabled gets a reference to the given bool and assigns it to the LockoutEnabled field.
func (o *IdentityUserDto) SetLockoutEnabled(v bool) {
	o.LockoutEnabled = &v
}

// GetLockoutEnd returns the LockoutEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetLockoutEnd() time.Time {
	if o == nil || o.LockoutEnd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LockoutEnd.Get()
}

// GetLockoutEndOk returns a tuple with the LockoutEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetLockoutEndOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LockoutEnd.Get(), o.LockoutEnd.IsSet()
}

// HasLockoutEnd returns a boolean if a field has been set.
func (o *IdentityUserDto) HasLockoutEnd() bool {
	if o != nil && o.LockoutEnd.IsSet() {
		return true
	}

	return false
}

// SetLockoutEnd gets a reference to the given NullableTime and assigns it to the LockoutEnd field.
func (o *IdentityUserDto) SetLockoutEnd(v time.Time) {
	o.LockoutEnd.Set(&v)
}
// SetLockoutEndNil sets the value for LockoutEnd to be an explicit nil
func (o *IdentityUserDto) SetLockoutEndNil() {
	o.LockoutEnd.Set(nil)
}

// UnsetLockoutEnd ensures that no value is present for LockoutEnd, not even an explicit nil
func (o *IdentityUserDto) UnsetLockoutEnd() {
	o.LockoutEnd.Unset()
}

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserDto) GetConcurrencyStamp() string {
	if o == nil || o.ConcurrencyStamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp.Get()
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserDto) GetConcurrencyStampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConcurrencyStamp.Get(), o.ConcurrencyStamp.IsSet()
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *IdentityUserDto) HasConcurrencyStamp() bool {
	if o != nil && o.ConcurrencyStamp.IsSet() {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given NullableString and assigns it to the ConcurrencyStamp field.
func (o *IdentityUserDto) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp.Set(&v)
}
// SetConcurrencyStampNil sets the value for ConcurrencyStamp to be an explicit nil
func (o *IdentityUserDto) SetConcurrencyStampNil() {
	o.ConcurrencyStamp.Set(nil)
}

// UnsetConcurrencyStamp ensures that no value is present for ConcurrencyStamp, not even an explicit nil
func (o *IdentityUserDto) UnsetConcurrencyStamp() {
	o.ConcurrencyStamp.Unset()
}

func (o IdentityUserDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtraProperties != nil {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreationTime != nil {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.CreatorId.IsSet() {
		toSerialize["creatorId"] = o.CreatorId.Get()
	}
	if o.LastModificationTime.IsSet() {
		toSerialize["lastModificationTime"] = o.LastModificationTime.Get()
	}
	if o.LastModifierId.IsSet() {
		toSerialize["lastModifierId"] = o.LastModifierId.Get()
	}
	if o.IsDeleted != nil {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if o.DeleterId.IsSet() {
		toSerialize["deleterId"] = o.DeleterId.Get()
	}
	if o.DeletionTime.IsSet() {
		toSerialize["deletionTime"] = o.DeletionTime.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Surname.IsSet() {
		toSerialize["surname"] = o.Surname.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.EmailConfirmed != nil {
		toSerialize["emailConfirmed"] = o.EmailConfirmed
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phoneNumber"] = o.PhoneNumber.Get()
	}
	if o.PhoneNumberConfirmed != nil {
		toSerialize["phoneNumberConfirmed"] = o.PhoneNumberConfirmed
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	if o.LockoutEnabled != nil {
		toSerialize["lockoutEnabled"] = o.LockoutEnabled
	}
	if o.LockoutEnd.IsSet() {
		toSerialize["lockoutEnd"] = o.LockoutEnd.Get()
	}
	if o.ConcurrencyStamp.IsSet() {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityUserDto struct {
	value *IdentityUserDto
	isSet bool
}

func (v NullableIdentityUserDto) Get() *IdentityUserDto {
	return v.value
}

func (v *NullableIdentityUserDto) Set(val *IdentityUserDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUserDto(val *IdentityUserDto) *NullableIdentityUserDto {
	return &NullableIdentityUserDto{value: val, isSet: true}
}

func (v NullableIdentityUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

