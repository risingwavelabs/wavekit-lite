/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Cluster } from '../models/Cluster';
import type { ClusterCreate } from '../models/ClusterCreate';
import type { Credentials } from '../models/Credentials';
import type { Database } from '../models/Database';
import type { DatabaseConnectInfo } from '../models/DatabaseConnectInfo';
import type { DDLProgress } from '../models/DDLProgress';
import type { DiagnosticConfig } from '../models/DiagnosticConfig';
import type { DiagnosticData } from '../models/DiagnosticData';
import type { QueryRequest } from '../models/QueryRequest';
import type { QueryResponse } from '../models/QueryResponse';
import type { RefreshTokenRequest } from '../models/RefreshTokenRequest';
import type { SignInRequest } from '../models/SignInRequest';
import type { Snapshot } from '../models/Snapshot';
import type { SnapshotConfig } from '../models/SnapshotConfig';
import type { SnapshotCreate } from '../models/SnapshotCreate';
import type { TestClusterConnectionPayload } from '../models/TestClusterConnectionPayload';
import type { TestClusterConnectionResult } from '../models/TestClusterConnectionResult';
import type { TestDatabaseConnectionPayload } from '../models/TestDatabaseConnectionPayload';
import type { TestDatabaseConnectionResult } from '../models/TestDatabaseConnectionResult';
import type { UpdateClusterRequest } from '../models/UpdateClusterRequest';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class DefaultService {
    /**
     * List all databases
     * Retrieve a list of all databases and their tables
     * @returns Database Successfully retrieved database list
     * @throws ApiError
     */
    public static listDatabases(): CancelablePromise<Array<Database>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/databases',
        });
    }
    /**
     * Create a new database
     * Create a new database
     * @param requestBody
     * @returns Database Database created successfully
     * @throws ApiError
     */
    public static createDatabase(
        requestBody: DatabaseConnectInfo,
    ): CancelablePromise<Database> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/databases',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Test database connection
     * Test a database connection
     * @param requestBody
     * @returns TestDatabaseConnectionResult Successfully tested database connection
     * @throws ApiError
     */
    public static testDatabaseConnection(
        requestBody: TestDatabaseConnectionPayload,
    ): CancelablePromise<TestDatabaseConnectionResult> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/databases/test-connection',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get database details
     * Retrieve details of a specific database
     * @param id
     * @returns Database Successfully retrieved database
     * @throws ApiError
     */
    public static getDatabase(
        id: number,
    ): CancelablePromise<Database> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/databases/{ID}',
            path: {
                'ID': id,
            },
        });
    }
    /**
     * Update database
     * Update a specific database
     * @param id
     * @param requestBody
     * @returns Database Database updated successfully
     * @throws ApiError
     */
    public static updateDatabase(
        id: number,
        requestBody: DatabaseConnectInfo,
    ): CancelablePromise<Database> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/databases/{ID}',
            path: {
                'ID': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete database
     * Delete a specific database
     * @param id
     * @returns void
     * @throws ApiError
     */
    public static deleteDatabase(
        id: number,
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/databases/{ID}',
            path: {
                'ID': id,
            },
        });
    }
    /**
     * Query database
     * Query a specific database
     * @param id
     * @param requestBody
     * @returns QueryResponse Query executed successfully
     * @throws ApiError
     */
    public static queryDatabase(
        id: number,
        requestBody: QueryRequest,
    ): CancelablePromise<QueryResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/databases/{ID}/query',
            path: {
                'ID': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get DDL progress
     * Get the progress of a DDL operation
     * @param id
     * @returns DDLProgress Successfully retrieved DDL progress
     * @throws ApiError
     */
    public static getDdlProgress(
        id: number,
    ): CancelablePromise<Array<DDLProgress>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/databases/{ID}/ddl-progress',
            path: {
                'ID': id,
            },
        });
    }
    /**
     * @param id
     * @param ddlId
     * @returns any Successfully canceled DDL operation
     * @throws ApiError
     */
    public static cancelDdlProgress(
        id: number,
        ddlId: number,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/databases/{ID}/ddl-progress/{ddlID}/cancel',
            path: {
                'ID': id,
                'ddlID': ddlId,
            },
        });
    }
    /**
     * Test cluster connection
     * Test a cluster connection
     * @param requestBody
     * @returns TestClusterConnectionResult Successfully tested cluster connection
     * @throws ApiError
     */
    public static testClusterConnection(
        requestBody: TestClusterConnectionPayload,
    ): CancelablePromise<TestClusterConnectionResult> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/test-cluster-connection',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * List all cluster versions
     * Retrieve a list of all cluster versions
     * @returns string Successfully retrieved cluster version list
     * @throws ApiError
     */
    public static listClusterVersions(): CancelablePromise<Array<string>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/cluster-versions',
        });
    }
    /**
     * List all clusters
     * Retrieve a list of all database clusters
     * @returns Cluster Successfully retrieved cluster list
     * @throws ApiError
     */
    public static listClusters(): CancelablePromise<Array<Cluster>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/clusters',
        });
    }
    /**
     * Create a new cluster
     * Create a new database cluster
     * @param requestBody
     * @returns Cluster Cluster created successfully
     * @throws ApiError
     */
    public static createCluster(
        requestBody: ClusterCreate,
    ): CancelablePromise<Cluster> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/clusters',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get cluster details
     * Retrieve details of a specific cluster
     * @param id
     * @returns Cluster Successfully retrieved cluster
     * @throws ApiError
     */
    public static getCluster(
        id: number,
    ): CancelablePromise<Cluster> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/clusters/{ID}',
            path: {
                'ID': id,
            },
        });
    }
    /**
     * Update cluster
     * Update a specific cluster
     * @param id
     * @param requestBody
     * @returns Cluster Cluster updated successfully
     * @throws ApiError
     */
    public static updateCluster(
        id: number,
        requestBody: UpdateClusterRequest,
    ): CancelablePromise<Cluster> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/clusters/{ID}',
            path: {
                'ID': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete cluster
     * Delete a specific cluster
     * @param id
     * @param cascade
     * @returns void
     * @throws ApiError
     */
    public static deleteCluster(
        id: number,
        cascade?: boolean,
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/clusters/{ID}',
            path: {
                'ID': id,
            },
            query: {
                'cascade': cascade,
            },
        });
    }
    /**
     * List cluster snapshots
     * Retrieve a list of all snapshots for a specific cluster
     * @param id
     * @returns Snapshot Successfully retrieved snapshot list
     * @throws ApiError
     */
    public static listClusterSnapshots(
        id: number,
    ): CancelablePromise<Array<Snapshot>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/clusters/{ID}/snapshots',
            path: {
                'ID': id,
            },
        });
    }
    /**
     * Create a new snapshot
     * Create a new metadata snapshot for a specific cluster
     * @param id
     * @param requestBody
     * @returns Snapshot Snapshot created successfully
     * @throws ApiError
     */
    public static createClusterSnapshot(
        id: number,
        requestBody: SnapshotCreate,
    ): CancelablePromise<Snapshot> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/clusters/{ID}/snapshots',
            path: {
                'ID': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete snapshot
     * Delete a specific snapshot
     * @param id
     * @param snapshotId
     * @returns void
     * @throws ApiError
     */
    public static deleteClusterSnapshot(
        id: number,
        snapshotId: number,
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/clusters/{ID}/snapshots/{snapshotId}',
            path: {
                'ID': id,
                'snapshotId': snapshotId,
            },
        });
    }
    /**
     * Restore snapshot
     * Restore cluster metadata from a specific snapshot
     * @param id
     * @param snapshotId
     * @returns any Snapshot restored successfully
     * @throws ApiError
     */
    public static restoreClusterSnapshot(
        id: number,
        snapshotId: number,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/clusters/{ID}/snapshots/{snapshotId}',
            path: {
                'ID': id,
                'snapshotId': snapshotId,
            },
        });
    }
    /**
     * Get snapshot configuration
     * Get automatic snapshot configuration for a cluster
     * @param id
     * @returns SnapshotConfig Successfully retrieved snapshot configuration
     * @throws ApiError
     */
    public static getClusterSnapshotConfig(
        id: number,
    ): CancelablePromise<SnapshotConfig> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/clusters/{ID}/snapshot-config',
            path: {
                'ID': id,
            },
        });
    }
    /**
     * Update snapshot configuration
     * Update automatic snapshot configuration for a cluster
     * @param id
     * @param requestBody
     * @returns SnapshotConfig Snapshot configuration updated successfully
     * @throws ApiError
     */
    public static updateClusterSnapshotConfig(
        id: number,
        requestBody: SnapshotConfig,
    ): CancelablePromise<SnapshotConfig> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/clusters/{ID}/snapshot-config',
            path: {
                'ID': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * List diagnostic data
     * Retrieve diagnostic data for a specific cluster with optional date range filtering
     * @param id
     * @param from Start date for filtering diagnostic data
     * @param to End date for filtering diagnostic data
     * @param page Page number for pagination
     * @param perPage Number of items per page
     * @returns any Successfully retrieved diagnostic data
     * @throws ApiError
     */
    public static listClusterDiagnostics(
        id: number,
        from?: string,
        to?: string,
        page: number = 1,
        perPage: number = 20,
    ): CancelablePromise<{
        items: Array<DiagnosticData>;
        /**
         * Total number of diagnostic entries
         */
        total: number;
        /**
         * Current page number
         */
        page: number;
        /**
         * Number of items per page
         */
        perPage: number;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/clusters/{ID}/diagnostics',
            path: {
                'ID': id,
            },
            query: {
                'from': from,
                'to': to,
                'page': page,
                'perPage': perPage,
            },
        });
    }
    /**
     * Get diagnostic configuration
     * Get diagnostic data collection configuration for a cluster
     * @param id
     * @returns DiagnosticConfig Successfully retrieved diagnostic configuration
     * @throws ApiError
     */
    public static getClusterDiagnosticConfig(
        id: number,
    ): CancelablePromise<DiagnosticConfig> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/clusters/{ID}/diagnostics/config',
            path: {
                'ID': id,
            },
        });
    }
    /**
     * Update diagnostic configuration
     * Update diagnostic data collection configuration for a cluster
     * @param id
     * @param requestBody
     * @returns DiagnosticConfig Diagnostic configuration updated successfully
     * @throws ApiError
     */
    public static updateClusterDiagnosticConfig(
        id: number,
        requestBody: DiagnosticConfig,
    ): CancelablePromise<DiagnosticConfig> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/clusters/{ID}/diagnostics/config',
            path: {
                'ID': id,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Sign in user
     * Authenticate user and return access token
     * @param requestBody
     * @returns Credentials Successfully authenticated
     * @throws ApiError
     */
    public static signIn(
        requestBody: SignInRequest,
    ): CancelablePromise<Credentials> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/sign-in',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                401: `Invalid credentials`,
            },
        });
    }
    /**
     * Refresh access token
     * Get a new access token using a refresh token
     * @param requestBody
     * @returns Credentials Successfully refreshed token
     * @throws ApiError
     */
    public static refreshToken(
        requestBody: RefreshTokenRequest,
    ): CancelablePromise<Credentials> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/refresh',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                401: `Invalid or expired refresh token`,
            },
        });
    }
}
