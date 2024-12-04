/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
HPE Aruba Networking UXI Configuration

This document outlines the API contracts for HPE Aruba Networking UXI.

API version: 5.19.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SensorPatchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensorPatchResponse{}

// SensorPatchResponse struct for SensorPatchResponse
type SensorPatchResponse struct {
	Id                 string                 `json:"id"`
	Serial             string                 `json:"serial"`
	Name               string                 `json:"name"`
	ModelNumber        string                 `json:"modelNumber"`
	WifiMacAddress     NullableString         `json:"wifiMacAddress"`
	EthernetMacAddress NullableString         `json:"ethernetMacAddress"`
	AddressNote        NullableString         `json:"addressNote"`
	Longitude          NullableFloat32        `json:"longitude"`
	Latitude           NullableFloat32        `json:"latitude"`
	Notes              NullableString         `json:"notes"`
	PcapMode           NullableSensorPcapMode `json:"pcapMode"`
	Type               string                 `json:"type"`
}

type _SensorPatchResponse SensorPatchResponse

// NewSensorPatchResponse instantiates a new SensorPatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorPatchResponse(
	id string,
	serial string,
	name string,
	modelNumber string,
	wifiMacAddress NullableString,
	ethernetMacAddress NullableString,
	addressNote NullableString,
	longitude NullableFloat32,
	latitude NullableFloat32,
	notes NullableString,
	pcapMode NullableSensorPcapMode,
	type_ string,
) *SensorPatchResponse {
	this := SensorPatchResponse{}
	this.Id = id
	this.Serial = serial
	this.Name = name
	this.ModelNumber = modelNumber
	this.WifiMacAddress = wifiMacAddress
	this.EthernetMacAddress = ethernetMacAddress
	this.AddressNote = addressNote
	this.Longitude = longitude
	this.Latitude = latitude
	this.Notes = notes
	this.PcapMode = pcapMode
	this.Type = type_
	return &this
}

// NewSensorPatchResponseWithDefaults instantiates a new SensorPatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorPatchResponseWithDefaults() *SensorPatchResponse {
	this := SensorPatchResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SensorPatchResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SensorPatchResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SensorPatchResponse) SetId(v string) {
	o.Id = v
}

// GetSerial returns the Serial field value
func (o *SensorPatchResponse) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *SensorPatchResponse) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *SensorPatchResponse) SetSerial(v string) {
	o.Serial = v
}

// GetName returns the Name field value
func (o *SensorPatchResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SensorPatchResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SensorPatchResponse) SetName(v string) {
	o.Name = v
}

// GetModelNumber returns the ModelNumber field value
func (o *SensorPatchResponse) GetModelNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value
// and a boolean to check if the value has been set.
func (o *SensorPatchResponse) GetModelNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelNumber, true
}

// SetModelNumber sets field value
func (o *SensorPatchResponse) SetModelNumber(v string) {
	o.ModelNumber = v
}

// GetWifiMacAddress returns the WifiMacAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SensorPatchResponse) GetWifiMacAddress() string {
	if o == nil || o.WifiMacAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.WifiMacAddress.Get()
}

// GetWifiMacAddressOk returns a tuple with the WifiMacAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorPatchResponse) GetWifiMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WifiMacAddress.Get(), o.WifiMacAddress.IsSet()
}

// SetWifiMacAddress sets field value
func (o *SensorPatchResponse) SetWifiMacAddress(v string) {
	o.WifiMacAddress.Set(&v)
}

// GetEthernetMacAddress returns the EthernetMacAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SensorPatchResponse) GetEthernetMacAddress() string {
	if o == nil || o.EthernetMacAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.EthernetMacAddress.Get()
}

// GetEthernetMacAddressOk returns a tuple with the EthernetMacAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorPatchResponse) GetEthernetMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EthernetMacAddress.Get(), o.EthernetMacAddress.IsSet()
}

// SetEthernetMacAddress sets field value
func (o *SensorPatchResponse) SetEthernetMacAddress(v string) {
	o.EthernetMacAddress.Set(&v)
}

