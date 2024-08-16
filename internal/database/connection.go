package database

import (
	"context"
	"task/internal/config"
	"task/internal/handlers"
	"task/internal/models"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
)

/*
Инициализирует подключение к БД
*/
func Connect(cfg *config.Config) *pgx.Conn {

	conn, err := pgx.Connect(context.Background(), config.GetConnectionString(cfg))

	handlers.HasError(err, "Не удалось подключится к базе данных! %s", err)

	return conn
}

/*
Общий метод на выполнение SQL-запроса
*/
func query(conn *pgx.Conn, sql string, args ...any) (pgx.Rows, error) {

	rows, err := conn.Query(context.Background(), sql, args...)

	handlers.HasError(err, "Ошибка выполнения SQL-запроса! %s", err)

	return rows, err
}

/*
Запрос на получение оборудования с транзакциями по ID/Названию оборудования/ID владельца оборудования с фильтрацией по статусу транзакций
*/
func GetInventoryWithOperations(conn *pgx.Conn, id int, id_owner int, name string, status string) (*models.InventoryWithOperations, error) {

	var err error

	result := new(models.InventoryWithOperations)

	result.Selected_Inventory = getInventory(conn, id, id_owner, name)
	result.Operations = getOperationsByInventory(conn, result.Selected_Inventory.ID, status)

	return result, err
}

/*
SQL-запрос на получение оборудования
*/
func getInventory(conn *pgx.Conn, id int, id_owner int, name string) models.Inventory {

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "\"name\"", "id_executor")
	sb.From("inventory")

	var inventoryRows pgx.Rows
	var err error
	var expr string

	if id_owner != 0 {

		expr = sb.Equal("id_executor", id_owner)

	} else if name != "" {

		expr = sb.Equal("name", name)

	} else if id != 0 {

		expr = sb.Equal("id", id)

	}

	sb.Where(expr)
	sqlResult, args := sb.Build()

	inventoryRows, err = query(conn, sqlResult, args...)

	handlers.HasError(err, "Ошибка выполнения запроса! %s", err)

	inventory, err := pgx.CollectOneRow(inventoryRows, pgx.RowToStructByPos[models.Inventory])

	handlers.HasError(err, "Ошибка получения оборудования! %s", err)

	return inventory
}

/*
SQL-запрос на получение транзакций по ID оборудования
*/
func getOperationsByInventory(conn *pgx.Conn, id_inventory int, status string) []models.InventoryOperation {

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("o.id", "src_executor", "dst_executor", "request_time", "status_time", "status")
	sb.From("inventory_operations_detail AS od")
	sb.Join("inventory_operations AS o", "o.id = od.id_inventory_operation")

	var operationsRows pgx.Rows
	var err error
	var expr string

	if id_inventory != 0 {

		expr = sb.Equal("id_inventory", id_inventory)
	}

	if status != "" {

		if expr != "" {

			expr = sb.And(expr, sb.Equal("status", status))

		} else {

			expr = sb.Equal("status", status)
		}

	}

	sb.Where(expr)
	sqlResult, args := sb.Build()

	operationsRows, err = query(conn, sqlResult, args...)

	handlers.HasError(err, "Ошибка выполнения запроса! %s", err)

	operations, err := pgx.CollectRows(operationsRows, pgx.RowToStructByPos[models.InventoryOperation])

	handlers.HasError(err, "Ошибка получения транзакций! %s", err)

	return operations
}

/*
Запрос на получение транзакций с оборудованием с фильтрацией по статусу и периоду
*/
func GetOperationsWithInventory(conn *pgx.Conn, id_operation int, src_executor int, dst_executor int, status string, created_date models.Period, updated_date models.Period) (*[]models.OperationWithInventory, error) {

	var err error

	var result []models.OperationWithInventory

	operations := getOperations(conn, id_operation, src_executor, dst_executor, status, created_date, updated_date)

	for i := 0; i < len(operations); i++ {

		tempOperation := new(models.OperationWithInventory)

		inventory := getInventoryByOperation(conn, operations[i].ID)

		tempOperation.Operation = operations[i]
		tempOperation.Invetory_Operation = inventory

		result = append(result, *tempOperation)
	}

	return &result, err
}

/*
SQL-запрос на получение транзакций
*/
func getOperations(conn *pgx.Conn, id_operation int, src_executor int, dst_executor int, status string, created_date models.Period, updated_date models.Period) []models.InventoryOperation {

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "src_executor", "dst_executor", "request_time", "status_time", "status")
	sb.From("public.inventory_operations")

	var expr string
	var operationsRows pgx.Rows
	var err error

	if id_operation != 0 {

		expr = sb.Equal("id", id_operation)

	} else if src_executor != 0 {

		expr = sb.Equal("src_executor", src_executor)

	} else if dst_executor != 0 {

		expr = sb.Equal("dst_executor", dst_executor)

	}

	if status != "" {

		if expr != "" {
			expr = sb.And(expr, sb.Equal("status", status))
		} else {
			expr = sb.Equal("status", status)
		}

	}

	if !created_date.Date_From.IsZero() && !created_date.Date_To.IsZero() {

		if expr != "" {

			expr = sb.And(expr, sb.Between("request_time", created_date.Date_From, created_date.Date_To))
		} else {
			expr = sb.Between("request_time", created_date.Date_From, created_date.Date_To)
		}

	}

	if !updated_date.Date_From.IsZero() && !updated_date.Date_To.IsZero() {

		if expr != "" {
			expr = sb.And(expr, sb.Between("status_time", updated_date.Date_From, updated_date.Date_To))
		} else {
			expr = sb.Between("status_time", updated_date.Date_From, updated_date.Date_To)
		}

	}

	if expr != "" {
		sb.Where(expr)
	}
	sqlResult, args := sb.Build()

	operationsRows, err = query(conn, sqlResult, args...)

	handlers.HasError(err, "Ошибка выполнения запроса! %s", err)

	operations, err := pgx.CollectRows(operationsRows, pgx.RowToStructByPos[models.InventoryOperation])

	handlers.HasError(err, "Ошибка получения транзакций! %s", err)

	return operations
}

/*
SQL-запрос на получение оборудования по транзакции
*/
func getInventoryByOperation(conn *pgx.Conn, id int) models.Inventory {

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("i.id", "\"name\"", "id_executor")
	sb.From("inventory_operations_detail as od ")
	sb.Join("public.inventory AS i", "i.id = od.id_inventory")
	sb.Where(sb.Equal("od.id_inventory_operation", id))
	sqlResult, args := sb.Build()

	inventoryRows, err := query(conn, sqlResult, args...)

	handlers.HasError(err, "Ошибка выполнения запроса! %s", err)

	inventory, err := pgx.CollectOneRow(inventoryRows, pgx.RowToStructByPos[models.Inventory])

	handlers.HasError(err, "Ошибка получения оборудования! %s", err)

	return inventory
}

/*
SQL-запрос на получение оборудования
*/
func GetAllInventory(conn *pgx.Conn) ([]models.Inventory, error) {

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "\"name\"", "id_executor")
	sb.From("inventory")

	var inventoryRows pgx.Rows
	var err error

	sqlResult, args := sb.Build()

	inventoryRows, err = query(conn, sqlResult, args...)

	handlers.HasError(err, "Ошибка выполнения запроса! %s", err)

	inventory, err := pgx.CollectRows(inventoryRows, pgx.RowToStructByPos[models.Inventory])

	handlers.HasError(err, "Ошибка получения оборудования! %s", err)

	return inventory, err
}
