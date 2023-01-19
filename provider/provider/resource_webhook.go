package provider

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/cliutil"
)

func resourceWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWebhookCreate,
		ReadContext:   resourceWebhookRead,
		UpdateContext: resourceWebhookUpdate,
		DeleteContext: resourceWebhookDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user-facing display name of the webhook shown in console",
			},
			"dry_run": {
				Type:        schema.TypeBool,
				Default:     false,
				Description: "Testing mode: indicates whether the webhook should be skipped instead of triggered",
				Optional:    true,
			},
			"header": {
				Type:        schema.TypeSet,
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
				Default:     "POST",
				Description: "HTTP method used to send HTTP request",
				Optional:    true,
			},
			"secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Secret value used to sign the request",
				Sensitive:   true,
			},
			"space_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Space the webhook is in",
			},
			"url": {
				Type:        schema.TypeString,
				Description: "URL of deployed Webhook; requested when a permission is granted or revoked",
				Optional:    true,
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique ID idenifying the webhook",
				Optional:    true,
			},
		},
	}
}

//gocyclo:ignore
func handlersToLabels(handlers []interface{}) (map[string]string, error) {
	labels := make(map[string]string)
	labels["indent.com/kind"] = "indent.blocks.v1.Webhook"

	var pullUpdate, applyUpdate, decision map[string]interface{}

	// pull the handlers out from the set of handlers for individual processing
	for _, h := range handlers {
		handler, ok := h.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to cast handler to map[string]interface{}")
		}
		if handler["type"] == strPullUpdate && pullUpdate == nil {
			pullUpdate = handler
		} else if handler["type"] == strApplyUpdate && applyUpdate == nil {
			applyUpdate = handler
		} else if handler["type"] == strDecision && decision == nil {
			decision = handler
		} else {
			// if there are any remaining handler definitions, something's wrong
			return nil, fmt.Errorf("invalid or duplicate handler of type %s", handler["type"])
		}
	}

	// ordering of handler processing determines final value of `kind` field, which
	// we prefer to be `pull update` over `apply update` over `decision`
	if decision != nil {
		resources, ok := decision["resource_kinds"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to cast decision.resource_kinds to []interface{}")
		}
		if len(resources) != 0 {
			return nil, fmt.Errorf("resources cannot be defined for decision handlers")
		}
		labels["indent.com/webhook/kind"] = strDecision
		labels["indent.com/webhook/kind:decision"] = strconv.FormatBool(true)
	}

	if applyUpdate != nil {
		resourcesInterfaceSlice, ok := applyUpdate["resource_kinds"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to cast applyUpdate.resource_kinds to []interface{}")
		}
		resources, err := interfaceSliceToStringSlice(resourcesInterfaceSlice)
		if err != nil {
			return nil, err
		}
		labels["indent.com/webhook/kind"] = strApplyUpdate
		labels["indent.com/webhook/kind:applyUpdate"] = strconv.FormatBool(true)
		labels["indent.com/webhook/applyUpdate:kind"] = strings.Join(resources, ",")
	}

	if pullUpdate != nil {
		resourcesInterfaceSlice, ok := pullUpdate["resource_kinds"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to cast pullUpdate.resources to []interface{}")
		}
		resources, err := interfaceSliceToStringSlice(resourcesInterfaceSlice)
		if err != nil {
			return nil, err
		}
		labels["indent.com/webhook/kind"] = strPullUpdate
		labels["indent.com/webhook/kind:pullUpdate"] = strconv.FormatBool(true)
		labels["indent.com/webhook/pullUpdate:kind"] = strings.Join(resources, ",")
	}

	return labels, nil
}

func interfaceSliceToStringSlice(interfaceSlice []interface{}) ([]string, error) {
	stringSlice := make([]string, len(interfaceSlice))
	var ok bool
	for i, d := range interfaceSlice {
		stringSlice[i], ok = d.(string)
		if !ok {
			return nil, errors.New("failed to cast interface slice to string slice")
		}
	}
	return stringSlice, nil
}

var secretChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func generateRandomSecret(length int) (string, error) {
	ll := big.NewInt(int64(len(secretChars)))
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, ll)
		if err != nil {
			return "", err
		}
		b[i] = secretChars[num.Int64()]
	}
	return "wks0" + string(b), nil
}

