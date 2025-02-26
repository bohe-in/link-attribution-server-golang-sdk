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

// checks if the ResponsesSdkLinkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesSdkLinkResponse{}

// ResponsesSdkLinkResponse struct for ResponsesSdkLinkResponse
type ResponsesSdkLinkResponse struct {
	ApiVersion float64 `json:"api_version"`
	Data ResponsesSdkLinkResponseData `json:"data"`
}

// NewResponsesSdkLinkResponse instantiates a new ResponsesSdkLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesSdkLinkResponse(apiVersion float64, data ResponsesSdkLinkResponseData) *ResponsesSdkLinkResponse {
	this := ResponsesSdkLinkResponse{}
	this.ApiVersion = apiVersion
	this.Data = data
	return &this
}

// NewResponsesSdkLinkResponseWithDefaults instantiates a new ResponsesSdkLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesSdkLinkResponseWithDefaults() *ResponsesSdkLinkResponse {
	this := ResponsesSdkLinkResponse{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ResponsesSdkLinkResponse) GetApiVersion() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ResponsesSdkLinkResponse) GetApiVersionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ResponsesSdkLinkResponse) SetApiVersion(v float64) {
	o.ApiVersion = v
}

// GetData returns the Data field value
func (o *ResponsesSdkLinkResponse) GetData() ResponsesSdkLinkResponseData {
	if o == nil {
		var ret ResponsesSdkLinkResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponsesSdkLinkResponse) GetDataOk() (*ResponsesSdkLinkResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponsesSdkLinkResponse) SetData(v ResponsesSdkLinkResponseData) {
	o.Data = v
}

func (o ResponsesSdkLinkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesSdkLinkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["api_version"] = o.ApiVersion
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableResponsesSdkLinkResponse struct {
	value *ResponsesSdkLinkResponse
	isSet bool
}

func (v NullableResponsesSdkLinkResponse) Get() *ResponsesSdkLinkResponse {
	return v.value
}

func (v *NullableResponsesSdkLinkResponse) Set(val *ResponsesSdkLinkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesSdkLinkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesSdkLinkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesSdkLinkResponse(val *ResponsesSdkLinkResponse) *NullableResponsesSdkLinkResponse {
	return &NullableResponsesSdkLinkResponse{value: val, isSet: true}
}

func (v NullableResponsesSdkLinkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesSdkLinkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


