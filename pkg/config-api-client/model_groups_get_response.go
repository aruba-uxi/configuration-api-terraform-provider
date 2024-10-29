/*
Configuration Api

Nice description goes here

API version: 3.0.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GroupsGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsGetResponse{}

// GroupsGetResponse struct for GroupsGetResponse
type GroupsGetResponse struct {
	Items []GroupsGetItem `json:"items"`
	Count int32           `json:"count"`
	Next  NullableString  `json:"next"`
}

type _GroupsGetResponse GroupsGetResponse

// NewGroupsGetResponse instantiates a new GroupsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsGetResponse(items []GroupsGetItem, count int32, next NullableString) *GroupsGetResponse {
	this := GroupsGetResponse{}
	this.Items = items
	this.Count = count
	this.Next = next
	return &this
}

// NewGroupsGetResponseWithDefaults instantiates a new GroupsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsGetResponseWithDefaults() *GroupsGetResponse {
	this := GroupsGetResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *GroupsGetResponse) GetItems() []GroupsGetItem {
	if o == nil {
		var ret []GroupsGetItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GroupsGetResponse) GetItemsOk() ([]GroupsGetItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GroupsGetResponse) SetItems(v []GroupsGetItem) {
	o.Items = v
}

// GetCount returns the Count field value
func (o *GroupsGetResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GroupsGetResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GroupsGetResponse) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupsGetResponse) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsGetResponse) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *GroupsGetResponse) SetNext(v string) {
	o.Next.Set(&v)
}

func (o GroupsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["count"] = o.Count
	toSerialize["next"] = o.Next.Get()
	return toSerialize, nil
}

func (o *GroupsGetResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"count",
		"next",
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

	varGroupsGetResponse := _GroupsGetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsGetResponse)

	if err != nil {
		return err
	}

	*o = GroupsGetResponse(varGroupsGetResponse)

	return err
}

type NullableGroupsGetResponse struct {
	value *GroupsGetResponse
	isSet bool
}

func (v NullableGroupsGetResponse) Get() *GroupsGetResponse {
	return v.value
}

func (v *NullableGroupsGetResponse) Set(val *GroupsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsGetResponse(val *GroupsGetResponse) *NullableGroupsGetResponse {
	return &NullableGroupsGetResponse{value: val, isSet: true}
}

func (v NullableGroupsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
