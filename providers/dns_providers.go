package dns

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/entrustcorporation/dv/dns01"
	"github.com/entrustcorporation/dv/providers/acmedns"
	"github.com/entrustcorporation/dv/providers/alidns"
	"github.com/entrustcorporation/dv/providers/allinkl"
	"github.com/entrustcorporation/dv/providers/arvancloud"
	"github.com/entrustcorporation/dv/providers/auroradns"
	"github.com/entrustcorporation/dv/providers/autodns"
	"github.com/entrustcorporation/dv/providers/azure"
	"github.com/entrustcorporation/dv/providers/azuredns"
	"github.com/entrustcorporation/dv/providers/bindman"
	"github.com/entrustcorporation/dv/providers/bluecat"
	"github.com/entrustcorporation/dv/providers/brandit"
	"github.com/entrustcorporation/dv/providers/bunny"
	"github.com/entrustcorporation/dv/providers/checkdomain"
	"github.com/entrustcorporation/dv/providers/civo"
	"github.com/entrustcorporation/dv/providers/clouddns"
	"github.com/entrustcorporation/dv/providers/cloudflare"
	"github.com/entrustcorporation/dv/providers/cloudns"
	"github.com/entrustcorporation/dv/providers/cloudru"
	"github.com/entrustcorporation/dv/providers/cloudxns"
	"github.com/entrustcorporation/dv/providers/conoha"
	"github.com/entrustcorporation/dv/providers/constellix"
	"github.com/entrustcorporation/dv/providers/cpanel"
	"github.com/entrustcorporation/dv/providers/derak"
	"github.com/entrustcorporation/dv/providers/desec"
	"github.com/entrustcorporation/dv/providers/designate"
	"github.com/entrustcorporation/dv/providers/digitalocean"
	"github.com/entrustcorporation/dv/providers/dnshomede"
	"github.com/entrustcorporation/dv/providers/dnsimple"
	"github.com/entrustcorporation/dv/providers/dnsmadeeasy"
	"github.com/entrustcorporation/dv/providers/dnspod"
	"github.com/entrustcorporation/dv/providers/dode"
	"github.com/entrustcorporation/dv/providers/domeneshop"
	"github.com/entrustcorporation/dv/providers/dreamhost"
	"github.com/entrustcorporation/dv/providers/duckdns"
	"github.com/entrustcorporation/dv/providers/dyn"
	"github.com/entrustcorporation/dv/providers/dynu"
	"github.com/entrustcorporation/dv/providers/easydns"
	"github.com/entrustcorporation/dv/providers/edgedns"
	"github.com/entrustcorporation/dv/providers/efficientip"
	"github.com/entrustcorporation/dv/providers/epik"
	"github.com/entrustcorporation/dv/providers/exec"
	"github.com/entrustcorporation/dv/providers/exoscale"
	"github.com/entrustcorporation/dv/providers/freemyip"
	"github.com/entrustcorporation/dv/providers/gandi"
	"github.com/entrustcorporation/dv/providers/gandiv5"
	"github.com/entrustcorporation/dv/providers/gcloud"
	"github.com/entrustcorporation/dv/providers/gcore"
	"github.com/entrustcorporation/dv/providers/glesys"
	"github.com/entrustcorporation/dv/providers/godaddy"
	"github.com/entrustcorporation/dv/providers/googledomains"
	"github.com/entrustcorporation/dv/providers/hetzner"
	"github.com/entrustcorporation/dv/providers/hostingde"
	"github.com/entrustcorporation/dv/providers/hosttech"
	"github.com/entrustcorporation/dv/providers/httpreq"
	"github.com/entrustcorporation/dv/providers/hurricane"
	"github.com/entrustcorporation/dv/providers/hyperone"
	"github.com/entrustcorporation/dv/providers/ibmcloud"
	"github.com/entrustcorporation/dv/providers/iij"
	"github.com/entrustcorporation/dv/providers/iijdpf"
	"github.com/entrustcorporation/dv/providers/infoblox"
	"github.com/entrustcorporation/dv/providers/infomaniak"
	"github.com/entrustcorporation/dv/providers/internetbs"
	"github.com/entrustcorporation/dv/providers/inwx"
	"github.com/entrustcorporation/dv/providers/ionos"
	"github.com/entrustcorporation/dv/providers/ipv64"
	"github.com/entrustcorporation/dv/providers/iwantmyname"
	"github.com/entrustcorporation/dv/providers/joker"
	"github.com/entrustcorporation/dv/providers/liara"
	"github.com/entrustcorporation/dv/providers/lightsail"
	"github.com/entrustcorporation/dv/providers/linode"
	"github.com/entrustcorporation/dv/providers/liquidweb"
	"github.com/entrustcorporation/dv/providers/loopia"
	"github.com/entrustcorporation/dv/providers/luadns"
	"github.com/entrustcorporation/dv/providers/mailinabox"
	"github.com/entrustcorporation/dv/providers/metaname"
	"github.com/entrustcorporation/dv/providers/mydnsjp"
	"github.com/entrustcorporation/dv/providers/mythicbeasts"
	"github.com/entrustcorporation/dv/providers/namecheap"
	"github.com/entrustcorporation/dv/providers/namedotcom"
	"github.com/entrustcorporation/dv/providers/namesilo"
	"github.com/entrustcorporation/dv/providers/nearlyfreespeech"
	"github.com/entrustcorporation/dv/providers/netcup"
	"github.com/entrustcorporation/dv/providers/netlify"
	"github.com/entrustcorporation/dv/providers/nicmanager"
	"github.com/entrustcorporation/dv/providers/nifcloud"
	"github.com/entrustcorporation/dv/providers/njalla"
	"github.com/entrustcorporation/dv/providers/nodion"
	"github.com/entrustcorporation/dv/providers/ns1"
	"github.com/entrustcorporation/dv/providers/oraclecloud"
	"github.com/entrustcorporation/dv/providers/otc"
	"github.com/entrustcorporation/dv/providers/ovh"
	"github.com/entrustcorporation/dv/providers/pdns"
	"github.com/entrustcorporation/dv/providers/plesk"
	"github.com/entrustcorporation/dv/providers/porkbun"
	"github.com/entrustcorporation/dv/providers/rackspace"
	"github.com/entrustcorporation/dv/providers/rcodezero"
	"github.com/entrustcorporation/dv/providers/regru"
	"github.com/entrustcorporation/dv/providers/rfc2136"
	"github.com/entrustcorporation/dv/providers/rimuhosting"
	"github.com/entrustcorporation/dv/providers/route53"
	"github.com/entrustcorporation/dv/providers/safedns"
	"github.com/entrustcorporation/dv/providers/sakuracloud"
	"github.com/entrustcorporation/dv/providers/scaleway"
	"github.com/entrustcorporation/dv/providers/selectel"
	"github.com/entrustcorporation/dv/providers/servercow"
	"github.com/entrustcorporation/dv/providers/shellrent"
	"github.com/entrustcorporation/dv/providers/simply"
	"github.com/entrustcorporation/dv/providers/sonic"
	"github.com/entrustcorporation/dv/providers/stackpath"
	"github.com/entrustcorporation/dv/providers/tencentcloud"
	"github.com/entrustcorporation/dv/providers/transip"
	"github.com/entrustcorporation/dv/providers/ultradns"
	"github.com/entrustcorporation/dv/providers/variomedia"
	"github.com/entrustcorporation/dv/providers/vegadns"
	"github.com/entrustcorporation/dv/providers/vercel"
	"github.com/entrustcorporation/dv/providers/versio"
	"github.com/entrustcorporation/dv/providers/vinyldns"
	"github.com/entrustcorporation/dv/providers/vkcloud"
	"github.com/entrustcorporation/dv/providers/vscale"
	"github.com/entrustcorporation/dv/providers/vultr"
	"github.com/entrustcorporation/dv/providers/webnames"
	"github.com/entrustcorporation/dv/providers/websupport"
	"github.com/entrustcorporation/dv/providers/wedos"
	"github.com/entrustcorporation/dv/providers/yandex"
	"github.com/entrustcorporation/dv/providers/yandex360"
	"github.com/entrustcorporation/dv/providers/yandexcloud"
	"github.com/entrustcorporation/dv/providers/zoneee"
	"github.com/entrustcorporation/dv/providers/zonomi"
)

