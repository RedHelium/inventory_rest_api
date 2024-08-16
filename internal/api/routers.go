package api

import (
	"net/http"
	"strconv"
	"task/internal/database"
	"task/internal/handlers"
	"task/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

/*
GET-запрос на получение оборудования с транзакциями по ID/Названию оборудования/ID владельца оборудования с фильтрацией по статусу транзакций
*/
func GetInventoryWithOperations(conn *pgx.Conn) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		id, err := strconv.Atoi(c.DefaultQuery("id", "0"))

		if handlers.HasHttpError(c, http.StatusBadRequest, "Ошибка. ID должен быть числовым значением!", err, nil) {
			return
		}

		status := c.Query("status")
		id_owner, err := strconv.Atoi(c.DefaultQuery("id_owner", "0"))

		if handlers.HasHttpError(c, http.StatusBadRequest, "Ошибка. ID владельца оборудования должен быть числовым значением!", err, nil) {
			return
		}

		name := c.Query("name")

		data, err := database.GetInventoryWithOperations(conn, id, id_owner, name, status)

		handlers.HasHttpError(c, http.StatusBadRequest, "Не удалось получить оборудование!", err, data)

		if data.Selected_Inventory.Name != "" {
			c.IndentedJSON(http.StatusOK, data)
		} else {
			status := http.StatusOK
			c.IndentedJSON(status, handlers.GetResponse("Оборудование не найдено!", status))
		}

	}

	return gin.HandlerFunc(fn)
}

/*
GET-запрос на получение транзакции с оборудованием с возможностью фильтрации по статусу и периоду
*/
func GetOperationWithInventory(conn *pgx.Conn) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		id_operaton, err := strconv.Atoi(c.DefaultQuery("id_operation", "0"))

		if handlers.HasHttpError(c, http.StatusBadRequest, "Ошибка. ID должен быть числовым значением!", err, nil) {
			return
		}

		src_executor, err := strconv.Atoi(c.DefaultQuery("src_executor", "0"))

		if handlers.HasHttpError(c, http.StatusBadRequest, "Ошибка. ID должен быть числовым значением!", err, nil) {
			return
		}

		dst_executor, err := strconv.Atoi(c.DefaultQuery("dst_executor", "0"))

		if handlers.HasHttpError(c, http.StatusBadRequest, "Ошибка. ID должен быть числовым значением!", err, nil) {
			return
		}

		status := c.Query("status")

		created_date_from := c.Query("created_date_from")
		created_date_to := c.Query("created_date_to")

		updated_date_from := c.Query("updated_date_from")
		updated_date_to := c.Query("updated_date_to")

		created_date := new(models.Period)
		updated_date := new(models.Period)

		created_date.Date_From, _ = time.Parse("2006-01-02", created_date_from)
		created_date.Date_To, _ = time.Parse("2006-01-02", created_date_to)

		updated_date.Date_From, _ = time.Parse("2006-01-02", updated_date_from)
		updated_date.Date_To, _ = time.Parse("2006-01-02", updated_date_to)

		data, err := database.GetOperationsWithInventory(conn, id_operaton, src_executor, dst_executor, status, *created_date, *updated_date)

		handlers.HasHttpError(c, http.StatusBadRequest, "Не удалось получить транзакции!", err, data)

		if len(*data) > 0 {
			c.IndentedJSON(http.StatusOK, data)
		} else {
			status := http.StatusOK
			c.IndentedJSON(status, handlers.GetResponse("Транзакций не найдено!", status))
		}

	}
	return gin.HandlerFunc(fn)
}

/*
GET-запрос на получение всего оборудования
*/
func GetAllInventory(conn *pgx.Conn) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		data, err := database.GetAllInventory(conn)

		handlers.HasHttpError(c, http.StatusBadRequest, "Не удалось получить оборудование!", err, data)

		if len(data) > 0 {
			c.IndentedJSON(http.StatusOK, data)
		} else {
			status := http.StatusOK
			c.IndentedJSON(status, handlers.GetResponse("Оборудование не найдено!", status))
		}

	}

	return gin.HandlerFunc(fn)
}
