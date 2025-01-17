package cloudflare

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_access_application",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Access Application resource.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The account to which the access application should be added. Conflicts with ` + "`" + `zone_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The DNS zone to which the access application should be added. Conflicts with ` + "`" + `account_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Friendly name of the Access Application.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The complete URL of the asset you wish to put Cloudflare Access in front of. Can include subdomains or paths. Or both.`,
				},
				resource.Attribute{
					Name:        "session_duration",
					Description: `(Optional) How often a user will be forced to re-authorise. Must be in the format ` + "`" + `"48h"` + "`" + ` or ` + "`" + `"2h45m"` + "`" + `. Valid time units are ` + "`" + `ns` + "`" + `, ` + "`" + `us` + "`" + ` (or ` + "`" + `µs` + "`" + `), ` + "`" + `ms` + "`" + `, ` + "`" + `s` + "`" + `, ` + "`" + `m` + "`" + `, ` + "`" + `h` + "`" + `. Defaults to ` + "`" + `24h` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cors_headers",
					Description: `(Optional) CORS configuration for the Access Application. See below for reference structure.`,
				},
				resource.Attribute{
					Name:        "allowed_idps",
					Description: `(Optional) The identity providers selected for the application.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `(Optional) List of methods to expose via CORS.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `(Optional) List of origins permitted to make CORS requests.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `(Optional) List of HTTP headers to expose via CORS.`,
				},
				resource.Attribute{
					Name:        "allow_all_methods",
					Description: `(Optional) Boolean value to determine whether all methods are exposed.`,
				},
				resource.Attribute{
					Name:        "allow_all_origins",
					Description: `(Optional) Boolean value to determine whether all origins are permitted to make CORS requests.`,
				},
				resource.Attribute{
					Name:        "allow_all_headers",
					Description: `(Optional) Boolean value to determine whether all HTTP headers are exposed.`,
				},
				resource.Attribute{
					Name:        "allow_credentials",
					Description: `(Optional) Boolean value to determine if credentials (cookies, authorization headers, or TLS client certificates) are included with requests.`,
				},
				resource.Attribute{
					Name:        "max_age",
					Description: `(Optional) Integer representing the maximum time a preflight request will be cached.`,
				},
				resource.Attribute{
					Name:        "auto_redirect_to_identity",
					Description: `(Optional) Option to skip identity provider selection if only one is configured in allowed_idps. Defaults to ` + "`" + `false` + "`" + ` (disabled).`,
				},
				resource.Attribute{
					Name:        "enable_binding_cookie",
					Description: `(Optional) Option to provide increased security against compromised authorization tokens and CSRF attacks by requiring an additional "binding" cookie on requests. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_deny_message",
					Description: `(Optional) Option that returns a custom error message when a user is denied access to the application.`,
				},
				resource.Attribute{
					Name:        "custom_deny_url",
					Description: `(Optional) Option that redirects to a custom URL when a user is denied access to the application. ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the application`,
				},
				resource.Attribute{
					Name:        "aud",
					Description: `Application Audience (AUD) Tag of the application`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the application`,
				},
				resource.Attribute{
					Name:        "session_duration",
					Description: `Length of session for the application before prompting for a sign in`,
				},
				resource.Attribute{
					Name:        "auto_redirect_to_identity",
					Description: `If the IdP selection page is skipped or not`,
				},
				resource.Attribute{
					Name:        "allowed_idps",
					Description: `The identity providers selected for the application ## Import Access Applications can be imported using a composite ID formed of account ID and application ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_access_application.staging cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the application`,
				},
				resource.Attribute{
					Name:        "aud",
					Description: `Application Audience (AUD) Tag of the application`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the application`,
				},
				resource.Attribute{
					Name:        "session_duration",
					Description: `Length of session for the application before prompting for a sign in`,
				},
				resource.Attribute{
					Name:        "auto_redirect_to_identity",
					Description: `If the IdP selection page is skipped or not`,
				},
				resource.Attribute{
					Name:        "allowed_idps",
					Description: `The identity providers selected for the application ## Import Access Applications can be imported using a composite ID formed of account ID and application ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_access_application.staging cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_access_ca_certificate",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Access CA certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"ca",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The account to which the Access CA certificate should be added. Conflicts with ` + "`" + `zone_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The DNS zone to which the Access CA certificate should be added. Conflicts with ` + "`" + `account_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required) The Access Application ID to associate with the CA certificate. ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the CA certificate`,
				},
				resource.Attribute{
					Name:        "aud",
					Description: `Application Audience (AUD) Tag of the CA certificate`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Cryptographic public key of the generated CA certificate ## Import Access CA certificates can be imported using a composite ID formed of identifer ("account" or "zone"), identifier ID and the CA certificate ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_access_ca_certificate.example account/1d5fdc9e88c8a8c4518b068cd94331fe/edc1e4e24567217764b4322669c44df985dddffdf03ac781 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the CA certificate`,
				},
				resource.Attribute{
					Name:        "aud",
					Description: `Application Audience (AUD) Tag of the CA certificate`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Cryptographic public key of the generated CA certificate ## Import Access CA certificates can be imported using a composite ID formed of identifer ("account" or "zone"), identifier ID and the CA certificate ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_access_ca_certificate.example account/1d5fdc9e88c8a8c4518b068cd94331fe/edc1e4e24567217764b4322669c44df985dddffdf03ac781 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_access_group",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Access Group resource.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The ID of the account the group is associated with. Conflicts with ` + "`" + `zone_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The ID of the zone the group is associated with. Conflicts with ` + "`" + `account_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Friendly name of the Access Group.`,
				},
				resource.Attribute{
					Name:        "require",
					Description: `(Optional) A series of access conditions, see below for full list.`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(Optional) A series of access conditions, see below for full list.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Required) A series of access conditions, see below for full list. ## Conditions ` + "`" + `require` + "`" + `, ` + "`" + `exclude` + "`" + ` and ` + "`" + `include` + "`" + ` arguments share the available conditions which can be applied. The conditions are:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) A list of IP addresses or ranges. Example: ` + "`" + `ip = ["1.2.3.4", "10.0.0.0/2"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) A list of email addresses. Example: ` + "`" + `email = ["test@example.com"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "email_domain",
					Description: `(Optional) A list of email domains. Example: ` + "`" + `email_domain = ["example.com"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "service_token",
					Description: `(Optional) A list of service token ids. Example: ` + "`" + `service_token = [cloudflare_access_service_token.demo.id]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "any_valid_service_token",
					Description: `(Optional) Boolean indicating if allow all tokens to be granted. Example: ` + "`" + `any_valid_service_token = true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) A list of access group ids. Example: ` + "`" + `group = [cloudflare_access_group.demo.id]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "everyone",
					Description: `(Optional) Boolean indicating permitting access for all requests. Example: ` + "`" + `everyone = true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional) Whether to use mTLS certificate authentication.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Optional) Use a certificate common name to authenticate with.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `(Optional) A string identifying the authentication method code. The list of codes are listed here: https://tools.ietf.org/html/rfc8176#section-2. Custom values are also supported.`,
				},
				resource.Attribute{
					Name:        "geo",
					Description: `(Optional) A list of country codes. Example: ` + "`" + `geo = ["US"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gsuite",
					Description: `(Optional) Use GSuite as the authentication mechanism. Example: ` + "`" + `` + "`" + `` + "`" + `hcl # ... other configuration include { gsuite { email = ["admins@example.com"] identity_provider_id = "ca298b82-93b5-41bf-bc2d-10493f09b761" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "github",
					Description: `(Optional) Use a GitHub organization as the ` + "`" + `include` + "`" + ` condition. Example: ` + "`" + `` + "`" + `` + "`" + `hcl # ... other configuration include { github { name = "my-github-org-name" # (Required) GitHub organization name team = ["my-github-team-name"] # (Optional) GitHub teams identity_provider_id = "ca298b82-93b5-41bf-bc2d-10493f09b761" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "azure",
					Description: `(Optional) Use Azure AD as the ` + "`" + `include` + "`" + ` condition. Example: ` + "`" + `` + "`" + `` + "`" + `hcl # ... other configuration include { azure { id = ["86773093-5feb-48dd-814b-7ccd3676ff50e"] identity_provider_id = "ca298b82-93b5-41bf-bc2d-10493f09b761" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "okta",
					Description: `(Optional) Use Okta as the ` + "`" + `include` + "`" + ` condition. Example: ` + "`" + `` + "`" + `` + "`" + `hcl # ... other configuration include { okta { name = ["admins"] identity_provider_id = "ca298b82-93b5-41bf-bc2d-10493f09b761" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "saml",
					Description: `(Optional) Use an external SAML setup as the ` + "`" + `include` + "`" + ` condition. Example: ` + "`" + `` + "`" + `` + "`" + `hcl # ... other configuration include { saml { attribute_name = "group" attribute_value = "admins" identity_provider_id = "ca298b82-93b5-41bf-bc2d-10493f09b761" } } ` + "`" + `` + "`" + `` + "`" + ` ## Import Access Groups can be imported using a composite ID formed of account ID and group ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_access_group.staging 975ecf5a45e3bcb680dba0722a420ad9/67ea780ce4982c1cfbe6b7293afc765d ` + "`" + `` + "`" + `` + "`" + ` where`,
				},
				resource.Attribute{
					Name:        "975ecf5a45e3bcb680dba0722a420ad9",
					Description: `Account ID`,
				},
				resource.Attribute{
					Name:        "67ea780ce4982c1cfbe6b7293afc765d",
					Description: `Access Group ID`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_access_identity_provider",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Access Identity Provider resource.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"identity",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The account ID the provider should be associated with. Conflicts with ` + "`" + `zone_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The zone ID the provider should be associated with. Conflicts with ` + "`" + `account_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Friendly name of the Access Identity Provider configuration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The provider type to use. Must be one of: ` + "`" + `"centrify"` + "`" + `, ` + "`" + `"facebook"` + "`" + `, ` + "`" + `"google-apps"` + "`" + `, ` + "`" + `"oidc"` + "`" + `, ` + "`" + `"github"` + "`" + `, ` + "`" + `"google"` + "`" + `, ` + "`" + `"saml"` + "`" + `, ` + "`" + `"linkedin"` + "`" + `, ` + "`" + `"azureAD"` + "`" + `, ` + "`" + `"okta"` + "`" + `, ` + "`" + `"onetimepin"` + "`" + `, ` + "`" + `"onelogin"` + "`" + `, ` + "`" + `"yandex"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) Provider configuration from the [developer documentation][access_identity_provider_guide]. ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Access Identity Provider`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Friendly name of the Access Identity Provider configuration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The provider type to use.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Access Identity Provider configuration. ## Import Access Identity Providers can be imported using a composite ID formed of account ID and Access Identity Provider ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_access_identity_provider.my_idp cb029e245cfdd66dc8d2e570d5dd3322/e00e1c13-e350-44fe-96c5-fb75c954871c ` + "`" + `` + "`" + `` + "`" + ` [access_identity_provider_guide]: https://developers.cloudflare.com/access/configuring-identity-providers/`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Access Identity Provider`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Friendly name of the Access Identity Provider configuration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The provider type to use.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Access Identity Provider configuration. ## Import Access Identity Providers can be imported using a composite ID formed of account ID and Access Identity Provider ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_access_identity_provider.my_idp cb029e245cfdd66dc8d2e570d5dd3322/e00e1c13-e350-44fe-96c5-fb75c954871c ` + "`" + `` + "`" + `` + "`" + ` [access_identity_provider_guide]: https://developers.cloudflare.com/access/configuring-identity-providers/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_access_policy",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Access Policy resource.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required) The ID of the application the policy is associated with.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The account to which the access rule should be added. Conflicts with ` + "`" + `zone_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The DNS zone to which the access rule should be added. Conflicts with ` + "`" + `account_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "decision",
					Description: `(Required) Defines the action Access will take if the policy matches the user. Allowed values: ` + "`" + `allow` + "`" + `, ` + "`" + `deny` + "`" + `, ` + "`" + `non_identity` + "`" + `, ` + "`" + `bypass` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Friendly name of the Access Application.`,
				},
				resource.Attribute{
					Name:        "precedence",
					Description: `(Required) The unique precedence for policies on a single application. Integer.`,
				},
				resource.Attribute{
					Name:        "require",
					Description: `(Optional) A series of access conditions, see [Access Groups](/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(Optional) A series of access conditions, see [Access Groups](/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Required) A series of access conditions, see [Access Groups](/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions). ## Import Access Policies can be imported using a composite ID formed of identifier type (` + "`" + `zone` + "`" + ` or ` + "`" + `account` + "`" + `), identifier ID (` + "`" + `zone_id` + "`" + ` or ` + "`" + `account_id` + "`" + `), application ID and policy ID. ` + "`" + `` + "`" + `` + "`" + ` # import a zone level Access policy $ terraform import cloudflare_access_policy.staging zone/cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e/67ea780ce4982c1cfbe6b7293afc765d # import an account level Access policy $ terraform import cloudflare_access_policy.production account/0d599f0ec05c3bda8c3b8a68c32a1b47/d41d8cd98f00b204e9800998ecf8427e/67ea780ce4982c1cfbe6b7293afc765d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_access_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare IP Firewall Access Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The DNS zone to which the access rule should be added.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The action to apply to a matched request. Allowed values: "block", "challenge", "whitelist", "js_challenge"`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) A personal note about the rule. Typically used as a reminder or explanation for the rule.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Required) Rule configuration to apply to a matched request. It's a complex value. See description below.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The request property to target. Allowed values: "ip", "ip6", "ip_range", "asn", "country"`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to target. Depends on target's type. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The access rule ID.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The DNS zone ID. ## Import Records can be imported using a composite ID formed of access rule type, access rule type identifier and identifer value, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_access_rule.default zone/cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `access rule type (` + "`" + `account` + "`" + `, ` + "`" + `zone` + "`" + ` or ` + "`" + `user` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "cb029e245cfdd66dc8d2e570d5dd3322",
					Description: `access rule type ID (i.e the zone ID or account ID you wish to target)`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `access rule ID as returned by respective API endpoint for the type you are attempting to import.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The access rule ID.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The DNS zone ID. ## Import Records can be imported using a composite ID formed of access rule type, access rule type identifier and identifer value, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_access_rule.default zone/cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `access rule type (` + "`" + `account` + "`" + `, ` + "`" + `zone` + "`" + ` or ` + "`" + `user` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "cb029e245cfdd66dc8d2e570d5dd3322",
					Description: `access rule type ID (i.e the zone ID or account ID you wish to target)`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `access rule ID as returned by respective API endpoint for the type you are attempting to import.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_access_service_token",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Access Service Token resource.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"service",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The ID of the account where the Access Service is being created. Conflicts with ` + "`" + `zone_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The ID of the zone where the Access Service is being created. Conflicts with ` + "`" + `account_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Friendly name of the token's intent. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `UUID client ID associated with the Service Token.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `A secret for interacting with Access protocols. ## Import ~>`,
				},
				resource.Attribute{
					Name:        "cb029e245cfdd66dc8d2e570d5dd3322",
					Description: `Account ID`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `Access Service Token ID`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `UUID client ID associated with the Service Token.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `A secret for interacting with Access protocols. ## Import ~>`,
				},
				resource.Attribute{
					Name:        "cb029e245cfdd66dc8d2e570d5dd3322",
					Description: `Account ID`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `Access Service Token ID`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_account_member",
			Category:         "Resources",
			ShortDescription: `Provides a resource which manages Cloudflare account members.`,
			Description:      ``,
			Keywords: []string{
				"account",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email_address",
					Description: `(Required) The email address of the user who you wish to manage. Note: Following creation, this field becomes read only via the API and cannot be updated.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `(Required) Array of account role IDs that you want to assign to a member. ## Import Account members can be imported using a composite ID formed of account ID and account member ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_account_member.example_user d41d8cd98f00b204e9800998ecf8427e/b58c6f14d292556214bd64909bcdb118 ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `account ID as returned by the [API](https://api.cloudflare.com/#accounts-account-details)`,
				},
				resource.Attribute{
					Name:        "b58c6f14d292556214bd64909bcdb118",
					Description: `account member ID as returned by the [API](https://api.cloudflare.com/#account-members-member-details)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_api_token",
			Category:         "Resources",
			ShortDescription: `Provides a resource which manages Cloudflare API tokens.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the APIToken.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) Permissions policy. Multiple policy blocks can be defined. See the definition below.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) Condition block. See the definition below. The`,
				},
				resource.Attribute{
					Name:        "permission_groups",
					Description: `(Required) List of permissions groups ids ([see official docs][1]).`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Map describes what operations against which resources are allowed or denied.`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Optional) Policy effect. Valid values are ` + "`" + `allow` + "`" + ` or ` + "`" + `deny` + "`" + `. ` + "`" + `allow` + "`" + ` is set as default. The`,
				},
				resource.Attribute{
					Name:        "request_ip",
					Description: `(Optional) Request IP related conditions. See the definition below. The`,
				},
				resource.Attribute{
					Name:        "in",
					Description: `(Optional) List of IPv4/IPv6 CIDR addresses where the Token can be used from.`,
				},
				resource.Attribute{
					Name:        "not_in",
					Description: `(Optional) List of IPv4/IPv6 CIDR addresses where the Token cannot be used from. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the API Token.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the API Token.`,
				},
				resource.Attribute{
					Name:        "issued_on",
					Description: `The RFC3339 timestamp of when the API Token was issued.`,
				},
				resource.Attribute{
					Name:        "modified_on",
					Description: `The RFC3339 timestamp of when the API Token was last modified. [1]: https://developers.cloudflare.com/api/tokens/create/permissions`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the API Token.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the API Token.`,
				},
				resource.Attribute{
					Name:        "issued_on",
					Description: `The RFC3339 timestamp of when the API Token was issued.`,
				},
				resource.Attribute{
					Name:        "modified_on",
					Description: `The RFC3339 timestamp of when the API Token was last modified. [1]: https://developers.cloudflare.com/api/tokens/create/permissions`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_argo",
			Category:         "Resources",
			ShortDescription: `Provides the ability to manage Cloudflare Argo features.`,
			Description:      ``,
			Keywords: []string{
				"argo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID that you wish to manage Argo on.`,
				},
				resource.Attribute{
					Name:        "tiered_caching",
					Description: `(Optional) Whether tiered caching is enabled. Valid values: ` + "`" + `on` + "`" + ` or ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "smart_routing",
					Description: `(Optional) Whether smart routing is enabled. Valid values: ` + "`" + `on` + "`" + ` or ` + "`" + `off` + "`" + `. ## Import Argo settings can be imported the zone ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_argo.example d41d8cd98f00b204e9800998ecf8427e ` + "`" + `` + "`" + `` + "`" + ` where ` + "`" + `d41d8cd98f00b204e9800998ecf8427e` + "`" + ` is the zone ID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_argo_tunnel",
			Category:         "Resources",
			ShortDescription: `Provides the ability to manage Cloudflare Argo Tunnels.`,
			Description:      ``,
			Keywords: []string{
				"argo",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The Cloudflare account ID that you wish to manage the Argo Tunnel on.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A user-friendly name chosen when the tunnel is created. Cannot be empty.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Required) 32 or more bytes, encoded as a base64 string. The Create Argo Tunnel endpoint sets this as the tunnel's password. Anyone wishing to run the tunnel needs this password. ## Import Argo Tunnels can be imported a composite ID of the account ID and tunnel UUID. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_byo_ip_prefix",
			Category:         "Resources",
			ShortDescription: `Provides the ability to manage Bring-Your-Own-IP prefixes (BYOIP) which are used with or without Magic Transit.`,
			Description:      ``,
			Keywords: []string{
				"byo",
				"ip",
				"prefix",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "prefix_id",
					Description: `(Required) The assigned Bring-Your-Own-IP prefix ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the prefix.`,
				},
				resource.Attribute{
					Name:        "advertisement",
					Description: `(Optional) Whether or not the prefix shall be announced. A prefix can be activated or deactivated once every 15 minutes (attempting more regular updates will trigger rate limiting). Valid values: ` + "`" + `on` + "`" + ` or ` + "`" + `off` + "`" + `. ## Import The current settings for Bring-Your-Own-IP prefixes can be imported using the prefix ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_byo_ip_prefix.example d41d8cd98f00b204e9800998ecf8427e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_certificate_pack",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Certificate Pack resource.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
				"pack",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone to which the certificate pack should be added.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Certificate pack configuration type. Allowed values: ` + "`" + `"custom"` + "`" + `, ` + "`" + `"dedicated_custom"` + "`" + `, ` + "`" + `"advanced"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `(Required) List of hostnames to provision the certificate pack for. Note: If using Let's Encrypt, you cannot use individual subdomains and only a wildcard for subdomain is available.`,
				},
				resource.Attribute{
					Name:        "validation_method",
					Description: `(Optional based on ` + "`" + `type` + "`" + `) Which validation method to use in order to prove domain ownership. Allowed values: ` + "`" + `"txt"` + "`" + `, ` + "`" + `"http"` + "`" + `, ` + "`" + `"email"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "validity_days",
					Description: `(Optional based on ` + "`" + `type` + "`" + `) How long the certificate is valid for. Note: If using Let's Encrypt, this value can only be 90 days. Allowed values: 14, 30, 90, 365.`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Optional based on ` + "`" + `type` + "`" + `) Which certificate authority to issue the certificate pack. Allowed values: ` + "`" + `"digicert"` + "`" + `, ` + "`" + `"lets_encrypt"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cloudflare_branding",
					Description: `(Optional based on ` + "`" + `type` + "`" + `) Whether or not to include Cloudflare branding. This will add ` + "`" + `sni.cloudflaressl.com` + "`" + ` as the Common Name if set to ` + "`" + `true` + "`" + `. ## Import Certificate packs can be imported using a composite ID of the zone ID and certificate pack ID. This isn't recommended and it is advised to replace the certificate entirely instead. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_certificate_pack.example cb029e245cfdd66dc8d2e570d5dd3322/8fda82e2-6af9-4eb2-992a-5ab65b792ef1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_custom_hostname",
			Category:         "Resources",
			ShortDescription: `!- Provides a Cloudflare custom hostname resource.`,
			Description:      ``,
			Keywords: []string{
				"custom",
				"hostname",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID where the custom hostname should be assigned.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Hostname you intend to request a certificate for.`,
				},
				resource.Attribute{
					Name:        "custom_origin_server",
					Description: `(Optional) The custom origin server used for certificates.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(Required) SSL configuration of the certificate. See further notes below.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) Domain control validation (DCV) method used for this hostname. Valid values are ` + "`" + `"txt"` + "`" + `, ` + "`" + `"http"` + "`" + ` and ` + "`" + `"email"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Level of validation to be used for this hostname. Domain validation ("dv") must be used.`,
				},
				resource.Attribute{
					Name:        "wildcard",
					Description: `(Required) Indicates whether the certificate covers a wildcard.`,
				},
				resource.Attribute{
					Name:        "custom_certificate",
					Description: `(Optional) If a custom uploaded certificate is used.`,
				},
				resource.Attribute{
					Name:        "custom_key",
					Description: `(Optional) The key for a custom uploaded certificate.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Required) SSL/TLS settings for the certificate. See further notes below.`,
				},
				resource.Attribute{
					Name:        "http2",
					Description: `(Optional) Whether or not HTTP2 should be supported. Valid values are ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls13",
					Description: `(Optional) Whether or not TLSv1.3 should be supported. Valid values are ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `(Optional) Lowest version of TLS this certificate should support. Valid values are ` + "`" + `"1.0"` + "`" + `, ` + "`" + `"1.1"` + "`" + `, ` + "`" + `"1.2"` + "`" + ` and ` + "`" + `"1.3"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) List of SSL/TLS ciphers to associate with this certificate. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ownership_verification.type",
					Description: `Domain control validation (DCV) method used for the hostname.`,
				},
				resource.Attribute{
					Name:        "ownership_verification.value",
					Description: `Domain control validation (DCV) value for confirming ownership. Example, "_cf-custom-hostname.example.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ownership_verification.name",
					Description: `Domain control validation (DCV) name confirming ownership. Example, "03f28e11-fa64-4966-bb1e-dd2423e16f36"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ownership_verification_http.http_url",
					Description: `Domain control validation (DCV) URL for confirming ownership. Example, ` + "`" + `http://hostname.example.com/.well-known/cf-custom-hostname-challenge/643395f9-de80-42f5-a2a0-e03ff60cf2a7` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ownership_verification_http.http_body",
					Description: `Domain control validation (DCV) body for confirming ownership. Example, ` + "`" + `03f28e11-fa64-4966-bb1e-dd2423e16f36` + "`" + ` ## Import Custom hostname certificates can be imported using a composite ID formed of the zone ID and [hostname ID](https://api.cloudflare.com/#custom-hostname-for-a-zone-properties), separated by a "/" e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_custom_hostname.example d41d8cd98f00b204e9800998ecf8427e/0d89c70d-ad9f-4843-b99f-6cc0252067e9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ownership_verification.type",
					Description: `Domain control validation (DCV) method used for the hostname.`,
				},
				resource.Attribute{
					Name:        "ownership_verification.value",
					Description: `Domain control validation (DCV) value for confirming ownership. Example, "_cf-custom-hostname.example.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ownership_verification.name",
					Description: `Domain control validation (DCV) name confirming ownership. Example, "03f28e11-fa64-4966-bb1e-dd2423e16f36"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ownership_verification_http.http_url",
					Description: `Domain control validation (DCV) URL for confirming ownership. Example, ` + "`" + `http://hostname.example.com/.well-known/cf-custom-hostname-challenge/643395f9-de80-42f5-a2a0-e03ff60cf2a7` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ownership_verification_http.http_body",
					Description: `Domain control validation (DCV) body for confirming ownership. Example, ` + "`" + `03f28e11-fa64-4966-bb1e-dd2423e16f36` + "`" + ` ## Import Custom hostname certificates can be imported using a composite ID formed of the zone ID and [hostname ID](https://api.cloudflare.com/#custom-hostname-for-a-zone-properties), separated by a "/" e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_custom_hostname.example d41d8cd98f00b204e9800998ecf8427e/0d89c70d-ad9f-4843-b99f-6cc0252067e9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_custom_pages",
			Category:         "Resources",
			ShortDescription: `Provides a resource which manages Cloudflare custom pages.`,
			Description:      ``,
			Keywords: []string{
				"custom",
				"pages",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The zone ID where the custom pages should be updated. Either ` + "`" + `zone_id` + "`" + ` or ` + "`" + `account_id` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The account ID where the custom pages should be updated. Either ` + "`" + `account_id` + "`" + ` or ` + "`" + `zone_id` + "`" + ` must be provided. If ` + "`" + `account_id` + "`" + ` is present, it will override the zone setting.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of custom page you wish to update. Must be one of ` + "`" + `basic_challenge` + "`" + `, ` + "`" + `waf_challenge` + "`" + `, ` + "`" + `waf_block` + "`" + `, ` + "`" + `ratelimit_block` + "`" + `, ` + "`" + `country_challenge` + "`" + `, ` + "`" + `ip_block` + "`" + `, ` + "`" + `under_attack` + "`" + `, ` + "`" + `500_errors` + "`" + `, ` + "`" + `1000_errors` + "`" + `, ` + "`" + `always_online` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) URL of where the custom page source is located.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Required) Managed state of the custom page. Must be one of ` + "`" + `default` + "`" + `, ` + "`" + `customized` + "`" + `. If the value is ` + "`" + `default` + "`" + ` it will be removed from the Terraform state management. ## Import Custom pages can be imported using a composite ID formed of:`,
				},
				resource.Attribute{
					Name:        "customPageLevel",
					Description: `Either ` + "`" + `account` + "`" + ` or ` + "`" + `zone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `The ID of the account or zone you intend to manage.`,
				},
				resource.Attribute{
					Name:        "pageType",
					Description: `The value from the ` + "`" + `type` + "`" + ` argument. Example for a zone: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_custom_pages.basic_challenge zone/d41d8cd98f00b204e9800998ecf8427e/basic_challenge ` + "`" + `` + "`" + `` + "`" + ` Example for an account: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_custom_pages.basic_challenge account/e268443e43d93dab7ebef303bbe9642f/basic_challenge ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_custom_ssl",
			Category:         "Resources",
			ShortDescription: `!- Provides a Cloudflare custom ssl resource.`,
			Description:      ``,
			Keywords: []string{
				"custom",
				"ssl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone id to the custom ssl cert should be added.`,
				},
				resource.Attribute{
					Name:        "custom_ssl_options",
					Description: `(Required) The certificate, private key and associated optional parameters, such as bundle_method, geo_restrictions, and type.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) Certificate certificate and the intermediate(s)`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) Certificate's private key`,
				},
				resource.Attribute{
					Name:        "bundle_method",
					Description: `(Optional) Method of building intermediate certificate chain. A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it. Valid values are ` + "`" + `ubiquitous` + "`" + ` (default), ` + "`" + `optimal` + "`" + `, ` + "`" + `force` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "geo_restrictions",
					Description: `(Optional) Specifies the region where your private key can be held locally. Valid values are ` + "`" + `us` + "`" + `, ` + "`" + `eu` + "`" + `, ` + "`" + `highest_security` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Whether to enable support for legacy clients which do not include SNI in the TLS handshake. Valid values are ` + "`" + `legacy_custom` + "`" + ` (default), ` + "`" + `sni_custom` + "`" + `. ## Import Custom SSL Certs can be imported using a composite ID formed of the zone ID and [certificate ID](https://api.cloudflare.com/#custom-ssl-for-a-zone-properties), separated by a "/" e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_custom_ssl.default 1d5fdc9e88c8a8c4518b068cd94331fe/0123f0ab-9cde-45b2-80bd-4da3010f1337 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_filter",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Filter expression that can be referenced across multiple features.`,
			Description:      ``,
			Keywords: []string{
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone to which the Filter should be added.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `(Optional) Whether this filter is currently paused. Boolean value.`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) The filter expression to be used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A note that you can use to describe the purpose of the filter.`,
				},
				resource.Attribute{
					Name:        "ref",
					Description: `(Optional) Short reference tag to quickly select related rules. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Filter identifier. ## Import Filter can be imported using a composite ID formed of zone ID and filter ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_filter.default d41d8cd98f00b204e9800998ecf8427e/9e107d9d372bb6826bd81d3542a419d6 ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID`,
				},
				resource.Attribute{
					Name:        "9e107d9d372bb6826bd81d3542a419d6",
					Description: `filter ID as returned by [API](https://api.cloudflare.com/#zone-firewall-filters)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Filter identifier. ## Import Filter can be imported using a composite ID formed of zone ID and filter ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_filter.default d41d8cd98f00b204e9800998ecf8427e/9e107d9d372bb6826bd81d3542a419d6 ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID`,
				},
				resource.Attribute{
					Name:        "9e107d9d372bb6826bd81d3542a419d6",
					Description: `filter ID as returned by [API](https://api.cloudflare.com/#zone-firewall-filters)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_firewall_rule",
			Category:         "Resources",
			ShortDescription: `Define Firewall rule using filter expression for more control over how traffic is matched to the rule.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone to which the Filter should be added.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to apply to a matched request. Allowed values: "block", "challenge", "allow", "js_challenge", "bypass". Enterprise plan also allows "log".`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the rule to allow control of processing order. A lower number indicates high priority. If not provided, any rules with a priority will be sequenced before those without.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `(Optional) Whether this filter based firewall rule is currently paused. Boolean value.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the rule to help identify it.`,
				},
				resource.Attribute{
					Name:        "products",
					Description: `(Optional) List of products to bypass for a request when the bypass action is used. Allowed values: "zoneLockdown", "uaBlock", "bic", "hot", "securityLevel", "rateLimit", "waf". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Firewall Rule identifier. ## Import Firewall Rule can be imported using a composite ID formed of zone ID and rule ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_firewall_rule.default d41d8cd98f00b204e9800998ecf8427e/9e107d9d372bb6826bd81d3542a419d6 ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID`,
				},
				resource.Attribute{
					Name:        "9e107d9d372bb6826bd81d3542a419d6",
					Description: `rule ID as returned by [API](https://api.cloudflare.com/#zone-firewall-filter-rules)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Firewall Rule identifier. ## Import Firewall Rule can be imported using a composite ID formed of zone ID and rule ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_firewall_rule.default d41d8cd98f00b204e9800998ecf8427e/9e107d9d372bb6826bd81d3542a419d6 ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID`,
				},
				resource.Attribute{
					Name:        "9e107d9d372bb6826bd81d3542a419d6",
					Description: `rule ID as returned by [API](https://api.cloudflare.com/#zone-firewall-filter-rules)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_healthcheck",
			Category:         "Resources",
			ShortDescription: `Provides the ability to create a Standalone Health Check without needing a Cloudflare Load Balancer.`,
			Description:      ``,
			Keywords: []string{
				"healthcheck",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID to which apply settings.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A short name to identify the health check. Only alphanumeric characters, hyphens and underscores are allowed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-readable description of the health check.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The hostname or IP address of the origin server to run health checks on.`,
				},
				resource.Attribute{
					Name:        "suspended",
					Description: `(Optional) If suspended, no health checks are sent to the origin. Valid values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` (Default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "check_regions",
					Description: `(Optional) A list of regions from which to run health checks. If not set Cloudflare will pick a default region. Valid values: ` + "`" + `WNAM` + "`" + `, ` + "`" + `ENAM` + "`" + `, ` + "`" + `WEU` + "`" + `, ` + "`" + `EEU` + "`" + `, ` + "`" + `NSAM` + "`" + `, ` + "`" + `SSAM` + "`" + `, ` + "`" + `OC` + "`" + `, ` + "`" + `ME` + "`" + `, ` + "`" + `NAF` + "`" + `, ` + "`" + `SAF` + "`" + `, ` + "`" + `IN` + "`" + `, ` + "`" + `SEAS` + "`" + `, ` + "`" + `NEAS` + "`" + `, ` + "`" + `ALL_REGIONS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notification_suspended",
					Description: `(Optional) Whether the notifications are suspended or not. Useful for maintenance periods. Valid values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` (Default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "notification_email_addresses",
					Description: `(Optional) A list of email addresses we want to send the notifications to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The protocol to use for the health check. Valid values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `TCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port number to connect to for the health check. Valid values are in the range ` + "`" + `0-65535` + "`" + ` (Default: ` + "`" + `80` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout (in seconds) before marking the health check as failed. (Default: ` + "`" + `5` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately. (Default: ` + "`" + `2` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The interval between each health check. Shorter intervals may give quicker notifications if the origin status changes, but will increase load on the origin as we check from multiple locations. (Default: ` + "`" + `60` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "consecutive_fails",
					Description: `(Optional) The number of consecutive fails required from a health check before changing the health to unhealthy. (Default: ` + "`" + `1` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "consecutive_successes",
					Description: `(Optional) The number of consecutive successes required from a health check before changing the health to healthy. (Default: ` + "`" + `1` + "`" + `) ### HTTP/HTTPS specific arguments`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The HTTP method to use for the health check. Valid values: ` + "`" + `GET` + "`" + ` or ` + "`" + `HEAD` + "`" + ` (Default: ` + "`" + `GET` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The endpoint path to health check against. (Default: ` + "`" + `/` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "expected_body",
					Description: `(Optional) A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `(Optional) The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all codes starting with 2) of the health check. (Default: ` + "`" + `["200"]` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "follow_redirects",
					Description: `(Optional) Follow redirects if the origin returns a 3xx status code. Valid values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` (Default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_insecure",
					Description: `(Optional) Do not validate the certificate when the health check uses HTTPS. Valid values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` (Default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Optional) The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Required) The header name.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of string values for the header. ### TCP specific arguments`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The TCP connection method to use for the health check. Valid values: ` + "`" + `connection_established` + "`" + ` (Default: ` + "`" + `connection_established` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_ip_list",
			Category:         "Resources",
			ShortDescription: `Provides IP Lists to be used in Firewall Rules across all zones within the same account.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The ID of the account where the IP List is being created.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the list (used in filter expressions). Valid pattern: ` + "`" + `^[a-zA-Z0-9_]+$` + "`" + `. Maximum Length: 50`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind of values in the List. Valid values: ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A note that can be used to annotate the List. Maximum Length: 500 The`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The IPv4 address, IPv4 CIDR or IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) A note that can be used to annotate the item. ## Import An existing IP List can be imported using the account ID and list ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_ip_list.example d41d8cd98f00b204e9800998ecf8427e/cb029e245cfdd66dc8d2e570d5dd3322 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_load_balancer",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone ID to add the load balancer to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The DNS name (FQDN, including the zone) to associate with the load balancer.`,
				},
				resource.Attribute{
					Name:        "fallback_pool_id",
					Description: `(Required) The pool ID to use when all other pools are detected as unhealthy.`,
				},
				resource.Attribute{
					Name:        "default_pool_ids",
					Description: `(Required) A list of pool IDs ordered by their failover priority. Used whenever region/pop pools are not defined.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Free text description.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live (TTL) of this load balancer's DNS ` + "`" + `name` + "`" + `. Conflicts with ` + "`" + `proxied` + "`" + ` - this cannot be set for proxied load balancers. Default is ` + "`" + `30` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "steering_policy",
					Description: `(Optional) Determine which method the load balancer uses to determine the fastest route to your origin. Valid values are: ` + "`" + `"off"` + "`" + `, ` + "`" + `"geo"` + "`" + `, ` + "`" + `"dynamic_latency"` + "`" + `, ` + "`" + `"random"` + "`" + ` or ` + "`" + `""` + "`" + `. Default is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxied",
					Description: `(Optional) Whether the hostname gets Cloudflare's origin protection. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable or disable the load balancer. Defaults to ` + "`" + `true` + "`" + ` (enabled).`,
				},
				resource.Attribute{
					Name:        "region_pools",
					Description: `(Optional) A set containing mappings of region/country codes to a list of pool IDs (ordered by their failover priority) for the given region. Fields documented below.`,
				},
				resource.Attribute{
					Name:        "pop_pools",
					Description: `(Optional) A set containing mappings of Cloudflare Point-of-Presence (PoP) identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). This feature is only available to enterprise customers. Fields documented below.`,
				},
				resource.Attribute{
					Name:        "session_affinity",
					Description: `(Optional) Associates all requests coming from an end-user with a single origin. Cloudflare will set a cookie on the initial response to the client, such that consequent requests with the cookie in the request will go to the same origin, so long as it is available. Valid values are: ` + "`" + `""` + "`" + `, ` + "`" + `"none"` + "`" + `, ` + "`" + `"cookie"` + "`" + `, and ` + "`" + `"ip_cookie"` + "`" + `. Default is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "session_affinity_ttl",
					Description: `(Optional) Time, in seconds, until this load balancers session affinity cookie expires after being created. This parameter is ignored unless a supported session affinity policy is set. The current default of 23 hours will be used unless ` + "`" + `session_affinity_ttl` + "`" + ` is explicitly set. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. Valid values are between 1800 and 604800.`,
				},
				resource.Attribute{
					Name:        "session_affinity_attributes",
					Description: `(Optional) Configure cookie attributes for session affinity cookie. See the field documentation below.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) A region code which must be in the list defined [here](https://support.cloudflare.com/hc/en-us/articles/115000540888-Load-Balancing-Geographic-Regions). Multiple entries should not be specified with the same region.`,
				},
				resource.Attribute{
					Name:        "pool_ids",
					Description: `(Required) A list of pool IDs in failover priority to use in the given region.`,
				},
				resource.Attribute{
					Name:        "pop",
					Description: `(Required) A 3-letter code for the Point-of-Presence. Allowed values can be found in the list of datacenters on the [status page](https://www.cloudflarestatus.com/). Multiple entries should not be specified with the same PoP.`,
				},
				resource.Attribute{
					Name:        "pool_ids",
					Description: `(Required) A list of pool IDs in failover priority to use for traffic reaching the given PoP.`,
				},
				resource.Attribute{
					Name:        "samesite",
					Description: `(Optional) Configures the SameSite attribute on session affinity cookie. Value "Auto" will be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note: when using value "None", the secure attribute can not be set to "Never". Valid values: ` + "`" + `"Auto"` + "`" + `, ` + "`" + `"Lax"` + "`" + `, ` + "`" + `"None"` + "`" + ` or ` + "`" + `"Strict"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secure",
					Description: `(Optional) Configures the Secure attribute on session affinity cookie. Value "Always" indicates the Secure attribute will be set in the Set-Cookie header, "Never" indicates the Secure attribute will not be set, and "Auto" will set the Secure attribute depending if Always Use HTTPS is enabled. Valid values: ` + "`" + `"Auto"` + "`" + `, ` + "`" + `"Always"` + "`" + ` or ` + "`" + `"Never"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "drain_duration",
					Description: `(Optional) Configures the drain duration in seconds. This field is only used when session affinity is enabled on the load balancer. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the load balancer.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `The RFC3339 timestamp of when the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "modified_on",
					Description: `The RFC3339 timestamp of when the load balancer was last modified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the load balancer.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `The RFC3339 timestamp of when the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "modified_on",
					Description: `The RFC3339 timestamp of when the load balancer was last modified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_load_balancer_monitor",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Load Balancer Monitor resource.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"monitor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "expected_body",
					Description: `(Optional) A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. Only valid if ` + "`" + `type` + "`" + ` is "http" or "https". Default: "".`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `(Optional) The expected HTTP response code or code range of the health check. Eg ` + "`" + `2xx` + "`" + `. Only valid and required if ` + "`" + `type` + "`" + ` is "http" or "https".`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The method to use for the health check. Valid values are any valid HTTP verb if ` + "`" + `type` + "`" + ` is "http" or "https", or ` + "`" + `connection_established` + "`" + ` if ` + "`" + `type` + "`" + ` is "tcp". Default: "GET" if ` + "`" + `type` + "`" + ` is "http" or "https", or "connection_established" if ` + "`" + `type` + "`" + ` is "tcp" .`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout (in seconds) before marking the health check as failed. Default: 5.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The endpoint path to health check against. Default: "/". Only valid if ` + "`" + `type` + "`" + ` is "http" or "https".`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations. Default: 60.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately. Default: 2.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Optional) The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden. Fields documented below. Only valid if ` + "`" + `type` + "`" + ` is "http" or "https".`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The protocol to use for the healthcheck. Currently supported protocols are 'HTTP', 'HTTPS' and 'TCP'. Default: "http".`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number to use for the healthcheck, required when creating a TCP monitor. Valid values are in the range ` + "`" + `0-65535` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Free text description.`,
				},
				resource.Attribute{
					Name:        "allow_insecure",
					Description: `(Optional) Do not validate the certificate when monitor use HTTPS. Only valid if ` + "`" + `type` + "`" + ` is "http" or "https".`,
				},
				resource.Attribute{
					Name:        "follow_redirects",
					Description: `(Optional) Follow redirects if returned by the origin. Only valid if ` + "`" + `type` + "`" + ` is "http" or "https".`,
				},
				resource.Attribute{
					Name:        "probe_zone",
					Description: `(Optional) Assign this monitor to emulate the specified zone while probing. Only valid if ` + "`" + `type` + "`" + ` is "http" or "https".`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Required) The header name.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of string values for the header. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Load balancer monitor ID.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `The RFC3339 timestamp of when the load balancer monitor was created.`,
				},
				resource.Attribute{
					Name:        "modified_on",
					Description: `The RFC3339 timestamp of when the load balancer monitor was last modified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Load balancer monitor ID.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `The RFC3339 timestamp of when the load balancer monitor was created.`,
				},
				resource.Attribute{
					Name:        "modified_on",
					Description: `The RFC3339 timestamp of when the load balancer monitor was last modified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_load_balancer_pool",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Load Balancer Pool resource.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A short name (tag) for the pool. Only alphanumeric characters, hyphens, and underscores are allowed.`,
				},
				resource.Attribute{
					Name:        "origins",
					Description: `(Required) The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy. It's a complex value. See description below.`,
				},
				resource.Attribute{
					Name:        "check_regions",
					Description: `(Optional) A list of regions (specified by region code) from which to run health checks. Empty means every Cloudflare data center (the default), but requires an Enterprise plan. Region codes can be found [here](https://support.cloudflare.com/hc/en-us/articles/115000540888-Load-Balancing-Geographic-Regions).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Free text description.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether to enable (the default) this pool. Disabled pools will not receive traffic and are excluded from health checks. Disabling a pool will cause any load balancers using it to failover to the next pool (if any).`,
				},
				resource.Attribute{
					Name:        "minimum_origins",
					Description: `(Optional) The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and we will failover to the next available pool. Default: 1.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional) The ID of the Monitor to use for health checking origins within this pool.`,
				},
				resource.Attribute{
					Name:        "notification_email",
					Description: `(Optional) The email address to send health status notifications to. This can be an individual mailbox or a mailing list. Multiple emails can be supplied as a comma delimited list. The`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-identifiable name for the origin.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The IP address (IPv4 or IPv6) of the origin, or the publicly addressable hostname. Hostnames entered here should resolve directly to the origin, and not be a hostname proxied by Cloudflare.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The weight (0.01 - 1.00) of this origin, relative to other origins in the pool. Equal values mean equal weighting. A weight of 0 means traffic will not be sent to this origin, but health is still checked. Default: 1.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether to enable (the default) this origin within the Pool. Disabled origins will not receive traffic and are excluded from health checks. The origin will only be disabled for the current pool. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID for this load balancer pool.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `The RFC3339 timestamp of when the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "modified_on",
					Description: `The RFC3339 timestamp of when the load balancer was last modified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID for this load balancer pool.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `The RFC3339 timestamp of when the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "modified_on",
					Description: `The RFC3339 timestamp of when the load balancer was last modified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_logpull_retention",
			Category:         "Resources",
			ShortDescription: `Allows management of the Logpull Retention settings used to control whether or not to retain HTTP request logs.`,
			Description:      ``,
			Keywords: []string{
				"logpull",
				"retention",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone ID to apply the log retention to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Whether you wish to retain logs or not. ## Import You can import existing Logpull Retention using the zone ID as the identifier. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_logpull_retention.example fb54f084ca7f7b732d3d3ecbd8ef7bf2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_logpush_job",
			Category:         "Resources",
			ShortDescription: `Provides a resource which manages Cloudflare logpush jobs.`,
			Description:      ``,
			Keywords: []string{
				"logpush",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the logpush job to create. Must match the regular expression ` + "`" + `^[a-zA-Z0-9\-\.]`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone ID where the logpush job should be created.`,
				},
				resource.Attribute{
					Name:        "destination_conf",
					Description: `(Required) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#destination).`,
				},
				resource.Attribute{
					Name:        "ownership_challenge",
					Description: `(Required) Ownership challenge token to prove destination ownership. See [Developer documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#usage).`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) Which type of dataset resource to use. Available values are ` + "`" + `"firewall_events"` + "`" + `, ` + "`" + `"http_requests"` + "`" + `, and ` + "`" + `"spectrum_events"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logpull_options",
					Description: `(Optional) Configuration string for the Logshare API. It specifies things like requested fields and timestamp formats. See [Logpull options documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#options).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether to enable the job.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_logpush_ownership_challenge",
			Category:         "Resources",
			ShortDescription: `Provides a resource which manages Cloudflare Logpush ownership challenges to use in a Logpush Job.`,
			Description:      ``,
			Keywords: []string{
				"logpush",
				"ownership",
				"challenge",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone ID where the logpush ownership challenge should be created.`,
				},
				resource.Attribute{
					Name:        "destination_conf",
					Description: `(Required) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#destination). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ownership_challenge_filename",
					Description: `The filename of the ownership challenge which contains the contents required for Logpush Job creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ownership_challenge_filename",
					Description: `The filename of the ownership challenge which contains the contents required for Logpush Job creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_magic_firewall_ruleset",
			Category:         "Resources",
			ShortDescription: `Provides the ability to manage a Magic Firewall Ruleset and it's firewall rules which are used with Magic Transit.`,
			Description:      ``,
			Keywords: []string{
				"magic",
				"firewall",
				"ruleset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The ID of the account where the ruleset is being created.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ruleset.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A note that can be used to annotate the ruleset. The`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Valid values: ` + "`" + `allow` + "`" + ` or ` + "`" + `block` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) A Firewall expression using Wireshark syntax.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A note that can be used to annotate the rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Whether the rule is enabled or not. Valid values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. ## Import An existing Magic Firewall Ruleset can be imported using the account ID and ruleset ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_magic_firewall_ruleset.example d41d8cd98f00b204e9800998ecf8427e/cb029e245cfdd66dc8d2e570d5dd3322 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_origin_ca_certificate",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Origin CA certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"origin",
				"ca",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostnames",
					Description: `(Required) An array of hostnames or wildcard names bound to the certificate.`,
				},
				resource.Attribute{
					Name:        "request_type",
					Description: `(Required) The signature type desired on the certificate.`,
				},
				resource.Attribute{
					Name:        "requested_validity",
					Description: `(Optional) The number of days for which the certificate should be valid. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The x509 serial number of the Origin CA certificate.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The Origin CA certificate`,
				},
				resource.Attribute{
					Name:        "expires_on",
					Description: `The datetime when the certificate will expire. ## Import Origin CA certificate resource can be imported using an ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_origin_ca_certificate.example 276266538771611802607153687288146423901027769273 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The x509 serial number of the Origin CA certificate.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The Origin CA certificate`,
				},
				resource.Attribute{
					Name:        "expires_on",
					Description: `The datetime when the certificate will expire. ## Import Origin CA certificate resource can be imported using an ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_origin_ca_certificate.example 276266538771611802607153687288146423901027769273 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_page_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare page rule resource.`,
			Description:      ``,
			Keywords: []string{
				"page",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID to which the page rule should be added.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The URL pattern to target with the page rule.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `(Required) The actions taken by the page rule, options given below.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the page rule among others for this target, the higher the number the higher the priority as per [API documentation](https://api.cloudflare.com/#page-rules-for-a-zone-create-page-rule).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Whether the page rule is active or disabled. Action blocks support the following:`,
				},
				resource.Attribute{
					Name:        "always_online",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "always_use_https",
					Description: `(Optional) Boolean of whether this action is enabled. Default: false.`,
				},
				resource.Attribute{
					Name:        "automatic_https_rewrites",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "browser_cache_ttl",
					Description: `(Optional) The Time To Live for the browser cache. ` + "`" + `0` + "`" + ` means 'Respect Existing Headers'`,
				},
				resource.Attribute{
					Name:        "browser_check",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bypass_cache_on_cookie",
					Description: `(Optional) String value of cookie name to conditionally bypass cache the page.`,
				},
				resource.Attribute{
					Name:        "cache_by_device_type",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_deception_armor",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_key_fields",
					Description: `(Optional) Controls how Cloudflare creates Cache Keys used to identify files in cache. [See below](#cache-key-fields) for full description.`,
				},
				resource.Attribute{
					Name:        "cache_level",
					Description: `(Optional) Whether to set the cache level to ` + "`" + `"bypass"` + "`" + `, ` + "`" + `"basic"` + "`" + `, ` + "`" + `"simplified"` + "`" + `, ` + "`" + `"aggressive"` + "`" + `, or ` + "`" + `"cache_everything"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_on_cookie",
					Description: `(Optional) String value of cookie name to conditionally cache the page.`,
				},
				resource.Attribute{
					Name:        "cache_ttl_by_status",
					Description: `(Optional) Set cache TTL based on the response status from the origin web server. Can be specified multiple times. [See below](#cache-ttl-by-status) for full description.`,
				},
				resource.Attribute{
					Name:        "disable_apps",
					Description: `(Optional) Boolean of whether this action is enabled. Default: false.`,
				},
				resource.Attribute{
					Name:        "disable_performance",
					Description: `(Optional) Boolean of whether this action is enabled. Default: false.`,
				},
				resource.Attribute{
					Name:        "disable_railgun",
					Description: `(Optional) Boolean of whether this action is enabled. Default: false.`,
				},
				resource.Attribute{
					Name:        "disable_security",
					Description: `(Optional) Boolean of whether this action is enabled. Default: false.`,
				},
				resource.Attribute{
					Name:        "edge_cache_ttl",
					Description: `(Optional) The Time To Live for the edge cache.`,
				},
				resource.Attribute{
					Name:        "email_obfuscation",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "explicit_cache_control",
					Description: `(Optional) Whether origin Cache-Control action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "forwarding_url",
					Description: `(Optional) The URL to forward to, and with what status. See below.`,
				},
				resource.Attribute{
					Name:        "host_header_override",
					Description: `(Optional) Value of the Host header to send.`,
				},
				resource.Attribute{
					Name:        "ip_geolocation",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minify",
					Description: `(Optional) The configuration for HTML, CSS and JS minification. See below for full list of options.`,
				},
				resource.Attribute{
					Name:        "mirage",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "opportunistic_encryption",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_error_page_pass_thru",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "polish",
					Description: `(Optional) Whether this action is ` + "`" + `"off"` + "`" + `, ` + "`" + `"lossless"` + "`" + ` or ` + "`" + `"lossy"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resolve_override",
					Description: `(Optional) Overridden origin server name.`,
				},
				resource.Attribute{
					Name:        "respect_strong_etag",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_buffering",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rocket_loader",
					Description: `(Optional) Whether to set the rocket loader to ` + "`" + `"on"` + "`" + `, ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_level",
					Description: `(Optional) Whether to set the security level to ` + "`" + `"off"` + "`" + `, ` + "`" + `"essentially_off"` + "`" + `, ` + "`" + `"low"` + "`" + `, ` + "`" + `"medium"` + "`" + `, ` + "`" + `"high"` + "`" + `, or ` + "`" + `"under_attack"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_side_exclude",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "smart_errors",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sort_query_string_for_cache",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(Optional) Whether to set the SSL mode to ` + "`" + `"off"` + "`" + `, ` + "`" + `"flexible"` + "`" + `, ` + "`" + `"full"` + "`" + `, ` + "`" + `"strict"` + "`" + `, or ` + "`" + `"origin_pull"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "true_client_ip_header",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "waf",
					Description: `(Optional) Whether this action is ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `. Forwarding URL actions support the following:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL to which the page rule should forward.`,
				},
				resource.Attribute{
					Name:        "status_code",
					Description: `(Required) The status code to use for the redirection. Minify actions support the following:`,
				},
				resource.Attribute{
					Name:        "html",
					Description: `(Required) Whether HTML should be minified. Valid values are ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "css",
					Description: `(Required) Whether CSS should be minified. Valid values are ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "js",
					Description: `(Required) Whether Javascript should be minified. Valid values are ` + "`" + `"on"` + "`" + ` or ` + "`" + `"off"` + "`" + `. ### Cache Key Fields -> This setting is available to Enterprise customers only. A Cache Key is an identifier that Cloudflare uses for a file in a cache. The Cache Key Template defines this identifier for a given HTTP request. For detailed description of use cases and semantics for the particular setting please refer to [Cloudflare Support article](https://support.cloudflare.com/hc/en-us/articles/115004290387-Creating-Cache-Keys). Example: ` + "`" + `` + "`" + `` + "`" + `hcl # Cache JavaScript files: # - ignore CORS Origin header (one copy regardless of requesting Host) # - ignore API key query string # - include browser language preference (e.g. string translations) resource "cloudflare_page_rule" "foobar" { zone_id = var.cloudflare_zone_id target = "embed.${var.cloudflare_zone}/`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `(Required, but allowed to be empty) Controls what cookies go into Cache Key:`,
				},
				resource.Attribute{
					Name:        "check_presence",
					Description: `(Optional, Array) Check for presence of specified cookies, without including their actual values.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Optional, Array) Use values of specified cookies in Cache Key.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Required, but allowed to be empty) Controls what HTTP headers go into Cache Key:`,
				},
				resource.Attribute{
					Name:        "check_presence",
					Description: `(Optional, Array) Check for presence of specified HTTP headers, without including their actual values.`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(Optional, Array) Exclude these HTTP headers from Cache Key. Currently, only the ` + "`" + `Origin` + "`" + ` header can be excluded.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Optional, Array) Use values of specified HTTP headers in Cache Key. Please refer to [Support article](https://support.cloudflare.com/hc/en-us/articles/115004290387-Creating-Cache-Keys) for the list of HTTP headers that cannot be included. The ` + "`" + `Origin` + "`" + ` header is always included unless explicitly excluded.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required, but allowed to be empty) Controls which Host header goes into Cache Key:`,
				},
				resource.Attribute{
					Name:        "resolved",
					Description: `(Optional, Boolean) ` + "`" + `false` + "`" + ` (default) - includes the Host header in the HTTP request sent to the origin; ` + "`" + `true` + "`" + ` - includes the Host header that was resolved to get the origin IP for the request (e.g. changed with Resolve Override Page Rule).`,
				},
				resource.Attribute{
					Name:        "query_string",
					Description: `(Required, but allowed to be empty) Controls which URL query string parameters go into the Cache Key.`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(Optional, Array) Exclude these query string parameters from Cache Key.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Optional, Array) Only use values of specified query string parameters in Cache Key.`,
				},
				resource.Attribute{
					Name:        "ignore",
					Description: `(Optional, Boolean) ` + "`" + `false` + "`" + ` (default) - all query string parameters are used for Cache Key, unless explicitly excluded; ` + "`" + `true` + "`" + ` - all query string parameters are ignored; value is ignored if any of ` + "`" + `exclude` + "`" + ` or ` + "`" + `include` + "`" + ` is non-empty.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required, but allowed to be empty) Controls which end user-related features go into the Cache Key.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Optional, Boolean) ` + "`" + `true` + "`" + ` - classifies a request as “mobile”, “desktop”, or “tablet” based on the User Agent; defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "geo",
					Description: `(Optional, Boolean) ` + "`" + `true` + "`" + ` - includes the client’s country, derived from the IP address; defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional, Boolean) ` + "`" + `true` + "`" + ` - includes the first language code contained in the ` + "`" + `Accept-Language` + "`" + ` header sent by the client; defaults to ` + "`" + `false` + "`" + `. Example: ` + "`" + `` + "`" + `` + "`" + `hcl # Unrealistic example with all features used resource "cloudflare_page_rule" "foobar" { zone_id = var.cloudflare_zone_id target = "${var.cloudflare_zone}/app/`,
				},
				resource.Attribute{
					Name:        "codes",
					Description: `(Required) A HTTP code (e.g. ` + "`" + `404` + "`" + `) or range of codes (e.g. ` + "`" + `400-499` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) Duration a resource lives in the Cloudflare cache.`,
				},
				resource.Attribute{
					Name:        "0",
					Description: `sets ` + "`" + `no-cache` + "`" + `, saved to cache, but expired immediately (revalidate from origin every time)`,
				},
				resource.Attribute{
					Name:        "-1",
					Description: `sets ` + "`" + `no-store` + "`" + `, never save to cache Example: ` + "`" + `` + "`" + `` + "`" + `hcl resource "cloudflare_page_rule" "test" { zone_id = var.cloudflare_zone_id target = "${var.cloudflare_zone}/app/`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The page rule ID.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The URL pattern targeted by the page rule.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `The actions applied by the page rule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the page rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Whether the page rule is active or disabled. ## Import Page rules can be imported using a composite ID formed of zone ID and page rule ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_page_rule.default d41d8cd98f00b204e9800998ecf8427e/ch8374ftwdghsif43 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The page rule ID.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The URL pattern targeted by the page rule.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `The actions applied by the page rule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the page rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Whether the page rule is active or disabled. ## Import Page rules can be imported using a composite ID formed of zone ID and page rule ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_page_rule.default d41d8cd98f00b204e9800998ecf8427e/ch8374ftwdghsif43 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_rate_limit",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare rate limit resource for a particular zone.`,
			Description:      ``,
			Keywords: []string{
				"rate",
				"limit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID to apply rate limiting to.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) The threshold that triggers the rate limit mitigations, combine with period. i.e. threshold per period (min: 2, max: 1,000,000).`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Required) The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed (min: 1, max: 86,400).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to be performed when the threshold of matched traffic within the period defined is exceeded.`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `(Optional) Determines which traffic the rate limit counts towards the threshold. By default matches all traffic in the zone. See definition below.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Whether this ratelimit is currently disabled. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed.`,
				},
				resource.Attribute{
					Name:        "bypass_url_patterns",
					Description: `(Optional) URLs matching the patterns specified here will be excluded from rate limiting.`,
				},
				resource.Attribute{
					Name:        "correlate",
					Description: `(Optional) Determines how rate limiting is applied. By default if not specified, rate limiting applies to the clients IP address. The`,
				},
				resource.Attribute{
					Name:        "request",
					Description: `(Optional) Matches HTTP requests (from the client to Cloudflare). See definition below.`,
				},
				resource.Attribute{
					Name:        "methods",
					Description: `(Optional) HTTP Methods, can be a subset ['POST','PUT'] or all ['\_ALL\_']. Default: ['\_ALL\_'].`,
				},
				resource.Attribute{
					Name:        "schemes",
					Description: `(Optional) HTTP Schemes, can be one ['HTTPS'], both ['HTTP','HTTPS'] or all ['\_ALL\_']. Default: ['\_ALL\_'].`,
				},
				resource.Attribute{
					Name:        "url_pattern",
					Description: `(Optional) The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use`,
				},
				resource.Attribute{
					Name:        "statuses",
					Description: `(Optional) HTTP Status codes, can be one [403], many [401,403] or indicate all by not providing this value.`,
				},
				resource.Attribute{
					Name:        "origin_traffic",
					Description: `(Optional) Only count traffic that has come from your origin servers. If true, cached items that Cloudflare serve will not count towards rate limiting. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) block is a list of maps with the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the response header to match.`,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(Required) The operator when matching. Allowable values are 'eq', 'ne' where ` + "`" + `eq` + "`" + ` means equals, ` + "`" + `ne` + "`" + ` means not equals.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the header, which will be exactly matched. The`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The type of action to perform. Allowable values are 'simulate', 'ban', 'challenge' and 'js_challenge'.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The time in seconds as an integer to perform the mitigation action. This field is required if the ` + "`" + `mode` + "`" + ` is either ` + "`" + `simulate` + "`" + ` or ` + "`" + `ban` + "`" + `. Must be the same or greater than the period (min: 1, max: 86400).`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Optional) Custom content-type and body to return, this overrides the custom error for the zone. This field is not required. Omission will result in default HTML error page. Definition below. The`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Required) The content-type of the body, must be one of: 'text/plain', 'text/xml', 'application/json'.`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `(Required) The body to return, the content here should conform to the content_type. The`,
				},
				resource.Attribute{
					Name:        "by",
					Description: `(Optional) If set to 'nat', NAT support will be enabled for rate limiting. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Rate limit ID. ## Import Rate limits can be imported using a composite ID formed of zone name and rate limit ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_rate_limit.default d41d8cd98f00b204e9800998ecf8427e/ch8374ftwdghsif43 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Rate limit ID. ## Import Rate limits can be imported using a composite ID formed of zone name and rate limit ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_rate_limit.default d41d8cd98f00b204e9800998ecf8427e/ch8374ftwdghsif43 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_record",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare record resource.`,
			Description:      ``,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID to add the record to`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the record`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The (string) value of the record. Either this or ` + "`" + `data` + "`" + ` must be specified`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional) Map of attributes that constitute the record value. Primarily used for LOC and SRV record types. Either this or ` + "`" + `value` + "`" + ` must be specified`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record ([automatic: '1'](https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record))`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the record`,
				},
				resource.Attribute{
					Name:        "proxied",
					Description: `(Optional) Whether the record gets Cloudflare's origin protection; defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the record`,
				},
				resource.Attribute{
					Name:        "proxiable",
					Description: `Shows whether this record can be proxied, must be true if setting ` + "`" + `proxied=true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `The RFC3339 timestamp of when the record was created`,
				},
				resource.Attribute{
					Name:        "modified_on",
					Description: `The RFC3339 timestamp of when the record was last modified`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A key-value map of string metadata Cloudflare associates with the record ## Import Records can be imported using a composite ID formed of zone ID and record ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_record.default ae36f999674d196762efcc5abb06b345/d41d8cd98f00b204e9800998ecf8427e ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "ae36f999674d196762efcc5abb06b345",
					Description: `the zone ID`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `record ID as returned by [API](https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the record`,
				},
				resource.Attribute{
					Name:        "proxiable",
					Description: `Shows whether this record can be proxied, must be true if setting ` + "`" + `proxied=true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `The RFC3339 timestamp of when the record was created`,
				},
				resource.Attribute{
					Name:        "modified_on",
					Description: `The RFC3339 timestamp of when the record was last modified`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A key-value map of string metadata Cloudflare associates with the record ## Import Records can be imported using a composite ID formed of zone ID and record ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_record.default ae36f999674d196762efcc5abb06b345/d41d8cd98f00b204e9800998ecf8427e ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "ae36f999674d196762efcc5abb06b345",
					Description: `the zone ID`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `record ID as returned by [API](https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_spectrum_application",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Spectrum Application resource.`,
			Description:      ``,
			Keywords: []string{
				"spectrum",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID to add the application to`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Required) The name and type of DNS record for the Spectrum application. Fields documented below.`,
				},
				resource.Attribute{
					Name:        "origin_direct",
					Description: `(Optional) A list of destination addresses to the origin. e.g. ` + "`" + `tcp://192.0.2.1:22` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_dns",
					Description: `(Optional) A destination DNS addresses to the origin. Fields documented below.`,
				},
				resource.Attribute{
					Name:        "origin_port",
					Description: `(Optional) If using ` + "`" + `origin_dns` + "`" + ` and not ` + "`" + `origin_port_range` + "`" + `, this is a required attribute. Origin port to proxy traffice to e.g. ` + "`" + `22` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_port_range",
					Description: `(Optional) If using ` + "`" + `origin_dns` + "`" + ` and not ` + "`" + `origin_port` + "`" + `, this is a required attribute. Origin port range to proxy traffice to. When using a range, the protocol field must also specify a range, e.g. ` + "`" + `tcp/22-23` + "`" + `. Fields documented below.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) TLS configuration option for Cloudflare to connect to your origin. Valid values are: ` + "`" + `off` + "`" + `, ` + "`" + `flexible` + "`" + `, ` + "`" + `full` + "`" + ` and ` + "`" + `strict` + "`" + `. Defaults to ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_firewall",
					Description: `(Optional) Enables the IP Firewall for this application. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `(Optional) Enables a proxy protocol to the origin. Valid values are: ` + "`" + `off` + "`" + `, ` + "`" + `v1` + "`" + `, ` + "`" + `v2` + "`" + `, and ` + "`" + `simple` + "`" + `. Defaults to ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "traffic_type",
					Description: `(Optional) Sets application type. Valid values are: ` + "`" + `direct` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `. Defaults to ` + "`" + `direct` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "argo_smart_routing",
					Description: `(Optional). Enables Argo Smart Routing. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "edge_ip_connectivity",
					Description: `(Optional). Choose which types of IP addresses will be provisioned for this subdomain. Valid values are: ` + "`" + `all` + "`" + `, ` + "`" + `ipv4` + "`" + `, ` + "`" + `ipv6` + "`" + `. Defaults to ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "edge_ips",
					Description: `(Optional). A list of edge IPs (IPv4 and/or IPv6) to configure Spectrum application to. Requires [Bring Your Own IP](https://developers.cloudflare.com/spectrum/getting-started/byoip/) provisioned.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of DNS record associated with the application. Valid values: ` + "`" + `CNAME` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the DNS record associated with the application.i.e. ` + "`" + `ssh.example.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Fully qualified domain name of the origin e.g. origin-ssh.example.com.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) Lower bound of the origin port range, e.g. ` + "`" + `1000` + "`" + ``,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) Upper bound of the origin port range, e.g. ` + "`" + `2000` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the spectrum application. ## Import Spectrum resource can be imported using a zone ID and Application ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_spectrum_application.example d41d8cd98f00b204e9800998ecf8427e/9a7806061c88ada191ed06f989cc3dac ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID, as returned from [API](https://api.cloudflare.com/#zone-list-zones)`,
				},
				resource.Attribute{
					Name:        "9a7806061c88ada191ed06f989cc3dac",
					Description: `Application ID`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the spectrum application. ## Import Spectrum resource can be imported using a zone ID and Application ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_spectrum_application.example d41d8cd98f00b204e9800998ecf8427e/9a7806061c88ada191ed06f989cc3dac ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID, as returned from [API](https://api.cloudflare.com/#zone-list-zones)`,
				},
				resource.Attribute{
					Name:        "9a7806061c88ada191ed06f989cc3dac",
					Description: `Application ID`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_waf_group",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare WAF rule group resource for a particular zone.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID to apply to.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The WAF Rule Group ID.`,
				},
				resource.Attribute{
					Name:        "package_id",
					Description: `(Optional) The ID of the WAF Rule Package that contains the group.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode of the group, can be one of ["on", "off"]. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The WAF Rule Group ID, the same as ` + "`" + `group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "package_id",
					Description: `The ID of the WAF Rule Package that contains the group. ## Import WAF Rule Groups can be imported using a composite ID formed of zone ID and the WAF Rule Group ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_waf_group.honey_pot ae36f999674d196762efcc5abb06b345/de677e5818985db1285d0e80225f06e5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The WAF Rule Group ID, the same as ` + "`" + `group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "package_id",
					Description: `The ID of the WAF Rule Package that contains the group. ## Import WAF Rule Groups can be imported using a composite ID formed of zone ID and the WAF Rule Group ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_waf_group.honey_pot ae36f999674d196762efcc5abb06b345/de677e5818985db1285d0e80225f06e5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_waf_override",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare WAF Override resource.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"override",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone to which the WAF override condition should be added.`,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(Required) An array of URLs to apply the WAF override to.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Required) A list of WAF rule ID to rule action you intend to apply.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `(Optional) Whether this package is currently paused.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of what the WAF override does.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Relative priority of this configuration when multiple configurations match a single URL.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Similar to ` + "`" + `rules` + "`" + `; which WAF groups you want to alter.`,
				},
				resource.Attribute{
					Name:        "rewrite_action",
					Description: `(Optional) When a WAF rule matches, substitute its configured action for a different action specified by this definition. ## Import WAF Overrides can be imported using a composite ID formed of zone ID and override ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_waf_override.my_example_waf_override 3abe5b950053dbddf1516d89f9ef1e8a/9d4e66d7649c178663bf62e06dbacb23 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_waf_package",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare WAF rule package resource for a particular zone.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID to apply to.`,
				},
				resource.Attribute{
					Name:        "package_id",
					Description: `(Required) The WAF Package ID.`,
				},
				resource.Attribute{
					Name:        "sensitivity",
					Description: `(Optional) The sensitivity of the package, can be one of ["high", "medium", "low", "off"].`,
				},
				resource.Attribute{
					Name:        "action_mode",
					Description: `(Optional) The action mode of the package, can be one of ["block", "challenge", "simulate"]. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The WAF Package ID, the same as package_id. ## Import Packages can be imported using a composite ID formed of zone ID and the WAF Package ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_waf_package.owasp ae36f999674d196762efcc5abb06b345/a25a9a7e9c00afc1fb2e0245519d725b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The WAF Package ID, the same as package_id. ## Import Packages can be imported using a composite ID formed of zone ID and the WAF Package ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_waf_package.owasp ae36f999674d196762efcc5abb06b345/a25a9a7e9c00afc1fb2e0245519d725b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_waf_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare WAF rule resource for a particular zone.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID to apply to.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) The WAF Rule ID.`,
				},
				resource.Attribute{
					Name:        "package_id",
					Description: `(Optional) The ID of the WAF Rule Package that contains the rule.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The mode of the rule, can be one of ["block", "challenge", "default", "disable", "simulate"] or ["on", "off"] depending on the WAF Rule type. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The WAF Rule ID, the same as rule_id.`,
				},
				resource.Attribute{
					Name:        "package_id",
					Description: `The ID of the WAF Rule Package that contains the rule.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The ID of the WAF Rule Group that contains the rule. ## Import Rules can be imported using a composite ID formed of zone ID and the WAF Rule ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_waf_rule.100000 ae36f999674d196762efcc5abb06b345/100000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The WAF Rule ID, the same as rule_id.`,
				},
				resource.Attribute{
					Name:        "package_id",
					Description: `The ID of the WAF Rule Package that contains the rule.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The ID of the WAF Rule Group that contains the rule. ## Import Rules can be imported using a composite ID formed of zone ID and the WAF Rule ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_waf_rule.100000 ae36f999674d196762efcc5abb06b345/100000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_worker_cron_trigger",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare Worker Cron Trigger resource.`,
			Description:      ``,
			Keywords: []string{
				"worker",
				"cron",
				"trigger",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "script_name",
					Description: `(Required) Worker script to target for the schedules`,
				},
				resource.Attribute{
					Name:        "schedules",
					Description: `(Required) List of cron expressions to execute the Worker Script ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `md5 checksum of the script name`,
				},
				resource.Attribute{
					Name:        "script_name",
					Description: `Name of the Worker Script being targeted`,
				},
				resource.Attribute{
					Name:        "schedules",
					Description: `List of cron expressions in use ## Import Worker Cron Triggers can be imported using the script name of the Worker they are targeting. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_worker_cron_trigger.example my-script ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `md5 checksum of the script name`,
				},
				resource.Attribute{
					Name:        "script_name",
					Description: `Name of the Worker Script being targeted`,
				},
				resource.Attribute{
					Name:        "schedules",
					Description: `List of cron expressions in use ## Import Worker Cron Triggers can be imported using the script name of the Worker they are targeting. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_worker_cron_trigger.example my-script ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_worker_route",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare worker route resource.`,
			Description:      ``,
			Keywords: []string{
				"worker",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone ID to add the route to.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) The [route pattern](https://developers.cloudflare.com/workers/about/routes/)`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID`,
				},
				resource.Attribute{
					Name:        "9a7806061c88ada191ed06f989cc3dac",
					Description: `route ID as returned by [API](https://api.cloudflare.com/#worker-filters-list-filters)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_worker_script",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare worker script resource.`,
			Description:      ``,
			Keywords: []string{
				"worker",
				"script",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the script.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The script content.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The global variable for the binding in your Worker code.`,
				},
				resource.Attribute{
					Name:        "kv_namespace_id",
					Description: `(Required) ID of the KV namespace you want to use.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The global variable for the binding in your Worker code.`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `(Required) The plain text you want to store.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The global variable for the binding in your Worker code.`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `(Required) The secret text you want to store.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The global variable for the binding in your Worker code.`,
				},
				resource.Attribute{
					Name:        "module",
					Description: `(Required) The base64 encoded wasm module you want to store. ## Import To import a script, use a script name, e.g. ` + "`" + `script_name` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_worker_script.default script_name ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "script_name",
					Description: `the script name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_workers_kv",
			Category:         "Resources",
			ShortDescription: `Provides the ability to manage Cloudflare Workers KV features.`,
			Description:      ``,
			Keywords: []string{
				"workers",
				"kv",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Required) The ID of the Workers KV namespace in which you want to create the KV pair`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key name`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The string value to be stored in the key ## Import Workers KV Namespace settings can be imported using it's ID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_workers_kv_namespace",
			Category:         "Resources",
			ShortDescription: `Provides the ability to manage Cloudflare Workers KV Namespace features.`,
			Description:      ``,
			Keywords: []string{
				"workers",
				"kv",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The name of the namespace you wish to create. ## Import Workers KV Namespace settings can be imported using it's ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_workers_kv_namespace.example beaeb6716c9443eaa4deef11763ccca6 ` + "`" + `` + "`" + `` + "`" + ` where: - ` + "`" + `beaeb6716c9443eaa4deef11763ccca6` + "`" + ` is the ID of the namespace`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_zone",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare resource to create and modify a zone.`,
			Description:      ``,
			Keywords: []string{
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The DNS zone name which will be added.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `(Optional) Boolean of whether this zone is paused (traffic bypasses Cloudflare). Default: false.`,
				},
				resource.Attribute{
					Name:        "jump_start",
					Description: `(Optional) Boolean of whether to scan for DNS records on creation. Ignored after zone is created. Default: false.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Optional) The name of the commercial plan to apply to the zone, can be updated once the one is created; one of ` + "`" + `free` + "`" + `, ` + "`" + `pro` + "`" + `, ` + "`" + `business` + "`" + `, ` + "`" + `enterprise` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `A full zone implies that DNS is hosted with Cloudflare. A partial zone is typically a partner-hosted zone or a CNAME setup. Valid values: ` + "`" + `full` + "`" + `, ` + "`" + `partial` + "`" + `. Default is ` + "`" + `full` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The zone ID.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The name of the commercial plan to apply to the zone.`,
				},
				resource.Attribute{
					Name:        "vanity_name_servers",
					Description: `List of Vanity Nameservers (if set).`,
				},
				resource.Attribute{
					Name:        "meta.wildcard_proxiable",
					Description: `Indicates whether wildcard DNS records can receive Cloudflare security and performance features.`,
				},
				resource.Attribute{
					Name:        "meta.phishing_detected",
					Description: `Indicates if URLs on the zone have been identified as hosting phishing content.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the zone. Valid values: ` + "`" + `active` + "`" + `, ` + "`" + `pending` + "`" + `, ` + "`" + `initializing` + "`" + `, ` + "`" + `moved` + "`" + `, ` + "`" + `deleted` + "`" + `, ` + "`" + `deactivated` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Cloudflare-assigned name servers. This is only populated for zones that use Cloudflare DNS.`,
				},
				resource.Attribute{
					Name:        "verification_key",
					Description: `Contains the TXT record value to validate domain ownership. This is only populated for zones of type ` + "`" + `partial` + "`" + `. ## Import Zone resource can be imported using a zone ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_zone.example d41d8cd98f00b204e9800998ecf8427e ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID, as returned from [API](https://api.cloudflare.com/#zone-list-zones)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The zone ID.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The name of the commercial plan to apply to the zone.`,
				},
				resource.Attribute{
					Name:        "vanity_name_servers",
					Description: `List of Vanity Nameservers (if set).`,
				},
				resource.Attribute{
					Name:        "meta.wildcard_proxiable",
					Description: `Indicates whether wildcard DNS records can receive Cloudflare security and performance features.`,
				},
				resource.Attribute{
					Name:        "meta.phishing_detected",
					Description: `Indicates if URLs on the zone have been identified as hosting phishing content.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the zone. Valid values: ` + "`" + `active` + "`" + `, ` + "`" + `pending` + "`" + `, ` + "`" + `initializing` + "`" + `, ` + "`" + `moved` + "`" + `, ` + "`" + `deleted` + "`" + `, ` + "`" + `deactivated` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Cloudflare-assigned name servers. This is only populated for zones that use Cloudflare DNS.`,
				},
				resource.Attribute{
					Name:        "verification_key",
					Description: `Contains the TXT record value to validate domain ownership. This is only populated for zones of type ` + "`" + `partial` + "`" + `. ## Import Zone resource can be imported using a zone ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_zone.example d41d8cd98f00b204e9800998ecf8427e ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID, as returned from [API](https://api.cloudflare.com/#zone-list-zones)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_zone_lockdown",
			Category:         "Resources",
			ShortDescription: `Provides a Cloudflare resource to lock down access to URLs by IP address or IP ranges.`,
			Description:      ``,
			Keywords: []string{
				"zone",
				"lockdown",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID to which the access rule should be added.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description about the lockdown entry. Typically used as a reminder or explanation for the lockdown.`,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(Required) A list of simple wildcard patterns to match requests against. The order of the urls is unimportant.`,
				},
				resource.Attribute{
					Name:        "configurations",
					Description: `(Required) A list of IP addresses or IP ranges to match the request against specified in target, value pairs. It's a complex value. See description below. The order of the configuration entries is unimportant.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `(Optional) Boolean of whether this zone lockdown is currently paused. Default: false.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The request property to target. Allowed values: "ip", "ip_range"`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to target. Depends on target's type. IP addresses should just be standard IPv4/IPv6 notation i.e. ` + "`" + `198.51.100.4` + "`" + ` or ` + "`" + `2001:db8::/32` + "`" + ` and IP ranges in CIDR format i.e. ` + "`" + `198.51.0.0/16` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The access rule ID. ## Import Records can be imported using a composite ID formed of zone name and record ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_zone_lockdown d41d8cd98f00b204e9800998ecf8427e/37cb64fe4a90adb5ca3afc04f2c82a2f ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID`,
				},
				resource.Attribute{
					Name:        "37cb64fe4a90adb5ca3afc04f2c82a2f",
					Description: `zone lockdown ID as returned by [API](https://api.cloudflare.com/#zone-lockdown-list-lockdown-rules)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The access rule ID. ## Import Records can be imported using a composite ID formed of zone name and record ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import cloudflare_zone_lockdown d41d8cd98f00b204e9800998ecf8427e/37cb64fe4a90adb5ca3afc04f2c82a2f ` + "`" + `` + "`" + `` + "`" + ` where:`,
				},
				resource.Attribute{
					Name:        "d41d8cd98f00b204e9800998ecf8427e",
					Description: `zone ID`,
				},
				resource.Attribute{
					Name:        "37cb64fe4a90adb5ca3afc04f2c82a2f",
					Description: `zone lockdown ID as returned by [API](https://api.cloudflare.com/#zone-lockdown-list-lockdown-rules)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_zone_settings_override",
			Category:         "Resources",
			ShortDescription: `Provides a resource which customizes Cloudflare zone settings.`,
			Description:      ``,
			Keywords: []string{
				"zone",
				"settings",
				"override",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The DNS zone ID to which apply settings.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Optional) Settings overrides that will be applied to the zone. If a setting is not specified the existing setting will be used. For a full list of available settings see below. The`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The zone ID.`,
				},
				resource.Attribute{
					Name:        "initial_settings",
					Description: `Settings present in the zone at the time the resource is created. This will be used to restore the original settings when this resource is destroyed. Shares the same schema as the ` + "`" + `settings` + "`" + ` attribute (Above).`,
				},
				resource.Attribute{
					Name:        "intial_settings_read_at",
					Description: `Time when this resource was created and the ` + "`" + `initial_settings` + "`" + ` were set.`,
				},
				resource.Attribute{
					Name:        "readonly_settings",
					Description: `Which of the current ` + "`" + `settings` + "`" + ` are not able to be set by the user. Which settings these are is determined by plan level and user permissions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The zone ID.`,
				},
				resource.Attribute{
					Name:        "initial_settings",
					Description: `Settings present in the zone at the time the resource is created. This will be used to restore the original settings when this resource is destroyed. Shares the same schema as the ` + "`" + `settings` + "`" + ` attribute (Above).`,
				},
				resource.Attribute{
					Name:        "intial_settings_read_at",
					Description: `Time when this resource was created and the ` + "`" + `initial_settings` + "`" + ` were set.`,
				},
				resource.Attribute{
					Name:        "readonly_settings",
					Description: `Which of the current ` + "`" + `settings` + "`" + ` are not able to be set by the user. Which settings these are is determined by plan level and user permissions.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"cloudflare_access_application":          0,
		"cloudflare_access_ca_certificate":       1,
		"cloudflare_access_group":                2,
		"cloudflare_access_identity_provider":    3,
		"cloudflare_access_policy":               4,
		"cloudflare_access_rule":                 5,
		"cloudflare_access_service_token":        6,
		"cloudflare_account_member":              7,
		"cloudflare_api_token":                   8,
		"cloudflare_argo":                        9,
		"cloudflare_argo_tunnel":                 10,
		"cloudflare_byo_ip_prefix":               11,
		"cloudflare_certificate_pack":            12,
		"cloudflare_custom_hostname":             13,
		"cloudflare_custom_pages":                14,
		"cloudflare_custom_ssl":                  15,
		"cloudflare_filter":                      16,
		"cloudflare_firewall_rule":               17,
		"cloudflare_healthcheck":                 18,
		"cloudflare_ip_list":                     19,
		"cloudflare_load_balancer":               20,
		"cloudflare_load_balancer_monitor":       21,
		"cloudflare_load_balancer_pool":          22,
		"cloudflare_logpull_retention":           23,
		"cloudflare_logpush_job":                 24,
		"cloudflare_logpush_ownership_challenge": 25,
		"cloudflare_magic_firewall_ruleset":      26,
		"cloudflare_origin_ca_certificate":       27,
		"cloudflare_page_rule":                   28,
		"cloudflare_rate_limit":                  29,
		"cloudflare_record":                      30,
		"cloudflare_spectrum_application":        31,
		"cloudflare_waf_group":                   32,
		"cloudflare_waf_override":                33,
		"cloudflare_waf_package":                 34,
		"cloudflare_waf_rule":                    35,
		"cloudflare_worker_cron_trigger":         36,
		"cloudflare_worker_route":                37,
		"cloudflare_worker_script":               38,
		"cloudflare_workers_kv":                  39,
		"cloudflare_workers_kv_namespace":        40,
		"cloudflare_zone":                        41,
		"cloudflare_zone_lockdown":               42,
		"cloudflare_zone_settings_override":      43,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
