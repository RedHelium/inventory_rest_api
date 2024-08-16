package models

import "time"

type Inventory struct {
	ID    int
	Name  string
	Owner int
}

type InventoryOperation struct {
	ID                   int
	Source_Executor      int
	Destination_Executor int
	Request_Time         *time.Time
	Status_Time          *time.Time
	Status               string
}

type InventoryWithOperations struct {
	Selected_Inventory Inventory
	Operations         []InventoryOperation
}

type OperationWithInventory struct {
	Operation          InventoryOperation
	Invetory_Operation Inventory
}

type MessageResponse struct {
	Status  int
	Message string
}

type Period struct {
	Date_From time.Time
	Date_To   time.Time
}
