// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type RelationStatus string

const (
	RelationStatusFollow RelationStatus = "follow"
)

func (e *RelationStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RelationStatus(s)
	case string:
		*e = RelationStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for RelationStatus: %T", src)
	}
	return nil
}

type NullRelationStatus struct {
	RelationStatus RelationStatus
	Valid          bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRelationStatus) Scan(value interface{}) error {
	if value == nil {
		ns.RelationStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RelationStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRelationStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.RelationStatus, nil
}

type Account struct {
	AccountID   uuid.UUID      `json:"account_id"`
	AccountName string         `json:"account_name"`
	Dob         time.Time      `json:"dob"`
	Address     sql.NullString `json:"address"`
	Description sql.NullString `json:"description"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	CreatedBy   uuid.NullUUID  `json:"created_by"`
}

type Relation struct {
	ObserverID uuid.UUID      `json:"observer_id"`
	ObservedID uuid.UUID      `json:"observed_id"`
	Status     RelationStatus `json:"status"`
}

type User struct {
	UserID    uuid.UUID      `json:"user_id"`
	FirstName string         `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}
