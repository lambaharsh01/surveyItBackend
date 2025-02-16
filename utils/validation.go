package utils

import (
	"fmt"
	"strings"
)

func ValidateSqlString(query string) error {
	query = strings.ToUpper(query)

	disallowedKeywords := []string{
        "INSERT", "UPDATE", "DELETE", "MERGE", "TRUNCATE",
        "ALTER", "DROP", "CREATE", "RENAME",
        "EXEC", "EXECUTE", "CALL", "GRANT", "REVOKE",
        "SET", "USE ", "DECLARE",
    }

	for _, keyword := range disallowedKeywords {
		if strings.Contains(query, keyword) {
			var errorString string = fmt.Sprintf( "disallowed operation (%s) detected in query", strings.ToLower(keyword) )
			return fmt.Errorf(errorString)
		}
	}

	return nil
}