# Automated Domain Control Validation Renewal

This application allows you to automatically (re)perform domain control validation using the Entrust API and your DNS provider.

**This project is made available by Entrust Corporation under the [MIT license](LICENSE).** As such the application is provided to you AS IS and without warranty, as further detailed in the MIT license.

## Config

| Environment variable            | Value                                                     |
|---------------------------------|-----------------------------------------------------------|
| `DNS_PROVIDER`                  | The name of your default DNS provider                     |
| `ENTRUST_API_CERTIFICATE`       | PEM encoded client certificate                            |
| `ENTRUST_API_PRIVATE_KEY`       | PEM encoded private key                                   |
| `ENTRUST_API_USERNAME`          | API Username                                              |
| `ENTRUST_API_PASSWORD`          | API Password (API Key)                                    |
| `DURATION_BEFORE_EXPIRY`        | Optional, remaining validity before starting revalidation |

Each DNS provider has its own set of environment variables.

## Obtaining credentials

### Entrust

These credentials can be obtained via the Entrust Certificate Services Enterprise Portal (Enterprise Portal, Administration > Advanced Settings > API). In order to obtain an account and related credentials, you must have an active subscription or other entitlement from Entrust for Entrust Certificate Services (ECS).

### DNS Provider(s)

This project is utilizing the DNS providers of the LEGO ACME client; documentation can be found here:
https://go-acme.github.io/lego/dns/#dns-providers

In addition to the default DNS provider specified in `DNS_PROVIDER`, alternative DNS providers can be specified on a per domain basis by setting the environment variable `DNS_PROVIDER_example_com` where the suffix `example_com` stands for `example.com` and can be replaced by any domain in your account.

# Security

Both the Entrust and the DNS provider credentials provide strong entitlements and can have a devastating impact to your organisation or service when not secured properly. It’s your sole responsibility to store these credentials securely and restrict API credentials to the resources and capabilities this tool requires.

# Example

```sh
export ENTRUST_API_CERTIFICATE=`cat cert.pem`
export ENTRUST_API_PRIVATE_KEY=`cat key.pem`
export ENTRUST_API_USERNAME=
export ENTRUST_API_PASSWORD=

# Provide the API credentials for your DNS provider
export CLOUDFLARE_DNS_API_TOKEN=
export CLOUDFLARE_ZONE_API_TOKEN=

docker run --name dvrenewal --rm \
    -e DNS_PROVIDER=cloudflare \
    -e CLOUDFLARE_DNS_API_TOKEN \
    -e CLOUDFLARE_ZONE_API_TOKEN \
    -e ENTRUST_API_CERTIFICATE \
    -e ENTRUST_API_PRIVATE_KEY \
    -e ENTRUST_API_USERNAME \
    -e ENTRUST_API_PASSWORD \
    -e DURATION_BEFORE_EXPIRY \
    ghcr.io/entrustcorporation/dv/dv:1.0.3

```
