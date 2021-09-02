package domain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/admpub/log"
	"github.com/admpub/nging/v3/application/library/ddnsmanager"
	"github.com/admpub/nging/v3/application/library/ddnsmanager/config"
	"github.com/admpub/nging/v3/application/library/ddnsmanager/domain/dnsdomain"
	"github.com/admpub/nging/v3/application/library/ddnsmanager/resolver"
	"github.com/admpub/nging/v3/application/library/ddnsmanager/utils"
	"golang.org/x/net/publicsuffix"
)

// Domains Ipv4/Ipv6 domains
type Domains struct {
	IPv4Addr    string
	IPv4Domains map[string][]*dnsdomain.Domain // {dnspod:[]}
	IPv6Addr    string
	IPv6Domains map[string][]*dnsdomain.Domain // {dnspod:[]}
}

func NewDomains() *Domains {
	return &Domains{
		IPv4Domains: map[string][]*dnsdomain.Domain{},
		IPv6Domains: map[string][]*dnsdomain.Domain{},
	}
}

// ParseDomain 接口获得ip并校验用户输入的域名
func ParseDomain(conf *config.Config) (*Domains, error) {
	domains := NewDomains()
	var err error
	// IPv4
	ipv4Addr := utils.GetIPv4Addr(conf.IPv4.NetInterface, conf.IPv4.NetIPApiUrl)
	if len(ipv4Addr) > 0 {
		domains.IPv4Addr = ipv4Addr
		for _, service := range conf.DNSServices {
			_, ok := domains.IPv4Domains[service.Provider]
			if !ok {
				domains.IPv4Domains[service.Provider] = []*dnsdomain.Domain{}
			}
			domains.IPv4Domains[service.Provider], err = parseDomainArr(service.IPv4Domains)
			if err != nil {
				return domains, err
			}
		}
	}
	// IPv6
	ipv6Addr := utils.GetIPv6Addr(conf.IPv6.NetInterface, conf.IPv6.NetIPApiUrl)
	if len(ipv6Addr) > 0 {
		domains.IPv6Addr = ipv6Addr
		for _, service := range conf.DNSServices {
			_, ok := domains.IPv6Domains[service.Provider]
			if !ok {
				domains.IPv6Domains[service.Provider] = []*dnsdomain.Domain{}
			}
			domains.IPv6Domains[service.Provider], err = parseDomainArr(service.IPv6Domains)
		}
	}
	return domains, err
}

func (domains *Domains) Update(conf *config.Config) error {
	var errs []error
	// IPv4
	if ipv4Addr := utils.GetIPv4Addr(conf.IPv4.NetInterface, conf.IPv4.NetIPApiUrl); len(ipv4Addr) > 0 && domains.IPv4Addr != ipv4Addr {
		domains.IPv4Addr = ipv4Addr
		for dnsProvider, dnsDomains := range domains.IPv4Domains {
			var _dnsDomains []*dnsdomain.Domain
			for _, dnsDomain := range dnsDomains {
				oldIP, err := resolver.ResolveDNS(dnsDomain.String(), conf.DNSResolver, `IPV4`)
				if err != nil {
					log.Errorf("[%s] ResolveDNS(%s): %s", dnsProvider, dnsDomain.String(), err.Error())
					errs = append(errs, err)
					continue
				}
				if oldIP != ipv4Addr {
					_dnsDomains = append(_dnsDomains, dnsDomain)
					continue
				}
				log.Infof("[%s] IP is the same as cached one (%s). Skip update (%s)", dnsProvider, ipv4Addr, dnsDomain.String())
			}
			if len(_dnsDomains) == 0 {
				continue
			}
			updater := ddnsmanager.Open(dnsProvider)
			if updater == nil {
				continue
			}
			dnsService := conf.FindService(dnsProvider)
			err := updater.Init(dnsService.Settings, _dnsDomains)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			log.Infof("[%s] %s - Start to update record IP...", dnsProvider, ipv4Addr)
			err = updater.Update(`A`, ipv4Addr)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	// IPv6
	if ipv6Addr := utils.GetIPv6Addr(conf.IPv6.NetInterface, conf.IPv6.NetIPApiUrl); len(ipv6Addr) > 0 && domains.IPv6Addr != ipv6Addr {
		domains.IPv6Addr = ipv6Addr
		for dnsProvider, dnsDomains := range domains.IPv6Domains {
			var _dnsDomains []*dnsdomain.Domain
			for _, dnsDomain := range dnsDomains {
				oldIP, err := resolver.ResolveDNS(dnsDomain.String(), conf.DNSResolver, `IPV6`)
				if err != nil {
					log.Errorf("[%s] ResolveDNS(%s): %s", dnsProvider, dnsDomain.String(), err.Error())
					errs = append(errs, err)
					continue
				}
				if oldIP != ipv6Addr {
					_dnsDomains = append(_dnsDomains, dnsDomain)
					continue
				}
				log.Infof("[%s] IP is the same as cached one (%s). Skip update (%s)", dnsProvider, ipv6Addr, dnsDomain.String())
			}
			if len(_dnsDomains) == 0 {
				continue
			}
			updater := ddnsmanager.Open(dnsProvider)
			if updater == nil {
				continue
			}
			dnsService := conf.FindService(dnsProvider)
			err := updater.Init(dnsService.Settings, _dnsDomains)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			log.Infof("[%s] %s - Start to update record IP...", dnsProvider, ipv6Addr)
			err = updater.Update(`AAAA`, ipv6Addr)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	var err error
	if len(errs) > 0 {
		errMessages := make([]string, len(errs))
		for index, err := range errs {
			errMessages[index] = err.Error()
		}
		err = errors.New(strings.Join(errMessages, "\n"))
	}
	return err
}

// parseDomainArr 校验用户输入的域名
func parseDomainArr(dnsDomains []*config.DNSDomain) (domains []*dnsdomain.Domain, err error) {
	for _, dnsDomain := range dnsDomains {
		_domain := strings.TrimSpace(dnsDomain.Domain)
		if len(_domain) == 0 {
			continue
		}
		domain := &dnsdomain.Domain{
			Port:         dnsDomain.Port,
			UpdateStatus: dnsdomain.UpdatedIdle,
		}
		sp := strings.Split(_domain, ".")
		length := len(sp)
		if length <= 1 {
			err = fmt.Errorf(`域名不正确: %s`, _domain)
			return
		}
		var topLevelDomain string
		// 处理域名
		topLevelDomain, err = publicsuffix.EffectiveTLDPlusOne(_domain)
		if err != nil {
			err = fmt.Errorf(`域名不正确: %w`, err)
			return
		}
		domain.DomainName = topLevelDomain
		domainLen := len(_domain) - len(domain.DomainName)
		if domainLen > 0 {
			domain.SubDomain = _domain[:domainLen-1]
		} else {
			domain.SubDomain = _domain[:domainLen]
		}
		domains = append(domains, domain)
	}
	return
}
