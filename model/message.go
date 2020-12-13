package model

import "time"

// Message represents the message send by client sent
type Message struct {
	ID       string                 `json:"id" bson:"_id"`
	Msg      string                 `json:"msg" bson:"msg"`
	Room     string                 `json:"room" bson:"room"`
	LoginID  string                 `json:"login_id" bson:"login_id"`
	Type     string                 `json:"type" bson:"type"`
	Metadata map[string]interface{} `json:"metadata,omitempty" bson:"metadata"`
	CreateAt time.Time              `json:"created_at,omitempty" bson:"created_at"`
}
