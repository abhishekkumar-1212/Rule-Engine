package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	initDB()

	http.HandleFunc("/rules", createRuleHandler)         // Create rule
	http.HandleFunc("/rules/list", getRulesHandler)      // List rules
	http.HandleFunc("/rules/evaluate", evaluateRuleHandler) // Evaluate rule
	http.HandleFunc("/rules/combine", combineRulesHandler)   // Combine rules

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
