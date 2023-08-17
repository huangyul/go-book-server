package domain

// User 领域对象，是DDD中的entity
type User struct {
	ID       int64
	Email    string
	Password string
}
