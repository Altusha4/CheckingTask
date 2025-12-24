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
