// Code generated by ent, DO NOT EDIT.

package accountuser

import (
	"entgo.io/ent/dialect/sql"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldID, id))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldUsername, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldPhone, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldEmail, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldPassword, v))
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldCreateAt, v))
}

// CreateIPAt applies equality check predicate on the "create_ip_at" field. It's identical to CreateIPAtEQ.
func CreateIPAt(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldCreateIPAt, v))
}

// LastLoginAt applies equality check predicate on the "last_login_at" field. It's identical to LastLoginAtEQ.
func LastLoginAt(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldLastLoginAt, v))
}

// LastLoginIPAt applies equality check predicate on the "last_login_ip_at" field. It's identical to LastLoginIPAtEQ.
func LastLoginIPAt(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldLastLoginIPAt, v))
}

// LoginTimes applies equality check predicate on the "login_times" field. It's identical to LoginTimesEQ.
func LoginTimes(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldLoginTimes, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldStatus, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameIsNil applies the IsNil predicate on the "username" field.
func UsernameIsNil() predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIsNull(FieldUsername))
}

// UsernameNotNil applies the NotNil predicate on the "username" field.
func UsernameNotNil() predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotNull(FieldUsername))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContainsFold(FieldUsername, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContainsFold(FieldPhone, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContainsFold(FieldEmail, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIsNull(FieldPassword))
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotNull(FieldPassword))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContainsFold(FieldPassword, v))
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldCreateAt, v))
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldCreateAt, v))
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldCreateAt, vs...))
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldCreateAt, vs...))
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldCreateAt, v))
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldCreateAt, v))
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldCreateAt, v))
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldCreateAt, v))
}

// CreateIPAtEQ applies the EQ predicate on the "create_ip_at" field.
func CreateIPAtEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldCreateIPAt, v))
}

// CreateIPAtNEQ applies the NEQ predicate on the "create_ip_at" field.
func CreateIPAtNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldCreateIPAt, v))
}

// CreateIPAtIn applies the In predicate on the "create_ip_at" field.
func CreateIPAtIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldCreateIPAt, vs...))
}

// CreateIPAtNotIn applies the NotIn predicate on the "create_ip_at" field.
func CreateIPAtNotIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldCreateIPAt, vs...))
}

// CreateIPAtGT applies the GT predicate on the "create_ip_at" field.
func CreateIPAtGT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldCreateIPAt, v))
}

// CreateIPAtGTE applies the GTE predicate on the "create_ip_at" field.
func CreateIPAtGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldCreateIPAt, v))
}

// CreateIPAtLT applies the LT predicate on the "create_ip_at" field.
func CreateIPAtLT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldCreateIPAt, v))
}

// CreateIPAtLTE applies the LTE predicate on the "create_ip_at" field.
func CreateIPAtLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldCreateIPAt, v))
}

// CreateIPAtContains applies the Contains predicate on the "create_ip_at" field.
func CreateIPAtContains(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContains(FieldCreateIPAt, v))
}

// CreateIPAtHasPrefix applies the HasPrefix predicate on the "create_ip_at" field.
func CreateIPAtHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasPrefix(FieldCreateIPAt, v))
}

// CreateIPAtHasSuffix applies the HasSuffix predicate on the "create_ip_at" field.
func CreateIPAtHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasSuffix(FieldCreateIPAt, v))
}

// CreateIPAtEqualFold applies the EqualFold predicate on the "create_ip_at" field.
func CreateIPAtEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEqualFold(FieldCreateIPAt, v))
}

// CreateIPAtContainsFold applies the ContainsFold predicate on the "create_ip_at" field.
func CreateIPAtContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContainsFold(FieldCreateIPAt, v))
}

// LastLoginAtEQ applies the EQ predicate on the "last_login_at" field.
func LastLoginAtEQ(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldLastLoginAt, v))
}

// LastLoginAtNEQ applies the NEQ predicate on the "last_login_at" field.
func LastLoginAtNEQ(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldLastLoginAt, v))
}

// LastLoginAtIn applies the In predicate on the "last_login_at" field.
func LastLoginAtIn(vs ...int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldLastLoginAt, vs...))
}

// LastLoginAtNotIn applies the NotIn predicate on the "last_login_at" field.
func LastLoginAtNotIn(vs ...int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldLastLoginAt, vs...))
}

