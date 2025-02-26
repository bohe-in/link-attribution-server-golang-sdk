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

// checks if the ResponsesLinkAnalyticTags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesLinkAnalyticTags{}

// ResponsesLinkAnalyticTags struct for ResponsesLinkAnalyticTags
type ResponsesLinkAnalyticTags struct {
	Campaign NullableString `json:"campaign"`
	Channel NullableString `json:"channel"`
	Feature NullableString `json:"feature"`
	Tags NullableString `json:"tags"`
}

// NewResponsesLinkAnalyticTags instantiates a new ResponsesLinkAnalyticTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesLinkAnalyticTags(campaign NullableString, channel NullableString, feature NullableString, tags NullableString) *ResponsesLinkAnalyticTags {
	this := ResponsesLinkAnalyticTags{}
	this.Campaign = campaign
	this.Channel = channel
	this.Feature = feature
	this.Tags = tags
	return &this
}

// NewResponsesLinkAnalyticTagsWithDefaults instantiates a new ResponsesLinkAnalyticTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesLinkAnalyticTagsWithDefaults() *ResponsesLinkAnalyticTags {
	this := ResponsesLinkAnalyticTags{}
	return &this
}

// GetCampaign returns the Campaign field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResponsesLinkAnalyticTags) GetCampaign() string {
	if o == nil || o.Campaign.Get() == nil {
		var ret string
		return ret
	}

	return *o.Campaign.Get()
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponsesLinkAnalyticTags) GetCampaignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaign.Get(), o.Campaign.IsSet()
}

// SetCampaign sets field value
func (o *ResponsesLinkAnalyticTags) SetCampaign(v string) {
	o.Campaign.Set(&v)
}

// GetChannel returns the Channel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResponsesLinkAnalyticTags) GetChannel() string {
	if o == nil || o.Channel.Get() == nil {
		var ret string
		return ret
	}

	return *o.Channel.Get()
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponsesLinkAnalyticTags) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channel.Get(), o.Channel.IsSet()
}

// SetChannel sets field value
func (o *ResponsesLinkAnalyticTags) SetChannel(v string) {
	o.Channel.Set(&v)
}

// GetFeature returns the Feature field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResponsesLinkAnalyticTags) GetFeature() string {
	if o == nil || o.Feature.Get() == nil {
		var ret string
		return ret
	}

	return *o.Feature.Get()
}

// GetFeatureOk returns a tuple with the Feature field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponsesLinkAnalyticTags) GetFeatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Feature.Get(), o.Feature.IsSet()
}

// SetFeature sets field value
func (o *ResponsesLinkAnalyticTags) SetFeature(v string) {
	o.Feature.Set(&v)
}

// GetTags returns the Tags field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResponsesLinkAnalyticTags) GetTags() string {
	if o == nil || o.Tags.Get() == nil {
		var ret string
		return ret
	}

	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponsesLinkAnalyticTags) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// SetTags sets field value
func (o *ResponsesLinkAnalyticTags) SetTags(v string) {
	o.Tags.Set(&v)
}

func (o ResponsesLinkAnalyticTags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesLinkAnalyticTags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaign"] = o.Campaign.Get()
	toSerialize["channel"] = o.Channel.Get()
	toSerialize["feature"] = o.Feature.Get()
	toSerialize["tags"] = o.Tags.Get()
	return toSerialize, nil
}

type NullableResponsesLinkAnalyticTags struct {
	value *ResponsesLinkAnalyticTags
	isSet bool
}

func (v NullableResponsesLinkAnalyticTags) Get() *ResponsesLinkAnalyticTags {
	return v.value
}

func (v *NullableResponsesLinkAnalyticTags) Set(val *ResponsesLinkAnalyticTags) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesLinkAnalyticTags) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesLinkAnalyticTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesLinkAnalyticTags(val *ResponsesLinkAnalyticTags) *NullableResponsesLinkAnalyticTags {
	return &NullableResponsesLinkAnalyticTags{value: val, isSet: true}
}

func (v NullableResponsesLinkAnalyticTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesLinkAnalyticTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


