package rfc3164_test

import "testing"

import "github.com/influxdata/go-syslog/rfc3164"

func TestParse(t *testing.T) {
	p := rfc3164.NewParser()

	m, err := p.Parse([]byte(`<34>Jan 12 06:30:00 xxx apache_server: 1.2.3.4 - - [12/Jan/2011:06:29:59 +0100] "GET /foo/bar.html HTTP/1.1" 301 96 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.9.2.12) Gecko/20101026 Firefox/3.6.12 ( .NET CLR 3.5.30729)" PID 18904 Time Taken 0`))

	if err != nil {
		t.Error("Expected error to be nil:", err)
	}

	t.Errorf("%+v", m)
	t.Errorf("Hostname: %s", *m.Hostname())
	t.Errorf("Message: %s", *m.Message())
}
