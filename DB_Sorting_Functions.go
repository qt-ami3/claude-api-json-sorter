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

func createBinDirectory() error { // createBinDirectory creates the ./bin directory if it doesn't exist.
	return os.MkdirAll("./bin", 0755)
}

func cleanupBinDirectory() error { // cleanupBinDirectory removes all files from ./bin directory.
	files, err := filepath.Glob("./bin/*")
	if err != nil {
		return err
	}
	
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("Warning: could not remove file %s: %v\n", file, err)
		}
	}
	return nil
}

func convertJSONToText(jsonData []byte, filename string) (string, error) {// convertJSONToText converts JSON data to formatted plaintext and saves to ./bin.
	if err := createBinDirectory(); err != nil { // Create ./bin directory if it doesn't exist.
		return "", fmt.Errorf("error creating bin directory: %v", err)
	}
	
	var prettyJSON interface{} // Pretty format the JSON for readability.
	if err := json.Unmarshal(jsonData, &prettyJSON); err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}
	
	formattedJSON, err := json.MarshalIndent(prettyJSON, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting JSON: %v", err)
	}
	
	txtPath := filepath.Join("./bin", filename+".txt") // Write to ./bin directory as .txt file
	if err := os.WriteFile(txtPath, formattedJSON, 0644); err != nil {
		return "", fmt.Errorf("error writing text file: %v", err)
	}
	
	return txtPath, nil
}
