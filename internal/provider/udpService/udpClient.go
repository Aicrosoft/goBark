package udpService

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/Aicrosoft/goBark/internal/setting"
)

// UDPServer struct.
type UDPClient struct {
	Configuration *setting.AppSetting
}

// inti the UdpServer struct.
func (server *UDPClient) Init(conf *setting.AppSetting) {
	server.Configuration = conf
}

// Start the udpServer.
func (server *UDPClient) Start() error {
	log.Debug("UDPClient Start.")
	config := server.Configuration
	ip := net.ParseIP(config.UDPClient.Host)
	port := config.UDPClient.Port
	srcAddr := &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: port}
	log.Debug(fmt.Sprintf("sending from %s to %s", srcAddr, dstAddr))
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	payload := "this is udp message."
	var r io.Reader = strings.NewReader(payload)
	if payload == "" {
		log.Error("Udp message must be not null.")
		r = os.Stdin
	}

	n, err := io.Copy(conn, r)
	if err != nil {
		log.Fatal(fmt.Sprintf("error sending data: %s", err))
	}

	log.Info(fmt.Sprintf("sent %d bytes", n))

	return nil
}
