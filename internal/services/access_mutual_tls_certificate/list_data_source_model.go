// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessMutualTLSCertificatesResultListDataSourceEnvelope struct {
	Result *[]*AccessMutualTLSCertificatesItemsDataSourceModel `json:"result,computed"`
}

type AccessMutualTLSCertificatesDataSourceModel struct {
	AccountID types.String                                        `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                        `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                         `tfsdk:"max_items"`
	Items     *[]*AccessMutualTLSCertificatesItemsDataSourceModel `tfsdk:"items"`
}

type AccessMutualTLSCertificatesItemsDataSourceModel struct {
	ID                  types.String    `tfsdk:"id" json:"id,computed"`
	AssociatedHostnames *[]types.String `tfsdk:"associated_hostnames" json:"associated_hostnames,computed"`
	CreatedAt           types.String    `tfsdk:"created_at" json:"created_at,computed"`
	ExpiresOn           types.String    `tfsdk:"expires_on" json:"expires_on,computed"`
	Fingerprint         types.String    `tfsdk:"fingerprint" json:"fingerprint,computed"`
	Name                types.String    `tfsdk:"name" json:"name,computed"`
	UpdatedAt           types.String    `tfsdk:"updated_at" json:"updated_at,computed"`
}
