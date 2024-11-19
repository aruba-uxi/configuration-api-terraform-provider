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

// checks if the WiredNetworksResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WiredNetworksResponse{}

// WiredNetworksResponse struct for WiredNetworksResponse
type WiredNetworksResponse struct {
	Items []WiredNetworksItem `json:"items"`
	Count int32               `json:"count"`
	Next  NullableString      `json:"next"`
}

type _WiredNetworksResponse WiredNetworksResponse

// NewWiredNetworksResponse instantiates a new WiredNetworksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWiredNetworksResponse(
	items []WiredNetworksItem,
	count int32,
	next NullableString,
) *WiredNetworksResponse {
	this := WiredNetworksResponse{}
	this.Items = items
	this.Count = count
	this.Next = next
	return &this
}

// NewWiredNetworksResponseWithDefaults instantiates a new WiredNetworksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWiredNetworksResponseWithDefaults() *WiredNetworksResponse {
	this := WiredNetworksResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *WiredNetworksResponse) GetItems() []WiredNetworksItem {
	if o == nil {
		var ret []WiredNetworksItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksResponse) GetItemsOk() ([]WiredNetworksItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *WiredNetworksResponse) SetItems(v []WiredNetworksItem) {
	o.Items = v
}

// GetCount returns the Count field value
func (o *WiredNetworksResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *WiredNetworksResponse) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WiredNetworksResponse) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WiredNetworksResponse) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *WiredNetworksResponse) SetNext(v string) {
	o.Next.Set(&v)
}

func (o WiredNetworksResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WiredNetworksResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["count"] = o.Count
	toSerialize["next"] = o.Next.Get()
	return toSerialize, nil
}

func (o *WiredNetworksResponse) UnmarshalJSON(data []byte) (err error) {
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

	varWiredNetworksResponse := _WiredNetworksResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWiredNetworksResponse)

	if err != nil {
		return err
	}

	*o = WiredNetworksResponse(varWiredNetworksResponse)

	return err
}

type NullableWiredNetworksResponse struct {
	value *WiredNetworksResponse
	isSet bool
}

func (v NullableWiredNetworksResponse) Get() *WiredNetworksResponse {
	return v.value
}

func (v *NullableWiredNetworksResponse) Set(val *WiredNetworksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWiredNetworksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWiredNetworksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWiredNetworksResponse(val *WiredNetworksResponse) *NullableWiredNetworksResponse {
	return &NullableWiredNetworksResponse{value: val, isSet: true}
}

func (v NullableWiredNetworksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWiredNetworksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
