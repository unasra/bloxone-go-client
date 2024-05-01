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

// checks if the UpdateHostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateHostResponse{}

// UpdateHostResponse The DNS Host object update response format.
type UpdateHostResponse struct {
	Result               *Host `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateHostResponse UpdateHostResponse

// NewUpdateHostResponse instantiates a new UpdateHostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateHostResponse() *UpdateHostResponse {
	this := UpdateHostResponse{}
	return &this
}

// NewUpdateHostResponseWithDefaults instantiates a new UpdateHostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateHostResponseWithDefaults() *UpdateHostResponse {
	this := UpdateHostResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateHostResponse) GetResult() Host {
	if o == nil || IsNil(o.Result) {
		var ret Host
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostResponse) GetResultOk() (*Host, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateHostResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Host and assigns it to the Result field.
func (o *UpdateHostResponse) SetResult(v Host) {
	o.Result = &v
}

func (o UpdateHostResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateHostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateHostResponse) UnmarshalJSON(data []byte) (err error) {
	varUpdateHostResponse := _UpdateHostResponse{}

	err = json.Unmarshal(data, &varUpdateHostResponse)

	if err != nil {
		return err
	}

	*o = UpdateHostResponse(varUpdateHostResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateHostResponse struct {
	value *UpdateHostResponse
	isSet bool
}

func (v NullableUpdateHostResponse) Get() *UpdateHostResponse {
	return v.value
}

func (v *NullableUpdateHostResponse) Set(val *UpdateHostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateHostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateHostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateHostResponse(val *UpdateHostResponse) *NullableUpdateHostResponse {
	return &NullableUpdateHostResponse{value: val, isSet: true}
}

func (v NullableUpdateHostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateHostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
