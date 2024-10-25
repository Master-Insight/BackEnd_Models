package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response encargado de gestionar las respuestas.
type Response struct {
	Context    echo.Context
	Estructura ResponseStruct
}

// New crea y devuelve una nueva instancia de Response.
func New(c echo.Context) *Response {
	params, _ := c.FormParams()

	// Estructura de respuesta inicial.
	return &Response{
		Context: c,
		Estructura: ResponseStruct{
			Data: nil,
			Metadata: Metadata{
				Request:    params,
				Origin:     c.Request().RequestURI,
				Pagination: nil,
			},
			State: State{
				Code:        http.StatusOK,
				Status:      http.StatusText(http.StatusOK),
				Message:     "Success",
				Description: "",
			},
		},
	}
}

// Send envía la respuesta con el código de estado y datos configurados.
func (r *Response) Send() error {
	return r.Context.JSON(r.Estructura.State.Code, r.Estructura)
} // ? Ejemplo: return res.New(c).Data("Success!").Send()

// * Métodos auxiliares para actualizar los datos básicos de la respuesta.

// Data setea la data de la respuesta
func (response *Response) Data(data interface{}) *Response {
	response.Estructura.Data = data
	return response
} // ? Ejemplo: return res.New(c).Data(map[string]string{"name": "John"}).Send()

// Message setea el mensaje de la respuesta
func (response *Response) Message(message string) *Response {
	response.Estructura.State.Message = message
	return response
}

// Description setea la descripción de la respuesta
func (response *Response) Description(description string) *Response {
	response.Estructura.State.Description = description
	return response
}

// Code setea el código de estado y el status de la respuesta
func (r *Response) Code(code int) *Response {
	r.Estructura.State.Code = code
	r.Estructura.State.Status = http.StatusText(code)
	return r
}

// * Métodos adicionales

// HideRequest oculta los parámetros de la solicitud en la respuesta
func (response *Response) HideRequest() *Response {
	response.Estructura.Metadata.Request = "hidden by API"
	return response
}

// Pagination setea la paginación en la respuesta
func (response *Response) Pagination(pagination interface{}) *Response {
	response.Estructura.Metadata.Pagination = pagination
	return response
}

// PageSize setea el tamaño de página en la paginación
func (response *Response) PageSize(pageSize int) *Response {
	response.Estructura.Metadata.PageSize = pageSize
	return response
}

// Page setea la página actual en la paginación
func (response *Response) Page(page int) *Response {
	response.Estructura.Metadata.Page = page
	return response
}

// Total setea el número total de elementos en la respuesta paginada
func (response *Response) Total(total int) *Response {
	response.Estructura.Metadata.Total = total
	return response
}
