/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWirelessNetworkResource(t *testing.T) {
	wirelessNetwork := util.GetWirelessNetwork(config.WirelessNetworkId)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a wireless_network is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "uxi_wireless_network" "wireless_network_0" {
						name = "name"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)creating a wireless_network is not supported; wireless_networks can only be\s*imported`,
				),
			},
			// Importing a wireless_network
			{
				Config: provider.ProviderConfig + `
					resource "uxi_wireless_network" "wireless_network_0" {
						name = "` + config.WirelessNetworkName + `"
					}

					import {
						to = uxi_wireless_network.wireless_network_0
						id = "` + config.WirelessNetworkId + `"
					}`,

				Check: shared.CheckStateAgainstWirelessNetwork(
					t,
					"uxi_wireless_network.wireless_network_0",
					wirelessNetwork,
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_wireless_network.wireless_network_0",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a wireless_network is not allowed
			{
				Config: provider.ProviderConfig + `
				resource "uxi_wireless_network" "wireless_network_0" {
					name = "` + config.WirelessNetworkName + `-updated-name"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)updating a wireless_network is not supported; wireless_networks can only be\s*updated through the dashboard`,
				),
			},
			// Deleting a wireless_network is not allowed
			{
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`(?s)deleting a wireless_network is not supported; wireless_networks can only\s*removed from state`,
				),
			},
			// Remove wireless_network from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_wireless_network.wireless_network_0

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}
