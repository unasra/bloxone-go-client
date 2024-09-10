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

// checks if the ListConfigProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConfigProfileResponse{}

// ListConfigProfileResponse The response format to retrieve config profiles.
type ListConfigProfileResponse struct {
	// Contains result-set depending on the type.
	Results              []Server `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListConfigProfileResponse ListConfigProfileResponse

// NewListConfigProfileResponse instantiates a new ListConfigProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfigProfileResponse() *ListConfigProfileResponse {
	this := ListConfigProfileResponse{}
	return &this
}

// NewListConfigProfileResponseWithDefaults instantiates a new ListConfigProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfigProfileResponseWithDefaults() *ListConfigProfileResponse {
	this := ListConfigProfileResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListConfigProfileResponse) GetResults() []Server {
	if o == nil || IsNil(o.Results) {
		var ret []Server
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfigProfileResponse) GetResultsOk() ([]Server, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListConfigProfileResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Server and assigns it to the Results field.
func (o *ListConfigProfileResponse) SetResults(v []Server) {
	o.Results = v
}

func (o ListConfigProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConfigProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListConfigProfileResponse) UnmarshalJSON(data []byte) (err error) {
	varListConfigProfileResponse := _ListConfigProfileResponse{}

	err = json.Unmarshal(data, &varListConfigProfileResponse)

	if err != nil {
		return err
	}

	*o = ListConfigProfileResponse(varListConfigProfileResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListConfigProfileResponse struct {
	value *ListConfigProfileResponse
	isSet bool
}

func (v NullableListConfigProfileResponse) Get() *ListConfigProfileResponse {
	return v.value
}

func (v *NullableListConfigProfileResponse) Set(val *ListConfigProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfigProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfigProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfigProfileResponse(val *ListConfigProfileResponse) *NullableListConfigProfileResponse {
	return &NullableListConfigProfileResponse{value: val, isSet: true}
}

func (v NullableListConfigProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfigProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}