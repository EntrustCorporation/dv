package dns

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/digitorus/dv/dns01"
	"github.com/digitorus/dv/providers/acmedns"
	"github.com/digitorus/dv/providers/alidns"
	"github.com/digitorus/dv/providers/arvancloud"
	"github.com/digitorus/dv/providers/auroradns"
	"github.com/digitorus/dv/providers/autodns"
	"github.com/digitorus/dv/providers/azure"
	"github.com/digitorus/dv/providers/bindman"
	"github.com/digitorus/dv/providers/bluecat"
	"github.com/digitorus/dv/providers/checkdomain"
	"github.com/digitorus/dv/providers/clouddns"
	"github.com/digitorus/dv/providers/cloudflare"
	"github.com/digitorus/dv/providers/cloudns"
	"github.com/digitorus/dv/providers/cloudxns"
	"github.com/digitorus/dv/providers/conoha"
	"github.com/digitorus/dv/providers/constellix"
	"github.com/digitorus/dv/providers/desec"
	"github.com/digitorus/dv/providers/designate"
	"github.com/digitorus/dv/providers/digitalocean"
	"github.com/digitorus/dv/providers/dnsimple"
	"github.com/digitorus/dv/providers/dnsmadeeasy"
	"github.com/digitorus/dv/providers/dnspod"
	"github.com/digitorus/dv/providers/dode"
	"github.com/digitorus/dv/providers/domeneshop"
	"github.com/digitorus/dv/providers/dreamhost"
	"github.com/digitorus/dv/providers/duckdns"
	"github.com/digitorus/dv/providers/dyn"
	"github.com/digitorus/dv/providers/dynu"
	"github.com/digitorus/dv/providers/easydns"
	"github.com/digitorus/dv/providers/edgedns"
	"github.com/digitorus/dv/providers/exec"
	"github.com/digitorus/dv/providers/exoscale"
	"github.com/digitorus/dv/providers/gandi"
	"github.com/digitorus/dv/providers/gandiv5"
	"github.com/digitorus/dv/providers/gcloud"
	"github.com/digitorus/dv/providers/glesys"
	"github.com/digitorus/dv/providers/godaddy"
	"github.com/digitorus/dv/providers/hetzner"
	"github.com/digitorus/dv/providers/hostingde"
	"github.com/digitorus/dv/providers/httpreq"
	"github.com/digitorus/dv/providers/hurricane"
	"github.com/digitorus/dv/providers/hyperone"
	"github.com/digitorus/dv/providers/iij"
	"github.com/digitorus/dv/providers/infomaniak"
	"github.com/digitorus/dv/providers/inwx"
	"github.com/digitorus/dv/providers/ionos"
	"github.com/digitorus/dv/providers/joker"
	"github.com/digitorus/dv/providers/lightsail"
	"github.com/digitorus/dv/providers/linode"
	"github.com/digitorus/dv/providers/liquidweb"
	"github.com/digitorus/dv/providers/loopia"
	"github.com/digitorus/dv/providers/luadns"
	"github.com/digitorus/dv/providers/mydnsjp"
	"github.com/digitorus/dv/providers/mythicbeasts"
	"github.com/digitorus/dv/providers/namecheap"
	"github.com/digitorus/dv/providers/namedotcom"
	"github.com/digitorus/dv/providers/namesilo"
	"github.com/digitorus/dv/providers/netcup"
	"github.com/digitorus/dv/providers/netlify"
	"github.com/digitorus/dv/providers/nifcloud"
	"github.com/digitorus/dv/providers/njalla"
	"github.com/digitorus/dv/providers/ns1"
	"github.com/digitorus/dv/providers/oraclecloud"
	"github.com/digitorus/dv/providers/otc"
	"github.com/digitorus/dv/providers/ovh"
	"github.com/digitorus/dv/providers/pdns"
	"github.com/digitorus/dv/providers/porkbun"
	"github.com/digitorus/dv/providers/rackspace"
	"github.com/digitorus/dv/providers/regru"
	"github.com/digitorus/dv/providers/rfc2136"
	"github.com/digitorus/dv/providers/rimuhosting"
	"github.com/digitorus/dv/providers/route53"
	"github.com/digitorus/dv/providers/sakuracloud"
	"github.com/digitorus/dv/providers/scaleway"
	"github.com/digitorus/dv/providers/selectel"
	"github.com/digitorus/dv/providers/servercow"
	"github.com/digitorus/dv/providers/sonic"
	"github.com/digitorus/dv/providers/stackpath"
	"github.com/digitorus/dv/providers/transip"
	"github.com/digitorus/dv/providers/vegadns"
	"github.com/digitorus/dv/providers/versio"
	"github.com/digitorus/dv/providers/vinyldns"
	"github.com/digitorus/dv/providers/vscale"
	"github.com/digitorus/dv/providers/vultr"
	"github.com/digitorus/dv/providers/wedos"
	"github.com/digitorus/dv/providers/yandex"
	"github.com/digitorus/dv/providers/zoneee"
	"github.com/digitorus/dv/providers/zonomi"
)

