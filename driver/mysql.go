package driver

import (
	"strings"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

type Mysql struct{}

func (m *Mysql) Parse(ddl string) ([]Table, error) {
	tables := strings.Split(ddl, ";")
	result := make([]Table, 0, len(tables)) // not accurate table count
	for _, ddl := range tables {
		if strings.TrimSpace(ddl) == "" {
			continue
		}

		stmt, err := sqlparser.Parse(ddl)
		if err != nil {
			return nil, err
		}

		createTable, ok := stmt.(*sqlparser.CreateTable)
		if !ok {
			continue
		}

		table := Table{
			Name:    createTable.NewName.Name.String(),
			Columns: make([]Column, 0, len(createTable.Columns)),
		}

		for _, option := range createTable.Options {
			if option.Type == sqlparser.TableOptionComment {
				table.Comment = option.StrValue
			}
		}

		var primaryKeyName string
		for _, constraint := range createTable.Constraints {
			if constraint.Type == sqlparser.ConstraintPrimaryKey {
				primaryKeyName = constraint.Keys[0].String()
			}
		}

		for _, column := range createTable.Columns {
			c := Column{
				Name: column.Name,
				Type: strings.Split(column.Type, "(")[0],
			}

			if column.Name == primaryKeyName {
				c.IsPrimaryKey = true
			}

			for _, option := range column.Options {
				if option.Type == sqlparser.ColumnOptionPrimaryKey {
					c.IsPrimaryKey = true
					continue
				}

				if option.Type != sqlparser.ColumnOptionComment {
					continue
				}

				c.Comment = strings.TrimSpace(option.Value)
			}

			table.Columns = append(table.Columns, c)
		}

		result = append(result, table)
	}

	return result, nil
}
