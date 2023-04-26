package dea

import (
	"net"
	"regexp"
	"strings"
)

var (
	disposableEmailPattern = regexp.MustCompile(`(?i)(test|temp|demo)[.-]?[\w\d]*@[A-Za-z\d\-.]+\.[A-Za-z]{2,}`)
	emailRegex             = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
)

func IsDisposableEmail(email string) (bool, error) {
	providers := NewEmailProviders()

	if !isEmail(email) {
		return true, ErrInvalidEmail
	}

	// Step 1: Check if provider is in list of blocked email providers
	for _, blocked := range providers.blockList {
		if strings.HasSuffix(email, blocked) {
			// If provider is in blocked list, check if provider is explicitly allowed
			for _, allowed := range providers.allowList {
				if strings.HasSuffix(email, allowed) {
					return false, nil
				}
			}
			// If provider is not explicitly allowed, return true
			return true, ErrDisposableEmail
		}
	}

	// Step 2: Check if the domain has an SPF record and check if the length of SPF record are generally greater than 24 chars for a well configured
	// email provider
	// TODO: Do more indepth checks against, MX and SPF records.
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return true, ErrInvalidEmail
	}
	domain := parts[1]

	// Check if domain name exists
	_, err := checkDomain(domain)
	if err != nil {
		// if no ip address are found, TXT (SPF) records
		txtRecords, err := net.LookupTXT(domain)
		if err != nil {
			return true, ErrCouldNotLookupSPFRecords
		}

		for _, spfRecord := range txtRecords {
			if strings.HasPrefix(spfRecord, "v=spf1 ") {
				if len(spfRecord) < 25 {
					return true, ErrDisposableEmail
				}
			}
		}
	}

	// Step 3: Look for email addresses that contain the words "test", "temp", or "demo" (case-insensitive)
	if disposableEmailPattern.MatchString(email) {
		return true, ErrDisposableEmail
	}

	return false, nil
}

func isEmail(input string) bool {
	match, _ := regexp.MatchString(emailRegex, input)
	return match
}

func checkDomain(domainName string) ([]string, error) {
	ips, err := net.LookupHost(domainName)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
