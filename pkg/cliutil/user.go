package cliutil

import (
	"context"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/common"
)

// kindOrder is the ordering kinds being chosen as the return user.
var kindOrder = [...]string{
	common.KindUser,
	common.KindSlackUser,
	common.KindOktaUser,
}

// lookupUser fetches details from the Resources representing a User.
func lookupUser(ctx context.Context, f Factory, user *auditv1.Resource) *auditv1.Resource {
	logger := f.Logger()
	client := f.API(ctx).Resources()

	emailRequirement, err := labels.NewRequirement(common.LabelResourceEmail, selection.Equals, []string{user.Email})
	if err != nil {
		logger.Fatal("Failed to create requirement to list Resources for user", zap.Error(err))
	}
	sel := labels.NewSelector().Add(*emailRequirement)

	resp, err := client.ListResources(ctx, &indentv1.ListResourcesRequest{
		SpaceName:     f.Config().Space,
		LabelSelector: sel.String(),
	})
	if err != nil {
		logger.Fatal("Failed to list Resources for user", zap.Error(err), zap.String("selector", sel.String()))
	} else if len(resp.Resources) < 1 {
		logger.Fatal("No resources returned for user", zap.Error(err), zap.String("selector", sel.String()))
	} else if len(resp.Resources) == 1 {
		return resp.Resources[0]
	}

	for _, kind := range kindOrder {
		for _, resource := range resp.Resources {
			if resource.Kind == kind {
				return resource
			}
		}
	}
	logger.Fatal("Multiple users returned, unclear which to use")
	return nil
}
