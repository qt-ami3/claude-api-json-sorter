package main

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

// API_KEY can be set here or loaded from environment variables.
var API_KEY = ""

// contains the instructions for Claude.
var promptAPI = `
You have two JSON files:
  1. clients.json - Contains client data with fields: id, firstName, lastName, location, income, shopingFor
  2. products.json - Contains product data with fields: id (company name), product, product Price, ad budget

TASK: Create a SQL file that matches each client with products they can afford based on their income.

For each client, find all products where the "product Price" is less than or equal to the client's income.

Output format should be SQL INSERT statements in this format:
INSERT INTO client_products (client_id, client_name, affordable_products) 
VALUES (1, 'Betta July', 'Devpulse: Bamboo Utensil Holder ($93.18), Feednation: Cranberry Almond Granola ($28.27)');

Requirements:
  - Parse the income field (remove $ sign and convert to number for comparison)
  - Parse the product Price field (remove $ sign and convert to number for comparison)
  - For each client, list ALL products they can afford in a comma-separated format
  - Include company name and product name with price
  - Create the table schema at the top of the SQL file
  - Make sure the SQL is valid and can be executed directly

Please generate a complete SQL file with:
  1. DROP TABLE IF EXISTS statement
  2. CREATE TABLE statement with appropriate columns
  3. INSERT statements for all clients with their affordable products
`

var client = anthropic.NewClient( // Initialize Claude client.
	option.WithAPIKey(API_KEY),
)

const modelNameAPI = "claude-sonnet-4-20250514" // 	Model in use.
