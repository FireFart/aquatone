package parsers

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/michenriksen/aquatone/core"
)

type CustomParser struct{}

func NewCustomParser() *CustomParser {
	return &CustomParser{}
}

func (p *CustomParser) Parse(r io.Reader) ([]string, error) {
	var targets []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		s := strings.Split(line, ":")
		if len(s) != 2 {
			return nil, fmt.Errorf("missing port")
		}
		host := s[0]
		ports := s[1]
		for _, portString := range strings.Split(ports, ",") {
			port, err := strconv.Atoi(portString)
			if err != nil {
				return nil, fmt.Errorf("invalid port %s", portString)
			}
			targets = append(targets, core.HostAndPortToURL(host, port, ""))
		}
	}
	return targets, nil
}
