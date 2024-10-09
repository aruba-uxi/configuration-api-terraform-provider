package datasources

import (
	"context"

	config_api_client "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/util"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &networkGroupAssignmentDataSource{}
	_ datasource.DataSourceWithConfigure = &networkGroupAssignmentDataSource{}
)

func NewNetworkGroupAssignmentDataSource() datasource.DataSource {
	return &networkGroupAssignmentDataSource{}
}

type networkGroupAssignmentDataSource struct {
	client *config_api_client.APIClient
}

type networkGroupAssignmentDataSourceModel struct {
	ID        types.String `tfsdk:"id"`
	NetworkID types.String `tfsdk:"network_id"`
	GroupID   types.String `tfsdk:"group_id"`
	Filter    struct {
		NetworkGroupAssignmentID string `tfsdk:"network_group_assignment_id"`
	} `tfsdk:"filter"`
}

func (d *networkGroupAssignmentDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_group_assignment"
}

func (d *networkGroupAssignmentDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"network_id": schema.StringAttribute{
				Computed: true,
			},
			"group_id": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"network_group_assignment_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *networkGroupAssignmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state networkGroupAssignmentDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		GetUxiV1alpha1NetworkGroupAssignmentsGet(ctx).
		Uid(state.Filter.NetworkGroupAssignmentID)
	networkGroupAssignmentResponse, _, err := util.RetryFor429(request.Execute)

	if err != nil || len(networkGroupAssignmentResponse.NetworkGroupAssignments) != 1 {
		resp.Diagnostics.AddError(
			"Error reading Network Group Assignment",
			"Could not retrieve Network Group Assignment, unexpected error: "+err.Error(),
		)
		return
	}

	networkGroupAssignment := networkGroupAssignmentResponse.NetworkGroupAssignments[0]
	state.ID = types.StringValue(networkGroupAssignment.Uid)
	state.NetworkID = types.StringValue(networkGroupAssignment.NetworkUid)
	state.GroupID = types.StringValue(networkGroupAssignment.GroupUid)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *networkGroupAssignmentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Data Source type: Network Group Assignment. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}