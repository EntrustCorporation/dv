module github.com/digitorus/dv

go 1.16

require (
	cloud.google.com/go v0.79.0
	github.com/Azure/azure-sdk-for-go v52.4.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.18
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.7
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/OpenDNS/vegadns2client v0.0.0-20180418235048-a3fa4a771d87
	github.com/akamai/AkamaiOPEN-edgegrid-golang v1.1.0
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.976
	github.com/aws/aws-sdk-go v1.37.28
	github.com/boombuler/barcode v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.1.0
	github.com/cloudflare/cloudflare-go v0.14.0
	github.com/cpu/goacmedns v0.1.1
	github.com/digitorus/entrust v0.0.0-20210311090410-c18455d75837
	github.com/dnsimple/dnsimple-go v0.63.0
	github.com/exoscale/egoscale v1.19.0
	github.com/go-acme/lego/v4 v4.3.0
	github.com/go-errors/errors v1.1.1 // indirect
	github.com/go-resty/resty/v2 v2.5.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-querystring v1.0.0
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/googleapis/gnostic v0.5.4 // indirect
	github.com/gophercloud/gophercloud v0.16.0
	github.com/gophercloud/utils v0.0.0-20210216074907-f6de111f2eae
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/iij/doapi v0.0.0-20190504054126-0bbf12d6d7df
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/kolo/xmlrpc v0.0.0-20201022064351-38db28db192b // indirect
	github.com/labbsr0x/bindman-dns-webhook v1.0.2
	github.com/linode/linodego v1.0.0
	github.com/liquidweb/liquidweb-cli v0.6.10 // indirect
	github.com/liquidweb/liquidweb-go v1.6.3
	github.com/miekg/dns v1.1.40
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/namedotcom/go v0.0.0-20180403034216-08470befbe04
	github.com/nrdcg/auroradns v1.0.1
	github.com/nrdcg/desec v0.5.0
	github.com/nrdcg/dnspod-go v0.4.0
	github.com/nrdcg/goinwx v0.8.1
	github.com/nrdcg/namesilo v0.2.1
	github.com/oracle/oci-go-sdk v24.3.0+incompatible
	github.com/ovh/go-ovh v1.1.0
	github.com/pquerna/otp v1.3.0
	github.com/sacloud/libsacloud v1.36.2
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/transip/gotransip/v6 v6.6.0
	github.com/vultr/govultr/v2 v2.4.0
	go.uber.org/ratelimit v0.2.0 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	golang.org/x/oauth2 v0.0.0-20210220000619-9bb904979d93
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2 // indirect
	golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/api v0.41.0
	gopkg.in/ns1/ns1-go.v2 v2.4.4
	gopkg.in/square/go-jose.v2 v2.5.1
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	k8s.io/client-go v0.20.4 // indirect
	k8s.io/klog/v2 v2.6.0 // indirect
	k8s.io/utils v0.0.0-20210305010621-2afb4311ab10 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.1.0 // indirect
)

// Azure SDK (autorest) depends on an old version witch pushes a depricated messages
replace contrib.go.opencensus.io/exporter/ocagent => contrib.go.opencensus.io/exporter/ocagent v0.7.0
