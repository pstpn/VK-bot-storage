package models

type (
	Login    string
	Password string
	ChatID   int64
)

type UserData struct {
	ChatID      ChatID   `json:"chat_id"`
	ServiceName string   `json:"service_name"`
	Login       Login    `json:"login"`
	Password    Password `json:"password"`
}

type SetUserDataInput struct {
	ServiceName string   `json:"service_name"`
	ChatID      ChatID   `json:"chat_id"`
	Login       Login    `json:"login"`
	Password    Password `json:"password"`
}

type GetUserDataInput struct {
	ServiceName string `json:"service_name"`
	ChatID      ChatID `json:"chat_id"`
}

type DelUserDataInput struct {
	ServiceName string `json:"service_name"`
	ChatID      ChatID `json:"chat_id"`
}
