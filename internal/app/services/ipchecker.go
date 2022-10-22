package services

import (
	"net/http"
	"net/netip"

	"github.com/belamov/ypgo-url-shortener/internal/app/config"
)

type IPCheckerInterface interface {
	IsRequestFromTrustedSubnet(r *http.Request) (bool, error)
}

type IPChecker struct {
	trustedIPNet netip.Prefix
}

const realIPHeader = "X-Real-IP"

func NewIPChecker(config *config.Config) (*IPChecker, error) {
	CIDR := config.TrustedSubnet
	network, err := netip.ParsePrefix(CIDR)
	if err != nil {
		return nil, err
	}

	return &IPChecker{trustedIPNet: network}, nil
}

func (c *IPChecker) IsRequestFromTrustedSubnet(r *http.Request) (bool, error) {
	ipStr := r.Header.Get(realIPHeader)
	ip, err := netip.ParseAddr(ipStr)
	if err != nil {
		return false, err
	}
	return c.trustedIPNet.Contains(ip), nil
}
