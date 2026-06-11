package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Timeout: time.Duration(80) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				},
				PreferServerCipherSuites: true,
				MinVersion:               tls.VersionTLS12,
				MaxVersion:               tls.VersionTLS13,
			},
			DialContext: (&net.Dialer{
				Timeout: time.Duration(80) * time.Second,
			}).DialContext,
		},
	}
}

// NewRequest create new request instance
/*func NewRequest() *Request {
	request := &Request{}
	request.Headers = make(map[string]string)
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Duration(80) * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					CipherSuites: []uint16{
						tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
						tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					},
					PreferServerCipherSuites: true,
					MinVersion:               tls.VersionTLS12,
					MaxVersion:               tls.VersionTLS13,
				},
				DialContext: (&net.Dialer{
					Timeout: time.Duration(80) * time.Second,
				}).DialContext,
			},
		}
	}
	request.Client = httpClient
	request.Logger = log.(os.Stdout)
	request.Logger = log.With(request.Logger,
		"ts", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
	return request
}
*/
