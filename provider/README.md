### terraform-provider-indent

Auth is currently handled with the `indent` cli. To log in:

`$ indent auth login --space my-space-name`

To verify you have authenticated successfully, you can check `auth view`:

```
$ indent auth view --space my-space-name 2>&1 | jq
{
  ...
  "msg": "Current User",
  "user": {
    "Id": "my-email@example.com",
    ...
    "Labels": [
        "indent.com/app/config/id": "space:my-space-name"
    ]
    ...
  }
}
```

Once you've successfully authenticated, you can initialize the provider:

```terraform
provider "indent" {
  space = "my-space-name"
}
```

And start creating webhooks:

```terraform
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
```
