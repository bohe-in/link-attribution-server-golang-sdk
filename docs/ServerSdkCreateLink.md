# ServerSdkCreateLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsTags** | Pointer to [**RequestsLinkAnalyticTags**](RequestsLinkAnalyticTags.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Redirects** | Pointer to [**RequestsLinkRedirects**](RequestsLinkRedirects.md) |  | [optional] 
**SocialMediaTags** | Pointer to [**RequestsLinkSocialMediaTags**](RequestsLinkSocialMediaTags.md) |  | [optional] 

## Methods

### NewServerSdkCreateLink

`func NewServerSdkCreateLink(name string, ) *ServerSdkCreateLink`

NewServerSdkCreateLink instantiates a new ServerSdkCreateLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerSdkCreateLinkWithDefaults

`func NewServerSdkCreateLinkWithDefaults() *ServerSdkCreateLink`

NewServerSdkCreateLinkWithDefaults instantiates a new ServerSdkCreateLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsTags

`func (o *ServerSdkCreateLink) GetAnalyticsTags() RequestsLinkAnalyticTags`

GetAnalyticsTags returns the AnalyticsTags field if non-nil, zero value otherwise.

### GetAnalyticsTagsOk

`func (o *ServerSdkCreateLink) GetAnalyticsTagsOk() (*RequestsLinkAnalyticTags, bool)`

GetAnalyticsTagsOk returns a tuple with the AnalyticsTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsTags

`func (o *ServerSdkCreateLink) SetAnalyticsTags(v RequestsLinkAnalyticTags)`

SetAnalyticsTags sets AnalyticsTags field to given value.

### HasAnalyticsTags

`func (o *ServerSdkCreateLink) HasAnalyticsTags() bool`

HasAnalyticsTags returns a boolean if a field has been set.

### GetData

`func (o *ServerSdkCreateLink) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServerSdkCreateLink) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServerSdkCreateLink) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ServerSdkCreateLink) HasData() bool`

HasData returns a boolean if a field has been set.

### GetName

`func (o *ServerSdkCreateLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerSdkCreateLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerSdkCreateLink) SetName(v string)`

SetName sets Name field to given value.


### GetRedirects

`func (o *ServerSdkCreateLink) GetRedirects() RequestsLinkRedirects`

GetRedirects returns the Redirects field if non-nil, zero value otherwise.

### GetRedirectsOk

`func (o *ServerSdkCreateLink) GetRedirectsOk() (*RequestsLinkRedirects, bool)`

GetRedirectsOk returns a tuple with the Redirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirects

`func (o *ServerSdkCreateLink) SetRedirects(v RequestsLinkRedirects)`

SetRedirects sets Redirects field to given value.

### HasRedirects

`func (o *ServerSdkCreateLink) HasRedirects() bool`

HasRedirects returns a boolean if a field has been set.

### GetSocialMediaTags

`func (o *ServerSdkCreateLink) GetSocialMediaTags() RequestsLinkSocialMediaTags`

GetSocialMediaTags returns the SocialMediaTags field if non-nil, zero value otherwise.

### GetSocialMediaTagsOk

`func (o *ServerSdkCreateLink) GetSocialMediaTagsOk() (*RequestsLinkSocialMediaTags, bool)`

GetSocialMediaTagsOk returns a tuple with the SocialMediaTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialMediaTags

`func (o *ServerSdkCreateLink) SetSocialMediaTags(v RequestsLinkSocialMediaTags)`

SetSocialMediaTags sets SocialMediaTags field to given value.

### HasSocialMediaTags

`func (o *ServerSdkCreateLink) HasSocialMediaTags() bool`

HasSocialMediaTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


