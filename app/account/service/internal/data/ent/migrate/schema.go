// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountUserColumns holds the columns for the "account_user" table.
	AccountUserColumns = []*schema.Column{
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
	// AccountUserTable holds the schema information for the "account_user" table.
	AccountUserTable = &schema.Table{
		Name:       "account_user",
		Columns:    AccountUserColumns,
		PrimaryKey: []*schema.Column{AccountUserColumns[0]},
	}
	// StaffColumns holds the columns for the "staff" table.
	StaffColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "int(11)unsigned"}},
		{Name: "uid", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int(11)unsigned"}},
		{Name: "name", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(30)"}},
		{Name: "phone", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(15)"}},
		{Name: "email", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(30)"}},
		{Name: "nickname", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(30)"}},
		{Name: "avatar", Type: field.TypeString, Default: "", SchemaType: map[string]string{"mysql": "varchar(255)"}},
		{Name: "gender", Type: field.TypeInt8, Default: 0, SchemaType: map[string]string{"mysql": "tinyint(1)unsigned"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "created_by", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int(11)unsigned"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int(11)unsigned"}},
	}
	// StaffTable holds the schema information for the "staff" table.
	StaffTable = &schema.Table{
		Name:       "staff",
		Columns:    StaffColumns,
		PrimaryKey: []*schema.Column{StaffColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountUserTable,
		StaffTable,
	}
)

func init() {
	AccountUserTable.Annotation = &entsql.Annotation{
		Table:     "account_user",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
	}
	StaffTable.Annotation = &entsql.Annotation{
		Table:     "staff",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
	}
}
