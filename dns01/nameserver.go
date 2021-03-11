package dns01

import "strings"

// FindZoneByFqdn is used by Lego DNS providers to determine the domain
func FindZoneByFqdn(fqdn string) (string, error) {
	return strings.TrimPrefix(fqdn, "_pki-validation."), nil
}

// FindZoneByFqdnCustom is used by Lego DNS providers to determine the domain
func FindZoneByFqdnCustom(fqdn string, ns []string) (string, error) {
	return FindZoneByFqdn(fqdn)
}
