package driver

type Driver interface {
	Parse(ddl string) ([]Table, error)
}

type Column struct {
	Name         string
	Type         string
	Comment      string
	IsPrimaryKey bool
}

type Table struct {
	Name    string
	Columns []Column
}
