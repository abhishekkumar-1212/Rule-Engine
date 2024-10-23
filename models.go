package main

import (
	"encoding/json"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Rule represents the structure of the rule in the database
type Rule struct {
	ID          uint            `gorm:"primaryKey"`
	RuleName    string          `json:"rule_name"`
	RuleAST     json.RawMessage `json:"rule_ast"` // Store AST as JSON
	Description string          `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
}

var db *gorm.DB

// Initialize the database
func initDB() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=rules_db port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	// Migrate the schema
	db.AutoMigrate(&Rule{})
}
