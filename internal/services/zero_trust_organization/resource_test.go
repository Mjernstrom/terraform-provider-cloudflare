package zero_trust_organization_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareAccessOrganization(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_organization.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	headerText := "My header text"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID, headerText),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", "terraform-cfapi.cloudflareaccess.com"),
					resource.TestCheckResourceAttr(name, "auth_domain", rnd+"-terraform-cfapi.cloudflareaccess.com"),
					resource.TestCheckResourceAttr(name, "is_ui_read_only", "false"),
					resource.TestCheckResourceAttr(name, "user_seat_expiration_inactive_time", "1460h"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
					resource.TestCheckResourceAttr(name, "login_design.background_color", "#FFFFFF"),
					resource.TestCheckResourceAttr(name, "login_design.text_color", "#000000"),
					resource.TestCheckResourceAttr(name, "login_design.logo_path", "https://example.com/logo.png"),
					resource.TestCheckResourceAttr(name, "login_design.header_text", headerText),
					resource.TestCheckResourceAttr(name, "login_design.footer_text", "My footer text"),
					resource.TestCheckResourceAttr(name, "session_duration", "12h"),
					resource.TestCheckResourceAttr(name, "warp_auth_session_duration", "36h"),
					resource.TestCheckResourceAttr(name, "allow_authenticate_via_warp", "false"),
				),
				ResourceName: name,
				// ImportState:      true,
				// ImportStateId:    accountID,
				// ImportStateCheck: accessOrgImportStateCheck,
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID, headerText+" updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", "terraform-cfapi.cloudflareaccess.com"),
					resource.TestCheckResourceAttr(name, "auth_domain", rnd+"-terraform-cfapi.cloudflareaccess.com"),
					resource.TestCheckResourceAttr(name, "is_ui_read_only", "false"),
					resource.TestCheckResourceAttr(name, "user_seat_expiration_inactive_time", "1460h"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
					resource.TestCheckResourceAttr(name, "login_design.background_color", "#FFFFFF"),
					resource.TestCheckResourceAttr(name, "login_design.text_color", "#000000"),
					resource.TestCheckResourceAttr(name, "login_design.logo_path", "https://example.com/logo.png"),
					resource.TestCheckResourceAttr(name, "login_design.header_text", headerText+" updated"),
					resource.TestCheckResourceAttr(name, "login_design.footer_text", "My footer text"),
					resource.TestCheckResourceAttr(name, "session_duration", "12h"),
					resource.TestCheckResourceAttr(name, "warp_auth_session_duration", "36h"),
					resource.TestCheckResourceAttr(name, "allow_authenticate_via_warp", "false"),
				),
				ResourceName: name,
				// ImportState:      true,
				// ImportStateId:    accountID,
				// ImportStateCheck: accessOrgImportStateCheck,
			},
		},
	})
}

func accessOrgImportStateCheck(instanceStates []*terraform.InstanceState) error {
	state := instanceStates[0]
	attrs := state.Attributes
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	stateChecks := []struct {
		field         string
		stateValue    string
		expectedValue string
	}{
		{field: "ID", stateValue: state.ID, expectedValue: accountID},
		{field: consts.AccountIDSchemaKey, stateValue: attrs[consts.AccountIDSchemaKey], expectedValue: accountID},
		{field: "name", stateValue: attrs["name"], expectedValue: "terraform-cfapi.cloudflareaccess.com"},
		{field: "auth_domain", stateValue: attrs["auth_domain"], expectedValue: "terraform-cfapi.cloudflareaccess.com"},
		{field: "is_ui_read_only", stateValue: attrs["is_ui_read_only"], expectedValue: "false"},
		{field: "ui_read_only_toggle_reason", stateValue: attrs["ui_read_only_toggle_reason"], expectedValue: ""}, // UI read only is off so no message returned
		{field: "user_seat_expiration_inactive_time", stateValue: attrs["user_seat_expiration_inactive_time"], expectedValue: "1460h"},
		{field: "auto_redirect_to_identity", stateValue: attrs["auto_redirect_to_identity"], expectedValue: "false"},
		{field: "login_design.#", stateValue: attrs["login_design.#"], expectedValue: "1"},
	}

	for _, check := range stateChecks {
		if check.stateValue != check.expectedValue {
			return fmt.Errorf("%s has value %q and does not match expected value %q", check.field, check.stateValue, check.expectedValue)
		}
	}

	return nil
}

func testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID, headerText string) string {
	return acctest.LoadTestCase("accessorganizationconfigbasic.tf", rnd, accountID, headerText)
}
