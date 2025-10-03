package esepunittests

type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

// Public API expected by tests: returns letter grade with inclusive thresholds
func (gc *GradeCalculator) GetFinalGrade() string {
	numerical := gc.calculateNumericalGrade() // float64

	switch {
	case numerical >= 90.0:
		return "A"
	case numerical >= 80.0:
		return "B"
	case numerical >= 70.0:
		return "C"
	case numerical >= 60.0:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{Name: name, Grade: grade, Type: Assignment})
	case Exam:
		gc.exams = append(gc.exams, Grade{Name: name, Grade: grade, Type: Exam})
	case Essay:
		gc.essays = append(gc.essays, Grade{Name: name, Grade: grade, Type: Essay})
	}
}

// Weights: Assignments 50%, Exams 35%, Essays 15%.
// Empty categories contribute 0 to the average.
func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	assignmentAverage := computeAverage(gc.assignments)
	examAverage := computeAverage(gc.exams)
	essayAverage := computeAverage(gc.essays)

	const wA, wE, wS = 0.50, 0.35, 0.15
	return assignmentAverage*wA + examAverage*wE + essayAverage*wS
}

// computeAverage safely handles empty slices (returns 0).
func computeAverage(grades []Grade) float64 {
	n := len(grades)
	if n == 0 {
		return 0.0
	}
	sum := 0.0
	for _, g := range grades {
		sum += float64(g.Grade)
	}
	return sum / float64(n)
}
