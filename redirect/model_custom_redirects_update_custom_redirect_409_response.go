/*
BloxOne Redirect API

You can configure BloxOne Threat Defense Cloud to redirect traffic to the Infoblox server that displays the default or customized redirect page. You can redirect traffic to a custom destination using custom redirects.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redirect

import (
	"encoding/json"
)

// checks if the CustomRedirectsUpdateCustomRedirect409Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomRedirectsUpdateCustomRedirect409Response{}

// CustomRedirectsUpdateCustomRedirect409Response struct for CustomRedirectsUpdateCustomRedirect409Response
type CustomRedirectsUpdateCustomRedirect409Response struct {
	Error                *CustomRedirectsUpdateCustomRedirect409ResponseError `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomRedirectsUpdateCustomRedirect409Response CustomRedirectsUpdateCustomRedirect409Response

// NewCustomRedirectsUpdateCustomRedirect409Response instantiates a new CustomRedirectsUpdateCustomRedirect409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomRedirectsUpdateCustomRedirect409Response() *CustomRedirectsUpdateCustomRedirect409Response {
	this := CustomRedirectsUpdateCustomRedirect409Response{}
	return &this
}

// NewCustomRedirectsUpdateCustomRedirect409ResponseWithDefaults instantiates a new CustomRedirectsUpdateCustomRedirect409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomRedirectsUpdateCustomRedirect409ResponseWithDefaults() *CustomRedirectsUpdateCustomRedirect409Response {
	this := CustomRedirectsUpdateCustomRedirect409Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CustomRedirectsUpdateCustomRedirect409Response) GetError() CustomRedirectsUpdateCustomRedirect409ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret CustomRedirectsUpdateCustomRedirect409ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRedirectsUpdateCustomRedirect409Response) GetErrorOk() (*CustomRedirectsUpdateCustomRedirect409ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CustomRedirectsUpdateCustomRedirect409Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CustomRedirectsUpdateCustomRedirect409ResponseError and assigns it to the Error field.
func (o *CustomRedirectsUpdateCustomRedirect409Response) SetError(v CustomRedirectsUpdateCustomRedirect409ResponseError) {
	o.Error = &v
}

func (o CustomRedirectsUpdateCustomRedirect409Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomRedirectsUpdateCustomRedirect409Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomRedirectsUpdateCustomRedirect409Response) UnmarshalJSON(data []byte) (err error) {
	varCustomRedirectsUpdateCustomRedirect409Response := _CustomRedirectsUpdateCustomRedirect409Response{}

	err = json.Unmarshal(data, &varCustomRedirectsUpdateCustomRedirect409Response)

	if err != nil {
		return err
	}

	*o = CustomRedirectsUpdateCustomRedirect409Response(varCustomRedirectsUpdateCustomRedirect409Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomRedirectsUpdateCustomRedirect409Response struct {
	value *CustomRedirectsUpdateCustomRedirect409Response
	isSet bool
}

func (v NullableCustomRedirectsUpdateCustomRedirect409Response) Get() *CustomRedirectsUpdateCustomRedirect409Response {
	return v.value
}

func (v *NullableCustomRedirectsUpdateCustomRedirect409Response) Set(val *CustomRedirectsUpdateCustomRedirect409Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomRedirectsUpdateCustomRedirect409Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomRedirectsUpdateCustomRedirect409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomRedirectsUpdateCustomRedirect409Response(val *CustomRedirectsUpdateCustomRedirect409Response) *NullableCustomRedirectsUpdateCustomRedirect409Response {
	return &NullableCustomRedirectsUpdateCustomRedirect409Response{value: val, isSet: true}
}

func (v NullableCustomRedirectsUpdateCustomRedirect409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomRedirectsUpdateCustomRedirect409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
