package udpService

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/Aicrosoft/goBark/internal/handler"
	"github.com/Aicrosoft/goBark/internal/setting"
	"github.com/Aicrosoft/goBark/pkg/lib"
)

// UDPServer struct.
type UDPServer struct {
	Configuration *setting.AppSetting
}

// inti the UdpServer struct.
func (server *UDPServer) Init(conf *setting.AppSetting) {
	server.Configuration = conf
}

// Start the udpServer.
func (server *UDPServer) Start() error {
	log.Debug("UDPServer Start.")
	config := server.Configuration
	ip := net.ParseIP(config.UDPServer.Host)
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: ip, Port: config.UDPServer.Port})
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Info(fmt.Sprintf("listening on addr=%s with blockSize=%d", listener.LocalAddr(), config.UDPServer.BlockSize))

	data := make([]byte, config.UDPServer.BlockSize)
	event := handler.EventHandler{}
	event.Init(config)
	for {
		n, remoteAddr, err := listener.ReadFrom(data)
		if err != nil {
			log.Error(fmt.Printf("Read error:%s", err))
		}

		log.Debug(fmt.Sprintf("<%s>:%s <- %v", remoteAddr, data[:n], data[:n]))

		if config.DisableCaptureMessage {
			log.Info(fmt.Sprintf("[%s] %s", remoteAddr, data[:n]))
		} else {
			capEvent := event.Recive(string(data[:n]))
			if capEvent != nil {
				log.Info(fmt.Sprintf("Capture Event Result:%+v", capEvent))
			}

			// execute webhook when it is enabled
			if config.Webhook.Enabled {
				if err := lib.GetWebhook(config).Execute(capEvent); err != nil {
					return err
				}
			}

		}

	}
}
