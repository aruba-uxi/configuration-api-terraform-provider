package resource_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-configuration/test/live/config"
	"github.com/aruba-uxi/terraform-provider-configuration/test/live/util"
	"github.com/aruba-uxi/terraform-provider-configuration/test/mocked/provider"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestServiceTestGroupAssignmentResource(t *testing.T) {
	const groupName = "tf_acceptance_test_service_test_group_assignment"
	const group2Name = "tf_acceptance_test_service_test_group_assignment_two"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a serviceTest group assignment
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "` + config.ServiceTestName + `"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "` + config.ServiceTestUid + `"
					}

					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						config.ServiceTestUid,
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						util.GetGroupByName(groupName).Id,
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						config.ServiceTestUid,
					),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_service_test_group_assignment.my_service_test_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "` + groupName + `"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "` + config.ServiceTestName + `"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from service_test/group to service_test/group_2
					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						config.ServiceTestUid,
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						util.GetGroupByName(group2Name).Id,
					),
				),
			},
			// Remove serviceTests from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_service_test.my_service_test

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}