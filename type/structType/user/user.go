package user

type UserInfoStruct struct {
	Email  string `db:"email" json:"email"`
	Name   string `db:"name" json:"name"`
	Phone  string `db:"phone" json:"phone"`
	JobID  int    `db:"job_id" json:"job_id"`
	Avatar string `db:"avatar" json:"avatar"`
	Role   string `db:"role" json:"role"`
}

// 部門數據結構
type Node struct {
	ID       int     `db:"id" json:"id"`
	Key      string  `db:"key" json:"key"`
	Title    string  `db:"title" json:"title"`
	ParentID *int    `db:"parent_id" json:"parent_id"`
	Children []*Node `json:"children,omitempty"`
}
