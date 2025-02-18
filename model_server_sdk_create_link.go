/*
Link Attribution Server SDK

This is the api used by link attribution

API version: 0.100000
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package linkattribution

import (
	"encoding/json"
)

// checks if the ServerSdkCreateLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerSdkCreateLink{}

// ServerSdkCreateLink struct for ServerSdkCreateLink
type ServerSdkCreateLink struct {
	AnalyticsTags *RequestsLinkAnalyticTags `json:"analyticsTags,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
	Name string `json:"name"`
	Redirects *RequestsLinkRedirects `json:"redirects,omitempty"`
	SocialMediaTags *RequestsLinkSocialMediaTags `json:"socialMediaTags,omitempty"`
}

// NewServerSdkCreateLink instantiates a new ServerSdkCreateLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerSdkCreateLink(name string) *ServerSdkCreateLink {
	this := ServerSdkCreateLink{}
	this.Name = name
	return &this
}

// NewServerSdkCreateLinkWithDefaults instantiates a new ServerSdkCreateLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerSdkCreateLinkWithDefaults() *ServerSdkCreateLink {
	this := ServerSdkCreateLink{}
	return &this
}

// GetAnalyticsTags returns the AnalyticsTags field value if set, zero value otherwise.
func (o *ServerSdkCreateLink) GetAnalyticsTags() RequestsLinkAnalyticTags {
	if o == nil || IsNil(o.AnalyticsTags) {
		var ret RequestsLinkAnalyticTags
		return ret
	}
	return *o.AnalyticsTags
}

// GetAnalyticsTagsOk returns a tuple with the AnalyticsTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerSdkCreateLink) GetAnalyticsTagsOk() (*RequestsLinkAnalyticTags, bool) {
	if o == nil || IsNil(o.AnalyticsTags) {
		return nil, false
	}
	return o.AnalyticsTags, true
}

// HasAnalyticsTags returns a boolean if a field has been set.
func (o *ServerSdkCreateLink) HasAnalyticsTags() bool {
	if o != nil && !IsNil(o.AnalyticsTags) {
		return true
	}

	return false
}

// SetAnalyticsTags gets a reference to the given RequestsLinkAnalyticTags and assigns it to the AnalyticsTags field.
func (o *ServerSdkCreateLink) SetAnalyticsTags(v RequestsLinkAnalyticTags) {
	o.AnalyticsTags = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServerSdkCreateLink) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerSdkCreateLink) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServerSdkCreateLink) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *ServerSdkCreateLink) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetName returns the Name field value
func (o *ServerSdkCreateLink) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServerSdkCreateLink) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServerSdkCreateLink) SetName(v string) {
	o.Name = v
}

// GetRedirects returns the Redirects field value if set, zero value otherwise.
func (o *ServerSdkCreateLink) GetRedirects() RequestsLinkRedirects {
	if o == nil || IsNil(o.Redirects) {
		var ret RequestsLinkRedirects
		return ret
	}
	return *o.Redirects
}

// GetRedirectsOk returns a tuple with the Redirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerSdkCreateLink) GetRedirectsOk() (*RequestsLinkRedirects, bool) {
	if o == nil || IsNil(o.Redirects) {
		return nil, false
	}
	return o.Redirects, true
}

// HasRedirects returns a boolean if a field has been set.
func (o *ServerSdkCreateLink) HasRedirects() bool {
	if o != nil && !IsNil(o.Redirects) {
		return true
	}

	return false
}

// SetRedirects gets a reference to the given RequestsLinkRedirects and assigns it to the Redirects field.
func (o *ServerSdkCreateLink) SetRedirects(v RequestsLinkRedirects) {
	o.Redirects = &v
}

// GetSocialMediaTags returns the SocialMediaTags field value if set, zero value otherwise.
func (o *ServerSdkCreateLink) GetSocialMediaTags() RequestsLinkSocialMediaTags {
	if o == nil || IsNil(o.SocialMediaTags) {
		var ret RequestsLinkSocialMediaTags
		return ret
	}
	return *o.SocialMediaTags
}

// GetSocialMediaTagsOk returns a tuple with the SocialMediaTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerSdkCreateLink) GetSocialMediaTagsOk() (*RequestsLinkSocialMediaTags, bool) {
	if o == nil || IsNil(o.SocialMediaTags) {
		return nil, false
	}
	return o.SocialMediaTags, true
}

// HasSocialMediaTags returns a boolean if a field has been set.
func (o *ServerSdkCreateLink) HasSocialMediaTags() bool {
	if o != nil && !IsNil(o.SocialMediaTags) {
		return true
	}

	return false
}

// SetSocialMediaTags gets a reference to the given RequestsLinkSocialMediaTags and assigns it to the SocialMediaTags field.
func (o *ServerSdkCreateLink) SetSocialMediaTags(v RequestsLinkSocialMediaTags) {
	o.SocialMediaTags = &v
}

func (o ServerSdkCreateLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerSdkCreateLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnalyticsTags) {
		toSerialize["analyticsTags"] = o.AnalyticsTags
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Redirects) {
		toSerialize["redirects"] = o.Redirects
	}
	if !IsNil(o.SocialMediaTags) {
		toSerialize["socialMediaTags"] = o.SocialMediaTags
	}
	return toSerialize, nil
}

type NullableServerSdkCreateLink struct {
	value *ServerSdkCreateLink
	isSet bool
}

func (v NullableServerSdkCreateLink) Get() *ServerSdkCreateLink {
	return v.value
}

func (v *NullableServerSdkCreateLink) Set(val *ServerSdkCreateLink) {
	v.value = val
	v.isSet = true
}

func (v NullableServerSdkCreateLink) IsSet() bool {
	return v.isSet
}

func (v *NullableServerSdkCreateLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerSdkCreateLink(val *ServerSdkCreateLink) *NullableServerSdkCreateLink {
	return &NullableServerSdkCreateLink{value: val, isSet: true}
}

func (v NullableServerSdkCreateLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerSdkCreateLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


