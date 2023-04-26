package dea

import (
	"net"
	"regexp"
	"strings"
)

func IsDisposableEmail(email string) (bool, error) {
	providers := NewEmailProviders()

	// Step 1: Check if provider is in list of blocked email providers
	for _, blocked := range providers.blockList {
		if strings.HasSuffix(email, blocked) {
			// If provider is in blocked list, check if provider is explicitly allowed
			for _, allowed := range providers.allowList {
				if strings.HasSuffix(email, allowed) {
					return false, nil
				}
			}
		}
	}

	// Step 2: Check if the domain has an SPF record and check if the length of SPF record are generally greater than 24 chars for a well configured
	// email provider
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return true, nil
	}
	domain := parts[1]

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		return true, nil
	}

	for _, spfRecord := range txtRecords {
		if strings.HasPrefix(spfRecord, "v=spf1 ") {
			if len(spfRecord) < 25 {
				return true, nil
			}
		}
	}

	// Step 3: check for email addresses that contain random strings of characters or misspelled words
	pattern := regexp.MustCompile(`^([a-z0-9._%+\-]+|([a-z0-9]+[.\-_])*[a-z0-9]+)@(10|anon|example|guerrilla|fake|disposable|temp)([.\-_][a-z]{2,})+$`)
	if pattern.MatchString(email) {
		return true, nil
	}

	return false, nil
}
