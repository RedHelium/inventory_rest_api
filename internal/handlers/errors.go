package handlers

import (
	"fmt"
	"net/http"
	"task/internal/models"

	"github.com/gin-gonic/gin"
)

/*
Вывод ошибки в консоль
*/
func HasError(err error, message string, args ...any) {

	if err != nil {
		fmt.Printf(message, args...)
	}
}

/*
Сигнализирует о наличии ошибки и возвращает в ответе на запрос JSON с кодом ошибки и её описанием
*/
func HasHttpError(ctx *gin.Context, HttpStatusErrorCode int, message string, err error, data any) (hasError bool) {

	hasError = false

	if err != nil {
		ctx.IndentedJSON(HttpStatusErrorCode, GetResponse(message, http.StatusInternalServerError, err))
		hasError = true
	}

	return hasError
}

/*
Формирует ответ со статусом и сообщением
*/
func GetResponse(message string, httpStatusCode int, args ...any) models.MessageResponse {

	response := new(models.MessageResponse)
	response.Status = httpStatusCode
	response.Message = fmt.Sprintf(message, args...)

	return *response
}
