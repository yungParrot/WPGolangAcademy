package academy

import (
  "math"
)

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
  noOfGrades := len(grades)
  if noOfGrades == 0 {
    return 0
  }
  sum := 0.0
  for _, grade := range grades {
    sum = sum + float64(grade)
  }
  avg := sum / float64(noOfGrades)
  roundedAvg := math.Round(avg)
  return int(roundedAvg)
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
  noOfAtt := len(attendance)
  if noOfAtt == 0 {
    return 0
  }
  sum := 0.0
  for _, attendance := range attendance {
    if attendance {
      sum = sum + 1.0
    }
  }
  attendancePercentage := sum / float64(noOfAtt)
  return attendancePercentage
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
  projectGrade := float64(s.Project)
  averageGrade := float64(AverageGrade(s.Grades))
  attendancePercentage := AttendancePercentage(s.Attendance)
  
  finalGrade := (projectGrade + averageGrade) / 2
  switch {
  case projectGrade == 1.0:
    finalGrade = 1.0
  case averageGrade == 1.0:
    finalGrade = 1.0
  case attendancePercentage < 0.6:
    finalGrade = 1.0
  case attendancePercentage < 0.8 && finalGrade > 1.0:
    finalGrade = finalGrade - 1.0
  }
  roundedFinalGrade := math.Round(finalGrade)
  return int(roundedFinalGrade)
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
  gradedStudents := make(map[string]uint8)

  for _, student := range students {
    gradedStudents[student.Name] = uint8(FinalGrade(student))
  }

  return gradedStudents
}
