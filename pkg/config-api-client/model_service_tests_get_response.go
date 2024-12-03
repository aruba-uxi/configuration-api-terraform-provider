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

// checks if the ServiceTestsGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTestsGetResponse{}

// ServiceTestsGetResponse struct for ServiceTestsGetResponse
type ServiceTestsGetResponse struct {
	Items []ServiceTestsGetItem `json:"items"`
	Count int32                 `json:"count"`
	Next  NullableString        `json:"next"`
}

type _ServiceTestsGetResponse ServiceTestsGetResponse

// NewServiceTestsGetResponse instantiates a new ServiceTestsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTestsGetResponse(
	items []ServiceTestsGetItem,
	count int32,
	next NullableString,
) *ServiceTestsGetResponse {
	this := ServiceTestsGetResponse{}
	this.Items = items
	this.Count = count
	this.Next = next
	return &this
}

// NewServiceTestsGetResponseWithDefaults instantiates a new ServiceTestsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTestsGetResponseWithDefaults() *ServiceTestsGetResponse {
	this := ServiceTestsGetResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ServiceTestsGetResponse) GetItems() []ServiceTestsGetItem {
	if o == nil {
		var ret []ServiceTestsGetItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ServiceTestsGetResponse) GetItemsOk() ([]ServiceTestsGetItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ServiceTestsGetResponse) SetItems(v []ServiceTestsGetItem) {
	o.Items = v
}

// GetCount returns the Count field value
func (o *ServiceTestsGetResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ServiceTestsGetResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ServiceTestsGetResponse) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServiceTestsGetResponse) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceTestsGetResponse) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *ServiceTestsGetResponse) SetNext(v string) {
	o.Next.Set(&v)
}

func (o ServiceTestsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTestsGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["count"] = o.Count
	toSerialize["next"] = o.Next.Get()
	return toSerialize, nil
}

func (o *ServiceTestsGetResponse) UnmarshalJSON(data []byte) (err error) {
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

	varServiceTestsGetResponse := _ServiceTestsGetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTestsGetResponse)

	if err != nil {
		return err
	}

	*o = ServiceTestsGetResponse(varServiceTestsGetResponse)

	return err
}

type NullableServiceTestsGetResponse struct {
	value *ServiceTestsGetResponse
	isSet bool
}

func (v NullableServiceTestsGetResponse) Get() *ServiceTestsGetResponse {
	return v.value
}

func (v *NullableServiceTestsGetResponse) Set(val *ServiceTestsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTestsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTestsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTestsGetResponse(
	val *ServiceTestsGetResponse,
) *NullableServiceTestsGetResponse {
	return &NullableServiceTestsGetResponse{value: val, isSet: true}
}

func (v NullableServiceTestsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTestsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}