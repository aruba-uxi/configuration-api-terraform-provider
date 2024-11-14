package resource_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAgentResource(t *testing.T) {
	const (
		agentName        = "tf_provider_acceptance_test_agent_resource"
		agentNameUpdated = agentName + "_updated"
	)

	// we provision an agent here so that we have something to delete later on
	agentId, err := util.ProvisionAgent{
		CustomerId:        config.CustomerId,
		ProvisionToken:    os.Getenv("UXI_PROVISION_TOKEN"),
		DeviceSerial:      config.AgentCreateSerial,
		DeviceGatewayHost: config.DeviceGatewayHost,
	}.Provision()
	if err != nil {
		panic(err)
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating an agent is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name = "` + agentName + `"
					}`,

				ExpectError: regexp.MustCompile(
					`creating an agent is not supported; agents can only be imported`,
				),
			},
			// Importing an agent
			{
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name  	  = "` + agentName + `"
						notes 	  = ""
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "` + agentId + `"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", agentName),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", ""),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "pcap_mode", "light"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "id", agentId),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_agent.my_agent",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update testing
			{
				Config: provider.ProviderConfig + `
				resource "uxi_agent" "my_agent" {
					name = "` + agentNameUpdated + `"
					notes = "notes"
					pcap_mode = "off"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "id", agentId),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", agentNameUpdated),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", "notes"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "pcap_mode", "off"),
				),
			},
			// Delete testing happen automatically
		},
	})
}
