package cliutil

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/manifoldco/promptui"
	"golang.org/x/crypto/ssh/terminal"
	"google.golang.org/protobuf/reflect/protoreflect"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	indentv1 "go.indent.com/indent-go/api/indent/v1"
)

const (
	promptLogo = "{{ \"▶\" | blue }} "

	selectSizePadding  = 9
	selectSizeFallback = 15
)

// NewSelect is a helper for interactive selection from a list.
func NewSelect[A proto.Message](items []A) (*Select[A], error) {
	if len(items) == 0 {
		return nil, errors.New("items cannot be empty")
	}
	first := items[0]

	desc := proto.MessageReflect(first).Descriptor()
	return &Select[A]{
		Select: promptui.Select{
			Label:             desc.Name() + "s",
			Items:             items,
			Size:              selectHeight(),
			Templates:         newSelectTemplates(desc, first),
			Searcher:          newSearcher(items),
			StartInSearchMode: true,
		},
		items: items,
	}, nil
}

// Select manages the lifecycle of a select prompt.
type Select[A proto.Message] struct {
	Select promptui.Select
	items  []A
}

// Run the select prompt.
func (s *Select[A]) Run() (a A, err error) {
	i, _, err := s.Select.Run()
	if err != nil {
		return a, err
	}
	return s.items[i], nil
}

// newSelectTemplates generates the templates for the select prompt.
func newSelectTemplates(desc protoreflect.MessageDescriptor, msg proto.Message) *promptui.SelectTemplates {
	kind := strconv.Quote(string(desc.FullName()))
	name := ".Name"

	switch msg.(type) {
	case *auditv1.Resource:
		kind, name = ".Kind", ".DisplayName"
	case *indentv1.Petition:
		kind = ".State.Status.Phase"
	}

	promptName := formatPromptName(kind, name)
	details := fmt.Sprintf(`
--------- %s ----------
{{ "Kind:" | faint }}	{{ %s }}
{{ "Name:" | faint }}	{{ %s }}
`, desc.Name(), kind, name)
	return &promptui.SelectTemplates{
		Label:    fmt.Sprintf("{{ %s }}", name),
		Active:   promptLogo + promptName,
		Inactive: "  " + promptName,
		Selected: promptLogo + promptName,
		Details:  details,
	}
}

// selectHeight returns the height of the select prompt based on terminal size, falls back to fallbackSelectSize.
func selectHeight() int {
	if _, height, err := terminal.GetSize(0); err == nil {
		return height - selectSizePadding
	}
	return selectSizeFallback
}

// newSearcher generates the searcher for the select prompt.
func newSearcher[A proto.Message](items []A) func(input string, index int) bool {
	return func(input string, index int) bool {
		item := items[index]
		searchStr := newSearchStr(item)
		for _, s := range strings.Split(input, " ") {
			if !strings.Contains(searchStr, normalize(s)) {
				return false
			}
		}
		return true
	}
}

// interfaces allow searcher to be generic over different types of items.
type searchNameGetter interface {
	GetName() string
}
type searchKindGetter interface {
	GetKind() string
}
type searchIDGetter interface {
	GetId() string
}
type displayNameGetter interface {
	GetDisplayName() string
}

func newSearchStr(item proto.Message) (searchStr string) {
	if name, ok := item.(searchNameGetter); ok {
		searchStr += name.GetName()
	}
	if kind, ok := item.(searchKindGetter); ok {
		searchStr += kind.GetKind()
	}
	if id, ok := item.(searchIDGetter); ok {
		searchStr += id.GetId()
	}
	if displayName, ok := item.(displayNameGetter); ok {
		searchStr += displayName.GetDisplayName()
	}
	return normalize(searchStr)
}

// normalize returns a string with all spaces removed and all characters lowercase.
func normalize(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "")
}

// formatPromptName returns the name as it appears in the select prompt.
func formatPromptName(a, b string) string {
	return fmt.Sprintf("{{ %s | yellow }} {{ %s | cyan }}", a, b)
}
