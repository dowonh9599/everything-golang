package arraysAndSlices

/*
The GetTutoringSubjects function returns a slice of strings that represent the subjects a tutor can teach.
@return A slice of strings representing the subjects a tutor can teach.
*/
// Tips: Slice in Go is just an array in Go without the length specified
// * scalable length
func GetTutoringSubjects() []string {
	subjects := []string{"Math", "Physics"}
	subjects = append(subjects, "Computer Science")
	subjects = append(subjects, "Mandarin")
	return subjects
}
