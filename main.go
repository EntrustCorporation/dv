package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/weppos/publicsuffix-go/publicsuffix"

	dns "github.com/digitorus/dv/providers"
	"github.com/digitorus/entrust"
)

func main() {
	log.SetLevel(log.DebugLevel)

	e, err := entrust.New()
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("DNS_PROVIDER") == "" {
		log.Fatal("DNS_PROVIDER not configured")
	}

	// By default, start renewing domains one month before they expire.
	//
	// Setting a duration in the environment variable DURATION_BEFORE_EXPIRY
	// will override the default (e.g., 48h, to set two days).
	expiry := time.Now().Add(31 * 24 * time.Hour)
	if duration := os.Getenv("DURATION_BEFORE_EXPIRY"); duration != "" {
		rt, err := time.ParseDuration(duration)
		if err != nil {
			log.Fatal(err)
		}
		expiry = time.Now().Add(rt)
	}

	offset := 0
	for {
		domains, err := e.GetDomains(offset, 250, &expiry)
		if err != nil {
			log.Fatal(err)
		}

		if offset == 0 {
			log.Infof("Processing %d domains", domains.Summary.Total)
		}

		for _, domain := range domains.Domains {
			log.WithFields(log.Fields{"domain": domain.DomainName}).Info("Processing")
			_, err := publicsuffix.DomainFromListWithOptions(publicsuffix.DefaultList, domain.DomainName, &publicsuffix.FindOptions{IgnorePrivate: true})
			if err != nil {
				log.WithFields(log.Fields{"domain": domain.DomainName}).Info("Skipping domain as it does not end with a Public Suffix")
				continue
			}

			switch domain.VerificationStatus {
			case "RE_VERIFICATION", "INITIAL_VERIFICATION":
				// Pending verification, check if DNS is configured correctly
				break
			case "CANCELLED", "DECLINED":
				// This domain has been deleted or is not invalid, don't try to renew it
				continue
			default:
				// Resubmit the domain for validation
				e.ReverifyDomain(domain.ClientID, domain.DomainName)
				time.Sleep(5 * time.Second) // wait for the code to be generated
			}

			log.WithFields(log.Fields{"domain": domain.DomainName}).Info("Getting verification code")

			domainResponse, err := e.GetDomain(domain.ClientID, domain.DomainName)
			if err != nil {
				log.Println(err)
				continue
			}
			if domainResponse.VerificationStatus == "APPROVED" {
				log.Warning("Unexpected status, not ready for re-verification, try again in a few minutes...")
				continue
			}

			log.WithFields(log.Fields{"domain": domain.DomainName}).Info("Updating DNS")
			renewDomainValidation(domainResponse)

			log.Println()
		}

		// New offset, break if we did them all
		offset = domains.Summary.Offset + domains.Summary.Limit
		if offset >= domains.Summary.Total {
			log.Info("Done")
			break
		}
	}
}

func renewDomainValidation(domain *entrust.Domain) {
	if domain == nil || domain.DNSMethod == nil {
		log.Error("Renew domain validation failed: invalid domain or method")
		return
	}

	// A default DNS provider can be set using the environment variable DNS_PROVIDER
	// and a domain specific provider can be set using the environment variable
	// DNS_PROVIDER_example_com.
	dnsProvider := os.Getenv("DNS_PROVIDER")
	if dnsProv := os.Getenv(fmt.Sprintf("DNS_PROVIDER_%s",
		strings.Replace(domain.DomainName, ".", "_", -1))); dnsProv != "" {

		dnsProvider = dnsProv
	}

	dnsClient, err := dns.NewDNSChallengeProviderByName(dnsProvider)
	if err != nil {
		log.Fatal(err)
	}

	err = dnsClient.Present(domain.DomainName, domain.DNSMethod.RecordDomain, domain.DNSMethod.RecordValue)
	if err != nil {
		log.Error("Present failed:", err)
		return
	}

	// Ideally we would run dnsClient.CleanUp here but this requires the validation to be completed
	// as the lego client does only delete records based on known id's and the Entrust client doesn't
	// allow us to initiate a direct verification.
}
