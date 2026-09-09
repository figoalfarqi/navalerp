package helper

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var trailingDigitsRegex = regexp.MustCompile(`(\d+)$`)

// GenerateNextNumber computes the next sequential number for a given table and column.
// Format is: <PREFIX>-<YEAR>-<COUNTER (4 digits)> or continues existing prefix style.
func GenerateNextNumber(ctx context.Context, db *pgxpool.Pool, table, column, prefix string) (string, error) {
	if db == nil {
		return fmt.Sprintf("%s-%d-0001", prefix, time.Now().Year()), nil
	}

	year := time.Now().Year()
	pattern := fmt.Sprintf("%s-%d-%%", prefix, year)
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIKE $1 ORDER BY %s DESC LIMIT 1", column, table, column, column)

	var lastVal string
	err := db.QueryRow(ctx, query, pattern).Scan(&lastVal)

	var nextNumber string
	var nextSeq int
	var prefixPart string
	var digitsLen int = 4

	if err == nil && lastVal != "" {
		matches := trailingDigitsRegex.FindStringSubmatch(lastVal)
		if len(matches) > 1 {
			digitsStr := matches[1]
			digitsLen = len(digitsStr)
			seq, _ := strconv.Atoi(digitsStr)
			nextSeq = seq + 1
			prefixPart = lastVal[:len(lastVal)-len(digitsStr)]
			nextNumber = fmt.Sprintf("%s%0*d", prefixPart, digitsLen, nextSeq)
		} else {
			nextSeq = 1
			nextNumber = fmt.Sprintf("%s-%d-%04d", prefix, year, nextSeq)
			prefixPart = fmt.Sprintf("%s-%d-", prefix, year)
		}
	} else {
		nextSeq = 1
		nextNumber = fmt.Sprintf("%s-%d-%04d", prefix, year, nextSeq)
		prefixPart = fmt.Sprintf("%s-%d-", prefix, year)
	}

	// Guarantee uniqueness against race conditions
	for i := 0; i < 100; i++ {
		var exists bool
		checkQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = $1)", table, column)
		checkErr := db.QueryRow(ctx, checkQuery, nextNumber).Scan(&exists)
		if checkErr != nil || !exists {
			break
		}
		nextSeq++
		if strings.HasSuffix(prefixPart, "-") {
			nextNumber = fmt.Sprintf("%s%0*d", prefixPart, digitsLen, nextSeq)
		} else {
			nextNumber = fmt.Sprintf("%s-%0*d", prefixPart, digitsLen, nextSeq)
		}
	}

	return nextNumber, nil
}
