package setting

import (
	"testing"
)

func TestLoadJSONSetting(t *testing.T) {
	var settings AppSetting
	err := LoadSetting("../../config/config_sample.json", &settings)

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
