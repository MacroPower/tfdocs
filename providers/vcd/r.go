package vcd

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director catalog resource. This can be used to create and delete a catalog.`,
			Description:      ``,
			Keywords: []string{
				"catalog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Catalog name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) - Description of catalog`,
				},
				resource.Attribute{
					Name:        "delete_recursive",
					Description: `(Required) - When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that normally allows removal`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog_item",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director catalog item resource. This can be used to upload and delete OVA file inside a catalog.`,
			Description:      ``,
			Keywords: []string{
				"catalog",
				"item",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) The name of the catalog where to upload OVA file`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Item name in catalog`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) - Description of item`,
				},
				resource.Attribute{
					Name:        "ova_path",
					Description: `(Required) - Absolute or relative path to file to upload`,
				},
				resource.Attribute{
					Name:        "upload_piece_size",
					Description: `(Optional) - Size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.`,
				},
				resource.Attribute{
					Name:        "show_upload_progress",
					Description: `(Optional) - Default false. Allows to see upload progress`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog_media",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director media resource. This can be used to upload and delete media (ISO) file inside a catalog.`,
			Description:      ``,
			Keywords: []string{
				"catalog",
				"media",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) The name of the catalog where to upload media file`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Media file name in catalog`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) - Description of media file`,
				},
				resource.Attribute{
					Name:        "media_path",
					Description: `(Required) - Absolute or relative path to file to upload`,
				},
				resource.Attribute{
					Name:        "upload_piece_size",
					Description: `(Optional) - size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.`,
				},
				resource.Attribute{
					Name:        "show_upload_progress",
					Description: `(Optional) - Default false. Allows to see upload progress`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "is_iso",
					Description: `(Computed) returns True if this media file is ISO`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Computed) returns owner name`,
				},
				resource.Attribute{
					Name:        "is_published",
					Description: `(Computed) returns True if this media file is in a published catalog`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `(Computed) returns creation date`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Computed) returns media storage in Bytes`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) returns media status`,
				},
				resource.Attribute{
					Name:        "storage_profile_name",
					Description: `(Computed) returns storage profile name ## Importing Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_edgegateway",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director edge gateway. This can be used to create and delete edge gateways connected to one or more external networks.`,
			Description:      ``,
			Keywords: []string{
				"edgegateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the VDC belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC that owns the edge gateway. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the edge gateway.`,
				},
				resource.Attribute{
					Name:        "external_network",
					Description: `(Required,`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Required) Configuration of the vShield edge VM for this gateway. One of: ` + "`" + `compact` + "`" + `, ` + "`" + `full` + "`" + ` ("Large"), ` + "`" + `x-large` + "`" + `, ` + "`" + `full4` + "`" + ` ("Quad Large").`,
				},
				resource.Attribute{
					Name:        "ha_enabled",
					Description: `(Optional) Enable high availability on this edge gateway. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "distributed_routing",
					Description: `(Optional) If advanced networking enabled, also enable distributed routing. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fips_mode_enabled",
					Description: `(Optional) When FIPS mode is enabled, any secure communication to or from the NSX Edge uses cryptographic algorithms or protocols that are allowed by United States Federal Information Processing Standards (FIPS). FIPS mode turns on the cipher suites that comply with FIPS. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_default_route_for_dns_relay",
					Description: `(Optional) When default route is set, it will be used for gateways' default routing and DNS forwarding. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lb_enabled",
					Description: `(Optional) Enable load balancing. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lb_acceleration_enabled",
					Description: `(Optional) Enable to configure the load balancer to use the faster L4 engine rather than L7 engine. The L4 TCP VIP is processed before the edge gateway firewall so no ` + "`" + `allow` + "`" + ` firewall rule is required. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lb_logging_enabled",
					Description: `(Optional) Enables the edge gateway load balancer to collect traffic logs. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lb_loglevel",
					Description: `(Optional) Choose the severity of events to be logged. One of ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `info` + "`" + `, ` + "`" + `debug` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_edgegateway_settings",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director edge gateway global settings. This can be used to update global edge gateways settings related to firewall and load balancing.`,
			Description:      ``,
			Keywords: []string{
				"edgegateway",
				"settings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the VDC belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC that owns the edge gateway. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_name",
					Description: `(Optional) A unique name for the edge gateway. (Required if ` + "`" + `edge_gateway_id` + "`" + ` is not set)`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Optional) The edge gateway ID. (Required if ` + "`" + `edge_gateway_name` + "`" + ` is not set)`,
				},
				resource.Attribute{
					Name:        "lb_enabled",
					Description: `(Optional) Enable load balancing. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lb_acceleration_enabled",
					Description: `(Optional) Enable to configure the load balancer.`,
				},
				resource.Attribute{
					Name:        "lb_logging_enabled",
					Description: `(Optional) Enables the edge gateway load balancer to collect traffic logs. Default is ` + "`" + `false` + "`" + `. Note:`,
				},
				resource.Attribute{
					Name:        "lb_loglevel",
					Description: `(Optional) Choose the severity of events to be logged. One of ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `info` + "`" + `, ` + "`" + `debug` + "`" + `. Note:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_edgegateway_vpn",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director IPsec VPN. This can be used to create, modify, and delete VPN settings and rules.`,
			Description:      ``,
			Keywords: []string{
				"edgegateway",
				"vpn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the Firewall Rules`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPN`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the VPN`,
				},
				resource.Attribute{
					Name:        "encryption_protocol",
					Description: `(Required) - E.g. ` + "`" + `AES256` + "`" + ``,
				},
				resource.Attribute{
					Name:        "local_ip_address",
					Description: `(Required) - Local IP Address`,
				},
				resource.Attribute{
					Name:        "local_id",
					Description: `(Required) - Local ID`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Required) - The MTU setting`,
				},
				resource.Attribute{
					Name:        "peer_ip_address",
					Description: `(Required) - Peer IP Address`,
				},
				resource.Attribute{
					Name:        "peer_id",
					Description: `(Required) - Peer ID`,
				},
				resource.Attribute{
					Name:        "shared_secret",
					Description: `(Required) - Shared Secret`,
				},
				resource.Attribute{
					Name:        "local_subnets",
					Description: `(Required) - List of Local Subnets see [Local Subnets](#localsubnets) below for details.`,
				},
				resource.Attribute{
					Name:        "peer_subnets",
					Description: `(Required) - List of Peer Subnets see [Peer Subnets](#peersubnets) below for details.`,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "local_subnet_name",
					Description: `(Required) Name of the local subnet`,
				},
				resource.Attribute{
					Name:        "local_subnet_gateway",
					Description: `(Required) Gateway of the local subnet`,
				},
				resource.Attribute{
					Name:        "local_subnet_mask",
					Description: `(Required) Subnet mask of the local subnet <a id="peersubnets"></a> ## Peer Subnets Each Peer Subnet supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "peer_subnet_name",
					Description: `(Required) Name of the peer subnet`,
				},
				resource.Attribute{
					Name:        "peer_subnet_gateway",
					Description: `(Required) Gateway of the peer subnet`,
				},
				resource.Attribute{
					Name:        "peer_subnet_mask",
					Description: `(Required) Subnet mask of the peer subnet`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_external_network",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director external network resource. This can be used to create and delete external networks.`,
			Description:      ``,
			Keywords: []string{
				"external",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Network friendly description`,
				},
				resource.Attribute{
					Name:        "ip_scope",
					Description: `(Required) A list of IP scopes for the network. See [IP Scope](#ipscope) below for details.`,
				},
				resource.Attribute{
					Name:        "vsphere_network",
					Description: `(Required) A list of DV_PORTGROUP or NETWORK objects names that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server registered with the system. See [vSphere Network](#vspherenetwork) below for details.`,
				},
				resource.Attribute{
					Name:        "retain_net_info_across_deployments",
					Description: `(Optional) Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false. <a id="ipscope"></a> ## IP Scope`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required) Gateway of the network`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Required) Network mask`,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) Primary DNS server`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Secondary DNS server`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Required) IP ranges used for static pool allocation in the network. See [IP Pool](#ip-pool) below for details. <a id="ip-pool"></a> ## IP Pool`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) Start address of the IP range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) End address of the IP range <a id="vspherenetwork"></a> ## vSphere Network`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Port group name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The vSphere type of the object. One of: DV_PORTGROUP (distributed virtual port group), NETWORK (standard switch port group)`,
				},
				resource.Attribute{
					Name:        "vcenter",
					Description: `(Required) The vCenter server name ## Importing Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_external_network_v2",
			Category:         "Resources",
			ShortDescription: `Provides a VMware Cloud Director External Network resource (version 2). New version of this resource uses new VCD API and is capable of creating NSX-T backed external networks as well as port group backed ones.`,
			Description:      ``,
			Keywords: []string{
				"external",
				"network",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Network friendly description`,
				},
				resource.Attribute{
					Name:        "ip_scope",
					Description: `(Required) One or more IP scopes for the network. See [IP Scope](#ipscope) below for details.`,
				},
				resource.Attribute{
					Name:        "vsphere_network",
					Description: `(Optional) One or more blocks of [vSphere Network](#vspherenetwork)..`,
				},
				resource.Attribute{
					Name:        "nsxt_network",
					Description: `(Optional) NSX-T network definition. See [NSX-T Network](#nsxtnetwork) below for details. <a id="ipscope"></a> ## IP Scope`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required) Gateway of the network.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Required) Network prefix.`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Required) IP ranges used for static pool allocation in the network. See [IP Pool](#ip-pool) below for details.`,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) Primary DNS server.`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Secondary DNS server.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Default is ` + "`" + `true` + "`" + `. <a id="ip-pool"></a> ## IP Pool`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) Start address of the IP range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) End address of the IP range <a id="vspherenetwork"></a> ## vSphere Network`,
				},
				resource.Attribute{
					Name:        "vcenter_id",
					Description: `(Required) vCenter ID. Can be looked up using [` + "`" + `vcd_vcenter` + "`" + `](/docs/providers/vcd/d/vcenter.html) data source.`,
				},
				resource.Attribute{
					Name:        "portgroup_id",
					Description: `(Required) vSphere portgroup ID. Can be looked up using [` + "`" + `vcd_portgroup` + "`" + `](/docs/providers/vcd/d/portgroup.html) data source. <a id="nsxtnetwork"></a> ## NSX-T Network`,
				},
				resource.Attribute{
					Name:        "nsxt_manager_id",
					Description: `(Required) NSX-T manager ID. Can be looked up using [` + "`" + `vcd_nsxt_manager` + "`" + `](/docs/providers/vcd/d/nsxt_manager.html) data source.`,
				},
				resource.Attribute{
					Name:        "nsxt_tier0_router_id",
					Description: `(Required) NSX-T Tier-0 router ID. Can be looked up using [` + "`" + `vcd_nsxt_tier0_router` + "`" + `](/docs/providers/vcd/d/nsxt_tier0_router.html) data source. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_independent_disk",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director independent disk resource. This can be used to create and delete independent disks.`,
			Description:      ``,
			Keywords: []string{
				"independent",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Disk name`,
				},
				resource.Attribute{
					Name:        "size_in_mb",
					Description: `(Required,`,
				},
				resource.Attribute{
					Name:        "bus_type",
					Description: `(Optional) Disk bus type. Values can be: ` + "`" + `IDE` + "`" + `, ` + "`" + `SCSI` + "`" + `, ` + "`" + `SATA` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bus_sub_type",
					Description: `(Optional) Disk bus subtype. Values can be: ` + "`" + `buslogic` + "`" + `, ` + "`" + `lsilogic` + "`" + `, ` + "`" + `lsilogicsas` + "`" + `, ` + "`" + `VirtualSCSI` + "`" + ` for ` + "`" + `SCSI` + "`" + ` and ` + "`" + `ahci` + "`" + ` for ` + "`" + `SATA` + "`" + ``,
				},
				resource.Attribute{
					Name:        "storage_profile",
					Description: `(Optional) The name of storage profile where disk will be created ## Attribute reference Supported in provider`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Computed) IOPS request for the created disk`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Computed) The owner name of the disk`,
				},
				resource.Attribute{
					Name:        "datastore_name",
					Description: `(Computed) Data store name. Readable only for system user.`,
				},
				resource.Attribute{
					Name:        "is_attached",
					Description: `(Computed) True if the disk is already attached ## Importing Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_inserted_media",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director resource for inserting or ejecting media (ISO) file for the VM. Create this resource for inserting the media, and destroy it for ejecting.`,
			Description:      ``,
			Keywords: []string{
				"inserted",
				"media",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) The name of the catalog where to find media file`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Media file name in catalog which will be inserted to VM`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) - The name of vApp to find`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Required) - The name of VM to be used to insert media file`,
				},
				resource.Attribute{
					Name:        "eject_force",
					Description: `(Optional;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_ip_set",
			Category:         "Resources",
			ShortDescription: `Provides an IP set resource.`,
			Description:      ``,
			Keywords: []string{
				"nsxv",
				"ip",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique IP set name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description for IP set.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Required) A set of IP addresses, CIDRs and ranges as strings.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of IP set ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of IP set ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_app_profile",
			Category:         "Resources",
			ShortDescription: `Provides an NSX edge gateway load balancer application profile resource.`,
			Description:      ``,
			Keywords: []string{
				"lb",
				"app",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the application profile is to be created`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Application profile name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Protocol type used to send requests to the server. One of ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `http` + "`" + `, or ` + "`" + `https` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enable_ssl_passthrough",
					Description: `(Optional) Enable SSL authentication to be passed through to the virtual server. Otherwise SSL authentication takes place at the destination address`,
				},
				resource.Attribute{
					Name:        "http_redirect_url",
					Description: `(Optional) The URL to which traffic that arrives at the destination address should be redirected. Only applies for types ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + ``,
				},
				resource.Attribute{
					Name:        "persistence_mechanism",
					Description: `(Optional) Persistence mechanism for the profile. One of 'cookie', 'ssl-sessionid', 'sourceip'`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) Used to uniquely identify the session the first time a client accesses the site. The load balancer refers to this cookie when connecting subsequent requests in the session, so that they all go to the same virtual server. Only applies for ` + "`" + `persistence_mechanism` + "`" + ` 'cookie'`,
				},
				resource.Attribute{
					Name:        "cookie_mode",
					Description: `(Optional) The mode by which the cookie should be inserted. One of 'insert', 'prefix', or 'appsession'`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) Length of time in seconds that persistence stays in effect`,
				},
				resource.Attribute{
					Name:        "insert_x_forwarded_http_header",
					Description: `(Optional) Enables 'X-Forwarded-For' header for identifying the originating IP address of a client connecting to a Web server through the load balancer. Only applies for types ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enable_pool_side_ssl",
					Description: `(Optional) Enable to define the certificate, CAs, or CRLs used to authenticate the load balancer from the server side.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The NSX ID of the load balancer application profile ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The NSX ID of the load balancer application profile ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_app_rule",
			Category:         "Resources",
			ShortDescription: `Provides an NSX edge gateway load balancer application rule resource.`,
			Description:      ``,
			Keywords: []string{
				"lb",
				"app",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the application rule is to be created`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Application rule name`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Required) A multiline application rule script. Terraform's [HEREDOC syntax](https://www.terraform.io/docs/configuration/expressions.html#string-literals) may be useful for multiline scripts.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The NSX ID of the load balancer application rule ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The NSX ID of the load balancer application rule ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_server_pool",
			Category:         "Resources",
			ShortDescription: `Provides an NSX edge gateway load balancer server pool resource.`,
			Description:      ``,
			Keywords: []string{
				"lb",
				"server",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the server pool is to be created`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Server Pool name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Server Pool description`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Required) Server Pool load balancing method. Can be one of ` + "`" + `ip-hash` + "`" + `, ` + "`" + `round-robin` + "`" + `, ` + "`" + `uri` + "`" + `, ` + "`" + `leastconn` + "`" + `, ` + "`" + `url` + "`" + `, or ` + "`" + `httpheader` + "`" + ``,
				},
				resource.Attribute{
					Name:        "algorithm_parameters",
					Description: `(Optional) Valid only when ` + "`" + `algorithm` + "`" + ` is ` + "`" + `httpheader` + "`" + ` or ` + "`" + `url` + "`" + `. The ` + "`" + `httpheader` + "`" + ` algorithm parameter has one option ` + "`" + `headerName=<name>` + "`" + ` while the ` + "`" + `url` + "`" + ` algorithm parameter has option ` + "`" + `urlParam=<url>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_transparency",
					Description: `(Optional) When transparency is ` + "`" + `false` + "`" + ` (default) backend servers see the IP address of the traffic source as the internal IP address of the load balancer. When it is ` + "`" + `true` + "`" + ` the source IP address is the actual IP address of the client and the edge gateway must be set as the default gateway to ensure that return packets go through the edge gateway.`,
				},
				resource.Attribute{
					Name:        "monitor_id",
					Description: `(Optional) ` + "`" + `vcd_lb_service_monitor` + "`" + ` resource ` + "`" + `id` + "`" + ` to attach to server pool for health check parameters`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Optional) A block to define server pool members. Multiple can be used. See [Member](#member) and example for usage details. <a id="member"></a> ## Member`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Required) State of member in a pool. One of ` + "`" + `enabled` + "`" + `, ` + "`" + `disabled` + "`" + `, or ` + "`" + `drain` + "`" + `. When member condition is set to ` + "`" + `drain` + "`" + ` it stops taking new connections and calls, while it allows its sessions on existing connections to continue until they naturally end. This allows to gracefully remove member node from load balancing rotation.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Member name`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Member IP address`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port at which the member is to receive traffic from the load balancer.`,
				},
				resource.Attribute{
					Name:        "monitor_port",
					Description: `(Required) Monitor Port at which the member is to receive health monitor requests.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) The proportion of traffic this member is to handle. Must be an integer in the range 1-256.`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `(Optional) The maximum number of concurrent connections the member can handle.`,
				},
				resource.Attribute{
					Name:        "min_connections",
					Description: `(Optional) The minimum number of concurrent connections a member must always accept. ## Attribute Reference The following attributes are exported on this resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The NSX ID of the load balancer server pool Additionally each of members defined in blocks expose their own ` + "`" + `id` + "`" + ` fields as well ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The NSX ID of the load balancer server pool Additionally each of members defined in blocks expose their own ` + "`" + `id` + "`" + ` fields as well ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_service_monitor",
			Category:         "Resources",
			ShortDescription: `Provides an NSX edge gateway load balancer service monitor resource.`,
			Description:      ``,
			Keywords: []string{
				"lb",
				"service",
				"monitor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the service monitor is to be created`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Service Monitor name`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Interval in seconds at which a server is to be monitored using the specified Method. Defaults to 10`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Maximum time in seconds within which a response from the server must be received. Defaults to 15`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `(Optional) Number of times the specified monitoring Method must fail sequentially before the server is declared down. Defaults to 3`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Select the way in which you want to send the health check request to the server — ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `udp` + "`" + `. Depending on the type selected, the remaining attributes are allowed or not`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) For types ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + `. Select http method to be used to detect server status. One of OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) For types ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + `. URL to be used in the server status request`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `(Optional) For types ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, and ` + "`" + `udp` + "`" + `. The data to be sent.`,
				},
				resource.Attribute{
					Name:        "expected",
					Description: `(Optional) For types ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + `. String that the monitor expects to match in the status line of the HTTP or HTTPS response (for example, ` + "`" + `HTTP/1.1` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "receive",
					Description: `(Optional) For types ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, and ` + "`" + `udp` + "`" + `. The string to be matched in the response content.`,
				},
				resource.Attribute{
					Name:        "extension",
					Description: `(Optional) A map of advanced monitor parameters as key=value pairs (i.e. ` + "`" + `max-age=SECONDS` + "`" + `, ` + "`" + `invert-regex` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The NSX ID of the load balancer service monitor ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The NSX ID of the load balancer service monitor ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_virtual_server",
			Category:         "Resources",
			ShortDescription: `Provides an NSX edge gateway load balancer virtual server resource.`,
			Description:      ``,
			Keywords: []string{
				"lb",
				"virtual",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the virtual server is to be created`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Virtual server name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Virtual server description`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Defines if the virtual server is enabled. Default ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enable_acceleration",
					Description: `(Optional) Defines if the virtual server uses acceleration. Default ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Set the IP address that the load balancer listens on`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Select the protocol that the virtual server accepts. One of ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `http` + "`" + `, or ` + "`" + `https` + "`" + ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port number that the load balancer listens on`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional) Maximum concurrent connections that the virtual server can process`,
				},
				resource.Attribute{
					Name:        "connection_rate_limit",
					Description: `(Optional) Maximum incoming new connection requests per second`,
				},
				resource.Attribute{
					Name:        "server_pool_id",
					Description: `(Optional) The server pool that the load balancer will use`,
				},
				resource.Attribute{
					Name:        "app_profile_id",
					Description: `(Optional) Application profile ID to be associated with the virtual server`,
				},
				resource.Attribute{
					Name:        "app_rule_ids",
					Description: `(Optional) List of attached application rule IDs ## Attribute Reference The following attributes are exported on the base level of this resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The NSX ID of the load balancer virtual server ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The NSX ID of the load balancer virtual server ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_direct",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Org VDC Network attached to an external one. This can be used to create, modify, and delete internal networks for vApps to connect.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"direct",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "external_network",
					Description: `(Required) The name of the external network.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Defines if this network is shared between multiple VDCs in the Org. Defaults to ` + "`" + `false` + "`" + `. ## Attribute reference Supported in provider`,
				},
				resource.Attribute{
					Name:        "external_network_gateway",
					Description: `(Computed) returns the gateway from the external network`,
				},
				resource.Attribute{
					Name:        "external_network_netmask",
					Description: `(Computed) returns the netmask from the external network`,
				},
				resource.Attribute{
					Name:        "external_network_dns1",
					Description: `(Computed) returns the first DNS from the external network`,
				},
				resource.Attribute{
					Name:        "external_network_dns2",
					Description: `(Computed) returns the second DNS from the external network`,
				},
				resource.Attribute{
					Name:        "external_network_dns_suffix",
					Description: `(Computed) returns the DNS suffix from the external network ## Importing Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_isolated",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Org VDC isolated Network. This can be used to create, modify, and delete internal networks for vApps to connect.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"isolated",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) The netmask for the new network. Defaults to ` + "`" + `255.255.255.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) First DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Second DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) A FQDN for the virtual machines on this network`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Defines if this network is shared between multiple VDCs in the Org. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dhcp_pool",
					Description: `(Optional) A range of IPs to issue to virtual machines that don't have a static IP; see [IP Pools](#ip-pools) below for details.`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Optional) A range of IPs permitted to be used as static IPs for virtual machines; see [IP Pools](#ip-pools) below for details. <a id="ip-pools"></a> ## IP Pools Static IP Pools and DHCP Pools support the following attributes:`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) The first address in the IP Range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) The final address in the IP Range DHCP Pools additionally support the following attributes:`,
				},
				resource.Attribute{
					Name:        "default_lease_time",
					Description: `(Optional) The default DHCP lease time to use. Defaults to ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_lease_time",
					Description: `(Optional) The maximum DHCP lease time to use. Defaults to ` + "`" + `7200` + "`" + `. ## Importing Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_isolated_v2",
			Category:         "Resources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC isolated Network. This can be used to create, modify, and delete isolated VDC networks (backed by NSX-T or NSX-V).`,
			Description:      ``,
			Keywords: []string{
				"network",
				"isolated",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the network`,
				},
				resource.Attribute{
					Name:        "is_shared",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Required) The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).`,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) First DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Second DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) A FQDN for the virtual machines on this network`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Optional) A range of IPs permitted to be used as static IPs for virtual machines; see [IP Pools](#ip-pools) below for details. <a id="ip-pools"></a> ## IP Pools Static IP Pools support the following attributes:`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) The first address in the IP Range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) The final address in the IP Range ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_routed",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Org VDC routed Network. This can be used to create, modify, and delete internal networks for vApps to connect.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"routed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) The netmask for the new network. Defaults to ` + "`" + `255.255.255.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) First DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Second DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) A FQDN for the virtual machines on this network`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Defines if this network is shared between multiple VDCs in the Org. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dhcp_pool",
					Description: `(Optional) A range of IPs to issue to virtual machines that don't have a static IP; see [IP Pools](#ip-pools) below for details.`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Optional) A range of IPs permitted to be used as static IPs for virtual machines; see [IP Pools](#ip-pools) below for details. <a id="ip-pools"></a> ## IP Pools Static IP Pools and DHCP Pools support the following attributes:`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) The first address in the IP Range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) The final address in the IP Range DHCP Pools additionally support the following attribute:`,
				},
				resource.Attribute{
					Name:        "max_lease_time",
					Description: `(Optional) The maximum DHCP lease time to use. Defaults to ` + "`" + `7200` + "`" + `. ## Importing Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_routed_v2",
			Category:         "Resources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC routed Network. This can be used to create, modify, and delete routed VDC networks (backed by NSX-T or NSX-V).`,
			Description:      ``,
			Keywords: []string{
				"network",
				"routed",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the network`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `(Optional) An interface for the network. One of ` + "`" + `internal` + "`" + ` (default), ` + "`" + `subinterface` + "`" + `, ` + "`" + `distributed` + "`" + ` (requires the edge gateway to support distributed networks). NSX-T supports only ` + "`" + `internal` + "`" + ``,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID name of the edge gateway (NSX-V or NSX-T)`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Required) The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).`,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) First DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Second DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) A FQDN for the virtual machines on this network`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Optional) A range of IPs permitted to be used as static IPs for virtual machines; see [IP Pools](#ip-pools) below for details. <a id="ip-pools"></a> ## IP Pools Static IP Pools and DHCP Pools support the following attributes:`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) The first address in the IP Range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) The final address in the IP Range ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_edgegateway",
			Category:         "Resources",
			ShortDescription: `Provides a VMware Cloud Director NSX-T edge gateway. This can be used to create, update, and delete NSX-T edge gateways connected to external networks.`,
			Description:      ``,
			Keywords: []string{
				"nsxt",
				"edgegateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the VDC belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC that owns the edge gateway. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the edge gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A unique name for the edge gateway.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `(Required) An external network ID.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) One or more [subnets](#edgegateway-subnet) defined for edge gateway.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `(Optional) Specific Edge Cluster ID if required`,
				},
				resource.Attribute{
					Name:        "dedicate_external_network",
					Description: `(Optional) Dedicating the External Network will enable Route Advertisement for this Edge Gateway. Default ` + "`" + `false` + "`" + `. <a id="edgegateway-subnet"></a> ## Edge Gateway Subnet`,
				},
				resource.Attribute{
					Name:        "primary_ip",
					Description: `Primary IP address exposed for an easy access without nesting. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_ip",
					Description: `Primary IP address exposed for an easy access without nesting. ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_network_dhcp",
			Category:         "Resources",
			ShortDescription: `Provides a resource to manage DHCP pools for NSX-T Org VDC Routed network.`,
			Description:      ``,
			Keywords: []string{
				"nsxt",
				"network",
				"dhcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "org_network_id",
					Description: `(Required) ID of parent Org VDC Routed network`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Required) One or more blocks to define DHCP pool ranges. See [Pools](#pools) and example for usage details. ## Pools`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) Start address of DHCP pool range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) End address of DHCP pool range ## Importing ~> The current implementation of Terraform import can only import resources into the state. It does not generate configuration. [More information.](https://www.terraform.io/docs/import/) An existing DHCP configuration can be [imported][docs-import] into this resource via supplying the full dot separated path for your Org VDC network. An example is below: [docs-import]: https://www.terraform.io/docs/import/ ` + "`" + `` + "`" + `` + "`" + ` terraform import vcd_nsxt_network_dhcp.imported my-org.my-org-vdc.my-nsxt-vdc-network-name ` + "`" + `` + "`" + `` + "`" + ` The above would import the DHCP config settings that are defined on VDC network ` + "`" + `my-nsxt-vdc-network-name` + "`" + ` which is configured in organization named ` + "`" + `my-org` + "`" + ` and VDC named ` + "`" + `my-org-vdc` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_network_imported",
			Category:         "Resources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC NSX-T Imported Network type. This can be used to create, modify, and delete NSX-T VDC networks of Imported type (backed by NSX-T).`,
			Description:      ``,
			Keywords: []string{
				"nsxt",
				"network",
				"imported",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "nsxt_logical_switch_name",
					Description: `(Required) Unique name of an existing NSX-T segment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the network`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Required) The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).`,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) First DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Second DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) A FQDN for the virtual machines on this network`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Optional) A range of IPs permitted to be used as static IPs for virtual machines; see [IP Pools](#ip-pools) below for details. <a id="ip-pools"></a> ## IP Pools Static IP Pools support the following attributes:`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) The first address in the IP Range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) The final address in the IP Range ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "nsxt_logical_switch_id",
					Description: `ID of an existing NSX-T segment ## Importing ~> After import the field ` + "`" + `nsxt_logical_switch_name` + "`" + ` will remain empty because it is impossible to read it in API once it is consumed by network. ~> The current implementation of Terraform import can only import resources into the state. It does not generate configuration. [More information.][docs-import] An existing NSX-T VDC Imported network can be [imported][docs-import] into this Terraform resource via supplying its path. The path for this resource is made of orgName.vdcName.networkName. For example, using this structure, representing an NSX-T Imported Network that was`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nsxt_logical_switch_id",
					Description: `ID of an existing NSX-T segment ## Importing ~> After import the field ` + "`" + `nsxt_logical_switch_name` + "`" + ` will remain empty because it is impossible to read it in API once it is consumed by network. ~> The current implementation of Terraform import can only import resources into the state. It does not generate configuration. [More information.][docs-import] An existing NSX-T VDC Imported network can be [imported][docs-import] into this Terraform resource via supplying its path. The path for this resource is made of orgName.vdcName.networkName. For example, using this structure, representing an NSX-T Imported Network that was`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_dhcp_relay",
			Category:         "Resources",
			ShortDescription: `Provides an NSX edge gateway DHCP relay configuration resource.`,
			Description:      ``,
			Keywords: []string{
				"nsxv",
				"dhcp",
				"relay",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which DHCP relay is to be configured.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Optional) A set of IP addresses.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Optional) A set of domain names.`,
				},
				resource.Attribute{
					Name:        "ip_sets",
					Description: `(Optional) A set of IP set names.`,
				},
				resource.Attribute{
					Name:        "relay_agent",
					Description: `(Required) One or more blocks to define Org network and optional IP address of edge gateway interfaces from which DHCP messages are to be relayed to the external DHCP relay server(s). See [Relay Agent](#relay-agent) and example for usage details. <a id="relay-agent"></a> ## Relay Agent`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `(Required) An existing Org network name from which DHCP messages are to be relayed.`,
				},
				resource.Attribute{
					Name:        "gateway_ip_address",
					Description: `(Optional) IP address on edge gateway to be used for relaying messages. Primary address of edge gateway interface will be picked if not specified. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_dnat",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director DNAT resource for advanced edge gateways (NSX-V). This can be used to create, modify, and delete destination NATs to map an external IP/port to an internal IP/port.`,
			Description:      ``,
			Keywords: []string{
				"nsxv",
				"dnat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the DNAT rule.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Required) Type of the network on which to apply the DNAT rule. Possible values ` + "`" + `org` + "`" + ` or ` + "`" + `ext` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `(Required) The name of the network on which to apply the DNAT rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Defines if the rule is enabaled. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging_enabled",
					Description: `(Optional) Defines if the logging for this rule is enabaled. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Free text description.`,
				},
				resource.Attribute{
					Name:        "rule_tag",
					Description: `(Optional) This can be used to specify user-controlled rule tag. If not specified, it will report rule ID after creation. Must be between 65537-131072.`,
				},
				resource.Attribute{
					Name:        "original_address",
					Description: `(Required) IP address, range or subnet. This address must be the public IP address of the edge gateway for which you are configuring the DNAT rule. In the packet being inspected, this IP address or range would be those that appear as the destination IP address of the packet. These packet destination addresses are the ones translated by this DNAT rule.`,
				},
				resource.Attribute{
					Name:        "original_port",
					Description: `(Optional) Select the port or port range that the incoming traffic uses on the edge gateway to connect to the internal network on which the virtual machines are connected. This selection is not available when the Protocol is set to ` + "`" + `icmp` + "`" + ` or ` + "`" + `any` + "`" + `. Default ` + "`" + `any` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `(Required) IP address, range or subnet. IP addresses to which destination addresses on inbound packets will be translated. These addresses are the IP addresses of the one or more virtual machines for which you are configuring DNAT so that they can receive traffic from the external network.`,
				},
				resource.Attribute{
					Name:        "translated_port",
					Description: `(Optional) Select the port or port range that inbound traffic is connecting to on the virtual machines on the internal network. These ports are the ones into which the DNAT rule is translating for the packets inbound to the virtual machines.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Select the protocol to which the rule applies. One of ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `any` + "`" + `. Default ` + "`" + `any` + "`" + ` protocols, select Any.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) Only when ` + "`" + `protocol` + "`" + ` is set to ` + "`" + `icmp` + "`" + `. One of ` + "`" + `any` + "`" + `, ` + "`" + `address-mask-request` + "`" + `, ` + "`" + `address-mask-reply` + "`" + `, ` + "`" + `destination-unreachable` + "`" + `, ` + "`" + `echo-request` + "`" + `, ` + "`" + `echo-reply` + "`" + `, ` + "`" + `parameter-problem` + "`" + `, ` + "`" + `redirect` + "`" + `, ` + "`" + `router-advertisement` + "`" + `, ` + "`" + `router-solicitation` + "`" + `, ` + "`" + `source-quench` + "`" + `, ` + "`" + `time-exceeded` + "`" + `, ` + "`" + `timestamp-request` + "`" + `, ` + "`" + `timestamp-reply` + "`" + `. Default ` + "`" + `any` + "`" + ` ## Attribute Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Possible values - ` + "`" + `user` + "`" + `, ` + "`" + `internal_high` + "`" + `. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_type",
					Description: `Possible values - ` + "`" + `user` + "`" + `, ` + "`" + `internal_high` + "`" + `. ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_firewall_rule",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director firewall rule resource for advanced edge gateways (NSX-V). This can be used to create, modify, and delete firewall rules.`,
			Description:      ``,
			Keywords: []string{
				"nsxv",
				"firewall",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the firewall rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Defines if the rule is set to ` + "`" + `accept` + "`" + ` or ` + "`" + `deny` + "`" + ` traffic. Default ` + "`" + `accept` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Defines if the rule is enabaled. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging_enabled",
					Description: `(Optional) Defines if the logging for this rule is enabaled. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Free text name. Can be duplicate.`,
				},
				resource.Attribute{
					Name:        "rule_tag",
					Description: `(Optional) This can be used to specify user-controlled rule tag. If not specified, it will report rule ID after creation. Must be between 65537-131072.`,
				},
				resource.Attribute{
					Name:        "above_rule_id",
					Description: `(Optional) This can be used to alter default rule placement order. By default every rule is appended to the end of firewall rule list. When a value of another rule is set - this rule will be placed above the specified rule.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Exactly one block to define source criteria for firewall. See [Endpoint](#endpoint) and example for usage details.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Exactly one block to define source criteria for firewall. See [Endpoint](#endpoint) and example for usage details.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) One or more blocks to define protocol and port details. Use multiple blocks if you want to define multiple port/protocol combinations for the same rule. See [Service](#service) and example for usage details. <a id="endpoint"></a> ## Endpoint (source or destination)`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(Optional) When the toggle exclusion is selected, the rule is applied to traffic on all sources except for the locations you excluded. When the toggle exclusion is not selected, the rule applies to traffic you specified. Default ` + "`" + `false` + "`" + `. This [example](#example-usage-3-use-exclusion-in-source-) uses it.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Optional) A set of IP addresses, CIDRs or ranges. A keyword ` + "`" + `any` + "`" + ` is also accepted as a parameter.`,
				},
				resource.Attribute{
					Name:        "gateway_interfaces",
					Description: `(Optional) A set of with either three keywords ` + "`" + `vse` + "`" + ` (UI names it as ` + "`" + `any` + "`" + `), ` + "`" + `internal` + "`" + `, ` + "`" + `external` + "`" + ` or an org network name. It automatically looks up vNic in the backend.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) A set of ` + "`" + `.id` + "`" + ` fields of ` + "`" + `vcd_vapp_vm` + "`" + ` resources.`,
				},
				resource.Attribute{
					Name:        "org_networks",
					Description: `(Optional) A set of org network names.`,
				},
				resource.Attribute{
					Name:        "ip_sets",
					Description: `(Optional) A set of existing IP set names (either created manually or configured using ` + "`" + `vcd_nsxv_ip_set` + "`" + ` resource) <a id="service"></a> ## Service`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) One of ` + "`" + `any` + "`" + `, ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + ` to apply.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port number or range separated by ` + "`" + `-` + "`" + ` for port number. Default 'any'.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) Port number or range separated by ` + "`" + `-` + "`" + ` for port number. Default 'any'. ## Attribute Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Possible values - ` + "`" + `user` + "`" + `, ` + "`" + `internal_high` + "`" + `. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_type",
					Description: `Possible values - ` + "`" + `user` + "`" + `, ` + "`" + `internal_high` + "`" + `. ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_snat",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director SNAT resource for advanced edge gateways (NSX-V). This can be used to create, modify, and delete source NATs to allow vApps to send external traffic.`,
			Description:      ``,
			Keywords: []string{
				"nsxv",
				"snat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the SNAT rule.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Required) Type of the network on which to apply the DNAT rule. Possible values ` + "`" + `org` + "`" + ` or ` + "`" + `ext` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `(Required) The name of the network on which to apply the SNAT rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Defines if the rule is enabaled. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging_enabled",
					Description: `(Optional) Defines if the logging for this rule is enabaled. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Free text description.`,
				},
				resource.Attribute{
					Name:        "rule_tag",
					Description: `(Optional) This can be used to specify user-controlled rule tag. If not specified, it will report rule ID after creation. Must be between 65537-131072.`,
				},
				resource.Attribute{
					Name:        "original_address",
					Description: `(Required) IP address, range or subnet. These addresses are the IP addresses of one or more virtual machines for which you are configuring the SNAT rule so that they can send traffic to the external network.`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `(Required) IP address, range or subnet. This address is always the public IP address of the gateway for which you are configuring the SNAT rule. Specifies the IP address to which source addresses (the virtual machines) on outbound packets are translated to when they send traffic to the external network. ## Attribute Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Possible values - ` + "`" + `user` + "`" + `, ` + "`" + `internal_high` + "`" + `. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_type",
					Description: `Possible values - ` + "`" + `user` + "`" + `, ` + "`" + `internal_high` + "`" + `. ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Organization resource. This can be used to create delete, and update an organization.`,
			Description:      ``,
			Keywords: []string{
				"org",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Org name`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `(Required) Org full name`,
				},
				resource.Attribute{
					Name:        "delete_recursive",
					Description: `(Required) - pass ` + "`" + `delete_recursive` + "`" + `=true as query parameter to remove an organization or VDC and any objects it contains that are in a state that normally allows removal.`,
				},
				resource.Attribute{
					Name:        "delete_force",
					Description: `(Required) - pass ` + "`" + `delete_force=true` + "`" + ` and ` + "`" + `delete_recursive=true` + "`" + ` to remove an organization or VDC and any objects it contains, regardless of their state.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) - True if this organization is enabled (allows login and all other operations). Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) - Org description. Default is empty.`,
				},
				resource.Attribute{
					Name:        "deployed_vm_quota",
					Description: `(Optional) - Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. Default is unlimited (0)`,
				},
				resource.Attribute{
					Name:        "stored_vm_quota",
					Description: `(Optional) - Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization. Default is unlimited (0)`,
				},
				resource.Attribute{
					Name:        "can_publish_catalogs",
					Description: `(Optional) - True if this organization is allowed to share catalogs. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_after_power_on_seconds",
					Description: `(Optional) - Specifies this organization's default for virtual machine boot delay after power on. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vapp_lease",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vapp_template_lease",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "maximum_runtime_lease_in_sec",
					Description: `(Required) - How long vApps can run before they are automatically stopped (in seconds). 0 means never expires. Values accepted from 3600+ <br>Note: Default when the whole ` + "`" + `vapp_lease` + "`" + ` block is omitted is 604800 (7 days) but may vary depending on vCD version`,
				},
				resource.Attribute{
					Name:        "power_off_on_runtime_lease_expiration",
					Description: `(Required) - When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires. <br>Note: Default when the whole ` + "`" + `vapp_lease` + "`" + ` block is omitted is false`,
				},
				resource.Attribute{
					Name:        "maximum_storage_lease_in_sec",
					Description: `(Required) - How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires. Regular values accepted from 3600+ <br>Note: Default when the whole ` + "`" + `vapp_lease` + "`" + ` block is omitted is 2592000 (30 days) but may vary depending on vCD version`,
				},
				resource.Attribute{
					Name:        "delete_on_storage_lease_expiration",
					Description: `(Required) - If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted. <br>Note: Default when the whole ` + "`" + `vapp_lease` + "`" + ` block is omitted is false <a id="vapp-template-lease"></a> ## vApp Template Lease The ` + "`" + `vapp_template_lease` + "`" + ` section contains lease parameters for vApp templates created in the current organization, as defined below:`,
				},
				resource.Attribute{
					Name:        "maximum_storage_lease_in_sec",
					Description: `(Required) - How long vApp templates are available before being automatically cleaned up (in seconds). 0 means never expires. Regular values accepted from 3600+ <br>Note: Default when the whole ` + "`" + `vapp_template_lease` + "`" + ` block is omitted is 2592000 (30 days) but may vary depending on vCD version`,
				},
				resource.Attribute{
					Name:        "delete_on_storage_lease_expiration",
					Description: `(Required) - If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted. <br>Note: Default when the whole ` + "`" + `vapp_template_lease` + "`" + ` block is omitted is false ## Importing Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org_group",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Organization group. This can be used to create, update, and delete organization groups defined in SAML or LDAP.`,
			Description:      ``,
			Keywords: []string{
				"org",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the VDC belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of Organization group`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `(Required) Identity provider type for this this group. One of ` + "`" + `SAML` + "`" + ` or ` + "`" + `INTEGRATED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role of the group. Role names can be retrieved from the organization. Both built-in roles and custom built can be used. The roles normally available are:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Organization group ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Organization group ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org_user",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Organization user. This can be used to create, update, and delete organization users.`,
			Description:      ``,
			Keywords: []string{
				"org",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the user belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, but required if ` + "`" + `password_file` + "`" + ` was not given) The user password. This value is never returned on read. It is inspected on create and modify. To modify, fill with a different value. Note that if you remove the password`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `(Optional) Identity provider type for this user. One of: ` + "`" + `INTEGRATED` + "`" + `, ` + "`" + `SAML` + "`" + `, ` + "`" + `OAUTH` + "`" + `. The default is ` + "`" + `INTEGRATED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role of the user. Role names can be retrieved from the organization. Both built-in roles and custom built can be used. The roles normally available are:`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `(Optional) The full name of the user.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the user.`,
				},
				resource.Attribute{
					Name:        "telephone",
					Description: `(Optional) The Org User telephone number.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `(Optional) The Org User email address. Needs to be a properly formatted email address.`,
				},
				resource.Attribute{
					Name:        "instant_messaging",
					Description: `(Optional) The Org User instant messaging.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) True if the user is enabled and can log in. The default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_group_role",
					Description: `(Optional) True if this user has a group role. The default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_locked",
					Description: `(Optional)aIf the user account has been locked due to too many invalid login attempts, the value will change to true (only the system can lock the user). To unlock the user re-set this flag to false.`,
				},
				resource.Attribute{
					Name:        "take_ownership",
					Description: `(Optional) Take ownership of user's objects on deletion.`,
				},
				resource.Attribute{
					Name:        "deployed_vm_quota",
					Description: `(Optional) Quota of vApps that this user can deploy. A value of 0 specifies an unlimited quota. The default is 10.`,
				},
				resource.Attribute{
					Name:        "stored_vm_quota",
					Description: `(Optional) Quota of vApps that this user can store. A value of 0 specifies an unlimited quota. The default is 10. ## Attribute Reference The following attributes are exported on this resource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Organization user ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Organization user ## Importing ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org_vdc",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Organization VDC resource. This can be used to create and delete an Organization VDC.`,
			Description:      ``,
			Keywords: []string{
				"org",
				"vdc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) Organization to create the VDC in, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) VDC name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) VDC friendly description`,
				},
				resource.Attribute{
					Name:        "provider_vdc_name",
					Description: `(Required, System Admin) Name of the Provider VDC from which this organization VDC is provisioned.`,
				},
				resource.Attribute{
					Name:        "allocation_model",
					Description: `(Required) The allocation model used by this VDC; must be one of`,
				},
				resource.Attribute{
					Name:        "compute_capacity",
					Description: `(Required) The compute capacity allocated to this VDC. See [Compute Capacity](#computecapacity) below for details.`,
				},
				resource.Attribute{
					Name:        "nic_quota",
					Description: `(Optional) Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.`,
				},
				resource.Attribute{
					Name:        "network_quota",
					Description: `(Optional) Maximum number of network objects that can be deployed in this VDC. Defaults to 0, which means no networks can be deployed.`,
				},
				resource.Attribute{
					Name:        "vm_quota",
					Description: `(Optional) The maximum number of VMs that can be created in this VDC. Includes deployed and undeployed VMs in vApps and vApp templates. Defaults to 0, which specifies an unlimited number.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) True if this VDC is enabled for use by the organization VDCs. Default is true.`,
				},
				resource.Attribute{
					Name:        "storage_profile",
					Description: `(Required, System Admin) Storage profiles supported by this VDC. See [Storage Profile](#storageprofile) below for details.`,
				},
				resource.Attribute{
					Name:        "memory_guaranteed",
					Description: `(Optional, System Admin) Percentage of allocated memory resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when ` + "`" + `allocation_model` + "`" + ` is AllocationVApp, AllocationPool or Flex. When Allocation model is AllocationPool minimum value is 0.2. If left empty, vCD sets a value.`,
				},
				resource.Attribute{
					Name:        "cpu_guaranteed",
					Description: `(Optional, System Admin) Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when ` + "`" + `allocation_model` + "`" + ` is AllocationVApp, AllocationPool or Flex. If left empty, vCD sets a value.`,
				},
				resource.Attribute{
					Name:        "cpu_speed",
					Description: `(Optional, System Admin) Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM. A VM with 2 vCPUs will consume twice as much of this value. Ignored for ReservationPool. Required when ` + "`" + `allocation_model` + "`" + ` is AllocationVApp, AllocationPool or Flex, and may not be less than 256 MHz. Defaults to 1000 MHz if value isn't provided.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "enable_thin_provisioning",
					Description: `(Optional, System Admin) Boolean to request thin provisioning. Request will be honored only if the underlying data store supports it. Thin provisioning saves storage space by committing it on demand. This allows over-allocation of storage.`,
				},
				resource.Attribute{
					Name:        "enable_fast_provisioning",
					Description: `(Optional, System Admin) Request fast provisioning. Request will be honored only if the underlying datastore supports it. Fast provisioning can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast provisioning, all provisioning operations will result in full clones.`,
				},
				resource.Attribute{
					Name:        "network_pool_name",
					Description: `(Optional, System Admin) Reference to a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.`,
				},
				resource.Attribute{
					Name:        "allow_over_commit",
					Description: `(Optional) Set to false to disallow creation of the VDC if the ` + "`" + `allocation_model` + "`" + ` is AllocationPool or ReservationPool and the ComputeCapacity you specified is greater than what the backing Provider VDC can supply. Default is true.`,
				},
				resource.Attribute{
					Name:        "enable_vm_discovery",
					Description: `(Optional) If true, discovery of vCenter VMs is enabled for resource pools backing this VDC. If false, discovery is disabled. If left unspecified, the actual behaviour depends on enablement at the organization level and at the system level.`,
				},
				resource.Attribute{
					Name:        "elasticity",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "include_vm_memory_overhead",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "delete_force",
					Description: `(Required) When destroying use ` + "`" + `delete_force=True` + "`" + ` to remove a VDC and any objects it contains, regardless of their state.`,
				},
				resource.Attribute{
					Name:        "delete_recursive",
					Description: `(Required) When destroying use ` + "`" + `delete_recursive=True` + "`" + ` to remove the VDC and any objects it contains that are in a state that normally allows removal.`,
				},
				resource.Attribute{
					Name:        "default_vm_sizing_policy_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "vm_sizing_policy_ids",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Provider VDC storage profile.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) True if this storage profile is enabled for use in the VDC. Default is true.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Required) Maximum number of MB allocated for this storage profile. A value of 0 specifies unlimited MB.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) True if this is default storage profile for this VDC. The default storage profile is used when an object that can specify a storage profile is created with no storage profile specified.`,
				},
				resource.Attribute{
					Name:        "storage_used_in_mb",
					Description: `(Computed,`,
				},
				resource.Attribute{
					Name:        "allocated",
					Description: `(Optional) Capacity that is committed to be available. Value in MB or MHz. Used with AllocationPool ("Allocation pool"), ReservationPool ("Reservation pool"), Flex.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Capacity limit relative to the value specified for Allocation. It must not be less than that value. If it is greater than that value, it implies over provisioning. A value of 0 specifies unlimited units. Value in MB or MHz. Used with AllocationVApp ("Pay as you go") or Flex (only for ` + "`" + `cpu` + "`" + `). ## Importing Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director vApp resource. This can be used to create, modify, and delete vApps.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the vApp`,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "power_on",
					Description: `(Optional) A boolean value stating if this vApp should be powered on. Default is ` + "`" + `false` + "`" + `. Works only on update when vApp already has VMs.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Key value map of metadata to assign to this vApp. Key and value can be any string. (Since`,
				},
				resource.Attribute{
					Name:        "guest_properties",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `(Computed) The vApp Hyper Reference`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed;`,
				},
				resource.Attribute{
					Name:        "status_text",
					Description: `(Computed;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_access_control",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Access Control structure for a vApp.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
				"access",
				"control",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the vApp belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of organization to which the vApp belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vapp_id",
					Description: `(Required) A unique identifier for the vApp.`,
				},
				resource.Attribute{
					Name:        "shared_with_everyone",
					Description: `(Required) Whether the vApp is shared with everyone. If any ` + "`" + `shared_with` + "`" + ` blocks are included, this property cannot be used.`,
				},
				resource.Attribute{
					Name:        "everyone_access_level",
					Description: `(Optional) Access level when the vApp is shared with everyone (one of ` + "`" + `ReadOnly` + "`" + `, ` + "`" + `Change` + "`" + `, ` + "`" + `FullControl` + "`" + `). Required if ` + "`" + `shared_with_everyone` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "shared_with",
					Description: `(Optional) one or more blocks defining a subject to which we are sharing. See [shared_with](#shared_with) below for detail. It cannot be used if ` + "`" + `shared_with_everyone` + "`" + ` is set. ## shared_with`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional) The ID of a user with which we are sharing. Required if ` + "`" + `group_id` + "`" + ` is not set.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) The ID of a group with which we are sharing. Required if ` + "`" + `user_id` + "`" + ` is not set.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Required) The access level for the user or group to which we are sharing. (One of ` + "`" + `ReadOnly` + "`" + `, ` + "`" + `Change` + "`" + `, ` + "`" + `FullControl` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "subject_name",
					Description: `(Computed) the name of the subject (group or user) with which we are sharing. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_firewall_rules",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director vApp Firewall resource. This can be used to create, modify, and delete firewall settings and rules.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
				"firewall",
				"rules",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vapp_id",
					Description: `(Required) The identifier of [vApp](/docs/providers/vcd/r/vapp.html).`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The identifier of [vApp network](/docs/providers/vcd/r/vapp_network.html).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable or disable firewall. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(Required) Either 'allow' or 'drop'. Specifies what to do should none of the rules match.`,
				},
				resource.Attribute{
					Name:        "log_default_action",
					Description: `(Optional) Flag to enable logging for default action. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Configures a firewall rule; see [Rules](#rules) below for details. <a id="rules"></a> ## Rules Each firewall rule supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the firewall rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) ` + "`" + `true` + "`" + ` value will enable firewall rule. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) Specifies what to do when this rule is matched. Either ` + "`" + `allow` + "`" + ` or ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol to match. One of ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `any` + "`" + ` or ` + "`" + `tcp&udp` + "`" + `. Default is ` + "`" + `any` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Optional) The destination port to match. Either a port number or ` + "`" + `any` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_ip",
					Description: `(Optional) The destination IP to match. Either an IP address, IP range or ` + "`" + `any` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_vm_id",
					Description: `(Optional) Destination VM identifier.`,
				},
				resource.Attribute{
					Name:        "destination_vm_ip_type",
					Description: `(Optional) The value can be one of: ` + "`" + `assigned` + "`" + ` - assigned internal IP will be automatically chosen, ` + "`" + `NAT` + "`" + ` - NATed external IP will be automatically chosen.`,
				},
				resource.Attribute{
					Name:        "destination_vm_nic_id",
					Description: `(Optional) VM NIC ID to which this rule applies.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) The source port to match. Either a port number or ` + "`" + `any` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `(Optional) The source IP to match. Either an IP address, IP range or ` + "`" + `any` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_vm_id",
					Description: `(Optional) Source VM identifier.`,
				},
				resource.Attribute{
					Name:        "source_vm_ip_type",
					Description: `(Optional) The value can be one of: ` + "`" + `assigned` + "`" + ` - assigned internal IP will be automatically chosen, ` + "`" + `NAT` + "`" + ` - NATed external IP will be automatically chosen.`,
				},
				resource.Attribute{
					Name:        "source_vm_nic_id",
					Description: `(Optional) VM NIC ID to which this rule applies.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_nat_rules",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director vApp NAT resource. This can be used to create, modify, and delete NAT rules.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
				"nat",
				"rules",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vapp_id",
					Description: `(Required) The identifier of [vApp](/docs/providers/vcd/r/vapp.html).`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The identifier of [vApp network](/docs/providers/vcd/r/vapp_network.html).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable or disable NAT. Default is ` + "`" + `true` + "`" + `. To enable the NAT service, [vcd_vapp_firewall_rules](/docs/providers/vcd/r/vapp_firewall_rules.html) needs to be enabled as well.`,
				},
				resource.Attribute{
					Name:        "nat_type",
					Description: `(Required) "One of: ` + "`" + `ipTranslation` + "`" + ` (use IP translation), ` + "`" + `portForwarding` + "`" + ` (use port forwarding). For ` + "`" + `ipTranslation` + "`" + ` fields ` + "`" + `vm_id` + "`" + `, ` + "`" + `vm_nic_id` + "`" + `, ` + "`" + `mapping_mode` + "`" + ` are required and ` + "`" + `external_ip` + "`" + ` is optional. For ` + "`" + `portForwarding` + "`" + ` fields ` + "`" + `vm_id` + "`" + `, ` + "`" + `vm_nic_id` + "`" + `, ` + "`" + `protocol` + "`" + `, ` + "`" + `external_port` + "`" + ` and ` + "`" + `forward_to_port` + "`" + ` are required.`,
				},
				resource.Attribute{
					Name:        "enable_ip_masquerade",
					Description: `(Optional) When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Configures a NAT rule; see [Rules](#rules) below for details. <a id="rules"></a> ## Rules Each NAT rule supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "mapping_mode",
					Description: `(Optional) Mapping mode. One of: ` + "`" + `automatic` + "`" + `, ` + "`" + `manual` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Required) VM to which this rule applies.`,
				},
				resource.Attribute{
					Name:        "vm_nic_id",
					Description: `(Required) VM NIC ID to which this rule applies.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `(Optional) External IP address to forward to or External IP address to map to VM.`,
				},
				resource.Attribute{
					Name:        "external_port",
					Description: `(Optional) External port to forward. ` + "`" + `-1` + "`" + ` value for any port.`,
				},
				resource.Attribute{
					Name:        "forward_to_port",
					Description: `(Optional) Internal port to forward. ` + "`" + `-1` + "`" + ` value for any port.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol to forward. One of: ` + "`" + `TCP` + "`" + ` (forward TCP packets), ` + "`" + `UDP` + "`" + ` (forward UDP packets), ` + "`" + `TCP_UDP` + "`" + ` (forward TCP and UDP packets). ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_network",
			Category:         "Resources",
			ShortDescription: `Allows to provision a vApp network and optionally connect it to an existing Org VDC network.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) The vApp this network belongs to.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) The netmask for the new network. Default is ` + "`" + `255.255.255.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) First DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Second DNS server to use.`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) A FQDN for the virtual machines on this network.`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Optional) A range of IPs permitted to be used as static IPs for virtual machines; see [IP Pools](#ip-pools) below for details.`,
				},
				resource.Attribute{
					Name:        "dhcp_pool",
					Description: `(Optional) A range of IPs to issue to virtual machines that don't have a static IP; see [IP Pools](#ip-pools) below for details.`,
				},
				resource.Attribute{
					Name:        "org_network_name",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "retain_ip_mac_enabled",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) The first address in the IP Range.`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) The final address in the IP Range. DHCP Pools additionally support the following attributes:`,
				},
				resource.Attribute{
					Name:        "default_lease_time",
					Description: `(Optional) The default DHCP lease time to use. Defaults to ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_lease_time",
					Description: `(Optional) The maximum DHCP lease time to use. Defaults to ` + "`" + `7200` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Allows to enable or disable service. Default is true. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_org_network",
			Category:         "Resources",
			ShortDescription: `Provides capability to attach an existing Org VDC Network to a vApp and toggle network features.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
				"org",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) The vApp this network belongs to.`,
				},
				resource.Attribute{
					Name:        "org_network_name",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "retain_ip_mac_enabled",
					Description: `(Optional) Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Configurable when ` + "`" + `is_fenced` + "`" + ` is true. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_static_routing",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director vApp static routing resource. This can be used to create, modify, and delete static routing rules.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
				"static",
				"routing",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vapp_id",
					Description: `(Required) The identifier of [vApp](/docs/providers/vcd/r/vapp.html).`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The identifier of [vApp network](/docs/providers/vcd/r/vapp_network.html).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable or disable static Routing. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Configures a static routing rule; see [Rules](#rules) below for details. <a id="rules"></a> ## Rules Each static routing rule supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the static route.`,
				},
				resource.Attribute{
					Name:        "network_cidr",
					Description: `(Required) Network specification in CIDR.`,
				},
				resource.Attribute{
					Name:        "next_hop_ip",
					Description: `(Required) IP address of Next Hop router/gateway. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_vm",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director VM resource. This can be used to create, modify, and delete VMs within a vApp.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
				"vm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) The vApp this VM belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the VM, unique within the vApp`,
				},
				resource.Attribute{
					Name:        "computer_name",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vm_name_in_template",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) The amount of RAM (in MB) to allocate to the VM. If ` + "`" + `memory_hot_add_enabled` + "`" + ` is true, then memory will be increased without VM power off.`,
				},
				resource.Attribute{
					Name:        "cpus",
					Description: `(Optional) The number of virtual CPUs to allocate to the VM. Socket count is a result of: virtual logical processors/cores per socket. If ` + "`" + `cpu_hot_add_enabled` + "`" + ` is true, then cpus will be increased without VM power off.`,
				},
				resource.Attribute{
					Name:        "cpu_cores",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "power_on",
					Description: `(Optional) A boolean value stating if this VM should be powered on. Default is ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "accept_all_eulas",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "expose_hardware_virtualization",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "customization",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "guest_properties",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "override_template_disk",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "network_dhcp_wait_seconds",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "hardware_version",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "boot_image",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "cpu_hot_add_enabled",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "memory_hot_add_enabled",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "prevent_update_power_off",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Independent disk name`,
				},
				resource.Attribute{
					Name:        "bus_number",
					Description: `(Required) Bus number on which to place the disk controller`,
				},
				resource.Attribute{
					Name:        "unit_number",
					Description: `(Required) Unit number (slot) on the bus specified by BusNumber. <a id="network-block"></a> ## Network`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Computed) Mac address of network interface.`,
				},
				resource.Attribute{
					Name:        "adapter_type",
					Description: `(Optional, Computed) Adapter type (names are case insensitive). Some known adapter types - ` + "`" + `VMXNET3` + "`" + `, ` + "`" + `E1000` + "`" + `, ` + "`" + `E1000E` + "`" + `, ` + "`" + `SRIOVETHERNETCARD` + "`" + `, ` + "`" + `VMXNET2` + "`" + `, ` + "`" + `PCNet32` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "POOL",
					Description: `Static IP address is allocated automatically from defined static pool in network.`,
				},
				resource.Attribute{
					Name:        "DHCP",
					Description: `IP address is obtained from a DHCP service. Field ` + "`" + `ip` + "`" + ` is not guaranteed to be populated. Because of this it may appear after multiple ` + "`" + `terraform refresh` + "`" + ` operations.`,
				},
				resource.Attribute{
					Name:        "MANUAL",
					Description: `IP address is assigned manually in the ` + "`" + `ip` + "`" + ` field. Must be valid IP address from static pool.`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `No IP address will be set because VM will have a NIC without network.`,
				},
				resource.Attribute{
					Name:        "connected",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "bus_type",
					Description: `(Required) The type of disk controller. Possible values: ` + "`" + `ide` + "`" + `, ` + "`" + `parallel` + "`" + `( LSI Logic Parallel SCSI), ` + "`" + `sas` + "`" + `(LSI Logic SAS (SCSI)), ` + "`" + `paravirtual` + "`" + `(Paravirtual (SCSI)), ` + "`" + `sata` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size_in_mb",
					Description: `(Required) The size of the disk in MB.`,
				},
				resource.Attribute{
					Name:        "bus_number",
					Description: `(Required) The number of the SCSI or IDE controller itself.`,
				},
				resource.Attribute{
					Name:        "unit_number",
					Description: `(Required) The device number on the SCSI or IDE controller of the disk.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) Specifies the IOPS for the disk. Default is 0.`,
				},
				resource.Attribute{
					Name:        "storage_profile",
					Description: `(Optional) Storage profile which overrides the VM default one. <a id="customization-block"></a> ## Customization When you customize your guest OS you can set up a virtual machine with the operating system that you want. vCloud Director can customize the network settings of the guest operating system of a virtual machine created from a vApp template. When you customize your guest operating system, you can create and deploy multiple unique virtual machines based on the same vApp template without machine name or network conflicts. When you configure a vApp template with the prerequisites for guest customization and add a virtual machine to a vApp based on that template, vCloud Director creates a package with guest customization tools. When you deploy and power on the virtual machine for the first time, vCloud Director copies the package, runs the tools, and deletes the package from the virtual machine. ~>`,
				},
				resource.Attribute{
					Name:        "internal_disk",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "disk.size_in_mb",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "bus_type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "size_in_mb",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "bus_number",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "unit_number",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "thin_provisioned",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "storage_profile",
					Description: `(`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "internal_disk",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "disk.size_in_mb",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "bus_type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "size_in_mb",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "bus_number",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "unit_number",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "thin_provisioned",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "storage_profile",
					Description: `(`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm",
			Category:         "Resources",
			ShortDescription: `Provides a VMware Cloud Director standalone VM resource. This can be used to create, modify, and delete Standalone VMs.`,
			Description:      ``,
			Keywords: []string{
				"vm",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm_affinity_rule",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director VM affinity rule resource. This can be used to create, modify, and delete VM affinity and anti-affinity rules.`,
			Description:      ``,
			Keywords: []string{
				"vm",
				"affinity",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of VM affinity rule. Duplicates are allowed, although the name can be used to retrieve the rule (as data source or when importing) only if it is unique.`,
				},
				resource.Attribute{
					Name:        "polarity",
					Description: `(Required) One of ` + "`" + `Affinity` + "`" + ` or ` + "`" + `Anti-Affinity` + "`" + `. This property cannot be changed. Once created, if we need to change polarity, we need to remove the rule and create a new one.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm_internal_disk",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director VM internal disk resource. This can be used to create and delete VM internal disks.`,
			Description:      ``,
			Keywords: []string{
				"vm",
				"internal",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) The vAPP this VM internal disk belongs to.`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Required) VM in vAPP in which internal disk is created.`,
				},
				resource.Attribute{
					Name:        "allow_vm_reboot",
					Description: `(Optional) Powers off VM when changing any attribute of an IDE disk or unit/bus number of other disk types, after the change is complete VM is powered back on. Without this setting enabled, such changes on a powered-on VM would fail. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "bus_type",
					Description: `(Required) The type of disk controller. Possible values: ` + "`" + `ide` + "`" + `, ` + "`" + `parallel` + "`" + `( LSI Logic Parallel SCSI), ` + "`" + `sas` + "`" + `(LSI Logic SAS (SCSI)), ` + "`" + `paravirtual` + "`" + `(Paravirtual (SCSI)), ` + "`" + `sata` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size_in_mb",
					Description: `(Required) The size of the disk in MB.`,
				},
				resource.Attribute{
					Name:        "bus_number",
					Description: `(Required) The number of the SCSI or IDE controller itself.`,
				},
				resource.Attribute{
					Name:        "unit_number",
					Description: `(Required) The device number on the SCSI or IDE controller of the disk.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) Specifies the IOPS for the disk. Default is 0.`,
				},
				resource.Attribute{
					Name:        "storage_profile",
					Description: `(Optional) Storage profile which overrides the VM default one. ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "thin_provisioned",
					Description: `Specifies whether the disk storage is pre-allocated or allocated on demand. ## Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm_sizing_policy",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director VM sizing policy resource. This can be used to create, modify, and delete VM sizing policy.`,
			Description:      ``,
			Keywords: []string{
				"vm",
				"sizing",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of VM sizing policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description of VM sizing policy.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional) Configures cpu policy; see [Cpu](#cpu) below for details.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Configures memory policy; see [Memory](#memory) below for details. <a id="cpu"></a> ## CPU Each VM sizing policy supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "shares",
					Description: `(Optional) Defines the number of CPU shares for a VM. Shares specify the relative importance of a VM within a virtual data center. If a VM has twice as many shares of CPU as another VM, it is entitled to consume twice as much CPU when these two virtual machines are competing for resources. If not defined in the VDC compute policy, normal shares are applied to the VM.`,
				},
				resource.Attribute{
					Name:        "limit_in_mhz",
					Description: `(Optional) Defines the CPU limit in MHz for a VM. If not defined in the VDC compute policy, CPU limit is equal to the vCPU speed multiplied by the number of vCPUs.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Required) Defines the number of vCPUs configured for a VM. This is a VM hardware configuration. When a tenant assigns the VM sizing policy to a VM, this count becomes the configured number of vCPUs for the VM.`,
				},
				resource.Attribute{
					Name:        "speed_in_mhz",
					Description: `(Optional) Defines the vCPU speed of a core in MHz.`,
				},
				resource.Attribute{
					Name:        "cores_per_socket",
					Description: `(Optional) The number of cores per socket for a VM. This is a VM hardware configuration. The number of vCPUs that is defined in the VM sizing policy must be divisible by the number of cores per socket. If the number of vCPUs is not divisible by the number of cores per socket, the number of cores per socket becomes invalid.`,
				},
				resource.Attribute{
					Name:        "reservation_guarantee",
					Description: `(Optional) Defines how much of the CPU resources of a VM are reserved. The allocated CPU for a VM equals the number of vCPUs times the vCPU speed in MHz. The value of the attribute ranges between 0 and one. Value of 0 CPU reservation guarantee defines no CPU reservation. Value of 1 defines 100% of CPU reserved. <a id="memory"></a> ## Memory Each VM sizing policy supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "shares",
					Description: `(Optional) Defines the number of memory shares for a VM. Shares specify the relative importance of a VM within a virtual data center. If a VM has twice as many shares of memory as another VM, it is entitled to consume twice as much memory when these two virtual machines are competing for resources. If not defined in the VDC compute policy, normal shares are applied to the VM.`,
				},
				resource.Attribute{
					Name:        "size_in_mb",
					Description: `(Optional) Defines the memory configured for a VM in MB. This is a VM hardware configuration. When a tenant assigns the VM sizing policy to a VM, the VM receives the amount of memory defined by this attribute.`,
				},
				resource.Attribute{
					Name:        "limit_in_mb",
					Description: `(Optional) Defines the memory limit in MB for a VM. If not defined in the VM sizing policy, memory limit is equal to the allocated memory for the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_guarantee",
					Description: `(Optional) Defines the reserved amount of memory that is configured for a VM. The value of the attribute ranges between 0 and one. Value of 0 memory reservation guarantee defines no memory reservation. Value of 1 defines 100% of memory reserved. # Importing ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"vcd_catalog":               0,
		"vcd_catalog_item":          1,
		"vcd_catalog_media":         2,
		"vcd_edgegateway":           3,
		"vcd_edgegateway_settings":  4,
		"vcd_edgegateway_vpn":       5,
		"vcd_external_network":      6,
		"vcd_external_network_v2":   7,
		"vcd_independent_disk":      8,
		"vcd_inserted_media":        9,
		"vcd_nsxv_ip_set":           10,
		"vcd_lb_app_profile":        11,
		"vcd_lb_app_rule":           12,
		"vcd_lb_server_pool":        13,
		"vcd_lb_service_monitor":    14,
		"vcd_lb_virtual_server":     15,
		"vcd_network_direct":        16,
		"vcd_network_isolated":      17,
		"vcd_network_isolated_v2":   18,
		"vcd_network_routed":        19,
		"vcd_network_routed_v2":     20,
		"vcd_nsxt_edgegateway":      21,
		"vcd_nsxt_network_dhcp":     22,
		"vcd_nsxt_network_imported": 23,
		"vcd_nsxv_dhcp_relay":       24,
		"vcd_nsxv_dnat":             25,
		"vcd_nsxv_firewall_rule":    26,
		"vcd_nsxv_snat":             27,
		"vcd_org":                   28,
		"vcd_org_group":             29,
		"vcd_org_user":              30,
		"vcd_org_vdc":               31,
		"vcd_vapp":                  32,
		"vcd_vapp_access_control":   33,
		"vcd_vapp_firewall_rules":   34,
		"vcd_vapp_nat_rules":        35,
		"vcd_vapp_network":          36,
		"vcd_vapp_org_network":      37,
		"vcd_vapp_static_routing":   38,
		"vcd_vapp_vm":               39,
		"vcd_vm":                    40,
		"vcd_vm_affinity_rule":      41,
		"vcd_vm_internal_disk":      42,
		"vcd_vm_sizing_policy":      43,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
