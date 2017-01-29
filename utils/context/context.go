package context

import "strings"

type Metadata map[string]string

var fmtMdKey = func(s string) string {
	return strings.ToUpper(strings.Replace(s, "_", "-", -1))
}

func (m Metadata) Set(key, val string) {
	m[fmtMdKey(key)] = val
}

func (m Metadata) Get(key string) string {
	if val, ok := m[fmtMdKey(key)]; ok {
		return val
	}
	return ""
}

type Context interface {
	Metadata() Metadata
}

type context struct {
	md Metadata
}

func (ctx *context) Metadata() Metadata {
	return ctx.md
}

func NewContext() Context {
	return &context{
		md: Metadata{},
	}
}
