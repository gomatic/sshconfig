package sshconfig

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

//
type Config interface {
	Find(Query) Config
}

//
type delimiter rune

//
const (
	Equals = delimiter('=')
	Space  = delimiter(' ')
)

//
type sshConfig struct {
	source    []string
	options   options
	hosts     Hosts
	delimiter delimiter
}

//
func (c *sshConfig) Generic() interface{} {
	i := []interface{}{c.options.slice(c.delimiter)}
	for _, h := range c.hosts {
		i = append(i, h.Marshal(c.delimiter))
	}
	return i
}

//
func (c *sshConfig) MarshalYAML() (interface{}, error) {
	return c.Generic(), nil
}

//
func (c *sshConfig) MarshalJSON() (interface{}, error) {
	return c.Generic(), nil
}

//
func (c *sshConfig) MarshalConfig() []string {
	i := []string{}
	for _, o := range c.options.slice(c.delimiter) {
		i = append(i, o)
	}
	for _, h := range c.hosts {
		i = append(i, "Host"+string(c.delimiter)+strings.Join(h.names, " "))
		for _, o := range h.options.slice(c.delimiter) {
			i = append(i, "\t"+o)
		}
	}
	return i
}

//
func (c *sshConfig) String() string {
	b := bytes.Buffer{}
	for _, l := range c.MarshalConfig() {
		fmt.Fprintln(&b, l)
	}
	return b.String()
}

//
func (c *sshConfig) parse() (*sshConfig, error) {

	ws := regexp.MustCompile(`\\s+`)
	kv := regexp.MustCompile(` += +| +=|= +| +|=`)

	var h host
	c.hosts = Hosts{}
	c.options = map[string]option{}

	for n, line := range c.source {

		//

		line = strings.TrimLeft(line, " \t")
		line = strings.TrimRight(line, " \t\n")
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		//

		line = ws.ReplaceAllString(line, " ")
		ps := kv.Split(line, 2)
		if len(ps) != 2 {
			return nil, fmt.Errorf("ERROR: Invalid key/value: line %d: %s", n, line)
		}
		switch k, v := ps[0], ps[1]; strings.ToLower(k) {
		case "host":
			h = host{
				names:   strings.Split(v, " "),
				line:    n,
				options: options{},
			}
			c.hosts = append(c.hosts, h)
		default:
			if h.options == nil {
				c.options[k] = option{
					order: len(h.options) + 1,
					line:  n,
					value: v,
				}
			} else {
				h.options[k] = option{
					order: len(h.options) + 1,
					line:  n,
					value: v,
				}
			}
		}
	}
	return c, nil
}

//
func (c *sshConfig) Find(q Query) Config {
	f := &sshConfig{
		source:    c.source,
		delimiter: c.delimiter,
		options:   options{},
		hosts:     Hosts{},
	}
	for _, p := range q {
		for _, h := range c.hosts {
			for _, n := range h.names {
				if p.key == "host" && strings.Contains(n, p.value) {
					f.hosts = append(f.hosts, h)
					break
				}
			}
		}
	}
	return f
}
