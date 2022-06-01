package driver

import (
	"strings"

	"github.com/jinzhu/inflection"
)

type Driver interface {
	Parse(ddl string) (Tables, error)
}

type Column struct {
	Name         string
	Type         string
	Comment      string
	IsPrimaryKey bool
}

type Table struct {
	Name    string
	Comment string
	Columns []Column
}

type Relation struct {
	TableName      string
	RelationTables []RelationTable
}

type RelationTable struct {
	TableName string
	Column    string
}

type Tables []Table

func (r Tables) Relationship() []Relation {
	relationMap := make(map[string][]RelationTable)

	for _, table := range r {
		if _, ok := relationMap[table.Name]; !ok {
			relationMap[table.Name] = []RelationTable{}
		}

		for _, column := range table.Columns {
			if strings.Index(column.Name, "_id") == -1 {
				continue
			}

			singularTableName := strings.ReplaceAll(column.Name, "_id", "")
			pluralTableName := inflection.Plural(singularTableName)

			if _, ok := relationMap[singularTableName]; ok {
				relationMap[singularTableName] = append(relationMap[singularTableName], RelationTable{
					TableName: table.Name,
					Column:    column.Name,
				})
				continue
			}

			if _, ok := relationMap[pluralTableName]; ok {
				relationMap[pluralTableName] = append(relationMap[pluralTableName], RelationTable{
					TableName: table.Name,
					Column:    column.Name,
				})
				continue
			}
		}

	}

	var relations []Relation
	for table, relationTables := range relationMap {
		if len(relationTables) == 0 {
			continue
		}

		var rts []RelationTable
		for _, relationTable := range relationTables {
			rts = append(rts, RelationTable{
				TableName: relationTable.TableName,
				Column:    relationTable.Column,
			})
		}

		relations = append(relations, Relation{
			TableName:      table,
			RelationTables: rts,
		})
	}

	return relations
}
