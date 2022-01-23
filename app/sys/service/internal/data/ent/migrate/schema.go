// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysDictColumns holds the columns for the "sys_dict" table.
	SysDictColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "int unsigned"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "created_by", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "bigint unsigned"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "bigint unsigned"}},
		{Name: "type", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "label", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "value", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(100)"}},
		{Name: "status", Type: field.TypeInt8, Default: 0, SchemaType: map[string]string{"mysql": "tinyint(1) unsigned"}},
		{Name: "remark", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(500)"}},
		{Name: "sort", Type: field.TypeInt8, Default: 0, SchemaType: map[string]string{"mysql": "tinyint(1) unsigned"}},
		{Name: "is_default", Type: field.TypeInt8, Default: 0, SchemaType: map[string]string{"mysql": "tinyint(1) unsigned"}},
		{Name: "is_deleted", Type: field.TypeInt8, Default: 0, SchemaType: map[string]string{"mysql": "tinyint(1) unsigned"}},
	}
	// SysDictTable holds the schema information for the "sys_dict" table.
	SysDictTable = &schema.Table{
		Name:       "sys_dict",
		Columns:    SysDictColumns,
		PrimaryKey: []*schema.Column{SysDictColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysDictTable,
	}
)

func init() {
	SysDictTable.Annotation = &entsql.Annotation{
		Table:     "sys_dict",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
	}
}