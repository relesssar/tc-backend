package tc

type User struct {
	Id         string `json:"-" db:"id"`
	Name       string `json:"name" db:"name" binding:"required"`
	Email      string `json:"email" db:"email" binding:"required"`
	Phone      string `json:"phone" db:"phone"`
	Password   string `json:"password" binding:"required"`
	ModuleName string `json:"module_name" db:"module_name"`
}
type UpdateUserInput struct {
	Id         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name" binding:"required"`
	Email      string `json:"email" db:"email" binding:"required"`
	Phone      string `json:"phone" db:"phone"`
	Password   string `json:"password"`
	ModuleName string `json:"module_name" db:"module_name"`
}
type UsersList struct {
	Id         string     `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	Email      string     `json:"email" db:"email"`
	Phone      string     `json:"phone" db:"phone"`
	ModuleName NullString `json:"module_name" db:"module_name" swaggertype:"string"`
}

type UserModule struct {
	Id         string `json:"-" db:"id"`
	UserId     string `json:"user_id" db:"user_id"`
	ModuleName string `json:"module_name" db:"module_name"`
	ModuleDesc string `json:"module_desc" db:"module_desc"`
}
