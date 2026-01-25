package metadata

import "testing"

func TestStatusDisplay(t *testing.T) {
	tests := []struct {
		status Status
		want   string
	}{
		{StatusDraft, "Draft"},
		{StatusInProgress, "In Progress"},
		{StatusNotStarted, "Not Started"},
		{StatusOnTrack, "On Track"},
	}

	for _, tt := range tests {
		got := tt.status.Display()
		if got != tt.want {
			t.Errorf("Status(%q).Display() = %q, want %q", tt.status, got, tt.want)
		}
	}
}

func TestStatusIsComplete(t *testing.T) {
	complete := []Status{StatusCompleted, StatusApproved, StatusArchived}
	incomplete := []Status{StatusDraft, StatusInProgress, StatusBlocked}

	for _, s := range complete {
		if !s.IsComplete() {
			t.Errorf("Status(%q).IsComplete() = false, want true", s)
		}
	}

	for _, s := range incomplete {
		if s.IsComplete() {
			t.Errorf("Status(%q).IsComplete() = true, want false", s)
		}
	}
}

func TestStatusIsActive(t *testing.T) {
	active := []Status{StatusInProgress, StatusActive, StatusReview}
	inactive := []Status{StatusDraft, StatusCompleted, StatusBlocked}

	for _, s := range active {
		if !s.IsActive() {
			t.Errorf("Status(%q).IsActive() = false, want true", s)
		}
	}

	for _, s := range inactive {
		if s.IsActive() {
			t.Errorf("Status(%q).IsActive() = true, want false", s)
		}
	}
}

func TestStatusIsBlocking(t *testing.T) {
	blocking := []Status{StatusBlocked, StatusOnHold, StatusCancelled}
	notBlocking := []Status{StatusDraft, StatusInProgress, StatusCompleted}

	for _, s := range blocking {
		if !s.IsBlocking() {
			t.Errorf("Status(%q).IsBlocking() = false, want true", s)
		}
	}

	for _, s := range notBlocking {
		if s.IsBlocking() {
			t.Errorf("Status(%q).IsBlocking() = true, want false", s)
		}
	}
}

func TestPriorityNormalize(t *testing.T) {
	tests := []struct {
		priority Priority
		want     Priority
	}{
		{PriorityP0, PriorityCritical},
		{PriorityP1, PriorityHigh},
		{PriorityP2, PriorityMedium},
		{PriorityP3, PriorityLow},
		{PriorityCritical, PriorityCritical},
		{PriorityHigh, PriorityHigh},
	}

	for _, tt := range tests {
		got := tt.priority.Normalize()
		if got != tt.want {
			t.Errorf("Priority(%q).Normalize() = %q, want %q", tt.priority, got, tt.want)
		}
	}
}

func TestPriorityLevel(t *testing.T) {
	tests := []struct {
		priority Priority
		want     int
	}{
		{PriorityCritical, 0},
		{PriorityP0, 0},
		{PriorityHigh, 1},
		{PriorityP1, 1},
		{PriorityMedium, 2},
		{PriorityLow, 3},
	}

	for _, tt := range tests {
		got := tt.priority.Level()
		if got != tt.want {
			t.Errorf("Priority(%q).Level() = %d, want %d", tt.priority, got, tt.want)
		}
	}
}
