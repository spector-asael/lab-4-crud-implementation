// Filename: internal/data/comments.go
package data

import "time"

// each name begins with uppercase so that they are exportable/public
type Comment struct {
	ID        int64     `json:"id"`         // unique value for each comment
	Content   string    `json:"content"`    // the comment data
	Author    string    `json:"author"`     // the person who wrote the comment
	CreatedAt time.Time `json:"created_at"` // database timestamp
	Version   int32     `json:"version"`    // incremented on each update
}
