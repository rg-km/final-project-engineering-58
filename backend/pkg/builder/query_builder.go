package builder

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

func INSERTStmt(tableName string, columns []string, rows int) string {
	inner := "( " + strings.TrimSuffix(strings.Repeat("?,", len(columns)), ",") + " ),"
	values := strings.TrimSuffix(strings.Repeat(inner, rows), ",")

	return fmt.Sprintf("INSERT INTO %s ( %s ) VALUES %s", tableName, strings.Join(columns, ","), values)
}

func UPDATEStmt(tableName string, columns map[string]interface{}, column string, id string) string {
	var values []string

	for k, v := range columns {
		values = append(values, fmt.Sprintf("%s = '%s'", k, reflect.ValueOf(v).String()))
	}

	return fmt.Sprintf("UPDATE %s SET %s WHERE %s = '%s'", tableName, strings.Join(values, ","), column, id)
}

func QueryExec(ctx context.Context, db *sql.DB, stmt string, args ...interface{}) (sql.Result, error) {
	for _, v := range args {
		if arg := reflect.ValueOf(v); arg.Kind() == reflect.Slice {
			args = FlattenArgs(args...)
			break
		}
	}

	res, err := db.ExecContext(ctx, stmt, args...)
	
	if err != nil {
		return nil, err
	}

	return res, nil
}

func FlattenArgs(args ...interface{}) []interface{} {
	out := make([]interface{}, 0, len(args))

	var sliceConv func(reflect.Value)
	sliceConv = func(arg reflect.Value) {
		if arg.Kind() == reflect.Slice {
			for i := 0; i < arg.Len(); i++ {
				sliceConv(reflect.ValueOf(arg.Index(i).Interface()))
			}
		} else if !arg.IsValid() {
			out = append(out, nil)
		} else {
			out = append(out, arg.Interface())
		}
	}

	for i := range args {
		arg := args[i]
		if rarg := reflect.ValueOf(arg); rarg.Kind() == reflect.Slice {
			sliceConv(rarg)
		} else {
			out = append(out, arg)
		}
	}

	return out
}

func RowsToJson(rows *sql.Rows) (res []map[string]interface{}) {
	res = [](map[string]interface{}){}
	var tmpInt int = 0
	columns, _ := rows.Columns()

	var values []interface{} = make([]interface{}, len(columns))
	var rawBytes [][]byte = make([][]byte, len(values))

	for v := range values {
		values[v] = &rawBytes[v]
	}

	for rows.Next() {
		rows.Scan(values...)

		if len(res) == tmpInt {
			res = append(res, make(map[string]interface{}))
		}

		for i, v := range columns {
			res[tmpInt][v] = string(rawBytes[i])
		}
		tmpInt++
	}
	return res
}