package service

import (
	"errors"
)

// Admin model (mock)
type Admin struct {
	ID       int
	Username string
	Password string
}

// AdminService interface
type AdminService interface {
	CreateAdmin(username, password string) (*Admin, error)
	GetAdminByID(id int) (*Admin, error)
	DeleteAdmin(id int) error
}

// implementation
type adminService struct {
	admins map[int]*Admin
	nextID int
}

// constructor
func NewAdminService() AdminService {
	return &adminService{
		admins: make(map[int]*Admin),
		nextID: 1,
	}
}

// Create admin
func (s *adminService) CreateAdmin(username, password string) (*Admin, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password required")
	}

	admin := &Admin{
		ID:       s.nextID,
		Username: username,
		Password: password,
	}

	s.admins[s.nextID] = admin
	s.nextID++

	return admin, nil
}

// Get admin by ID
func (s *adminService) GetAdminByID(id int) (*Admin, error) {
	admin, ok := s.admins[id]
	if !ok {
		return nil, errors.New("admin not found")
	}
	return admin, nil
}

// Delete admin
func (s *adminService) DeleteAdmin(id int) error {
	_, ok := s.admins[id]
	if !ok {
		return errors.New("admin not found")
	}
	delete(s.admins, id)
	return nil
}