// LastLoginAtGT applies the GT predicate on the "last_login_at" field.
func LastLoginAtGT(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldLastLoginAt, v))
}

// LastLoginAtGTE applies the GTE predicate on the "last_login_at" field.
func LastLoginAtGTE(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldLastLoginAt, v))
}

// LastLoginAtLT applies the LT predicate on the "last_login_at" field.
func LastLoginAtLT(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldLastLoginAt, v))
}

// LastLoginAtLTE applies the LTE predicate on the "last_login_at" field.
func LastLoginAtLTE(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldLastLoginAt, v))
}

// LastLoginIPAtEQ applies the EQ predicate on the "last_login_ip_at" field.
func LastLoginIPAtEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldLastLoginIPAt, v))
}

// LastLoginIPAtNEQ applies the NEQ predicate on the "last_login_ip_at" field.
func LastLoginIPAtNEQ(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldLastLoginIPAt, v))
}

// LastLoginIPAtIn applies the In predicate on the "last_login_ip_at" field.
func LastLoginIPAtIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldLastLoginIPAt, vs...))
}

// LastLoginIPAtNotIn applies the NotIn predicate on the "last_login_ip_at" field.
func LastLoginIPAtNotIn(vs ...string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldLastLoginIPAt, vs...))
}

// LastLoginIPAtGT applies the GT predicate on the "last_login_ip_at" field.
func LastLoginIPAtGT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldLastLoginIPAt, v))
}

// LastLoginIPAtGTE applies the GTE predicate on the "last_login_ip_at" field.
func LastLoginIPAtGTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldLastLoginIPAt, v))
}

// LastLoginIPAtLT applies the LT predicate on the "last_login_ip_at" field.
func LastLoginIPAtLT(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldLastLoginIPAt, v))
}

// LastLoginIPAtLTE applies the LTE predicate on the "last_login_ip_at" field.
func LastLoginIPAtLTE(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldLastLoginIPAt, v))
}

// LastLoginIPAtContains applies the Contains predicate on the "last_login_ip_at" field.
func LastLoginIPAtContains(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContains(FieldLastLoginIPAt, v))
}

// LastLoginIPAtHasPrefix applies the HasPrefix predicate on the "last_login_ip_at" field.
func LastLoginIPAtHasPrefix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasPrefix(FieldLastLoginIPAt, v))
}

// LastLoginIPAtHasSuffix applies the HasSuffix predicate on the "last_login_ip_at" field.
func LastLoginIPAtHasSuffix(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldHasSuffix(FieldLastLoginIPAt, v))
}

// LastLoginIPAtEqualFold applies the EqualFold predicate on the "last_login_ip_at" field.
func LastLoginIPAtEqualFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEqualFold(FieldLastLoginIPAt, v))
}

// LastLoginIPAtContainsFold applies the ContainsFold predicate on the "last_login_ip_at" field.
func LastLoginIPAtContainsFold(v string) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldContainsFold(FieldLastLoginIPAt, v))
}

// LoginTimesEQ applies the EQ predicate on the "login_times" field.
func LoginTimesEQ(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldLoginTimes, v))
}

// LoginTimesNEQ applies the NEQ predicate on the "login_times" field.
func LoginTimesNEQ(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldLoginTimes, v))
}

// LoginTimesIn applies the In predicate on the "login_times" field.
func LoginTimesIn(vs ...int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldLoginTimes, vs...))
}

// LoginTimesNotIn applies the NotIn predicate on the "login_times" field.
func LoginTimesNotIn(vs ...int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldLoginTimes, vs...))
}

// LoginTimesGT applies the GT predicate on the "login_times" field.
func LoginTimesGT(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldLoginTimes, v))
}

// LoginTimesGTE applies the GTE predicate on the "login_times" field.
func LoginTimesGTE(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldLoginTimes, v))
}

// LoginTimesLT applies the LT predicate on the "login_times" field.
func LoginTimesLT(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldLoginTimes, v))
}

// LoginTimesLTE applies the LTE predicate on the "login_times" field.
func LoginTimesLTE(v int64) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldLoginTimes, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.AccountUser {
	return predicate.AccountUser(sql.FieldLTE(FieldStatus, v))
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
