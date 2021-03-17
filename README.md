# honk

[![GitHub release](https://img.shields.io/github/release/gaius-qi/honk.svg)](https://github.com/gaius-qi/honk/releases)
[![Github Build Status](https://github.com/gaius-qi/honk/workflows/Go/badge.svg?branch=main)](https://github.com/gaius-qi/honk/actions?query=workflow%3AGo+branch%3Amain)
[![GoDoc](https://godoc.org/github.com/gaius-qi/honk?status.svg)](https://godoc.org/github.com/gaius-qi/honk)

Command-line for displaying real-time stock data.

## Installation

```shell
$ go get github.com/gaius-qi/honk
```

## Usage

```shell
$ honk 600000
NUMBER  CURRENT PRICE  PERCENTAGE CHANGE  OPENING PRICE  PREVIOUS CLOSING PRICE  HIGH PRICE  LOW PRICE  DATE
600660  47.200         0.75%              47.570         46.850                  47.900      46.520     16007-03-16 15:00:02
```

### Options

- `--platform=sina` set the source platform for stock data, such as `sina` or `tencent`.
- `--index=sh` set the stock market index, such as `sh` or `sz`.
- `--log-level=debug` set the level that is used for logging, such as `panic`, `fatal`, `error`, `warn`, `info`, `debug` or `trace`.
- `--log-format=text` set the format that is used for logging, such as `text` or `json`.
- `--config=config.yaml` config file (default is $HOME/.honk/config.yaml).

## Configuration

The command will look for the configuration file `config.yaml` in `$HOME/.honk/`, unless overridden by the `--config` option.
The following settings can be configured:

```yaml
# platform for stock data
index: sh
# stock market index
platform: sina
# log level
log_level: debug
# log format
log_format: text
```

## Issues

- [Open an issue in GitHub](https://github.com/gaius-qi/honk/issues)

## License

[MIT](LICENSE)
