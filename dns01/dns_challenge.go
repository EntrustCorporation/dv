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
func GetRecord(domain, keyAuth string) (fqdn, value string) {
	return fmt.Sprintf("_pki-validation.%s.", domain), keyAuth
}
