package hcp

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hcp_aws_network_peering",
			Category:         "Data Sources",
			ShortDescription: `The AWS Network peering data source provides information about an existing Network peering between an HVN and a peer AWS VPC.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_aws_transit_gateway_attachment",
			Category:         "Data Sources",
			ShortDescription: `The AWS Transit Gateway Attachment data source provides information about an existing transit gateway attachment.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_agent_helm_config",
			Category:         "Data Sources",
			ShortDescription: `The Consul agent Helm config data source provides Helm values for a Consul agent running in Kubernetes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_agent_kubernetes_secret",
			Category:         "Data Sources",
			ShortDescription: `The agent config Kubernetes secret data source provides Consul agents running in Kubernetes the configuration needed to connect to the Consul cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_cluster",
			Category:         "Data Sources",
			ShortDescription: `The cluster data source provides information about an existing HCP Consul cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_versions",
			Category:         "Data Sources",
			ShortDescription: `The Consul versions data source provides the Consul versions supported by HCP.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_hvn",
			Category:         "Data Sources",
			ShortDescription: `The HVN data source provides information about an existing HashiCorp Virtual Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"hcp_aws_network_peering":            0,
		"hcp_aws_transit_gateway_attachment": 1,
		"hcp_consul_agent_helm_config":       2,
		"hcp_consul_agent_kubernetes_secret": 3,
		"hcp_consul_cluster":                 4,
		"hcp_consul_versions":                5,
		"hcp_hvn":                            6,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
