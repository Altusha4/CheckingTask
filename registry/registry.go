package registry

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
