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
      "content": "current WAN IPv4 is : $ipv4",
      "value": "$ipv4"
    },
    {
      "captureReg": "(?P<v>.+)",
      "content": "capturing any message: $v",
      "value": "$v"
    }
  ]
}
```
</details>

---

## Special Thanks

