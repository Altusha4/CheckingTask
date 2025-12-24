package registry

import "fmt"

type Student struct {
	ID      uint64
	Name    string
	Courses []string
}

type Registry struct {
	Students map[uint64]Student
}

func NewRegistry() *Registry {
	return &Registry{
		Students: make(map[uint64]Student),
	}
}

func (r *Registry) AddStudent(student Student) error {
	if student.Name != "" {
		return fmt.Errorf("name cannot be empty")
	}
	if r.Students[student.ID].ID != 0 {
		return fmt.Errorf("id exists")
	}
	r.Students[student.ID] = student
	return nil
}

func (r *Registry) EnrollCourse(id uint64, course string) error {
	if course == "" {
		return fmt.Errorf("empty course")
	}

	s := r.Students[id]
	if s.ID == 0 {
		return fmt.Errorf("Student not found")
	}

	for _, c := range s.Courses {
		if c == course {
			return fmt.Errorf("already enrolled")
		}
	}

	s.Courses = append(s.Courses, course)
	r.Students[id] = s
	return nil
}

func (r *Registry) RemoveCourse(id uint64, course string) error {
	s := r.Students[id]
	if s.ID == 0 {
		return fmt.Errorf("Student not found")
	}

	for i, c := range s.Courses {
		if c == course {
			s.Courses = append(s.Courses[:i], s.Courses[i+1:]...)
			r.Students[id] = s
			return nil
		}
	}
	return fmt.Errorf("Course not found")
}

func (r *Registry) ListStudents() []Student {
	result := []Student{}
	for _, s := range r.Students {
		result = append(result, s)
	}
	return result
}

func (r *Registry) CoursesCount() map[string]int {
	m := make(map[string]int)

	for _, s := range r.Students {
		for _, c := range s.Courses {
			m[c]++
		}
	}
	return m
}
