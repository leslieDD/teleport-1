package service

import (
	"context"
	"net"
	"strings"

	"github.com/gravitational/teleport/lib/utils"

	"golang.org/x/crypto/acme/autocert"

	"github.com/gravitational/trace"
)

// hostPolicy approves certificates that are subdomains of public addresses
func hostPolicy(addrs []utils.NetAddr) (autocert.HostPolicy, error) {
	dnsNames := make([]string, 0, len(addrs))

	for _, addr := range addrs {
		host, err := utils.Host(addr.Addr)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		if ip := net.ParseIP(host); ip == nil {
			dnsNames = append(dnsNames, host)
		}
	}

	return func(_ context.Context, host string) error {
		for _, dnsName := range dnsNames {
			if dnsName == host || strings.HasSuffix(host, "."+dnsName) {
				return nil
			}
		}
		return trace.BadParameter(
			"acme does not recognize domain %q, please add it in proxy public_addr section, supported domains are: %v (including subdomains)",
			host, strings.Join(dnsNames, ","))
	}, nil
}
