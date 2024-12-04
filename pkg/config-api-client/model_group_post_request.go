/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
HPE Aruba Networking UXI Configuration

This document outlines the API contracts for HPE Aruba Networking UXI.

API version: 5.19.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GroupPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupPostRequest{}

// GroupPostRequest struct for GroupPostRequest
type GroupPostRequest struct {
	ParentId NullableString `json:"parentId,omitempty"`
	Name     string         `json:"name"`
}

type _GroupPostRequest GroupPostRequest

// NewGroupPostRequest instantiates a new GroupPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPostRequest(name string) *GroupPostRequest {
	this := GroupPostRequest{}
	this.Name = name
	return &this
}

// NewGroupPostRequestWithDefaults instantiates a new GroupPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPostRequestWithDefaults() *GroupPostRequest {
	this := GroupPostRequest{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupPostRequest) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupPostRequest) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *GroupPostRequest) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *GroupPostRequest) SetParentId(v string) {
	o.ParentId.Set(&v)
}

// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *GroupPostRequest) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *GroupPostRequest) UnsetParentId() {
	o.ParentId.Unset()
}

// GetName returns the Name field value
func (o *GroupPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupPostRequest) SetName(v string) {
	o.Name = v
}

func (o GroupPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GroupPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varGroupPostRequest := _GroupPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupPostRequest)

	if err != nil {
		return err
	}

	*o = GroupPostRequest(varGroupPostRequest)

	return err
}

type NullableGroupPostRequest struct {
	value *GroupPostRequest
	isSet bool
}

func (v NullableGroupPostRequest) Get() *GroupPostRequest {
	return v.value
}

func (v *NullableGroupPostRequest) Set(val *GroupPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPostRequest(val *GroupPostRequest) *NullableGroupPostRequest {
	return &NullableGroupPostRequest{value: val, isSet: true}
}

func (v NullableGroupPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
