/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"
	"time"

	"github.com/h2non/gock"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func GenerateWirelessNetworksGetResponse(
	id string,
	postfix string,
) config_api_client.WirelessNetworksGetResponse {
	createdAt, _ := time.Parse(time.RFC3339, "2024-09-11T12:00:00.000Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2024-09-11T12:00:00.000Z")

	return config_api_client.WirelessNetworksGetResponse{
		Items: []config_api_client.WirelessNetworksGetItem{
			{
				Id:                   id,
				Ssid:                 "ssid" + postfix,
				CreatedAt:            createdAt,
				UpdatedAt:            updatedAt,
				Name:                 "name" + postfix,
				IpVersion:            config_api_client.IPVERSION_IPV4,
				Security:             *config_api_client.NewNullableString(config_api_client.PtrString("security" + postfix)),
				Hidden:               false,
				BandLocking:          "band_locking" + postfix,
				DnsLookupDomain:      *config_api_client.NewNullableString(config_api_client.PtrString("dns_lookup_domain" + postfix)),
				DisableEdns:          false,
				UseDns64:             false,
				ExternalConnectivity: false,
				Type:                 shared.WirelessNetworkType,
			},
		},
	}
}

func MockGetWirelessNetwork(id string, response interface{}, times int) {
	gock.New(MockUXIURL).
		Get(shared.WirelessNetworkPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}