func inflateHeaders(headers []interface{}) (map[string]*indentv1.Header, error) {
	// on the wire, we send a map of header name -> credential, value
	v1Headers := make(map[string]*indentv1.Header)
	for _, h := range headers {
		header, ok := h.(map[string]interface{})
		if !ok {
			return nil, errors.New("failed to cast header to map[string]interface{}")
		}
		headerName, ok := header["name"].(string)
		if !ok {
			return nil, errors.New("failed to cast header[\"name\"] to string")
		}
		value, ok := header["value"].(string)
		if !ok {
			return nil, errors.New("failed to cast header[\"value\"] to string")
		}

		credential, ok := header["credential"].(string)
		if !ok {
			return nil, errors.New("failed to cast header[\"credential\"] to string")
		}

		v1Headers[headerName] = &indentv1.Header{
			Value:      value,
			Credential: credential,
		}
	}
	return v1Headers, nil
}

// inflateWebhook inflates a Webhook from a given ResourceData. If an existing
// webhook is passed in, it will inflate into that to update values; if no
// webhook is passed in a new one will be created.
func inflateWebhook(w *indentv1.Webhook, d *schema.ResourceData) (*indentv1.Webhook, error) {
	var ok bool
	var err error

	if w == nil {
		w = &indentv1.Webhook{}
	}

	if w.DisplayName, ok = d.Get("display_name").(string); !ok {
		return nil, errors.New("failed to cast display_name to string")
	}
	if w.Url, ok = d.Get("url").(string); !ok {
		return nil, errors.New("failed to cast url to string")
	}
	if w.DryRun, ok = d.Get("dry_run").(bool); !ok {
		return nil, errors.New("failed to cast dry_run to bool")
	}
	if w.Method, ok = d.Get("method").(string); !ok {
		return nil, errors.New("failed to cast method to string")
	}
	if w.Secret, ok = d.Get("secret").(string); !ok {
		return nil, errors.New("failed to cast secret to string")
	}

	var handler *schema.Set
	if handler, ok = d.Get("handler").(*schema.Set); !ok {
		return nil, errors.New("failed to cast handler to *schema.Set")
	}

	var labels map[string]string
	handlers := handler.List()
	if labels, err = handlersToLabels(handlers); err != nil {
		return nil, err
	}

	w.Labels = labels

	h, ok := d.Get("header").(*schema.Set)
	if !ok {
		return nil, errors.New("failed to cast header to *schema.Set")
	}

	headers, err := inflateHeaders(h.List())
	if err != nil {
		return nil, err
	}
	w.Headers = headers

	return w, nil
}

func resourceWebhookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	f, ok := m.(cliutil.Factory)
	if !ok {
		return diag.Errorf("failed to initial API factory from cast")
	}
	client := f.API(ctx).Webhooks()

	providedSecret, ok := d.Get("secret").(string)
	if !ok {
		return diag.Errorf("failed to cast secret to string")
	}

	if providedSecret == "" {
		defaultLength := 40 // length set to match secrets created in frontend
		secret, err := generateRandomSecret(defaultLength)
		if err != nil {
			return diag.FromErr(err)
		}

		if err = d.Set("secret", secret); err != nil {
			return diag.FromErr(err)
		}
	}

	webhook, err := inflateWebhook(nil, d)
	if err != nil {
		return diag.FromErr(err)
	}

	webhook.SpaceName = f.Config().Space

	out, err := client.CreateWebhook(ctx, &indentv1.CreateWebhookRequest{
		SpaceName: f.Config().Space,
		Webhook:   webhook,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("space_name", f.Config().Space); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("id", out.Name); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(out.Name)
	return nil
}

func resourceWebhookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	f, ok := m.(cliutil.Factory)
	if !ok {
		return diag.Errorf("failed to initial API factory from cast")
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

func resourceWebhookUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	f, ok := m.(cliutil.Factory)
	if !ok {
		return diag.Errorf("failed to initial API factory from cast")
	}
	client := f.API(ctx).Webhooks()

	webhook, err := client.GetWebhook(ctx, &indentv1.GetWebhookRequest{
		SpaceName:   f.Config().Space,
		WebhookName: d.Get("id").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	webhook, err = inflateWebhook(webhook, d)
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = client.UpdateWebhook(ctx, &indentv1.UpdateWebhookRequest{
		SpaceName:   f.Config().Space,
		WebhookName: d.Get("id").(string),
		Webhook:     webhook,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceWebhookRead(ctx, d, m)
}

func resourceWebhookDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	f, ok := m.(cliutil.Factory)
	if !ok {
		return diag.Errorf("failed to initial API factory from cast")
	}
	blockClient := f.API(ctx).Blocks()

	// should use DeleteWebhook, but it's currently broken -- ID-1926
	_, err := blockClient.DeleteBlock(ctx, &indentv1.DeleteBlockRequest{
		SpaceName: f.Config().Space,
		BlockName: d.Get("id").(string),
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
