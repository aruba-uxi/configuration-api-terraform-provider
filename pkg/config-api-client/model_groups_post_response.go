/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
Configuration Api

Nice description goes here

API version: 5.13.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GroupsPostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsPostResponse{}

// GroupsPostResponse struct for GroupsPostResponse
type GroupsPostResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Parent Parent `json:"parent"`
	Type   string `json:"type"`
}

type _GroupsPostResponse GroupsPostResponse

// NewGroupsPostResponse instantiates a new GroupsPostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsPostResponse(
	id string,
	name string,
	path string,
	parent Parent,
	type_ string,
) *GroupsPostResponse {
	this := GroupsPostResponse{}
	this.Id = id
	this.Name = name
	this.Path = path
	this.Parent = parent
	this.Type = type_
	return &this
}

// NewGroupsPostResponseWithDefaults instantiates a new GroupsPostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsPostResponseWithDefaults() *GroupsPostResponse {
	this := GroupsPostResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GroupsPostResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupsPostResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupsPostResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GroupsPostResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupsPostResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupsPostResponse) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *GroupsPostResponse) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *GroupsPostResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *GroupsPostResponse) SetPath(v string) {
	o.Path = v
}

// GetParent returns the Parent field value
func (o *GroupsPostResponse) GetParent() Parent {
	if o == nil {
		var ret Parent
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *GroupsPostResponse) GetParentOk() (*Parent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *GroupsPostResponse) SetParent(v Parent) {
	o.Parent = v
}

// GetType returns the Type field value
func (o *GroupsPostResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GroupsPostResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GroupsPostResponse) SetType(v string) {
	o.Type = v
}

func (o GroupsPostResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsPostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	toSerialize["parent"] = o.Parent
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *GroupsPostResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"path",
		"parent",
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

	varGroupsPostResponse := _GroupsPostResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsPostResponse)

	if err != nil {
		return err
	}

	*o = GroupsPostResponse(varGroupsPostResponse)

	return err
}

type NullableGroupsPostResponse struct {
	value *GroupsPostResponse
	isSet bool
}

func (v NullableGroupsPostResponse) Get() *GroupsPostResponse {
	return v.value
}

func (v *NullableGroupsPostResponse) Set(val *GroupsPostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsPostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsPostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsPostResponse(val *GroupsPostResponse) *NullableGroupsPostResponse {
	return &NullableGroupsPostResponse{value: val, isSet: true}
}

func (v NullableGroupsPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsPostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
