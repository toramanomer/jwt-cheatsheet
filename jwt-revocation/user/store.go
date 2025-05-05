package user

import "sync"

type UserStore struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]*User),
	}
}

func (s *UserStore) CreateUser(username string) (*User, error) {
	if user := s.GetUser(username); user != nil {
		return nil, ErrUserAlreadyExists
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	user := NewUser(username)
	s.users[username] = user

	return user, nil
}

func (s *UserStore) GetUser(username string) *User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[username]
	if !ok {
		return nil
	}
	return user
}
