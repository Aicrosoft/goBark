```text
            ______              _     
            (____  \            | |    
  ____  ___  ____)  )_____  ____| |  _ 
 / _  |/ _ \|  __  ((____ |/ ___) |_/ )
( (_| | |_| | |__)  ) ___ | |   |  _ ( 
 \___ |\___/|______/\_____|_|   |_| \_)
(_____|                                
 ```

[![MIT licensed][1]][2] 
[![Docker Image][3]][4] 
[![Go Report Card][5]][6] 
[![GoDoc][9]][10]

[1]: https://img.shields.io/badge/license-MIT-blue.svg
[2]: LICENSE
[3]: https://img.shields.io/docker/image-size/aicrosoft/gobark/latest
[4]: https://hub.docker.com/r/aicrosoft/gobark
[5]: https://goreportcard.com/badge/github.com/Aicrosoft/goBark
[6]: https://goreportcard.com/report/github.com/Aicrosoft/goBark
[9]: https://godoc.org/github.com/Aicrosoft/goBark?status.svg
[10]: https://godoc.org/github.com/Aicrosoft/goBark


[goBark](https://github.com/Aicrosoft/goBark) is a analyses udp message by regex tool . You can receive messages from the router and then define specific messages as events for subsequent actions.

---
- [Supported Platforms](#supported-platforms)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
  - [Overview](#overview)
  - [Configuration examples](#configuration-examples)
    - [Default](#default)
- [Running goBark](#running-gobark)
  - [Manually](#manually)
  - [As a manual daemon](#as-a-manual-daemon)
  - [As a managed daemon (with upstart)](#as-a-managed-daemon-with-upstart)
  - [As a managed daemon (with systemd)](#as-a-managed-daemon-with-systemd)
  - [As a Docker container](#as-a-docker-container)
  - [As a Windows service](#as-a-windows-service)
- [Special Thanks](#special-thanks)    

## Supported Platforms

* Linux
* MacOS
* ARM Linux (Raspberry Pi, etc.)
* Windows
* MIPS32 platform

  To compile binaries for MIPS (mips or mipsle), run:

  ```bash
  GOOS=linux GOARCH=mips/mipsle GOMIPS=softfloat go build -a
  ```

  The binary can run on routers as well.


## Installation

Build goBark by running (from the root of the repository):

```bash
# get dependencies
go mod download     
# build
go build cmd\goBark.go

# Or build for debug 
go build -ldflags="-X 'main.Version=0.1.3' -X 'main.DebugMode=true'" cmd\goBark.go
```

You can also download a compiled binary from the [releases](https://github.com/Aicrosoft/goBark/releases).


## Usage

Print usage/help by running:

```bash
$ ./gobark -h
Usage of ./gobark:
  -c string
        Specify a config file (default "./config.json")
  -h    Show help
```
## Configuration

### Overview 
 overview 
 
* Make a copy of [config_sample.json](configs/config_sample.json) and name it as `config.json`.
* Configure your  credentials, etc.
* Configure a notification medium (e.g. SMTP to receive emails) to get notified when your receive sepcial messages.
* Place the file in the same directory of goBark or use the `-c=path/to/your/config.json` option .

### Configuration file format

goBark supports 1 configuration file formats:

* JSON


### Configuration properties
* `disable_capture_message` - Just use the goBark as UdpServer when true.
* `debug_info` - Whether to output debugging information.
* `udpServer` - Set the UdpServer information
* `event_message_ignore_keys` - Ignore those keys message.
* `event_messages` - Capture the udp first expression caputred of array.

### Configuration examples

#### Default
This is a simple default configuration.
<details>
<summary>Example</summary>

```json
{
  "debug_info": true,
  "udpServer": {
    "host": "0.0.0.0",
    "port": 996,
    "blockSize": 1024
  },
  "disable_capture_message": false,
  "event_message_ignore_keys": ["dnsmasq-dhcp"],
  "event_messages": [
    {
      "captureReg": "local  IP address (?P<ipv4>((25[0-5]|2[0-4]\\d|[01]?\\d\\d?)\\.){3}(25[0-5]|2[0-4]\\d|[01]?\\d\\d?))",
      "title": "PPoE dial success",
      "content": "current WAN IPv4 is : $ipv4",
      "value": "$ipv4"
    }
  ],
  "use_proxy": false,
  "socks5_proxy": "",
  "webhook": {
    "enabled": false,
    "url": "https://api.day.app/{change to your bark token}/{{.Title}}/{{.Content}}",
    "request_body": ""
  }
}
```
</details>

---


## Running goBark

There are a few ways to run goBark.

### Manually

Note: make sure to set the `run_once` parameter in your config file so the program will quit after the first run (the default is `false`).

Can be added to `cron` or attached to other events on your system.

```json
{
  "...": "...",
  "run_once": true
}
```
Then run

```bash
./gobark
```

### As a manual daemon

```bash
nohup ./gobark &
```

Note: when the program stops, it will not be restarted.

### As a managed daemon (with upstart)

1. Install `upstart` first (if not available already)
2. Copy `./upstart/gobark.conf` to `/etc/init` (and tweak it to your needs)
3. Start the service:

   ```bash
   sudo start gobark
   ```

### As a managed daemon (with systemd)

1. Install `systemd` first (it not available already)
2. Copy `./systemd/gobark.service` to `/lib/systemd/system` (and tweak it to your needs)
3. Start the service:

   ```bash
   sudo systemctl enable gobark
   sudo systemctl start gobark
   ```

### As a Docker container

Available docker registries:
* https://hub.docker.com/repository/docker/aicrosoft/gobark#

Visit https://hub.docker.com/repository/docker/aicrosoft/gobark# to fetch the latest docker image.  
With `/path/to/config.json` your local configuration file, run:

```bash
docker run \
-p 996:996/udp \
-d --name gobark --restart=always \
-v /path/to/config.json:/config.json \
aicrosoft/gobark:latest
```

### As a Windows service

1. Download the latest version of [NSSM](https://nssm.cc/download)

2. In an administrative prompt, from the folder where NSSM was downloaded, e.g. `C:\Downloads\nssm\` **win64**, run:

   ```
   nssm install YOURSERVICENAME
   ```

3. Follow the interface to configure the service. In the "Application" tab just indicate where the `gobark.exe` file is. Optionally you can also define a description on the "Details" tab and define a log file on the "I/O" tab. Finish by clicking on the "Install service" button.

4. The service will now start along Windows.

Note: you can uninstall the service by running:

```
nssm remove YOURSERVICENAME
```


## Special Thanks

[gobark][110] : good to learn golang example.

[110]: https://github.com/TimothyYe/gobark

