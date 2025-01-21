// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querier

import (
	"time"
)

type AutoBackupConfig struct {
	ClusterID      int32
	Enabled        bool
	CronExpression string
	KeepLast       int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type AutoDiagnosticsConfig struct {
	ClusterID         int32
	Enabled           bool
	CronExpression    string
	RetentionDuration *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Cluster struct {
	ID             int32
	OrganizationID int32
	Name           string
	Host           string
	SqlPort        int32
	MetaPort       int32
	HttpPort       int32
	Version        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ClusterDiagnostic struct {
	ID        int32
	ClusterID int32
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ClusterSnapshot struct {
	ClusterID  int32
	SnapshotID int64
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type DatabaseConnection struct {
	ID             int32
	OrganizationID int32
	Name           string
	ClusterID      int32
	Username       string
	Password       *string
	Database       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type OrgSetting struct {
	OrganizationID int32
	Timezone       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Organization struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrganizationOwner struct {
	UserID         int32
	OrganizationID int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type RefreshToken struct {
	ID        int32
	UserID    int32
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID             int32
	Name           string
	PasswordHash   string
	PasswordSalt   string
	OrganizationID int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
