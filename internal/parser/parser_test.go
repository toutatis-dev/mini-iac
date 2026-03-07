package parser

import (
	"mini-iac/internal/ast"
	"mini-iac/internal/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseManifest(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected ast.Manifest
	}{
		{"Valid Resource", "resource \"file\" \"test1.md\" {content = \"# My Project\";}", ast.Manifest{
			Blocks: []ast.Block{
				&ast.Resource{
					Provider:     "file",
					ResourceName: "test1.md",
					Properties:   map[string]string{"content": "# My Project"},
				},
			},
		}},
		{"Invalid Resource", "resource file \"test1.md\" {content = \"# My Project\";}", ast.Manifest{Blocks: []ast.Block{}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := string(tt.input)
			lex := lexer.New(input)
			p := New(lex)
			manifest := p.ParseManifest()

			assert.Equal(t, tt.expected, manifest)
		})
	}
}
