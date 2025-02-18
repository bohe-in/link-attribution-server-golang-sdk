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

// checks if the ResponsesSdkLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesSdkLink{}

// ResponsesSdkLink struct for ResponsesSdkLink
type ResponsesSdkLink struct {
	AnalyticsTags ResponsesLinkAnalyticTags `json:"analyticsTags"`
	Data map[string]interface{} `json:"data"`
	Slug string `json:"slug"`
	SocialMediaTags ResponsesLinkSocialMediaTags `json:"socialMediaTags"`
}

// NewResponsesSdkLink instantiates a new ResponsesSdkLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesSdkLink(analyticsTags ResponsesLinkAnalyticTags, data map[string]interface{}, slug string, socialMediaTags ResponsesLinkSocialMediaTags) *ResponsesSdkLink {
	this := ResponsesSdkLink{}
	this.AnalyticsTags = analyticsTags
	this.Data = data
	this.Slug = slug
	this.SocialMediaTags = socialMediaTags
	return &this
}

// NewResponsesSdkLinkWithDefaults instantiates a new ResponsesSdkLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesSdkLinkWithDefaults() *ResponsesSdkLink {
	this := ResponsesSdkLink{}
	return &this
}

// GetAnalyticsTags returns the AnalyticsTags field value
func (o *ResponsesSdkLink) GetAnalyticsTags() ResponsesLinkAnalyticTags {
	if o == nil {
		var ret ResponsesLinkAnalyticTags
		return ret
	}

	return o.AnalyticsTags
}

// GetAnalyticsTagsOk returns a tuple with the AnalyticsTags field value
// and a boolean to check if the value has been set.
func (o *ResponsesSdkLink) GetAnalyticsTagsOk() (*ResponsesLinkAnalyticTags, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyticsTags, true
}

// SetAnalyticsTags sets field value
func (o *ResponsesSdkLink) SetAnalyticsTags(v ResponsesLinkAnalyticTags) {
	o.AnalyticsTags = v
}

// GetData returns the Data field value
func (o *ResponsesSdkLink) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponsesSdkLink) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ResponsesSdkLink) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetSlug returns the Slug field value
func (o *ResponsesSdkLink) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *ResponsesSdkLink) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *ResponsesSdkLink) SetSlug(v string) {
	o.Slug = v
}

// GetSocialMediaTags returns the SocialMediaTags field value
func (o *ResponsesSdkLink) GetSocialMediaTags() ResponsesLinkSocialMediaTags {
	if o == nil {
		var ret ResponsesLinkSocialMediaTags
		return ret
	}

	return o.SocialMediaTags
}

// GetSocialMediaTagsOk returns a tuple with the SocialMediaTags field value
// and a boolean to check if the value has been set.
func (o *ResponsesSdkLink) GetSocialMediaTagsOk() (*ResponsesLinkSocialMediaTags, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SocialMediaTags, true
}

// SetSocialMediaTags sets field value
func (o *ResponsesSdkLink) SetSocialMediaTags(v ResponsesLinkSocialMediaTags) {
	o.SocialMediaTags = v
}

func (o ResponsesSdkLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesSdkLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["analyticsTags"] = o.AnalyticsTags
	toSerialize["data"] = o.Data
	toSerialize["slug"] = o.Slug
	toSerialize["socialMediaTags"] = o.SocialMediaTags
	return toSerialize, nil
}

type NullableResponsesSdkLink struct {
	value *ResponsesSdkLink
	isSet bool
}

func (v NullableResponsesSdkLink) Get() *ResponsesSdkLink {
	return v.value
}

func (v *NullableResponsesSdkLink) Set(val *ResponsesSdkLink) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesSdkLink) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesSdkLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesSdkLink(val *ResponsesSdkLink) *NullableResponsesSdkLink {
	return &NullableResponsesSdkLink{value: val, isSet: true}
}

func (v NullableResponsesSdkLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesSdkLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


