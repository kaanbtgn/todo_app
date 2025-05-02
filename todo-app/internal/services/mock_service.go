package services

import (
	"log"
	"sync"
	"time"
	"todo-app/internal/models"
)

type MockService struct {
	users     []models.User
	todoLists []models.TodoList
	todoItems []models.TodoItem
	mu        sync.RWMutex
}

func NewMockService() *MockService {
	// Initialize with some test data
	users := []models.User{
		{ID: 1, Username: "user1", Password: "password1", Role: "user"},
		{ID: 2, Username: "admin", Password: "admin123", Role: "admin"},
	}

	return &MockService{
		users: users,
	}
}

// User methods
func (s *MockService) GetUserByUsername(username string) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, nil
}

// TodoList methods
func (s *MockService) CreateTodoList(list *models.TodoList) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	list.ID = len(s.todoLists) + 1
	list.CreatedAt = time.Now()
	list.UpdatedAt = time.Now()
	list.CompletionPercentage = 0

	s.todoLists = append(s.todoLists, *list)
	return nil
}

func (s *MockService) GetTodoLists(userID int, isAdmin bool) ([]models.TodoList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var lists []models.TodoList
	for _, list := range s.todoLists {
		if list.DeletedAt == nil && (isAdmin || list.UserID == userID) {
			lists = append(lists, list)
		}
	}
	return lists, nil
}

func (s *MockService) GetTodoList(id, userID int, isAdmin bool) (*models.TodoList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, list := range s.todoLists {
		if list.ID == id && list.DeletedAt == nil && (isAdmin || list.UserID == userID) {
			return &list, nil
		}
	}
	return nil, nil
}

func (s *MockService) UpdateTodoList(list *models.TodoList) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, l := range s.todoLists {
		if l.ID == list.ID {
			list.UpdatedAt = time.Now()
			s.todoLists[i] = *list
			return nil
		}
	}
	return nil
}

func (s *MockService) DeleteTodoList(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for i, list := range s.todoLists {
		if list.ID == id {
			s.todoLists[i].DeletedAt = &now
			return nil
		}
	}
	return nil
}

// TodoItem methods
func (s *MockService) CreateTodoItem(item *models.TodoItem) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Debug logging
	log.Printf("Creating todo item in service: %+v", item)

	item.ID = len(s.todoItems) + 1
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	s.todoItems = append(s.todoItems, *item)

	// Debug logging
	log.Printf("Todo items after creation: %+v", s.todoItems)

	return nil
}

func (s *MockService) GetTodoItems(listID int) ([]models.TodoItem, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var items []models.TodoItem
	for _, item := range s.todoItems {
		if item.ListID == listID && item.DeletedAt == nil {
			items = append(items, item)
		}
	}
	return items, nil
}

func (s *MockService) UpdateTodoItem(item *models.TodoItem) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, it := range s.todoItems {
		if it.ID == item.ID {
			item.UpdatedAt = time.Now()
			s.todoItems[i] = *item
			return nil
		}
	}
	return nil
}

func (s *MockService) DeleteTodoItem(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for i, item := range s.todoItems {
		if item.ID == id {
			s.todoItems[i].DeletedAt = &now
			return nil
		}
	}
	return nil
}
