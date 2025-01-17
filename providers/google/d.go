package google

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "google_active_folder",
			Category:         "Data Sources",
			ShortDescription: `Get a folder within GCP.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The folder's display name.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Required) The resource name of the parent Folder or Organization. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the Folder. This uniquely identifies the folder.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the Folder. This uniquely identifies the folder.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_app_engine_default_service_account",
			Category:         "Data Sources",
			ShortDescription: `Retrieve the default App Engine service account used in this project`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project ID. If it is not provided, the provider project is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email address of the default service account used by App Engine in this project.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `The unique id of the service account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully-qualified name of the service account.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the service account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `Email address of the default service account used by App Engine in this project.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `The unique id of the service account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully-qualified name of the service account.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the service account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_bigquery_default_service_account",
			Category:         "Data Sources",
			ShortDescription: `Get the email address of the project's BigQuery service account`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project the unique service account was created for. If it is not provided, the provider project is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email address of the service account. This value is often used to refer to the service account in order to grant IAM permissions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The email address of the service account. This value is often used to refer to the service account in order to grant IAM permissions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_billing_account",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Billing Account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The billing account ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the billing account in the form ` + "`" + `billingAccounts/{billing_account_id}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_ids",
					Description: `The IDs of any projects associated with the billing account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The billing account ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the billing account in the form ` + "`" + `billingAccounts/{billing_account_id}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_ids",
					Description: `The IDs of any projects associated with the billing account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_client_config",
			Category:         "Data Sources",
			ShortDescription: `Get information about the configuration of the Google Cloud provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `The ID of the project to apply any resources to.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region to operate under.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The zone to operate under.`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `The OAuth2 access token used by the client to authenticate against the Google Cloud API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `The ID of the project to apply any resources to.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region to operate under.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The zone to operate under.`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `The OAuth2 access token used by the client to authenticate against the Google Cloud API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_client_openid_userinfo",
			Category:         "Data Sources",
			ShortDescription: `Get OpenID userinfo about the credentials used with the Google provider, specifically the email.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The email of the account used by the provider to authenticate with GCP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The email of the account used by the provider to authenticate with GCP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_cloud_identity_group_memberships",
			Category:         "Data Sources",
			ShortDescription: `Get list of the Cloud Identity Group Memberships within a Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group",
					Description: `The parent Group resource under which to lookup the Membership names. Must be of the form groups/{group_id}. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "memberships",
					Description: `The list of memberships under the given group. Structure is documented below. The ` + "`" + `memberships` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `The MembershipRoles that apply to the Membership. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "member_key",
					Description: `(Optional) EntityKey of the member. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "preferred_member_key",
					Description: `(Optional) EntityKey of the member. Structure is documented below. The ` + "`" + `roles` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the MembershipRole. One of OWNER, MANAGER, MEMBER. The ` + "`" + `member_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the entity. For Google-managed entities, the id is the email address of an existing group or user. For external-identity-mapped entities, the id is a string conforming to the Identity Source's requirements.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace in which the entity exists. If not populated, the EntityKey represents a Google-managed entity such as a Google user or a Google Group. If populated, the EntityKey represents an external-identity-mapped group. The ` + "`" + `preferred_member_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the entity. For Google-managed entities, the id is the email address of an existing group or user. For external-identity-mapped entities, the id is a string conforming to the Identity Source's requirements.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace in which the entity exists. If not populated, the EntityKey represents a Google-managed entity such as a Google user or a Google Group. If populated, the EntityKey represents an external-identity-mapped group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "memberships",
					Description: `The list of memberships under the given group. Structure is documented below. The ` + "`" + `memberships` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `The MembershipRoles that apply to the Membership. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "member_key",
					Description: `(Optional) EntityKey of the member. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "preferred_member_key",
					Description: `(Optional) EntityKey of the member. Structure is documented below. The ` + "`" + `roles` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the MembershipRole. One of OWNER, MANAGER, MEMBER. The ` + "`" + `member_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the entity. For Google-managed entities, the id is the email address of an existing group or user. For external-identity-mapped entities, the id is a string conforming to the Identity Source's requirements.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace in which the entity exists. If not populated, the EntityKey represents a Google-managed entity such as a Google user or a Google Group. If populated, the EntityKey represents an external-identity-mapped group. The ` + "`" + `preferred_member_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the entity. For Google-managed entities, the id is the email address of an existing group or user. For external-identity-mapped entities, the id is a string conforming to the Identity Source's requirements.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace in which the entity exists. If not populated, the EntityKey represents a Google-managed entity such as a Google user or a Google Group. If populated, the EntityKey represents an external-identity-mapped group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_cloud_identity_groups",
			Category:         "Data Sources",
			ShortDescription: `Get list of the Cloud Identity Groups under a customer or namespace.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent",
					Description: `(Required) The parent resource under which to list all Groups. Must be of the form identitysources/{identity_source_id} for external- identity-mapped groups or customers/{customer_id} for Google Groups. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `The list of groups under the provided customer or namespace. Structure is documented below. The ` + "`" + `groups` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Resource name of the Group in the format: groups/{group_id}, where ` + "`" + `group_id` + "`" + ` is the unique ID assigned to the Group.`,
				},
				resource.Attribute{
					Name:        "group_key",
					Description: `EntityKey of the Group. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An extended description to help users determine the purpose of a Group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the entity. For Google-managed entities, the id is the email address of an existing group or user. For external-identity-mapped entities, the id is a string conforming to the Identity Source's requirements.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace in which the entity exists. If not populated, the EntityKey represents a Google-managed entity such as a Google user or a Google Group. If populated, the EntityKey represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of ` + "`" + `identitysources/{identity_source_id}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "groups",
					Description: `The list of groups under the provided customer or namespace. Structure is documented below. The ` + "`" + `groups` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Resource name of the Group in the format: groups/{group_id}, where ` + "`" + `group_id` + "`" + ` is the unique ID assigned to the Group.`,
				},
				resource.Attribute{
					Name:        "group_key",
					Description: `EntityKey of the Group. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An extended description to help users determine the purpose of a Group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the entity. For Google-managed entities, the id is the email address of an existing group or user. For external-identity-mapped entities, the id is a string conforming to the Identity Source's requirements.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace in which the entity exists. If not populated, the EntityKey represents a Google-managed entity such as a Google user or a Google Group. If populated, the EntityKey represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of ` + "`" + `identitysources/{identity_source_id}` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_cloud_run_locations",
			Category:         "Data Sources",
			ShortDescription: `Get Cloud Run locations available for a project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project to list versions for. If it is not provided, the provider project is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `The list of Cloud Run locations available for the given project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "locations",
					Description: `The list of Cloud Run locations available for the given project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_cloud_run_service",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Cloud Run Service.`,
			Description:      ``,
			Icon:             "Compute/Cloud_Run.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Cloud Run Service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location of the cloud run instance. eg us-central1 - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_cloud_run_service](https://www.terraform.io/docs/providers/google/r/cloud_run_service.html#argument-reference) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_cloudfunctions_function",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Cloud Function.`,
			Description:      ``,
			Icon:             "Compute/Cloud_Functions.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of a Cloud Function. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which the resource belongs. If it is not provided, the provider region is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Cloud Function.`,
				},
				resource.Attribute{
					Name:        "source_archive_bucket",
					Description: `The GCS bucket containing the zip archive which contains the function.`,
				},
				resource.Attribute{
					Name:        "source_archive_object",
					Description: `The source archive object (file) in archive bucket.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the function.`,
				},
				resource.Attribute{
					Name:        "available_memory_mb",
					Description: `Available memory (in MB) to the function.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Function execution timeout (in seconds).`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `The runtime in which the function is running.`,
				},
				resource.Attribute{
					Name:        "entry_point",
					Description: `Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.`,
				},
				resource.Attribute{
					Name:        "trigger_http",
					Description: `If function is triggered by HTTP, this boolean is set.`,
				},
				resource.Attribute{
					Name:        "event_trigger",
					Description: `A source that fires events in response to a condition in another service. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "https_trigger_url",
					Description: `If function is triggered by HTTP, trigger URL is set here.`,
				},
				resource.Attribute{
					Name:        "ingress_settings",
					Description: `Controls what traffic can reach the function.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this function.`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `The service account email to be assumed by the cloud function.`,
				},
				resource.Attribute{
					Name:        "vpc_connector",
					Description: `The VPC Network Connector that this cloud function can connect to.`,
				},
				resource.Attribute{
					Name:        "vpc_connector_egress_settings",
					Description: `The egress settings for the connector, controlling what traffic is diverted through it.`,
				},
				resource.Attribute{
					Name:        "max_instances",
					Description: `The limit on the maximum number of function instances that may coexist at a given time.`,
				},
				resource.Attribute{
					Name:        "source_repository",
					Description: `The URL of the Cloud Source Repository that the function is deployed from. Structure is documented below. The ` + "`" + `event_trigger` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `The type of event to observe. For example: ` + "`" + `"google.storage.object.finalize"` + "`" + `. See the documentation on [calling Cloud Functions](https://cloud.google.com/functions/docs/calling/) for a full reference of accepted triggers.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `The name of the resource whose events are being observed, for example, ` + "`" + `"myBucket"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "failure_policy",
					Description: `Policy for failed executions. Structure is documented below. The ` + "`" + `failure_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `Whether the function should be retried on failure. The ` + "`" + `source_repository` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL pointing to the hosted repository where the function is defined.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Cloud Function.`,
				},
				resource.Attribute{
					Name:        "source_archive_bucket",
					Description: `The GCS bucket containing the zip archive which contains the function.`,
				},
				resource.Attribute{
					Name:        "source_archive_object",
					Description: `The source archive object (file) in archive bucket.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the function.`,
				},
				resource.Attribute{
					Name:        "available_memory_mb",
					Description: `Available memory (in MB) to the function.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Function execution timeout (in seconds).`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `The runtime in which the function is running.`,
				},
				resource.Attribute{
					Name:        "entry_point",
					Description: `Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.`,
				},
				resource.Attribute{
					Name:        "trigger_http",
					Description: `If function is triggered by HTTP, this boolean is set.`,
				},
				resource.Attribute{
					Name:        "event_trigger",
					Description: `A source that fires events in response to a condition in another service. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "https_trigger_url",
					Description: `If function is triggered by HTTP, trigger URL is set here.`,
				},
				resource.Attribute{
					Name:        "ingress_settings",
					Description: `Controls what traffic can reach the function.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this function.`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `The service account email to be assumed by the cloud function.`,
				},
				resource.Attribute{
					Name:        "vpc_connector",
					Description: `The VPC Network Connector that this cloud function can connect to.`,
				},
				resource.Attribute{
					Name:        "vpc_connector_egress_settings",
					Description: `The egress settings for the connector, controlling what traffic is diverted through it.`,
				},
				resource.Attribute{
					Name:        "max_instances",
					Description: `The limit on the maximum number of function instances that may coexist at a given time.`,
				},
				resource.Attribute{
					Name:        "source_repository",
					Description: `The URL of the Cloud Source Repository that the function is deployed from. Structure is documented below. The ` + "`" + `event_trigger` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `The type of event to observe. For example: ` + "`" + `"google.storage.object.finalize"` + "`" + `. See the documentation on [calling Cloud Functions](https://cloud.google.com/functions/docs/calling/) for a full reference of accepted triggers.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `The name of the resource whose events are being observed, for example, ` + "`" + `"myBucket"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "failure_policy",
					Description: `Policy for failed executions. Structure is documented below. The ` + "`" + `failure_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `Whether the function should be retried on failure. The ` + "`" + `source_repository` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL pointing to the hosted repository where the function is defined.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_composer_environment",
			Category:         "Data Sources",
			ShortDescription: `Provides Cloud Composer environment configuration data.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the environment.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The location or Compute Engine region of the environment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier for the resource in format ` + "`" + `projects/{{project}}/locations/{{region}}/environments/{{name}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Configuration parameters for the environment. Full structure is provided by [composer environment resource documentation](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/composer_environment#config).`,
				},
				resource.Attribute{
					Name:        "config.0.gke_cluster",
					Description: `The Kubernetes Engine cluster used to run the environment.`,
				},
				resource.Attribute{
					Name:        "config.0.dag_gcs_prefix",
					Description: `The Cloud Storage prefix of the DAGs for the environment.`,
				},
				resource.Attribute{
					Name:        "config.0.airflow_uri",
					Description: `The URI of the Apache Airflow Web UI hosted within the environment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier for the resource in format ` + "`" + `projects/{{project}}/locations/{{region}}/environments/{{name}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Configuration parameters for the environment. Full structure is provided by [composer environment resource documentation](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/composer_environment#config).`,
				},
				resource.Attribute{
					Name:        "config.0.gke_cluster",
					Description: `The Kubernetes Engine cluster used to run the environment.`,
				},
				resource.Attribute{
					Name:        "config.0.dag_gcs_prefix",
					Description: `The Cloud Storage prefix of the DAGs for the environment.`,
				},
				resource.Attribute{
					Name:        "config.0.airflow_uri",
					Description: `The URI of the Apache Airflow Web UI hosted within the environment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_composer_image_versions",
			Category:         "Data Sources",
			ShortDescription: `Provides available Cloud Composer versions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project to list versions in. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The location to list versions in. If it is not provider, the provider region is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "image_versions",
					Description: `A list of composer image versions available in the given project and location. Each ` + "`" + `image_version` + "`" + ` contains:`,
				},
				resource.Attribute{
					Name:        "image_version_id",
					Description: `The string identifier of the image version, in the form: "composer-x.y.z-airflow-a.b(.c)"`,
				},
				resource.Attribute{
					Name:        "supported_python_versions",
					Description: `Supported python versions for this image version`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "image_versions",
					Description: `A list of composer image versions available in the given project and location. Each ` + "`" + `image_version` + "`" + ` contains:`,
				},
				resource.Attribute{
					Name:        "image_version_id",
					Description: `The string identifier of the image version, in the form: "composer-x.y.z-airflow-a.b(.c)"`,
				},
				resource.Attribute{
					Name:        "supported_python_versions",
					Description: `Supported python versions for this image version`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_address",
			Category:         "Data Sources",
			ShortDescription: `Get the IP address from a static address.`,
			Description:      ``,
			Icon:             "Storage/Persistent_Disk.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource, required by GCE. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The Region in which the created address reside. If it is not provided, the provider region is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP of the created resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates if the address is used. Possible values are: RESERVED or IN_USE.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP of the created resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates if the address is used. Possible values are: RESERVED or IN_USE.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_backend_bucket",
			Category:         "Data Sources",
			ShortDescription: `Get information about a BackendBucket.`,
			Description:      ``,
			Icon:             "Storage/Cloud_Storage.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_compute_backend_bucket](https://www.terraform.io/docs/providers/google/r/compute_backend_bucket.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_backend_service",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Backend Service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Backend Service. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "connection_draining_timeout_sec",
					Description: `Time for which instance will be drained (not accept new connections, but still work to finish started ones).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Textual description for the Backend Service.`,
				},
				resource.Attribute{
					Name:        "enable_cdn",
					Description: `Whether or not Cloud CDN is enabled on the Backend Service.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the Backend Service.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `projects/{{project}}/global/backendServices/{{name}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `The name of a service that has been added to an instance group in this backend.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for incoming requests.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the Backend Service.`,
				},
				resource.Attribute{
					Name:        "session_affinity",
					Description: `The Backend Service session stickiness configuration.`,
				},
				resource.Attribute{
					Name:        "timeout_sec",
					Description: `The number of seconds to wait for a backend to respond to a request before considering the request failed.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `The set of backends that serve this Backend Service.`,
				},
				resource.Attribute{
					Name:        "health_checks",
					Description: `The set of HTTP/HTTPS health checks used by the Backend Service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_draining_timeout_sec",
					Description: `Time for which instance will be drained (not accept new connections, but still work to finish started ones).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Textual description for the Backend Service.`,
				},
				resource.Attribute{
					Name:        "enable_cdn",
					Description: `Whether or not Cloud CDN is enabled on the Backend Service.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the Backend Service.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `projects/{{project}}/global/backendServices/{{name}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `The name of a service that has been added to an instance group in this backend.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for incoming requests.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the Backend Service.`,
				},
				resource.Attribute{
					Name:        "session_affinity",
					Description: `The Backend Service session stickiness configuration.`,
				},
				resource.Attribute{
					Name:        "timeout_sec",
					Description: `The number of seconds to wait for a backend to respond to a request before considering the request failed.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `The set of backends that serve this Backend Service.`,
				},
				resource.Attribute{
					Name:        "health_checks",
					Description: `The set of HTTP/HTTPS health checks used by the Backend Service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_default_service_account",
			Category:         "Data Sources",
			ShortDescription: `Retrieve default service account used by VMs running in this project`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project ID. If it is not provided, the provider project is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email address of the default service account used by VMs running in this project`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `The unique id of the service account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully-qualified name of the service account.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the service account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `Email address of the default service account used by VMs running in this project`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `The unique id of the service account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully-qualified name of the service account.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the service account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_forwarding_rule",
			Category:         "Data Sources",
			ShortDescription: `Get a regional forwarding rule within GCE.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the forwarding rule. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which the resource belongs. If it is not provided, the project region is used. ## Attributes Reference See [google_compute_forwarding_rule](https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_global_address",
			Category:         "Data Sources",
			ShortDescription: `Get the IP address from a static address reserved for a Global Forwarding Rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource, required by GCE. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP of the created resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates if the address is used. Possible values are: RESERVED or IN_USE.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP of the created resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates if the address is used. Possible values are: RESERVED or IN_USE.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_global_forwarding_rule",
			Category:         "Data Sources",
			ShortDescription: `Get a global forwarding rule within GCE.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the global forwarding rule. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_compute_global_forwarding_rule](https://www.terraform.io/docs/providers/google/r/compute_global_forwarding_rule.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_health_check",
			Category:         "Data Sources",
			ShortDescription: `Get information about a HealthCheck.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_compute_health_check](https://www.terraform.io/docs/providers/google/r/compute_health_check.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_image",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Compute Image.`,
			Description:      ``,
			Icon:             "Storage/Persistent_Disk.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. If you are using a [public base image][pubimg], be sure to specify the correct Image Project. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the data source with format ` + "`" + `projects/{{project}}/global/images/{{name}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the image.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The family name of the image.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `The size of the image when restored onto a persistent disk in gigabytes.`,
				},
				resource.Attribute{
					Name:        "archive_size_bytes",
					Description: `The size of the image tar.gz archive stored in Google Cloud Storage in bytes.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The unique identifier for the image.`,
				},
				resource.Attribute{
					Name:        "image_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this image.`,
				},
				resource.Attribute{
					Name:        "source_image_id",
					Description: `The ID value of the image used to create this image.`,
				},
				resource.Attribute{
					Name:        "source_disk",
					Description: `The URL of the source disk used to create this image.`,
				},
				resource.Attribute{
					Name:        "source_disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this image.`,
				},
				resource.Attribute{
					Name:        "source_disk_id",
					Description: `The ID value of the disk used to create this image.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `The creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of this image.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this image.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `A fingerprint for the labels being applied to this image.`,
				},
				resource.Attribute{
					Name:        "licenses",
					Description: `A list of applicable license URI.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image. Possible values are`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the data source with format ` + "`" + `projects/{{project}}/global/images/{{name}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the image.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The family name of the image.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `The size of the image when restored onto a persistent disk in gigabytes.`,
				},
				resource.Attribute{
					Name:        "archive_size_bytes",
					Description: `The size of the image tar.gz archive stored in Google Cloud Storage in bytes.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The unique identifier for the image.`,
				},
				resource.Attribute{
					Name:        "image_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this image.`,
				},
				resource.Attribute{
					Name:        "source_image_id",
					Description: `The ID value of the image used to create this image.`,
				},
				resource.Attribute{
					Name:        "source_disk",
					Description: `The URL of the source disk used to create this image.`,
				},
				resource.Attribute{
					Name:        "source_disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this image.`,
				},
				resource.Attribute{
					Name:        "source_disk_id",
					Description: `The ID value of the disk used to create this image.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `The creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of this image.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this image.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `A fingerprint for the labels being applied to this image.`,
				},
				resource.Attribute{
					Name:        "licenses",
					Description: `A list of applicable license URI.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image. Possible values are`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_instance",
			Category:         "Data Sources",
			ShortDescription: `Get a VM instance within GCE.`,
			Description:      ``,
			Icon:             "Compute/Compute_Engine.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `(Optional) The self link of the instance. One of ` + "`" + `name` + "`" + ` or ` + "`" + `self_link` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the instance. One of ` + "`" + `name` + "`" + ` or ` + "`" + `self_link` + "`" + ` must be provided. ---`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If ` + "`" + `self_link` + "`" + ` is provided, this value is ignored. If neither ` + "`" + `self_link` + "`" + ` nor ` + "`" + `project` + "`" + ` are provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone of the instance. If ` + "`" + `self_link` + "`" + ` is provided, this value is ignored. If neither ` + "`" + `self_link` + "`" + ` nor ` + "`" + `zone` + "`" + ` are provided, the provider zone is used. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `The boot disk for the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `The machine type to create.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `The networks attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "attached_disk",
					Description: `List of disks attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "can_ip_forward",
					Description: `Whether sending and receiving of packets with non-matching source or destination IPs is allowed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A brief description of the resource.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `Whether deletion protection is enabled on this instance.`,
				},
				resource.Attribute{
					Name:        "guest_accelerator",
					Description: `List of the type and count of accelerator cards attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata key/value pairs made available within the instance.`,
				},
				resource.Attribute{
					Name:        "min_cpu_platform",
					Description: `The minimum CPU platform specified for the VM instance.`,
				},
				resource.Attribute{
					Name:        "scheduling",
					Description: `The scheduling strategy being used by the instance.`,
				},
				resource.Attribute{
					Name:        "scratch_disk",
					Description: `The scratch disks attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `The service account to attach to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The list of tags attached to the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The server-assigned unique identifier of this instance.`,
				},
				resource.Attribute{
					Name:        "metadata_fingerprint",
					Description: `The unique fingerprint of the metadata.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "tags_fingerprint",
					Description: `The unique fingerprint of the tags.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The unique fingerprint of the labels.`,
				},
				resource.Attribute{
					Name:        "cpu_platform",
					Description: `The CPU platform used by this instance.`,
				},
				resource.Attribute{
					Name:        "shielded_instance_config",
					Description: `The shielded vm config being used by the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.network_ip",
					Description: `The internal ip address of the instance, either manually or dynamically assigned.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.access_config.0.nat_ip",
					Description: `If the instance has an access config, either the given external ip (in the ` + "`" + `nat_ip` + "`" + ` field) or the ephemeral (generated) ip (if you didn't provide one).`,
				},
				resource.Attribute{
					Name:        "attached_disk.0.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource.`,
				},
				resource.Attribute{
					Name:        "boot_disk.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource.`,
				},
				resource.Attribute{
					Name:        "disk.0.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `Whether the disk will be auto-deleted when the instance is deleted.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Name with which attached disk will be accessible under ` + "`" + `/dev/disk/by-id/` + "`" + ``,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `Parameters with which a disk was created alongside the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The name or self_link of an existing disk (such as those managed by ` + "`" + `google_compute_disk` + "`" + `) that was attached to the instance. The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in gigabytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The GCE disk type. One of ` + "`" + `pd-standard` + "`" + ` or ` + "`" + `pd-ssd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The image from which this disk was initialised. The ` + "`" + `scratch_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `The disk interface used for attaching this disk. One of ` + "`" + `SCSI` + "`" + ` or ` + "`" + `NVME` + "`" + `. The ` + "`" + `attached_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The name or self_link of the disk attached to this instance.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Name with which the attached disk is accessible under ` + "`" + `/dev/disk/by-id/` + "`" + ``,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Read/write mode for the disk. One of ` + "`" + `"READ_ONLY"` + "`" + ` or ` + "`" + `"READ_WRITE"` + "`" + `. The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The name or self_link of the network attached to this interface.`,
				},
				resource.Attribute{
					Name:        "network_ip",
					Description: `The private IP address assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "access_config",
					Description: `Access configurations, i.e. IPs via which this instance can be accessed via the Internet. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "alias_ip_range",
					Description: `An array of alias IP ranges for this network interface. Structure documented below. The ` + "`" + `access_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `The IP address that is be 1:1 mapped to the instance's network ip.`,
				},
				resource.Attribute{
					Name:        "public_ptr_domain_name",
					Description: `The DNS domain name for the public PTR record.`,
				},
				resource.Attribute{
					Name:        "network_tier",
					Description: `The [networking tier][network-tier] used for configuring this instance. One of ` + "`" + `PREMIUM` + "`" + ` or ` + "`" + `STANDARD` + "`" + `. The ` + "`" + `alias_ip_range` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `The IP CIDR range represented by this alias IP range.`,
				},
				resource.Attribute{
					Name:        "subnetwork_range_name",
					Description: `The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range. The ` + "`" + `service_account` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The service account e-mail address.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `A list of service scopes. The ` + "`" + `scheduling` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `Whether the instance is preemptible.`,
				},
				resource.Attribute{
					Name:        "on_host_maintenance",
					Description: `Describes maintenance behavior for the instance. One of ` + "`" + `MIGRATE` + "`" + ` or ` + "`" + `TERMINATE` + "`" + `, for more info, read [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options)`,
				},
				resource.Attribute{
					Name:        "automatic_restart",
					Description: `Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user). The ` + "`" + `guest_accelerator` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The accelerator type resource exposed to this instance. E.g. ` + "`" + `nvidia-tesla-k80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `The number of the guest accelerator cards exposed to this instance. [network-tier]: https://cloud.google.com/network-tiers/docs/overview The ` + "`" + `shielded_instance_config` + "`" + ` block supports:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "boot_disk",
					Description: `The boot disk for the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `The machine type to create.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `The networks attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "attached_disk",
					Description: `List of disks attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "can_ip_forward",
					Description: `Whether sending and receiving of packets with non-matching source or destination IPs is allowed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A brief description of the resource.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `Whether deletion protection is enabled on this instance.`,
				},
				resource.Attribute{
					Name:        "guest_accelerator",
					Description: `List of the type and count of accelerator cards attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata key/value pairs made available within the instance.`,
				},
				resource.Attribute{
					Name:        "min_cpu_platform",
					Description: `The minimum CPU platform specified for the VM instance.`,
				},
				resource.Attribute{
					Name:        "scheduling",
					Description: `The scheduling strategy being used by the instance.`,
				},
				resource.Attribute{
					Name:        "scratch_disk",
					Description: `The scratch disks attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `The service account to attach to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The list of tags attached to the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The server-assigned unique identifier of this instance.`,
				},
				resource.Attribute{
					Name:        "metadata_fingerprint",
					Description: `The unique fingerprint of the metadata.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "tags_fingerprint",
					Description: `The unique fingerprint of the tags.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The unique fingerprint of the labels.`,
				},
				resource.Attribute{
					Name:        "cpu_platform",
					Description: `The CPU platform used by this instance.`,
				},
				resource.Attribute{
					Name:        "shielded_instance_config",
					Description: `The shielded vm config being used by the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.network_ip",
					Description: `The internal ip address of the instance, either manually or dynamically assigned.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.access_config.0.nat_ip",
					Description: `If the instance has an access config, either the given external ip (in the ` + "`" + `nat_ip` + "`" + ` field) or the ephemeral (generated) ip (if you didn't provide one).`,
				},
				resource.Attribute{
					Name:        "attached_disk.0.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource.`,
				},
				resource.Attribute{
					Name:        "boot_disk.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource.`,
				},
				resource.Attribute{
					Name:        "disk.0.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `Whether the disk will be auto-deleted when the instance is deleted.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Name with which attached disk will be accessible under ` + "`" + `/dev/disk/by-id/` + "`" + ``,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `Parameters with which a disk was created alongside the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The name or self_link of an existing disk (such as those managed by ` + "`" + `google_compute_disk` + "`" + `) that was attached to the instance. The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in gigabytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The GCE disk type. One of ` + "`" + `pd-standard` + "`" + ` or ` + "`" + `pd-ssd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The image from which this disk was initialised. The ` + "`" + `scratch_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `The disk interface used for attaching this disk. One of ` + "`" + `SCSI` + "`" + ` or ` + "`" + `NVME` + "`" + `. The ` + "`" + `attached_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The name or self_link of the disk attached to this instance.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Name with which the attached disk is accessible under ` + "`" + `/dev/disk/by-id/` + "`" + ``,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Read/write mode for the disk. One of ` + "`" + `"READ_ONLY"` + "`" + ` or ` + "`" + `"READ_WRITE"` + "`" + `. The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The name or self_link of the network attached to this interface.`,
				},
				resource.Attribute{
					Name:        "network_ip",
					Description: `The private IP address assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "access_config",
					Description: `Access configurations, i.e. IPs via which this instance can be accessed via the Internet. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "alias_ip_range",
					Description: `An array of alias IP ranges for this network interface. Structure documented below. The ` + "`" + `access_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `The IP address that is be 1:1 mapped to the instance's network ip.`,
				},
				resource.Attribute{
					Name:        "public_ptr_domain_name",
					Description: `The DNS domain name for the public PTR record.`,
				},
				resource.Attribute{
					Name:        "network_tier",
					Description: `The [networking tier][network-tier] used for configuring this instance. One of ` + "`" + `PREMIUM` + "`" + ` or ` + "`" + `STANDARD` + "`" + `. The ` + "`" + `alias_ip_range` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `The IP CIDR range represented by this alias IP range.`,
				},
				resource.Attribute{
					Name:        "subnetwork_range_name",
					Description: `The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range. The ` + "`" + `service_account` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The service account e-mail address.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `A list of service scopes. The ` + "`" + `scheduling` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `Whether the instance is preemptible.`,
				},
				resource.Attribute{
					Name:        "on_host_maintenance",
					Description: `Describes maintenance behavior for the instance. One of ` + "`" + `MIGRATE` + "`" + ` or ` + "`" + `TERMINATE` + "`" + `, for more info, read [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options)`,
				},
				resource.Attribute{
					Name:        "automatic_restart",
					Description: `Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user). The ` + "`" + `guest_accelerator` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The accelerator type resource exposed to this instance. E.g. ` + "`" + `nvidia-tesla-k80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `The number of the guest accelerator cards exposed to this instance. [network-tier]: https://cloud.google.com/network-tiers/docs/overview The ` + "`" + `shielded_instance_config` + "`" + ` block supports:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_instance_group",
			Category:         "Data Sources",
			ShortDescription: `Get a Compute Instance Group within GCE.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the instance group. Either ` + "`" + `name` + "`" + ` or ` + "`" + `self_link` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `(Optional) The self link of the instance group. Either ` + "`" + `name` + "`" + ` or ` + "`" + `self_link` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone of the instance group. If referencing the instance group by name and ` + "`" + `zone` + "`" + ` is not provided, the provider zone is used. ## Attributes Reference The following arguments are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Textual description of the instance group.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `List of instances in the group.`,
				},
				resource.Attribute{
					Name:        "named_port",
					Description: `List of named ports in the group.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The URL of the network the instance group is in.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Textual description of the instance group.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `List of instances in the group.`,
				},
				resource.Attribute{
					Name:        "named_port",
					Description: `List of named ports in the group.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The URL of the network the instance group is in.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_instance_serial_port",
			Category:         "Data Sources",
			ShortDescription: `Get the serial port output from a Compute Instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name of the Compute Instance to read output from.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The number of the serial port to read output from. Possible values are 1-4. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the Compute Instance exists. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone in which the Compute Instance exists. If it is not provided, the provider zone is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "contents",
					Description: `The output of the serial port. Serial port output is available only when the VM instance is running, and logs are limited to the most recent 1 MB of output per port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "contents",
					Description: `The output of the serial port. Serial port output is available only when the VM instance is running, and logs are limited to the most recent 1 MB of output per port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_instance_template",
			Category:         "Data Sources",
			ShortDescription: `Get a VM instance template within GCE.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If ` + "`" + `project` + "`" + ` is not provideded, the provider project is used. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disks to attach to instances created from this template. This can be specified multiple times for multiple disks. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `The machine type to create. To create a machine with a [custom type][custom-vm-types] (such as extended memory), format the value like ` + "`" + `custom-VCPUS-MEM_IN_MB` + "`" + ` like ` + "`" + `custom-6-20480` + "`" + ` for 6 vCPU and 20GB of RAM.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance template. If you leave this blank, Terraform will auto-generate a unique name.`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `Creates a unique name beginning with the specified prefix. Conflicts with ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "can_ip_forward",
					Description: `Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A brief description of this resource.`,
				},
				resource.Attribute{
					Name:        "instance_description",
					Description: `A brief description to use for instances created from this template.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to instances created from this template,`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata key/value pairs to make available from within instances created from this template.`,
				},
				resource.Attribute{
					Name:        "metadata_startup_script",
					Description: `An alternative to using the startup-script metadata key, mostly to match the compute_instance resource. This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `Networks to attach to instances created from this template. This can be specified multiple times for multiple networks. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `An instance template is a global resource that is not bound to a zone or a region. However, you can still specify some regional resources in an instance template, which restricts the template to the region where that resource resides. For example, a custom ` + "`" + `subnetwork` + "`" + ` resource is tied to a specific region. Defaults to the region of the Provider if no value is given.`,
				},
				resource.Attribute{
					Name:        "scheduling",
					Description: `The scheduling strategy to use. More details about this configuration option are detailed below.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `Service account to attach to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags to attach to the instance.`,
				},
				resource.Attribute{
					Name:        "guest_accelerator",
					Description: `List of the type and count of accelerator cards attached to the instance. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "min_cpu_platform",
					Description: `Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as ` + "`" + `Intel Haswell` + "`" + ` or ` + "`" + `Intel Skylake` + "`" + `. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).`,
				},
				resource.Attribute{
					Name:        "shielded_instance_config",
					Description: `Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "enable_display",
					Description: `Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.`,
				},
				resource.Attribute{
					Name:        "confidential_instance_config",
					Description: `Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. The ` + "`" + `disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `Whether or not the disk should be auto-deleted. This defaults to true.`,
				},
				resource.Attribute{
					Name:        "boot",
					Description: `Indicates that this is a boot disk.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance. If not specified, the server chooses a default device name to apply to this disk.`,
				},
				resource.Attribute{
					Name:        "disk_name",
					Description: `Name of the disk. When not provided, this defaults to the name of the instance.`,
				},
				resource.Attribute{
					Name:        "source_image",
					Description: `The image from which to initialize this disk. This can be one of: the image's ` + "`" + `self_link` + "`" + `, ` + "`" + `projects/{project}/global/images/{image}` + "`" + `, ` + "`" + `projects/{project}/global/images/family/{family}` + "`" + `, ` + "`" + `global/images/{image}` + "`" + `, ` + "`" + `global/images/family/{family}` + "`" + `, ` + "`" + `family/{family}` + "`" + `, ` + "`" + `{project}/{family}` + "`" + `, ` + "`" + `{project}/{image}` + "`" + `, ` + "`" + `{family}` + "`" + `, or ` + "`" + `{image}` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI. Persistent disks must always use SCSI and the request will fail if you attempt to attach a persistent disk in any other format than SCSI. Local SSDs can use either NVME or SCSI.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If you are attaching or creating a boot disk, this must read-write mode.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The name (`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The GCE disk type. Can be either ` + "`" + `"pd-ssd"` + "`" + `, ` + "`" + `"local-ssd"` + "`" + `, ` + "`" + `"pd-balanced"` + "`" + ` or ` + "`" + `"pd-standard"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `The size of the image in gigabytes. If not specified, it will inherit the size of its base image. For SCRATCH disks, the size must be exactly 375GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of GCE disk, can be either ` + "`" + `"SCRATCH"` + "`" + ` or ` + "`" + `"PERSISTENT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_key",
					Description: `Encrypts or decrypts a disk using a customer-supplied encryption key. If you are creating a new disk, this field encrypts the new disk using an encryption key that you provide. If you are attaching an existing disk that is already encrypted, this field decrypts the disk using the customer-supplied encryption key. If you encrypt a disk using a customer-supplied key, you must provide the same key again when you attempt to use this resource at a later time. For example, you must provide the key when you create a snapshot or an image from the disk or when you attach the disk to a virtual machine instance. If you do not provide an encryption key, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later. Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt disks in a managed instance group.`,
				},
				resource.Attribute{
					Name:        "kms_key_self_link",
					Description: `The self link of the encryption key that is stored in Google Cloud KMS The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The name or self_link of the network to attach this interface to. Use ` + "`" + `network` + "`" + ` attribute for Legacy or Auto subnetted networks and ` + "`" + `subnetwork` + "`" + ` for custom subnetted networks.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `the name of the subnetwork to attach this interface to. The subnetwork must exist in the same ` + "`" + `region` + "`" + ` this instance will be created in. Either ` + "`" + `network` + "`" + ` or ` + "`" + `subnetwork` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "subnetwork_project",
					Description: `The ID of the project in which the subnetwork belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "network_ip",
					Description: `The private IP address to assign to the instance. If empty, the address will be automatically assigned.`,
				},
				resource.Attribute{
					Name:        "access_config",
					Description: `Access configurations, i.e. IPs via which this instance can be accessed via the Internet. Omit to ensure that the instance is not accessible from the Internet (this means that ssh provisioners will not work unless you are running Terraform can send traffic to the instance's network (e.g. via tunnel or because it is running on another cloud instance on that network). This block can be repeated multiple times. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "alias_ip_range",
					Description: `An array of alias IP ranges for this network interface. Can only be specified for network interfaces on subnet-mode networks. Structure documented below. The ` + "`" + `access_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `The IP address that will be 1:1 mapped to the instance's network ip. If not given, one will be generated.`,
				},
				resource.Attribute{
					Name:        "network_tier",
					Description: `The [networking tier][network-tier] used for configuring this instance template. This field can take the following values: PREMIUM or STANDARD. If this field is not specified, it is assumed to be PREMIUM. The ` + "`" + `alias_ip_range` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `The IP CIDR range represented by this alias IP range. This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces. At the time of writing only a netmask (e.g. /24) may be supplied, with a CIDR format resulting in an API error.`,
				},
				resource.Attribute{
					Name:        "subnetwork_range_name",
					Description: `The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range. If left unspecified, the primary range of the subnetwork will be used. The ` + "`" + `service_account` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The service account e-mail address. If not given, the default Google Compute Engine service account is used.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `A list of service scopes. Both OAuth2 URLs and gcloud short names are supported. To allow full access to all Cloud APIs, use the ` + "`" + `cloud-platform` + "`" + ` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes). The [service accounts documentation](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam) explains that access scopes are the legacy method of specifying permissions for your instance. If you are following best practices and using IAM roles to grant permissions to service accounts, then you can define this field as an empty list. The ` + "`" + `scheduling` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "automatic_restart",
					Description: `Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user). This defaults to true.`,
				},
				resource.Attribute{
					Name:        "on_host_maintenance",
					Description: `Defines the maintenance behavior for this instance.`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `Allows instance to be preempted. This defaults to false. Read more on this [here](https://cloud.google.com/compute/docs/instances/preemptible).`,
				},
				resource.Attribute{
					Name:        "node_affinities",
					Description: `Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems. Read more on sole-tenant node creation [here](https://cloud.google.com/compute/docs/nodes/create-nodes). Structure documented below. The ` + "`" + `guest_accelerator` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The accelerator type resource to expose to this instance. E.g. ` + "`" + `nvidia-tesla-k80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `The number of the guest accelerator cards exposed to this instance. The ` + "`" + `node_affinities` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key for the node affinity label.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `The operator. Can be ` + "`" + `IN` + "`" + ` for node-affinities or ` + "`" + `NOT_IN` + "`" + ` for anti-affinities.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The values for the node affinity label. The ` + "`" + `shielded_instance_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `projects/{{project}}/global/instanceTemplates/{{name}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "metadata_fingerprint",
					Description: `The unique fingerprint of the metadata.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "tags_fingerprint",
					Description: `The unique fingerprint of the tags. [1]: /docs/providers/google/r/compute_instance_group_manager.html [2]: /docs/configuration/resources.html#lifecycle`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disk",
					Description: `Disks to attach to instances created from this template. This can be specified multiple times for multiple disks. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `The machine type to create. To create a machine with a [custom type][custom-vm-types] (such as extended memory), format the value like ` + "`" + `custom-VCPUS-MEM_IN_MB` + "`" + ` like ` + "`" + `custom-6-20480` + "`" + ` for 6 vCPU and 20GB of RAM.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance template. If you leave this blank, Terraform will auto-generate a unique name.`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `Creates a unique name beginning with the specified prefix. Conflicts with ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "can_ip_forward",
					Description: `Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A brief description of this resource.`,
				},
				resource.Attribute{
					Name:        "instance_description",
					Description: `A brief description to use for instances created from this template.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to instances created from this template,`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata key/value pairs to make available from within instances created from this template.`,
				},
				resource.Attribute{
					Name:        "metadata_startup_script",
					Description: `An alternative to using the startup-script metadata key, mostly to match the compute_instance resource. This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `Networks to attach to instances created from this template. This can be specified multiple times for multiple networks. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `An instance template is a global resource that is not bound to a zone or a region. However, you can still specify some regional resources in an instance template, which restricts the template to the region where that resource resides. For example, a custom ` + "`" + `subnetwork` + "`" + ` resource is tied to a specific region. Defaults to the region of the Provider if no value is given.`,
				},
				resource.Attribute{
					Name:        "scheduling",
					Description: `The scheduling strategy to use. More details about this configuration option are detailed below.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `Service account to attach to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags to attach to the instance.`,
				},
				resource.Attribute{
					Name:        "guest_accelerator",
					Description: `List of the type and count of accelerator cards attached to the instance. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "min_cpu_platform",
					Description: `Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as ` + "`" + `Intel Haswell` + "`" + ` or ` + "`" + `Intel Skylake` + "`" + `. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).`,
				},
				resource.Attribute{
					Name:        "shielded_instance_config",
					Description: `Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "enable_display",
					Description: `Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.`,
				},
				resource.Attribute{
					Name:        "confidential_instance_config",
					Description: `Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. The ` + "`" + `disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `Whether or not the disk should be auto-deleted. This defaults to true.`,
				},
				resource.Attribute{
					Name:        "boot",
					Description: `Indicates that this is a boot disk.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance. If not specified, the server chooses a default device name to apply to this disk.`,
				},
				resource.Attribute{
					Name:        "disk_name",
					Description: `Name of the disk. When not provided, this defaults to the name of the instance.`,
				},
				resource.Attribute{
					Name:        "source_image",
					Description: `The image from which to initialize this disk. This can be one of: the image's ` + "`" + `self_link` + "`" + `, ` + "`" + `projects/{project}/global/images/{image}` + "`" + `, ` + "`" + `projects/{project}/global/images/family/{family}` + "`" + `, ` + "`" + `global/images/{image}` + "`" + `, ` + "`" + `global/images/family/{family}` + "`" + `, ` + "`" + `family/{family}` + "`" + `, ` + "`" + `{project}/{family}` + "`" + `, ` + "`" + `{project}/{image}` + "`" + `, ` + "`" + `{family}` + "`" + `, or ` + "`" + `{image}` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI. Persistent disks must always use SCSI and the request will fail if you attempt to attach a persistent disk in any other format than SCSI. Local SSDs can use either NVME or SCSI.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If you are attaching or creating a boot disk, this must read-write mode.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The name (`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The GCE disk type. Can be either ` + "`" + `"pd-ssd"` + "`" + `, ` + "`" + `"local-ssd"` + "`" + `, ` + "`" + `"pd-balanced"` + "`" + ` or ` + "`" + `"pd-standard"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `The size of the image in gigabytes. If not specified, it will inherit the size of its base image. For SCRATCH disks, the size must be exactly 375GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of GCE disk, can be either ` + "`" + `"SCRATCH"` + "`" + ` or ` + "`" + `"PERSISTENT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_key",
					Description: `Encrypts or decrypts a disk using a customer-supplied encryption key. If you are creating a new disk, this field encrypts the new disk using an encryption key that you provide. If you are attaching an existing disk that is already encrypted, this field decrypts the disk using the customer-supplied encryption key. If you encrypt a disk using a customer-supplied key, you must provide the same key again when you attempt to use this resource at a later time. For example, you must provide the key when you create a snapshot or an image from the disk or when you attach the disk to a virtual machine instance. If you do not provide an encryption key, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later. Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt disks in a managed instance group.`,
				},
				resource.Attribute{
					Name:        "kms_key_self_link",
					Description: `The self link of the encryption key that is stored in Google Cloud KMS The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The name or self_link of the network to attach this interface to. Use ` + "`" + `network` + "`" + ` attribute for Legacy or Auto subnetted networks and ` + "`" + `subnetwork` + "`" + ` for custom subnetted networks.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `the name of the subnetwork to attach this interface to. The subnetwork must exist in the same ` + "`" + `region` + "`" + ` this instance will be created in. Either ` + "`" + `network` + "`" + ` or ` + "`" + `subnetwork` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "subnetwork_project",
					Description: `The ID of the project in which the subnetwork belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "network_ip",
					Description: `The private IP address to assign to the instance. If empty, the address will be automatically assigned.`,
				},
				resource.Attribute{
					Name:        "access_config",
					Description: `Access configurations, i.e. IPs via which this instance can be accessed via the Internet. Omit to ensure that the instance is not accessible from the Internet (this means that ssh provisioners will not work unless you are running Terraform can send traffic to the instance's network (e.g. via tunnel or because it is running on another cloud instance on that network). This block can be repeated multiple times. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "alias_ip_range",
					Description: `An array of alias IP ranges for this network interface. Can only be specified for network interfaces on subnet-mode networks. Structure documented below. The ` + "`" + `access_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `The IP address that will be 1:1 mapped to the instance's network ip. If not given, one will be generated.`,
				},
				resource.Attribute{
					Name:        "network_tier",
					Description: `The [networking tier][network-tier] used for configuring this instance template. This field can take the following values: PREMIUM or STANDARD. If this field is not specified, it is assumed to be PREMIUM. The ` + "`" + `alias_ip_range` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `The IP CIDR range represented by this alias IP range. This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces. At the time of writing only a netmask (e.g. /24) may be supplied, with a CIDR format resulting in an API error.`,
				},
				resource.Attribute{
					Name:        "subnetwork_range_name",
					Description: `The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range. If left unspecified, the primary range of the subnetwork will be used. The ` + "`" + `service_account` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The service account e-mail address. If not given, the default Google Compute Engine service account is used.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `A list of service scopes. Both OAuth2 URLs and gcloud short names are supported. To allow full access to all Cloud APIs, use the ` + "`" + `cloud-platform` + "`" + ` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes). The [service accounts documentation](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam) explains that access scopes are the legacy method of specifying permissions for your instance. If you are following best practices and using IAM roles to grant permissions to service accounts, then you can define this field as an empty list. The ` + "`" + `scheduling` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "automatic_restart",
					Description: `Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user). This defaults to true.`,
				},
				resource.Attribute{
					Name:        "on_host_maintenance",
					Description: `Defines the maintenance behavior for this instance.`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `Allows instance to be preempted. This defaults to false. Read more on this [here](https://cloud.google.com/compute/docs/instances/preemptible).`,
				},
				resource.Attribute{
					Name:        "node_affinities",
					Description: `Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems. Read more on sole-tenant node creation [here](https://cloud.google.com/compute/docs/nodes/create-nodes). Structure documented below. The ` + "`" + `guest_accelerator` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The accelerator type resource to expose to this instance. E.g. ` + "`" + `nvidia-tesla-k80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `The number of the guest accelerator cards exposed to this instance. The ` + "`" + `node_affinities` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key for the node affinity label.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `The operator. Can be ` + "`" + `IN` + "`" + ` for node-affinities or ` + "`" + `NOT_IN` + "`" + ` for anti-affinities.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The values for the node affinity label. The ` + "`" + `shielded_instance_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `projects/{{project}}/global/instanceTemplates/{{name}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "metadata_fingerprint",
					Description: `The unique fingerprint of the metadata.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "tags_fingerprint",
					Description: `The unique fingerprint of the tags. [1]: /docs/providers/google/r/compute_instance_group_manager.html [2]: /docs/configuration/resources.html#lifecycle`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_lb_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Get information about the IP ranges used when health-checking load balancers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network",
					Description: `The IP ranges used for health checks when`,
				},
				resource.Attribute{
					Name:        "http_ssl_tcp_internal",
					Description: `The IP ranges used for health checks when`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network",
					Description: `The IP ranges used for health checks when`,
				},
				resource.Attribute{
					Name:        "http_ssl_tcp_internal",
					Description: `The IP ranges used for health checks when`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_network",
			Category:         "Data Sources",
			ShortDescription: `Get a network within GCE.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the network. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format projects/{{project}}/global/networks/{{name}}`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this network.`,
				},
				resource.Attribute{
					Name:        "gateway_ipv4",
					Description: `The IP address of the gateway.`,
				},
				resource.Attribute{
					Name:        "subnetworks_self_links",
					Description: `the list of subnetworks which belong to the network`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format projects/{{project}}/global/networks/{{name}}`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this network.`,
				},
				resource.Attribute{
					Name:        "gateway_ipv4",
					Description: `The IP address of the gateway.`,
				},
				resource.Attribute{
					Name:        "subnetworks_self_links",
					Description: `the list of subnetworks which belong to the network`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_network_endpoint_group",
			Category:         "Data Sources",
			ShortDescription: `Retrieve Network Endpoint Group's details.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project to list versions in. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Network Endpoint Group name. Provide either this or a ` + "`" + `self_link` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The Network Endpoint Group availability zone.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `(Optional) The Network Endpoint Group self\_link. ## Attributes Reference In addition the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The network to which all network endpoints in the NEG belong.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `subnetwork to which all network endpoints in the NEG belong.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The NEG description.`,
				},
				resource.Attribute{
					Name:        "network_endpoint_type",
					Description: `Type of network endpoints in this network endpoint group.`,
				},
				resource.Attribute{
					Name:        "default_port",
					Description: `The NEG default port.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Number of network endpoints in the network endpoint group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network",
					Description: `The network to which all network endpoints in the NEG belong.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `subnetwork to which all network endpoints in the NEG belong.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The NEG description.`,
				},
				resource.Attribute{
					Name:        "network_endpoint_type",
					Description: `Type of network endpoints in this network endpoint group.`,
				},
				resource.Attribute{
					Name:        "default_port",
					Description: `The NEG default port.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Number of network endpoints in the network endpoint group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_node_types",
			Category:         "Data Sources",
			ShortDescription: `Provides list of available Google Compute Engine node types for sole-tenant nodes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of node types available in the given zone and project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of node types available in the given zone and project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_region_instance_group",
			Category:         "Data Sources",
			ShortDescription: `Get the instances inside a Compute Region Instance Group within GCE.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the instance group. One of ` + "`" + `name` + "`" + ` or ` + "`" + `self_link` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `(Optional) The link to the instance group. One of ` + "`" + `name` + "`" + ` or ` + "`" + `self_link` + "`" + ` must be provided. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If ` + "`" + `self_link` + "`" + ` is provided, this value is ignored. If neither ` + "`" + `self_link` + "`" + ` nor ` + "`" + `project` + "`" + ` are provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which the resource belongs. If ` + "`" + `self_link` + "`" + ` is provided, this value is ignored. If neither ` + "`" + `self_link` + "`" + ` nor ` + "`" + `region` + "`" + ` are provided, the provider region is used. ## Attributes Reference The following arguments are exported:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the group.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `List of instances in the group, as a list of resources, each containing:`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `URL to the instance.`,
				},
				resource.Attribute{
					Name:        "named_ports",
					Description: `List of named ports in the group, as a list of resources, each containing:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Integer port number`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `String port name`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `String description of current state of the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the group.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `List of instances in the group, as a list of resources, each containing:`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `URL to the instance.`,
				},
				resource.Attribute{
					Name:        "named_ports",
					Description: `List of named ports in the group, as a list of resources, each containing:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Integer port number`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `String port name`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `String description of current state of the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_region_ssl_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get info about a Regional Google Compute SSL Certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which the resource belongs. If it is not provided, the provider region is used. ## Attributes Reference See [google_compute_region_ssl_certificate](https://www.terraform.io/docs/providers/google/r/compute_region_ssl_certificate.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_regions",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available Google Compute regions`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of regions available in the given project`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of regions available in the given project`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_resource_policy",
			Category:         "Data Sources",
			ShortDescription: `Provide access to a Resource Policy's attributes`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of this Resource Policy.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of this Resource Policy.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_router",
			Category:         "Data Sources",
			ShortDescription: `Get a Cloud Router within GCE.`,
			Description:      ``,
			Icon:             "Networking/Cloud_Router.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the router.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The VPC network on which this router lives.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region this router has been created in. If unspecified, this defaults to the region configured in the provider. ## Attributes Reference See [google_compute_router](https://www.terraform.io/docs/providers/google/r/compute_router.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_ssl_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get info about a Google Compute SSL Certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_compute_ssl_certificate](https://www.terraform.io/docs/providers/google/r/compute_ssl_certificate.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_ssl_policy",
			Category:         "Data Sources",
			ShortDescription: `Gets an SSL Policy within GCE, for use with Target HTTPS and Target SSL Proxies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSL Policy. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "enabled_features",
					Description: `The set of enabled encryption ciphers as a result of the policy config`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this SSL Policy.`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `The minimum supported TLS version of this policy.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `The Google-curated or custom profile used by this policy.`,
				},
				resource.Attribute{
					Name:        "custom_features",
					Description: `If the ` + "`" + `profile` + "`" + ` is ` + "`" + `CUSTOM` + "`" + `, these are the custom encryption ciphers supported by the profile. If the ` + "`" + `profile` + "`" + ` is`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled_features",
					Description: `The set of enabled encryption ciphers as a result of the policy config`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this SSL Policy.`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `The minimum supported TLS version of this policy.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `The Google-curated or custom profile used by this policy.`,
				},
				resource.Attribute{
					Name:        "custom_features",
					Description: `If the ` + "`" + `profile` + "`" + ` is ` + "`" + `CUSTOM` + "`" + `, these are the custom encryption ciphers supported by the profile. If the ` + "`" + `profile` + "`" + ` is`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_subnetwork",
			Category:         "Data Sources",
			ShortDescription: `Get a subnetwork within GCE.`,
			Description:      ``,
			Icon:             "Networking/Cloud_Network.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `(Optional) The self link of the subnetwork. If ` + "`" + `self_link` + "`" + ` is specified, ` + "`" + `name` + "`" + `, ` + "`" + `project` + "`" + `, and ` + "`" + `region` + "`" + ` are ignored.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnetwork. One of ` + "`" + `name` + "`" + ` or ` + "`" + `self_link` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region this subnetwork has been created in. If unspecified, this defaults to the region configured in the provider. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The network name or resource link to the parent network of this subnetwork.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this subnetwork.`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `The IP address range that machines in this network are assigned to, represented as a CIDR block.`,
				},
				resource.Attribute{
					Name:        "gateway_address",
					Description: `The IP address of the gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip_google_access",
					Description: `Whether the VMs in this subnet can access Google services without assigned external IP addresses.`,
				},
				resource.Attribute{
					Name:        "secondary_ip_range",
					Description: `An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. Structure is documented below. The ` + "`" + `secondary_ip_range` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "range_name",
					Description: `The name associated with this subnetwork secondary range, used when adding an alias IP range to a VM instance.`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `The range of IP addresses belonging to this subnetwork secondary range.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network",
					Description: `The network name or resource link to the parent network of this subnetwork.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this subnetwork.`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `The IP address range that machines in this network are assigned to, represented as a CIDR block.`,
				},
				resource.Attribute{
					Name:        "gateway_address",
					Description: `The IP address of the gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip_google_access",
					Description: `Whether the VMs in this subnet can access Google services without assigned external IP addresses.`,
				},
				resource.Attribute{
					Name:        "secondary_ip_range",
					Description: `An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. Structure is documented below. The ` + "`" + `secondary_ip_range` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "range_name",
					Description: `The name associated with this subnetwork secondary range, used when adding an alias IP range to a VM instance.`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `The range of IP addresses belonging to this subnetwork secondary range.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_vpn_gateway",
			Category:         "Data Sources",
			ShortDescription: `Get a VPN gateway within GCE.`,
			Description:      ``,
			Icon:             "Networking/Cloud_VPN.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPN gateway. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which the resource belongs. If it is not provided, the project region is used. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The network of this VPN gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this VPN gateway.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of this VPN gateway.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network",
					Description: `The network of this VPN gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this VPN gateway.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of this VPN gateway.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available Google Compute zones`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of zones available in the given region`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of zones available in the given region`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_container_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get info about a Google Kubernetes Engine cluster.`,
			Description:      ``,
			Icon:             "Compute/Kubernetes_Engine.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_container_cluster](https://www.terraform.io/docs/providers/google/r/container_cluster.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_container_engine_versions",
			Category:         "Data Sources",
			ShortDescription: `Provides lists of available Google Kubernetes Engine versions for masters and nodes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "valid_master_versions",
					Description: `A list of versions available in the given zone for use with master instances.`,
				},
				resource.Attribute{
					Name:        "valid_node_versions",
					Description: `A list of versions available in the given zone for use with node instances.`,
				},
				resource.Attribute{
					Name:        "latest_master_version",
					Description: `The latest version available in the given zone for use with master instances.`,
				},
				resource.Attribute{
					Name:        "latest_node_version",
					Description: `The latest version available in the given zone for use with node instances.`,
				},
				resource.Attribute{
					Name:        "default_cluster_version",
					Description: `Version of Kubernetes the service deploys by default.`,
				},
				resource.Attribute{
					Name:        "release_channel_default_version",
					Description: `A map from a release channel name to the channel's default version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "valid_master_versions",
					Description: `A list of versions available in the given zone for use with master instances.`,
				},
				resource.Attribute{
					Name:        "valid_node_versions",
					Description: `A list of versions available in the given zone for use with node instances.`,
				},
				resource.Attribute{
					Name:        "latest_master_version",
					Description: `The latest version available in the given zone for use with master instances.`,
				},
				resource.Attribute{
					Name:        "latest_node_version",
					Description: `The latest version available in the given zone for use with node instances.`,
				},
				resource.Attribute{
					Name:        "default_cluster_version",
					Description: `Version of Kubernetes the service deploys by default.`,
				},
				resource.Attribute{
					Name:        "release_channel_default_version",
					Description: `A map from a release channel name to the channel's default version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_container_registry_image",
			Category:         "Data Sources",
			ShortDescription: `Get URLs for a given project's container registry image.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_container_registry_repository",
			Category:         "Data Sources",
			ShortDescription: `Get URLs for a given project's container registry repository.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dns_keys",
			Category:         "Data Sources",
			ShortDescription: `Get DNSKEY and DS records of DNSSEC-signed managed zones.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "managed_zone",
					Description: `(Required) The name or id of the Cloud DNS managed zone.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If ` + "`" + `project` + "`" + ` is not provided, the provider project is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_signing_keys",
					Description: `A list of Key-signing key (KSK) records. Structure is documented below. Additionally, the DS record is provided:`,
				},
				resource.Attribute{
					Name:        "ds_record",
					Description: `The DS record based on the KSK record. This is used when [delegating](https://cloud.google.com/dns/docs/dnssec-advanced#subdelegation) DNSSEC-signed subdomains.`,
				},
				resource.Attribute{
					Name:        "zone_signing_keys",
					Description: `A list of Zone-signing key (ZSK) records. Structure is documented below. --- The ` + "`" + `key_signing_keys` + "`" + ` and ` + "`" + `zone_signing_keys` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `String mnemonic specifying the DNSSEC algorithm of this key. Immutable after creation time. Possible values are ` + "`" + `ecdsap256sha256` + "`" + `, ` + "`" + `ecdsap384sha384` + "`" + `, ` + "`" + `rsasha1` + "`" + `, ` + "`" + `rsasha256` + "`" + `, and ` + "`" + `rsasha512` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `The time that this resource was created in the control plane. This is in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A mutable string of at most 1024 characters associated with this resource for the user's convenience.`,
				},
				resource.Attribute{
					Name:        "digests",
					Description: `A list of cryptographic hashes of the DNSKEY resource record associated with this DnsKey. These digests are needed to construct a DS record that points at this DNS key. Each contains: - ` + "`" + `digest` + "`" + ` - The base-16 encoded bytes of this digest. Suitable for use in a DS resource record. - ` + "`" + `type` + "`" + ` - Specifies the algorithm used to calculate this digest. Possible values are ` + "`" + `sha1` + "`" + `, ` + "`" + `sha256` + "`" + ` and ` + "`" + `sha384` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier for the resource; defined by the server.`,
				},
				resource.Attribute{
					Name:        "is_active",
					Description: `Active keys will be used to sign subsequent changes to the ManagedZone. Inactive keys will still be present as DNSKEY Resource Records for the use of resolvers validating existing signatures.`,
				},
				resource.Attribute{
					Name:        "key_length",
					Description: `Length of the key in bits. Specified at creation time then immutable.`,
				},
				resource.Attribute{
					Name:        "key_tag",
					Description: `The key tag is a non-cryptographic hash of the a DNSKEY resource record associated with this DnsKey. The key tag can be used to identify a DNSKEY more quickly (but it is not a unique identifier). In particular, the key tag is used in a parent zone's DS record to point at the DNSKEY in this child ManagedZone. The key tag is a number in the range [0, 65535] and the algorithm to calculate it is specified in RFC4034 Appendix B.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Base64 encoded public half of this key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_signing_keys",
					Description: `A list of Key-signing key (KSK) records. Structure is documented below. Additionally, the DS record is provided:`,
				},
				resource.Attribute{
					Name:        "ds_record",
					Description: `The DS record based on the KSK record. This is used when [delegating](https://cloud.google.com/dns/docs/dnssec-advanced#subdelegation) DNSSEC-signed subdomains.`,
				},
				resource.Attribute{
					Name:        "zone_signing_keys",
					Description: `A list of Zone-signing key (ZSK) records. Structure is documented below. --- The ` + "`" + `key_signing_keys` + "`" + ` and ` + "`" + `zone_signing_keys` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `String mnemonic specifying the DNSSEC algorithm of this key. Immutable after creation time. Possible values are ` + "`" + `ecdsap256sha256` + "`" + `, ` + "`" + `ecdsap384sha384` + "`" + `, ` + "`" + `rsasha1` + "`" + `, ` + "`" + `rsasha256` + "`" + `, and ` + "`" + `rsasha512` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `The time that this resource was created in the control plane. This is in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A mutable string of at most 1024 characters associated with this resource for the user's convenience.`,
				},
				resource.Attribute{
					Name:        "digests",
					Description: `A list of cryptographic hashes of the DNSKEY resource record associated with this DnsKey. These digests are needed to construct a DS record that points at this DNS key. Each contains: - ` + "`" + `digest` + "`" + ` - The base-16 encoded bytes of this digest. Suitable for use in a DS resource record. - ` + "`" + `type` + "`" + ` - Specifies the algorithm used to calculate this digest. Possible values are ` + "`" + `sha1` + "`" + `, ` + "`" + `sha256` + "`" + ` and ` + "`" + `sha384` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier for the resource; defined by the server.`,
				},
				resource.Attribute{
					Name:        "is_active",
					Description: `Active keys will be used to sign subsequent changes to the ManagedZone. Inactive keys will still be present as DNSKEY Resource Records for the use of resolvers validating existing signatures.`,
				},
				resource.Attribute{
					Name:        "key_length",
					Description: `Length of the key in bits. Specified at creation time then immutable.`,
				},
				resource.Attribute{
					Name:        "key_tag",
					Description: `The key tag is a non-cryptographic hash of the a DNSKEY resource record associated with this DnsKey. The key tag can be used to identify a DNSKEY more quickly (but it is not a unique identifier). In particular, the key tag is used in a parent zone's DS record to point at the DNSKEY in this child ManagedZone. The key tag is a number in the range [0, 65535] and the algorithm to calculate it is specified in RFC4034 Appendix B.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Base64 encoded public half of this key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dns_managed_zone",
			Category:         "Data Sources",
			ShortDescription: `Provides access to the attributes of a zone within Google Cloud DNS`,
			Description:      ``,
			Icon:             "Networking/Cloud_DNS.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project for the Google Cloud DNS zone. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The fully qualified DNS name of this zone, e.g. ` + "`" + `terraform.io.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A textual description field.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `The list of nameservers that will be authoritative for this domain. Use NS records to redirect from your DNS provider to these names, thus making Google Cloud DNS authoritative for this zone.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dns_name",
					Description: `The fully qualified DNS name of this zone, e.g. ` + "`" + `terraform.io.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A textual description field.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `The list of nameservers that will be authoritative for this domain. Use NS records to redirect from your DNS provider to these names, thus making Google Cloud DNS authoritative for this zone.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_firebase_web_app",
			Category:         "Data Sources",
			ShortDescription: `A Google Cloud Firebase web application instance`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) The app_ip of name of the Firebase webApp. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `{{name}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully qualified resource name of the App, for example: projects/projectId/webApps/appId`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque token, as the data format is not specified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `{{name}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully qualified resource name of the App, for example: projects/projectId/webApps/appId`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque token, as the data format is not specified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_firebase_web_app_config",
			Category:         "Data Sources",
			ShortDescription: `A Google Cloud Firebase web application configuration`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "web_app_id",
					Description: `(Required) the id of the firebase web app - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `The API key associated with the web App.`,
				},
				resource.Attribute{
					Name:        "auth_domain",
					Description: `The domain Firebase Auth configures for OAuth redirects, in the format: projectId.firebaseapp.com`,
				},
				resource.Attribute{
					Name:        "database_url",
					Description: `The default Firebase Realtime Database URL.`,
				},
				resource.Attribute{
					Name:        "storage_bucket",
					Description: `The default Cloud Storage for Firebase storage bucket name.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `The ID of the project's default GCP resource location. The location is one of the available GCP resource locations. This field is omitted if the default GCP resource location has not been finalized yet. To set your project's default GCP resource location, call defaultLocation.finalize after you add Firebase services to your project.`,
				},
				resource.Attribute{
					Name:        "messaging_sender_id",
					Description: `The sender ID for use with Firebase Cloud Messaging.`,
				},
				resource.Attribute{
					Name:        "measurement_id",
					Description: `The unique Google-assigned identifier of the Google Analytics web stream associated with the Firebase Web App. Firebase SDKs use this ID to interact with Google Analytics APIs. This field is only present if the App is linked to a web stream in a Google Analytics App + Web property. Learn more about this ID and Google Analytics web streams in the Analytics documentation. To generate a measurementId and link the Web App with a Google Analytics web stream, call projects.addGoogleAnalytics.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_key",
					Description: `The API key associated with the web App.`,
				},
				resource.Attribute{
					Name:        "auth_domain",
					Description: `The domain Firebase Auth configures for OAuth redirects, in the format: projectId.firebaseapp.com`,
				},
				resource.Attribute{
					Name:        "database_url",
					Description: `The default Firebase Realtime Database URL.`,
				},
				resource.Attribute{
					Name:        "storage_bucket",
					Description: `The default Cloud Storage for Firebase storage bucket name.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `The ID of the project's default GCP resource location. The location is one of the available GCP resource locations. This field is omitted if the default GCP resource location has not been finalized yet. To set your project's default GCP resource location, call defaultLocation.finalize after you add Firebase services to your project.`,
				},
				resource.Attribute{
					Name:        "messaging_sender_id",
					Description: `The sender ID for use with Firebase Cloud Messaging.`,
				},
				resource.Attribute{
					Name:        "measurement_id",
					Description: `The unique Google-assigned identifier of the Google Analytics web stream associated with the Firebase Web App. Firebase SDKs use this ID to interact with Google Analytics APIs. This field is only present if the App is linked to a web stream in a Google Analytics App + Web property. Learn more about this ID and Google Analytics web streams in the Analytics documentation. To generate a measurementId and link the Web App with a Google Analytics web stream, call projects.addGoogleAnalytics.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_folder",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Cloud Folder.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Folder ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the Folder in the form ` + "`" + `folders/{folder_id}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `The resource name of the parent Folder or Organization.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The folder's display name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `The Folder's current lifecycle state.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `If ` + "`" + `lookup_organization` + "`" + ` is enable, the resource name of the Organization that the folder belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Folder ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the Folder in the form ` + "`" + `folders/{folder_id}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `The resource name of the parent Folder or Organization.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The folder's display name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `The Folder's current lifecycle state.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `If ` + "`" + `lookup_organization` + "`" + ` is enable, the resource name of the Organization that the folder belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_folder_organization_policy",
			Category:         "Data Sources",
			ShortDescription: `Retrieve Organization policies for a Google Folder`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder",
					Description: `(Required) The resource name of the folder to set the policy for. Its format is folders/{folder_id}.`,
				},
				resource.Attribute{
					Name:        "constraint",
					Description: `(Required) (Required) The name of the Constraint the Policy is configuring, for example, ` + "`" + `serviceuser.services` + "`" + `. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints). ## Attributes Reference See [google_folder_organization_policy](https://www.terraform.io/docs/providers/google/r/google_folder_organization_policy.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_game_services_game_server_deployment_rollout",
			Category:         "Data Sources",
			ShortDescription: `Get the rollout state.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "deployment_id",
					Description: `(Required) The deployment to get the rollout state from. Only 1 rollout must be associated with each deployment. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "default_game_server_config",
					Description: `This field points to the game server config that is applied by default to all realms and clusters. For example, ` + "`" + `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "game_server_config_overrides",
					Description: `The game_server_config_overrides contains the per game server config overrides. The overrides are processed in the order they are listed. As soon as a match is found for a cluster, the rest of the list is not processed. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `game_server_config_overrides` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "realms_selector",
					Description: `Selection by realms. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "config_version",
					Description: `Version of the configuration. The ` + "`" + `realms_selector` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "realms",
					Description: `List of realms to match against.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource id of the game server deployment eg: ` + "`" + `projects/my-project/locations/global/gameServerDeployments/my-deployment/rollout` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "default_game_server_config",
					Description: `This field points to the game server config that is applied by default to all realms and clusters. For example, ` + "`" + `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "game_server_config_overrides",
					Description: `The game_server_config_overrides contains the per game server config overrides. The overrides are processed in the order they are listed. As soon as a match is found for a cluster, the rest of the list is not processed. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `game_server_config_overrides` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "realms_selector",
					Description: `Selection by realms. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "config_version",
					Description: `Version of the configuration. The ` + "`" + `realms_selector` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "realms",
					Description: `List of realms to match against.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource id of the game server deployment eg: ` + "`" + `projects/my-project/locations/global/gameServerDeployments/my-deployment/rollout` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_iam_policy",
			Category:         "Data Sources",
			ShortDescription: `Generates an IAM policy that can be referenced by other resources, applying the policy to them.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding. Structure is documented below. The ` + "`" + `condition` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) Textual representation of an expression in Common Expression Language syntax.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) A title for the expression, i.e. a short string describing its purpose.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI. ## Attributes Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `The above bindings serialized in a format suitable for referencing from a resource that supports IAM.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_data",
					Description: `The above bindings serialized in a format suitable for referencing from a resource that supports IAM.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_iam_role",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google IAM Role.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `is a friendly title for the role, such as "Role Viewer"`,
				},
				resource.Attribute{
					Name:        "included_permissions",
					Description: `specifies the list of one or more permissions to include in the custom role, such as - ` + "`" + `iam.roles.get` + "`" + ``,
				},
				resource.Attribute{
					Name:        "stage",
					Description: `indicates the stage of a role in the launch lifecycle, such as ` + "`" + `GA` + "`" + `, ` + "`" + `BETA` + "`" + ` or ` + "`" + `ALPHA` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `is a friendly title for the role, such as "Role Viewer"`,
				},
				resource.Attribute{
					Name:        "included_permissions",
					Description: `specifies the list of one or more permissions to include in the custom role, such as - ` + "`" + `iam.roles.get` + "`" + ``,
				},
				resource.Attribute{
					Name:        "stage",
					Description: `indicates the stage of a role in the launch lifecycle, such as ` + "`" + `GA` + "`" + `, ` + "`" + `BETA` + "`" + ` or ` + "`" + `ALPHA` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_iam_testable_permissions",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of testable permissions for a resource. Testable permissions mean the permissions that user can add or remove in a role at a given resource. The resource can be referenced either via the full resource name or via a URI.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "full_resource_name",
					Description: `(Required) See [full resource name documentation](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more detail.`,
				},
				resource.Attribute{
					Name:        "stages",
					Description: `(Optional) The acceptable release stages of the permission in the output. Note that ` + "`" + `BETA` + "`" + ` does not include permissions in ` + "`" + `GA` + "`" + `, but you can specify both with ` + "`" + `["GA", "BETA"]` + "`" + ` for example. Can be a list of ` + "`" + `"ALPHA"` + "`" + `, ` + "`" + `"BETA"` + "`" + `, ` + "`" + `"GA"` + "`" + `, ` + "`" + `"DEPRECATED"` + "`" + `. Default is ` + "`" + `["GA"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_support_level",
					Description: `(Optional) The level of support for custom roles. Can be one of ` + "`" + `"NOT_SUPPORTED"` + "`" + `, ` + "`" + `"SUPPORTED"` + "`" + `, ` + "`" + `"TESTING"` + "`" + `. Default is ` + "`" + `"SUPPORTED"` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `A list of permissions matching the provided input. Structure is defined below. The ` + "`" + `permissions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the permission.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Human readable title of the permission.`,
				},
				resource.Attribute{
					Name:        "stage",
					Description: `Release stage of the permission.`,
				},
				resource.Attribute{
					Name:        "custom_support_level",
					Description: `The the support level of this permission for custom roles.`,
				},
				resource.Attribute{
					Name:        "api_disabled",
					Description: `Whether the corresponding API has been enabled for the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "permissions",
					Description: `A list of permissions matching the provided input. Structure is defined below. The ` + "`" + `permissions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the permission.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Human readable title of the permission.`,
				},
				resource.Attribute{
					Name:        "stage",
					Description: `Release stage of the permission.`,
				},
				resource.Attribute{
					Name:        "custom_support_level",
					Description: `The the support level of this permission for custom roles.`,
				},
				resource.Attribute{
					Name:        "api_disabled",
					Description: `Whether the corresponding API has been enabled for the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_iam_workload_identity_pool",
			Category:         "Data Sources",
			ShortDescription: `Get a IAM workload identity pool from Google Cloud`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "workload_identity_pool_id",
					Description: `(Required) The id of the pool which is the final component of the resource name. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_iam_workload_identity_pool](https://www.terraform.io/docs/providers/google/r/iam_workload_identity_pool.html) resource for details of all the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_iam_workload_identity_pool_provider",
			Category:         "Data Sources",
			ShortDescription: `Get a IAM workload identity pool provider from Google Cloud`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "workload_identity_pool_id",
					Description: `(Required) The id of the pool which is the final component of the pool resource name.`,
				},
				resource.Attribute{
					Name:        "workload_identity_pool_provider_id",
					Description: `(Required) The id of the provider which is the final component of the resource name. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_iam_workload_identity_pool_provider](https://www.terraform.io/docs/providers/google/r/iam_workload_identity_pool_provider.html) resource for details of all the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_iap_client",
			Category:         "Data Sources",
			ShortDescription: `Contains the data that describes an Identity Aware Proxy owned client.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "brand",
					Description: `(Required) The name of the brand.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) The client_id of the brand. ## Attributes Reference See [google_iap_client](https://www.terraform.io/docs/providers/google/r/iap_client.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_crypto_key",
			Category:         "Data Sources",
			ShortDescription: `Provides access to KMS key data with Google Cloud KMS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The CryptoKey's name. A CryptoKey’s name belonging to the specified Google Cloud Platform KeyRing and match the regular expression ` + "`" + `[a-zA-Z0-9_-]{1,63}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "key_ring",
					Description: `(Required) The ` + "`" + `self_link` + "`" + ` of the Google Cloud Platform KeyRing to which the key belongs. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rotation_period",
					Description: `Every time this period passes, generate a new CryptoKeyVersion and set it as the primary. The first rotation will take place after the specified period. The rotation period has the format of a decimal number with up to 9 fractional digits, followed by the letter s (seconds).`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `Defines the cryptographic capabilities of the key.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The self link of the created CryptoKey. Its format is ` + "`" + `projects/{projectId}/locations/{location}/keyRings/{keyRingName}/cryptoKeys/{cryptoKeyName}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rotation_period",
					Description: `Every time this period passes, generate a new CryptoKeyVersion and set it as the primary. The first rotation will take place after the specified period. The rotation period has the format of a decimal number with up to 9 fractional digits, followed by the letter s (seconds).`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `Defines the cryptographic capabilities of the key.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The self link of the created CryptoKey. Its format is ` + "`" + `projects/{projectId}/locations/{location}/keyRings/{keyRingName}/cryptoKeys/{cryptoKeyName}` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_crypto_key_version",
			Category:         "Data Sources",
			ShortDescription: `Provides access to KMS key version data with Google Cloud KMS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "crypto_key",
					Description: `(Required) The ` + "`" + `self_link` + "`" + ` of the Google Cloud Platform CryptoKey to which the key version belongs.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version number for this CryptoKeyVersion. Defaults to ` + "`" + `1` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `//cloudkms.googleapis.com/v1/{{crypto_key}}/cryptoKeyVersions/{{version}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the CryptoKeyVersion. See the [state reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions#CryptoKeyVersion.CryptoKeyVersionState) for possible outputs.`,
				},
				resource.Attribute{
					Name:        "protection_level",
					Description: `The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion. See the [protection_level reference](https://cloud.google.com/kms/docs/reference/rest/v1/ProtectionLevel) for possible outputs.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports. See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible outputs.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `If the enclosing CryptoKey has purpose ` + "`" + `ASYMMETRIC_SIGN` + "`" + ` or ` + "`" + `ASYMMETRIC_DECRYPT` + "`" + `, this block contains details about the public key associated to this CryptoKeyVersion. Structure is documented below. The ` + "`" + `public_key` + "`" + ` block, if present, contains:`,
				},
				resource.Attribute{
					Name:        "pem",
					Description: `The public key, encoded in PEM format. For more information, see the RFC 7468 sections for General Considerations and Textual Encoding of Subject Public Key Info.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `//cloudkms.googleapis.com/v1/{{crypto_key}}/cryptoKeyVersions/{{version}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the CryptoKeyVersion. See the [state reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions#CryptoKeyVersion.CryptoKeyVersionState) for possible outputs.`,
				},
				resource.Attribute{
					Name:        "protection_level",
					Description: `The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion. See the [protection_level reference](https://cloud.google.com/kms/docs/reference/rest/v1/ProtectionLevel) for possible outputs.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports. See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible outputs.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `If the enclosing CryptoKey has purpose ` + "`" + `ASYMMETRIC_SIGN` + "`" + ` or ` + "`" + `ASYMMETRIC_DECRYPT` + "`" + `, this block contains details about the public key associated to this CryptoKeyVersion. Structure is documented below. The ` + "`" + `public_key` + "`" + ` block, if present, contains:`,
				},
				resource.Attribute{
					Name:        "pem",
					Description: `The public key, encoded in PEM format. For more information, see the RFC 7468 sections for General Considerations and Textual Encoding of Subject Public Key Info.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_key_ring",
			Category:         "Data Sources",
			ShortDescription: `Provides access to KMS key ring data with Google Cloud KMS.`,
			Description:      ``,
			Icon:             "Security/Key_Management_Service.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The KeyRing's name. A KeyRing name must exist within the provided location and match the regular expression ` + "`" + `[a-zA-Z0-9_-]{1,63}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The Google Cloud Platform location for the KeyRing. A full list of valid locations can be found by running ` + "`" + `gcloud kms locations list` + "`" + `. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The self link of the created KeyRing. Its format is ` + "`" + `projects/{projectId}/locations/{location}/keyRings/{keyRingName}` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The self link of the created KeyRing. Its format is ` + "`" + `projects/{projectId}/locations/{location}/keyRings/{keyRingName}` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_secret",
			Category:         "Data Sources",
			ShortDescription: `Provides access to secret data encrypted with Google Cloud KMS`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plaintext",
					Description: `Contains the result of decrypting the provided ciphertext.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "plaintext",
					Description: `Contains the result of decrypting the provided ciphertext.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_secret_asymmetric",
			Category:         "Data Sources",
			ShortDescription: `Provides access to secret data encrypted with Google Cloud KMS asymmetric key`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plaintext",
					Description: `Contains the result of decrypting the provided ciphertext.`,
				},
				resource.Attribute{
					Name:        "crc32",
					Description: `Contains the crc32 checksum of the provided ciphertext.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "plaintext",
					Description: `Contains the result of decrypting the provided ciphertext.`,
				},
				resource.Attribute{
					Name:        "crc32",
					Description: `Contains the crc32 checksum of the provided ciphertext.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_secret_ciphertext",
			Category:         "Data Sources",
			ShortDescription: `Encrypts secret data with Google Cloud KMS and provides access to the ciphertext`,
			Description:      ``,
			Icon:             "Security/Key_Management_Service.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ciphertext",
					Description: `Contains the result of encrypting the provided plaintext, encoded in base64. ## User Project Overrides This data source supports [User Project Overrides](https://www.terraform.io/docs/providers/google/guides/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ciphertext",
					Description: `Contains the result of encrypting the provided plaintext, encoded in base64. ## User Project Overrides This data source supports [User Project Overrides](https://www.terraform.io/docs/providers/google/guides/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_monitoring_app_engine_service",
			Category:         "Data Sources",
			ShortDescription: `An Monitoring Service resource created automatically by GCP to monitor an App Engine service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "module_id",
					Description: `(Required) The ID of the App Engine module underlying this service. Corresponds to the moduleId resource label in the [gae_app](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name. - - - Other optional fields include:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The full REST resource name for this channel. The syntax is: ` + "`" + `projects/[PROJECT_ID]/services/[SERVICE_ID]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Name used for UI elements listing this (Monitoring) Service.`,
				},
				resource.Attribute{
					Name:        "telemetry",
					Description: `Configuration for how to query telemetry on the Service. Structure is documented below. The ` + "`" + `telemetry` + "`" + ` block includes:`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Optional) The full name of the resource that defines this service. Formatted as described in https://cloud.google.com/apis/design/resource_names.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The full REST resource name for this channel. The syntax is: ` + "`" + `projects/[PROJECT_ID]/services/[SERVICE_ID]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Name used for UI elements listing this (Monitoring) Service.`,
				},
				resource.Attribute{
					Name:        "telemetry",
					Description: `Configuration for how to query telemetry on the Service. Structure is documented below. The ` + "`" + `telemetry` + "`" + ` block includes:`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Optional) The full name of the resource that defines this service. Formatted as described in https://cloud.google.com/apis/design/resource_names.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_monitoring_cluster_istio_service",
			Category:         "Data Sources",
			ShortDescription: `An Monitoring Service resource created automatically by GCP to monitor an Cluster Istio service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location of the Kubernetes cluster in which this Istio service is defined. Corresponds to the location resource label in k8s_cluster resources.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Kubernetes cluster in which this Istio service is defined. Corresponds to the clusterName resource label in k8s_cluster resources.`,
				},
				resource.Attribute{
					Name:        "service_namespace",
					Description: `(Required) The namespace of the Istio service underlying this service. Corresponds to the destination_service_namespace metric label in Istio metrics.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The name of the Istio service underlying this service. Corresponds to the destination_service_name metric label in Istio metrics. - - - Other optional fields include:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The full REST resource name for this channel. The syntax is: ` + "`" + `projects/[PROJECT_ID]/services/[SERVICE_ID]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Name used for UI elements listing this (Monitoring) Service.`,
				},
				resource.Attribute{
					Name:        "telemetry",
					Description: `Configuration for how to query telemetry on the Service. Structure is documented below. The ` + "`" + `telemetry` + "`" + ` block includes:`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Optional) The full name of the resource that defines this service. Formatted as described in https://cloud.google.com/apis/design/resource_names.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The full REST resource name for this channel. The syntax is: ` + "`" + `projects/[PROJECT_ID]/services/[SERVICE_ID]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Name used for UI elements listing this (Monitoring) Service.`,
				},
				resource.Attribute{
					Name:        "telemetry",
					Description: `Configuration for how to query telemetry on the Service. Structure is documented below. The ` + "`" + `telemetry` + "`" + ` block includes:`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Optional) The full name of the resource that defines this service. Formatted as described in https://cloud.google.com/apis/design/resource_names.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_monitoring_mesh_istio_service",
			Category:         "Data Sources",
			ShortDescription: `An Monitoring Service resource created automatically by GCP to monitor an Mesh Istio service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mesh_uid",
					Description: `(Required) Identifier for the mesh in which this Istio service is defined. Corresponds to the meshUid metric label in Istio metrics.`,
				},
				resource.Attribute{
					Name:        "service_namespace",
					Description: `(Required) The namespace of the Istio service underlying this service. Corresponds to the destination_service_namespace metric label in Istio metrics.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The name of the Istio service underlying this service. Corresponds to the destination_service_name metric label in Istio metrics. - - - Other optional fields include:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The full REST resource name for this channel. The syntax is: ` + "`" + `projects/[PROJECT_ID]/services/[SERVICE_ID]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Name used for UI elements listing this (Monitoring) Service.`,
				},
				resource.Attribute{
					Name:        "telemetry",
					Description: `Configuration for how to query telemetry on the Service. Structure is documented below. The ` + "`" + `telemetry` + "`" + ` block includes:`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Optional) The full name of the resource that defines this service. Formatted as described in https://cloud.google.com/apis/design/resource_names.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The full REST resource name for this channel. The syntax is: ` + "`" + `projects/[PROJECT_ID]/services/[SERVICE_ID]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Name used for UI elements listing this (Monitoring) Service.`,
				},
				resource.Attribute{
					Name:        "telemetry",
					Description: `Configuration for how to query telemetry on the Service. Structure is documented below. The ` + "`" + `telemetry` + "`" + ` block includes:`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Optional) The full name of the resource that defines this service. Formatted as described in https://cloud.google.com/apis/design/resource_names.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_monitoring_notification_channel",
			Category:         "Data Sources",
			ShortDescription: `A NotificationChannel is a medium through which an alert is delivered when a policy violation is detected.`,
			Description:      ``,
			Icon:             "Management_Tools/Monitoring.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name for this notification channel.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the notification channel. - - - Other optional fields include:`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels (corresponding to the NotificationChannelDescriptor schema) to filter the notification channels by.`,
				},
				resource.Attribute{
					Name:        "user_labels",
					Description: `(Optional) User-provided key-value labels to filter by.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The full REST resource name for this channel. The syntax is: ` + "`" + `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_status",
					Description: `Indicates whether this channel has been verified or not.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Configuration fields that define the channel and its behavior.`,
				},
				resource.Attribute{
					Name:        "user_labels",
					Description: `User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional human-readable description of this notification channel.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether notifications are forwarded to the described channel.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The full REST resource name for this channel. The syntax is: ` + "`" + `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_status",
					Description: `Indicates whether this channel has been verified or not.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Configuration fields that define the channel and its behavior.`,
				},
				resource.Attribute{
					Name:        "user_labels",
					Description: `User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional human-readable description of this notification channel.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether notifications are forwarded to the described channel.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_monitoring_uptime_check_ips",
			Category:         "Data Sources",
			ShortDescription: `Returns the list of IP addresses Stackdriver Monitoring uses for uptime checking.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uptime_check_ips",
					Description: `A list of uptime check IPs used by Stackdriver Monitoring. Each ` + "`" + `uptime_check_ip` + "`" + ` contains:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `A broad region category in which the IP address is located.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `A more specific location within the region that typically encodes a particular city/town/metro (and its containing state/province or country) within the broader umbrella region category.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address from which the Uptime check originates. This is a fully specified IP address (not an IP address range). Most IP addresses, as of this publication, are in IPv4 format; however, one should not rely on the IP addresses being in IPv4 format indefinitely, and should support interpreting this field in either IPv4 or IPv6 format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_netblock_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get the IP addresses from different special IP ranges on Google Cloud Platform.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud-netblocks",
					Description: `Corresponds to the IP addresses used for resources on Google Cloud Platform. [More details.](https://cloud.google.com/compute/docs/faq#where_can_i_find_product_name_short_ip_ranges)`,
				},
				resource.Attribute{
					Name:        "google-netblocks",
					Description: `Corresponds to IP addresses used for Google services. [More details.](https://cloud.google.com/compute/docs/faq#where_can_i_find_product_name_short_ip_ranges)`,
				},
				resource.Attribute{
					Name:        "restricted-googleapis",
					Description: `Corresponds to the IP addresses used for Private Google Access only for services that support VPC Service Controls API access. [More details.](https://cloud.google.com/vpc/docs/private-access-options#domain-vips)`,
				},
				resource.Attribute{
					Name:        "private-googleapis",
					Description: `Corresponds to the IP addresses used for Private Google Access for services that do not support VPC Service Controls. [More details.](https://cloud.google.com/vpc/docs/private-access-options#domain-vips)`,
				},
				resource.Attribute{
					Name:        "dns-forwarders",
					Description: `Corresponds to the IP addresses used to originate Cloud DNS outbound forwarding. [More details.](https://cloud.google.com/dns/zones/#creating-forwarding-zones)`,
				},
				resource.Attribute{
					Name:        "iap-forwarders",
					Description: `Corresponds to the IP addresses used for Cloud IAP for TCP forwarding. [More details.](https://cloud.google.com/iap/docs/using-tcp-forwarding)`,
				},
				resource.Attribute{
					Name:        "health-checkers",
					Description: `Corresponds to the IP addresses used for health checking in Cloud Load Balancing. [More details.](https://cloud.google.com/load-balancing/docs/health-checks)`,
				},
				resource.Attribute{
					Name:        "legacy-health-checkers",
					Description: `Corresponds to the IP addresses used for legacy style health checkers (used by Network Load Balancing). [ More details.](https://cloud.google.com/load-balancing/docs/health-checks) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `Retrieve list of all CIDR blocks.`,
				},
				resource.Attribute{
					Name:        "cidr_blocks_ipv4",
					Description: `Retrieve list of the IPv4 CIDR blocks`,
				},
				resource.Attribute{
					Name:        "cidr_blocks_ipv6",
					Description: `Retrieve list of the IPv6 CIDR blocks, if available.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `Retrieve list of all CIDR blocks.`,
				},
				resource.Attribute{
					Name:        "cidr_blocks_ipv4",
					Description: `Retrieve list of the IPv4 CIDR blocks`,
				},
				resource.Attribute{
					Name:        "cidr_blocks_ipv6",
					Description: `Retrieve list of the IPv6 CIDR blocks, if available.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_organization",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Cloud Organization.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `The Organization ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the Organization in the form ` + "`" + `organizations/{organization_id}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory_customer_id",
					Description: `The Google for Work customer ID of the Organization.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `The Organization's current lifecycle state.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `The Organization ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the Organization in the form ` + "`" + `organizations/{organization_id}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory_customer_id",
					Description: `The Google for Work customer ID of the Organization.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `The Organization's current lifecycle state.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_project",
			Category:         "Data Sources",
			ShortDescription: `Retrieve project details`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project ID. If it is not provided, the provider project is used. ## Attributes Reference The following attributes are exported: See [google_project](https://www.terraform.io/docs/providers/google/r/google_project.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_project_organization_policy",
			Category:         "Data Sources",
			ShortDescription: `Retrieve Organization policies for a Google Project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The project ID.`,
				},
				resource.Attribute{
					Name:        "constraint",
					Description: `(Required) (Required) The name of the Constraint the Policy is configuring, for example, ` + "`" + `serviceuser.services` + "`" + `. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints). ## Attributes Reference See [google_project_organization_policy](https://www.terraform.io/docs/providers/google/r/google_project_organization_policy.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_projects",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a set of projects based on a filter.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) A string filter as defined in the [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#query-parameters). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "projects",
					Description: `A list of projects matching the provided filter. Structure is defined below. The ` + "`" + `projects` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project id of the project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "projects",
					Description: `A list of projects matching the provided filter. Structure is defined below. The ` + "`" + `projects` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project id of the project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_pubsub_topic",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Cloud Pub/Sub Topic.`,
			Description:      ``,
			Icon:             "Data_Analytics/Cloud_PubSub.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Cloud Pub/Sub Topic. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_pubsub_topic](https://www.terraform.io/docs/providers/google/r/pubsub_topic.html#argument-reference) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_redis_instance",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Cloud Redis instance.`,
			Description:      ``,
			Icon:             "Databases/Cloud_Memorystore.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of a Redis instance. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which the resource belongs. If it is not provided, the provider region is used. ## Attributes Reference See [google_redis_instance](https://www.terraform.io/docs/providers/google/r/redis_instance.html) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_runtimeconfig_config",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Cloud RuntimeConfig.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Runtime Configurator configuration. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_runtimeconfig_config](https://www.terraform.io/docs/providers/google/r/runtimeconfig_config.html#argument-reference) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_runtimeconfig_variable",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Cloud RuntimeConfig varialbe.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Runtime Configurator configuration.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Required) The name of the RuntimeConfig resource containing this variable. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_runtimeconfig_variable](https://www.terraform.io/docs/providers/google/r/runtimeconfig_variable.html#argument-reference) resource for details of the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_secret_manager_secret_version",
			Category:         "Data Sources",
			ShortDescription: `Get a Secret Manager secret's version.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project to get the secret version for. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Required) The secret to get the secret version for.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version of the secret to get. If it is not provided, the latest version is retrieved. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "secret_data",
					Description: `The secret data. No larger than 64KiB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the SecretVersion. Format: ` + "`" + `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time at which the Secret was created.`,
				},
				resource.Attribute{
					Name:        "destroy_time",
					Description: `The time at which the Secret was destroyed. Only present if state is DESTROYED.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `True if the current state of the SecretVersion is enabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_data",
					Description: `The secret data. No larger than 64KiB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the SecretVersion. Format: ` + "`" + `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time at which the Secret was created.`,
				},
				resource.Attribute{
					Name:        "destroy_time",
					Description: `The time at which the Secret was destroyed. Only present if state is DESTROYED.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `True if the current state of the SecretVersion is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_service_account",
			Category:         "Data Sources",
			ShortDescription: `Get the service account from a project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The Google service account ID. This be one of:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project that the service account is present in. Defaults to the provider project configuration. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address of the service account. This value should be referenced from any ` + "`" + `google_iam_policy` + "`" + ` data sources that would grant the service account privileges.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `The unique id of the service account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully-qualified name of the service account.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the service account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address of the service account. This value should be referenced from any ` + "`" + `google_iam_policy` + "`" + ` data sources that would grant the service account privileges.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `The unique id of the service account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully-qualified name of the service account.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the service account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_service_account_access_token",
			Category:         "Data Sources",
			ShortDescription: `Produces access_token for impersonated service accounts`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_token",
					Description: `The ` + "`" + `access_token` + "`" + ` representing the new generated identity.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_token",
					Description: `The ` + "`" + `access_token` + "`" + ` representing the new generated identity.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_service_account_id_token",
			Category:         "Data Sources",
			ShortDescription: `Produces OpenID Connect token for service accounts`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id_token",
					Description: `The ` + "`" + `id_token` + "`" + ` representing the new generated identity.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id_token",
					Description: `The ` + "`" + `id_token` + "`" + ` representing the new generated identity.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_service_account_key",
			Category:         "Data Sources",
			ShortDescription: `Get a Google Cloud Platform service account Public Key`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the service account key. This must have format ` + "`" + `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{KEYID}` + "`" + `, where ` + "`" + `{ACCOUNT}` + "`" + ` is the email address or unique id of the service account.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project that the service account will be created in. Defaults to the provider project configuration.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The public key, base64 encoded`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_key",
					Description: `The public key, base64 encoded`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_object_signed_url",
			Category:         "Data Sources",
			ShortDescription: `Provides signed URL to Google Cloud Storage object.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to read the object from`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The full path to the object inside the bucket`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `(Optional) What HTTP Method will the signed URL allow (defaults to ` + "`" + `GET` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) For how long shall the signed URL be valid (defaults to 1 hour - i.e. ` + "`" + `1h` + "`" + `). See [here](https://golang.org/pkg/time/#ParseDuration) for info on valid duration formats.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Optional) What Google service account credentials json should be used to sign the URL. This data source checks the following locations for credentials, in order of preference: data source ` + "`" + `credentials` + "`" + ` attribute, provider ` + "`" + `credentials` + "`" + ` attribute and finally the GOOGLE_APPLICATION_CREDENTIALS environment variable. >`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) If you specify this in the datasource, the client must provide the ` + "`" + `Content-Type` + "`" + ` HTTP header with the same value in its request.`,
				},
				resource.Attribute{
					Name:        "content_md5",
					Description: `(Optional) The [MD5 digest](https://cloud.google.com/storage/docs/hashes-etags#_MD5) value in Base64. Typically retrieved from ` + "`" + `google_storage_bucket_object.object.md5hash` + "`" + ` attribute. If you provide this in the datasource, the client (e.g. browser, curl) must provide the ` + "`" + `Content-MD5` + "`" + ` HTTP header with this same value in its request.`,
				},
				resource.Attribute{
					Name:        "extension_headers",
					Description: `(Optional) As needed. The server checks to make sure that the client provides matching values in requests using the signed URL. Any header starting with ` + "`" + `x-goog-` + "`" + ` is accepted but see the [Google Docs](https://cloud.google.com/storage/docs/xml-api/reference-headers) for list of headers that are supported by Google. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "signed_url",
					Description: `The signed URL that can be used to access the storage object without authentication.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "signed_url",
					Description: `The signed URL that can be used to access the storage object without authentication.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_spanner_instance",
			Category:         "Data Sources",
			ShortDescription: `Get a spanner instance from Google Cloud`,
			Description:      ``,
			Icon:             "Databases/Cloud_Spanner.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the spanner instance. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference See [google_spanner_instance](https://www.terraform.io/docs/providers/google/r/spanner_instance.html) resource for details of all the available attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_sql_backup_run",
			Category:         "Data Sources",
			ShortDescription: `Get a SQL backup run in Google Cloud SQL.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `(required) The name of the instance the backup is taken from.`,
				},
				resource.Attribute{
					Name:        "backup_id",
					Description: `(optional) The identifier for this backup run. Unique only for a specific Cloud SQL instance. If left empty and multiple backups exist for the instance, ` + "`" + `most_recent` + "`" + ` must be set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(optional) Toggles use of the most recent backup run if multiple backups exist for a Cloud SQL instance. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of the backups.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The time the backup operation actually started in UTC timezone in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this run. Refer to [API reference](https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/backupRuns#SqlBackupRunStatus) for possible status values.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `Location of the backups.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The time the backup operation actually started in UTC timezone in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this run. Refer to [API reference](https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/backupRuns#SqlBackupRunStatus) for possible status values.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_sql_ca_certs",
			Category:         "Data Sources",
			ShortDescription: `Get all of the trusted Certificate Authorities (CAs) for the specified SQL database instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name or self link of the instance. ---`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If ` + "`" + `project` + "`" + ` is not provided, the provider project is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "active_version",
					Description: `SHA1 fingerprint of the currently active CA certificate.`,
				},
				resource.Attribute{
					Name:        "certs",
					Description: `A list of server CA certificates for the instance. Each contains:`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `The CA certificate used to connect to the SQL instance via SSL.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `The CN valid for the CA cert.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the CA cert.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Expiration time of the CA cert.`,
				},
				resource.Attribute{
					Name:        "sha1_fingerprint",
					Description: `SHA1 fingerprint of the CA cert.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "active_version",
					Description: `SHA1 fingerprint of the currently active CA certificate.`,
				},
				resource.Attribute{
					Name:        "certs",
					Description: `A list of server CA certificates for the instance. Each contains:`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `The CA certificate used to connect to the SQL instance via SSL.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `The CN valid for the CA cert.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the CA cert.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Expiration time of the CA cert.`,
				},
				resource.Attribute{
					Name:        "sha1_fingerprint",
					Description: `SHA1 fingerprint of the CA cert.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_sql_database_instance",
			Category:         "Data Sources",
			ShortDescription: `Get a SQL database instance in Google Cloud SQL.`,
			Description:      ``,
			Icon:             "Databases/Cloud_SQL.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(required) The name of the instance.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(optional) The ID of the project in which the resource belongs. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `The settings to use for the database. The configuration is detailed below.`,
				},
				resource.Attribute{
					Name:        "database_version",
					Description: `The MySQL, PostgreSQL or SQL Server (beta) version to use.`,
				},
				resource.Attribute{
					Name:        "master_instance_name",
					Description: `The name of the instance that will act as the master in the replication setup.`,
				},
				resource.Attribute{
					Name:        "replica_configuration",
					Description: `The configuration for replication. The configuration is detailed below.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.`,
				},
				resource.Attribute{
					Name:        "encryption_key_name",
					Description: `[Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html)) The full path to the encryption key used for the CMEK disk encryption. The ` + "`" + `settings` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `The machine type to use.`,
				},
				resource.Attribute{
					Name:        "activation_policy",
					Description: `This specifies when the instance should be active. Can be either ` + "`" + `ALWAYS` + "`" + `, ` + "`" + `NEVER` + "`" + ` or ` + "`" + `ON_DEMAND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "authorized_gae_applications",
					Description: `(Deprecated) This property is only applicable to First Generation instances. First Generation instances are now deprecated, see [here](https://cloud.google.com/sql/docs/mysql/upgrade-2nd-gen) for information on how to upgrade to Second Generation instances. A list of Google App Engine (GAE) project names that are allowed to access this instance.`,
				},
				resource.Attribute{
					Name:        "availability_type",
					Description: `The availability type of the Cloud SQL instance, high availability (` + "`" + `REGIONAL` + "`" + `) or single zone (` + "`" + `ZONAL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "crash_safe_replication",
					Description: `(Deprecated) This property is only applicable to First Generation instances. First Generation instances are now deprecated, see [here](https://cloud.google.com/sql/docs/mysql/upgrade-2nd-gen)`,
				},
				resource.Attribute{
					Name:        "disk_autoresize",
					Description: `Configuration to increase storage size automatically.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The size of data disk, in GB.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The type of data disk.`,
				},
				resource.Attribute{
					Name:        "pricing_plan",
					Description: `Pricing plan for this instance.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `This property is only applicable to First Generation instances. First Generation instances are now deprecated, see [here](https://cloud.google.com/sql/docs/mysql/upgrade-2nd-gen)`,
				},
				resource.Attribute{
					Name:        "user_labels",
					Description: `A set of key/value user label pairs to assign to the instance. The ` + "`" + `settings.database_flags` + "`" + ` sublist contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the flag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the flag. The ` + "`" + `settings.backup_configuration` + "`" + ` subblock contains:`,
				},
				resource.Attribute{
					Name:        "binary_log_enabled",
					Description: `True if binary logging is enabled.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `True if backup configuration is enabled.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `` + "`" + `HH:MM` + "`" + ` format time indicating when backup configuration starts. The ` + "`" + `settings.ip_configuration` + "`" + ` subblock contains:`,
				},
				resource.Attribute{
					Name:        "ipv4_enabled",
					Description: `Whether this Cloud SQL instance should be assigned a public IPV4 address.`,
				},
				resource.Attribute{
					Name:        "private_network",
					Description: `The VPC network from which the Cloud SQL instance is accessible for private IP.`,
				},
				resource.Attribute{
					Name:        "require_ssl",
					Description: `True if mysqld default to ` + "`" + `REQUIRE X509` + "`" + ` for users connecting over IP. The ` + "`" + `settings.ip_configuration.authorized_networks[]` + "`" + ` sublist contains:`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `The [RFC 3339](https://tools.ietf.org/html/rfc3339) formatted date time string indicating when this whitelist expires.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A name for this whitelist entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `A CIDR notation IPv4 or IPv6 address that is allowed to access this instance. The ` + "`" + `settings.location_preference` + "`" + ` subblock contains:`,
				},
				resource.Attribute{
					Name:        "follow_gae_application",
					Description: `A GAE application whose zone to remain in.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The preferred compute engine. The ` + "`" + `settings.maintenance_window` + "`" + ` subblock for instances declares a one-hour [maintenance window](https://cloud.google.com/sql/docs/instance-settings?hl=en#maintenance-window-2ndgen) when an Instance can automatically restart to apply updates. The maintenance window is specified in UTC time. It contains:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `Day of week (` + "`" + `1-7` + "`" + `), starting on Monday.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Hour of day (` + "`" + `0-23` + "`" + `), ignored if ` + "`" + `day` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "update_track",
					Description: `Receive updates earlier (` + "`" + `canary` + "`" + `) or later (` + "`" + `stable` + "`" + `). The ` + "`" + `settings.insights_config` + "`" + ` subblock for instances declares [Query Insights](https://cloud.google.com/sql/docs/postgres/insights-overview) configuration. It contains:`,
				},
				resource.Attribute{
					Name:        "query_insights_enabled",
					Description: `True if Query Insights feature is enabled.`,
				},
				resource.Attribute{
					Name:        "query_string_length",
					Description: `Maximum query length stored in bytes. Between 256 and 4500. Default to 1024.`,
				},
				resource.Attribute{
					Name:        "record_application_tags",
					Description: `True if Query Insights will record application tags from query when enabled.`,
				},
				resource.Attribute{
					Name:        "record_client_address",
					Description: `True if Query Insights will record client address when enabled. The ` + "`" + `replica_configuration` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `PEM representation of the trusted CA's x509 certificate.`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `PEM representation of the replica's x509 certificate.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `PEM representation of the replica's private key.`,
				},
				resource.Attribute{
					Name:        "connect_retry_interval",
					Description: `The number of seconds between connect retries.`,
				},
				resource.Attribute{
					Name:        "dump_file_path",
					Description: `Path to a SQL file in GCS from which replica instances are created.`,
				},
				resource.Attribute{
					Name:        "failover_target",
					Description: `Specifies if the replica is the failover target.`,
				},
				resource.Attribute{
					Name:        "master_heartbeat_period",
					Description: `Time in ms between replication heartbeats.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the replication connection.`,
				},
				resource.Attribute{
					Name:        "sslCipher",
					Description: `Permissible ciphers for use in SSL encryption.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for replication connection.`,
				},
				resource.Attribute{
					Name:        "verify_server_certificate",
					Description: `True if the master's common name value is checked during the SSL handshake.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `The connection name of the instance to be used in connection strings.`,
				},
				resource.Attribute{
					Name:        "service_account_email_address",
					Description: `The service account email address assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.ip_address",
					Description: `The IPv4 address assigned.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.time_to_retire",
					Description: `The time this IP address will be retired, in RFC 3339 format.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.type",
					Description: `The type of this IP address.`,
				},
				resource.Attribute{
					Name:        "first_ip_address",
					Description: `The first IPv4 address of any type assigned.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `The first public (` + "`" + `PRIMARY` + "`" + `) IPv4 address assigned.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The first private (` + "`" + `PRIVATE` + "`" + `) IPv4 address assigned.`,
				},
				resource.Attribute{
					Name:        "settings.version",
					Description: `Used to make sure changes to the ` + "`" + `settings` + "`" + ` block are atomic.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.cert",
					Description: `The CA Certificate used to connect to the SQL Instance via SSL.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.common_name",
					Description: `The CN valid for the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.create_time",
					Description: `Creation time of the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.expiration_time",
					Description: `Expiration time of the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.sha1_fingerprint",
					Description: `SHA Fingerprint of the CA Cert.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "settings",
					Description: `The settings to use for the database. The configuration is detailed below.`,
				},
				resource.Attribute{
					Name:        "database_version",
					Description: `The MySQL, PostgreSQL or SQL Server (beta) version to use.`,
				},
				resource.Attribute{
					Name:        "master_instance_name",
					Description: `The name of the instance that will act as the master in the replication setup.`,
				},
				resource.Attribute{
					Name:        "replica_configuration",
					Description: `The configuration for replication. The configuration is detailed below.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.`,
				},
				resource.Attribute{
					Name:        "encryption_key_name",
					Description: `[Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html)) The full path to the encryption key used for the CMEK disk encryption. The ` + "`" + `settings` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `The machine type to use.`,
				},
				resource.Attribute{
					Name:        "activation_policy",
					Description: `This specifies when the instance should be active. Can be either ` + "`" + `ALWAYS` + "`" + `, ` + "`" + `NEVER` + "`" + ` or ` + "`" + `ON_DEMAND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "authorized_gae_applications",
					Description: `(Deprecated) This property is only applicable to First Generation instances. First Generation instances are now deprecated, see [here](https://cloud.google.com/sql/docs/mysql/upgrade-2nd-gen) for information on how to upgrade to Second Generation instances. A list of Google App Engine (GAE) project names that are allowed to access this instance.`,
				},
				resource.Attribute{
					Name:        "availability_type",
					Description: `The availability type of the Cloud SQL instance, high availability (` + "`" + `REGIONAL` + "`" + `) or single zone (` + "`" + `ZONAL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "crash_safe_replication",
					Description: `(Deprecated) This property is only applicable to First Generation instances. First Generation instances are now deprecated, see [here](https://cloud.google.com/sql/docs/mysql/upgrade-2nd-gen)`,
				},
				resource.Attribute{
					Name:        "disk_autoresize",
					Description: `Configuration to increase storage size automatically.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The size of data disk, in GB.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The type of data disk.`,
				},
				resource.Attribute{
					Name:        "pricing_plan",
					Description: `Pricing plan for this instance.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `This property is only applicable to First Generation instances. First Generation instances are now deprecated, see [here](https://cloud.google.com/sql/docs/mysql/upgrade-2nd-gen)`,
				},
				resource.Attribute{
					Name:        "user_labels",
					Description: `A set of key/value user label pairs to assign to the instance. The ` + "`" + `settings.database_flags` + "`" + ` sublist contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the flag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the flag. The ` + "`" + `settings.backup_configuration` + "`" + ` subblock contains:`,
				},
				resource.Attribute{
					Name:        "binary_log_enabled",
					Description: `True if binary logging is enabled.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `True if backup configuration is enabled.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `` + "`" + `HH:MM` + "`" + ` format time indicating when backup configuration starts. The ` + "`" + `settings.ip_configuration` + "`" + ` subblock contains:`,
				},
				resource.Attribute{
					Name:        "ipv4_enabled",
					Description: `Whether this Cloud SQL instance should be assigned a public IPV4 address.`,
				},
				resource.Attribute{
					Name:        "private_network",
					Description: `The VPC network from which the Cloud SQL instance is accessible for private IP.`,
				},
				resource.Attribute{
					Name:        "require_ssl",
					Description: `True if mysqld default to ` + "`" + `REQUIRE X509` + "`" + ` for users connecting over IP. The ` + "`" + `settings.ip_configuration.authorized_networks[]` + "`" + ` sublist contains:`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `The [RFC 3339](https://tools.ietf.org/html/rfc3339) formatted date time string indicating when this whitelist expires.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A name for this whitelist entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `A CIDR notation IPv4 or IPv6 address that is allowed to access this instance. The ` + "`" + `settings.location_preference` + "`" + ` subblock contains:`,
				},
				resource.Attribute{
					Name:        "follow_gae_application",
					Description: `A GAE application whose zone to remain in.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The preferred compute engine. The ` + "`" + `settings.maintenance_window` + "`" + ` subblock for instances declares a one-hour [maintenance window](https://cloud.google.com/sql/docs/instance-settings?hl=en#maintenance-window-2ndgen) when an Instance can automatically restart to apply updates. The maintenance window is specified in UTC time. It contains:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `Day of week (` + "`" + `1-7` + "`" + `), starting on Monday.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Hour of day (` + "`" + `0-23` + "`" + `), ignored if ` + "`" + `day` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "update_track",
					Description: `Receive updates earlier (` + "`" + `canary` + "`" + `) or later (` + "`" + `stable` + "`" + `). The ` + "`" + `settings.insights_config` + "`" + ` subblock for instances declares [Query Insights](https://cloud.google.com/sql/docs/postgres/insights-overview) configuration. It contains:`,
				},
				resource.Attribute{
					Name:        "query_insights_enabled",
					Description: `True if Query Insights feature is enabled.`,
				},
				resource.Attribute{
					Name:        "query_string_length",
					Description: `Maximum query length stored in bytes. Between 256 and 4500. Default to 1024.`,
				},
				resource.Attribute{
					Name:        "record_application_tags",
					Description: `True if Query Insights will record application tags from query when enabled.`,
				},
				resource.Attribute{
					Name:        "record_client_address",
					Description: `True if Query Insights will record client address when enabled. The ` + "`" + `replica_configuration` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `PEM representation of the trusted CA's x509 certificate.`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `PEM representation of the replica's x509 certificate.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `PEM representation of the replica's private key.`,
				},
				resource.Attribute{
					Name:        "connect_retry_interval",
					Description: `The number of seconds between connect retries.`,
				},
				resource.Attribute{
					Name:        "dump_file_path",
					Description: `Path to a SQL file in GCS from which replica instances are created.`,
				},
				resource.Attribute{
					Name:        "failover_target",
					Description: `Specifies if the replica is the failover target.`,
				},
				resource.Attribute{
					Name:        "master_heartbeat_period",
					Description: `Time in ms between replication heartbeats.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the replication connection.`,
				},
				resource.Attribute{
					Name:        "sslCipher",
					Description: `Permissible ciphers for use in SSL encryption.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for replication connection.`,
				},
				resource.Attribute{
					Name:        "verify_server_certificate",
					Description: `True if the master's common name value is checked during the SSL handshake.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `The connection name of the instance to be used in connection strings.`,
				},
				resource.Attribute{
					Name:        "service_account_email_address",
					Description: `The service account email address assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.ip_address",
					Description: `The IPv4 address assigned.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.time_to_retire",
					Description: `The time this IP address will be retired, in RFC 3339 format.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.type",
					Description: `The type of this IP address.`,
				},
				resource.Attribute{
					Name:        "first_ip_address",
					Description: `The first IPv4 address of any type assigned.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `The first public (` + "`" + `PRIMARY` + "`" + `) IPv4 address assigned.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The first private (` + "`" + `PRIVATE` + "`" + `) IPv4 address assigned.`,
				},
				resource.Attribute{
					Name:        "settings.version",
					Description: `Used to make sure changes to the ` + "`" + `settings` + "`" + ` block are atomic.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.cert",
					Description: `The CA Certificate used to connect to the SQL Instance via SSL.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.common_name",
					Description: `The CN valid for the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.create_time",
					Description: `Creation time of the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.expiration_time",
					Description: `Expiration time of the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.sha1_fingerprint",
					Description: `SHA Fingerprint of the CA Cert.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_bucket_object",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Google Cloud Storage bucket object.`,
			Description:      ``,
			Icon:             "Storage/Cloud_Storage.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the containing bucket.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the object. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Computed) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2) directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Computed) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Computed) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `(Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Computed) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".`,
				},
				resource.Attribute{
					Name:        "crc32c",
					Description: `(Computed) Base 64 CRC32 hash of the uploaded data.`,
				},
				resource.Attribute{
					Name:        "md5hash",
					Description: `(Computed) Base 64 MD5 hash of the uploaded data.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `(Computed) A url reference to this object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Computed) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object. Supported values include: ` + "`" + `MULTI_REGIONAL` + "`" + `, ` + "`" + `REGIONAL` + "`" + `, ` + "`" + `NEARLINE` + "`" + `, ` + "`" + `COLDLINE` + "`" + `, ` + "`" + `ARCHIVE` + "`" + `. If not provided, this defaults to the bucket's default storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.`,
				},
				resource.Attribute{
					Name:        "media_link",
					Description: `(Computed) A url reference to download this object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Computed) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2) directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Computed) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Computed) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `(Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Computed) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".`,
				},
				resource.Attribute{
					Name:        "crc32c",
					Description: `(Computed) Base 64 CRC32 hash of the uploaded data.`,
				},
				resource.Attribute{
					Name:        "md5hash",
					Description: `(Computed) Base 64 MD5 hash of the uploaded data.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `(Computed) A url reference to this object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Computed) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object. Supported values include: ` + "`" + `MULTI_REGIONAL` + "`" + `, ` + "`" + `REGIONAL` + "`" + `, ` + "`" + `NEARLINE` + "`" + `, ` + "`" + `COLDLINE` + "`" + `, ` + "`" + `ARCHIVE` + "`" + `. If not provided, this defaults to the bucket's default storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.`,
				},
				resource.Attribute{
					Name:        "media_link",
					Description: `(Computed) A url reference to download this object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_bucket_object_content",
			Category:         "Data Sources",
			ShortDescription: `Get content of a Google Cloud Storage bucket object.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the containing bucket.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the object. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object content.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `(Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object content.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_project_service_account",
			Category:         "Data Sources",
			ShortDescription: `Get the email address of the project's Google Cloud Storage service account`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project the unique service account was created for. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "user_project",
					Description: `(Optional) The project the lookup originates from. This field is used if you are making the request from a different account than the one you are finding the service account for. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `The email address of the service account. This value is often used to refer to the service account in order to grant IAM permissions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email_address",
					Description: `The email address of the service account. This value is often used to refer to the service account in order to grant IAM permissions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_transfer_project_service_account",
			Category:         "Data Sources",
			ShortDescription: `Retrieve default service account used by Storage Transfer Jobs running in this project`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project ID. If it is not provided, the provider project is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email address of the default service account used by Storage Transfer Jobs running in this project`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `Email address of the default service account used by Storage Transfer Jobs running in this project`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_tpu_tensorflow_versions",
			Category:         "Data Sources",
			ShortDescription: `Get available TensorFlow versions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project to list versions for. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone to list versions for. If it is not provided, the provider zone is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `The list of TensorFlow versions available for the given project and zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "versions",
					Description: `The list of TensorFlow versions available for the given project and zone.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"google_active_folder":                                0,
		"google_app_engine_default_service_account":           1,
		"google_bigquery_default_service_account":             2,
		"google_billing_account":                              3,
		"google_client_config":                                4,
		"google_client_openid_userinfo":                       5,
		"google_cloud_identity_group_memberships":             6,
		"google_cloud_identity_groups":                        7,
		"google_cloud_run_locations":                          8,
		"google_cloud_run_service":                            9,
		"google_cloudfunctions_function":                      10,
		"google_composer_environment":                         11,
		"google_composer_image_versions":                      12,
		"google_compute_address":                              13,
		"google_compute_backend_bucket":                       14,
		"google_compute_backend_service":                      15,
		"google_compute_default_service_account":              16,
		"google_compute_forwarding_rule":                      17,
		"google_compute_global_address":                       18,
		"google_compute_global_forwarding_rule":               19,
		"google_compute_health_check":                         20,
		"google_compute_image":                                21,
		"google_compute_instance":                             22,
		"google_compute_instance_group":                       23,
		"google_compute_instance_serial_port":                 24,
		"google_compute_instance_template":                    25,
		"google_compute_lb_ip_ranges":                         26,
		"google_compute_network":                              27,
		"google_compute_network_endpoint_group":               28,
		"google_compute_node_types":                           29,
		"google_compute_region_instance_group":                30,
		"google_compute_region_ssl_certificate":               31,
		"google_compute_regions":                              32,
		"google_compute_resource_policy":                      33,
		"google_compute_router":                               34,
		"google_compute_ssl_certificate":                      35,
		"google_compute_ssl_policy":                           36,
		"google_compute_subnetwork":                           37,
		"google_compute_vpn_gateway":                          38,
		"google_compute_zones":                                39,
		"google_container_cluster":                            40,
		"google_container_engine_versions":                    41,
		"google_container_registry_image":                     42,
		"google_container_registry_repository":                43,
		"google_dns_keys":                                     44,
		"google_dns_managed_zone":                             45,
		"google_firebase_web_app":                             46,
		"google_firebase_web_app_config":                      47,
		"google_folder":                                       48,
		"google_folder_organization_policy":                   49,
		"google_game_services_game_server_deployment_rollout": 50,
		"google_iam_policy":                                   51,
		"google_iam_role":                                     52,
		"google_iam_testable_permissions":                     53,
		"google_iam_workload_identity_pool":                   54,
		"google_iam_workload_identity_pool_provider":          55,
		"google_iap_client":                                   56,
		"google_kms_crypto_key":                               57,
		"google_kms_crypto_key_version":                       58,
		"google_kms_key_ring":                                 59,
		"google_kms_secret":                                   60,
		"google_kms_secret_asymmetric":                        61,
		"google_kms_secret_ciphertext":                        62,
		"google_monitoring_app_engine_service":                63,
		"google_monitoring_cluster_istio_service":             64,
		"google_monitoring_mesh_istio_service":                65,
		"google_monitoring_notification_channel":              66,
		"google_monitoring_uptime_check_ips":                  67,
		"google_netblock_ip_ranges":                           68,
		"google_organization":                                 69,
		"google_project":                                      70,
		"google_project_organization_policy":                  71,
		"google_projects":                                     72,
		"google_pubsub_topic":                                 73,
		"google_redis_instance":                               74,
		"google_runtimeconfig_config":                         75,
		"google_runtimeconfig_variable":                       76,
		"google_secret_manager_secret_version":                77,
		"google_service_account":                              78,
		"google_service_account_access_token":                 79,
		"google_service_account_id_token":                     80,
		"google_service_account_key":                          81,
		"google_storage_object_signed_url":                    82,
		"google_spanner_instance":                             83,
		"google_sql_backup_run":                               84,
		"google_sql_ca_certs":                                 85,
		"google_sql_database_instance":                        86,
		"google_storage_bucket_object":                        87,
		"google_storage_bucket_object_content":                88,
		"google_storage_project_service_account":              89,
		"google_storage_transfer_project_service_account":     90,
		"google_tpu_tensorflow_versions":                      91,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
