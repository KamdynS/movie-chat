package model

type Room struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
