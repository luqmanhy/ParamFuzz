<h1 align="center">ParamFuzz</h1>

<p align="center">
A fast, stdin-friendly URL parameter fuzzing tool in Go with safe query value replacement.
</p>

<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-red.svg"></a>
<a href="https://goreportcard.com/badge/github.com/luqmanhy/paramfuzz"><img src="https://goreportcard.com/badge/github.com/luqmanhy/paramfuzz"></a>
<a href="https://github.com/luqmanhy/paramfuzz/releases"><img src="https://img.shields.io/github/release/luqmanhy/paramfuzz"></a>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#credits">Credits</a> 
</p>

<p align="center">
<a href="https://github.com/luqmanhy/paramfuzz"><img src="/static/paramfuzz-demo.png" alt="ParamFuzz Demo"></a>
</p>

---

## Overview

ParamFuzz is a command-line tool written in Go that safely and efficiently replaces the values of all query parameters in URLs with a provided string (defaulting to "FUZZ"). It reads URLs from standard input and outputs the results to standard output.


## Features
- **Flexible Fuzzing**:
  - `-only`: fuzz specific parameters.
  - `-skip`: bypass certain parameters.
- **Quiet by Default**: Runs via stdin → stdout, making it easy to chain with other tools.
- **Cross-Platform**: Runs smoothly on Windows, Linux, and macOS.

## Installation
### Binaries
You can download a pre-built binary from [here](https://github.com/luqmanhy/paramfuzz/releases) and use it right away.

### Go
```sh
go install -v github.com/luqmanhy/paramfuzz/cmd/paramfuzz@latest
````

## Usage
Refer to the tool’s help output using:

```bash
paramfuzz -h
```

Here's what the help message looks like:

```
ParamFuzz - URL parameter fuzzing tool
Version: 0.0.1

Usage:
  cat urls.txt | paramfuzz [OPTIONS]

Options:
  -payload string   String to replace parameter values (default: FUZZ)
  -only string      Comma-separated list of parameters to fuzz
  -skip string      Comma-separated list of parameters to skip
  -help, -h         Show this help message

Examples:
  # Fuzz all parameters with default "FUZZ"
  echo "http://test.com/page?id=1&user=admin" | paramfuzz

  # Fuzz with custom payload
  echo "http://test.com/page?id=1&user=admin" | paramfuzz -payload "XSS"

  # Fuzz only 'id' parameter
  echo "http://test.com/page?id=1&user=admin" | paramfuzz -only id -payload "TEST"

  # Fuzz all except 'user' parameter
  echo "http://test.com/page?id=1&user=admin" | paramfuzz -skip user -payload "XSS"
```

## Examples
```
# Default fuzzing with payload "FUZZ"
echo "http://example.com?page=1&user=admin" | paramfuzz

# Custom payload
echo "http://example.com?page=1&user=admin" | paramfuzz -payload "XSS"

# Only fuzz `page`
echo "http://example.com?page=1&user=admin" | paramfuzz -only page -payload "TEST"

# Fuzz everything except `user`
echo "http://example.com?page=1&user=admin" | paramfuzz -skip user -payload "FUZZ"
```
## Credits
### Contributing

We welcome contributions! Feel free to submit [Pull Requests](https://github.com/luqmanhy/paramfuzz/pulls) or report [Issues](https://github.com/luqmanhy/paramfuzz/issues).

### Licensing

This utility is licensed under the [MIT license](https://opensource.org/license/mit). You are free to use, modify, and distribute it, as long as you follow the terms of the license. You can find the full license text in the repository - [Full MIT license text](https://github.com/luqmanhy/paramfuzz/blob/master/LICENSE).
