package component

import (
	"net"

	log "github.com/sirupsen/logrus"
)

func DNSResolver(domain string) error {
	_, err := net.LookupHost(domain)
	if err != nil {
		log.Error("Domain lookup failed for: ", domain, ", error:", err)
		return err
	}

	return nil
}
