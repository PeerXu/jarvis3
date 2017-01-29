package executor

import "strings"

type Data interface {
	Set(val string)
	Get() string
}

type data struct {
	val string
}

func NewData() Data {
	return &data{}
}

func (d *data) Set(val string) {
	d.val = val
}

func (d *data) Get() string {
	return d.val
}

type Metadata interface {
	Set(key, val string)
	Get(key string) string
}

type metadata struct {
	vals map[string]string
}

func NewMetadata() Metadata {
	return &metadata{vals: make(map[string]string)}
}

func fmtMetaKey(s string) string {
	return strings.Replace(strings.ToUpper(s), "_", "-", -1)
}

func (md *metadata) Set(key, val string) {
	md.vals[fmtMetaKey(key)] = val
}

func (md *metadata) Get(key string) string {
	if val, ok := md.vals[fmtMetaKey(key)]; ok {
		return val
	}
	return ""
}
