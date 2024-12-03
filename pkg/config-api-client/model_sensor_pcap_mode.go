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
	"encoding/json"
	"fmt"
)

// SensorPcapMode the model 'SensorPcapMode'
type SensorPcapMode string

// List of SensorPcapMode
const (
	SENSORPCAPMODE_LIGHT SensorPcapMode = "light"
	SENSORPCAPMODE_FULL  SensorPcapMode = "full"
	SENSORPCAPMODE_OFF   SensorPcapMode = "off"
)

// All allowed values of SensorPcapMode enum
var AllowedSensorPcapModeEnumValues = []SensorPcapMode{
	"light",
	"full",
	"off",
}

func (v *SensorPcapMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SensorPcapMode(value)
	for _, existing := range AllowedSensorPcapModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SensorPcapMode", value)
}

// NewSensorPcapModeFromValue returns a pointer to a valid SensorPcapMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSensorPcapModeFromValue(v string) (*SensorPcapMode, error) {
	ev := SensorPcapMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SensorPcapMode: valid values are %v", v, AllowedSensorPcapModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SensorPcapMode) IsValid() bool {
	for _, existing := range AllowedSensorPcapModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensorPcapMode value
func (v SensorPcapMode) Ptr() *SensorPcapMode {
	return &v
}

type NullableSensorPcapMode struct {
	value *SensorPcapMode
	isSet bool
}

func (v NullableSensorPcapMode) Get() *SensorPcapMode {
	return v.value
}

func (v *NullableSensorPcapMode) Set(val *SensorPcapMode) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorPcapMode) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorPcapMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorPcapMode(val *SensorPcapMode) *NullableSensorPcapMode {
	return &NullableSensorPcapMode{value: val, isSet: true}
}

func (v NullableSensorPcapMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorPcapMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
