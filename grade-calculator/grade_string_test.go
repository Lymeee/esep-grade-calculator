package esepunittests

import "testing"

func TestGradeTypeString(t *testing.T) {
if Assignment.String() != "assignment" {
t.Errorf("Assignment.String() = %q, want %q", Assignment.String(), "assignment")
}
if Exam.String() != "exam" {
t.Errorf("Exam.String() = %q, want %q", Exam.String(), "exam")
}
if Essay.String() != "essay" {
t.Errorf("Essay.String() = %q, want %q", Essay.String(), "essay")
}
}
