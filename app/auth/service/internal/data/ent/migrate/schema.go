// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthSystemColumns holds the columns for the "auth_system" table.
	AuthSystemColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "int(11)unsigned"}},
		{Name: "username", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(30)"}},
		{Name: "phone", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(15)"}},
		{Name: "email", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(30)"}},
		{Name: "password", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(32)"}},
		{Name: "create_at", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int(11)"}},
		{Name: "create_ip_at", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(12)"}},
		{Name: "last_login_at", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int(11)"}},
		{Name: "last_login_ip_at", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(12)"}},
		{Name: "login_times", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int(11)"}},
		{Name: "status", Type: field.TypeInt8, Default: 0, SchemaType: map[string]string{"mysql": "tinyint(1)"}},
	}
	// AuthSystemTable holds the schema information for the "auth_system" table.
	AuthSystemTable = &schema.Table{
		Name:       "auth_system",
		Columns:    AuthSystemColumns,
		PrimaryKey: []*schema.Column{AuthSystemColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthSystemTable,
	}
)

func init() {
	AuthSystemTable.Annotation = &entsql.Annotation{
		Table:     "auth_system",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
	}
}
