/*
Configuration Api

Nice description goes here

API version: 1.15.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the WirelessNetworksItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WirelessNetworksItem{}

// WirelessNetworksItem struct for WirelessNetworksItem
type WirelessNetworksItem struct {
	Uid                  string         `json:"uid"`
	Alias                string         `json:"alias"`
	Ssid                 string         `json:"ssid"`
	Security             NullableString `json:"security"`
	IpVersion            string         `json:"ip_version"`
	DatetimeCreated      time.Time      `json:"datetime_created"`
	DatetimeUpdated      time.Time      `json:"datetime_updated"`
	Hidden               bool           `json:"hidden"`
	BandLocking          string         `json:"band_locking"`
	DnsLookupDomain      NullableString `json:"dns_lookup_domain"`
	DisableEdns          bool           `json:"disable_edns"`
	UseDns64             bool           `json:"use_dns64"`
	ExternalConnectivity bool           `json:"external_connectivity"`
}

type _WirelessNetworksItem WirelessNetworksItem

// NewWirelessNetworksItem instantiates a new WirelessNetworksItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelessNetworksItem(uid string, alias string, ssid string, security NullableString, ipVersion string, datetimeCreated time.Time, datetimeUpdated time.Time, hidden bool, bandLocking string, dnsLookupDomain NullableString, disableEdns bool, useDns64 bool, externalConnectivity bool) *WirelessNetworksItem {
	this := WirelessNetworksItem{}
	this.Uid = uid
	this.Alias = alias
	this.Ssid = ssid
	this.Security = security
	this.IpVersion = ipVersion
	this.DatetimeCreated = datetimeCreated
	this.DatetimeUpdated = datetimeUpdated
	this.Hidden = hidden
	this.BandLocking = bandLocking
	this.DnsLookupDomain = dnsLookupDomain
	this.DisableEdns = disableEdns
	this.UseDns64 = useDns64
	this.ExternalConnectivity = externalConnectivity
	return &this
}

// NewWirelessNetworksItemWithDefaults instantiates a new WirelessNetworksItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelessNetworksItemWithDefaults() *WirelessNetworksItem {
	this := WirelessNetworksItem{}
	return &this
}

// GetUid returns the Uid field value
func (o *WirelessNetworksItem) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *WirelessNetworksItem) SetUid(v string) {
	o.Uid = v
}

// GetAlias returns the Alias field value
func (o *WirelessNetworksItem) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *WirelessNetworksItem) SetAlias(v string) {
	o.Alias = v
}

// GetSsid returns the Ssid field value
func (o *WirelessNetworksItem) GetSsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *WirelessNetworksItem) SetSsid(v string) {
	o.Ssid = v
}

// GetSecurity returns the Security field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WirelessNetworksItem) GetSecurity() string {
	if o == nil || o.Security.Get() == nil {
		var ret string
		return ret
	}

	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessNetworksItem) GetSecurityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// SetSecurity sets field value
func (o *WirelessNetworksItem) SetSecurity(v string) {
	o.Security.Set(&v)
}

// GetIpVersion returns the IpVersion field value
func (o *WirelessNetworksItem) GetIpVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetIpVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpVersion, true
}

// SetIpVersion sets field value
func (o *WirelessNetworksItem) SetIpVersion(v string) {
	o.IpVersion = v
}

// GetDatetimeCreated returns the DatetimeCreated field value
func (o *WirelessNetworksItem) GetDatetimeCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DatetimeCreated
}

// GetDatetimeCreatedOk returns a tuple with the DatetimeCreated field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetDatetimeCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatetimeCreated, true
}

// SetDatetimeCreated sets field value
func (o *WirelessNetworksItem) SetDatetimeCreated(v time.Time) {
	o.DatetimeCreated = v
}

// GetDatetimeUpdated returns the DatetimeUpdated field value
func (o *WirelessNetworksItem) GetDatetimeUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DatetimeUpdated
}

// GetDatetimeUpdatedOk returns a tuple with the DatetimeUpdated field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetDatetimeUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatetimeUpdated, true
}

// SetDatetimeUpdated sets field value
func (o *WirelessNetworksItem) SetDatetimeUpdated(v time.Time) {
	o.DatetimeUpdated = v
}

// GetHidden returns the Hidden field value
func (o *WirelessNetworksItem) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *WirelessNetworksItem) SetHidden(v bool) {
	o.Hidden = v
}

// GetBandLocking returns the BandLocking field value
func (o *WirelessNetworksItem) GetBandLocking() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BandLocking
}

// GetBandLockingOk returns a tuple with the BandLocking field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetBandLockingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BandLocking, true
}

// SetBandLocking sets field value
func (o *WirelessNetworksItem) SetBandLocking(v string) {
	o.BandLocking = v
}

// GetDnsLookupDomain returns the DnsLookupDomain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WirelessNetworksItem) GetDnsLookupDomain() string {
	if o == nil || o.DnsLookupDomain.Get() == nil {
		var ret string
		return ret
	}

	return *o.DnsLookupDomain.Get()
}

// GetDnsLookupDomainOk returns a tuple with the DnsLookupDomain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessNetworksItem) GetDnsLookupDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsLookupDomain.Get(), o.DnsLookupDomain.IsSet()
}

// SetDnsLookupDomain sets field value
func (o *WirelessNetworksItem) SetDnsLookupDomain(v string) {
	o.DnsLookupDomain.Set(&v)
}

// GetDisableEdns returns the DisableEdns field value
func (o *WirelessNetworksItem) GetDisableEdns() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableEdns
}

// GetDisableEdnsOk returns a tuple with the DisableEdns field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetDisableEdnsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableEdns, true
}

// SetDisableEdns sets field value
func (o *WirelessNetworksItem) SetDisableEdns(v bool) {
	o.DisableEdns = v
}

// GetUseDns64 returns the UseDns64 field value
func (o *WirelessNetworksItem) GetUseDns64() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseDns64
}

// GetUseDns64Ok returns a tuple with the UseDns64 field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetUseDns64Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseDns64, true
}

// SetUseDns64 sets field value
func (o *WirelessNetworksItem) SetUseDns64(v bool) {
	o.UseDns64 = v
}

// GetExternalConnectivity returns the ExternalConnectivity field value
func (o *WirelessNetworksItem) GetExternalConnectivity() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExternalConnectivity
}

// GetExternalConnectivityOk returns a tuple with the ExternalConnectivity field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksItem) GetExternalConnectivityOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalConnectivity, true
}

// SetExternalConnectivity sets field value
func (o *WirelessNetworksItem) SetExternalConnectivity(v bool) {
	o.ExternalConnectivity = v
}

func (o WirelessNetworksItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WirelessNetworksItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uid"] = o.Uid
	toSerialize["alias"] = o.Alias
	toSerialize["ssid"] = o.Ssid
	toSerialize["security"] = o.Security.Get()
	toSerialize["ip_version"] = o.IpVersion
	toSerialize["datetime_created"] = o.DatetimeCreated
	toSerialize["datetime_updated"] = o.DatetimeUpdated
	toSerialize["hidden"] = o.Hidden
	toSerialize["band_locking"] = o.BandLocking
	toSerialize["dns_lookup_domain"] = o.DnsLookupDomain.Get()
	toSerialize["disable_edns"] = o.DisableEdns
	toSerialize["use_dns64"] = o.UseDns64
	toSerialize["external_connectivity"] = o.ExternalConnectivity
	return toSerialize, nil
}

func (o *WirelessNetworksItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uid",
		"alias",
		"ssid",
		"security",
		"ip_version",
		"datetime_created",
		"datetime_updated",
		"hidden",
		"band_locking",
		"dns_lookup_domain",
		"disable_edns",
		"use_dns64",
		"external_connectivity",
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

	varWirelessNetworksItem := _WirelessNetworksItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWirelessNetworksItem)

	if err != nil {
		return err
	}

	*o = WirelessNetworksItem(varWirelessNetworksItem)

	return err
}

type NullableWirelessNetworksItem struct {
	value *WirelessNetworksItem
	isSet bool
}

func (v NullableWirelessNetworksItem) Get() *WirelessNetworksItem {
	return v.value
}

func (v *NullableWirelessNetworksItem) Set(val *WirelessNetworksItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessNetworksItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessNetworksItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessNetworksItem(val *WirelessNetworksItem) *NullableWirelessNetworksItem {
	return &NullableWirelessNetworksItem{value: val, isSet: true}
}

func (v NullableWirelessNetworksItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessNetworksItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
