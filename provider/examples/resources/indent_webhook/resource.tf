resource "indent_webhook" "integration-okta-webhook" {
  display_name = "integration-okta-webhook"
  url          = "https://xxxxxxxx.lambda-url.us-west-2.on.aws/"
  dry_run      = false

  header {
    name  = "X-MY-CUSTOM-HEADER"
    value = "header-value"
  }

  handler {
    type      = "pullUpdate"
    resources = ["okta.v1.Group", "okta.v1.User"]
  }

  handler {
    type      = "applyUpdate"
    resources = ["okta.v1.User"]
  }
}

output "okta_webhook_id" {
  description = "block ID of the new webhook"
  value       = indent_webhook.integration-okta-webhook.id
}

output "okta_webhook_secret" {
  description = "signing secret of the new webhook"
  value       = indent_webhook.integration-okta-webhook.secret
  sensitive   = true
}
