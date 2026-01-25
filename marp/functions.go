package marp

import (
	"fmt"
	"html/template"
	"strings"
)

// CommonFuncMap provides template functions used across all Marp renderers.
// Import this into your template.FuncMap to ensure consistent behavior.
var CommonFuncMap = template.FuncMap{
	"add":             Add,
	"sub":             Sub,
	"mul":             Mul,
	"div":             Div,
	"truncate":        Truncate,
	"progressBar":     ProgressBar,
	"progressBarLen":  ProgressBarLen,
	"progressPercent": ProgressPercent,
	"scorePercent":    ScorePercent,
	"statusIcon":      StatusIcon,
	"statusEmoji":     StatusEmoji,
	"statusLabel":     StatusLabel,
	"priorityIcon":    PriorityIcon,
	"priorityLabel":   PriorityLabel,
	"severityIcon":    SeverityIcon,
	"join":            strings.Join,
	"upper":           strings.ToUpper,
	"lower":           strings.ToLower,
	"title":           strings.Title, //nolint:staticcheck
	"hasPrefix":       strings.HasPrefix,
	"hasSuffix":       strings.HasSuffix,
	"contains":        strings.Contains,
	"replace":         strings.ReplaceAll,
	"trim":            strings.TrimSpace,
	"default":         Default,
	"coalesce":        Coalesce,
	"ternary":         Ternary,
	"seq":             Seq,
}

// Add returns a + b.
func Add(a, b int) int {
	return a + b
}

// Sub returns a - b.
func Sub(a, b int) int {
	return a - b
}

// Mul returns a * b.
func Mul(a, b int) int {
	return a * b
}

// Div returns a / b. Returns 0 if b is 0.
func Div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// Truncate shortens a string to max characters, adding "..." if truncated.
func Truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max]
	}
	return s[:max-3] + "..."
}

// ProgressBar returns a text-based progress bar of default length 10.
// Progress should be between 0.0 and 1.0.
func ProgressBar(progress float64) string {
	return ProgressBarLen(progress, 10)
}

// ProgressBarLen returns a text-based progress bar of specified length.
// Progress should be between 0.0 and 1.0.
func ProgressBarLen(progress float64, length int) string {
	if progress < 0 {
		progress = 0
	}
	if progress > 1 {
		progress = 1
	}
	filled := int(progress * float64(length))
	empty := length - filled
	return strings.Repeat("█", filled) + strings.Repeat("░", empty)
}

// ScorePercent formats a 0.0-1.0 score as a percentage string.
func ScorePercent(score float64) string {
	return fmt.Sprintf("%.0f%%", score*100)
}

// ProgressPercent formats a 0.0-1.0 progress value as a percentage string.
// Alias for ScorePercent for semantic clarity.
func ProgressPercent(progress float64) string {
	return ScorePercent(progress)
}

// StatusIcon returns an emoji icon for common status values.
// Supports: completed, done, in_progress, active, planned, pending, future, blocked, at_risk.
func StatusIcon(status string) string {
	switch strings.ToLower(strings.ReplaceAll(status, " ", "_")) {
	case "completed", "done", "complete", "achieved":
		return "✅"
	case "in_progress", "active", "in-progress", "inprogress":
		return "🚧"
	case "planned", "pending", "todo", "planning":
		return "📋"
	case "future", "backlog", "idea", "proposed", "not_started":
		return "💡"
	case "blocked", "failed", "missed":
		return "🚫"
	case "at_risk", "at-risk", "atrisk", "warning", "behind":
		return "⚠️"
	case "cancelled", "canceled", "dropped":
		return "❌"
	case "on_hold", "on-hold", "onhold", "paused":
		return "⏸️"
	default:
		return "○"
	}
}

// StatusEmoji returns a text-based status indicator (not emoji) for use in
// contexts where emoji may not render well. Used by V2MOM and similar documents.
func StatusEmoji(status string) string {
	switch strings.ToLower(strings.ReplaceAll(status, " ", "_")) {
	case "completed", "done", "complete", "achieved":
		return "[DONE]"
	case "in_progress", "active", "in-progress", "inprogress":
		return "[IN PROGRESS]"
	case "at_risk", "at-risk", "atrisk", "behind":
		return "[AT RISK]"
	case "planning":
		return "[PLANNING]"
	case "not_started", "proposed", "pending":
		return "[NOT STARTED]"
	case "cancelled", "canceled", "missed":
		return "[CANCELLED]"
	case "on_hold", "on-hold", "onhold", "paused":
		return "[ON HOLD]"
	case "blocked":
		return "[BLOCKED]"
	default:
		return ""
	}
}

// StatusLabel returns a human-readable label for a status value.
// Converts snake_case/kebab-case to Title Case.
func StatusLabel(status string) string {
	s := strings.ReplaceAll(status, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	return strings.Title(s) //nolint:staticcheck
}

// PriorityIcon returns an icon for priority levels.
// Supports: p0/critical, p1/high, p2/medium, p3/low.
func PriorityIcon(priority string) string {
	switch strings.ToLower(priority) {
	case "p0", "critical":
		return "🔴"
	case "p1", "high":
		return "🟠"
	case "p2", "medium":
		return "🟡"
	case "p3", "low":
		return "🟢"
	default:
		return "⚪"
	}
}

// PriorityLabel returns a human-readable label for a priority value.
// Converts P0/P1/P2/P3 to descriptive labels.
func PriorityLabel(priority string) string {
	switch strings.ToUpper(priority) {
	case "P0":
		return "P0 (Critical)"
	case "P1":
		return "P1 (High)"
	case "P2":
		return "P2 (Medium)"
	case "P3":
		return "P3 (Low)"
	case "CRITICAL":
		return "Critical"
	case "HIGH":
		return "High"
	case "MEDIUM":
		return "Medium"
	case "LOW":
		return "Low"
	default:
		return priority
	}
}

// SeverityIcon returns an icon for severity levels (used in security contexts).
// Supports: critical, high, medium, low, informational.
func SeverityIcon(severity string) string {
	switch strings.ToLower(severity) {
	case "critical":
		return "🔴"
	case "high":
		return "🟠"
	case "medium":
		return "🟡"
	case "low":
		return "🟢"
	case "informational", "info":
		return "🔵"
	default:
		return "⚪"
	}
}

// Default returns the value if non-empty, otherwise returns the default.
func Default(defaultVal, val string) string {
	if val == "" {
		return defaultVal
	}
	return val
}

// Coalesce returns the first non-empty string from the arguments.
func Coalesce(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// Ternary returns trueVal if condition is true, otherwise falseVal.
func Ternary(condition bool, trueVal, falseVal any) any {
	if condition {
		return trueVal
	}
	return falseVal
}

// Seq generates a sequence of integers from start to end (inclusive).
func Seq(start, end int) []int {
	if end < start {
		return nil
	}
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = start + i
	}
	return result
}
