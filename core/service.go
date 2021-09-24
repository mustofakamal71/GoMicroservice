package core

import "github.com/mustofakamal71/GoMicroservice/MyMicroservices"

type service struct {
	// a database dependency would go here but instead we're going to have a static map
	m map[int64]MyMicroservices.User
}

// NewService instantiates a new Service.
func NewService( /* a database connection would be injected here */ ) MyMicroservices.Service {
	return &service{
		m: map[int64]MyMicroservices.User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
			3: {ID: 3, Name: "Carol"},
		},
	}
}
func (s *service) GetUser(id int64) (result MyMicroservices.User, err error) {
	// instead of querying a database, we just query our static map
	if result, ok := s.m[id]; ok {
		return result, nil
	}
	return result, MyMicroservices.ErrNotFound
}
func (s *service) GetUsers(ids []int64) (result map[int64]MyMicroservices.User, err error) {
	// always a good idea to return non-nil maps to avoid nil pointer dereferences
	result = map[int64]MyMicroservices.User{}
	for _, id := range ids {
		if u, ok := s.m[id]; ok {
			result[id] = u
		}
	}
	return
}