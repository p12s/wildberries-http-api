package common

type User struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	// остальные поля добавил в дополнение к задаче, для авторизации
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
