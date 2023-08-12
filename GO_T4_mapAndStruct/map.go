package mapAndStruct

import "fmt"

func GetStudentSchools() map[string]string {
	studentSchools := map[string]string{
		"Zoe":     "KGV",
		"Conner":  "KGV",
		"Matthew": "RCHK",
		"Ryan":    "HKIS",
		"Irene":   "HKIS",
	}
	return studentSchools
}

func PrintMap(dict map[any]any) {
	for k, v := range dict {
		fmt.Printf("%s - %s", k, v)
	}
}
