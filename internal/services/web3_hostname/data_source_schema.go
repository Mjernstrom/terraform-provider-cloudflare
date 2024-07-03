// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &Web3HostnameDataSource{}
var _ datasource.DataSourceWithValidateConfig = &Web3HostnameDataSource{}

func (r Web3HostnameDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Optional: true,
			},
			"description": schema.StringAttribute{
				Description: "An optional description of the hostname.",
				Optional:    true,
			},
			"dnslink": schema.StringAttribute{
				Description: "DNSLink value used if the target is ipfs.",
				Optional:    true,
			},
			"modified_on": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Description: "The hostname that will point to the target gateway via CNAME.",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the hostname's activation.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "pending", "deleting", "error"),
				},
			},
			"target": schema.StringAttribute{
				Description: "Target gateway of the hostname.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ethereum", "ipfs", "ipfs_universal_path"),
				},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (r *Web3HostnameDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *Web3HostnameDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
