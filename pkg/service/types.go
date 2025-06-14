package service

import (
	"github.com/risingwavelabs/risingwave-console/pkg/zgen/apigen"
	"github.com/risingwavelabs/risingwave-console/pkg/zgen/querier"
)

func clusterToApi(cluster *querier.Cluster) *apigen.Cluster {
	return &apigen.Cluster{
		ID:             cluster.ID,
		OrgID:          cluster.OrgID,
		Name:           cluster.Name,
		Host:           cluster.Host,
		Version:        cluster.Version,
		SqlPort:        cluster.SqlPort,
		MetaPort:       cluster.MetaPort,
		HttpPort:       cluster.HttpPort,
		CreatedAt:      cluster.CreatedAt,
		UpdatedAt:      cluster.UpdatedAt,
		MetricsStoreID: cluster.MetricsStoreID,
	}
}
