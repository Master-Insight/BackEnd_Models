package res

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// "Clase Response" encargada gestionar las respuestas
type Response struct {
	Context    echo.Context
	Estructura ResponseStruct
}

// "Clase Response" - Constructor - Crea y devuelve una instancia de Response
func New(c echo.Context) *Response {

	var response Response
	response.Context = c

	// Parámetros de la solicitud
	params, _ := c.FormParams()

	// Respuesta por defecto
	response.Estructura = ResponseStruct{
		Data: "",
		Metadata: Metadata{
			Request:    params,                 // Parámetros de la solicitud
			Origin:     c.Request().RequestURI, // URI de la solicitud
			Pagination: "",
		},
		State: State{
			Code:        http.StatusOK,                  // Código 200 por defecto
			Status:      http.StatusText(http.StatusOK), // Texto "OK"
			Message:     "",
			Description: "",
		},
	}

	return &response
}

// "Clase Response" - Metodo Send - Envia la respuesta
func (response *Response) Send() error {
	return response.Context.JSON(response.Estructura.State.Code, response.Estructura)
}

// * "Clase Response" - Metodos para modficar variables de la respuesta

// "Clase Response" - Metodo Data - setea la data de la respuesta
func (response *Response) Data(data interface{}) *Response {
	response.Estructura.Data = data
	return response
}

// "Clase Response" - Metodo Message - setea el mensaje de la respuesta
func (response *Response) Message(message string) *Response {
	response.Estructura.State.Message = message
	return response
}

// "Clase Response" - Metodo Description - setea el mensaje de la respuesta
func (response *Response) Description(description string) *Response {
	response.Estructura.State.Description = description
	return response
}

// "Clase Response" - Metodo Code - setea el código de respuesta
func (response *Response) Code(code int) *Response {
	response.Estructura.State.Code = code
	response.Estructura.State.Status = http.StatusText(code)
	return response
}

// "Clase Response" - Metodo HideRequest - envia un Request vacio
func (response *Response) HideRequest() *Response {
	response.Estructura.Metadata.Request = "hidden by api"
	return response
}

// "Clase Response" - Metodo Pagination - setea la paginacion de la respuesta
func (response *Response) Pagination(pagination interface{}) *Response {
	response.Estructura.Metadata.Pagination = pagination
	return response
}

// "Clase Response" - Metodo PageSize - setea la paginacion de la respuesta
func (response *Response) PageSize(pageSize int) *Response {
	response.Estructura.Metadata.PageSize = pageSize
	return response
}

// "Clase Response" - Metodo Page - setea la paginacion de la respuesta
func (response *Response) Page(page int) *Response {
	response.Estructura.Metadata.Page = page
	return response
}

// "Clase Response" - Metodo Total - setea la paginacion de la respuesta
func (response *Response) Total(total int) *Response {
	response.Estructura.Metadata.Total = total
	return response
}
