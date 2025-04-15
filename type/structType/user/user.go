package user

type UserInfoStruct struct {
	Id       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Name     string `db:"name" json:"name"`
	Phone    string `db:"phone" json:"phone"`
	JobID    int    `db:"job_id" json:"job_id"`
	Avatar   string `db:"avatar" json:"avatar"`
	Role     string `db:"role" json:"role"`
	Status   bool   `db:"status" json:"status"`
	Org      string `db:"org" json:"org"`
	UserName string `db:"username" json:"username"`
	Position string `db:"position" json:"position"`
	Gender   string `db:"gender" json:"gender"`
}

// 部門數據結構
type Node struct {
	ID       int     `db:"id" json:"id"`
	Key      string  `db:"key" json:"key"`
	Title    string  `db:"title" json:"title"`
	ParentID *int    `db:"parent_id" json:"parent_id"`
	Children []*Node `json:"children,omitempty"`
}
