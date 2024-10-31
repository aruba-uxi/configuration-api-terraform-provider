/*
Configuration Api

Nice description goes here

API version: 5.2.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GroupsGetItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsGetItem{}

// GroupsGetItem struct for GroupsGetItem
type GroupsGetItem struct {
	Id     string         `json:"id"`
	Name   string         `json:"name"`
	Parent NullableParent `json:"parent"`
	Path   string         `json:"path"`
	Type   string         `json:"type"`
}

type _GroupsGetItem GroupsGetItem

// NewGroupsGetItem instantiates a new GroupsGetItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsGetItem(
	id string,
	name string,
	parent NullableParent,
	path string,
	type_ string,
) *GroupsGetItem {
	this := GroupsGetItem{}
	this.Id = id
	this.Name = name
	this.Parent = parent
	this.Path = path
	this.Type = type_
	return &this
}

// NewGroupsGetItemWithDefaults instantiates a new GroupsGetItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsGetItemWithDefaults() *GroupsGetItem {
	this := GroupsGetItem{}
	return &this
}

// GetId returns the Id field value
func (o *GroupsGetItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupsGetItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupsGetItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GroupsGetItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupsGetItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupsGetItem) SetName(v string) {
	o.Name = v
}

// GetParent returns the Parent field value
// If the value is explicit nil, the zero value for Parent will be returned
func (o *GroupsGetItem) GetParent() Parent {
	if o == nil || o.Parent.Get() == nil {
		var ret Parent
		return ret
	}

	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGetItem) GetParentOk() (*Parent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// SetParent sets field value
func (o *GroupsGetItem) SetParent(v Parent) {
	o.Parent.Set(&v)
}

// GetPath returns the Path field value
func (o *GroupsGetItem) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *GroupsGetItem) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *GroupsGetItem) SetPath(v string) {
	o.Path = v
}

// GetType returns the Type field value
func (o *GroupsGetItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GroupsGetItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GroupsGetItem) SetType(v string) {
	o.Type = v
}

func (o GroupsGetItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsGetItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["parent"] = o.Parent.Get()
	toSerialize["path"] = o.Path
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *GroupsGetItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"parent",
		"path",
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

	varGroupsGetItem := _GroupsGetItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsGetItem)

	if err != nil {
		return err
	}

	*o = GroupsGetItem(varGroupsGetItem)

	return err
}

type NullableGroupsGetItem struct {
	value *GroupsGetItem
	isSet bool
}

func (v NullableGroupsGetItem) Get() *GroupsGetItem {
	return v.value
}

func (v *NullableGroupsGetItem) Set(val *GroupsGetItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsGetItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsGetItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsGetItem(val *GroupsGetItem) *NullableGroupsGetItem {
	return &NullableGroupsGetItem{value: val, isSet: true}
}

func (v NullableGroupsGetItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsGetItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
