// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_custom_page

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessCustomPageResultDataSourceEnvelope struct {
	Result AccessCustomPageDataSourceModel `json:"result,computed"`
}

type AccessCustomPageResultListDataSourceEnvelope struct {
	Result *[]*AccessCustomPageDataSourceModel `json:"result,computed"`
}

type AccessCustomPageDataSourceModel struct {
	AccountID    types.String                              `tfsdk:"account_id" path:"account_id"`
	CustomPageID types.String                              `tfsdk:"custom_page_id" path:"custom_page_id"`
	CustomHTML   types.String                              `tfsdk:"custom_html" json:"custom_html"`
	Name         types.String                              `tfsdk:"name" json:"name"`
	Type         types.String                              `tfsdk:"type" json:"type"`
	AppCount     types.Int64                               `tfsdk:"app_count" json:"app_count"`
	CreatedAt    types.String                              `tfsdk:"created_at" json:"created_at"`
	UID          types.String                              `tfsdk:"uid" json:"uid"`
	UpdatedAt    types.String                              `tfsdk:"updated_at" json:"updated_at"`
	FindOneBy    *AccessCustomPageFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type AccessCustomPageFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
