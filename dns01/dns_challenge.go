package dns01

import (
	"fmt"
	"time"
)

const (
	// DefaultPropagationTimeout default propagation timeout.
	DefaultPropagationTimeout = 60 * time.Second

	// DefaultPollingInterval default polling interval.
	DefaultPollingInterval = 2 * time.Second

	// DefaultTTL default TTL.
	DefaultTTL = 120
)

// GetRecord returns a DNS record which will fulfill the `dns-01` challenge.
// Modified to function for non ACME domain validations
// Deprecated: use GetChallengeInfo instead.
func GetRecord(domain, keyAuth string) (fqdn, value string) {
	return fmt.Sprintf("_pki-validation.%s.", domain), keyAuth
}

// ChallengeInfo contains the information use to create the TXT record.
type ChallengeInfo struct {
	// FQDN is the full-qualified challenge domain (i.e. `_acme-challenge.[domain].`)
	FQDN string

	// EffectiveFQDN contains the resulting FQDN after the CNAMEs resolutions.
	EffectiveFQDN string

	// Value contains the value for the TXT record.
	Value string
}

// GetChallengeInfo returns information used to create a DNS record which will fulfill the `dns-01` challenge.
func GetChallengeInfo(domain, keyAuth string) ChallengeInfo {
	return ChallengeInfo{
		Value:         keyAuth,
		FQDN:          fmt.Sprintf("_pki-validation.%s.", domain),
		EffectiveFQDN: fmt.Sprintf("_pki-validation.%s.", domain),
	}
}
