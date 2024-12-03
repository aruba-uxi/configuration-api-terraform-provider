/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
Configuration Api

Nice description goes here

API version: 5.18.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NetworkGroupAssignmentsGetNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkGroupAssignmentsGetNetwork{}

// NetworkGroupAssignmentsGetNetwork struct for NetworkGroupAssignmentsGetNetwork
type NetworkGroupAssignmentsGetNetwork struct {
	Id string `json:"id"`
}

type _NetworkGroupAssignmentsGetNetwork NetworkGroupAssignmentsGetNetwork

// NewNetworkGroupAssignmentsGetNetwork instantiates a new NetworkGroupAssignmentsGetNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkGroupAssignmentsGetNetwork(id string) *NetworkGroupAssignmentsGetNetwork {
	this := NetworkGroupAssignmentsGetNetwork{}
	this.Id = id
	return &this
}

// NewNetworkGroupAssignmentsGetNetworkWithDefaults instantiates a new NetworkGroupAssignmentsGetNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkGroupAssignmentsGetNetworkWithDefaults() *NetworkGroupAssignmentsGetNetwork {
	this := NetworkGroupAssignmentsGetNetwork{}
	return &this
}

// GetId returns the Id field value
func (o *NetworkGroupAssignmentsGetNetwork) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsGetNetwork) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworkGroupAssignmentsGetNetwork) SetId(v string) {
	o.Id = v
}

func (o NetworkGroupAssignmentsGetNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkGroupAssignmentsGetNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *NetworkGroupAssignmentsGetNetwork) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varNetworkGroupAssignmentsGetNetwork := _NetworkGroupAssignmentsGetNetwork{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkGroupAssignmentsGetNetwork)

	if err != nil {
		return err
	}

	*o = NetworkGroupAssignmentsGetNetwork(varNetworkGroupAssignmentsGetNetwork)

	return err
}

type NullableNetworkGroupAssignmentsGetNetwork struct {
	value *NetworkGroupAssignmentsGetNetwork
	isSet bool
}

func (v NullableNetworkGroupAssignmentsGetNetwork) Get() *NetworkGroupAssignmentsGetNetwork {
	return v.value
}

func (v *NullableNetworkGroupAssignmentsGetNetwork) Set(val *NetworkGroupAssignmentsGetNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkGroupAssignmentsGetNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkGroupAssignmentsGetNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkGroupAssignmentsGetNetwork(
	val *NetworkGroupAssignmentsGetNetwork,
) *NullableNetworkGroupAssignmentsGetNetwork {
	return &NullableNetworkGroupAssignmentsGetNetwork{value: val, isSet: true}
}

func (v NullableNetworkGroupAssignmentsGetNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkGroupAssignmentsGetNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}