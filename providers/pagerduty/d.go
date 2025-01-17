package pagerduty

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_business_service",
			Category:         "Data Sources",
			ShortDescription: `Get information about a business service that you have created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The business service name to use to find a business service in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found business service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found business service. [1]: https://api-reference.pagerduty.com/#!/Business_Services/get_business_services`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found business service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found business service. [1]: https://api-reference.pagerduty.com/#!/Business_Services/get_business_services`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_escalation_policy",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Escalation Policy. This data source can be helpful when an escalation policy is handled outside Terraform but you still want to reference it in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to use to find an escalation policy in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found escalation policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found escalation policy. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Escalation_Policies/get_escalation_policies`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found escalation policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found escalation policy. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Escalation_Policies/get_escalation_policies`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_extension_schema",
			Category:         "Data Sources",
			ShortDescription: `Get information about an extension vendor that you can use for a service (e.g: Slack, Generic Webhook, ServiceNow).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The extension name to use to find an extension vendor in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found extension vendor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found extension vendor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The generic service type for this extension vendor. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Extension_Schemas/get_extension_schemas`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found extension vendor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found extension vendor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The generic service type for this extension vendor. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Extension_Schemas/get_extension_schemas`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_priority",
			Category:         "Data Sources",
			ShortDescription: `Get information about a priority that you can use with ruleset_rules, etc.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the priority to find in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found priority.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found priority.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the found priority. [1]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1priorities/get`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found priority.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found priority.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the found priority. [1]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1priorities/get`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_ruleset",
			Category:         "Data Sources",
			ShortDescription: `Get information about a ruleset that you have created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ruleset to find in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found ruleset.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found ruleset.`,
				},
				resource.Attribute{
					Name:        "routing_keys",
					Description: `Routing keys routed to this ruleset. [1]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1rulesets/get [2]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1rulesets~1%7Bid%7D~1rules/get`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found ruleset.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found ruleset.`,
				},
				resource.Attribute{
					Name:        "routing_keys",
					Description: `Routing keys routed to this ruleset. [1]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1rulesets/get [2]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1rulesets~1%7Bid%7D~1rules/get`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_schedule",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Schedule. This data source can be helpful when a schedule is handled outside Terraform but you still want to reference it in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to use to find a schedule in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found schedule. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Schedules/get_schedules`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found schedule. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Schedules/get_schedules`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_service",
			Category:         "Data Sources",
			ShortDescription: `Get information about a service that you have created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service name to use to find a service in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found service. [1]: https://api-reference.pagerduty.com/#!/Services/get_services`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found service. [1]: https://api-reference.pagerduty.com/#!/Services/get_services`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_team",
			Category:         "Data Sources",
			ShortDescription: `Get information about a team that you can use with escalation_policies, schedules etc.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the team to find in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found team.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the found team. [1]: https://v1.developer.pagerduty.com/documentation/rest/teams/list`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found team.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the found team. [1]: https://v1.developer.pagerduty.com/documentation/rest/teams/list`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_user",
			Category:         "Data Sources",
			ShortDescription: `Get information about a user that you can use for a service integration (e.g Amazon Cloudwatch, Splunk, Datadog).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The email to use to find a user in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found user. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Users/get_users`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found user. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Users/get_users`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_user_contact_method",
			Category:         "Data Sources",
			ShortDescription: `Get information about a contact method of a PagerDuty user (email, phone, SMS or push notification).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The ID of the user.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The contact method type. May be (` + "`" + `email_contact_method` + "`" + `, ` + "`" + `phone_contact_method` + "`" + `, ` + "`" + `sms_contact_method` + "`" + `, ` + "`" + `push_notification_contact_method` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label (e.g., "Work", "Mobile", "Ashley's iPhone", etc.). ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the found contact method. May be (` + "`" + `email_contact_method` + "`" + `, ` + "`" + `phone_contact_method` + "`" + `, ` + "`" + `sms_contact_method` + "`" + `, ` + "`" + `push_notification_contact_method` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "send_short_email",
					Description: `Send an abbreviated email message instead of the standard email output. (Email contact method only.)`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `The 1-to-3 digit country calling code. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label (e.g., "Work", "Mobile", "Ashley's iPhone", etc.).`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The "address" to deliver to: ` + "`" + `email` + "`" + `, ` + "`" + `phone number` + "`" + `, etc., depending on the type.`,
				},
				resource.Attribute{
					Name:        "blacklisted",
					Description: `If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If true, this phone is capable of receiving SMS messages. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Either ` + "`" + `ios` + "`" + ` or ` + "`" + `android` + "`" + `, depending on the type of the device receiving notifications. (Push notification contact method only.) [1]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1users~1%7Bid%7D~1contact_methods~1%7Bcontact_method_id%7D/get [2]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1users~1%7Bid%7D/get`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the found contact method. May be (` + "`" + `email_contact_method` + "`" + `, ` + "`" + `phone_contact_method` + "`" + `, ` + "`" + `sms_contact_method` + "`" + `, ` + "`" + `push_notification_contact_method` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "send_short_email",
					Description: `Send an abbreviated email message instead of the standard email output. (Email contact method only.)`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `The 1-to-3 digit country calling code. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label (e.g., "Work", "Mobile", "Ashley's iPhone", etc.).`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The "address" to deliver to: ` + "`" + `email` + "`" + `, ` + "`" + `phone number` + "`" + `, etc., depending on the type.`,
				},
				resource.Attribute{
					Name:        "blacklisted",
					Description: `If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If true, this phone is capable of receiving SMS messages. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Either ` + "`" + `ios` + "`" + ` or ` + "`" + `android` + "`" + `, depending on the type of the device receiving notifications. (Push notification contact method only.) [1]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1users~1%7Bid%7D~1contact_methods~1%7Bcontact_method_id%7D/get [2]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1users~1%7Bid%7D/get`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_vendor",
			Category:         "Data Sources",
			ShortDescription: `Get information about a vendor that you can use for a service integration (e.g Amazon Cloudwatch, Splunk, Datadog).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The vendor name to use to find a vendor in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found vendor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found vendor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The generic service type for this vendor. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Vendors/get_vendors`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found vendor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found vendor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The generic service type for this vendor. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Vendors/get_vendors`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"pagerduty_business_service":    0,
		"pagerduty_escalation_policy":   1,
		"pagerduty_extension_schema":    2,
		"pagerduty_priority":            3,
		"pagerduty_ruleset":             4,
		"pagerduty_schedule":            5,
		"pagerduty_service":             6,
		"pagerduty_team":                7,
		"pagerduty_user":                8,
		"pagerduty_user_contact_method": 9,
		"pagerduty_vendor":              10,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
