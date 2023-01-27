package petitioncfg

import (
	"strings"

	"go.indent.com/indent-go/pkg/common"
)

const (
	// LabelApp determines the app type for a petition.
	LabelApp = "indent.com/app"
)

// ConfigLabels creates labels for an access-manager app config with a
// config ID comprising idParts.
func ConfigLabels(idParts ...string) (cfgLabels map[string]string) {
	cfgLabels = map[string]string{
		LabelApp:          "access-manager",
		common.IndentKind: common.LabelKindPetitionAppValue,
	}
	// only add config ID selector if string is not empty
	if len(idParts) != 0 && strings.Join(idParts, "") != "" {
		cfgLabels[common.LabelAppConfigID] = Combine("-", idParts...)
	}
	return
}

// Combine acts like strings.Join except skips empty entries.
func Combine(sep string, a ...string) (config string) {
	for _, part := range a {
		if part == "" {
			continue
		} else if config != "" {
			config += sep
		}
		config += part
	}
	return
}
