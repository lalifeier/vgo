// Code generated by entc, DO NOT EDIT.

package accountuser

import (
	"entgo.io/ent/dialect/sql"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateIPAt applies equality check predicate on the "create_ip_at" field. It's identical to CreateIPAtEQ.
func CreateIPAt(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateIPAt), v))
	})
}

// LastLoginAt applies equality check predicate on the "last_login_at" field. It's identical to LastLoginAtEQ.
func LastLoginAt(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastLoginAt), v))
	})
}

// LastLoginIPAt applies equality check predicate on the "last_login_ip_at" field. It's identical to LastLoginIPAtEQ.
func LastLoginIPAt(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastLoginIPAt), v))
	})
}

// LoginTimes applies equality check predicate on the "login_times" field. It's identical to LoginTimesEQ.
func LoginTimes(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLoginTimes), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...int32) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...int32) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// CreateIPAtEQ applies the EQ predicate on the "create_ip_at" field.
func CreateIPAtEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateIPAt), v))
	})
}

// CreateIPAtNEQ applies the NEQ predicate on the "create_ip_at" field.
func CreateIPAtNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateIPAt), v))
	})
}

// CreateIPAtIn applies the In predicate on the "create_ip_at" field.
func CreateIPAtIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateIPAt), v...))
	})
}

// CreateIPAtNotIn applies the NotIn predicate on the "create_ip_at" field.
func CreateIPAtNotIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateIPAt), v...))
	})
}

// CreateIPAtGT applies the GT predicate on the "create_ip_at" field.
func CreateIPAtGT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateIPAt), v))
	})
}

// CreateIPAtGTE applies the GTE predicate on the "create_ip_at" field.
func CreateIPAtGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateIPAt), v))
	})
}

// CreateIPAtLT applies the LT predicate on the "create_ip_at" field.
func CreateIPAtLT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateIPAt), v))
	})
}

// CreateIPAtLTE applies the LTE predicate on the "create_ip_at" field.
func CreateIPAtLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateIPAt), v))
	})
}

// CreateIPAtContains applies the Contains predicate on the "create_ip_at" field.
func CreateIPAtContains(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreateIPAt), v))
	})
}

// CreateIPAtHasPrefix applies the HasPrefix predicate on the "create_ip_at" field.
func CreateIPAtHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreateIPAt), v))
	})
}

// CreateIPAtHasSuffix applies the HasSuffix predicate on the "create_ip_at" field.
func CreateIPAtHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreateIPAt), v))
	})
}

// CreateIPAtEqualFold applies the EqualFold predicate on the "create_ip_at" field.
func CreateIPAtEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreateIPAt), v))
	})
}

// CreateIPAtContainsFold applies the ContainsFold predicate on the "create_ip_at" field.
func CreateIPAtContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreateIPAt), v))
	})
}

// LastLoginAtEQ applies the EQ predicate on the "last_login_at" field.
func LastLoginAtEQ(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastLoginAt), v))
	})
}

// LastLoginAtNEQ applies the NEQ predicate on the "last_login_at" field.
func LastLoginAtNEQ(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastLoginAt), v))
	})
}

// LastLoginAtIn applies the In predicate on the "last_login_at" field.
func LastLoginAtIn(vs ...int32) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastLoginAt), v...))
	})
}

// LastLoginAtNotIn applies the NotIn predicate on the "last_login_at" field.
func LastLoginAtNotIn(vs ...int32) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastLoginAt), v...))
	})
}

// LastLoginAtGT applies the GT predicate on the "last_login_at" field.
func LastLoginAtGT(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastLoginAt), v))
	})
}

// LastLoginAtGTE applies the GTE predicate on the "last_login_at" field.
func LastLoginAtGTE(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastLoginAt), v))
	})
}

// LastLoginAtLT applies the LT predicate on the "last_login_at" field.
func LastLoginAtLT(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastLoginAt), v))
	})
}

// LastLoginAtLTE applies the LTE predicate on the "last_login_at" field.
func LastLoginAtLTE(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastLoginAt), v))
	})
}

// LastLoginIPAtEQ applies the EQ predicate on the "last_login_ip_at" field.
func LastLoginIPAtEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastLoginIPAt), v))
	})
}

// LastLoginIPAtNEQ applies the NEQ predicate on the "last_login_ip_at" field.
func LastLoginIPAtNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastLoginIPAt), v))
	})
}

// LastLoginIPAtIn applies the In predicate on the "last_login_ip_at" field.
func LastLoginIPAtIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastLoginIPAt), v...))
	})
}

// LastLoginIPAtNotIn applies the NotIn predicate on the "last_login_ip_at" field.
func LastLoginIPAtNotIn(vs ...string) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastLoginIPAt), v...))
	})
}

// LastLoginIPAtGT applies the GT predicate on the "last_login_ip_at" field.
func LastLoginIPAtGT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastLoginIPAt), v))
	})
}

// LastLoginIPAtGTE applies the GTE predicate on the "last_login_ip_at" field.
func LastLoginIPAtGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastLoginIPAt), v))
	})
}

// LastLoginIPAtLT applies the LT predicate on the "last_login_ip_at" field.
func LastLoginIPAtLT(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastLoginIPAt), v))
	})
}

// LastLoginIPAtLTE applies the LTE predicate on the "last_login_ip_at" field.
func LastLoginIPAtLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastLoginIPAt), v))
	})
}

// LastLoginIPAtContains applies the Contains predicate on the "last_login_ip_at" field.
func LastLoginIPAtContains(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLastLoginIPAt), v))
	})
}

// LastLoginIPAtHasPrefix applies the HasPrefix predicate on the "last_login_ip_at" field.
func LastLoginIPAtHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLastLoginIPAt), v))
	})
}

// LastLoginIPAtHasSuffix applies the HasSuffix predicate on the "last_login_ip_at" field.
func LastLoginIPAtHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLastLoginIPAt), v))
	})
}

// LastLoginIPAtEqualFold applies the EqualFold predicate on the "last_login_ip_at" field.
func LastLoginIPAtEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLastLoginIPAt), v))
	})
}

// LastLoginIPAtContainsFold applies the ContainsFold predicate on the "last_login_ip_at" field.
func LastLoginIPAtContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLastLoginIPAt), v))
	})
}

// LoginTimesEQ applies the EQ predicate on the "login_times" field.
func LoginTimesEQ(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLoginTimes), v))
	})
}

// LoginTimesNEQ applies the NEQ predicate on the "login_times" field.
func LoginTimesNEQ(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLoginTimes), v))
	})
}

// LoginTimesIn applies the In predicate on the "login_times" field.
func LoginTimesIn(vs ...int32) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLoginTimes), v...))
	})
}

// LoginTimesNotIn applies the NotIn predicate on the "login_times" field.
func LoginTimesNotIn(vs ...int32) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLoginTimes), v...))
	})
}

// LoginTimesGT applies the GT predicate on the "login_times" field.
func LoginTimesGT(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLoginTimes), v))
	})
}

// LoginTimesGTE applies the GTE predicate on the "login_times" field.
func LoginTimesGTE(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLoginTimes), v))
	})
}

// LoginTimesLT applies the LT predicate on the "login_times" field.
func LoginTimesLT(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLoginTimes), v))
	})
}

// LoginTimesLTE applies the LTE predicate on the "login_times" field.
func LoginTimesLTE(v int32) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLoginTimes), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.AccountUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccountUser) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccountUser) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccountUser) predicate.AccountUser {
	return predicate.AccountUser(func(s *sql.Selector) {
		p(s.Not())
	})
}
