package conditionalIterative

/*
The LegalToDrinkAlcohol function takes an age integer as an argument and returns a boolean indicating whether the person of that age is legally allowed to drink alcohol.
* @param age The age of the person.
* @return true if the person is legally allowed to drink alcohol, false otherwise.
*/
// Note: variable expression in if statement / switch statement
// * nothing different from locally declaring variable 'koreanAge' however,
// * it bears a semantic meaning that koreanAge is used for this if block only.
func LegalToDrinkAlcohol(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}
