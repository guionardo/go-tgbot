# go-tgbot
Go Telegram bot framework

[![CodeQL](https://github.com/guionardo/go-tgbot/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/guionardo/go-tgbot/actions/workflows/codeql-analysis.yml)
[![Go](https://github.com/guionardo/go-tgbot/actions/workflows/go.yml/badge.svg)](https://github.com/guionardo/go-tgbot/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/guionardo/go-tgbot?style=flat-square)](https://goreportcard.com/report/github.com/guionardo/go-tgbot)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/guionardo/go-tgbot) 
[![Release](https://img.shields.io/github/release/guionardo/go-tgbot.svg?style=flat-square)](https://github.com/guionardo/go-tgbot/releases/latest)

## Features

* Listening to messages sent into chat -> event triggering
* Background worker for scheduled services
* Message persistence with expiration time

## Configuration

### Json file

``` json
{
    "bot": {
        "token": "1234567890:ABCDEFGHIJKLMNOPQRSTUVWXYZ012345678",
        "name": "tgbot_test",
        "hello_world": "Hello World!"
    },
    "repository": {
        "connection_string": "tgbot.db",
        "house_keeping_max_age": "24h"
    },
    "logging": {
        "level": "warn",
        "format_time_stamp": "",
        "log_format": ""
    }
}
```

### YAML file

```yaml
---
bot:
  token: 1234567890:ABCDEFGHIJKLMNOPQRSTUVWXYZ012345678
  name: tgbot_test
  hello_world: "Hello World!"
repository:
  connection_string: "tgbot.db"
  house_keeping_max_age: 24h
logging:
  level: warn
  format_time_stamp: ""
  log_format: ""
```

### Environment variables

```
TG_BOT_TOKEN=1234567890:ABCDEFGHIJKLMNOPQRSTUVWXYZ012345678
TG_LOG_LEVEL=debug
TG_REPOSITORY_CONNECTION_STRING=tgbot.db
TG_BOT_NAME='GO Bot'
TG_BOT_HOUSEKEEPING_MAX_AGE=6h
```