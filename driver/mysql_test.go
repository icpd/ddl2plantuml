package driver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMysql_Parse(t *testing.T) {
	ddl := `create table user (
		id int(11) not null auto_increment,
		name varchar(255) not null comment '名称',
		primary key (id)
	) comment '用户表' engine=innodb default charset=utf8 `

	d := &Mysql{}
	tables, err := d.Parse(ddl)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tables))
	assert.Equal(t, 2, len(tables[0].Columns))
	assert.Equal(t, "user", tables[0].Name)
	assert.Equal(t, `用户表`, tables[0].Comment)
	assert.Equal(t, "id", tables[0].Columns[0].Name)
	assert.Empty(t, tables[0].Columns[0].Comment)
	assert.True(t, tables[0].Columns[0].IsPrimaryKey)
	assert.Equal(t, "name", tables[0].Columns[1].Name)
	assert.Equal(t, `"名称"`, tables[0].Columns[1].Comment)
	assert.False(t, tables[0].Columns[1].IsPrimaryKey)
}
