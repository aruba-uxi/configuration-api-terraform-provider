/*
Configuration Api

Nice description goes here

API version: 5.7.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GroupsPatchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsPatchResponse{}

// GroupsPatchResponse struct for GroupsPatchResponse
type GroupsPatchResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Parent Parent `json:"parent"`
	Type   string `json:"type"`
}

type _GroupsPatchResponse GroupsPatchResponse

// NewGroupsPatchResponse instantiates a new GroupsPatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsPatchResponse(
	id string,
	name string,
	path string,
	parent Parent,
	type_ string,
) *GroupsPatchResponse {
	this := GroupsPatchResponse{}
	this.Id = id
	this.Name = name
	this.Path = path
	this.Parent = parent
	this.Type = type_
	return &this
}

// NewGroupsPatchResponseWithDefaults instantiates a new GroupsPatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsPatchResponseWithDefaults() *GroupsPatchResponse {
	this := GroupsPatchResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GroupsPatchResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupsPatchResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupsPatchResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GroupsPatchResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupsPatchResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupsPatchResponse) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *GroupsPatchResponse) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *GroupsPatchResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *GroupsPatchResponse) SetPath(v string) {
	o.Path = v
}

// GetParent returns the Parent field value
func (o *GroupsPatchResponse) GetParent() Parent {
	if o == nil {
		var ret Parent
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *GroupsPatchResponse) GetParentOk() (*Parent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *GroupsPatchResponse) SetParent(v Parent) {
	o.Parent = v
}

// GetType returns the Type field value
func (o *GroupsPatchResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GroupsPatchResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GroupsPatchResponse) SetType(v string) {
	o.Type = v
}

func (o GroupsPatchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsPatchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	toSerialize["parent"] = o.Parent
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *GroupsPatchResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGroupsPatchResponse := _GroupsPatchResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsPatchResponse)

	if err != nil {
		return err
	}

	*o = GroupsPatchResponse(varGroupsPatchResponse)

	return err
}

type NullableGroupsPatchResponse struct {
	value *GroupsPatchResponse
	isSet bool
}

func (v NullableGroupsPatchResponse) Get() *GroupsPatchResponse {
	return v.value
}

func (v *NullableGroupsPatchResponse) Set(val *GroupsPatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsPatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsPatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsPatchResponse(val *GroupsPatchResponse) *NullableGroupsPatchResponse {
	return &NullableGroupsPatchResponse{value: val, isSet: true}
}

func (v NullableGroupsPatchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsPatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
