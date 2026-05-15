package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// I18n holds flattened listmonk locale strings (key -> translated).
type I18n struct {
	m map[string]string
}

// LoadI18n reads a listmonk-style JSON locale file.
func LoadI18n(path string) (*I18n, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m map[string]string
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return &I18n{m: m}, nil
}

// T returns the translation for key, or key if missing.
func (i *I18n) T(key string) string {
	if i == nil || i.m == nil {
		return key
	}
	if s, ok := i.m[key]; ok {
		return s
	}
	return key
}

// Ts returns T(key) with {placeholder} replaced from variadic pairs (name, value, ...).
func (i *I18n) Ts(key string, args ...any) string {
	s := i.T(key)
	for j := 0; j+1 < len(args); j += 2 {
		ph := fmt.Sprint(args[j])
		val := fmt.Sprint(args[j+1])
		s = strings.ReplaceAll(s, "{"+ph+"}", val)
	}
	return s
}
