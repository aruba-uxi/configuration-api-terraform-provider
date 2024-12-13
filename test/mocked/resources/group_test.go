/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/stretchr/testify/assert"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func setupGroupCreateMocks(groupID string, parentID *string, name string) {
	groupPath := util.MockRootGroupID + "." + groupID
	realParent := util.MockRootGroupID
	if parentID != nil {
		realParent := *parentID
		groupPath = realParent + "." + groupID
	}

	postRequest := createGroupPostRequest(name, parentID)
	postResponse := createGroupPostResponse(groupID, groupPath, name, realParent)
	util.MockPostGroup(postRequest, postResponse, 1)

	getResponse := createGroupGetResponse(groupID, realParent, groupPath, name)
	util.MockGetGroup(groupID, getResponse, 1)
}

func setupGroupUpdateMocks(groupID string, parentID string, name string) {
	path := parentID + "." + groupID
	if parentID == util.MockRootGroupID {
		util.MockGetGroup(util.MockRootGroupID, util.GenerateRootGroupGetResponse(), 1)
	} else {
		util.MockGetGroup(
			parentID,
			createGroupGetResponse(
				parentID,
				"fake_parent_id",
				"fake_parent_id."+parentID,
				"fake_parent",
			),
			1,
		)
	}

	util.MockPatchGroup(
		"id",
		createGroupPatchRequest(name),
		createGroupPatchResponse(groupID, name, path, parentID),
		1,
	)

	util.MockGetGroup(
		"id",
		createGroupGetResponse(groupID, parentID, path, name),
		1,
	)
}

func setupGroupDeleteMocks(groupID string, parentID *string, name string) {
	groupPath := util.MockRootGroupID + "." + groupID
	realParent := util.MockRootGroupID
	if parentID != nil {
		realParent := *parentID
		groupPath = realParent + "." + groupID
	}
	getResponse := createGroupGetResponse(groupID, realParent, groupPath, name)
	util.MockGetGroup(
		groupID,
		getResponse,
		1,
	)
	util.MockDeleteGroup(groupID, 1)
}

func setupGroupImportMocks(groupID string, parentID *string, name string) {
	if parentID != nil {
		realParent := *parentID
		groupPath := realParent + "." + groupID

		getParentResponse := createGroupGetResponse(
			realParent,
			"fake_parent_id",
			"fake_parent_id."+realParent,
			"fake parent",
		)

		util.MockGetGroup(realParent, getParentResponse, 1)

		getResponse := createGroupGetResponse(groupID, realParent, groupPath, name)
		util.MockGetGroup(groupID, getResponse, 1)
	} else {
		groupPath := util.MockRootGroupID + "." + groupID

		rootResponse := util.GenerateRootGroupGetResponse()
		util.MockGetGroup(util.MockRootGroupID, rootResponse, 1)

		getResponse := createGroupGetResponse(groupID, util.MockRootGroupID, groupPath, name)
		util.MockGetGroup(groupID, getResponse, 1)
	}
}

func createGroupPostRequest(name string, parentID *string) config_api_client.GroupPostRequest {
	postRequest := config_api_client.NewGroupPostRequest(name)
	if parentID != nil {
		realParent := *parentID
		postRequest.SetParentId(realParent)
	}

	return *postRequest
}

func createGroupPostResponse(
	groupID string,
	groupPath string,
	name string,
	parentID string,
) config_api_client.GroupPostResponse {
	return *config_api_client.NewGroupPostResponse(
		groupID,
		name,
		groupPath,
		*config_api_client.NewGroupPostParent(parentID),
		shared.GroupType,
	)
}

func createGroupGetResponse(
	groupID string,
	parentID string,
	groupPath string,
	name string,
) config_api_client.GroupsGetResponse {
	getResponseItems := []config_api_client.GroupsGetItem{*config_api_client.NewGroupsGetItem(
		groupID,
		name,
		*config_api_client.NewNullableGroupsGetParent(config_api_client.NewGroupsGetParent(parentID)),
		groupPath,
		shared.GroupType,
	)}
	getResponse := config_api_client.NewGroupsGetResponse(
		getResponseItems,
		1,
		*config_api_client.NewNullableString(nil),
	)

	return *getResponse
}

