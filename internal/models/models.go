package models

import "time"

type (
	Login    string
	Password string
)

type Service struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UsersData struct {
	ID        int       `json:"user_id"`
	ServiceID int       `json:"service_id"`
	Login     Login     `json:"login"`
	Password  Password  `json:"password"`
	Time      time.Time `json:"time"`
}

type SetUsersDataInput struct {
	ServiceID int      `json:"service_id"`
	UserID    int      `json:"user_id"`
	Login     Login    `json:"login"`
	Password  Password `json:"password"`
}

type GetUsersDataInput struct {
	ServiceID int `json:"service_id"`
	UserID    int `json:"user_id"`
}

type GetServiceInput struct {
	ServiceName string `json:"service_name"`
}

type DelUsersDataInput struct {
	ServiceID int    `json:"service_name"`
	UserID    int    `json:"user_id"`
	Login     string `json:"login"`
}
