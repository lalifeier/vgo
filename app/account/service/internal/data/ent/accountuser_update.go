// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/accountuser"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/predicate"
)

// AccountUserUpdate is the builder for updating AccountUser entities.
type AccountUserUpdate struct {
	config
	hooks     []Hook
	mutation  *AccountUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AccountUserUpdate builder.
func (auu *AccountUserUpdate) Where(ps ...predicate.AccountUser) *AccountUserUpdate {
	auu.mutation.Where(ps...)
	return auu
}

// SetUsername sets the "username" field.
func (auu *AccountUserUpdate) SetUsername(s string) *AccountUserUpdate {
	auu.mutation.SetUsername(s)
	return auu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (auu *AccountUserUpdate) SetNillableUsername(s *string) *AccountUserUpdate {
	if s != nil {
		auu.SetUsername(*s)
	}
	return auu
}

// ClearUsername clears the value of the "username" field.
func (auu *AccountUserUpdate) ClearUsername() *AccountUserUpdate {
	auu.mutation.ClearUsername()
	return auu
}

// SetPhone sets the "phone" field.
func (auu *AccountUserUpdate) SetPhone(s string) *AccountUserUpdate {
	auu.mutation.SetPhone(s)
	return auu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (auu *AccountUserUpdate) SetNillablePhone(s *string) *AccountUserUpdate {
	if s != nil {
		auu.SetPhone(*s)
	}
	return auu
}

// ClearPhone clears the value of the "phone" field.
func (auu *AccountUserUpdate) ClearPhone() *AccountUserUpdate {
	auu.mutation.ClearPhone()
	return auu
}

// SetEmail sets the "email" field.
func (auu *AccountUserUpdate) SetEmail(s string) *AccountUserUpdate {
	auu.mutation.SetEmail(s)
	return auu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auu *AccountUserUpdate) SetNillableEmail(s *string) *AccountUserUpdate {
	if s != nil {
		auu.SetEmail(*s)
	}
	return auu
}

// ClearEmail clears the value of the "email" field.
func (auu *AccountUserUpdate) ClearEmail() *AccountUserUpdate {
	auu.mutation.ClearEmail()
	return auu
}

// SetPassword sets the "password" field.
func (auu *AccountUserUpdate) SetPassword(s string) *AccountUserUpdate {
	auu.mutation.SetPassword(s)
	return auu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auu *AccountUserUpdate) SetNillablePassword(s *string) *AccountUserUpdate {
	if s != nil {
		auu.SetPassword(*s)
	}
	return auu
}

// ClearPassword clears the value of the "password" field.
func (auu *AccountUserUpdate) ClearPassword() *AccountUserUpdate {
	auu.mutation.ClearPassword()
	return auu
}

// SetCreateAt sets the "create_at" field.
func (auu *AccountUserUpdate) SetCreateAt(i int64) *AccountUserUpdate {
	auu.mutation.ResetCreateAt()
	auu.mutation.SetCreateAt(i)
	return auu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auu *AccountUserUpdate) SetNillableCreateAt(i *int64) *AccountUserUpdate {
	if i != nil {
		auu.SetCreateAt(*i)
	}
	return auu
}

// AddCreateAt adds i to the "create_at" field.
func (auu *AccountUserUpdate) AddCreateAt(i int64) *AccountUserUpdate {
	auu.mutation.AddCreateAt(i)
	return auu
}

// SetCreateIPAt sets the "create_ip_at" field.
func (auu *AccountUserUpdate) SetCreateIPAt(s string) *AccountUserUpdate {
	auu.mutation.SetCreateIPAt(s)
	return auu
}

// SetNillableCreateIPAt sets the "create_ip_at" field if the given value is not nil.
func (auu *AccountUserUpdate) SetNillableCreateIPAt(s *string) *AccountUserUpdate {
	if s != nil {
		auu.SetCreateIPAt(*s)
	}
	return auu
}

// SetLastLoginAt sets the "last_login_at" field.
func (auu *AccountUserUpdate) SetLastLoginAt(i int64) *AccountUserUpdate {
	auu.mutation.ResetLastLoginAt()
	auu.mutation.SetLastLoginAt(i)
	return auu
}

// SetNillableLastLoginAt sets the "last_login_at" field if the given value is not nil.
func (auu *AccountUserUpdate) SetNillableLastLoginAt(i *int64) *AccountUserUpdate {
	if i != nil {
		auu.SetLastLoginAt(*i)
	}
	return auu
}

// AddLastLoginAt adds i to the "last_login_at" field.
func (auu *AccountUserUpdate) AddLastLoginAt(i int64) *AccountUserUpdate {
	auu.mutation.AddLastLoginAt(i)
	return auu
}

// SetLastLoginIPAt sets the "last_login_ip_at" field.
func (auu *AccountUserUpdate) SetLastLoginIPAt(s string) *AccountUserUpdate {
	auu.mutation.SetLastLoginIPAt(s)
	return auu
}

// SetNillableLastLoginIPAt sets the "last_login_ip_at" field if the given value is not nil.
func (auu *AccountUserUpdate) SetNillableLastLoginIPAt(s *string) *AccountUserUpdate {
	if s != nil {
		auu.SetLastLoginIPAt(*s)
	}
	return auu
}

// SetLoginTimes sets the "login_times" field.
func (auu *AccountUserUpdate) SetLoginTimes(i int64) *AccountUserUpdate {
	auu.mutation.ResetLoginTimes()
	auu.mutation.SetLoginTimes(i)
	return auu
}

// SetNillableLoginTimes sets the "login_times" field if the given value is not nil.
func (auu *AccountUserUpdate) SetNillableLoginTimes(i *int64) *AccountUserUpdate {
	if i != nil {
		auu.SetLoginTimes(*i)
	}
	return auu
}

// AddLoginTimes adds i to the "login_times" field.
func (auu *AccountUserUpdate) AddLoginTimes(i int64) *AccountUserUpdate {
	auu.mutation.AddLoginTimes(i)
	return auu
}

// SetStatus sets the "status" field.
func (auu *AccountUserUpdate) SetStatus(i int8) *AccountUserUpdate {
	auu.mutation.ResetStatus()
	auu.mutation.SetStatus(i)
	return auu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auu *AccountUserUpdate) SetNillableStatus(i *int8) *AccountUserUpdate {
	if i != nil {
		auu.SetStatus(*i)
	}
	return auu
}

// AddStatus adds i to the "status" field.
func (auu *AccountUserUpdate) AddStatus(i int8) *AccountUserUpdate {
	auu.mutation.AddStatus(i)
	return auu
}

// Mutation returns the AccountUserMutation object of the builder.
func (auu *AccountUserUpdate) Mutation() *AccountUserMutation {
	return auu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (auu *AccountUserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AccountUserMutation](ctx, auu.sqlSave, auu.mutation, auu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auu *AccountUserUpdate) SaveX(ctx context.Context) int {
	affected, err := auu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (auu *AccountUserUpdate) Exec(ctx context.Context) error {
	_, err := auu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auu *AccountUserUpdate) ExecX(ctx context.Context) {
	if err := auu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auu *AccountUserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AccountUserUpdate {
	auu.modifiers = append(auu.modifiers, modifiers...)
	return auu
}

func (auu *AccountUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(accountuser.Table, accountuser.Columns, sqlgraph.NewFieldSpec(accountuser.FieldID, field.TypeUint32))
	if ps := auu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auu.mutation.Username(); ok {
		_spec.SetField(accountuser.FieldUsername, field.TypeString, value)
	}
	if auu.mutation.UsernameCleared() {
		_spec.ClearField(accountuser.FieldUsername, field.TypeString)
	}
	if value, ok := auu.mutation.Phone(); ok {
		_spec.SetField(accountuser.FieldPhone, field.TypeString, value)
	}
	if auu.mutation.PhoneCleared() {
		_spec.ClearField(accountuser.FieldPhone, field.TypeString)
	}
	if value, ok := auu.mutation.Email(); ok {
		_spec.SetField(accountuser.FieldEmail, field.TypeString, value)
	}
	if auu.mutation.EmailCleared() {
		_spec.ClearField(accountuser.FieldEmail, field.TypeString)
	}
	if value, ok := auu.mutation.Password(); ok {
		_spec.SetField(accountuser.FieldPassword, field.TypeString, value)
	}
	if auu.mutation.PasswordCleared() {
		_spec.ClearField(accountuser.FieldPassword, field.TypeString)
	}
	if value, ok := auu.mutation.CreateAt(); ok {
		_spec.SetField(accountuser.FieldCreateAt, field.TypeInt64, value)
	}
	if value, ok := auu.mutation.AddedCreateAt(); ok {
		_spec.AddField(accountuser.FieldCreateAt, field.TypeInt64, value)
	}
	if value, ok := auu.mutation.CreateIPAt(); ok {
		_spec.SetField(accountuser.FieldCreateIPAt, field.TypeString, value)
	}
	if value, ok := auu.mutation.LastLoginAt(); ok {
		_spec.SetField(accountuser.FieldLastLoginAt, field.TypeInt64, value)
	}
	if value, ok := auu.mutation.AddedLastLoginAt(); ok {
		_spec.AddField(accountuser.FieldLastLoginAt, field.TypeInt64, value)
	}
	if value, ok := auu.mutation.LastLoginIPAt(); ok {
		_spec.SetField(accountuser.FieldLastLoginIPAt, field.TypeString, value)
	}
	if value, ok := auu.mutation.LoginTimes(); ok {
		_spec.SetField(accountuser.FieldLoginTimes, field.TypeInt64, value)
	}
	if value, ok := auu.mutation.AddedLoginTimes(); ok {
		_spec.AddField(accountuser.FieldLoginTimes, field.TypeInt64, value)
	}
	if value, ok := auu.mutation.Status(); ok {
		_spec.SetField(accountuser.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := auu.mutation.AddedStatus(); ok {
		_spec.AddField(accountuser.FieldStatus, field.TypeInt8, value)
	}
	_spec.AddModifiers(auu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, auu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accountuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	auu.mutation.done = true
	return n, nil
}

// AccountUserUpdateOne is the builder for updating a single AccountUser entity.
type AccountUserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AccountUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUsername sets the "username" field.
func (auuo *AccountUserUpdateOne) SetUsername(s string) *AccountUserUpdateOne {
	auuo.mutation.SetUsername(s)
	return auuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (auuo *AccountUserUpdateOne) SetNillableUsername(s *string) *AccountUserUpdateOne {
	if s != nil {
		auuo.SetUsername(*s)
	}
	return auuo
}

// ClearUsername clears the value of the "username" field.
func (auuo *AccountUserUpdateOne) ClearUsername() *AccountUserUpdateOne {
	auuo.mutation.ClearUsername()
	return auuo
}

// SetPhone sets the "phone" field.
func (auuo *AccountUserUpdateOne) SetPhone(s string) *AccountUserUpdateOne {
	auuo.mutation.SetPhone(s)
	return auuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (auuo *AccountUserUpdateOne) SetNillablePhone(s *string) *AccountUserUpdateOne {
	if s != nil {
		auuo.SetPhone(*s)
	}
	return auuo
}

// ClearPhone clears the value of the "phone" field.
func (auuo *AccountUserUpdateOne) ClearPhone() *AccountUserUpdateOne {
	auuo.mutation.ClearPhone()
	return auuo
}

// SetEmail sets the "email" field.
func (auuo *AccountUserUpdateOne) SetEmail(s string) *AccountUserUpdateOne {
	auuo.mutation.SetEmail(s)
	return auuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auuo *AccountUserUpdateOne) SetNillableEmail(s *string) *AccountUserUpdateOne {
	if s != nil {
		auuo.SetEmail(*s)
	}
	return auuo
}

// ClearEmail clears the value of the "email" field.
func (auuo *AccountUserUpdateOne) ClearEmail() *AccountUserUpdateOne {
	auuo.mutation.ClearEmail()
	return auuo
}

// SetPassword sets the "password" field.
func (auuo *AccountUserUpdateOne) SetPassword(s string) *AccountUserUpdateOne {
	auuo.mutation.SetPassword(s)
	return auuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auuo *AccountUserUpdateOne) SetNillablePassword(s *string) *AccountUserUpdateOne {
	if s != nil {
		auuo.SetPassword(*s)
	}
	return auuo
}

// ClearPassword clears the value of the "password" field.
func (auuo *AccountUserUpdateOne) ClearPassword() *AccountUserUpdateOne {
	auuo.mutation.ClearPassword()
	return auuo
}

// SetCreateAt sets the "create_at" field.
func (auuo *AccountUserUpdateOne) SetCreateAt(i int64) *AccountUserUpdateOne {
	auuo.mutation.ResetCreateAt()
	auuo.mutation.SetCreateAt(i)
	return auuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auuo *AccountUserUpdateOne) SetNillableCreateAt(i *int64) *AccountUserUpdateOne {
	if i != nil {
		auuo.SetCreateAt(*i)
	}
	return auuo
}

// AddCreateAt adds i to the "create_at" field.
func (auuo *AccountUserUpdateOne) AddCreateAt(i int64) *AccountUserUpdateOne {
	auuo.mutation.AddCreateAt(i)
	return auuo
}

// SetCreateIPAt sets the "create_ip_at" field.
func (auuo *AccountUserUpdateOne) SetCreateIPAt(s string) *AccountUserUpdateOne {
	auuo.mutation.SetCreateIPAt(s)
	return auuo
}

// SetNillableCreateIPAt sets the "create_ip_at" field if the given value is not nil.
func (auuo *AccountUserUpdateOne) SetNillableCreateIPAt(s *string) *AccountUserUpdateOne {
	if s != nil {
		auuo.SetCreateIPAt(*s)
	}
	return auuo
}

// SetLastLoginAt sets the "last_login_at" field.
func (auuo *AccountUserUpdateOne) SetLastLoginAt(i int64) *AccountUserUpdateOne {
	auuo.mutation.ResetLastLoginAt()
	auuo.mutation.SetLastLoginAt(i)
	return auuo
}

// SetNillableLastLoginAt sets the "last_login_at" field if the given value is not nil.
func (auuo *AccountUserUpdateOne) SetNillableLastLoginAt(i *int64) *AccountUserUpdateOne {
	if i != nil {
		auuo.SetLastLoginAt(*i)
	}
	return auuo
}

// AddLastLoginAt adds i to the "last_login_at" field.
func (auuo *AccountUserUpdateOne) AddLastLoginAt(i int64) *AccountUserUpdateOne {
	auuo.mutation.AddLastLoginAt(i)
	return auuo
}

// SetLastLoginIPAt sets the "last_login_ip_at" field.
func (auuo *AccountUserUpdateOne) SetLastLoginIPAt(s string) *AccountUserUpdateOne {
	auuo.mutation.SetLastLoginIPAt(s)
	return auuo
}

// SetNillableLastLoginIPAt sets the "last_login_ip_at" field if the given value is not nil.
func (auuo *AccountUserUpdateOne) SetNillableLastLoginIPAt(s *string) *AccountUserUpdateOne {
	if s != nil {
		auuo.SetLastLoginIPAt(*s)
	}
	return auuo
}

// SetLoginTimes sets the "login_times" field.
func (auuo *AccountUserUpdateOne) SetLoginTimes(i int64) *AccountUserUpdateOne {
	auuo.mutation.ResetLoginTimes()
	auuo.mutation.SetLoginTimes(i)
	return auuo
}

// SetNillableLoginTimes sets the "login_times" field if the given value is not nil.
func (auuo *AccountUserUpdateOne) SetNillableLoginTimes(i *int64) *AccountUserUpdateOne {
	if i != nil {
		auuo.SetLoginTimes(*i)
	}
	return auuo
}

// AddLoginTimes adds i to the "login_times" field.
func (auuo *AccountUserUpdateOne) AddLoginTimes(i int64) *AccountUserUpdateOne {
	auuo.mutation.AddLoginTimes(i)
	return auuo
}

// SetStatus sets the "status" field.
func (auuo *AccountUserUpdateOne) SetStatus(i int8) *AccountUserUpdateOne {
	auuo.mutation.ResetStatus()
	auuo.mutation.SetStatus(i)
	return auuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auuo *AccountUserUpdateOne) SetNillableStatus(i *int8) *AccountUserUpdateOne {
	if i != nil {
		auuo.SetStatus(*i)
	}
	return auuo
}

// AddStatus adds i to the "status" field.
func (auuo *AccountUserUpdateOne) AddStatus(i int8) *AccountUserUpdateOne {
	auuo.mutation.AddStatus(i)
	return auuo
}

// Mutation returns the AccountUserMutation object of the builder.
func (auuo *AccountUserUpdateOne) Mutation() *AccountUserMutation {
	return auuo.mutation
}

// Where appends a list predicates to the AccountUserUpdate builder.
func (auuo *AccountUserUpdateOne) Where(ps ...predicate.AccountUser) *AccountUserUpdateOne {
	auuo.mutation.Where(ps...)
	return auuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auuo *AccountUserUpdateOne) Select(field string, fields ...string) *AccountUserUpdateOne {
	auuo.fields = append([]string{field}, fields...)
	return auuo
}

// Save executes the query and returns the updated AccountUser entity.
func (auuo *AccountUserUpdateOne) Save(ctx context.Context) (*AccountUser, error) {
	return withHooks[*AccountUser, AccountUserMutation](ctx, auuo.sqlSave, auuo.mutation, auuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auuo *AccountUserUpdateOne) SaveX(ctx context.Context) *AccountUser {
	node, err := auuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auuo *AccountUserUpdateOne) Exec(ctx context.Context) error {
	_, err := auuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auuo *AccountUserUpdateOne) ExecX(ctx context.Context) {
	if err := auuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auuo *AccountUserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AccountUserUpdateOne {
	auuo.modifiers = append(auuo.modifiers, modifiers...)
	return auuo
}

func (auuo *AccountUserUpdateOne) sqlSave(ctx context.Context) (_node *AccountUser, err error) {
	_spec := sqlgraph.NewUpdateSpec(accountuser.Table, accountuser.Columns, sqlgraph.NewFieldSpec(accountuser.FieldID, field.TypeUint32))
	id, ok := auuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AccountUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountuser.FieldID)
		for _, f := range fields {
			if !accountuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != accountuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auuo.mutation.Username(); ok {
		_spec.SetField(accountuser.FieldUsername, field.TypeString, value)
	}
	if auuo.mutation.UsernameCleared() {
		_spec.ClearField(accountuser.FieldUsername, field.TypeString)
	}
	if value, ok := auuo.mutation.Phone(); ok {
		_spec.SetField(accountuser.FieldPhone, field.TypeString, value)
	}
	if auuo.mutation.PhoneCleared() {
		_spec.ClearField(accountuser.FieldPhone, field.TypeString)
	}
	if value, ok := auuo.mutation.Email(); ok {
		_spec.SetField(accountuser.FieldEmail, field.TypeString, value)
	}
	if auuo.mutation.EmailCleared() {
		_spec.ClearField(accountuser.FieldEmail, field.TypeString)
	}
	if value, ok := auuo.mutation.Password(); ok {
		_spec.SetField(accountuser.FieldPassword, field.TypeString, value)
	}
	if auuo.mutation.PasswordCleared() {
		_spec.ClearField(accountuser.FieldPassword, field.TypeString)
	}
	if value, ok := auuo.mutation.CreateAt(); ok {
		_spec.SetField(accountuser.FieldCreateAt, field.TypeInt64, value)
	}
	if value, ok := auuo.mutation.AddedCreateAt(); ok {
		_spec.AddField(accountuser.FieldCreateAt, field.TypeInt64, value)
	}
	if value, ok := auuo.mutation.CreateIPAt(); ok {
		_spec.SetField(accountuser.FieldCreateIPAt, field.TypeString, value)
	}
	if value, ok := auuo.mutation.LastLoginAt(); ok {
		_spec.SetField(accountuser.FieldLastLoginAt, field.TypeInt64, value)
	}
	if value, ok := auuo.mutation.AddedLastLoginAt(); ok {
		_spec.AddField(accountuser.FieldLastLoginAt, field.TypeInt64, value)
	}
	if value, ok := auuo.mutation.LastLoginIPAt(); ok {
		_spec.SetField(accountuser.FieldLastLoginIPAt, field.TypeString, value)
	}
	if value, ok := auuo.mutation.LoginTimes(); ok {
		_spec.SetField(accountuser.FieldLoginTimes, field.TypeInt64, value)
	}
	if value, ok := auuo.mutation.AddedLoginTimes(); ok {
		_spec.AddField(accountuser.FieldLoginTimes, field.TypeInt64, value)
	}
	if value, ok := auuo.mutation.Status(); ok {
		_spec.SetField(accountuser.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := auuo.mutation.AddedStatus(); ok {
		_spec.AddField(accountuser.FieldStatus, field.TypeInt8, value)
	}
	_spec.AddModifiers(auuo.modifiers...)
	_node = &AccountUser{config: auuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accountuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auuo.mutation.done = true
	return _node, nil
}
