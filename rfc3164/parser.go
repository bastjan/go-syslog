package rfc3164

import (
	"sync"

	syslog "github.com/influxdata/go-syslog"
)

// parser represent a RFC3162 parser with mutex capabilities.
type parser struct {
	sync.Mutex
	*machine

	bestEffort bool
}

// NewParser creates a syslog.Machine that parses RFC3162 syslog messages.
func NewParser(options ...syslog.MachineOption) *parser {
	p := &parser{
		machine: NewMachine(),
	}

	return p
}

// HasBestEffort tells whether the receiving parser has best effort mode on or off.
func (p *parser) HasBestEffort() bool {
	return p.bestEffort
}

// Parse parses the input RFC3162 syslog message using its FSM.
//
// Best effort mode enables the partial parsing.
func (p *parser) Parse(input []byte) (syslog.Message, error) {
	p.Lock()
	defer p.Unlock()

	msg, err := p.machine.Parse(input, &p.bestEffort)

	return msg, err
}
