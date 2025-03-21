
resource "cloudflare_notification_policy" "%[1]s" {
	name        = "traffic anomalies alert"
	account_id  = "%[2]s"
	description = "test description"
	enabled     =  true
	alert_type  = "traffic_anomalies_alert"
	mechanisms = {
		"email": [{"id": "test-updated@example.com"}, {"id": "test2-updated@example.com"}]
	}
	filters = {
		alert_trigger_preferences = [
			"zscore_drop"
		]
		group_by = [
			"zone"
		]
		selectors = [
			"total"
		]
		where = [
			"(origin_status_code eq 200)"
		]
		zones = ["%[3]s"]
	}
}