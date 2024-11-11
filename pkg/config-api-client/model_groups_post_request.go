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

// checks if the GroupsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsPostRequest{}

// GroupsPostRequest struct for GroupsPostRequest
type GroupsPostRequest struct {
	ParentId NullableString `json:"parentId,omitempty"`
	Name     string         `json:"name"`
}

type _GroupsPostRequest GroupsPostRequest

// NewGroupsPostRequest instantiates a new GroupsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsPostRequest(name string) *GroupsPostRequest {
	this := GroupsPostRequest{}
	this.Name = name
	return &this
}

// NewGroupsPostRequestWithDefaults instantiates a new GroupsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsPostRequestWithDefaults() *GroupsPostRequest {
	this := GroupsPostRequest{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupsPostRequest) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsPostRequest) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *GroupsPostRequest) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *GroupsPostRequest) SetParentId(v string) {
	o.ParentId.Set(&v)
}

// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *GroupsPostRequest) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *GroupsPostRequest) UnsetParentId() {
	o.ParentId.Unset()
}

// GetName returns the Name field value
func (o *GroupsPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupsPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupsPostRequest) SetName(v string) {
	o.Name = v
}

func (o GroupsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GroupsPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGroupsPostRequest := _GroupsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsPostRequest)

	if err != nil {
		return err
	}

	*o = GroupsPostRequest(varGroupsPostRequest)

	return err
}

type NullableGroupsPostRequest struct {
	value *GroupsPostRequest
	isSet bool
}

func (v NullableGroupsPostRequest) Get() *GroupsPostRequest {
	return v.value
}

func (v *NullableGroupsPostRequest) Set(val *GroupsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsPostRequest(val *GroupsPostRequest) *NullableGroupsPostRequest {
	return &NullableGroupsPostRequest{value: val, isSet: true}
}

func (v NullableGroupsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
