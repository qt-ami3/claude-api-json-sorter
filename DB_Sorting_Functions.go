package main
import (	
	"strings"
)

func extractSQL(response string) string { // extractSQL removes markdown code blocks and extracts pure SQL.
	// Remove ```sql and ``` marker.
	sql := strings.ReplaceAll(response, "```sql", "")
	sql = strings.ReplaceAll(sql, "```", "")
	sql = strings.TrimSpace(sql)
	return sql
}

func min(a, b int) int { // min helper function
	if a < b {
		return a
	}
	return b
}


