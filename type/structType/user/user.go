package user

type UserInfoStruct struct {
	Email  string `db:"email" json:"email"`
	Name   string `db:"name" json:"name"`
	Phone  string `db:"phone" json:"phone"`
	JobID  int    `db:"job_id" json:"job_id"`
	Avatar string `db:"avatar" json:"avatar"`
	Role   string `db:"role" json:"role"`
}