// NewDNSChallengeProviderByName Factory for DNS providers.
func NewDNSChallengeProviderByName(name string) (challenge.Provider, error) {
	switch name {
	case "acme-dns": // TODO(ldez): remove "-" in v5
		return acmedns.NewDNSProvider()
	case "alidns":
		return alidns.NewDNSProvider()
	case "allinkl":
		return allinkl.NewDNSProvider()
	case "arvancloud":
		return arvancloud.NewDNSProvider()
	case "azure":
		return azure.NewDNSProvider()
	case "azuredns":
		return azuredns.NewDNSProvider()
	case "auroradns":
		return auroradns.NewDNSProvider()
	case "autodns":
		return autodns.NewDNSProvider()
	case "bindman":
		return bindman.NewDNSProvider()
	case "bluecat":
		return bluecat.NewDNSProvider()
	case "brandit":
		return brandit.NewDNSProvider()
	case "bunny":
		return bunny.NewDNSProvider()
	case "checkdomain":
		return checkdomain.NewDNSProvider()
	case "civo":
		return civo.NewDNSProvider()
	case "clouddns":
		return clouddns.NewDNSProvider()
	case "cloudflare":
		return cloudflare.NewDNSProvider()
	case "cloudns":
		return cloudns.NewDNSProvider()
	case "cloudru":
		return cloudru.NewDNSProvider()
	case "cloudxns":
		return cloudxns.NewDNSProvider()
	case "conoha":
		return conoha.NewDNSProvider()
	case "constellix":
		return constellix.NewDNSProvider()
	case "cpanel":
		return cpanel.NewDNSProvider()
	case "derak":
		return derak.NewDNSProvider()
	case "desec":
		return desec.NewDNSProvider()
	case "designate":
		return designate.NewDNSProvider()
	case "digitalocean":
		return digitalocean.NewDNSProvider()
	case "dnshomede":
		return dnshomede.NewDNSProvider()
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
	case "efficientip":
		return efficientip.NewDNSProvider()
	case "epik":
		return epik.NewDNSProvider()
	case "exec":
		return exec.NewDNSProvider()
	case "exoscale":
		return exoscale.NewDNSProvider()
	case "freemyip":
		return freemyip.NewDNSProvider()
	case "gandi":
		return gandi.NewDNSProvider()
	case "gandiv5":
		return gandiv5.NewDNSProvider()
	case "gcloud":
		return gcloud.NewDNSProvider()
	case "gcore":
		return gcore.NewDNSProvider()
	case "glesys":
		return glesys.NewDNSProvider()
	case "godaddy":
		return godaddy.NewDNSProvider()
	case "googledomains":
		return googledomains.NewDNSProvider()
	case "hetzner":
		return hetzner.NewDNSProvider()
	case "hostingde":
		return hostingde.NewDNSProvider()
	case "hosttech":
		return hosttech.NewDNSProvider()
	case "httpreq":
		return httpreq.NewDNSProvider()
	case "hurricane":
		return hurricane.NewDNSProvider()
	case "hyperone":
		return hyperone.NewDNSProvider()
	case "ibmcloud":
		return ibmcloud.NewDNSProvider()
	case "iij":
		return iij.NewDNSProvider()
	case "iijdpf":
		return iijdpf.NewDNSProvider()
	case "infoblox":
		return infoblox.NewDNSProvider()
	case "infomaniak":
		return infomaniak.NewDNSProvider()
	case "internetbs":
		return internetbs.NewDNSProvider()
	case "inwx":
		return inwx.NewDNSProvider()
	case "ionos":
		return ionos.NewDNSProvider()
	case "ipv64":
		return ipv64.NewDNSProvider()
	case "iwantmyname":
		return iwantmyname.NewDNSProvider()
	case "joker":
		return joker.NewDNSProvider()
	case "liara":
		return liara.NewDNSProvider()
	case "lightsail":
		return lightsail.NewDNSProvider()
	case "linode", "linodev4": // "linodev4" is for compatibility with v3, must be dropped in v5
		return linode.NewDNSProvider()
	case "liquidweb":
		return liquidweb.NewDNSProvider()
	case "loopia":
		return loopia.NewDNSProvider()
	case "luadns":
		return luadns.NewDNSProvider()
	case "mailinabox":
		return mailinabox.NewDNSProvider()
	case "manual":
		return dns01.NewDNSProviderManual()
	case "metaname":
		return metaname.NewDNSProvider()
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
	case "nearlyfreespeech":
		return nearlyfreespeech.NewDNSProvider()
	case "netcup":
		return netcup.NewDNSProvider()
	case "netlify":
		return netlify.NewDNSProvider()
	case "nicmanager":
		return nicmanager.NewDNSProvider()
	case "nifcloud":
		return nifcloud.NewDNSProvider()
	case "njalla":
		return njalla.NewDNSProvider()
	case "nodion":
		return nodion.NewDNSProvider()
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
	case "plesk":
		return plesk.NewDNSProvider()
	case "porkbun":
		return porkbun.NewDNSProvider()
	case "rackspace":
		return rackspace.NewDNSProvider()
	case "rcodezero":
		return rcodezero.NewDNSProvider()
	case "regru":
		return regru.NewDNSProvider()
	case "rfc2136":
		return rfc2136.NewDNSProvider()
	case "rimuhosting":
		return rimuhosting.NewDNSProvider()
	case "route53":
		return route53.NewDNSProvider()
	case "safedns":
		return safedns.NewDNSProvider()
	case "sakuracloud":
		return sakuracloud.NewDNSProvider()
	case "scaleway":
		return scaleway.NewDNSProvider()
	case "selectel":
		return selectel.NewDNSProvider()
	case "servercow":
		return servercow.NewDNSProvider()
	case "shellrent":
		return shellrent.NewDNSProvider()
	case "simply":
		return simply.NewDNSProvider()
	case "sonic":
		return sonic.NewDNSProvider()
	case "stackpath":
		return stackpath.NewDNSProvider()
	case "tencentcloud":
		return tencentcloud.NewDNSProvider()
	case "transip":
		return transip.NewDNSProvider()
	case "ultradns":
		return ultradns.NewDNSProvider()
	case "variomedia":
		return variomedia.NewDNSProvider()
	case "vegadns":
		return vegadns.NewDNSProvider()
	case "vercel":
		return vercel.NewDNSProvider()
	case "versio":
		return versio.NewDNSProvider()
	case "vinyldns":
		return vinyldns.NewDNSProvider()
	case "vkcloud":
		return vkcloud.NewDNSProvider()
	case "vscale":
		return vscale.NewDNSProvider()
	case "vultr":
		return vultr.NewDNSProvider()
	case "webnames":
		return webnames.NewDNSProvider()
	case "websupport":
		return websupport.NewDNSProvider()
	case "wedos":
		return wedos.NewDNSProvider()
	case "yandex":
		return yandex.NewDNSProvider()
	case "yandex360":
		return yandex360.NewDNSProvider()
	case "yandexcloud":
		return yandexcloud.NewDNSProvider()
	case "zoneee":
		return zoneee.NewDNSProvider()
	case "zonomi":
		return zonomi.NewDNSProvider()
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
