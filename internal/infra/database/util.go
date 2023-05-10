package database

import (
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateTableAndConnectionBD[T entity.Expense | entity.ExpenseOrigin |
	entity.ExpenseLevel | entity.User | entity.Income | entity.Period](nameTable T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(nameTable)
	db.AutoMigrate(nameTable)
	return db
}
