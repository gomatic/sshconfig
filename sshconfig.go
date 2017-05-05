package sshconfig

import (
	"bufio"
	"io"
	"os"

	"github.com/prometheus/common/log"
)

//
func MustNew(file string) Config {
	config, err := New(file)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

//
func Read(source io.Reader) (Config, error) {
	r := bufio.NewReader(source)
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

//
func New(file string) (Config, error) {
	source, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer source.Close()
	return Read(source)
}
