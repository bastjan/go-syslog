package rfc3164

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	syslogtesting "github.com/influxdata/go-syslog/testing"
)

type test struct {
	line string

	expectedMessage *SyslogMessage
	expectedError   error
}

var tests = []test{
	test{
		line: `<34>Jan 12 06:30:00 xxx apache_server: 1.2.3.4 - - [12/Jan/2011:06:29:59 +0100] "GET /foo/bar.html HTTP/1.1" 301 96 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.9.2.12) Gecko/20101026 Firefox/3.6.12 ( .NET CLR 3.5.30729)" PID 18904 Time Taken 0`,
		expectedMessage: &SyslogMessage{
			priority:  syslogtesting.Uint8Address(34),
			facility:  syslogtesting.Uint8Address(4),
			severity:  syslogtesting.Uint8Address(2),
			timestamp: syslogtesting.TimeParse(time.Stamp, `Jan 12 06:30:00`),
			hostname:  syslogtesting.StringAddress("xxx"),

			message: syslogtesting.StringAddress(`1.2.3.4 - - [12/Jan/2011:06:29:59 +0100] "GET /foo/bar.html HTTP/1.1" 301 96 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.9.2.12) Gecko/20101026 Firefox/3.6.12 ( .NET CLR 3.5.30729)" PID 18904 Time Taken 0`),
			appname: syslogtesting.StringAddress("apache_server"),
			procID:  nil,
		}},
	test{
		line: `<34>Jan 12 06:30:00 xxx apache_server[12343]: 1.2.3.4 - - [12/Jan/2011:06:29:59 +0100] "GET /foo/bar.html HTTP/1.1" 301 96 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.9.2.12) Gecko/20101026 Firefox/3.6.12 ( .NET CLR 3.5.30729)" PID 18904 Time Taken 0`,
		expectedMessage: &SyslogMessage{
			priority:  syslogtesting.Uint8Address(34),
			facility:  syslogtesting.Uint8Address(4),
			severity:  syslogtesting.Uint8Address(2),
			timestamp: syslogtesting.TimeParse(time.Stamp, `Jan 12 06:30:00`),
			hostname:  syslogtesting.StringAddress("xxx"),

			message: syslogtesting.StringAddress(`1.2.3.4 - - [12/Jan/2011:06:29:59 +0100] "GET /foo/bar.html HTTP/1.1" 301 96 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.9.2.12) Gecko/20101026 Firefox/3.6.12 ( .NET CLR 3.5.30729)" PID 18904 Time Taken 0`),
			appname: syslogtesting.StringAddress("apache_server"),
			procID:  syslogtesting.StringAddress("12343"),
		}},
	test{
		line: `<34>Jan 12 06:30:00 xxx 1.2.3.4 - - [12/Jan/2011:06:29:59 +0100] "GET /foo/bar.html HTTP/1.1" 301 96 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.9.2.12) Gecko/20101026 Firefox/3.6.12 ( .NET CLR 3.5.30729)" PID 18904 Time Taken 0`,
		expectedMessage: &SyslogMessage{
			priority: syslogtesting.Uint8Address(34),
			facility: syslogtesting.Uint8Address(4),
			severity: syslogtesting.Uint8Address(2),
			hostname: syslogtesting.StringAddress("xxx"),

			timestamp: syslogtesting.TimeParse(time.Stamp, `Jan 12 06:30:00`),
			message:   syslogtesting.StringAddress(`1.2.3.4 - - [12/Jan/2011:06:29:59 +0100] "GET /foo/bar.html HTTP/1.1" 301 96 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.9.2.12) Gecko/20101026 Firefox/3.6.12 ( .NET CLR 3.5.30729)" PID 18904 Time Taken 0`),
			appname:   nil,
			procID:    nil,
		}},
}

func TestParse(t *testing.T) {
	p := NewParser()

	for _, test := range tests {
		parsed, err := p.Parse([]byte(test.line))

		assert.Equal(t, test.expectedError, err)
		assert.Equal(t, test.expectedMessage, parsed)
	}
}
