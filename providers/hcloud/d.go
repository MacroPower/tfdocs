package hcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hcloud_certificate",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_datacenter",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Datacenter.`,
			Description: `
Provides details about a specific Hetzner Cloud Datacenter.
Use this resource to get detailed information about specific datacenter.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_datacenters",
			Category:         "Data Sources",
			ShortDescription: `List all available Hetzner Cloud Datacenters.`,
			Description: `
Provides a list of available Hetzner Cloud Datacenters.
This resource may be useful to create highly available infrastructure, distributed across several datacenters.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_firewall",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Firewall.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_floating_ip",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Floating IP.`,
			Description: `

Provides details about a Hetzner Cloud Floating IP.

This resource can be useful when you need to determine a Floating IP ID based on the IP address.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_image",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Image.`,
			Description: `
Provides details about a Hetzner Cloud Image.
This resource is useful if you want to use a non-terraform managed image.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_load_balancer",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Load Balancer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_location",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Location.`,
			Description: `
Provides details about a specific Hetzner Cloud Location.
Use this resource to get detailed information about specific location.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_locations",
			Category:         "Data Sources",
			ShortDescription: `List all available Hetzner Cloud Locations.`,
			Description: `
Provides a list of available Hetzner Cloud Locations.
This resource may be useful to create highly available infrastructure, distributed across several locations.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_network",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud network.`,
			Description: `
Provides details about a Hetzner Cloud network.
This resource is useful if you want to use a non-terraform managed network.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Server.`,
			Description: `
Provides details about a Hetzner Cloud Server.
This resource is useful if you want to use a non-terraform managed server.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server_type",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Server Type.`,
			Description: `
Provides details about a specific Hetzner Cloud Server Type.
Use this resource to get detailed information about specific Server Type.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server_types",
			Category:         "Data Sources",
			ShortDescription: `List all available Hetzner Cloud Server Types.`,
			Description: `
Provides a list of available Hetzner Cloud Server Types.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud SSH Key.`,
			Description: `
Provides details about a Hetzner Cloud SSH Key.
This resource is useful if you want to use a non-terraform managed SSH Key.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_ssh_keys",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud SSH Keys.`,
			Description: `

Provides details about Hetzner Cloud SSH Keys.
This resource is useful if you want to use a non-terraform managed SSH Key.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_volume",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud volume.`,
			Description: `
Provides details about a Hetzner Cloud volume.
This resource is useful if you want to use a non-terraform managed volume.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"hcloud_certificate":   0,
		"hcloud_datacenter":    1,
		"hcloud_datacenters":   2,
		"hcloud_firewall":      3,
		"hcloud_floating_ip":   4,
		"hcloud_image":         5,
		"hcloud_load_balancer": 6,
		"hcloud_location":      7,
		"hcloud_locations":     8,
		"hcloud_network":       9,
		"hcloud_server":        10,
		"hcloud_server_type":   11,
		"hcloud_server_types":  12,
		"hcloud_ssh_key":       13,
		"hcloud_ssh_keys":      14,
		"hcloud_volume":        15,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
