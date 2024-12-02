package user

import "time"

type User struct {
	ID            int64      `db:"id"`
	Email         string     `db:"email"`
	Password      string     `db:"password"`
	RoleID        int        `db:"role_id"`
	LastLogin     *time.Time `db:"last_login"`
	SentEmails    int        `db:"sent_emails"`
	LastEmailTime *time.Time `db:"last_email_time"`
	CreatedBy     int64      `db:"created_by"`
	UpdatedBy     *int       `db:"updated_by"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type CreateAdminRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type BulkCreateUserRequest struct {
	// Users []struct {
	// 	Email    string `json:"email" validate:"required,email"`
	// 	Password string `json:"password" validate:"required,min=6"`
	// } `json:"users" validate:"required,dive"`
	BaseName string `json:"base_name"`
	Domain   string `json:"domain"`
	Password string `json:"password" validate:"required,min=6"`
	Quantity int    `json:"quantity" validate:"required"`
	SendTo   string `json:"send_to" validate:"required"`
}

type PaginatedUsers struct {
	Users      []User `json:"users"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
	TotalPages int    `json:"total_pages"`
}
