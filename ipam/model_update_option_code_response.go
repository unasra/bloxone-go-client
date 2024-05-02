/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the UpdateOptionCodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOptionCodeResponse{}

// UpdateOptionCodeResponse The response format to update the __OptionCode__ object.
type UpdateOptionCodeResponse struct {
	Result               *OptionCode `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateOptionCodeResponse UpdateOptionCodeResponse

// NewUpdateOptionCodeResponse instantiates a new UpdateOptionCodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOptionCodeResponse() *UpdateOptionCodeResponse {
	this := UpdateOptionCodeResponse{}
	return &this
}

// NewUpdateOptionCodeResponseWithDefaults instantiates a new UpdateOptionCodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOptionCodeResponseWithDefaults() *UpdateOptionCodeResponse {
	this := UpdateOptionCodeResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateOptionCodeResponse) GetResult() OptionCode {
	if o == nil || IsNil(o.Result) {
		var ret OptionCode
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOptionCodeResponse) GetResultOk() (*OptionCode, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateOptionCodeResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OptionCode and assigns it to the Result field.
func (o *UpdateOptionCodeResponse) SetResult(v OptionCode) {
	o.Result = &v
}

func (o UpdateOptionCodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOptionCodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateOptionCodeResponse) UnmarshalJSON(data []byte) (err error) {
	varUpdateOptionCodeResponse := _UpdateOptionCodeResponse{}

	err = json.Unmarshal(data, &varUpdateOptionCodeResponse)

	if err != nil {
		return err
	}

	*o = UpdateOptionCodeResponse(varUpdateOptionCodeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateOptionCodeResponse struct {
	value *UpdateOptionCodeResponse
	isSet bool
}

func (v NullableUpdateOptionCodeResponse) Get() *UpdateOptionCodeResponse {
	return v.value
}

func (v *NullableUpdateOptionCodeResponse) Set(val *UpdateOptionCodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOptionCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOptionCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOptionCodeResponse(val *UpdateOptionCodeResponse) *NullableUpdateOptionCodeResponse {
	return &NullableUpdateOptionCodeResponse{value: val, isSet: true}
}

func (v NullableUpdateOptionCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOptionCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}