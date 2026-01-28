package gophitor

import (
	"net"
	"time"
)

type ActionCategory string

const (
	CategoryAuthentication   ActionCategory = "authentication"
	CategoryAuthorization    ActionCategory = "authorization"
	CategoryDataAccess       ActionCategory = "data_access"
	CategoryDataModification ActionCategory = "data_modification"
	CategoryDataDeletion     ActionCategory = "data_deletion"
	CategoryConfiguration    ActionCategory = "configuration"
	CategoryAdministration   ActionCategory = "administration"
	CategorySecurity         ActionCategory = "security"
	CategorySystem           ActionCategory = "system"
)

type ActionStatus string

const (
	StatusSuccesded ActionStatus = "succeeded"
	StatusFailed    ActionStatus = "failed"
	StatusDenied    ActionStatus = "denied"
)

const ()

type Entry struct {
	ID string `json:"id"`

	// Actor
	ActorID        string `json:"actor_id"`
	ActorType      string `json:"actor_type"`
	ActorIP        net.IP `json:"actor_ip,omitempty"`         // Redactable
	ActorUserAgent string `json:"actor_user_agent,omitempty"` // Redactable

	// Tenant
	TenantID    string `json:"tenant_id"`
	WorkspaceID string `json:"workspace_id"`

	// Action
	Action         string         `json:"action"`
	ActionCategory ActionCategory `json:"action_category"`
	ActionStatus   ActionStatus   `json:"action_status"`

	// Target
	ResourceType string `json:"resource_type"`
	ResourceID   string `json:"resource_id,omitempty"`
	ResouceName  string `json:"resource_name,omitempty"`

	// Cahges (State)
	PreviousState map[string]any `json:"previous_state,omitempty"`

	NewState      map[string]any `json:"new_state,omitempty"`
	CahgesSummary string         `json:"changes_summary,omitempty"`

	// Metadata
	Metadata  map[string]any `json:"metadata,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
	TraceID   string         `json:"trace_id,omitempty"`

	// Integrity
	Hash         string `json:"hash,omitempty"`
	PreviousHash string `json:"previous_hash,omitempty"`

	Timestamp  time.Time `json:"timestamp"`
	DurationMS int       `json:"duration_ms,omitempty"`
}
