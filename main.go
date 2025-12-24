package main

import (
	"CourseRegistry/registry"
	"fmt"
)

func main() {
	r := registry.NewRegistry()

	r.AddStudent(registry.Student{ID: 1, Name: "Alice"})
	r.AddStudent(registry.Student{ID: 2, Name: "Bob"})
	r.AddStudent(registry.Student{ID: 3, Name: "Charlie"})
	r.EnrollCourse(1, "Go")
	r.EnrollCourse(1, "Databases")
	r.EnrollCourse(2, "Go")

	for {
		fmt.Println("\n1. Add Student")
		fmt.Println("2. Enroll Course")
		fmt.Println("3. Remove Course")
		fmt.Println("4. List Students")
		fmt.Println("5. Course Statistics")
		fmt.Println("6. Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var id uint64
			var name string
			fmt.Print("ID: ")
			fmt.Scanln(&id)
			fmt.Print("Name: ")
			fmt.Scanln(&name)

			if err := r.AddStudent(registry.Student{ID: id, Name: name}); err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			var id uint64
			var course string
			fmt.Print("Student ID: ")
			fmt.Scanln(&id)
			fmt.Print("Course: ")
			fmt.Scanln(&course)

			if err := r.EnrollCourse(id, course); err != nil {
				fmt.Println("Error:", err)
			}
		case 3:
			var id uint64
			var course string
			fmt.Print("Student ID: ")
			fmt.Scanln(&id)
			fmt.Print("Course: ")
			fmt.Scanln(&course)

			if err := r.RemoveCourse(id, course); err != nil {
				fmt.Println("Error:", err)
			}
		case 4:
			for _, s := range r.ListStudents() {
				fmt.Printf("ID: %d | Name: %s | Courses: %v\n",
					s.ID, s.Name, s.Courses)
			}
		case 5:
			for c, count := range r.CoursesCount() {
				fmt.Println(c, "->", count)
			}
		case 6:
			return
		}
	}
}
