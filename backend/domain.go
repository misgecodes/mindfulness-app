package main

type Topic struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Content     []Content `json:"content"`
}

type Content struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	TopicID string `json:"topic_id"`
	Type    string `json:"type"`
	Source  string `json:"source"`
	URI     string `json:"uri"`
	// Duration int    `json:"duration"`
}

type UserTopic struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	TopicID   string `json:"topic_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsActive  bool   `json:"is_active"`
}
