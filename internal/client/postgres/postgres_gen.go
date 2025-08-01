// Package postgres provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package postgres

import (
	"time"
)

// Defines values for PostgresPlans.
const (
	Accelerated1024gb PostgresPlans = "accelerated_1024gb"
	Accelerated128gb  PostgresPlans = "accelerated_128gb"
	Accelerated16gb   PostgresPlans = "accelerated_16gb"
	Accelerated256gb  PostgresPlans = "accelerated_256gb"
	Accelerated32gb   PostgresPlans = "accelerated_32gb"
	Accelerated384gb  PostgresPlans = "accelerated_384gb"
	Accelerated512gb  PostgresPlans = "accelerated_512gb"
	Accelerated64gb   PostgresPlans = "accelerated_64gb"
	Accelerated768gb  PostgresPlans = "accelerated_768gb"
	Basic1gb          PostgresPlans = "basic_1gb"
	Basic256mb        PostgresPlans = "basic_256mb"
	Basic4gb          PostgresPlans = "basic_4gb"
	Custom            PostgresPlans = "custom"
	Free              PostgresPlans = "free"
	Pro               PostgresPlans = "pro"
	Pro128gb          PostgresPlans = "pro_128gb"
	Pro16gb           PostgresPlans = "pro_16gb"
	Pro192gb          PostgresPlans = "pro_192gb"
	Pro256gb          PostgresPlans = "pro_256gb"
	Pro32gb           PostgresPlans = "pro_32gb"
	Pro384gb          PostgresPlans = "pro_384gb"
	Pro4gb            PostgresPlans = "pro_4gb"
	Pro512gb          PostgresPlans = "pro_512gb"
	Pro64gb           PostgresPlans = "pro_64gb"
	Pro8gb            PostgresPlans = "pro_8gb"
	ProPlus           PostgresPlans = "pro_plus"
	Standard          PostgresPlans = "standard"
	Starter           PostgresPlans = "starter"
)

// Defines values for RecoveryInfoRecoveryStatus.
const (
	AVAILABLE      RecoveryInfoRecoveryStatus = "AVAILABLE"
	BACKUPNOTREADY RecoveryInfoRecoveryStatus = "BACKUP_NOT_READY"
	NOTAVAILABLE   RecoveryInfoRecoveryStatus = "NOT_AVAILABLE"
)

// PostgresExport defines model for postgresExport.
type PostgresExport struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`

	// Url URL to download the Postgres export
	Url *string `json:"url,omitempty"`
}

// PostgresPlans defines model for postgresPlans.
type PostgresPlans string

// RecoveryInfo defines model for recoveryInfo.
type RecoveryInfo struct {
	// RecoveryStatus Availability of point-in-time recovery.
	RecoveryStatus RecoveryInfoRecoveryStatus `json:"recoveryStatus"`
	StartsAt       *time.Time                 `json:"startsAt,omitempty"`
}

// RecoveryInfoRecoveryStatus Availability of point-in-time recovery.
type RecoveryInfoRecoveryStatus string

// RecoveryInput defines model for recoveryInput.
type RecoveryInput struct {
	// DatadogApiKey Datadog API key to use for monitoring the new database. Defaults to the API key of the original database. Use an empty string to prevent copying of the API key to the new database.
	DatadogApiKey *string `json:"datadogApiKey,omitempty"`

	// DatadogSite Datadog region code to use for monitoring the new database. Defaults to the region code of the original database. Use an empty string to prevent copying of the region code to the new database.
	DatadogSite *string `json:"datadogSite,omitempty"`

	// EnvironmentId The environment to create the new database in. Defaults to the environment of the original database.
	EnvironmentId *string `json:"environmentId,omitempty"`

	// Plan The plan to use for the new database. Defaults to the same plan as the original database. Cannot be a lower tier plan than the original database.
	Plan *string `json:"plan,omitempty"`

	// RestoreName Name of the new database.
	RestoreName *string `json:"restoreName,omitempty"`

	// RestoreTime The point in time to restore the database to. See `/recovery-info` for restore availability
	RestoreTime time.Time `json:"restoreTime"`
}
