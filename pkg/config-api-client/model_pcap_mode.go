/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
Configuration Api

Nice description goes here

API version: 5.15.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"encoding/json"
	"fmt"
)

// PcapMode the model 'PcapMode'
type PcapMode string

// List of PcapMode
const (
	LIGHT PcapMode = "light"
	FULL  PcapMode = "full"
	OFF   PcapMode = "off"
)

// All allowed values of PcapMode enum
var AllowedPcapModeEnumValues = []PcapMode{
	"light",
	"full",
	"off",
}

func (v *PcapMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PcapMode(value)
	for _, existing := range AllowedPcapModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PcapMode", value)
}

// NewPcapModeFromValue returns a pointer to a valid PcapMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPcapModeFromValue(v string) (*PcapMode, error) {
	ev := PcapMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PcapMode: valid values are %v", v, AllowedPcapModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PcapMode) IsValid() bool {
	for _, existing := range AllowedPcapModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PcapMode value
func (v PcapMode) Ptr() *PcapMode {
	return &v
}

type NullablePcapMode struct {
	value *PcapMode
	isSet bool
}

func (v NullablePcapMode) Get() *PcapMode {
	return v.value
}

func (v *NullablePcapMode) Set(val *PcapMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePcapMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePcapMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcapMode(val *PcapMode) *NullablePcapMode {
	return &NullablePcapMode{value: val, isSet: true}
}

func (v NullablePcapMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcapMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
