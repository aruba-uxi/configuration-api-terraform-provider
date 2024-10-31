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

// checks if the ServiceTestGroupAssignmentsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTestGroupAssignmentsPostRequest{}

// ServiceTestGroupAssignmentsPostRequest struct for ServiceTestGroupAssignmentsPostRequest
type ServiceTestGroupAssignmentsPostRequest struct {
	GroupId       string `json:"groupId"`
	ServiceTestId string `json:"serviceTestId"`
}

type _ServiceTestGroupAssignmentsPostRequest ServiceTestGroupAssignmentsPostRequest

// NewServiceTestGroupAssignmentsPostRequest instantiates a new ServiceTestGroupAssignmentsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTestGroupAssignmentsPostRequest(
	groupId string,
	serviceTestId string,
) *ServiceTestGroupAssignmentsPostRequest {
	this := ServiceTestGroupAssignmentsPostRequest{}
	this.GroupId = groupId
	this.ServiceTestId = serviceTestId
	return &this
}

// NewServiceTestGroupAssignmentsPostRequestWithDefaults instantiates a new ServiceTestGroupAssignmentsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTestGroupAssignmentsPostRequestWithDefaults() *ServiceTestGroupAssignmentsPostRequest {
	this := ServiceTestGroupAssignmentsPostRequest{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *ServiceTestGroupAssignmentsPostRequest) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentsPostRequest) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ServiceTestGroupAssignmentsPostRequest) SetGroupId(v string) {
	o.GroupId = v
}

// GetServiceTestId returns the ServiceTestId field value
func (o *ServiceTestGroupAssignmentsPostRequest) GetServiceTestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceTestId
}

// GetServiceTestIdOk returns a tuple with the ServiceTestId field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentsPostRequest) GetServiceTestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceTestId, true
}

// SetServiceTestId sets field value
func (o *ServiceTestGroupAssignmentsPostRequest) SetServiceTestId(v string) {
	o.ServiceTestId = v
}

func (o ServiceTestGroupAssignmentsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTestGroupAssignmentsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	toSerialize["serviceTestId"] = o.ServiceTestId
	return toSerialize, nil
}

func (o *ServiceTestGroupAssignmentsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupId",
		"serviceTestId",
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

	varServiceTestGroupAssignmentsPostRequest := _ServiceTestGroupAssignmentsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTestGroupAssignmentsPostRequest)

	if err != nil {
		return err
	}

	*o = ServiceTestGroupAssignmentsPostRequest(varServiceTestGroupAssignmentsPostRequest)

	return err
}

type NullableServiceTestGroupAssignmentsPostRequest struct {
	value *ServiceTestGroupAssignmentsPostRequest
	isSet bool
}

func (v NullableServiceTestGroupAssignmentsPostRequest) Get() *ServiceTestGroupAssignmentsPostRequest {
	return v.value
}

func (v *NullableServiceTestGroupAssignmentsPostRequest) Set(
	val *ServiceTestGroupAssignmentsPostRequest,
) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTestGroupAssignmentsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTestGroupAssignmentsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTestGroupAssignmentsPostRequest(
	val *ServiceTestGroupAssignmentsPostRequest,
) *NullableServiceTestGroupAssignmentsPostRequest {
	return &NullableServiceTestGroupAssignmentsPostRequest{value: val, isSet: true}
}

func (v NullableServiceTestGroupAssignmentsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTestGroupAssignmentsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
