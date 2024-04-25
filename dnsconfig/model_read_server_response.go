/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
)

// checks if the ReadServerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadServerResponse{}

// ReadServerResponse The Server object read response format.
type ReadServerResponse struct {
	Result *Server `json:"result,omitempty"`
}

// NewReadServerResponse instantiates a new ReadServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadServerResponse() *ReadServerResponse {
	this := ReadServerResponse{}
	return &this
}

// NewReadServerResponseWithDefaults instantiates a new ReadServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadServerResponseWithDefaults() *ReadServerResponse {
	this := ReadServerResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadServerResponse) GetResult() Server {
	if o == nil || IsNil(o.Result) {
		var ret Server
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadServerResponse) GetResultOk() (*Server, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadServerResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Server and assigns it to the Result field.
func (o *ReadServerResponse) SetResult(v Server) {
	o.Result = &v
}

func (o ReadServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadServerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableReadServerResponse struct {
	value *ReadServerResponse
	isSet bool
}

func (v NullableReadServerResponse) Get() *ReadServerResponse {
	return v.value
}

func (v *NullableReadServerResponse) Set(val *ReadServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadServerResponse(val *ReadServerResponse) *NullableReadServerResponse {
	return &NullableReadServerResponse{value: val, isSet: true}
}

func (v NullableReadServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
