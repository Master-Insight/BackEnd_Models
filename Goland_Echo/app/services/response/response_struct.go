package response

// ResponseStruct estructura de la respuesta JSON.
type ResponseStruct struct {
	Data     interface{} `json:"data"`
	Metadata Metadata    `json:"metadata"`
	State    State       `json:"state"`
}

// Metadata contiene información adicional sobre la solicitud y paginación.
type Metadata struct {
	Request    interface{} `json:"request"`
	Origin     string      `json:"origin"`
	Pagination interface{} `json:"pagination"`
	PageSize   int         `json:"page_size,omitempty"`
	Page       int         `json:"page,omitempty"`
	Total      int         `json:"total,omitempty"`
}

// State contiene el estado de la respuesta, como código y mensaje.
type State struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}
