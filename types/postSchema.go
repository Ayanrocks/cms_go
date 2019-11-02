package types

// Post -> Schema for db
type Post struct {
	// ID      int64  `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	// CreatedAt  time.Time `json:"created_at"`
	// UpdatedAt  time.Time `json:"updated_at"`
	ImageLinks []string `json:"image_links"`
}
