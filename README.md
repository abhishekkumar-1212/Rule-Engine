
Rule Engine with AST
Objective
The goal of this project is to develop a simple 3-tier rule engine application to determine user eligibility based on various attributes like age, department, income, spend, etc. The system uses an Abstract Syntax Tree (AST) to represent conditional rules, allowing for dynamic creation, combination, and modification of these rules.

Features
Dynamic Rule Creation: Users can define eligibility rules using conditional expressions.
AST Representation: Rules are represented as ASTs, making it easy to parse, combine, and evaluate.
Efficient Rule Combination: The system minimizes redundant checks when combining rules.
Simple UI & API: Frontend and backend components to manage rule creation, combination, and evaluation.
Backend in Golang: Built using Go for efficient backend performance and scalability.
Data Structure
The rule engine uses an AST to represent rules. Each rule is parsed into a tree-like structure, with the following data structure:

Node Structure
go
Copy code
type Node struct {
    Type  string  // "operator" for AND/OR, "operand" for conditions
    Left  *Node   // Reference to the left child
    Right *Node   // Reference to the right child (for operators)
    Value string  // Optional value for operand nodes (e.g., "age > 30")
}
Type: Indicates whether the node is an "operator" (AND/OR) or an "operand" (condition).
Left & Right: References to child nodes.
Value: Contains the condition for operand nodes.
Data Storage
Database
The rules and application metadata are stored in a database. The choice of database can be adjusted based on scalability and performance requirements.

Sample Schema
Table Name	Columns
rules	id (PK), rule_string (text), created_at
rule_metadata	id (PK), rule_id (FK), attribute, condition
Sample Rules
rule1 = "((age > 30 AND department = 'Sales') OR (age < 25 AND department = 'Marketing')) AND (salary > 50000 OR experience > 5)"
rule2 = "((age > 30 AND department = 'Marketing')) AND (salary > 20000 OR experience > 5)"
API Design
The backend provides the following core APIs:

create_rule(rule_string)

Description: Takes a string representing a rule and returns a Node object representing the corresponding AST.
Input: rule_string (e.g., "age > 30 AND department = 'Sales'")
Output: AST representation of the rule.
combine_rules(rules)

Description: Combines multiple rules into a single AST, minimizing redundant checks.
Input: List of rule strings.
Output: Root node of the combined AST.
evaluate_rule(JSON data)

Description: Takes a JSON representation of the combined rule's AST and a dictionary containing user attributes. Evaluates the rule against the data and returns True if the user is eligible, False otherwise.
Input: JSON data (e.g., {"age": 35, "department": "Sales", "salary": 60000, "experience": 3})
Output: Boolean result (True/False).
Test Cases
Creating Individual Rules: Verify the AST representation of individual rules using create_rule.
Combining Rules: Use combine_rules and ensure the resulting AST reflects the combined logic.
Evaluating Rules: Implement sample JSON data and test evaluate_rule for different scenarios.
Extending Functionality: Combine additional rules and test the overall functionality.
Backend (Golang)
Setup Instructions
Prerequisites: Ensure Go is installed.
Install Dependencies: Run go mod tidy to install required packages.
Run the Application:
bash
Copy code
go run main.go
Sample Code
go
Copy code
// CreateRule parses a rule string and returns an AST node
func CreateRule(ruleString string) *Node {
    // Logic to parse rule string into AST
}

// CombineRules merges multiple rules into a single AST
func CombineRules(rules []string) *Node {
    // Logic to combine ASTs
}

// EvaluateRule checks user data against a rule and returns eligibility
func EvaluateRule(data map[string]interface{}, rule *Node) bool {
    // Logic to evaluate data using the AST
}
Contributing
Feel free to contribute to this project by creating issues or submitting pull requests. Please make sure to follow the code style and add tests for new features.

License
This project is licensed under the MIT License. See the LICENSE file for details.
