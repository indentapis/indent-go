// Package provider implements the provider datasources and resource actions
package provider

import (
	"context"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/cliutil"
)

const (
	strPullUpdate  = "pullUpdate"
	strApplyUpdate = "applyUpdate"
	strDecision    = "decision"
)

var webhookSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The user-facing display name of the webhook shown in console",
		Optional:    true,
	},
	"dry_run": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Testing mode: indicates whether the webhook should be skipped instead of triggered",
		Optional:    true,
	},
	"header": {
		Type:        schema.TypeSet,
		Computed:    true,
		Description: "HTTP headers included when request to webhook is made",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"credential": {
					Type:        schema.TypeString,
					Description: "Name of Credential to use as header value (instead of string value)",
					Optional:    true,
				},
				"value": {
					Type:        schema.TypeString,
					Description: "String to use as header value (instead of Credential)",
					Optional:    true,
				},
				"name": {
					Type:        schema.TypeString,
					Description: "Header name",
					Required:    true,
				},
			},
		},
	},
	"handler": {
		Type:        schema.TypeSet,
		Computed:    true,
		Description: "Types of requests the webhook should handle",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:        schema.TypeString,
					Description: "Type of webhook (pullUpdate/applyUpdate/decision)",
					Optional:    true,
				},
				"resource_kinds": {
					Type:        schema.TypeList,
					Description: "Resource kinds webhook supports",
					Optional:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	},
	"method": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "HTTP method used to send HTTP request",
		Optional:    true,
	},
	"id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Unique ID idenifying the webhook",
		Optional:    true,
	},
	"secret": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Secret value used to sign the request",
		Optional:    true,
		Sensitive:   true,
	},
	"space_name": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Name of the Space the webhook is in",
		Optional:    true,
	},
	"url": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "URL of deployed Webhook; requested when a permission is granted or revoked",
		Optional:    true,
	},
}

func dataSourceWebhook() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWebhookRead,
		Schema:      webhookSchema,
	}
}

func dataSourceWebhookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	f, ok := m.(cliutil.Factory)
	if !ok {
		diag.Errorf("failed to fetch API factory")
	}
	client := f.API(ctx).Webhooks()

	webhook, err := client.GetWebhook(ctx, &indentv1.GetWebhookRequest{
		SpaceName:   f.Config().Space,
		WebhookName: d.Get("id").(string),
	})

	if err != nil {
		return diag.FromErr(err)
	}

	flat := flattenWebhook(webhook)
	for k, v := range flat {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(webhook.Name)
	return nil
}

func flattenHandlers(labels map[string]string) []map[string]interface{} {
	handlers := make([]map[string]interface{}, 0)

	valStr, ok := labels["indent.com/webhook/kind:pullUpdate"]
	if isTrue, _ := strconv.ParseBool(valStr); ok && isTrue {
		pullUpdate := make(map[string]interface{})
		pullUpdate["type"] = strPullUpdate

		valStr, ok = labels["indent.com/webhook/pullUpdate:kind"]
		if ok {
			pullUpdate["resource_kinds"] = strings.Split(valStr, ",")
		}

		handlers = append(handlers, pullUpdate)
	}

	valStr, ok = labels["indent.com/webhook/kind:applyUpdate"]
	if isTrue, _ := strconv.ParseBool(valStr); ok && isTrue {
		applyUpdate := make(map[string]interface{})
		applyUpdate["type"] = strApplyUpdate

		if valStr, ok = labels["indent.com/webhook/applyUpdate:kind"]; ok {
			applyUpdate["resource_kinds"] = strings.Split(valStr, ",")
		}

		handlers = append(handlers, applyUpdate)
	}

	valStr, ok = labels["indent.com/webhook/kind:decision"]
	if isTrue, _ := strconv.ParseBool(valStr); ok && isTrue {
		decision := make(map[string]interface{})
		decision["type"] = strDecision

		handlers = append(handlers, decision)
	}
	return handlers
}

// flattenWebhook converts a Webhook object into a generic map that fits
// Terraform's expectations for loading into a schema.ResourceData.
// As this is used for both the data source and the resource schema definition,
// we shouldn't just set directly into a ResourceData.
func flattenWebhook(wh *indentv1.Webhook) map[string]interface{} {
	m := make(map[string]interface{})

	m["url"] = wh.Url
	m["space_name"] = wh.SpaceName
	m["display_name"] = wh.DisplayName
	m["secret"] = wh.Secret
	m["id"] = wh.Name
	m["method"] = wh.Method
	m["dry_run"] = wh.DryRun
	m["handler"] = flattenHandlers(wh.Labels)

	headers := make([]map[string]interface{}, len(wh.Headers))
	i := 0
	for hName, hValue := range wh.Headers {
		headers[i] = make(map[string]interface{})
		headers[i]["name"] = hName
		headers[i]["value"] = hValue.Value
		headers[i]["credential"] = hValue.Credential
		i++
	}

	m["header"] = headers

	return m
}
