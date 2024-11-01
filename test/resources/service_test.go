package resource_test

import (
	"github.com/aruba-uxi/configuration-api-terraform-provider/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/test/util"
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/internal/provider/resources"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestServiceTestResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a service_test is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)creating a service_test is not supported; service_tests can only be\s*imported`,
				),
			},
			// Importing a service_test
			{
				PreConfig: func() {
					resources.GetServiceTest = func(uid string) resources.ServiceTestResponseModel {
						return util.GenerateServiceTestResponseModel(uid, "")
					}
				},
				Config: provider.ProviderConfig + `
					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "uid"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_service_test.my_service_test",
						"name",
						"name",
					),
					resource.TestCheckResourceAttr("uxi_service_test.my_service_test", "id", "uid"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_service_test.my_service_test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a service_test is not allowed
			{
				Config: provider.ProviderConfig + `
				resource "uxi_service_test" "my_service_test" {
					name = "updated_name"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)updating a service_test is not supported; service_tests can only be updated\s*through the dashboard`,
				),
			},
			// Deleting a service_test is not allowed
			{
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`(?s)deleting a service_test is not supported; service_tests can only removed from\s*state`,
				),
			},
			// Remove service_test from state
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

	mockOAuth.Mock.Disable()
}