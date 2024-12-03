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

// checks if the AgentGroupAssignmentsGetItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentGroupAssignmentsGetItem{}

// AgentGroupAssignmentsGetItem struct for AgentGroupAssignmentsGetItem
type AgentGroupAssignmentsGetItem struct {
	Id    string                        `json:"id"`
	Group AgentGroupAssignmentsGetGroup `json:"group"`
	Agent AgentGroupAssignmentsGetAgent `json:"agent"`
	Type  string                        `json:"type"`
}

type _AgentGroupAssignmentsGetItem AgentGroupAssignmentsGetItem

// NewAgentGroupAssignmentsGetItem instantiates a new AgentGroupAssignmentsGetItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentGroupAssignmentsGetItem(
	id string,
	group AgentGroupAssignmentsGetGroup,
	agent AgentGroupAssignmentsGetAgent,
	type_ string,
) *AgentGroupAssignmentsGetItem {
	this := AgentGroupAssignmentsGetItem{}
	this.Id = id
	this.Group = group
	this.Agent = agent
	this.Type = type_
	return &this
}

// NewAgentGroupAssignmentsGetItemWithDefaults instantiates a new AgentGroupAssignmentsGetItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentGroupAssignmentsGetItemWithDefaults() *AgentGroupAssignmentsGetItem {
	this := AgentGroupAssignmentsGetItem{}
	return &this
}

// GetId returns the Id field value
func (o *AgentGroupAssignmentsGetItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentsGetItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentGroupAssignmentsGetItem) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *AgentGroupAssignmentsGetItem) GetGroup() AgentGroupAssignmentsGetGroup {
	if o == nil {
		var ret AgentGroupAssignmentsGetGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentsGetItem) GetGroupOk() (*AgentGroupAssignmentsGetGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *AgentGroupAssignmentsGetItem) SetGroup(v AgentGroupAssignmentsGetGroup) {
	o.Group = v
}

// GetAgent returns the Agent field value
func (o *AgentGroupAssignmentsGetItem) GetAgent() AgentGroupAssignmentsGetAgent {
	if o == nil {
		var ret AgentGroupAssignmentsGetAgent
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentsGetItem) GetAgentOk() (*AgentGroupAssignmentsGetAgent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent, true
}

// SetAgent sets field value
func (o *AgentGroupAssignmentsGetItem) SetAgent(v AgentGroupAssignmentsGetAgent) {
	o.Agent = v
}

// GetType returns the Type field value
func (o *AgentGroupAssignmentsGetItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentsGetItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AgentGroupAssignmentsGetItem) SetType(v string) {
	o.Type = v
}

func (o AgentGroupAssignmentsGetItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentGroupAssignmentsGetItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["agent"] = o.Agent
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *AgentGroupAssignmentsGetItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"group",
		"agent",
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

	varAgentGroupAssignmentsGetItem := _AgentGroupAssignmentsGetItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentGroupAssignmentsGetItem)

	if err != nil {
		return err
	}

	*o = AgentGroupAssignmentsGetItem(varAgentGroupAssignmentsGetItem)

	return err
}

type NullableAgentGroupAssignmentsGetItem struct {
	value *AgentGroupAssignmentsGetItem
	isSet bool
}

func (v NullableAgentGroupAssignmentsGetItem) Get() *AgentGroupAssignmentsGetItem {
	return v.value
}

func (v *NullableAgentGroupAssignmentsGetItem) Set(val *AgentGroupAssignmentsGetItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentGroupAssignmentsGetItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentGroupAssignmentsGetItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentGroupAssignmentsGetItem(
	val *AgentGroupAssignmentsGetItem,
) *NullableAgentGroupAssignmentsGetItem {
	return &NullableAgentGroupAssignmentsGetItem{value: val, isSet: true}
}

func (v NullableAgentGroupAssignmentsGetItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentGroupAssignmentsGetItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}