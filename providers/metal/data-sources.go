package metal

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "metal_device",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal device datasource. This can be used to read existing devices.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `The device name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The id of the project in which the devices exists`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `Device ID User can lookup devices either by ` + "`" + `device_id` + "`" + ` or ` + "`" + `project_id` + "`" + ` and ` + "`" + `hostname` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_private_ipv4",
					Description: `The ipv4 private IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv4",
					Description: `The ipv4 management IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv6",
					Description: `The ipv6 management IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle of the device (monthly or hourly)`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility where the device is deployed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device`,
				},
				resource.Attribute{
					Name:        "hardware_reservation_id",
					Description: `The id of hardware reservation which this device occupies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the device`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The device's private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 networks:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `IPv4 or IPv6 address string`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Bit length of the network mask of the address`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Address of router`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether the address is routable from the Internet`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP version - "4" or "6"`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `L2 network type of the device, one of "layer3", "layer2-bonded", "layer2-individual", "hybrid"`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `The operating system running on the device`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The hardware config of the device`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `Ports assigned to the device`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port (e.g. ` + "`" + `eth0` + "`" + `, or ` + "`" + `bond0` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the port`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the port (e.g. ` + "`" + `NetworkPort` + "`" + ` or ` + "`" + `NetworkBondPort` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address assigned to the port`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Whether this port is part of a bond in bonded network setup`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Root password to the server (if still available)`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `List of IDs of SSH keys deployed in the device, can be both user or project SSH keys`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the device`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_private_ipv4",
					Description: `The ipv4 private IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv4",
					Description: `The ipv4 management IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv6",
					Description: `The ipv6 management IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle of the device (monthly or hourly)`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility where the device is deployed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device`,
				},
				resource.Attribute{
					Name:        "hardware_reservation_id",
					Description: `The id of hardware reservation which this device occupies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the device`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The device's private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 networks:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `IPv4 or IPv6 address string`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Bit length of the network mask of the address`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Address of router`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether the address is routable from the Internet`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP version - "4" or "6"`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `L2 network type of the device, one of "layer3", "layer2-bonded", "layer2-individual", "hybrid"`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `The operating system running on the device`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The hardware config of the device`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `Ports assigned to the device`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port (e.g. ` + "`" + `eth0` + "`" + `, or ` + "`" + `bond0` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the port`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the port (e.g. ` + "`" + `NetworkPort` + "`" + ` or ` + "`" + `NetworkBondPort` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address assigned to the port`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Whether this port is part of a bond in bonded network setup`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Root password to the server (if still available)`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `List of IDs of SSH keys deployed in the device, can be both user or project SSH keys`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the device`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_device_bgp_neighbors",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource for listing BGP neighbors of an Equinix Metal device`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `UUID of BGP-enabled device whose neighbors to list ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bgp_neighbors",
					Description: `array of BGP neighbor records with attributes:`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `IP address version, 4 or 6`,
				},
				resource.Attribute{
					Name:        "customer_as",
					Description: `Local autonomous system number`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `Local used peer IP address`,
				},
				resource.Attribute{
					Name:        "md5_enabled",
					Description: `Whether BGP session is password enabled`,
				},
				resource.Attribute{
					Name:        "md5_password",
					Description: `BGP session password in plaintext (not a checksum)`,
				},
				resource.Attribute{
					Name:        "multihop",
					Description: `Whether the neighbor is in EBGP multihop session`,
				},
				resource.Attribute{
					Name:        "peer_as",
					Description: `Peer AS number (different than customer_as for EBGP)`,
				},
				resource.Attribute{
					Name:        "peer_ips",
					Description: `Array of IP addresses of this neighbor's peers`,
				},
				resource.Attribute{
					Name:        "routes_in",
					Description: `Array of incoming routes. Each route has attributes:`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `CIDR expression of route (IP/mask)`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(bool) Whether the route is exact`,
				},
				resource.Attribute{
					Name:        "routes_out",
					Description: `Array of outgoing routes in the same format`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp_neighbors",
					Description: `array of BGP neighbor records with attributes:`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `IP address version, 4 or 6`,
				},
				resource.Attribute{
					Name:        "customer_as",
					Description: `Local autonomous system number`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `Local used peer IP address`,
				},
				resource.Attribute{
					Name:        "md5_enabled",
					Description: `Whether BGP session is password enabled`,
				},
				resource.Attribute{
					Name:        "md5_password",
					Description: `BGP session password in plaintext (not a checksum)`,
				},
				resource.Attribute{
					Name:        "multihop",
					Description: `Whether the neighbor is in EBGP multihop session`,
				},
				resource.Attribute{
					Name:        "peer_as",
					Description: `Peer AS number (different than customer_as for EBGP)`,
				},
				resource.Attribute{
					Name:        "peer_ips",
					Description: `Array of IP addresses of this neighbor's peers`,
				},
				resource.Attribute{
					Name:        "routes_in",
					Description: `Array of incoming routes. Each route has attributes:`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `CIDR expression of route (IP/mask)`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(bool) Whether the route is exact`,
				},
				resource.Attribute{
					Name:        "routes_out",
					Description: `Array of outgoing routes in the same format`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_facility",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal facility datasource. This can be used to read facilities.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "code",
					Description: `The facility code Facilities can be looked up by ` + "`" + `code` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the facility`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the facilityg system running on the device`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `The features of the facility`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the facility`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the facilityg system running on the device`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `The features of the facility`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_ip_block_ranges",
			Category:         "Data Sources",
			ShortDescription: `List IP address ranges allocated to a project`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of the project from which to list the blocks.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) Facility code filtering the IP blocks. Global IPv4 blcoks will be listed anyway. If you omit this, all the block from the project will be listed. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "global_ipv4",
					Description: `list of CIDR expressions for Global IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "public_ipv4",
					Description: `list of CIDR expressions for Public IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "private_ipv4",
					Description: `list of CIDR expressions for Private IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `list of CIDR expressions for IPv6 blocks in the project`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "global_ipv4",
					Description: `list of CIDR expressions for Global IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "public_ipv4",
					Description: `list of CIDR expressions for Public IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "private_ipv4",
					Description: `list of CIDR expressions for Private IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `list of CIDR expressions for IPv6 blocks in the project`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_operating_system",
			Category:         "Data Sources",
			ShortDescription: `Get an Equinix Metal operating system image`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "distro",
					Description: `(Optional) Name of the OS distribution.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name or part of the name of the distribution. Case insensitive.`,
				},
				resource.Attribute{
					Name:        "provisionable_on",
					Description: `(Optional) Plan name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the distribution ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Operating system slug`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Operating system slug (same as ` + "`" + `id` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Operating system slug`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Operating system slug (same as ` + "`" + `id` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_organization",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal Organization datasource. This can be used to read existing Organizations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The organization name`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The UUID of the organization resource Exactly one of ` + "`" + `name` + "`" + ` or ` + "`" + `organization_id` + "`" + ` must be given. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_ids",
					Description: `UUIDs of project resources which belong to this organization`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `Website link`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `Twitter handle`,
				},
				resource.Attribute{
					Name:        "logo",
					Description: `Logo URL`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_ids",
					Description: `UUIDs of project resources which belong to this organization`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `Website link`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `Twitter handle`,
				},
				resource.Attribute{
					Name:        "logo",
					Description: `Logo URL`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_precreated_ip_block",
			Category:         "Data Sources",
			ShortDescription: `Load automatically created IP blocks from your Equinix Metal project`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of the project where the searched block should be.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Required) 4 or 6, depending on which block you are looking for.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(Required) Whether to look for public or private block.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional) Whether to look for global block. Default is false for backward compatibility.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) Facility of the searched block. (Optional) Only allowed for non-global blocks. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `CIDR notation of the looked up block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `CIDR notation of the looked up block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_project",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal Project datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name which is used to look up the project`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The UUID by which to look up the project ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `The UUID of payment method for this project`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The UUID of this project's parent organization`,
				},
				resource.Attribute{
					Name:        "backend_transfer",
					Description: `Whether Backend Transfer is enabled for this project`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the project was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the project was updated`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `List of UUIDs of user accounts which belong to this project`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/). The ` + "`" + `bgp_config` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `Autonomous System Number for local BGP deployment`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `Password for BGP session in plaintext (not a checksum)`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `private` + "`" + ` or ` + "`" + `public` + "`" + `, the ` + "`" + `private` + "`" + ` is likely to be usable immediately, the ` + "`" + `public` + "`" + ` will need to be review by Equinix Metal engineers`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of BGP configuration in the project`,
				},
				resource.Attribute{
					Name:        "max_prefix",
					Description: `The maximum number of route filters allowed per server`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `The UUID of payment method for this project`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The UUID of this project's parent organization`,
				},
				resource.Attribute{
					Name:        "backend_transfer",
					Description: `Whether Backend Transfer is enabled for this project`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the project was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the project was updated`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `List of UUIDs of user accounts which belong to this project`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/). The ` + "`" + `bgp_config` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `Autonomous System Number for local BGP deployment`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `Password for BGP session in plaintext (not a checksum)`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `private` + "`" + ` or ` + "`" + `public` + "`" + `, the ` + "`" + `private` + "`" + ` is likely to be usable immediately, the ` + "`" + `public` + "`" + ` will need to be review by Equinix Metal engineers`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of BGP configuration in the project`,
				},
				resource.Attribute{
					Name:        "max_prefix",
					Description: `The maximum number of route filters allowed per server`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_project_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal Project SSH Key datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "search",
					Description: `(Optional) The name, fingerprint, or public_key of the SSH Key to search for in the Equinix Metal project`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the SSH Key to search for in the Equinix Metal project`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The Equinix Metal project id of the Equinix Metal SSH Key One of either ` + "`" + `search` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided along with ` + "`" + `project_id` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of parent project`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of parent project (same as project_id)`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of parent project`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of parent project (same as project_id)`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_spot_market_price",
			Category:         "Data Sources",
			ShortDescription: `Get an Equinix Metal Spot Market Price`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "facility",
					Description: `(Required) Name of the facility.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Name of the plan. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Current spot market price for given plan in given facility.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "price",
					Description: `Current spot market price for given plan in given facility.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_spot_market_request",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource for existing Spot Market Requests in the Equinix Metal host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "request_id",
					Description: `(Required) The id of the Spot Market Request ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "device_ids",
					Description: `List of IDs of devices spawned by the referenced Spot Market Request`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "device_ids",
					Description: `List of IDs of devices spawned by the referenced Spot Market Request`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_volume",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal Block Storage Volume Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of volume for lookup`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID the parent Equinix Metal project (for lookup by name) Either ` + "`" + `volume_id` + "`" + ` or both ` + "`" + `project_id` + "`" + ` and ` + "`" + `name` + "`" + ` must be specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the volume`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project id the volume is in`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size in GB of the volume`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Performance plan the volume is on`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle, defaults to hourly`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility slug the volume resides in`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the volume is locked or not`,
				},
				resource.Attribute{
					Name:        "device_ids",
					Description: `UUIDs of devices to which this volume is attached`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the volume`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project id the volume is in`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size in GB of the volume`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Performance plan the volume is on`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle, defaults to hourly`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility slug the volume resides in`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the volume is locked or not`,
				},
				resource.Attribute{
					Name:        "device_ids",
					Description: `UUIDs of devices to which this volume is attached`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"metal_device":               0,
		"metal_device_bgp_neighbors": 1,
		"metal_facility":             2,
		"metal_ip_block_ranges":      3,
		"metal_operating_system":     4,
		"metal_organization":         5,
		"metal_precreated_ip_block":  6,
		"metal_project":              7,
		"metal_project_ssh_key":      8,
		"metal_spot_market_price":    9,
		"metal_spot_market_request":  10,
		"metal_volume":               11,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
