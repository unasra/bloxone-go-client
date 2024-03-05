/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
)

// checks if the InternalDomainListsUpdateInternalDomains404Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalDomainListsUpdateInternalDomains404Response{}

// InternalDomainListsUpdateInternalDomains404Response struct for InternalDomainListsUpdateInternalDomains404Response
type InternalDomainListsUpdateInternalDomains404Response struct {
	Error *InternalDomainListsUpdateInternalDomains404ResponseError `json:"error,omitempty"`
}

// NewInternalDomainListsUpdateInternalDomains404Response instantiates a new InternalDomainListsUpdateInternalDomains404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalDomainListsUpdateInternalDomains404Response() *InternalDomainListsUpdateInternalDomains404Response {
	this := InternalDomainListsUpdateInternalDomains404Response{}
	return &this
}

// NewInternalDomainListsUpdateInternalDomains404ResponseWithDefaults instantiates a new InternalDomainListsUpdateInternalDomains404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalDomainListsUpdateInternalDomains404ResponseWithDefaults() *InternalDomainListsUpdateInternalDomains404Response {
	this := InternalDomainListsUpdateInternalDomains404Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InternalDomainListsUpdateInternalDomains404Response) GetError() InternalDomainListsUpdateInternalDomains404ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret InternalDomainListsUpdateInternalDomains404ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalDomainListsUpdateInternalDomains404Response) GetErrorOk() (*InternalDomainListsUpdateInternalDomains404ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InternalDomainListsUpdateInternalDomains404Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given InternalDomainListsUpdateInternalDomains404ResponseError and assigns it to the Error field.
func (o *InternalDomainListsUpdateInternalDomains404Response) SetError(v InternalDomainListsUpdateInternalDomains404ResponseError) {
	o.Error = &v
}

func (o InternalDomainListsUpdateInternalDomains404Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalDomainListsUpdateInternalDomains404Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableInternalDomainListsUpdateInternalDomains404Response struct {
	value *InternalDomainListsUpdateInternalDomains404Response
	isSet bool
}

func (v NullableInternalDomainListsUpdateInternalDomains404Response) Get() *InternalDomainListsUpdateInternalDomains404Response {
	return v.value
}

func (v *NullableInternalDomainListsUpdateInternalDomains404Response) Set(val *InternalDomainListsUpdateInternalDomains404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalDomainListsUpdateInternalDomains404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalDomainListsUpdateInternalDomains404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalDomainListsUpdateInternalDomains404Response(val *InternalDomainListsUpdateInternalDomains404Response) *NullableInternalDomainListsUpdateInternalDomains404Response {
	return &NullableInternalDomainListsUpdateInternalDomains404Response{value: val, isSet: true}
}

func (v NullableInternalDomainListsUpdateInternalDomains404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalDomainListsUpdateInternalDomains404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
