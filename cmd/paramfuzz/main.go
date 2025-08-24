package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const version = "0.0.1"

func showHelp() {
	fmt.Printf(`ParamFuzz - URL parameter fuzzing tool
Version: %s

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
`, version)
}

func main() {
	flag.Usage = func() {}

	// Define custom flags
	payload := flag.String("payload", "FUZZ", "String to replace parameter values")
	onlyParams := flag.String("only", "", "Comma-separated list of parameters to fuzz")
	skipParams := flag.String("skip", "", "Comma-separated list of parameters to skip")
	helpFlag := flag.Bool("help", false, "Show help message")
	hFlag := flag.Bool("h", false, "Show help message (alias of -help)")

	flag.Parse()

	// Handle help flags
	if *helpFlag || *hFlag {
		showHelp()
		os.Exit(0)
	}

	// Check if stdin has data (avoid freeze)
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) != 0 {
		// No stdin provided, show help
		showHelp()
		os.Exit(0)
	}

	// Encode payload for URLs
	encoded := url.QueryEscape(*payload)

	// Parse only/skip sets
	only := make(map[string]bool)
	skip := make(map[string]bool)

	if *onlyParams != "" {
		for _, p := range strings.Split(*onlyParams, ",") {
			only[strings.TrimSpace(p)] = true
		}
	}
	if *skipParams != "" {
		for _, p := range strings.Split(*skipParams, ",") {
			skip[strings.TrimSpace(p)] = true
		}
	}

	// Regex for query params
	re := regexp.MustCompile(`([?&])([^=]+)=([^&]*)`)

	// Read URLs from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		result := re.ReplaceAllStringFunc(line, func(match string) string {
			parts := re.FindStringSubmatch(match)
			if len(parts) != 4 {
				return match
			}
			sep, key := parts[1], parts[2]

			// Apply only rule
			if len(only) > 0 && !only[key] {
				return match
			}
			// Apply skip rule
			if skip[key] {
				return match
			}

			// Replace param value with payload
			return fmt.Sprintf("%s%s=%s", sep, key, encoded)
		})

		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		os.Exit(127)
	}
}
