package setting

import (
	"testing"
)

func TestLoadJSONSetting(t *testing.T) {
	var settings AppSetting
	err := LoadSetting("../../config/config_sample.json", &settings)

	regStr := settings.EventMessages[0].CaptureReg
	//golang can use `` include a block strings, but in json doc must use "" .
	if regStr != `(?P<name>[a-zA-Z]+)\s+(?P<age>\d+)\s+(?P<email>\w+@\w+(?:\.\w+)+)` {
		t.Fatal("can't read regstring correctly")
	}

	if err != nil {
		t.Fatal(err.Error())
	}

	if settings.UDPServer.Port == 0 || len(settings.UDPServer.Host) < 7 || settings.UDPServer.BlockSize == 0 {
		t.Fatal("UdpServer setting error,please check it first.")
	}

	err = LoadSetting("./file/does/not/exists", &settings)
	if err == nil {
		t.Fatal("file doesn't exist, should return error")
	}
}
