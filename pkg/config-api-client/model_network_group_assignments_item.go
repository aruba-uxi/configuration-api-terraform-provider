/*
Configuration Api

Nice description goes here

API version: 5.6.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NetworkGroupAssignmentsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkGroupAssignmentsItem{}

// NetworkGroupAssignmentsItem struct for NetworkGroupAssignmentsItem
type NetworkGroupAssignmentsItem struct {
	Id      string  `json:"id"`
	Group   Group   `json:"group"`
	Network Network `json:"network"`
	Type    string  `json:"type"`
}

type _NetworkGroupAssignmentsItem NetworkGroupAssignmentsItem

// NewNetworkGroupAssignmentsItem instantiates a new NetworkGroupAssignmentsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkGroupAssignmentsItem(
	id string,
	group Group,
	network Network,
	type_ string,
) *NetworkGroupAssignmentsItem {
	this := NetworkGroupAssignmentsItem{}
	this.Id = id
	this.Group = group
	this.Network = network
	this.Type = type_
	return &this
}

// NewNetworkGroupAssignmentsItemWithDefaults instantiates a new NetworkGroupAssignmentsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkGroupAssignmentsItemWithDefaults() *NetworkGroupAssignmentsItem {
	this := NetworkGroupAssignmentsItem{}
	return &this
}

// GetId returns the Id field value
func (o *NetworkGroupAssignmentsItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworkGroupAssignmentsItem) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *NetworkGroupAssignmentsItem) GetGroup() Group {
	if o == nil {
		var ret Group
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsItem) GetGroupOk() (*Group, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *NetworkGroupAssignmentsItem) SetGroup(v Group) {
	o.Group = v
}

// GetNetwork returns the Network field value
func (o *NetworkGroupAssignmentsItem) GetNetwork() Network {
	if o == nil {
		var ret Network
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsItem) GetNetworkOk() (*Network, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *NetworkGroupAssignmentsItem) SetNetwork(v Network) {
	o.Network = v
}

// GetType returns the Type field value
func (o *NetworkGroupAssignmentsItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworkGroupAssignmentsItem) SetType(v string) {
	o.Type = v
}

func (o NetworkGroupAssignmentsItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkGroupAssignmentsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["network"] = o.Network
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *NetworkGroupAssignmentsItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"group",
		"network",
		"type",
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

	varNetworkGroupAssignmentsItem := _NetworkGroupAssignmentsItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkGroupAssignmentsItem)

	if err != nil {
		return err
	}

	*o = NetworkGroupAssignmentsItem(varNetworkGroupAssignmentsItem)

	return err
}

type NullableNetworkGroupAssignmentsItem struct {
	value *NetworkGroupAssignmentsItem
	isSet bool
}

func (v NullableNetworkGroupAssignmentsItem) Get() *NetworkGroupAssignmentsItem {
	return v.value
}

func (v *NullableNetworkGroupAssignmentsItem) Set(val *NetworkGroupAssignmentsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkGroupAssignmentsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkGroupAssignmentsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkGroupAssignmentsItem(
	val *NetworkGroupAssignmentsItem,
) *NullableNetworkGroupAssignmentsItem {
	return &NullableNetworkGroupAssignmentsItem{value: val, isSet: true}
}

func (v NullableNetworkGroupAssignmentsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkGroupAssignmentsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
