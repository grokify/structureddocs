package metadata

import "strings"

// Status represents the status of a document or item.
type Status string

// Common document statuses.
const (
	StatusDraft      Status = "draft"
	StatusReview     Status = "review"
	StatusApproved   Status = "approved"
	StatusActive     Status = "active"
	StatusDeprecated Status = "deprecated"
	StatusArchived   Status = "archived"
)

// Common item/task statuses.
const (
	StatusNotStarted Status = "not_started"
	StatusPlanning   Status = "planning"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusBlocked    Status = "blocked"
	StatusOnHold     Status = "on_hold"
	StatusCancelled  Status = "cancelled"
)

// Common tracking statuses.
const (
	StatusOnTrack Status = "on_track"
	StatusAtRisk  Status = "at_risk"
	StatusBehind  Status = "behind"
)

// DocumentStatuses returns the list of valid document statuses.
func DocumentStatuses() []Status {
	return []Status{
		StatusDraft,
		StatusReview,
		StatusApproved,
		StatusActive,
		StatusDeprecated,
		StatusArchived,
	}
}

// ItemStatuses returns the list of valid item/task statuses.
func ItemStatuses() []Status {
	return []Status{
		StatusNotStarted,
		StatusPlanning,
		StatusInProgress,
		StatusCompleted,
		StatusBlocked,
		StatusOnHold,
		StatusCancelled,
	}
}

// TrackingStatuses returns the list of valid tracking statuses.
func TrackingStatuses() []Status {
	return []Status{
		StatusOnTrack,
		StatusAtRisk,
		StatusBehind,
	}
}

// String returns the status as a string.
func (s Status) String() string {
	return string(s)
}

// Display returns a human-readable display string.
func (s Status) Display() string {
	str := string(s)
	str = strings.ReplaceAll(str, "_", " ")
	return strings.Title(str) //nolint:staticcheck
}

// IsComplete returns true if the status indicates completion.
func (s Status) IsComplete() bool {
	switch s {
	case StatusCompleted, StatusApproved, StatusArchived:
		return true
	default:
		return false
	}
}

// IsActive returns true if the status indicates active work.
func (s Status) IsActive() bool {
	switch s {
	case StatusInProgress, StatusActive, StatusReview:
		return true
	default:
		return false
	}
}

// IsBlocking returns true if the status indicates a blocking condition.
func (s Status) IsBlocking() bool {
	switch s {
	case StatusBlocked, StatusOnHold, StatusCancelled:
		return true
	default:
		return false
	}
}

// Priority represents the priority level of an item.
type Priority string

// Priority levels.
const (
	PriorityCritical Priority = "critical"
	PriorityHigh     Priority = "high"
	PriorityMedium   Priority = "medium"
	PriorityLow      Priority = "low"
)

// P-notation priority levels.
const (
	PriorityP0 Priority = "P0"
	PriorityP1 Priority = "P1"
	PriorityP2 Priority = "P2"
	PriorityP3 Priority = "P3"
)

// Priorities returns the list of standard priorities.
func Priorities() []Priority {
	return []Priority{
		PriorityCritical,
		PriorityHigh,
		PriorityMedium,
		PriorityLow,
	}
}

// String returns the priority as a string.
func (p Priority) String() string {
	return string(p)
}

// Normalize converts P-notation to standard priority.
func (p Priority) Normalize() Priority {
	switch p {
	case PriorityP0:
		return PriorityCritical
	case PriorityP1:
		return PriorityHigh
	case PriorityP2:
		return PriorityMedium
	case PriorityP3:
		return PriorityLow
	default:
		return p
	}
}

// Level returns a numeric level (0 = highest priority).
func (p Priority) Level() int {
	switch p.Normalize() {
	case PriorityCritical:
		return 0
	case PriorityHigh:
		return 1
	case PriorityMedium:
		return 2
	case PriorityLow:
		return 3
	default:
		return 99
	}
}
