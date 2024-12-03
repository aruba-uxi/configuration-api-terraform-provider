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

// checks if the AgentGroupAssignmentPostGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentGroupAssignmentPostGroup{}

// AgentGroupAssignmentPostGroup struct for AgentGroupAssignmentPostGroup
type AgentGroupAssignmentPostGroup struct {
	Id string `json:"id"`
}

type _AgentGroupAssignmentPostGroup AgentGroupAssignmentPostGroup

// NewAgentGroupAssignmentPostGroup instantiates a new AgentGroupAssignmentPostGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentGroupAssignmentPostGroup(id string) *AgentGroupAssignmentPostGroup {
	this := AgentGroupAssignmentPostGroup{}
	this.Id = id
	return &this
}

// NewAgentGroupAssignmentPostGroupWithDefaults instantiates a new AgentGroupAssignmentPostGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentGroupAssignmentPostGroupWithDefaults() *AgentGroupAssignmentPostGroup {
	this := AgentGroupAssignmentPostGroup{}
	return &this
}

// GetId returns the Id field value
func (o *AgentGroupAssignmentPostGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentPostGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentGroupAssignmentPostGroup) SetId(v string) {
	o.Id = v
}

func (o AgentGroupAssignmentPostGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentGroupAssignmentPostGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AgentGroupAssignmentPostGroup) UnmarshalJSON(data []byte) (err error) {
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

	varAgentGroupAssignmentPostGroup := _AgentGroupAssignmentPostGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentGroupAssignmentPostGroup)

	if err != nil {
		return err
	}

	*o = AgentGroupAssignmentPostGroup(varAgentGroupAssignmentPostGroup)

	return err
}

type NullableAgentGroupAssignmentPostGroup struct {
	value *AgentGroupAssignmentPostGroup
	isSet bool
}

func (v NullableAgentGroupAssignmentPostGroup) Get() *AgentGroupAssignmentPostGroup {
	return v.value
}

func (v *NullableAgentGroupAssignmentPostGroup) Set(val *AgentGroupAssignmentPostGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentGroupAssignmentPostGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentGroupAssignmentPostGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentGroupAssignmentPostGroup(
	val *AgentGroupAssignmentPostGroup,
) *NullableAgentGroupAssignmentPostGroup {
	return &NullableAgentGroupAssignmentPostGroup{value: val, isSet: true}
}

func (v NullableAgentGroupAssignmentPostGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentGroupAssignmentPostGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}