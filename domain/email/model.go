package email

import "time"

type Email struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	SenderEmail string    `db:"sender_email"`
	SenderName  string    `db:"sender_name"`
	Subject     string    `db:"subject"`
	Preview     string    `db:"preview"`
	Body        string    `db:"body"`
	EmailType   string    `db:"email_type"`
	Attachments string    `db:"attachments"` // JSON format
	Timestamp   time.Time `db:"timestamp"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type SendEmailRequest struct {
	UserID      int          `json:"user_id"`
	To          string       `json:"to" validate:"required,email"`
	Subject     string       `json:"subject"`
	Body        string       `json:"body"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

// Convert timestamps to relative time
type EmailResponse struct {
	Email
	RelativeTime string `json:"RelativeTime"`
}
