package sshconfig

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

//
type Hosts []host

//
func (h Hosts) String() string {
	y, _ := yaml.Marshal(h)
	return string(y)
}

//
type host struct {
	names   []string
	line    int
	options options
}

//
func (h host) Marshal(delimiter delimiter) interface{} {
	host := []string{"Host" + string(delimiter) + strings.Join(h.names, " ")}
	return append(host, h.options.slice(delimiter)...)
}

//
type options map[string]option

//
func (o options) slice(delimiter delimiter) []string {
	os := make([]string, len(o))
	for k, v := range o {
		os[v.order-1] = fmt.Sprintf("%s%c%s", k, delimiter, v.value)
	}
	return os
}

//
type option struct {
	order int
	line  int
	value string
}
