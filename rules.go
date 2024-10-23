package main

import (
	"strings"
)

// NodeType defines the type of the AST node (operator or operand)
type NodeType string

const (
	Operator NodeType = "operator"
	Operand  NodeType = "operand"
)

// Node struct representing AST
type Node struct {
	Type  NodeType `json:"type"`
	Value string   `json:"value,omitempty"`
	Left  *Node    `json:"left,omitempty"`
	Right *Node    `json:"right,omitempty"`
}

// CreateRule: Takes a rule string and converts it into an AST
func CreateRule(ruleString string) *Node {
	if strings.Contains(ruleString, "AND") {
		parts := strings.Split(ruleString, "AND")
		return &Node{
			Type:  Operator,
			Value: "AND",
			Left:  CreateRule(strings.TrimSpace(parts[0])),
			Right: CreateRule(strings.TrimSpace(parts[1])),
		}
	}
	if strings.Contains(ruleString, "OR") {
		parts := strings.Split(ruleString, "OR")
		return &Node{
			Type:  Operator,
			Value: "OR",
			Left:  CreateRule(strings.TrimSpace(parts[0])),
			Right: CreateRule(strings.TrimSpace(parts[1])),
		}
	}
	// Leaf node (operand)
	return &Node{
		Type:  Operand,
		Value: ruleString,
	}
}

// EvaluateRule: Evaluates the rule AST against input data
func EvaluateRule(node *Node, data map[string]interface{}) bool {
	if node.Type == Operand {
		// Parse the condition (simplified for demo)
		return evaluateCondition(node.Value, data)
	}

	leftResult := EvaluateRule(node.Left, data)
	rightResult := EvaluateRule(node.Right, data)

	if node.Value == "AND" {
		return leftResult && rightResult
	}
	return leftResult || rightResult
}

// Helper to evaluate condition (simplified)
func evaluateCondition(condition string, data map[string]interface{}) bool {
	if condition == "age > 30" {
		return data["age"].(int) > 30
	}
	if condition == "department = Sales" {
		return data["department"] == "Sales"
	}
	return false
}

// Helper function to create operator node
func createOperatorNode(operatorVal string, leftNode, rightNode *Node) *Node {
	return &Node{
		Type:  Operator,
		Value: operatorVal,
		Left:  leftNode,
		Right: rightNode,
	}
}
