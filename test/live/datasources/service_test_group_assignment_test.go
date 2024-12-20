/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package data_source_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
)

func TestServiceTestGroupAssignmentDataSource(t *testing.T) {
	const groupName = "tf_provider_acceptance_test_service_test_group_assignment_datasource"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					// create the resource to be used as a datasource
					resource "hpeuxi_group" "my_group" {
						name = "` + groupName + `"
					}

					data "hpeuxi_service_test" "my_service_test" {
						filter = {
							id = "` + config.ServiceTestID + `"
						}
					}

					resource "hpeuxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = data.hpeuxi_service_test.my_service_test.id
						group_id   = hpeuxi_group.my_group.id
					}

					// the actual datasource
					data "hpeuxi_service_test_group_assignment" "my_service_test_group_assignment" {
						filter = {
							id = hpeuxi_service_test_group_assignment.my_service_test_group_assignment.id
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						resourceName := "hpeuxi_service_test_group_assignment.my_service_test_group_assignment"
						rs := s.RootModule().Resources[resourceName]

						return util.CheckStateAgainstServiceTestGroupAssignment(
							t,
							"data.hpeuxi_service_test_group_assignment.my_service_test_group_assignment",
							util.GetServiceTestGroupAssignment(rs.Primary.ID),
						)(
							s,
						)
					},
				),
			},
		},
	})
}