// GetAddressNote returns the AddressNote field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SensorPatchResponse) GetAddressNote() string {
	if o == nil || o.AddressNote.Get() == nil {
		var ret string
		return ret
	}

	return *o.AddressNote.Get()
}

// GetAddressNoteOk returns a tuple with the AddressNote field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorPatchResponse) GetAddressNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressNote.Get(), o.AddressNote.IsSet()
}

// SetAddressNote sets field value
func (o *SensorPatchResponse) SetAddressNote(v string) {
	o.AddressNote.Set(&v)
}

// GetLongitude returns the Longitude field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SensorPatchResponse) GetLongitude() float32 {
	if o == nil || o.Longitude.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorPatchResponse) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// SetLongitude sets field value
func (o *SensorPatchResponse) SetLongitude(v float32) {
	o.Longitude.Set(&v)
}

// GetLatitude returns the Latitude field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SensorPatchResponse) GetLatitude() float32 {
	if o == nil || o.Latitude.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorPatchResponse) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// SetLatitude sets field value
func (o *SensorPatchResponse) SetLatitude(v float32) {
	o.Latitude.Set(&v)
}

// GetNotes returns the Notes field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SensorPatchResponse) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}

	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorPatchResponse) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// SetNotes sets field value
func (o *SensorPatchResponse) SetNotes(v string) {
	o.Notes.Set(&v)
}

// GetPcapMode returns the PcapMode field value
// If the value is explicit nil, the zero value for SensorPcapMode will be returned
func (o *SensorPatchResponse) GetPcapMode() SensorPcapMode {
	if o == nil || o.PcapMode.Get() == nil {
		var ret SensorPcapMode
		return ret
	}

	return *o.PcapMode.Get()
}

// GetPcapModeOk returns a tuple with the PcapMode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorPatchResponse) GetPcapModeOk() (*SensorPcapMode, bool) {
	if o == nil {
		return nil, false
	}
	return o.PcapMode.Get(), o.PcapMode.IsSet()
}

// SetPcapMode sets field value
func (o *SensorPatchResponse) SetPcapMode(v SensorPcapMode) {
	o.PcapMode.Set(&v)
}

// GetType returns the Type field value
func (o *SensorPatchResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SensorPatchResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SensorPatchResponse) SetType(v string) {
	o.Type = v
}

func (o SensorPatchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensorPatchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serial"] = o.Serial
	toSerialize["name"] = o.Name
	toSerialize["modelNumber"] = o.ModelNumber
	toSerialize["wifiMacAddress"] = o.WifiMacAddress.Get()
	toSerialize["ethernetMacAddress"] = o.EthernetMacAddress.Get()
	toSerialize["addressNote"] = o.AddressNote.Get()
	toSerialize["longitude"] = o.Longitude.Get()
	toSerialize["latitude"] = o.Latitude.Get()
	toSerialize["notes"] = o.Notes.Get()
	toSerialize["pcapMode"] = o.PcapMode.Get()
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *SensorPatchResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"serial",
		"name",
		"modelNumber",
		"wifiMacAddress",
		"ethernetMacAddress",
		"addressNote",
		"longitude",
		"latitude",
		"notes",
		"pcapMode",
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

	varSensorPatchResponse := _SensorPatchResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSensorPatchResponse)

	if err != nil {
		return err
	}

	*o = SensorPatchResponse(varSensorPatchResponse)

	return err
}

type NullableSensorPatchResponse struct {
	value *SensorPatchResponse
	isSet bool
}

func (v NullableSensorPatchResponse) Get() *SensorPatchResponse {
	return v.value
}

func (v *NullableSensorPatchResponse) Set(val *SensorPatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorPatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorPatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorPatchResponse(val *SensorPatchResponse) *NullableSensorPatchResponse {
	return &NullableSensorPatchResponse{value: val, isSet: true}
}

func (v NullableSensorPatchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorPatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
