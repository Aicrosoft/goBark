package handler

import (
	"testing"

	"github.com/Aicrosoft/goBark/internal/setting"
)

func TestOneEventCapture(t *testing.T) {
	msg := "Aug 27 08:32:04 dnsmasq-dhcp[25431]: DHCPACK(br0) 192.168.10.196 d8:ce:3a:8c:48:94 MI-9-SE"
	event := EventHandler{}
	event.Init(&setting.AppSetting{
		EventMessages: []setting.UDPEventSetting{{
			CaptureReg: `(?P<val>(.+dnsmasq-dhcp.+DHCPACK.+))`,
			MessageSetting: setting.MessageSetting{
				Title:   "Device online",
				Content: "$val",
			},
		}},
	})

	rst := event.Recive(msg)
	t.Logf("rst:%+v", rst)
}

func TestEventCapture(t *testing.T) {
	msgs := [...]string{"", "Aug 27 09:06:00 dnsmasq-dhcp[26877]: read /etc/ethers - 9 addresses", "Aug 27 07:42:14 dnsmasq[25075]: using 68616 more nameservers", "Aug 27 09:05:59 pppd[26882]: local  IP address 110.165.101.102"}
	var settings setting.AppSetting
	err := setting.LoadSetting("../../config/config_sample.json", &settings)
	if err != nil {
		t.Error("load config failed.")
	}

	event := EventHandler{}
	event.Init(&settings)
	for i, msg := range msgs {
		rst := event.Recive(msg)
		if i < 2 {
			t.Logf("Ignore:%v", msg)
		} else {
			t.Logf("Captured An Event:%+v\n", rst)
			if i == 3 && rst.Value == "110.165.101.102" {
				t.Log("IPV4 Pass.")
			}
		}

	}

}
