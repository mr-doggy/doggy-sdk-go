# Go API client for doggy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./doggy"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppApi* | [**ApiAppAppFileOrCredentialsGet**](docs/AppApi.md#apiappappfileorcredentialsget) | **Get** /api/app/app/file-or-credentials | 
*AppApi* | [**ApiAppAppGet**](docs/AppApi.md#apiappappget) | **Get** /api/app/app | 
*AppApi* | [**ApiAppAppIdDelete**](docs/AppApi.md#apiappappiddelete) | **Delete** /api/app/app/{id} | 
*AppApi* | [**ApiAppAppIdGet**](docs/AppApi.md#apiappappidget) | **Get** /api/app/app/{id} | 
*AppApi* | [**ApiAppAppIdPut**](docs/AppApi.md#apiappappidput) | **Put** /api/app/app/{id} | 
*AppApi* | [**ApiAppAppPost**](docs/AppApi.md#apiappapppost) | **Post** /api/app/app | 
*AppReleaseApi* | [**ApiAppAppReleaseGet**](docs/AppReleaseApi.md#apiappappreleaseget) | **Get** /api/app/app-release | 
*AppReleaseApi* | [**ApiAppAppReleaseIdDelete**](docs/AppReleaseApi.md#apiappappreleaseiddelete) | **Delete** /api/app/app-release/{id} | 
*AppReleaseApi* | [**ApiAppAppReleaseIdGet**](docs/AppReleaseApi.md#apiappappreleaseidget) | **Get** /api/app/app-release/{id} | 
*AppReleaseApi* | [**ApiAppAppReleaseIdPut**](docs/AppReleaseApi.md#apiappappreleaseidput) | **Put** /api/app/app-release/{id} | 
*AppReleaseApi* | [**ApiAppAppReleasePost**](docs/AppReleaseApi.md#apiappappreleasepost) | **Post** /api/app/app-release | 
*DeviceApi* | [**ApiAppDeviceGet**](docs/DeviceApi.md#apiappdeviceget) | **Get** /api/app/device | 
*DeviceApi* | [**ApiAppDeviceIdDelete**](docs/DeviceApi.md#apiappdeviceiddelete) | **Delete** /api/app/device/{id} | 
*DeviceApi* | [**ApiAppDeviceRefreshPost**](docs/DeviceApi.md#apiappdevicerefreshpost) | **Post** /api/app/device/refresh | 
*FileApi* | [**ApiAppFileFileOrCredentialsGet**](docs/FileApi.md#apiappfilefileorcredentialsget) | **Get** /api/app/file/file-or-credentials | 
*FileApi* | [**ApiAppFilePreSignUrlPost**](docs/FileApi.md#apiappfilepresignurlpost) | **Post** /api/app/file/pre-sign-url | 
*FileApi* | [**ApiAppFileUrlPost**](docs/FileApi.md#apiappfileurlpost) | **Post** /api/app/file/url | 
*ItemApi* | [**ApiAppItemPullGet**](docs/ItemApi.md#apiappitempullget) | **Get** /api/app/item/pull | 
*ItemApi* | [**ApiAppItemPushPost**](docs/ItemApi.md#apiappitempushpost) | **Post** /api/app/item/push | 
*ItemApi* | [**ApiAppItemSpecialItemsGet**](docs/ItemApi.md#apiappitemspecialitemsget) | **Get** /api/app/item/special-items | 
*KeyValueApi* | [**ApiAppKeyValueBoolGet**](docs/KeyValueApi.md#apiappkeyvalueboolget) | **Get** /api/app/key-value/bool | 
*KeyValueApi* | [**ApiAppKeyValueDateTimeGet**](docs/KeyValueApi.md#apiappkeyvaluedatetimeget) | **Get** /api/app/key-value/date-time | 
*KeyValueApi* | [**ApiAppKeyValueDecimalGet**](docs/KeyValueApi.md#apiappkeyvaluedecimalget) | **Get** /api/app/key-value/decimal | 
*KeyValueApi* | [**ApiAppKeyValueDoubleGet**](docs/KeyValueApi.md#apiappkeyvaluedoubleget) | **Get** /api/app/key-value/double | 
*KeyValueApi* | [**ApiAppKeyValueIntGet**](docs/KeyValueApi.md#apiappkeyvalueintget) | **Get** /api/app/key-value/int | 
*KeyValueApi* | [**ApiAppKeyValueSetBoolPost**](docs/KeyValueApi.md#apiappkeyvaluesetboolpost) | **Post** /api/app/key-value/set-bool | 
*KeyValueApi* | [**ApiAppKeyValueSetDateTimePost**](docs/KeyValueApi.md#apiappkeyvaluesetdatetimepost) | **Post** /api/app/key-value/set-date-time | 
*KeyValueApi* | [**ApiAppKeyValueSetDecimalPost**](docs/KeyValueApi.md#apiappkeyvaluesetdecimalpost) | **Post** /api/app/key-value/set-decimal | 
*KeyValueApi* | [**ApiAppKeyValueSetDoublePost**](docs/KeyValueApi.md#apiappkeyvaluesetdoublepost) | **Post** /api/app/key-value/set-double | 
*KeyValueApi* | [**ApiAppKeyValueSetIntPost**](docs/KeyValueApi.md#apiappkeyvaluesetintpost) | **Post** /api/app/key-value/set-int | 
*KeyValueApi* | [**ApiAppKeyValueSetStringPost**](docs/KeyValueApi.md#apiappkeyvaluesetstringpost) | **Post** /api/app/key-value/set-string | 
*KeyValueApi* | [**ApiAppKeyValueStringGet**](docs/KeyValueApi.md#apiappkeyvaluestringget) | **Get** /api/app/key-value/string | 
*MemberApi* | [**ApiAppMemberGet**](docs/MemberApi.md#apiappmemberget) | **Get** /api/app/member | 
*NoteApi* | [**ApiAppNoteNoteSpecsGet**](docs/NoteApi.md#apiappnotenotespecsget) | **Get** /api/app/note/note-specs | 
*NotificationApi* | [**ApiAppNotificationBarkApiKeyMessageGet**](docs/NotificationApi.md#apiappnotificationbarkapikeymessageget) | **Get** /api/app/notification/bark/{apiKey}/{message} | 
*NotificationApi* | [**ApiAppNotificationGet**](docs/NotificationApi.md#apiappnotificationget) | **Get** /api/app/notification | 
*NotificationApi* | [**ApiAppNotificationPushPost**](docs/NotificationApi.md#apiappnotificationpushpost) | **Post** /api/app/notification/push | 
*SettingsApi* | [**ApiAppSettingsGet**](docs/SettingsApi.md#apiappsettingsget) | **Get** /api/app/settings | 
*SettingsApi* | [**ApiAppSettingsSetPost**](docs/SettingsApi.md#apiappsettingssetpost) | **Post** /api/app/settings/set | 
*SimpleDataApi* | [**ApiAppSimpleDataGet**](docs/SimpleDataApi.md#apiappsimpledataget) | **Get** /api/app/simple-data | 
*SimpleDataApi* | [**ApiAppSimpleDataIdDelete**](docs/SimpleDataApi.md#apiappsimpledataiddelete) | **Delete** /api/app/simple-data/{id} | 
*SimpleDataApi* | [**ApiAppSimpleDataIdGet**](docs/SimpleDataApi.md#apiappsimpledataidget) | **Get** /api/app/simple-data/{id} | 
*SimpleDataApi* | [**ApiAppSimpleDataSavePost**](docs/SimpleDataApi.md#apiappsimpledatasavepost) | **Post** /api/app/simple-data/save | 
*SmsApi* | [**ApiAppSmsSendChangePhoneCodePost**](docs/SmsApi.md#apiappsmssendchangephonecodepost) | **Post** /api/app/sms/send-change-phone-code | 
*SmsApi* | [**ApiAppSmsSendLoginCodePost**](docs/SmsApi.md#apiappsmssendlogincodepost) | **Post** /api/app/sms/send-login-code | 
*StorageApi* | [**ApiAppStorageGet**](docs/StorageApi.md#apiappstorageget) | **Get** /api/app/storage | 
*SyncStateApi* | [**ApiAppSyncStateGet**](docs/SyncStateApi.md#apiappsyncstateget) | **Get** /api/app/sync-state | 
*SyncStateApi* | [**ApiAppSyncStateItemChangedEtoPost**](docs/SyncStateApi.md#apiappsyncstateitemchangedetopost) | **Post** /api/app/sync-state/item-changed-eto | 


## Documentation For Models

 - [AbpLoginResult](docs/AbpLoginResult.md)
 - [ActionApiDescriptionModel](docs/ActionApiDescriptionModel.md)
 - [AppDto](docs/AppDto.md)
 - [AppDtoPagedResultDto](docs/AppDtoPagedResultDto.md)
 - [AppReleaseDto](docs/AppReleaseDto.md)
 - [AppReleaseDtoPagedResultDto](docs/AppReleaseDtoPagedResultDto.md)
 - [AppTheme](docs/AppTheme.md)
 - [ApplicationApiDescriptionModel](docs/ApplicationApiDescriptionModel.md)
 - [ApplicationAuthConfigurationDto](docs/ApplicationAuthConfigurationDto.md)
 - [ApplicationConfigurationDto](docs/ApplicationConfigurationDto.md)
 - [ApplicationFeatureConfigurationDto](docs/ApplicationFeatureConfigurationDto.md)
 - [ApplicationLocalizationConfigurationDto](docs/ApplicationLocalizationConfigurationDto.md)
 - [ApplicationSettingConfigurationDto](docs/ApplicationSettingConfigurationDto.md)
 - [BooleanKeyValue](docs/BooleanKeyValue.md)
 - [BooleanSetKeyValueDto](docs/BooleanSetKeyValueDto.md)
 - [ChangePasswordInput](docs/ChangePasswordInput.md)
 - [ClockDto](docs/ClockDto.md)
 - [ControllerApiDescriptionModel](docs/ControllerApiDescriptionModel.md)
 - [ControllerInterfaceApiDescriptionModel](docs/ControllerInterfaceApiDescriptionModel.md)
 - [CreateOrUpdateAppDto](docs/CreateOrUpdateAppDto.md)
 - [CreateOrUpdateAppReleaseDto](docs/CreateOrUpdateAppReleaseDto.md)
 - [CreatePushNotificationDto](docs/CreatePushNotificationDto.md)
 - [CreateUpdateItemDto](docs/CreateUpdateItemDto.md)
 - [CreateUpdateNotificationDto](docs/CreateUpdateNotificationDto.md)
 - [CurrentCultureDto](docs/CurrentCultureDto.md)
 - [CurrentTenantDto](docs/CurrentTenantDto.md)
 - [CurrentUserDto](docs/CurrentUserDto.md)
 - [DateTimeFormatDto](docs/DateTimeFormatDto.md)
 - [DateTimeKeyValue](docs/DateTimeKeyValue.md)
 - [DateTimeSetKeyValueDto](docs/DateTimeSetKeyValueDto.md)
 - [DecimalKeyValue](docs/DecimalKeyValue.md)
 - [DecimalSetKeyValueDto](docs/DecimalSetKeyValueDto.md)
 - [DeviceDto](docs/DeviceDto.md)
 - [DeviceDtoPagedResultDto](docs/DeviceDtoPagedResultDto.md)
 - [DoubleKeyValue](docs/DoubleKeyValue.md)
 - [DoubleSetKeyValueDto](docs/DoubleSetKeyValueDto.md)
 - [EmailSettingsDto](docs/EmailSettingsDto.md)
 - [EntityExtensionDto](docs/EntityExtensionDto.md)
 - [ExtensionEnumDto](docs/ExtensionEnumDto.md)
 - [ExtensionEnumFieldDto](docs/ExtensionEnumFieldDto.md)
 - [ExtensionPropertyApiCreateDto](docs/ExtensionPropertyApiCreateDto.md)
 - [ExtensionPropertyApiDto](docs/ExtensionPropertyApiDto.md)
 - [ExtensionPropertyApiGetDto](docs/ExtensionPropertyApiGetDto.md)
 - [ExtensionPropertyApiUpdateDto](docs/ExtensionPropertyApiUpdateDto.md)
 - [ExtensionPropertyAttributeDto](docs/ExtensionPropertyAttributeDto.md)
 - [ExtensionPropertyDto](docs/ExtensionPropertyDto.md)
 - [ExtensionPropertyUiDto](docs/ExtensionPropertyUiDto.md)
 - [ExtensionPropertyUiFormDto](docs/ExtensionPropertyUiFormDto.md)
 - [ExtensionPropertyUiLookupDto](docs/ExtensionPropertyUiLookupDto.md)
 - [ExtensionPropertyUiTableDto](docs/ExtensionPropertyUiTableDto.md)
 - [FeatureDto](docs/FeatureDto.md)
 - [FeatureGroupDto](docs/FeatureGroupDto.md)
 - [FeatureProviderDto](docs/FeatureProviderDto.md)
 - [FileDto](docs/FileDto.md)
 - [FileOrCredentialsDto](docs/FileOrCredentialsDto.md)
 - [FindTenantResultDto](docs/FindTenantResultDto.md)
 - [GetFeatureListResultDto](docs/GetFeatureListResultDto.md)
 - [GetPermissionListResultDto](docs/GetPermissionListResultDto.md)
 - [IStringValueType](docs/IStringValueType.md)
 - [IValueValidator](docs/IValueValidator.md)
 - [IanaTimeZone](docs/IanaTimeZone.md)
 - [IdentityRoleCreateDto](docs/IdentityRoleCreateDto.md)
 - [IdentityRoleDto](docs/IdentityRoleDto.md)
 - [IdentityRoleDtoListResultDto](docs/IdentityRoleDtoListResultDto.md)
 - [IdentityRoleDtoPagedResultDto](docs/IdentityRoleDtoPagedResultDto.md)
 - [IdentityRoleUpdateDto](docs/IdentityRoleUpdateDto.md)
 - [IdentityUserCreateDto](docs/IdentityUserCreateDto.md)
 - [IdentityUserDto](docs/IdentityUserDto.md)
 - [IdentityUserDtoPagedResultDto](docs/IdentityUserDtoPagedResultDto.md)
 - [IdentityUserUpdateDto](docs/IdentityUserUpdateDto.md)
 - [IdentityUserUpdateRolesDto](docs/IdentityUserUpdateRolesDto.md)
 - [Int32KeyValue](docs/Int32KeyValue.md)
 - [Int32SetKeyValueDto](docs/Int32SetKeyValueDto.md)
 - [ItemDto](docs/ItemDto.md)
 - [ItemDtoPagedResultDto](docs/ItemDtoPagedResultDto.md)
 - [LanguageInfo](docs/LanguageInfo.md)
 - [LocalizableStringDto](docs/LocalizableStringDto.md)
 - [LoginResultType](docs/LoginResultType.md)
 - [MemberDto](docs/MemberDto.md)
 - [MethodParameterApiDescriptionModel](docs/MethodParameterApiDescriptionModel.md)
 - [ModuleApiDescriptionModel](docs/ModuleApiDescriptionModel.md)
 - [ModuleExtensionDto](docs/ModuleExtensionDto.md)
 - [MultiTenancyInfoDto](docs/MultiTenancyInfoDto.md)
 - [NameValue](docs/NameValue.md)
 - [NoteSpecDto](docs/NoteSpecDto.md)
 - [NotificationDto](docs/NotificationDto.md)
 - [NotificationDtoPagedResultDto](docs/NotificationDtoPagedResultDto.md)
 - [NotificationInfoDto](docs/NotificationInfoDto.md)
 - [NotificationInfoDtoPagedResultDto](docs/NotificationInfoDtoPagedResultDto.md)
 - [ObjectExtensionsDto](docs/ObjectExtensionsDto.md)
 - [ParameterApiDescriptionModel](docs/ParameterApiDescriptionModel.md)
 - [PermissionGrantInfoDto](docs/PermissionGrantInfoDto.md)
 - [PermissionGroupDto](docs/PermissionGroupDto.md)
 - [ProfileDto](docs/ProfileDto.md)
 - [PropertyApiDescriptionModel](docs/PropertyApiDescriptionModel.md)
 - [ProviderInfoDto](docs/ProviderInfoDto.md)
 - [RefreshDeviceStatusDto](docs/RefreshDeviceStatusDto.md)
 - [RegisterDto](docs/RegisterDto.md)
 - [RemoteServiceErrorInfo](docs/RemoteServiceErrorInfo.md)
 - [RemoteServiceErrorResponse](docs/RemoteServiceErrorResponse.md)
 - [RemoteServiceValidationErrorInfo](docs/RemoteServiceValidationErrorInfo.md)
 - [ResetPasswordDto](docs/ResetPasswordDto.md)
 - [ReturnValueApiDescriptionModel](docs/ReturnValueApiDescriptionModel.md)
 - [SendPasswordResetCodeDto](docs/SendPasswordResetCodeDto.md)
 - [SendSmsCodeDto](docs/SendSmsCodeDto.md)
 - [SettingsDto](docs/SettingsDto.md)
 - [SimpleDataDto](docs/SimpleDataDto.md)
 - [SimpleDataDtoPagedResultDto](docs/SimpleDataDtoPagedResultDto.md)
 - [SpecialItemDto](docs/SpecialItemDto.md)
 - [StringKeyValue](docs/StringKeyValue.md)
 - [StringSetKeyValueDto](docs/StringSetKeyValueDto.md)
 - [SyncItemChangedEto](docs/SyncItemChangedEto.md)
 - [SyncStateDto](docs/SyncStateDto.md)
 - [TenantCreateDto](docs/TenantCreateDto.md)
 - [TenantDto](docs/TenantDto.md)
 - [TenantDtoPagedResultDto](docs/TenantDtoPagedResultDto.md)
 - [TenantUpdateDto](docs/TenantUpdateDto.md)
 - [TimeZone](docs/TimeZone.md)
 - [TimingDto](docs/TimingDto.md)
 - [TodoOrderBy](docs/TodoOrderBy.md)
 - [TodoSettingsDto](docs/TodoSettingsDto.md)
 - [TypeApiDescriptionModel](docs/TypeApiDescriptionModel.md)
 - [UpdateEmailSettingsDto](docs/UpdateEmailSettingsDto.md)
 - [UpdateFeatureDto](docs/UpdateFeatureDto.md)
 - [UpdateFeaturesDto](docs/UpdateFeaturesDto.md)
 - [UpdatePermissionDto](docs/UpdatePermissionDto.md)
 - [UpdatePermissionsDto](docs/UpdatePermissionsDto.md)
 - [UpdateProfileDto](docs/UpdateProfileDto.md)
 - [UploadCredentials](docs/UploadCredentials.md)
 - [UserData](docs/UserData.md)
 - [UserDataListResultDto](docs/UserDataListResultDto.md)
 - [UserLoginInfo](docs/UserLoginInfo.md)
 - [UserStorageDto](docs/UserStorageDto.md)
 - [UserStorageItemDto](docs/UserStorageItemDto.md)
 - [WindowsTimeZone](docs/WindowsTimeZone.md)


## Documentation For Authorization



### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://localhost:44342/connect/authorize
- **Scopes**: 
 - **Doggy**: Doggy API

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