// NewDNSChallengeProviderByName Factory for DNS providers.
func NewDNSChallengeProviderByName(name string) (challenge.Provider, error) {
	switch name {
	case "acme-dns":
		return acmedns.NewDNSProvider()
	case "alidns":
		return alidns.NewDNSProvider()
	case "arvancloud":
		return arvancloud.NewDNSProvider()
	case "azure":
		return azure.NewDNSProvider()
	case "auroradns":
		return auroradns.NewDNSProvider()
	case "autodns":
		return autodns.NewDNSProvider()
	case "bindman":
		return bindman.NewDNSProvider()
	case "bluecat":
		return bluecat.NewDNSProvider()
	case "checkdomain":
		return checkdomain.NewDNSProvider()
	case "clouddns":
		return clouddns.NewDNSProvider()
	case "cloudflare":
		return cloudflare.NewDNSProvider()
	case "cloudns":
		return cloudns.NewDNSProvider()
	case "cloudxns":
		return cloudxns.NewDNSProvider()
	case "conoha":
		return conoha.NewDNSProvider()
	case "constellix":
		return constellix.NewDNSProvider()
	case "desec":
		return desec.NewDNSProvider()
	case "designate":
		return designate.NewDNSProvider()
	case "digitalocean":
		return digitalocean.NewDNSProvider()
	case "dnsimple":
		return dnsimple.NewDNSProvider()
	case "dnsmadeeasy":
		return dnsmadeeasy.NewDNSProvider()
	case "dnspod":
		return dnspod.NewDNSProvider()
	case "dode":
		return dode.NewDNSProvider()
	case "domeneshop", "domainnameshop":
		return domeneshop.NewDNSProvider()
	case "dreamhost":
		return dreamhost.NewDNSProvider()
	case "duckdns":
		return duckdns.NewDNSProvider()
	case "dyn":
		return dyn.NewDNSProvider()
	case "dynu":
		return dynu.NewDNSProvider()
	case "easydns":
		return easydns.NewDNSProvider()
	case "edgedns", "fastdns": // "fastdns" is for compatibility with v3, must be dropped in v5
		return edgedns.NewDNSProvider()
	case "exec":
		return exec.NewDNSProvider()
	case "exoscale":
		return exoscale.NewDNSProvider()
	case "gandi":
		return gandi.NewDNSProvider()
	case "gandiv5":
		return gandiv5.NewDNSProvider()
	case "glesys":
		return glesys.NewDNSProvider()
	case "gcloud":
		return gcloud.NewDNSProvider()
	case "godaddy":
		return godaddy.NewDNSProvider()
	case "hetzner":
		return hetzner.NewDNSProvider()
	case "hostingde":
		return hostingde.NewDNSProvider()
	case "httpreq":
		return httpreq.NewDNSProvider()
	case "hurricane":
		return hurricane.NewDNSProvider()
	case "hyperone":
		return hyperone.NewDNSProvider()
	case "iij":
		return iij.NewDNSProvider()
	case "infomaniak":
		return infomaniak.NewDNSProvider()
	case "inwx":
		return inwx.NewDNSProvider()
	case "ionos":
		return ionos.NewDNSProvider()
	case "joker":
		return joker.NewDNSProvider()
	case "lightsail":
		return lightsail.NewDNSProvider()
	case "linode", "linodev4": // "linodev4" is for compatibility with v3, must be dropped in v5
		return linode.NewDNSProvider()
	case "liquidweb":
		return liquidweb.NewDNSProvider()
	case "luadns":
		return luadns.NewDNSProvider()
	case "loopia":
		return loopia.NewDNSProvider()
	case "manual":
		return dns01.NewDNSProviderManual()
	case "mydnsjp":
		return mydnsjp.NewDNSProvider()
	case "mythicbeasts":
		return mythicbeasts.NewDNSProvider()
	case "namecheap":
		return namecheap.NewDNSProvider()
	case "namedotcom":
		return namedotcom.NewDNSProvider()
	case "namesilo":
		return namesilo.NewDNSProvider()
	case "netcup":
		return netcup.NewDNSProvider()
	case "netlify":
		return netlify.NewDNSProvider()
	case "nifcloud":
		return nifcloud.NewDNSProvider()
	case "njalla":
		return njalla.NewDNSProvider()
	case "ns1":
		return ns1.NewDNSProvider()
	case "oraclecloud":
		return oraclecloud.NewDNSProvider()
	case "otc":
		return otc.NewDNSProvider()
	case "ovh":
		return ovh.NewDNSProvider()
	case "pdns":
		return pdns.NewDNSProvider()
	case "porkbun":
		return porkbun.NewDNSProvider()
	case "rackspace":
		return rackspace.NewDNSProvider()
	case "regru":
		return regru.NewDNSProvider()
	case "rfc2136":
		return rfc2136.NewDNSProvider()
	case "rimuhosting":
		return rimuhosting.NewDNSProvider()
	case "route53":
		return route53.NewDNSProvider()
	case "sakuracloud":
		return sakuracloud.NewDNSProvider()
	case "scaleway":
		return scaleway.NewDNSProvider()
	case "selectel":
		return selectel.NewDNSProvider()
	case "servercow":
		return servercow.NewDNSProvider()
	case "sonic":
		return sonic.NewDNSProvider()
	case "stackpath":
		return stackpath.NewDNSProvider()
	case "transip":
		return transip.NewDNSProvider()
	case "vegadns":
		return vegadns.NewDNSProvider()
	case "versio":
		return versio.NewDNSProvider()
	case "vinyldns":
		return vinyldns.NewDNSProvider()
	case "vultr":
		return vultr.NewDNSProvider()
	case "vscale":
		return vscale.NewDNSProvider()
	case "wedos":
		return wedos.NewDNSProvider()
	case "yandex":
		return yandex.NewDNSProvider()
	case "zoneee":
		return zoneee.NewDNSProvider()
	case "zonomi":
		return zonomi.NewDNSProvider()
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