func createGroupPatchRequest(name string) config_api_client.GroupPatchRequest {
	request := config_api_client.NewGroupPatchRequest()
	request.SetName(name)

	return *request
}

func createGroupPatchResponse(
	groupID string,
	name string,
	parentID string,
	path string,
) config_api_client.GroupPatchResponse {
	response := config_api_client.NewGroupPatchResponse(
		groupID,
		name,
		path,
		*config_api_client.NewGroupPatchParent(parentID),
		shared.GroupType,
	)

	return *response
}

func Test_CreateGroupResource_ShouldSucceed(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()
	parentID := new(string)
	*parentID = "create_parent"
	testGroupID := "create_id"
	testName := "test_name"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					setupGroupCreateMocks(testGroupID, parentID, testName)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "test_name"
					parent_group_id = "create_parent"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"hpeuxi_group.my_group",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectUnknownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("id"),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("parent_group_id"),
							knownvalue.StringExact("create_parent"),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("name"),
							knownvalue.StringExact(testName),
						),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", testName),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						"create_parent",
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", testGroupID),
				),
			},
			// Delete
			{
				PreConfig: func() {
					setupGroupDeleteMocks(testGroupID, parentID, testName)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_ImportGroupResource_ShouldSucceed(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()
	parentID := new(string)
	*parentID = "import_parent"
	testName := "import_name"
	testID := "import_id"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					setupGroupCreateMocks(testID, parentID, testName)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "import_name"
					parent_group_id = "import_parent"
				}`,
			},
			// ImportState
			{
				PreConfig: func() {
					setupGroupImportMocks(testID, parentID, testName)
				},
				ResourceName:      "hpeuxi_group.my_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete
			{
				PreConfig: func() {
					setupGroupDeleteMocks(testID, parentID, testName)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_CreateGroupResource_WithRootParent(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a group attached to the root
			{
				PreConfig: func() {
					setupGroupCreateMocks("id", nil, "name")
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name = "name"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"hpeuxi_group.my_group",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectUnknownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("id"),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("parent_group_id"),
							knownvalue.Null(),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("name"),
							knownvalue.StringExact("name"),
						),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name"),
					resource.TestCheckNoResourceAttr("hpeuxi_group.my_group", "parent_group_id"),
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					// existing group
					setupGroupDeleteMocks("id", nil, "name")
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_ImportGroupResource_WithRoot_ShouldFail(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Importing the root group does not work
			{
				PreConfig: func() {
					setupGroupImportMocks(util.MockRootGroupID, nil, "name")
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_root_group" {
					name            = "name"
				}

				import {
					to = hpeuxi_group.my_root_group
					id = "` + util.MockRootGroupID + `"
				}`,
				ExpectError: regexp.MustCompile(`The root group cannot be used as a resource`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_UpdateGroupResource_WithoutRecreate_ShouldSucceed(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					parentID := new(string)
					*parentID = "parent_id"
					setupGroupCreateMocks("id", parentID, "name")
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
			},
			// Update that does not trigger a recreate
			{
				PreConfig: func() {
					parentID := new(string)
					*parentID = "parent_id"
					setupGroupImportMocks("id", parentID, "name")

					// updated group
					setupGroupUpdateMocks("id", *parentID, "name_2")
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_id"
					}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"hpeuxi_group.my_group",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("name"),
							knownvalue.StringExact("name_2"),
						),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name_2"),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						"parent_id",
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
				),
			},
			// Delete
			{
				PreConfig: func() {
					parentID := new(string)
					*parentID = "parent_id"
					setupGroupDeleteMocks("id", parentID, "name_2")
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_UpdateGroupResource_WithoutParent_ShouldSucceed(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					setupGroupCreateMocks("id", nil, "name")
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
				}`,
			},
			// Update that does not trigger a recreate
			{
				PreConfig: func() {
					setupGroupImportMocks("id", nil, "name")
					setupGroupUpdateMocks("id", util.MockRootGroupID, "name_2")
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name_2"
					}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"hpeuxi_group.my_group",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("name"),
							knownvalue.StringExact("name_2"),
						),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name_2"),
					resource.TestCheckNoResourceAttr("hpeuxi_group.my_group", "parent_group_id"),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
				),
			},
			// Delete
			{
				PreConfig: func() {
					setupGroupDeleteMocks("id", nil, "name_2")
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_UpdateGroupResource_WithRecreate_ShouldSucceed(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					parentID := new(string)
					*parentID = "parent_id"
					setupGroupCreateMocks("id", parentID, "name")
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
			},
			// Update that does trigger a recreate
			{
				PreConfig: func() {
					oldParentID := new(string)
					*oldParentID = "parent_id"
					setupGroupImportMocks("id", oldParentID, "name")
					setupGroupDeleteMocks("id", oldParentID, "name")

					newParentID := new(string)
					*newParentID = "parent_id_2"
					setupGroupCreateMocks("new_id", newParentID, "name")
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id_2"
					}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"hpeuxi_group.my_group",
							plancheck.ResourceActionDestroyBeforeCreate,
						),
						plancheck.ExpectUnknownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("id"),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("name"),
							knownvalue.StringExact("name"),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("parent_group_id"),
							knownvalue.StringExact("parent_id_2"),
						),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						"parent_id_2",
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "new_id"),
				),
			},
			// Delete
			{
				PreConfig: func() {
					parentID := new(string)
					*parentID = "parent_id_2"
					setupGroupDeleteMocks("new_id", parentID, "name")
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestGroupResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Post("/networking-uxi/v1alpha1/groups").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("id", "", ""),
						util.GenerateGroupPostResponse("id", "", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 2)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Read
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/groups").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				ResourceName:      "hpeuxi_group.my_group",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Update
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// new group
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Patch("/networking-uxi/v1alpha1/groups/id").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPatchGroup(
						"id",
						util.GenerateGroupPatchRequest("_2"),
						util.GenerateGroupPatchResponse("id", "_2", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "_2", ""), 3)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_id"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name_2"),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Delete
			{
				PreConfig: func() {
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					util.MockDeleteGroup("id", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestGroupResourceHttpErrorHandling(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// read HTTP error
			{
				PreConfig: func() {
					gock.New(util.MockUXIURL).
						Get("/networking-uxi/v1alpha1/groups").
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					import {
						to = hpeuxi_group.my_group
						id = "id"
					}
				`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Read not found
			{
				PreConfig: func() {
					util.MockGetGroup("id", util.EmptyGetListResponse, 1)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					import {
						to = hpeuxi_group.my_group
						id = "id"
					}
				`,

				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Create HTTP error
			{
				PreConfig: func() {
					gock.New(util.MockUXIURL).
						Post("/networking-uxi/v1alpha1/groups").
						Reply(http.StatusBadRequest).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusBadRequest,
							"errorCode":      "HPE_GL_ERROR_BAD_REQUEST",
							"message":        "Validation error - bad request",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Create group in prep for next step
			{
				PreConfig: func() {
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("id", "", ""),
						util.GenerateGroupPostResponse("id", "", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 2)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
				),
			},
			// Update HTTP error
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// new group - with error
					gock.New(util.MockUXIURL).
						Patch("/networking-uxi/v1alpha1/groups/id").
						Reply(http.StatusUnprocessableEntity).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusUnprocessableEntity,
							"errorCode":      "HPE_GL_UXI_DUPLICATE_SIBLING_GROUP_NAME",
							"message":        "Unable to create group - a sibling group already has the specified name",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_id"
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Unable to create group - a sibling group already has the specified name\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Delete HTTP error
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// delete group - with error
					gock.New(util.MockUXIURL).
						Delete("/networking-uxi/v1alpha1/groups/id").
						Reply(http.StatusUnprocessableEntity).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusUnprocessableEntity,
							"errorCode":      "HPE_GL_UXI_GROUP_CANNOT_BE_DELETED",
							"message":        "Unable to delete group",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig,
				ExpectError: regexp.MustCompile(
					`(?s)Unable to delete group\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Actually delete group for cleanup reasons
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// delete group
					util.MockDeleteGroup("id", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
