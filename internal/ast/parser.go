package ast

import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
)

// Language represents a supported source language.
type Language string

const (
	LangGo   Language = "go"
	LangCPP  Language = "cpp"
	LangJava Language = "java"
	LangPHP  Language = "php"
)

// ParsedTree holds the tree-sitter result for a source file.
type ParsedTree struct {
	Tree     *sitter.Tree
	Source   []byte
	Language Language
}

// Parse parses source code bytes for the given language.
func Parse(source []byte, lang Language) (*ParsedTree, error) {
	sitterLang, err := resolveLanguage(lang)
	if err != nil {
		return nil, err
	}

	parser := sitter.NewParser()
	parser.SetLanguage(sitterLang)

	tree, err := parser.ParseCtx(context.Background(), nil, source)
	if err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	return &ParsedTree{
		Tree:     tree,
		Source:   source,
		Language: lang,
	}, nil
}

// resolveLanguage maps Language to the tree-sitter grammar.
// Grammars are registered here as they are added.
func resolveLanguage(lang Language) (*sitter.Language, error) {
	switch lang {
	default:
		return nil, fmt.Errorf("unsupported language: %s", lang)
	}
}
