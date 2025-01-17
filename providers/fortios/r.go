package fortios

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_address",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall addresses used in firewall policies of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of address(Support ipmask, iprange, fqdn and geography).`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP address and subnet mask of address.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `IP addresses associated to a specific country.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "associated_interface",
					Description: `Network interface associated with address.`,
				},
				resource.Attribute{
					Name:        "show_in_address_list",
					Description: `Enable/disable address visibility in the GUI. default is "enable".`,
				},
				resource.Attribute{
					Name:        "static_route_configure",
					Description: `Enable/disable use of this address in the static route configuration. default is "disable". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the address item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of address(Support ipmask, iprange, fqdn and geography).`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP address and subnet mask of address.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `IP addresses associated to a specific country.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "associated_interface",
					Description: `Network interface associated with address.`,
				},
				resource.Attribute{
					Name:        "show_in_address_list",
					Description: `Enable/disable address visibility in the GUI. default is "enable".`,
				},
				resource.Attribute{
					Name:        "static_route_configure",
					Description: `Enable/disable use of this address in the static route configuration. default is "disable".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the address item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of address(Support ipmask, iprange, fqdn and geography).`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP address and subnet mask of address.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `IP addresses associated to a specific country.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "associated_interface",
					Description: `Network interface associated with address.`,
				},
				resource.Attribute{
					Name:        "show_in_address_list",
					Description: `Enable/disable address visibility in the GUI. default is "enable".`,
				},
				resource.Attribute{
					Name:        "static_route_configure",
					Description: `Enable/disable use of this address in the static route configuration. default is "disable".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_addressgroup",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall address group used in firewall policies of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"addressgroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Address group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Address objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall address group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Address objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall address group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Address objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_ippool",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure IPv4 IP address pools of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"ippool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) IP pool name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) IP pool type(Support overload and one-to-one).`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `(Required) First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `(Required) Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP requests when an IP Pool is added to a policy.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP pool item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IP pool name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP pool type(Support overload and one-to-one).`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP requests when an IP Pool is added to a policy.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP pool item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IP pool name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP pool type(Support overload and one-to-one).`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP requests when an IP Pool is added to a policy.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_service",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall service of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Service category.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type based on IANA numbers.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `Multiple TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `Multiple UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `Multiple SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `ICMP type.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `ICMP code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "session_ttl",
					Description: `Custom tcp session TTL.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Service category.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type based on IANA numbers.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `Multiple TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `Multiple UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `Multiple SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `ICMP type.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `ICMP code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "session_ttl",
					Description: `Custom tcp session TTL.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Service category.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type based on IANA numbers.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `Multiple TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `Multiple UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `Multiple SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `ICMP type.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `ICMP code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "session_ttl",
					Description: `Custom tcp session TTL.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_servicecategory",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall service categories of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"servicecategory",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Category name.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service item.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service item.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_servicegroup",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall service group of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"servicegroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Service group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Service objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Service group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Service objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Service group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Service objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_vip",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall virtual IPs (VIPs) of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"vip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "extip",
					Description: `(Required) IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedip",
					Description: `(Required) IP address or address range on the destination network to which the external IP address is mapped.`,
				},
				resource.Attribute{
					Name:        "extintf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "portforward",
					Description: `Enable/disable port forwarding.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol to use when forwarding packets.`,
				},
				resource.Attribute{
					Name:        "extport",
					Description: `Incoming port number range that you want to map to a port number range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedport",
					Description: `Port number range on the destination network to which the external port number range is mapped.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall virtual IPs item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "extip",
					Description: `IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedip",
					Description: `IP address or address range on the destination network to which the external IP address is mapped.`,
				},
				resource.Attribute{
					Name:        "extintf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "portforward",
					Description: `Enable/disable port forwarding.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol to use when forwarding packets.`,
				},
				resource.Attribute{
					Name:        "extport",
					Description: `Incoming port number range that you want to map to a port number range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedport",
					Description: `Port number range on the destination network to which the external port number range is mapped.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall virtual IPs item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "extip",
					Description: `IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedip",
					Description: `IP address or address range on the destination network to which the external IP address is mapped.`,
				},
				resource.Attribute{
					Name:        "extintf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "portforward",
					Description: `Enable/disable port forwarding.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol to use when forwarding packets.`,
				},
				resource.Attribute{
					Name:        "extport",
					Description: `Incoming port number range that you want to map to a port number range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedport",
					Description: `Port number range on the destination network to which the external port number range is mapped.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_vipgroup",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure virtual IP groups of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"vipgroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) VIP group name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Member VIP objects of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual IP groups item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VIP group name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Member VIP objects of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual IP groups item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VIP group name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Member VIP objects of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_security_policy",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall policies of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"security",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy name.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `(Required) Incoming (ingress) interface.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `(Required) Outgoing (egress) interface.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `(Required) Source address and address group names.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `(Required) Destination address and address group names.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Required) Schedule name.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) Service and service group names..`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging. Log all sessions or security profile sessions.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `Device type category.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ssl_ssh_profile",
					Description: `Name of an existing SSL SSH profile.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Internet Service source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable or disable this policy.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall policy item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `Incoming (ingress) interface.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `Outgoing (egress) interface.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `Source address and address group names.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `Destination address and address group names.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Schedule name.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service and service group names..`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging. Log all sessions or security profile sessions.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `Device type category.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ssl_ssh_profile",
					Description: `Name of an existing SSL SSH profile.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Internet Service source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable or disable this policy.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall policy item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `Incoming (ingress) interface.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `Outgoing (egress) interface.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `Source address and address group names.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `Destination address and address group names.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Schedule name.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service and service group names..`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging. Log all sessions or security profile sessions.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `Device type category.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ssl_ssh_profile",
					Description: `Name of an existing SSL SSH profile.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Internet Service source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable or disable this policy.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_security_policyseq",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to alter firewall security policy sequence`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"security",
				"policyseq",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_src_id",
					Description: `(Required) The policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "policy_dst_id",
					Description: `(Required) The dest policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "alter_position",
					Description: `(Required) The alter position: should only be "before" or "after"`,
				},
				resource.Attribute{
					Name:        "enable_state_checking",
					Description: `Enable status detection for policy_src_id and policy_dst_id, to detect whether they have been changed outside of terraform, the default is false. If this parameter is true, when policy_src_id and policy_dst_id are modified or deleted outside the terraform, the resource will be able to detect the change, in other words, terraform plan can detect the change. When this parameter is false, if policy_src_id and policy_dst_id are modified or deleted outside the terraform, the change will not be detected by the resource, which means that the terraform plan will think that the state has not changed. The main purpose of this argument is to make the resource compatible with the old version. If you use this resource for the first time now, it is recommended to set it to true. For detailed usage, see "Attributes Reference" and "A More Detailed Example" below.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "policy_src_id",
					Description: `The policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "policy_dst_id",
					Description: `The dest policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "alter_position",
					Description: `The alter position: should only be "before" or "after"`,
				},
				resource.Attribute{
					Name:        "enable_state_checking",
					Description: `Enable status detection for policy_src_id and policy_dst_id`,
				},
				resource.Attribute{
					Name:        "state_policy_srcdst_pos",
					Description: `The parameter is read-only, it is used to get the lastest relative position of the policy with policy_src_id and the policy with policy_dst_id when enable_state_checking is set to true. This can help check whether the latest relative position of the two plicies matches the configuration, and help check whether they have been deleted. This is generally used in the following situations: These two policies are deleted or moved outside of Terraform. Terraform plan will determine the consistency of the state based on this attribute. It includs the following states:`,
				},
				resource.Attribute{
					Name:        "state_policy_list",
					Description: `The parameter is read-only, it is used to get the latest policy list(by sequence) for reference when enable_state_checking is set to true, in the list, the policy with policy_src_id and the policy with policy_dst_id will be marked with`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment ## A More Detailed Example: 1 Let us assume that there are the following existing policies (by sequence) ` + "`" + `` + "`" + `` + "`" + `hcl { action = "accept" name = "e234552" policyid = "6" }, { action = "accept" name = "fdsafcew222" policyid = "7" }, { action = "accept" name = "fdscer435" policyid = "4" }, { action = "accept" name = "22" policyid = "5" }, { action = "accept" name = "e3232" policyid = "8" }, ` + "`" + `` + "`" + `` + "`" + ` 2 We are going to move 7 to after 5 ` + "`" + `` + "`" + `` + "`" + `hcl resource "fortios_firewall_security_policyseq" "test1" { policy_src_id = 7 policy_dst_id = 5 alter_position = "after" enable_state_checking = true } ` + "`" + `` + "`" + `` + "`" + ` 3 After executing terraform apply, we execute terraform show ` + "`" + `` + "`" + `` + "`" + `hcl # terraform show 2020/09/14 01:49:32 [WARN] Log levels other than TRACE are currently unreliable, and are supported only for backward compatibility. Use TF_LOG=TRACE to see Terraform's internal logs. ---- # fortios_firewall_security_policyseq.test1: resource "fortios_firewall_security_policyseq" "test1" { alter_position = "after" enable_state_checking = true id = "7" policy_dst_id = 5 policy_src_id = 7 state_policy_list = [ { action = "accept" name = "e234552" policyid = "6" }, { action = "accept" name = "fdscer435" policyid = "4" }, { action = "accept" name = "22" policyid = "`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "policy_src_id",
					Description: `The policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "policy_dst_id",
					Description: `The dest policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "alter_position",
					Description: `The alter position: should only be "before" or "after"`,
				},
				resource.Attribute{
					Name:        "enable_state_checking",
					Description: `Enable status detection for policy_src_id and policy_dst_id`,
				},
				resource.Attribute{
					Name:        "state_policy_srcdst_pos",
					Description: `The parameter is read-only, it is used to get the lastest relative position of the policy with policy_src_id and the policy with policy_dst_id when enable_state_checking is set to true. This can help check whether the latest relative position of the two plicies matches the configuration, and help check whether they have been deleted. This is generally used in the following situations: These two policies are deleted or moved outside of Terraform. Terraform plan will determine the consistency of the state based on this attribute. It includs the following states:`,
				},
				resource.Attribute{
					Name:        "state_policy_list",
					Description: `The parameter is read-only, it is used to get the latest policy list(by sequence) for reference when enable_state_checking is set to true, in the list, the policy with policy_src_id and the policy with policy_dst_id will be marked with`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment ## A More Detailed Example: 1 Let us assume that there are the following existing policies (by sequence) ` + "`" + `` + "`" + `` + "`" + `hcl { action = "accept" name = "e234552" policyid = "6" }, { action = "accept" name = "fdsafcew222" policyid = "7" }, { action = "accept" name = "fdscer435" policyid = "4" }, { action = "accept" name = "22" policyid = "5" }, { action = "accept" name = "e3232" policyid = "8" }, ` + "`" + `` + "`" + `` + "`" + ` 2 We are going to move 7 to after 5 ` + "`" + `` + "`" + `` + "`" + `hcl resource "fortios_firewall_security_policyseq" "test1" { policy_src_id = 7 policy_dst_id = 5 alter_position = "after" enable_state_checking = true } ` + "`" + `` + "`" + `` + "`" + ` 3 After executing terraform apply, we execute terraform show ` + "`" + `` + "`" + `` + "`" + `hcl # terraform show 2020/09/14 01:49:32 [WARN] Log levels other than TRACE are currently unreliable, and are supported only for backward compatibility. Use TF_LOG=TRACE to see Terraform's internal logs. ---- # fortios_firewall_security_policyseq.test1: resource "fortios_firewall_security_policyseq" "test1" { alter_position = "after" enable_state_checking = true id = "7" policy_dst_id = 5 policy_src_id = 7 state_policy_list = [ { action = "accept" name = "e234552" policyid = "6" }, { action = "accept" name = "fdscer435" policyid = "4" }, { action = "accept" name = "22" policyid = "`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_log_fortianalyzer_setting",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure configure logging to FortiAnalyzer log management devices.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"log",
				"fortianalyzer",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Enable/disable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The remote FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "upload_option",
					Description: `Enable/disable logging to hard disk and then uploading to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "hmac_algorithm",
					Description: `FortiAnalyzer IPsec tunnel HMAC algorithm.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `Enable/disable sending FortiAnalyzer log data with SSL encryption. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The remote FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "upload_option",
					Description: `Enable/disable logging to hard disk and then uploading to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "hmac_algorithm",
					Description: `FortiAnalyzer IPsec tunnel HMAC algorithm.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `Enable/disable sending FortiAnalyzer log data with SSL encryption.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The remote FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "upload_option",
					Description: `Enable/disable logging to hard disk and then uploading to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "hmac_algorithm",
					Description: `FortiAnalyzer IPsec tunnel HMAC algorithm.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `Enable/disable sending FortiAnalyzer log data with SSL encryption.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_log_syslog_setting",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure logging to remote Syslog logging servers.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"log",
				"syslog",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Enable/disable remote syslog logging.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Address of remote syslog server.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Remote syslog logging over UDP/Reliable TCP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server listen port.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP address of syslog.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Log format. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable remote syslog logging.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Address of remote syslog server.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Remote syslog logging over UDP/Reliable TCP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server listen port.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP address of syslog.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Log format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable remote syslog logging.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Address of remote syslog server.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Remote syslog logging over UDP/Reliable TCP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server listen port.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP address of syslog.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Log format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_networking_interface_port",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure interface settings of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"networking",
				"interface",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) If the interface is physical, the argument is the name of the interface.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Interface type (support physical, vlan, loopback).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface IPv4 address and subnet mask, syntax` + "`" + ` - X.X.X.X X.X.X.X.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Alias will be displayed with the interface name to make it easier to distinguish.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Bring the interface up or shut the interface down.`,
				},
				resource.Attribute{
					Name:        "device_identification",
					Description: `Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `TCP maximum segment size. 0 means do not change segment size.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Interface speed. The default setting and the options available depend on the interface hardware.`,
				},
				resource.Attribute{
					Name:        "mtu_override",
					Description: `Enable to set a custom MTU for this interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU value for this interface.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Interface role.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Permitted types of management access to this interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) Addressing mode.`,
				},
				resource.Attribute{
					Name:        "dns_server_override",
					Description: `Enable/disable use DNS acquired by DHCP or PPPoE.`,
				},
				resource.Attribute{
					Name:        "defaultgw",
					Description: `Enable to get the gateway IP from the DHCP or PPPoE server.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Interface is in this virtual domain (VDOM).`,
				},
				resource.Attribute{
					Name:        "vlanid",
					Description: `VLAN ID. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Name of the interface.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface IPv4 address and subnet mask, syntax` + "`" + ` - X.X.X.X X.X.X.X.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Alias will be displayed with the interface name to make it easier to distinguish.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Bring the interface up or shut the interface down.`,
				},
				resource.Attribute{
					Name:        "device_identification",
					Description: `Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `TCP maximum segment size. 0 means do not change segment size.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Interface speed. The default setting and the options available depend on the interface hardware.`,
				},
				resource.Attribute{
					Name:        "mtu_override",
					Description: `Enable to set a custom MTU for this interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU value for this interface.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Interface role.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Permitted types of management access to this interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Addressing mode.`,
				},
				resource.Attribute{
					Name:        "dns_server_override",
					Description: `Enable/disable use DNS acquired by DHCP or PPPoE.`,
				},
				resource.Attribute{
					Name:        "defaultgw",
					Description: `Enable to get the gateway IP from the DHCP or PPPoE server.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Interface type (support physical, vlan, loopback).`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Interface is in this virtual domain (VDOM).`,
				},
				resource.Attribute{
					Name:        "vlanid",
					Description: `VLAN ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Name of the interface.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface IPv4 address and subnet mask, syntax` + "`" + ` - X.X.X.X X.X.X.X.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Alias will be displayed with the interface name to make it easier to distinguish.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Bring the interface up or shut the interface down.`,
				},
				resource.Attribute{
					Name:        "device_identification",
					Description: `Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `TCP maximum segment size. 0 means do not change segment size.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Interface speed. The default setting and the options available depend on the interface hardware.`,
				},
				resource.Attribute{
					Name:        "mtu_override",
					Description: `Enable to set a custom MTU for this interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU value for this interface.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Interface role.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Permitted types of management access to this interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Addressing mode.`,
				},
				resource.Attribute{
					Name:        "dns_server_override",
					Description: `Enable/disable use DNS acquired by DHCP or PPPoE.`,
				},
				resource.Attribute{
					Name:        "defaultgw",
					Description: `Enable to get the gateway IP from the DHCP or PPPoE server.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Interface type (support physical, vlan, loopback).`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Interface is in this virtual domain (VDOM).`,
				},
				resource.Attribute{
					Name:        "vlanid",
					Description: `VLAN ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_networking_route_static",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure static route of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"networking",
				"route",
				"static",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dst",
					Description: `(Required) Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required) Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "blackhole",
					Description: `Enable/disable black hole.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Administrative distance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Administrative weight.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Administrative priority.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Required) Gateway out interface or tunnel.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Optional comments.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Application ID in the Internet service database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable this static route. default is "enable". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the static route item.`,
				},
				resource.Attribute{
					Name:        "dst",
					Description: `Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "blackhole",
					Description: `Enable/disable black hole.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Administrative distance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Administrative weight.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Administrative priority.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Gateway out interface or tunnel.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Optional comments.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Application ID in the Internet service database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable this static route.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the static route item.`,
				},
				resource.Attribute{
					Name:        "dst",
					Description: `Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "blackhole",
					Description: `Enable/disable black hole.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Administrative distance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Administrative weight.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Administrative priority.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Gateway out interface or tunnel.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Optional comments.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Application ID in the Internet service database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable this static route.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_admin_administrator",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure administrator accounts of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"admin",
				"administrator",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Admin user password.`,
				},
				resource.Attribute{
					Name:        "trusthostN",
					Description: `Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domain(s) that the administrator can access.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `(Required) Access profile for this administrator. Access profiles control administrator access to FortiGate features.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the administrator account item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Admin user password.`,
				},
				resource.Attribute{
					Name:        "trusthostN",
					Description: `Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domain(s) that the administrator can access.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Access profile for this administrator. Access profiles control administrator access to FortiGate features.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the administrator account item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Admin user password.`,
				},
				resource.Attribute{
					Name:        "trusthostN",
					Description: `Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domain(s) that the administrator can access.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Access profile for this administrator. Access profiles control administrator access to FortiGate features.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_admin_profiles",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure access profiles of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"admin",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Profile name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `Scope of admin access.`,
				},
				resource.Attribute{
					Name:        "secfabgrp",
					Description: `Security Fabric.`,
				},
				resource.Attribute{
					Name:        "ftviewgrp",
					Description: `FortiView.`,
				},
				resource.Attribute{
					Name:        "authgrp",
					Description: `Administrator access to Users and Devices.`,
				},
				resource.Attribute{
					Name:        "sysgrp",
					Description: `System Configuration.`,
				},
				resource.Attribute{
					Name:        "netgrp",
					Description: `Network Configuration.`,
				},
				resource.Attribute{
					Name:        "loggrp",
					Description: `Administrator access to Logging and Reporting including viewing log messages.`,
				},
				resource.Attribute{
					Name:        "fwgrp",
					Description: `Administrator access to the Firewall configuration.`,
				},
				resource.Attribute{
					Name:        "vpngrp",
					Description: `Administrator access to IPsec, SSL, PPTP, and L2TP VPN.`,
				},
				resource.Attribute{
					Name:        "utmgrp",
					Description: `Administrator access to Security Profiles.`,
				},
				resource.Attribute{
					Name:        "wanoptgrp",
					Description: `Administrator access to WAN Opt & Cache.`,
				},
				resource.Attribute{
					Name:        "wifi",
					Description: `Administrator access to the WiFi controller and Switch controller.`,
				},
				resource.Attribute{
					Name:        "admintimeout_override",
					Description: `Enable/disable overriding the global administrator idle timeout.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the access profile item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Profile name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `Scope of admin access.`,
				},
				resource.Attribute{
					Name:        "secfabgrp",
					Description: `Security Fabric.`,
				},
				resource.Attribute{
					Name:        "ftviewgrp",
					Description: `FortiView.`,
				},
				resource.Attribute{
					Name:        "authgrp",
					Description: `Administrator access to Users and Devices.`,
				},
				resource.Attribute{
					Name:        "sysgrp",
					Description: `System Configuration.`,
				},
				resource.Attribute{
					Name:        "netgrp",
					Description: `Network Configuration.`,
				},
				resource.Attribute{
					Name:        "loggrp",
					Description: `Administrator access to Logging and Reporting including viewing log messages.`,
				},
				resource.Attribute{
					Name:        "fwgrp",
					Description: `Administrator access to the Firewall configuration.`,
				},
				resource.Attribute{
					Name:        "vpngrp",
					Description: `Administrator access to IPsec, SSL, PPTP, and L2TP VPN.`,
				},
				resource.Attribute{
					Name:        "utmgrp",
					Description: `Administrator access to Security Profiles.`,
				},
				resource.Attribute{
					Name:        "wanoptgrp",
					Description: `Administrator access to WAN Opt & Cache.`,
				},
				resource.Attribute{
					Name:        "wifi",
					Description: `Administrator access to the WiFi controller and Switch controller.`,
				},
				resource.Attribute{
					Name:        "admintimeout_override",
					Description: `Enable/disable overriding the global administrator idle timeout.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the access profile item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Profile name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `Scope of admin access.`,
				},
				resource.Attribute{
					Name:        "secfabgrp",
					Description: `Security Fabric.`,
				},
				resource.Attribute{
					Name:        "ftviewgrp",
					Description: `FortiView.`,
				},
				resource.Attribute{
					Name:        "authgrp",
					Description: `Administrator access to Users and Devices.`,
				},
				resource.Attribute{
					Name:        "sysgrp",
					Description: `System Configuration.`,
				},
				resource.Attribute{
					Name:        "netgrp",
					Description: `Network Configuration.`,
				},
				resource.Attribute{
					Name:        "loggrp",
					Description: `Administrator access to Logging and Reporting including viewing log messages.`,
				},
				resource.Attribute{
					Name:        "fwgrp",
					Description: `Administrator access to the Firewall configuration.`,
				},
				resource.Attribute{
					Name:        "vpngrp",
					Description: `Administrator access to IPsec, SSL, PPTP, and L2TP VPN.`,
				},
				resource.Attribute{
					Name:        "utmgrp",
					Description: `Administrator access to Security Profiles.`,
				},
				resource.Attribute{
					Name:        "wanoptgrp",
					Description: `Administrator access to WAN Opt & Cache.`,
				},
				resource.Attribute{
					Name:        "wifi",
					Description: `Administrator access to the WiFi controller and Switch controller.`,
				},
				resource.Attribute{
					Name:        "admintimeout_override",
					Description: `Enable/disable overriding the global administrator idle timeout.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_apiuser_setting",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure API users of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"apiuser",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User name.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `(Required) Admin user access profile.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `(Required) Virtual domains.`,
				},
				resource.Attribute{
					Name:        "trusthost-Type",
					Description: `(Required) Trusthost type.`,
				},
				resource.Attribute{
					Name:        "trusthost-ipv4_trusthost",
					Description: `(Required) IPv4 trusted host address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Admin user access profile.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domains.`,
				},
				resource.Attribute{
					Name:        "trusthost-Type",
					Description: `Trusthost type.`,
				},
				resource.Attribute{
					Name:        "trusthost-ipv4_trusthost",
					Description: `IPv4 trusted host address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Admin user access profile.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domains.`,
				},
				resource.Attribute{
					Name:        "trusthost-Type",
					Description: `Trusthost type.`,
				},
				resource.Attribute{
					Name:        "trusthost-ipv4_trusthost",
					Description: `IPv4 trusted host address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_license_forticare",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to add a FortiCare license for FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"license",
				"forticare",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "registration_code",
					Description: `(Required) Registration code.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_license_vdom",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to add a VDOM license for FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"license",
				"vdom",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "license",
					Description: `(Required) Registration code.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_license_vm",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to update VM license using uploaded file for FortiOS. Reboots immediately if successful.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"license",
				"vm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_content",
					Description: `(Required) The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_setting_dns",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure DNS of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"setting",
				"dns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "dns_over_tls",
					Description: `Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ] ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "dns_over_tls",
					Description: `Enable/disable/enforce DNS over TLS(available since v6.2.0).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "dns_over_tls",
					Description: `Enable/disable/enforce DNS over TLS(available since v6.2.0).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_setting_global",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure options related to the overall operation of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"setting",
				"global",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) FortiGate unit's hostname.`,
				},
				resource.Attribute{
					Name:        "admintimeout",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Number corresponding to your time zone from 00 to 86.`,
				},
				resource.Attribute{
					Name:        "admin_sport",
					Description: `Administrative access port for HTTPS.`,
				},
				resource.Attribute{
					Name:        "admin_ssh_port",
					Description: `Administrative access port for SSH.`,
				},
				resource.Attribute{
					Name:        "admin_scp",
					Description: `Enable SCP over SSH ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admintimeout",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Number corresponding to your time zone from 00 to 86.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `FortiGate unit's hostname.`,
				},
				resource.Attribute{
					Name:        "admin_sport",
					Description: `Administrative access port for HTTPS.`,
				},
				resource.Attribute{
					Name:        "admin_ssh_port",
					Description: `Administrative access port for SSH.`,
				},
				resource.Attribute{
					Name:        "admin_scp",
					Description: `Enable SCP over SSH`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admintimeout",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Number corresponding to your time zone from 00 to 86.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `FortiGate unit's hostname.`,
				},
				resource.Attribute{
					Name:        "admin_sport",
					Description: `Administrative access port for HTTPS.`,
				},
				resource.Attribute{
					Name:        "admin_ssh_port",
					Description: `Administrative access port for SSH.`,
				},
				resource.Attribute{
					Name:        "admin_scp",
					Description: `Enable SCP over SSH`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_setting_ntp",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure Network Time Protocol (NTP) servers of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"setting",
				"ntp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Use the FortiGuard NTP server or any other available NTP Server.`,
				},
				resource.Attribute{
					Name:        "ntpserver",
					Description: `Configure the FortiGate to connect to any available third-party NTP server.`,
				},
				resource.Attribute{
					Name:        "ntpsync",
					Description: `Enable/disable setting the FortiGate system time by synchronizing with an NTP Server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Use the FortiGuard NTP server or any other available NTP Server.`,
				},
				resource.Attribute{
					Name:        "ntpserver",
					Description: `Configure the FortiGate to connect to any available third-party NTP server.`,
				},
				resource.Attribute{
					Name:        "ntpsync",
					Description: `Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `Use the FortiGuard NTP server or any other available NTP Server.`,
				},
				resource.Attribute{
					Name:        "ntpserver",
					Description: `Configure the FortiGate to connect to any available third-party NTP server.`,
				},
				resource.Attribute{
					Name:        "ntpsync",
					Description: `Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_vdom_setting",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure VDOM of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"vdom",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) VDOM name.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `VDOM short name.`,
				},
				resource.Attribute{
					Name:        "temporary",
					Description: `Temporary. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VDOM.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VDOM name.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `VDOM short name.`,
				},
				resource.Attribute{
					Name:        "temporary",
					Description: `Temporary.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VDOM.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VDOM name.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `VDOM short name.`,
				},
				resource.Attribute{
					Name:        "temporary",
					Description: `Temporary.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_vpn_ipsec_phase1interface",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to use phase1-interface to define a phase 1 definition for a route-based (interface mode) IPsec VPN tunnel that generates authentication and encryption keys automatically.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"vpn",
				"ipsec",
				"phase1interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) IPsec remote gateway name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Remote gateway type.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) Local physical, aggregate, or VLAN outgoing interface.`,
				},
				resource.Attribute{
					Name:        "peertype",
					Description: `Accept this peer type.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase1 proposal.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "wizard_type",
					Description: `GUI VPN Wizard Type.`,
				},
				resource.Attribute{
					Name:        "remote_gw",
					Description: `(Required) IPv4 address of the remote gateway's external interface.`,
				},
				resource.Attribute{
					Name:        "psksecret",
					Description: `(Required) Pre-shared secret for PSK authentication.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Names of signed personal certificates.`,
				},
				resource.Attribute{
					Name:        "peerid",
					Description: `Accept this peer identity.`,
				},
				resource.Attribute{
					Name:        "peer",
					Description: `Accept this peer certificate.`,
				},
				resource.Attribute{
					Name:        "peergrp",
					Description: `Accept this peer certificate group.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_include",
					Description: `IPv4 split-include subnets.`,
				},
				resource.Attribute{
					Name:        "split_include_service",
					Description: `Split-include services.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_exclude",
					Description: `IPv4 subnets that should not be sent over the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "authmethod",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "authmethod_remote",
					Description: `Authentication method (remote side).`,
				},
				resource.Attribute{
					Name:        "mode_cfg",
					Description: `Enable/disable configuration method. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase1-interface item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec remote gateway name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Remote gateway type.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Local physical, aggregate, or VLAN outgoing interface.`,
				},
				resource.Attribute{
					Name:        "peertype",
					Description: `Accept this peer type.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase1 proposal.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "wizard_type",
					Description: `GUI VPN Wizard Type.`,
				},
				resource.Attribute{
					Name:        "remote_gw",
					Description: `IPv4 address of the remote gateway's external interface.`,
				},
				resource.Attribute{
					Name:        "psksecret",
					Description: `Pre-shared secret for PSK authentication.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Names of signed personal certificates.`,
				},
				resource.Attribute{
					Name:        "peerid",
					Description: `Accept this peer identity.`,
				},
				resource.Attribute{
					Name:        "peer",
					Description: `Accept this peer certificate.`,
				},
				resource.Attribute{
					Name:        "peergrp",
					Description: `Accept this peer certificate group.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_include",
					Description: `IPv4 split-include subnets.`,
				},
				resource.Attribute{
					Name:        "split_include_service",
					Description: `Split-include services.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_exclude",
					Description: `IPv4 subnets that should not be sent over the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "authmethod",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "authmethod_remote",
					Description: `Authentication method (remote side).`,
				},
				resource.Attribute{
					Name:        "mode_cfg",
					Description: `Enable/disable configuration method.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase1-interface item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec remote gateway name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Remote gateway type.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Local physical, aggregate, or VLAN outgoing interface.`,
				},
				resource.Attribute{
					Name:        "peertype",
					Description: `Accept this peer type.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase1 proposal.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "wizard_type",
					Description: `GUI VPN Wizard Type.`,
				},
				resource.Attribute{
					Name:        "remote_gw",
					Description: `IPv4 address of the remote gateway's external interface.`,
				},
				resource.Attribute{
					Name:        "psksecret",
					Description: `Pre-shared secret for PSK authentication.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Names of signed personal certificates.`,
				},
				resource.Attribute{
					Name:        "peerid",
					Description: `Accept this peer identity.`,
				},
				resource.Attribute{
					Name:        "peer",
					Description: `Accept this peer certificate.`,
				},
				resource.Attribute{
					Name:        "peergrp",
					Description: `Accept this peer certificate group.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_include",
					Description: `IPv4 split-include subnets.`,
				},
				resource.Attribute{
					Name:        "split_include_service",
					Description: `Split-include services.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_exclude",
					Description: `IPv4 subnets that should not be sent over the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "authmethod",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "authmethod_remote",
					Description: `Authentication method (remote side).`,
				},
				resource.Attribute{
					Name:        "mode_cfg",
					Description: `Enable/disable configuration method.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_vpn_ipsec_phase2interface",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to use phase2-interface to add or edit a phase 2 configuration on a route-based (interface mode) IPsec tunnel.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"vpn",
				"ipsec",
				"phase2interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) IPsec tunnel name.`,
				},
				resource.Attribute{
					Name:        "phase1name",
					Description: `(Required) Phase 1 determines the options required for phase 2.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase2 proposal.`,
				},
				resource.Attribute{
					Name:        "src_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_start_ip",
					Description: `Local proxy ID start.`,
				},
				resource.Attribute{
					Name:        "src_end_ip",
					Description: `Local proxy ID end.`,
				},
				resource.Attribute{
					Name:        "src_subnet",
					Description: `Local proxy ID subnet.`,
				},
				resource.Attribute{
					Name:        "dst_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_name",
					Description: `Local proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_name",
					Description: `Remote proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_start_ip",
					Description: `Remote proxy ID IPv4 start.`,
				},
				resource.Attribute{
					Name:        "dst_end_ip",
					Description: `Remote proxy ID IPv4 end.`,
				},
				resource.Attribute{
					Name:        "dst_subnet",
					Description: `Remote proxy ID IPv4 subnet.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase2-interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec tunnel name.`,
				},
				resource.Attribute{
					Name:        "phase1name",
					Description: `Phase 1 determines the options required for phase 2.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase2 proposal.`,
				},
				resource.Attribute{
					Name:        "src_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_start_ip",
					Description: `Local proxy ID start.`,
				},
				resource.Attribute{
					Name:        "src_end_ip",
					Description: `Local proxy ID end.`,
				},
				resource.Attribute{
					Name:        "src_subnet",
					Description: `Local proxy ID subnet.`,
				},
				resource.Attribute{
					Name:        "dst_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_name",
					Description: `Local proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_name",
					Description: `Remote proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_start_ip",
					Description: `Remote proxy ID IPv4 start.`,
				},
				resource.Attribute{
					Name:        "dst_end_ip",
					Description: `Remote proxy ID IPv4 end.`,
				},
				resource.Attribute{
					Name:        "dst_subnet",
					Description: `Remote proxy ID IPv4 subnet.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase2-interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec tunnel name.`,
				},
				resource.Attribute{
					Name:        "phase1name",
					Description: `Phase 1 determines the options required for phase 2.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase2 proposal.`,
				},
				resource.Attribute{
					Name:        "src_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_start_ip",
					Description: `Local proxy ID start.`,
				},
				resource.Attribute{
					Name:        "src_end_ip",
					Description: `Local proxy ID end.`,
				},
				resource.Attribute{
					Name:        "src_subnet",
					Description: `Local proxy ID subnet.`,
				},
				resource.Attribute{
					Name:        "dst_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_name",
					Description: `Local proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_name",
					Description: `Remote proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_start_ip",
					Description: `Remote proxy ID IPv4 start.`,
				},
				resource.Attribute{
					Name:        "dst_end_ip",
					Description: `Remote proxy ID IPv4 end.`,
				},
				resource.Attribute{
					Name:        "dst_subnet",
					Description: `Remote proxy ID IPv4 subnet.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"fortios_firewall_object_address":         0,
		"fortios_firewall_object_addressgroup":    1,
		"fortios_firewall_object_ippool":          2,
		"fortios_firewall_object_service":         3,
		"fortios_firewall_object_servicecategory": 4,
		"fortios_firewall_object_servicegroup":    5,
		"fortios_firewall_object_vip":             6,
		"fortios_firewall_object_vipgroup":        7,
		"fortios_firewall_security_policy":        8,
		"fortios_firewall_security_policyseq":     9,
		"fortios_log_fortianalyzer_setting":       10,
		"fortios_log_syslog_setting":              11,
		"fortios_networking_interface_port":       12,
		"fortios_networking_route_static":         13,
		"fortios_system_admin_administrator":      14,
		"fortios_system_admin_profiles":           15,
		"fortios_system_apiuser_setting":          16,
		"fortios_system_license_forticare":        17,
		"fortios_system_license_vdom":             18,
		"fortios_system_license_vm":               19,
		"fortios_system_setting_dns":              20,
		"fortios_system_setting_global":           21,
		"fortios_system_setting_ntp":              22,
		"fortios_system_vdom_setting":             23,
		"fortios_vpn_ipsec_phase1interface":       24,
		"fortios_vpn_ipsec_phase2interface":       25,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
