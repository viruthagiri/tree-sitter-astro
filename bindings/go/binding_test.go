package tree_sitter_astro_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_astro "github.com/tree-sitter/tree-sitter-astro/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_astro.Language())
	if language == nil {
		t.Errorf("Error loading astro grammar")
	}
}
