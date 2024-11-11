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

// checks if the ServiceTestGroupAssignmentsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTestGroupAssignmentsItem{}

// ServiceTestGroupAssignmentsItem struct for ServiceTestGroupAssignmentsItem
type ServiceTestGroupAssignmentsItem struct {
	Id          string      `json:"id"`
	Group       Group       `json:"group"`
	ServiceTest ServiceTest `json:"serviceTest"`
	Type        string      `json:"type"`
}

type _ServiceTestGroupAssignmentsItem ServiceTestGroupAssignmentsItem

// NewServiceTestGroupAssignmentsItem instantiates a new ServiceTestGroupAssignmentsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTestGroupAssignmentsItem(
	id string,
	group Group,
	serviceTest ServiceTest,
	type_ string,
) *ServiceTestGroupAssignmentsItem {
	this := ServiceTestGroupAssignmentsItem{}
	this.Id = id
	this.Group = group
	this.ServiceTest = serviceTest
	this.Type = type_
	return &this
}

// NewServiceTestGroupAssignmentsItemWithDefaults instantiates a new ServiceTestGroupAssignmentsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTestGroupAssignmentsItemWithDefaults() *ServiceTestGroupAssignmentsItem {
	this := ServiceTestGroupAssignmentsItem{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceTestGroupAssignmentsItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentsItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceTestGroupAssignmentsItem) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *ServiceTestGroupAssignmentsItem) GetGroup() Group {
	if o == nil {
		var ret Group
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentsItem) GetGroupOk() (*Group, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *ServiceTestGroupAssignmentsItem) SetGroup(v Group) {
	o.Group = v
}

// GetServiceTest returns the ServiceTest field value
func (o *ServiceTestGroupAssignmentsItem) GetServiceTest() ServiceTest {
	if o == nil {
		var ret ServiceTest
		return ret
	}

	return o.ServiceTest
}

// GetServiceTestOk returns a tuple with the ServiceTest field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentsItem) GetServiceTestOk() (*ServiceTest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceTest, true
}

// SetServiceTest sets field value
func (o *ServiceTestGroupAssignmentsItem) SetServiceTest(v ServiceTest) {
	o.ServiceTest = v
}

// GetType returns the Type field value
func (o *ServiceTestGroupAssignmentsItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentsItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceTestGroupAssignmentsItem) SetType(v string) {
	o.Type = v
}

func (o ServiceTestGroupAssignmentsItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTestGroupAssignmentsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["serviceTest"] = o.ServiceTest
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ServiceTestGroupAssignmentsItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"group",
		"serviceTest",
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

	varServiceTestGroupAssignmentsItem := _ServiceTestGroupAssignmentsItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTestGroupAssignmentsItem)

	if err != nil {
		return err
	}

	*o = ServiceTestGroupAssignmentsItem(varServiceTestGroupAssignmentsItem)

	return err
}

type NullableServiceTestGroupAssignmentsItem struct {
	value *ServiceTestGroupAssignmentsItem
	isSet bool
}

func (v NullableServiceTestGroupAssignmentsItem) Get() *ServiceTestGroupAssignmentsItem {
	return v.value
}

func (v *NullableServiceTestGroupAssignmentsItem) Set(val *ServiceTestGroupAssignmentsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTestGroupAssignmentsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTestGroupAssignmentsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTestGroupAssignmentsItem(
	val *ServiceTestGroupAssignmentsItem,
) *NullableServiceTestGroupAssignmentsItem {
	return &NullableServiceTestGroupAssignmentsItem{value: val, isSet: true}
}

func (v NullableServiceTestGroupAssignmentsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTestGroupAssignmentsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
