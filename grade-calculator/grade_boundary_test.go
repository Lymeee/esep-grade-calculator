package esepunittests

import "testing"

func TestBoundary90isA(t *testing.T) {
gc := NewGradeCalculator()
gc.AddGrade("assign", 90, Assignment)
gc.AddGrade("exam", 90, Exam)
gc.AddGrade("essay", 90, Essay)
if got := gc.GetFinalGrade(); got != "A" {
t.Errorf("expected A at 90 cutoff; got %s", got)
}
}

func TestBoundary80isB(t *testing.T) {
gc := NewGradeCalculator()
gc.AddGrade("assign", 80, Assignment)
gc.AddGrade("exam", 80, Exam)
gc.AddGrade("essay", 80, Essay)
if got := gc.GetFinalGrade(); got != "B" {
t.Errorf("expected B at 80 cutoff; got %s", got)
}
}

func TestBoundary70isC(t *testing.T) {
gc := NewGradeCalculator()
gc.AddGrade("assign", 70, Assignment)
gc.AddGrade("exam", 70, Exam)
gc.AddGrade("essay", 70, Essay)
if got := gc.GetFinalGrade(); got != "C" {
t.Errorf("expected C at 70 cutoff; got %s", got)
}
}

func TestBoundary60isD(t *testing.T) {
gc := NewGradeCalculator()
gc.AddGrade("assign", 60, Assignment)
gc.AddGrade("exam", 60, Exam)
gc.AddGrade("essay", 60, Essay)
if got := gc.GetFinalGrade(); got != "D" {
t.Errorf("expected D at 60 cutoff; got %s", got)
}
}

func TestBelow60isF(t *testing.T) {
gc := NewGradeCalculator()
// 50%*50 + 35%*55 + 15%*40 = 50.25 -> F
gc.AddGrade("assign", 50, Assignment)
gc.AddGrade("exam", 55, Exam)
gc.AddGrade("essay", 40, Essay)
if got := gc.GetFinalGrade(); got != "F" {
t.Errorf("expected F below 60; got %s", got)
}
}

func TestEmptyCategoriesAreZero(t *testing.T) {
gc := NewGradeCalculator()
if got := gc.GetFinalGrade(); got != "F" {
t.Errorf("expected F when all categories empty; got %s", got)
}
}
