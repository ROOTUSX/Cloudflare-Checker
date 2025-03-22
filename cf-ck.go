package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// ANSI Color Codes for Terminal Output
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Reset  = "\033[0m"
)

// Get Real IP Address of a Domain
func getRealIP(domain string) string {
	ips, err := net.LookupHost(domain)
	if err != nil {
		return Red + "[Real IP: Not Found]" + Reset
	}
	return Cyan + "[Real IP: " + strings.Join(ips, ", ") + "]" + Reset
}

// Check if a Website is Protected by Cloudflare
func isCloudflare(domain string) (bool, int, time.Duration, string, string) {
	url := domain
	if !strings.HasPrefix(domain, "http") {
		url = "https://" + domain
	}

	client := http.Client{Timeout: 7 * time.Second}

	start := time.Now()
	resp, err := client.Get(url)
	responseTime := time.Since(start)

	// If request fails, return error info
	if err != nil {
		return false, 0, responseTime, "", err.Error()
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := strings.ToLower(string(bodyBytes))

	// Headers indicating Cloudflare protection
	serverHeader := strings.ToLower(resp.Header.Get("Server"))
	cfRay := resp.Header.Get("CF-RAY")
	cfCacheStatus := resp.Header.Get("CF-Cache-Status")

	isCF := strings.Contains(serverHeader, "cloudflare") || cfRay != "" || cfCacheStatus != ""
	var additionalInfo string

	// Detect Cloudflare security challenges
	if strings.Contains(body, "checking your browser") {
		isCF = true
		additionalInfo = "(Cloudflare JS Challenge)"
	} else if strings.Contains(body, "one more step") {
		isCF = true
		additionalInfo = "(Cloudflare CAPTCHA Page)"
	}

	return isCF, resp.StatusCode, responseTime, fmt.Sprintf("CF-RAY: %s, CF-Cache: %s", cfRay, cfCacheStatus), additionalInfo
}

// Colorize Status Code for Output
func colorizeStatusCode(status int) string {
	if status >= 500 {
		return Red + fmt.Sprintf("[Status: %d]", status) + Reset
	} else if status >= 400 {
		return Yellow + fmt.Sprintf("[Status: %d]", status) + Reset
	} else {
		return Green + fmt.Sprintf("[Status: %d]", status) + Reset
	}
}

func main() {
	// Parse command-line flags
	outputFile := flag.String("o", "", "Specify output file to save results")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println(Red + "Usage: ./cf-ck websites.txt [-o output.txt]" + Reset)
		return
	}

	inputFile := flag.Arg(0)
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(Red + "Error opening file:", err, Reset)
		return
	}
	defer file.Close()

	// Read websites from the input file
	var websites []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		website := strings.TrimSpace(scanner.Text())
		if website != "" {
			websites = append(websites, website)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(Red + "Error reading file:", err, Reset)
		return
	}

	// Open output file if specified
	var output *os.File
	if *outputFile != "" {
		output, err = os.Create(*outputFile)
		if err != nil {
			fmt.Println(Red + "Error creating output file:", err, Reset)
			return
		}
		defer output.Close()
	}

	// Start checking websites
	fmt.Println(Blue + "🔵Cloudflare Checker - Created by RootUserX" + Reset)
	
	fmt.Println(Blue + "🔵 Starting Cloudflare Protection Check...\n" + Reset)

	var wg sync.WaitGroup
	for _, website := range websites {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()

			// Perform Cloudflare Check
			isCF, statusCode, responseTime, headers, additionalInfo := isCloudflare(w)
			realIP := getRealIP(w)
			statusColored := colorizeStatusCode(statusCode)

			// Format additional info
			if additionalInfo != "" {
				additionalInfo = " " + additionalInfo
			}

			// Generate output message
			var result string
			if isCF {
				result = fmt.Sprintf(Green+"✅ %s → Protected by Cloudflare %s "+Blue+"(Time: %v) %s\n%s %s"+Reset,
					w, statusColored, responseTime, additionalInfo, headers, realIP)
			} else if statusCode == 0 {
				result = fmt.Sprintf(Yellow+"⚠ %s → [Error] %s"+Reset, w, headers)
			} else {
				result = fmt.Sprintf(Red+"❌ %s → NOT protected by Cloudflare %s "+Blue+"(Time: %v) %s"+Reset,
					w, statusColored, responseTime, realIP)
			}

			// Print output to console
			fmt.Println(result)

			// Save output to file if specified
			if *outputFile != "" {
				cleanResult := strings.ReplaceAll(result, "\033[", "")
				cleanResult = strings.Split(cleanResult, "m")[1]
				output.WriteString(cleanResult + "\n")
			}

		}(website)
	}

	wg.Wait()

	fmt.Println(Blue + "\n✅ Scan Complete!" + Reset)
}
