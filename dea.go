package dea

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

var (
	ErrDisposableEmail          = errors.New("email is a disposable email address")
	ErrInvalidEmail             = errors.New("not valid email format")
	ErrCouldNotLookupSPFRecords = errors.New("could not lookup SPF records")
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
			// If provider is not explicitly allowed, return true
			return true, ErrDisposableEmail
		}
	}

	// Check if provider is in list of part of the allowed email provider. If yes then return
	for _, allowed := range providers.allowList {
		if strings.HasSuffix(email, allowed) {
			return false, nil
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

	// Step 3: check for email addresses that contain random strings of characters or misspelled words
	pattern := regexp.MustCompile(`^([a-z0-9._%+\-]+|([a-z0-9]+[.\-_])*[a-z0-9]+)@(10|anon|example|guerrilla|fake|disposable|temp)([.\-_][a-z]{2,})+$`)
	if pattern.MatchString(email) {
		return true, ErrDisposableEmail
	}

	return false, nil
}
