package sshconfig

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
)

//
func MustNew(file string) Config {
	config, err := New(file)
	if err != nil {
		panic(err)
	}
	return config
}

//
func New(file string) (Config, error) {
	source, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(bytes.NewReader(source))
	lines := []string{}
	for {

		//

		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}

	config := &sshConfig{
		source:    lines,
		delimiter: Space,
	}
	return config.parse()
}
