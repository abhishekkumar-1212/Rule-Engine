package main

import (
	"encoding/json"
	"net/http"
)

// Handler to create a new rule
func createRuleHandler(w http.ResponseWriter, r *http.Request) {
	var newRule Rule
	err := json.NewDecoder(r.Body).Decode(&newRule)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Save the rule to the database
	if err := db.Create(&newRule).Error; err != nil {
		http.Error(w, "Failed to save rule", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newRule)
}

// Handler to list all rules
func getRulesHandler(w http.ResponseWriter, r *http.Request) {
	var rules []Rule
	if err := db.Find(&rules).Error; err != nil {
		http.Error(w, "Failed to retrieve rules", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rules)
}

// Handler to combine rules based on their IDs
func combineRulesHandler(w http.ResponseWriter, r *http.Request) {
	var ruleIDs []int
	err := json.NewDecoder(r.Body).Decode(&ruleIDs)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Logic to combine rules from the database using ruleIDs
	// Here you need to retrieve rules, merge their AST, and return or save the new rule

	// Placeholder response:
	w.Write([]byte("Combine rules logic goes here"))
}

// Handler to evaluate rules against user data
func evaluateRuleHandler(w http.ResponseWriter, r *http.Request) {
	var inputData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Fetch the rule you want to evaluate from the database
	// Example: var rule Rule
	// db.First(&rule, <rule_id>)

	// Placeholder response:
	evaluationResult := true // Replace with actual evaluation logic
	json.NewEncoder(w).Encode(map[string]bool{"result": evaluationResult})
}
