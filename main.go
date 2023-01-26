package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Get domain name from command line argument
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a domain name as a command line argument")
		fmt.Println("Usage: ./domainAvailabilityChecker <domain_name> <optional_comma_separated_list_of_TLDs>")
		os.Exit(1)
	}

	// Get domain name from command line argument
	domain := os.Args[1]

	// Check if domain name is valid
	valid := isValidDomain(domain)
	if !valid {
		fmt.Println("Error: Invalid domain name.")
		os.Exit(1)
	}

	var tlds []string
	if len(os.Args) > 2 {
		tldArg := os.Args[2]
		// Check if TLD argument is valid
		validTLDs := isValidTLDs(tldArg)
		if !validTLDs {
			fmt.Println("Error: Invalid TLD argument. Please provide a comma-separated list of TLDs")
			os.Exit(1)
		}
		tlds = strings.Split(tldArg, ",")
	} else {
		tlds = getTLDs()
	}

	var registered []string

	fmt.Printf("Checking availability for domain '%s' with %d TLDs\n", domain, len(tlds))
	fmt.Println("-------------------------------------")
	fmt.Println("Available domains:")

	for _, tld := range tlds {
		// Check if domain is registered
		_, err := net.LookupNS(domain + "." + tld)
		if err == nil {
			registered = append(registered, tld)
		} else {
			if strings.Contains(err.Error(), "no such host") {
				fmt.Printf("\t%s.%s is available.\n", domain, tld)
			} else {
				fmt.Println(err)
			}
		}
	}
	if len(registered) == 0 {
		fmt.Println("No domain is registered.")
	} else {
		fmt.Println("-------------------------------------")
		fmt.Print("Registered domains:\n")
		for _, r := range registered {
			fmt.Printf("\t%s.%s is registered.\n", domain, r)
		}
	}
}

// Function to check if domain name is valid
func isValidDomain(domain string) bool {
	// Regular expression to match valid domain name (not include tld)
	r, _ := regexp.Compile("^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]).)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9-]*[A-Za-z0-9])$")
	return r.MatchString(domain)
}

// Function to check if string is comma-separated list of TLDs
func isValidTLDs(tlds string) bool {
	// Regular expression to match valid domain name (not include tld)
	r, _ := regexp.Compile("^([a-zA-Z0-9]+,)*[a-zA-Z0-9]+$")
	return r.MatchString(tlds)
}

// fetch the full list of all TLDs (top-level domains) on the internet
// https://data.iana.org/TLD/tlds-alpha-by-domain.txt  and return a slice
// of strings containing all the TLDs (without the dot)
func getTLDs() []string {
	// fetch the list of TLDs
	resp, err := http.Get("https://data.iana.org/TLD/tlds-alpha-by-domain.txt")
	if err != nil {
		fmt.Println("Error: Could not fetch the list of TLDs")
		os.Exit(1)
	}
	defer resp.Body.Close()
	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Could not read the response body")
		os.Exit(1)
	}

	// convert the response body to a string
	tldList := string(body)

	// remove all lines that start with a hash (#)
	r, _ := regexp.Compile("#.*")
	tldList = r.ReplaceAllString(tldList, "")

	// match and remove any line that starts with whitespace and a newline character,
	// or any line that is between two newline characters with whitespace in between
	r, _ = regexp.Compile("^\\s*\\n|\\n\\s*\\n")
	tldList = r.ReplaceAllString(tldList, "")

	// lowercase all TLDs
	tldList = strings.ToLower(tldList)

	// split each line of the string into a slice of TLDs
	r, _ = regexp.Compile("\\n")
	tlds := r.Split(tldList, -1)

	// remove the last element from the slice because it will be an empty string
	tlds = tlds[:len(tlds)-1]

	return tlds
}
