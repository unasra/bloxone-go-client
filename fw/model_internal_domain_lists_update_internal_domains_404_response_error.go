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

// checks if the InternalDomainListsUpdateInternalDomains404ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalDomainListsUpdateInternalDomains404ResponseError{}

// InternalDomainListsUpdateInternalDomains404ResponseError struct for InternalDomainListsUpdateInternalDomains404ResponseError
type InternalDomainListsUpdateInternalDomains404ResponseError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *string `json:"status,omitempty"`
}

// NewInternalDomainListsUpdateInternalDomains404ResponseError instantiates a new InternalDomainListsUpdateInternalDomains404ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalDomainListsUpdateInternalDomains404ResponseError() *InternalDomainListsUpdateInternalDomains404ResponseError {
	this := InternalDomainListsUpdateInternalDomains404ResponseError{}
	return &this
}

// NewInternalDomainListsUpdateInternalDomains404ResponseErrorWithDefaults instantiates a new InternalDomainListsUpdateInternalDomains404ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalDomainListsUpdateInternalDomains404ResponseErrorWithDefaults() *InternalDomainListsUpdateInternalDomains404ResponseError {
	this := InternalDomainListsUpdateInternalDomains404ResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InternalDomainListsUpdateInternalDomains404ResponseError) SetStatus(v string) {
	o.Status = &v
}

func (o InternalDomainListsUpdateInternalDomains404ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalDomainListsUpdateInternalDomains404ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableInternalDomainListsUpdateInternalDomains404ResponseError struct {
	value *InternalDomainListsUpdateInternalDomains404ResponseError
	isSet bool
}

func (v NullableInternalDomainListsUpdateInternalDomains404ResponseError) Get() *InternalDomainListsUpdateInternalDomains404ResponseError {
	return v.value
}

func (v *NullableInternalDomainListsUpdateInternalDomains404ResponseError) Set(val *InternalDomainListsUpdateInternalDomains404ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalDomainListsUpdateInternalDomains404ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalDomainListsUpdateInternalDomains404ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalDomainListsUpdateInternalDomains404ResponseError(val *InternalDomainListsUpdateInternalDomains404ResponseError) *NullableInternalDomainListsUpdateInternalDomains404ResponseError {
	return &NullableInternalDomainListsUpdateInternalDomains404ResponseError{value: val, isSet: true}
}

func (v NullableInternalDomainListsUpdateInternalDomains404ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalDomainListsUpdateInternalDomains404ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
