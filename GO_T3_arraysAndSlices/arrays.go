package arraysAndSlices

/*
The GetStudentNames function returns an array of five strings that represent the names of some students.
@return An array of five strings representing the names of some students.
*/
// Note: Array in Go is:
// * must specify length
// * must specify element types of array
// * length of array is pre-defined and immutable
// * optionally can initialize the elements
func GetStudentNames() [5]string {
	names := [5]string{"Zoe", "Connect", "Matthew", "Ryan"}
	names[4] = "Irene"
	return names
}
