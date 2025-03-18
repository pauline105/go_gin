package login

// LoginType 登錄用戶結構
type LoginType struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User 數據庫用戶表結構
type User struct {
	ID        int    `db:"id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	ProfileID int    `db:"profile_id"`
}
