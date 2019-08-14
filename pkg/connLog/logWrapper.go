package connLog

import (
	"RestSQL/pkg/config"
	//"log/syslog"
	"net/http"
	"net/http/httptrace"
	"time"

	log "github.com/sirupsen/logrus"
	//lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

/*type transport struct {
	current *http.Request
}*/

func FuncTimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

/*func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req
	return http.DefaultTransport.RoundTrip(req)
}

func (t *transport) GetConn(info httptrace.GotConnInfo) {
	log.Printf("Connection reused for %v? %v\n", t.current.URL, info.Reused)
}

func ConnTimeGet(req *http.Request) {
	t := &transport{}

	trace := &httptrace.ClientTrace{
		GotConn: t.GetConn,
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	client := &http.Client{Transport: t}
	if _, err := client.Do(req); err != nil {
		log.Panic(err.Error())
	}
}*/

func ConnTimeGet(url string) {
	//make part of handler to make the actual calls
	/*log := logrus.New()
	hook, err := lSyslog.NewSyslogHook("", "", syslog.LOG_INFO, "")

	if err == nil {
		log.Hooks.Add(hook)
	}*/

	req, _ := http.NewRequest("GET", url, nil)
	//req, _ := http.NewRequest(httpMethod, url, postBody)

	var start, connect, dns time.Time

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			log.Printf("DNS Done: %v\n", time.Since(dns))
		},

		ConnectStart: func(network, addr string) { connect = time.Now() },
		ConnectDone: func(network, addr string, err error) {
			log.Printf("DB: %s; Connect time: %v\n", config.DbName, time.Since(connect))
		},

		GotFirstResponseByte: func() {
			log.Printf("DB: %s; Time from start to first byte: %v\n", config.DbName, time.Since(start))
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		log.Panic(err.Error())
	}
	log.Printf("Total time: %v\n", time.Since(start))
}
