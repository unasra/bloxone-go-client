/*
Discovery Configuration API V2

The Discovery configuration service is a BloxOne Service that provides configuration for accessing and syncing the Cloud assets   Base Paths:  1. provider: **_/api/cloud_discovery/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clouddiscovery

import (
	"encoding/json"
)

// checks if the DNSConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNSConfig{}

// DNSConfig struct for DNSConfig
type DNSConfig struct {
	ConsolidatedZoneDataEnabled *bool `json:"consolidated_zone_data_enabled,omitempty"`
	// split_view_enabled consolidates private zones into a single view, which is separate from the public zone view.
	SplitViewEnabled     *bool   `json:"split_view_enabled,omitempty"`
	SyncType             *string `json:"sync_type,omitempty"`
	ViewId               *string `json:"view_id,omitempty"`
	ViewName             *string `json:"view_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DNSConfig DNSConfig

// NewDNSConfig instantiates a new DNSConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSConfig() *DNSConfig {
	this := DNSConfig{}
	return &this
}

// NewDNSConfigWithDefaults instantiates a new DNSConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSConfigWithDefaults() *DNSConfig {
	this := DNSConfig{}
	return &this
}

// GetConsolidatedZoneDataEnabled returns the ConsolidatedZoneDataEnabled field value if set, zero value otherwise.
func (o *DNSConfig) GetConsolidatedZoneDataEnabled() bool {
	if o == nil || IsNil(o.ConsolidatedZoneDataEnabled) {
		var ret bool
		return ret
	}
	return *o.ConsolidatedZoneDataEnabled
}

// GetConsolidatedZoneDataEnabledOk returns a tuple with the ConsolidatedZoneDataEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetConsolidatedZoneDataEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsolidatedZoneDataEnabled) {
		return nil, false
	}
	return o.ConsolidatedZoneDataEnabled, true
}

// HasConsolidatedZoneDataEnabled returns a boolean if a field has been set.
func (o *DNSConfig) HasConsolidatedZoneDataEnabled() bool {
	if o != nil && !IsNil(o.ConsolidatedZoneDataEnabled) {
		return true
	}

	return false
}

// SetConsolidatedZoneDataEnabled gets a reference to the given bool and assigns it to the ConsolidatedZoneDataEnabled field.
func (o *DNSConfig) SetConsolidatedZoneDataEnabled(v bool) {
	o.ConsolidatedZoneDataEnabled = &v
}

// GetSplitViewEnabled returns the SplitViewEnabled field value if set, zero value otherwise.
func (o *DNSConfig) GetSplitViewEnabled() bool {
	if o == nil || IsNil(o.SplitViewEnabled) {
		var ret bool
		return ret
	}
	return *o.SplitViewEnabled
}

// GetSplitViewEnabledOk returns a tuple with the SplitViewEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetSplitViewEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SplitViewEnabled) {
		return nil, false
	}
	return o.SplitViewEnabled, true
}

// HasSplitViewEnabled returns a boolean if a field has been set.
func (o *DNSConfig) HasSplitViewEnabled() bool {
	if o != nil && !IsNil(o.SplitViewEnabled) {
		return true
	}

	return false
}

// SetSplitViewEnabled gets a reference to the given bool and assigns it to the SplitViewEnabled field.
func (o *DNSConfig) SetSplitViewEnabled(v bool) {
	o.SplitViewEnabled = &v
}

// GetSyncType returns the SyncType field value if set, zero value otherwise.
func (o *DNSConfig) GetSyncType() string {
	if o == nil || IsNil(o.SyncType) {
		var ret string
		return ret
	}
	return *o.SyncType
}

// GetSyncTypeOk returns a tuple with the SyncType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetSyncTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SyncType) {
		return nil, false
	}
	return o.SyncType, true
}

// HasSyncType returns a boolean if a field has been set.
func (o *DNSConfig) HasSyncType() bool {
	if o != nil && !IsNil(o.SyncType) {
		return true
	}

	return false
}

// SetSyncType gets a reference to the given string and assigns it to the SyncType field.
func (o *DNSConfig) SetSyncType(v string) {
	o.SyncType = &v
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *DNSConfig) GetViewId() string {
	if o == nil || IsNil(o.ViewId) {
		var ret string
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.ViewId) {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *DNSConfig) HasViewId() bool {
	if o != nil && !IsNil(o.ViewId) {
		return true
	}

	return false
}

// SetViewId gets a reference to the given string and assigns it to the ViewId field.
func (o *DNSConfig) SetViewId(v string) {
	o.ViewId = &v
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *DNSConfig) GetViewName() string {
	if o == nil || IsNil(o.ViewName) {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetViewNameOk() (*string, bool) {
	if o == nil || IsNil(o.ViewName) {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *DNSConfig) HasViewName() bool {
	if o != nil && !IsNil(o.ViewName) {
		return true
	}

	return false
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *DNSConfig) SetViewName(v string) {
	o.ViewName = &v
}

func (o DNSConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsolidatedZoneDataEnabled) {
		toSerialize["consolidated_zone_data_enabled"] = o.ConsolidatedZoneDataEnabled
	}
	if !IsNil(o.SplitViewEnabled) {
		toSerialize["split_view_enabled"] = o.SplitViewEnabled
	}
	if !IsNil(o.SyncType) {
		toSerialize["sync_type"] = o.SyncType
	}
	if !IsNil(o.ViewId) {
		toSerialize["view_id"] = o.ViewId
	}
	if !IsNil(o.ViewName) {
		toSerialize["view_name"] = o.ViewName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DNSConfig) UnmarshalJSON(data []byte) (err error) {
	varDNSConfig := _DNSConfig{}

	err = json.Unmarshal(data, &varDNSConfig)

	if err != nil {
		return err
	}

	*o = DNSConfig(varDNSConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "consolidated_zone_data_enabled")
		delete(additionalProperties, "split_view_enabled")
		delete(additionalProperties, "sync_type")
		delete(additionalProperties, "view_id")
		delete(additionalProperties, "view_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDNSConfig struct {
	value *DNSConfig
	isSet bool
}

func (v NullableDNSConfig) Get() *DNSConfig {
	return v.value
}

func (v *NullableDNSConfig) Set(val *DNSConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSConfig(val *DNSConfig) *NullableDNSConfig {
	return &NullableDNSConfig{value: val, isSet: true}
}

func (v NullableDNSConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
