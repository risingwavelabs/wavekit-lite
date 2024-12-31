"use client"

import { useState } from "react"
import { ClusterList, Cluster } from "@/components/ui/cluster-list"
import { ClusterDialog, ClusterFormData } from "@/components/ui/new-cluster-dialog"

const sampleClusters: Cluster[] = [
  {
    id: "cluster-1",
    name: "Production DB",
    status: "running",
    host: "localhost",
    sqlPort: 8080,
    metaNodePort: 9090,
    user: "admin",
    database: "prod_db"
  },
  {
    id: "cluster-2",
    name: "Staging Environment",
    status: "running",
    host: "localhost",
    sqlPort: 8081,
    metaNodePort: 9091,
    user: "admin",
    database: "staging_db"
  },
  {
    id: "cluster-3",
    name: "Development DB",
    status: "stopped",
    host: "localhost",
    sqlPort: 8082,
    metaNodePort: 9092,
    user: "admin",
    database: "dev_db"
  },
  {
    id: "cluster-4",
    name: "Analytics Cluster",
    status: "running",
    host: "localhost",
    sqlPort: 8083,
    metaNodePort: 9093,
    user: "admin",
    database: "analytics_db"
  },
  {
    id: "cluster-5",
    name: "Testing Environment",
    status: "error",
    host: "localhost",
    sqlPort: 8084,
    metaNodePort: 9094,
    user: "admin",
    database: "test_db"
  },
  {
    id: "cluster-6",
    name: "Backup DB",
    status: "stopped",
    host: "localhost",
    sqlPort: 8085,
    metaNodePort: 9095,
    user: "admin",
    database: "backup_db"
  },
  {
    id: "cluster-7",
    name: "Data Warehouse",
    status: "running",
    host: "localhost",
    sqlPort: 8086,
    metaNodePort: 9096,
    user: "admin",
    database: "warehouse_db"
  },
  {
    id: "cluster-8",
    name: "Reporting DB",
    status: "running",
    host: "localhost",
    sqlPort: 8087,
    metaNodePort: 9097,
    user: "admin",
    database: "reporting_db"
  },
  {
    id: "cluster-9",
    name: "QA Environment",
    status: "stopped",
    host: "localhost",
    sqlPort: 8088,
    metaNodePort: 9098,
    user: "admin",
    database: "qa_db"
  },
  {
    id: "cluster-10",
    name: "Archive DB",
    status: "running",
    host: "localhost",
    sqlPort: 8089,
    metaNodePort: 9099,
    user: "admin",
    database: "archive_db"
  },
  {
    id: "cluster-11",
    name: "ML Training Cluster",
    status: "error",
    host: "localhost",
    sqlPort: 8090,
    metaNodePort: 9100,
    user: "admin",
    database: "ml_db"
  },
  {
    id: "cluster-12",
    name: "Cache Cluster",
    status: "running",
    host: "localhost",
    sqlPort: 8091,
    metaNodePort: 9101,
    user: "admin",
    database: "cache_db"
  }
]

export default function ClustersPage() {
  const [clusters, setClusters] = useState<Cluster[]>(sampleClusters)

  const handleCreateCluster = (data: ClusterFormData) => {
    const newCluster: Cluster = {
      id: `cluster-${clusters.length + 1}`,
      status: "stopped",
      ...data
    }
    setClusters(prev => [...prev, newCluster])
  }

  const handleEditCluster = (id: string, data: ClusterFormData) => {
    setClusters(prev => prev.map(cluster => 
      cluster.id === id 
        ? { ...cluster, ...data }
        : cluster
    ))
  }

  return (
    <div className="p-8">
      <div className="flex items-center justify-between mb-8">
        <div>
          <h1 className="text-2xl font-semibold">Clusters</h1>
          <p className="text-sm text-muted-foreground">
            Manage your database clusters
          </p>
        </div>
        <ClusterDialog mode="create" onSubmit={handleCreateCluster} />
      </div>
      <ClusterList 
        clusters={clusters}
        onEdit={(cluster) => handleEditCluster(cluster.id, cluster)}
      />
    </div>
  )
}
