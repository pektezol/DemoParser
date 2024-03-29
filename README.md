# DemoParser [![Go Report Card](https://goreportcard.com/badge/github.com/pektezol/demoparser)](https://goreportcard.com/report/github.com/pektezol/demoparser) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/pektezol/DemoParser/blob/main/LICENSE)

Source Demo Parser for Portal 2 written in Golang.

## Couldn't Do This Without Them

- [@UncraftedName](https://github.com/UncraftedName): For [UntitledParser](https://github.com/UncraftedName/UntitledParser)
- [@NeKzor](https://github.com/NeKzor): For [nekz.me/dem](https://nekz.me/dem)

## Usage

```bash
$ ./parser demo.dem

$ ./parser demos/
```

## Currently Supports

- File or folder input using the CLI.
- Parsing of demo headers and each (most) message type.
- Parsing of packet classes.
- Custom injected SAR data parsing.

## TODO

- In-depth packet class parsing for each class. ([#7][i7])
- Entity parsing. ([#17][i17])

[i7]: https://github.com/pektezol/DemoParser/issues/7
[i17]: https://github.com/pektezol/DemoParser/issues/17