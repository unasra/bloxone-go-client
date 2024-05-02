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

// checks if the ListViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListViewResponse{}

// ListViewResponse The View object list response format.
type ListViewResponse struct {
	// List of View objects.
	Results              []View `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListViewResponse ListViewResponse

// NewListViewResponse instantiates a new ListViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListViewResponse() *ListViewResponse {
	this := ListViewResponse{}
	return &this
}

// NewListViewResponseWithDefaults instantiates a new ListViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListViewResponseWithDefaults() *ListViewResponse {
	this := ListViewResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListViewResponse) GetResults() []View {
	if o == nil || IsNil(o.Results) {
		var ret []View
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListViewResponse) GetResultsOk() ([]View, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListViewResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []View and assigns it to the Results field.
func (o *ListViewResponse) SetResults(v []View) {
	o.Results = v
}

func (o ListViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListViewResponse) UnmarshalJSON(data []byte) (err error) {
	varListViewResponse := _ListViewResponse{}

	err = json.Unmarshal(data, &varListViewResponse)

	if err != nil {
		return err
	}

	*o = ListViewResponse(varListViewResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListViewResponse struct {
	value *ListViewResponse
	isSet bool
}

func (v NullableListViewResponse) Get() *ListViewResponse {
	return v.value
}

func (v *NullableListViewResponse) Set(val *ListViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListViewResponse(val *ListViewResponse) *NullableListViewResponse {
	return &NullableListViewResponse{value: val, isSet: true}
}

func (v NullableListViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}