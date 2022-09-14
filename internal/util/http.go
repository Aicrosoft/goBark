package util

import (
	"context"
	"net"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Aicrosoft/goBark/internal/setting"
	"golang.org/x/net/proxy"
)

// GetHTTPClient creates the HTTP client and return it.
func GetHTTPClient(conf *setting.AppSetting) *http.Client {
	client := &http.Client{
		Timeout: time.Second * defaultTimeout,
	}

	if conf.UseProxy && conf.Socks5Proxy != "" {
		log.Debug("use socks5 proxy:" + conf.Socks5Proxy)
		dialer, err := proxy.SOCKS5("tcp", conf.Socks5Proxy, nil, proxy.Direct)
		if err != nil {
			log.Error("can't connect to the proxy:", err)
			return nil
		}

		dialContext := func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.Dial(network, address)
		}

		httpTransport := &http.Transport{}
		client.Transport = httpTransport
		httpTransport.DialContext = dialContext
	}

	return client
}
