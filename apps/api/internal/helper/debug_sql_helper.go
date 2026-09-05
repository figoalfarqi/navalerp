package helper

import (
	"fmt"
	"strings"
	"time"
)

// cara penggunaan
// fmt.Println(helper.DebugSQL(query, args))
func DebugSQL(query string, args []interface{}) string {
	q := query
	for i, v := range args {
		placeholder := fmt.Sprintf("$%d", i+1)

		var val string
		switch t := v.(type) {
		case string:
			val = fmt.Sprintf("'%s'", t)
		case time.Time:
			val = fmt.Sprintf("'%s'", t.Format(time.RFC3339))
		default:
			val = fmt.Sprintf("%v", t)
		}

		q = strings.Replace(q, placeholder, val, 1)
	}
	return q
}
