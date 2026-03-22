package repository

type Teacher struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Email   string `json:"email"`
}

type TeacherRepository interface {
	GetAllTeachers() ([]Teacher, error)
	GetTeacherByID(id string) (*Teacher, error)
}

func (r *Repository) GetAllTeachers() ([]Teacher, error) {
	// ตัวอย่าง mock data
	teachers := []Teacher{
		{
			Id:      "T001",
			Name:    "Mr. Smith",
			Subject: "Math",
			Email:   "smith@example.com",
		},
		{
			Id:      "T002",
			Name:    "Ms. Johnson",
			Subject: "Science",
			Email:   "johnson@example.com",
		},
	}
	return teachers, nil
}

func (r *Repository) GetTeacherByID(id string) (*Teacher, error) {
	all, _ := r.GetAllTeachers()
	for _, t := range all {
		if t.Id == id {
			return &t, nil
		}
	}
	return nil, nil
}
