package resource_test

import (
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAgentGroupAssignmentResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a agent group assignment
			{
				PreConfig: func() {
					// required for agent import
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						1,
					)

					// required for agent group assignment create
					util.MockPostAgentGroupAssignment(
						util.GenerateAgentGroupAssignmentRequest("agent_group_assignment_uid", ""),
						util.GenerateAgentGroupAssignmentResponse("agent_group_assignment_uid", ""),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_agent" "my_agent" {
						name 			= "name"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_uid"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id       = uxi_agent.my_agent.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						"agent_uid",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						"group_uid",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"id",
						"agent_group_assignment_uid",
					),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
				},
				ResourceName:      "uxi_agent_group_assignment.my_agent_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid_2", "_2"),
							},
						),
						2,
					)
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
							},
						),
						2,
					)
					util.MockGetGroup(
						"group_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
							},
						),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						2,
					)

					// required for creating another group
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid_2", "_2", "_2"),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
						),
						1,
					)

					// required for agent group assignment create
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid_2",
									"_2",
								),
							},
						),
						2,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
					util.MockPostAgentGroupAssignment(
						util.GenerateAgentGroupAssignmentRequest(
							"agent_group_assignment_uid_2",
							"_2",
						),
						util.GenerateAgentGroupAssignmentResponse(
							"agent_group_assignment_uid_2",
							"_2",
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_agent" "my_agent" {
						name 			= "name"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_uid"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "name_2"
						parent_group_id = "parent_uid_2"
					}

					resource "uxi_agent" "my_agent_2" {
						name 			= "name_2"
						notes 			= "notes_2"
						pcap_mode 		= "light_2"
					}

					import {
						to = uxi_agent.my_agent_2
						id = "agent_uid_2"
					}

					// the assignment update, updated from agent/group to agent_2/group_2
					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id       = uxi_agent.my_agent_2.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						"agent_uid_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						"group_uid_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"id",
						"agent_group_assignment_uid_2",
					),
				),
			},
			// Delete testing automatically occurs in TestCase
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
							},
						),
						2,
					)
					util.MockGetAgent(
						"agent_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid_2", "_2"),
							},
						),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						2,
					)
					util.MockGetGroup(
						"group_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
							},
						),
						1,
					)
					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteGroup("group_uid_2", 1)
					util.MockDeleteAgent("agent_uid", 1)
					util.MockDeleteAgent("agent_uid_2", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
