// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/accountuser"
)

// AccountUser is the model entity for the AccountUser schema.
type AccountUser struct {
	config `json:"-"`
	// ID of the ent.
	// 自增id
	ID uint32 `json:"id,omitempty"`
	// 用户名
	Username *string `json:"username,omitempty"`
	// 手机号
	Phone *string `json:"phone,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty"`
	// 密码
	Password *string `json:"password,omitempty"`
	// 创建时间
	CreateAt int64 `json:"create_at,omitempty"`
	// 创建ip
	CreateIPAt string `json:"create_ip_at,omitempty"`
	// 最后一次登录时间
	LastLoginAt int64 `json:"last_login_at,omitempty"`
	// 最后一次登录ip
	LastLoginIPAt string `json:"last_login_ip_at,omitempty"`
	// 登录次数
	LoginTimes int64 `json:"login_times,omitempty"`
	// 状态 0:禁用 1:启用 -1:删除
	Status int8 `json:"status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccountUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accountuser.FieldID, accountuser.FieldCreateAt, accountuser.FieldLastLoginAt, accountuser.FieldLoginTimes, accountuser.FieldStatus:
			values[i] = new(sql.NullInt64)
		case accountuser.FieldUsername, accountuser.FieldPhone, accountuser.FieldEmail, accountuser.FieldPassword, accountuser.FieldCreateIPAt, accountuser.FieldLastLoginIPAt:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AccountUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccountUser fields.
func (au *AccountUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accountuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			au.ID = uint32(value.Int64)
		case accountuser.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				au.Username = new(string)
				*au.Username = value.String
			}
		case accountuser.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				au.Phone = new(string)
				*au.Phone = value.String
			}
		case accountuser.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				au.Email = new(string)
				*au.Email = value.String
			}
		case accountuser.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				au.Password = new(string)
				*au.Password = value.String
			}
		case accountuser.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				au.CreateAt = value.Int64
			}
		case accountuser.FieldCreateIPAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field create_ip_at", values[i])
			} else if value.Valid {
				au.CreateIPAt = value.String
			}
		case accountuser.FieldLastLoginAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_at", values[i])
			} else if value.Valid {
				au.LastLoginAt = value.Int64
			}
		case accountuser.FieldLastLoginIPAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_ip_at", values[i])
			} else if value.Valid {
				au.LastLoginIPAt = value.String
			}
		case accountuser.FieldLoginTimes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field login_times", values[i])
			} else if value.Valid {
				au.LoginTimes = value.Int64
			}
		case accountuser.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				au.Status = int8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AccountUser.
// Note that you need to call AccountUser.Unwrap() before calling this method if this AccountUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (au *AccountUser) Update() *AccountUserUpdateOne {
	return NewAccountUserClient(au.config).UpdateOne(au)
}

// Unwrap unwraps the AccountUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (au *AccountUser) Unwrap() *AccountUser {
	_tx, ok := au.config.driver.(*txDriver)
	if !ok {
		panic("ent: AccountUser is not a transactional entity")
	}
	au.config.driver = _tx.drv
	return au
}

// String implements the fmt.Stringer.
func (au *AccountUser) String() string {
	var builder strings.Builder
	builder.WriteString("AccountUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", au.ID))
	if v := au.Username; v != nil {
		builder.WriteString("username=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := au.Phone; v != nil {
		builder.WriteString("phone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := au.Email; v != nil {
		builder.WriteString("email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := au.Password; v != nil {
		builder.WriteString("password=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(fmt.Sprintf("%v", au.CreateAt))
	builder.WriteString(", ")
	builder.WriteString("create_ip_at=")
	builder.WriteString(au.CreateIPAt)
	builder.WriteString(", ")
	builder.WriteString("last_login_at=")
	builder.WriteString(fmt.Sprintf("%v", au.LastLoginAt))
	builder.WriteString(", ")
	builder.WriteString("last_login_ip_at=")
	builder.WriteString(au.LastLoginIPAt)
	builder.WriteString(", ")
	builder.WriteString("login_times=")
	builder.WriteString(fmt.Sprintf("%v", au.LoginTimes))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", au.Status))
	builder.WriteByte(')')
	return builder.String()
}

// AccountUsers is a parsable slice of AccountUser.
type AccountUsers []*AccountUser
