package telegram

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}
type Update struct {
	ID      int    `json:"update_id"`
	Nessage string `json:"message"`
}
