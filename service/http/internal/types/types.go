// Code generated by goctl. DO NOT EDIT.
package types

type MessageRequest struct {
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
}

type MessageReply struct {
	Reply string `json:"reply"`
}
