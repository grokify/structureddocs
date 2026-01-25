package marp

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{100, 200, 300},
	}

	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		s    string
		max  int
		want string
	}{
		{"hello", 10, "hello"},
		{"hello world", 8, "hello..."},
		{"hi", 2, "hi"},
		{"test", 3, "tes"},
		{"", 5, ""},
	}

	for _, tt := range tests {
		got := Truncate(tt.s, tt.max)
		if got != tt.want {
			t.Errorf("Truncate(%q, %d) = %q, want %q", tt.s, tt.max, got, tt.want)
		}
	}
}

func TestProgressBar(t *testing.T) {
	tests := []struct {
		progress float64
		want     string
	}{
		{0.0, "░░░░░░░░░░"},
		{0.5, "█████░░░░░"},
		{1.0, "██████████"},
		{0.3, "███░░░░░░░"},
		{-0.1, "░░░░░░░░░░"}, // Clamped to 0
		{1.5, "██████████"},  // Clamped to 1
	}

	for _, tt := range tests {
		got := ProgressBar(tt.progress)
		if got != tt.want {
			t.Errorf("ProgressBar(%v) = %q, want %q", tt.progress, got, tt.want)
		}
	}
}

func TestStatusIcon(t *testing.T) {
	tests := []struct {
		status string
		want   string
	}{
		{"completed", "✅"},
		{"done", "✅"},
		{"in_progress", "🚧"},
		{"in-progress", "🚧"},
		{"planned", "📋"},
		{"future", "💡"},
		{"blocked", "🚫"},
		{"at_risk", "⚠️"},
		{"cancelled", "❌"},
		{"on_hold", "⏸️"},
		{"unknown", "○"},
	}

	for _, tt := range tests {
		got := StatusIcon(tt.status)
		if got != tt.want {
			t.Errorf("StatusIcon(%q) = %q, want %q", tt.status, got, tt.want)
		}
	}
}

func TestPriorityIcon(t *testing.T) {
	tests := []struct {
		priority string
		want     string
	}{
		{"p0", "🔴"},
		{"critical", "🔴"},
		{"p1", "🟠"},
		{"high", "🟠"},
		{"p2", "🟡"},
		{"medium", "🟡"},
		{"p3", "🟢"},
		{"low", "🟢"},
		{"unknown", "⚪"},
	}

	for _, tt := range tests {
		got := PriorityIcon(tt.priority)
		if got != tt.want {
			t.Errorf("PriorityIcon(%q) = %q, want %q", tt.priority, got, tt.want)
		}
	}
}

func TestScorePercent(t *testing.T) {
	tests := []struct {
		score float64
		want  string
	}{
		{0.0, "0%"},
		{0.5, "50%"},
		{1.0, "100%"},
		{0.75, "75%"},
	}

	for _, tt := range tests {
		got := ScorePercent(tt.score)
		if got != tt.want {
			t.Errorf("ScorePercent(%v) = %q, want %q", tt.score, got, tt.want)
		}
	}
}

func TestSeq(t *testing.T) {
	got := Seq(1, 5)
	want := []int{1, 2, 3, 4, 5}

	if len(got) != len(want) {
		t.Fatalf("Seq(1, 5) returned %d elements, want %d", len(got), len(want))
	}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("Seq(1, 5)[%d] = %d, want %d", i, v, want[i])
		}
	}

	// Test empty sequence
	empty := Seq(5, 1)
	if empty != nil {
		t.Errorf("Seq(5, 1) = %v, want nil", empty)
	}
}

func TestCoalesce(t *testing.T) {
	tests := []struct {
		values []string
		want   string
	}{
		{[]string{"", "", "c"}, "c"},
		{[]string{"a", "b", "c"}, "a"},
		{[]string{"", "", ""}, ""},
		{[]string{}, ""},
	}

	for _, tt := range tests {
		got := Coalesce(tt.values...)
		if got != tt.want {
			t.Errorf("Coalesce(%v) = %q, want %q", tt.values, got, tt.want)
		}
	}
}
