package appfoundation

type TransactionType string

const (
	TTypeMigration  TransactionType = "migration"
	TTypeAllocation TransactionType = "allocation"
)
