package artifactory

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_access_token",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require ` + "`" + `groups` + "`" + ` to be set.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the token is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Valid time units are "s", "m", "h".`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of groups. The token is granted access based on the permissions of the groups. Specify ` + "`" + `["`,
				},
				resource.Attribute{
					Name:        "admin_token",
					Description: `(Optional) Specify the ` + "`" + `instance_id` + "`" + ` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. ` + "`" + `admin_token` + "`" + ` cannot be specified with ` + "`" + `groups` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "refreshable",
					Description: `(Optional) Is this token refreshable? Defaults to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `(Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set ` + "`" + `"jfrt@`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `Returns the access token to authenciate to Artifactory`,
				},
				resource.Attribute{
					Name:        "refresh_token",
					Description: `Returns the refresh token when ` + "`" + `refreshable` + "`" + ` is true, or an empty string when ` + "`" + `refreshable` + "`" + ` is false ## References - https://www.jfrog.com/confluence/display/ACC1X/Access+Tokens - https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateToken ## Import Artifactory`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_api_key",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_key",
					Description: `The API key. ## Import A user's API key can be imported using any identifier, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_api_key.test import ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_certificate",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alias",
					Description: `(Required) Name of certificate.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) PEM-encoded client certificate and private key. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `SHA256 fingerprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "issued_by",
					Description: `Name of the certificate authority that issued the certificate.`,
				},
				resource.Attribute{
					Name:        "issued_on",
					Description: `The time & date when the certificate is valid from.`,
				},
				resource.Attribute{
					Name:        "issued_to",
					Description: `Name of whom the certificate has been issued to.`,
				},
				resource.Attribute{
					Name:        "valid_until",
					Description: `The time & date when the certificate expires. ## Import Certificates can be imported using their alias, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_certificate.my-cert my-cert ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fingerprint",
					Description: `SHA256 fingerprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "issued_by",
					Description: `Name of the certificate authority that issued the certificate.`,
				},
				resource.Attribute{
					Name:        "issued_on",
					Description: `The time & date when the certificate is valid from.`,
				},
				resource.Attribute{
					Name:        "issued_to",
					Description: `Name of whom the certificate has been issued to.`,
				},
				resource.Attribute{
					Name:        "valid_until",
					Description: `The time & date when the certificate expires. ## Import Certificates can be imported using their alias, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_certificate.my-cert my-cert ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_local_repository",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "package_type",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "repo_layout_ref",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "debian_trivial_layout",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "max_unique_tags",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "blacked_out",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "property_sets",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "archive_browsing_enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "calculate_yum_metadata",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "yum_root_depth",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "docker_api_version",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "enable_file_lists_indexing",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `(Optional, Nuget repos only) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_repository.my-local my-local ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_permission_target",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of permission`,
				},
				resource.Attribute{
					Name:        "repo",
					Description: `(Optional) Repository permission configuration`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `(Optional) Pattern of artifacts to include`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `(Optional) Pattern of artifacts to exclude`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) List of repositories this permission target is applicable for`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: ``,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) Users this permission target applies for.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Groups this permission applies for.`,
				},
				resource.Attribute{
					Name:        "build",
					Description: `(Optional) As for repo but for artifactory-build-info permssions. ## Import Permission targets can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_permission_target.terraform-test-permission mypermission ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_permission_target_v1",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of permission`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `(Optional) Pattern of artifacts to include`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `(Optional) Pattern of artifacts to exclude`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) List of repositories this permission target is applicable for`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) Users this permission target applies for.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Groups this permission applies for. The permissions can be set to a combination of m=admin; d=delete; w=deploy; n=annotate; r=read ## Import Permission targets can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_permission_targets.terraform-test-permission mypermission ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_remote_repository",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "package_type",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "repo_layout_ref",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Requires password encryption to be turned off ` + "`" + `POST /api/system/decrypt` + "`" + ``,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "hard_fail",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "offline",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "blacked_out",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "store_artifacts_locally",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "socket_timeout_millis",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "local_address",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "missed_cache_period_seconds",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "unused_artifacts_cleanup_period_hours",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "share_configuration",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "synchronize_properties",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "block_mismatching_mime_types",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "property_sets",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "allow_any_host_auth",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "enable_cookie_management",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_tls_certificate",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "pypi_registry_url",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bypass_head_requests",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "enable_token_authentication",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "feed_context_path",
					Description: `(Optional, Nuget repos only)`,
				},
				resource.Attribute{
					Name:        "download_context_path",
					Description: `(Optional, Nuget repos only)`,
				},
				resource.Attribute{
					Name:        "v3_feed_url",
					Description: `(Optional, Nuget repos only)`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `(Optional, Nuget repos only)`,
				},
				resource.Attribute{
					Name:        "nuget",
					Description: `(Optional) Deprecated since 6.9.0+ Nuget repository special configuration`,
				},
				resource.Attribute{
					Name:        "feed_context_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "download_context_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "v3_feed_url",
					Description: `(Optional) ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_repository.my-remote my-remote ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_replication_config",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repo_key",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "cron_exp",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "enable_event_replication",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "replications",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "socket_timeout_millis",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Requires password encryption to be turned off ` + "`" + `POST /api/system/decrypt` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_deletes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_properties",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_statistics",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `(Optional) ## Import Replication configs can be imported using their repo key, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_replication_config.foo-rep provider_test_source ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_single_replication_config",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repo_key",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "cron_exp",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "enable_event_replication",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "socket_timeout_millis",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Requires password encryption to be turned off ` + "`" + `POST /api/system/decrypt` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_deletes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_properties",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_statistics",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `(Optional) ## Import Replication configs can be imported using their repo key, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_single_replication_config.foo-rep repository-key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Username for user`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email for user`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for the user`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "profile_updatable",
					Description: `(Optional) When set, this user can update his profile details (except for the password. Only an administrator can update the password).`,
				},
				resource.Attribute{
					Name:        "disable_ui_access",
					Description: `(Optional) When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges.`,
				},
				resource.Attribute{
					Name:        "internal_password_disabled",
					Description: `(Optional) When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of groups this user is a part of ## Import Users can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_user.test-user myusername ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_virtual_repository",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "package_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "repo_layout_ref",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "debian_trivial_layout",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "artifactory_requests_can_retrieve_remote_artifacts",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "default_deployment_repo",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `(Optional, Nuget repos only) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_repository.foo foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_xray_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the policy (must be unique)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the policy`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) More verbose description of the policy`,
				},
				resource.Attribute{
					Name:        "author",
					Description: `(Optional) Name of the policy author`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Required) Nested block describing the policy rules. Described below. ### Rules The top-level ` + "`" + `rules` + "`" + ` block is a list of one or more rules that each supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the rule`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Integer describing the rule priority`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Required) Nested block describing the criteria for the policy. Described below.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `(Required) Nested block describing the actions to be applied by the policy. Described below. #### criteria ~>`,
				},
				resource.Attribute{
					Name:        "min_severity",
					Description: `(Optional) The minimum security vulnerability severity that will be impacted by the policy.`,
				},
				resource.Attribute{
					Name:        "cvss_range",
					Description: `(Optional) Nested block describing a CVS score range to be impacted. Defined below. ###### cvss_range The nested ` + "`" + `cvss_range` + "`" + ` block is a list of one object that contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Required) The beginning of the range of CVS scores (from 1-10) to flag.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Required) The end of the range of CVS scores (from 1-10) to flag. ##### License criteria`,
				},
				resource.Attribute{
					Name:        "allow_unknown",
					Description: `(Optional) Whether or not to allow components whose license cannot be determined (` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "banned_licenses",
					Description: `(Optional) A list of OSS license names that may not be attached to a component.`,
				},
				resource.Attribute{
					Name:        "allowed_licenses",
					Description: `(Optional) A list of OSS license names that may be attached to a component. #### actions ~>`,
				},
				resource.Attribute{
					Name:        "mails",
					Description: `(Optional) A list of email addressed that will get emailed when a violation is triggered.`,
				},
				resource.Attribute{
					Name:        "fail_build",
					Description: `(Optional) Whether or not the related CI build should be marked as failed if a violation is triggered. This option is only available when the policy is applied to an ` + "`" + `xray_watch` + "`" + ` resource with a ` + "`" + `type` + "`" + ` of ` + "`" + `builds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "block_download",
					Description: `(Optional) Nested block describing artifacts that should be blocked for download if a violation is triggered. Described below.`,
				},
				resource.Attribute{
					Name:        "webhooks",
					Description: `(Optional) A list of Xray-configured webhook URLs to be invoked if a violation is triggered.`,
				},
				resource.Attribute{
					Name:        "custom_severity",
					Description: `(Optional) The severity of violation to be triggered if the ` + "`" + `criteria` + "`" + ` are met. ###### block_download ~>`,
				},
				resource.Attribute{
					Name:        "unscanned",
					Description: `Whether or not to block download of artifacts that meet the artifact ` + "`" + `filters` + "`" + ` for the associated ` + "`" + `xray_watch` + "`" + ` resource but have not been scanned yet.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `Whether or not to block download of artifacts that meet the artifact and severity ` + "`" + `filters` + "`" + ` for the associated ` + "`" + `xray_watch` + "`" + ` resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Timestamp of when the policy was first created`,
				},
				resource.Attribute{
					Name:        "modified",
					Description: `Timestamp of when the policy was last modified ## Import A policy can be imported by using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import xray_policy.example policy-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_xray_watch",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the watch (must be unique)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the watch`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Whether or not the watch will be active`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Nested argument describing the resources to be watched. Defined below.`,
				},
				resource.Attribute{
					Name:        "assigned_policies",
					Description: `(Required) Nested argument describing policies that will be applied. Defined below. ### resources The top-level ` + "`" + `resources` + "`" + ` block contains a list of one or more resource objects that each support the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of resource to be watched`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name describing the resource`,
				},
				resource.Attribute{
					Name:        "bin_mgr_id",
					Description: `(Optional) The ID number of a binary manager resource`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `(Optional) Nested argument describing filters to be applied. Defined below. #### filters The nested ` + "`" + `filters` + "`" + ` block contains a list of one or more filters to be applied, each of which supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of filter, such as ` + "`" + `regex` + "`" + ` or ` + "`" + `package-type` + "`" + ``,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the filter, such as the text of the regex or name of the package type ### assigned_policies The top-level ` + "`" + `assigned_policies` + "`" + ` block contains a list of one or more policy objects that each support the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy that will be applied`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the policy ## Import Watches can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import xray_watch.example watch-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"artifactory_artifactory_access_token":              0,
		"artifactory_artifactory_api_key":                   1,
		"artifactory_artifactory_certificate":               2,
		"artifactory_artifactory_group":                     3,
		"artifactory_artifactory_local_repository":          4,
		"artifactory_artifactory_permission_target":         5,
		"artifactory_artifactory_permission_target_v1":      6,
		"artifactory_artifactory_remote_repository":         7,
		"artifactory_artifactory_replication_config":        8,
		"artifactory_artifactory_single_replication_config": 9,
		"artifactory_artifactory_user":                      10,
		"artifactory_artifactory_virtual_repository":        11,
		"artifactory_artifactory_xray_policy":               12,
		"artifactory_artifactory_xray_watch":                13,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
