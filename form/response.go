package form

type Response struct {
	MetaData MetaData    `json:"metadata,omitempty,omitzero"`
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	Items    interface{} `json:"items,omitempty,omitzero"`
	Error    interface{} `json:"error,omitempty,omitzero"`
	Token    string      `json:"token,omitempty,omitzero"`
}

type MetaData struct {
	Total interface{} `json:"total_row,omitempty,omitzero"`
}
