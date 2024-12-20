/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"bytes"
	"context"
	"crypto/md5" // #nosec G501
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func GetAgent(id string) config_api_client.AgentsGetItem {
	result, response, err := Client.ConfigurationAPI.
		AgentsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if len(result.Items) != 1 {
		panic("agent with id `" + id + "` could not be found")
	}

	return result.Items[0]
}

type ProvisionAgent struct {
	CustomerID        string
	ProvisionToken    string
	DeviceSerial      string
	DeviceGatewayHost string
}

type provisionAgentRequest struct {
	UID            string `json:"uid"`
	CustomerUID    string `json:"customer_uid"`
	ProvisionToken string `json:"provision_token"`
	PlatformName   string `json:"platform_name"`
	DeviceSerial   string `json:"device_serial"`
}

func (p ProvisionAgent) Provision() (string, error) {
	url := p.DeviceGatewayHost + "/provision-zebra-device"
	id, err := p.generateID()
	if err != nil {
		return id, err
	}

	request := provisionAgentRequest{
		UID:            id,
		CustomerUID:    p.CustomerID,
		ProvisionToken: p.ProvisionToken,
		PlatformName:   "zebra",
		DeviceSerial:   p.DeviceSerial,
	}
	jsonData, err := json.Marshal(request)
	if err != nil {
		return id, fmt.Errorf("ProvisionAgent.Provision marshal json failled: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return id, fmt.Errorf("ProvisionAgent.Provision build request failled: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return id, fmt.Errorf("ProvisionAgent.Provision request Do failled: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		return id, fmt.Errorf(
			"unexpected status code returned: %d\nresponse: %s",
			resp.StatusCode,
			string(body),
		)
	}

	return id, nil
}

func (p ProvisionAgent) generateID() (string, error) {
	// Create an MD5 hash of the serial string
	hasher := md5.New() // #nosec G401
	hasher.Write([]byte(p.DeviceSerial))
	digest := hasher.Sum(nil)

	// Use the first 16 bytes of the digest to create a UUID v3
	uuid, err := uuid.FromBytes(digest[:16])
	if err != nil {
		return "", fmt.Errorf("Failed to create hash from bytes: %w", err)
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x30 // Set the version to 3 (MD5-based UUID)
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Set the variant to RFC 4122

	return uuid.String(), nil
}
