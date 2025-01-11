package main

type Student struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Grade int8   `json:"grade"`
}

var students = []*Student{}

func init() {
	students = append(students, &Student{Id: "S001", Name: "Fani alfirdaus", Grade: 3})
	students = append(students, &Student{Id: "S002", Name: "Burneo", Grade: 4})
	students = append(students, &Student{Id: "S003", Name: "Joko Widodo", Grade: 5})
	students = append(students, &Student{Id: "S004", Name: "Agus Salim", Grade: 6})
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, student := range students {
		if student.Id == id {
			return student
		}
	}
	return nil
}
