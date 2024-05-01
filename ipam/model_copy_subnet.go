/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
)

// checks if the CopySubnet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopySubnet{}

// CopySubnet struct for CopySubnet
type CopySubnet struct {
	// The description for the copied subnet. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Indicates whether dhcp options should be copied or not when _ipam/subnet_ object is copied.  Defaults to _false_.
	CopyDhcpOptions *bool `json:"copy_dhcp_options,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The name for the copied subnet. May contain 1 to 256 characters. Can include UTF-8.
	Name *string `json:"name,omitempty"`
	// Indicates whether child objects should be copied or not.  Defaults to _false_.
	Recursive *bool `json:"recursive,omitempty"`
	// Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_.
	SkipOnError *bool `json:"skip_on_error,omitempty"`
	// The resource identifier.
	Space                string `json:"space"`
	AdditionalProperties map[string]interface{}
}

type _CopySubnet CopySubnet

// NewCopySubnet instantiates a new CopySubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopySubnet(space string) *CopySubnet {
	this := CopySubnet{}
	this.Space = space
	return &this
}

// NewCopySubnetWithDefaults instantiates a new CopySubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopySubnetWithDefaults() *CopySubnet {
	this := CopySubnet{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CopySubnet) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopySubnet) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CopySubnet) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CopySubnet) SetComment(v string) {
	o.Comment = &v
}

// GetCopyDhcpOptions returns the CopyDhcpOptions field value if set, zero value otherwise.
func (o *CopySubnet) GetCopyDhcpOptions() bool {
	if o == nil || IsNil(o.CopyDhcpOptions) {
		var ret bool
		return ret
	}
	return *o.CopyDhcpOptions
}

// GetCopyDhcpOptionsOk returns a tuple with the CopyDhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopySubnet) GetCopyDhcpOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyDhcpOptions) {
		return nil, false
	}
	return o.CopyDhcpOptions, true
}

// HasCopyDhcpOptions returns a boolean if a field has been set.
func (o *CopySubnet) HasCopyDhcpOptions() bool {
	if o != nil && !IsNil(o.CopyDhcpOptions) {
		return true
	}

	return false
}

// SetCopyDhcpOptions gets a reference to the given bool and assigns it to the CopyDhcpOptions field.
func (o *CopySubnet) SetCopyDhcpOptions(v bool) {
	o.CopyDhcpOptions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CopySubnet) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopySubnet) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CopySubnet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CopySubnet) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CopySubnet) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopySubnet) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CopySubnet) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CopySubnet) SetName(v string) {
	o.Name = &v
}

// GetRecursive returns the Recursive field value if set, zero value otherwise.
func (o *CopySubnet) GetRecursive() bool {
	if o == nil || IsNil(o.Recursive) {
		var ret bool
		return ret
	}
	return *o.Recursive
}

// GetRecursiveOk returns a tuple with the Recursive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopySubnet) GetRecursiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Recursive) {
		return nil, false
	}
	return o.Recursive, true
}

// HasRecursive returns a boolean if a field has been set.
func (o *CopySubnet) HasRecursive() bool {
	if o != nil && !IsNil(o.Recursive) {
		return true
	}

	return false
}

// SetRecursive gets a reference to the given bool and assigns it to the Recursive field.
func (o *CopySubnet) SetRecursive(v bool) {
	o.Recursive = &v
}

// GetSkipOnError returns the SkipOnError field value if set, zero value otherwise.
func (o *CopySubnet) GetSkipOnError() bool {
	if o == nil || IsNil(o.SkipOnError) {
		var ret bool
		return ret
	}
	return *o.SkipOnError
}

// GetSkipOnErrorOk returns a tuple with the SkipOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopySubnet) GetSkipOnErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipOnError) {
		return nil, false
	}
	return o.SkipOnError, true
}

// HasSkipOnError returns a boolean if a field has been set.
func (o *CopySubnet) HasSkipOnError() bool {
	if o != nil && !IsNil(o.SkipOnError) {
		return true
	}

	return false
}

// SetSkipOnError gets a reference to the given bool and assigns it to the SkipOnError field.
func (o *CopySubnet) SetSkipOnError(v bool) {
	o.SkipOnError = &v
}

// GetSpace returns the Space field value
func (o *CopySubnet) GetSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *CopySubnet) GetSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *CopySubnet) SetSpace(v string) {
	o.Space = v
}

func (o CopySubnet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopySubnet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CopyDhcpOptions) {
		toSerialize["copy_dhcp_options"] = o.CopyDhcpOptions
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Recursive) {
		toSerialize["recursive"] = o.Recursive
	}
	if !IsNil(o.SkipOnError) {
		toSerialize["skip_on_error"] = o.SkipOnError
	}
	toSerialize["space"] = o.Space

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CopySubnet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"space",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCopySubnet := _CopySubnet{}

	err = json.Unmarshal(data, &varCopySubnet)

	if err != nil {
		return err
	}

	*o = CopySubnet(varCopySubnet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "copy_dhcp_options")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "recursive")
		delete(additionalProperties, "skip_on_error")
		delete(additionalProperties, "space")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCopySubnet struct {
	value *CopySubnet
	isSet bool
}

func (v NullableCopySubnet) Get() *CopySubnet {
	return v.value
}

func (v *NullableCopySubnet) Set(val *CopySubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableCopySubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableCopySubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopySubnet(val *CopySubnet) *NullableCopySubnet {
	return &NullableCopySubnet{value: val, isSet: true}
}

func (v NullableCopySubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopySubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
