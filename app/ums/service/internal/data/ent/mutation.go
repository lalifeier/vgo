// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/data/ent/accountuser"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/data/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAccountUser = "AccountUser"
)

// AccountUserMutation represents an operation that mutates the AccountUser nodes in the graph.
type AccountUserMutation struct {
	config
	op               Op
	typ              string
	id               *int64
	email            *string
	phone            *string
	username         *string
	password         *string
	create_at        *int32
	addcreate_at     *int32
	create_ip_at     *string
	last_login_at    *int32
	addlast_login_at *int32
	last_login_ip_at *string
	login_times      *int32
	addlogin_times   *int32
	status           *int8
	addstatus        *int8
	clearedFields    map[string]struct{}
	done             bool
	oldValue         func(context.Context) (*AccountUser, error)
	predicates       []predicate.AccountUser
}

var _ ent.Mutation = (*AccountUserMutation)(nil)

// accountuserOption allows management of the mutation configuration using functional options.
type accountuserOption func(*AccountUserMutation)

// newAccountUserMutation creates new mutation for the AccountUser entity.
func newAccountUserMutation(c config, op Op, opts ...accountuserOption) *AccountUserMutation {
	m := &AccountUserMutation{
		config:        c,
		op:            op,
		typ:           TypeAccountUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAccountUserID sets the ID field of the mutation.
func withAccountUserID(id int64) accountuserOption {
	return func(m *AccountUserMutation) {
		var (
			err   error
			once  sync.Once
			value *AccountUser
		)
		m.oldValue = func(ctx context.Context) (*AccountUser, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().AccountUser.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAccountUser sets the old AccountUser of the mutation.
func withAccountUser(node *AccountUser) accountuserOption {
	return func(m *AccountUserMutation) {
		m.oldValue = func(context.Context) (*AccountUser, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AccountUserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AccountUserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of AccountUser entities.
func (m *AccountUserMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AccountUserMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetEmail sets the "email" field.
func (m *AccountUserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *AccountUserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the AccountUser entity.
// If the AccountUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountUserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *AccountUserMutation) ResetEmail() {
	m.email = nil
}

// SetPhone sets the "phone" field.
func (m *AccountUserMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the value of the "phone" field in the mutation.
func (m *AccountUserMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old "phone" field's value of the AccountUser entity.
// If the AccountUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountUserMutation) OldPhone(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPhone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ResetPhone resets all changes to the "phone" field.
func (m *AccountUserMutation) ResetPhone() {
	m.phone = nil
}

// SetUsername sets the "username" field.
func (m *AccountUserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *AccountUserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the AccountUser entity.
// If the AccountUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountUserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *AccountUserMutation) ResetUsername() {
	m.username = nil
}

// SetPassword sets the "password" field.
func (m *AccountUserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *AccountUserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the AccountUser entity.
// If the AccountUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountUserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *AccountUserMutation) ResetPassword() {
	m.password = nil
}

// SetCreateAt sets the "create_at" field.
func (m *AccountUserMutation) SetCreateAt(i int32) {
	m.create_at = &i
	m.addcreate_at = nil
}

// CreateAt returns the value of the "create_at" field in the mutation.
func (m *AccountUserMutation) CreateAt() (r int32, exists bool) {
	v := m.create_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateAt returns the old "create_at" field's value of the AccountUser entity.
// If the AccountUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountUserMutation) OldCreateAt(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateAt: %w", err)
	}
	return oldValue.CreateAt, nil
}

// AddCreateAt adds i to the "create_at" field.
func (m *AccountUserMutation) AddCreateAt(i int32) {
	if m.addcreate_at != nil {
		*m.addcreate_at += i
	} else {
		m.addcreate_at = &i
	}
}

// AddedCreateAt returns the value that was added to the "create_at" field in this mutation.
func (m *AccountUserMutation) AddedCreateAt() (r int32, exists bool) {
	v := m.addcreate_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreateAt resets all changes to the "create_at" field.
func (m *AccountUserMutation) ResetCreateAt() {
	m.create_at = nil
	m.addcreate_at = nil
}

// SetCreateIPAt sets the "create_ip_at" field.
func (m *AccountUserMutation) SetCreateIPAt(s string) {
	m.create_ip_at = &s
}

// CreateIPAt returns the value of the "create_ip_at" field in the mutation.
func (m *AccountUserMutation) CreateIPAt() (r string, exists bool) {
	v := m.create_ip_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateIPAt returns the old "create_ip_at" field's value of the AccountUser entity.
// If the AccountUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountUserMutation) OldCreateIPAt(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateIPAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateIPAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateIPAt: %w", err)
	}
	return oldValue.CreateIPAt, nil
}

// ResetCreateIPAt resets all changes to the "create_ip_at" field.
func (m *AccountUserMutation) ResetCreateIPAt() {
	m.create_ip_at = nil
}

// SetLastLoginAt sets the "last_login_at" field.
func (m *AccountUserMutation) SetLastLoginAt(i int32) {
	m.last_login_at = &i
	m.addlast_login_at = nil
}

// LastLoginAt returns the value of the "last_login_at" field in the mutation.
func (m *AccountUserMutation) LastLoginAt() (r int32, exists bool) {
	v := m.last_login_at
	if v == nil {
		return
	}
	return *v, true
}

// OldLastLoginAt returns the old "last_login_at" field's value of the AccountUser entity.
// If the AccountUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountUserMutation) OldLastLoginAt(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLastLoginAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLastLoginAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastLoginAt: %w", err)
	}
	return oldValue.LastLoginAt, nil
}

// AddLastLoginAt adds i to the "last_login_at" field.
func (m *AccountUserMutation) AddLastLoginAt(i int32) {
	if m.addlast_login_at != nil {
		*m.addlast_login_at += i
	} else {
		m.addlast_login_at = &i
	}
}

// AddedLastLoginAt returns the value that was added to the "last_login_at" field in this mutation.
func (m *AccountUserMutation) AddedLastLoginAt() (r int32, exists bool) {
	v := m.addlast_login_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetLastLoginAt resets all changes to the "last_login_at" field.
func (m *AccountUserMutation) ResetLastLoginAt() {
	m.last_login_at = nil
	m.addlast_login_at = nil
}

// SetLastLoginIPAt sets the "last_login_ip_at" field.
func (m *AccountUserMutation) SetLastLoginIPAt(s string) {
	m.last_login_ip_at = &s
}

// LastLoginIPAt returns the value of the "last_login_ip_at" field in the mutation.
func (m *AccountUserMutation) LastLoginIPAt() (r string, exists bool) {
	v := m.last_login_ip_at
	if v == nil {
		return
	}
	return *v, true
}

// OldLastLoginIPAt returns the old "last_login_ip_at" field's value of the AccountUser entity.
// If the AccountUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountUserMutation) OldLastLoginIPAt(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLastLoginIPAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLastLoginIPAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastLoginIPAt: %w", err)
	}
	return oldValue.LastLoginIPAt, nil
}

// ResetLastLoginIPAt resets all changes to the "last_login_ip_at" field.
func (m *AccountUserMutation) ResetLastLoginIPAt() {
	m.last_login_ip_at = nil
}

// SetLoginTimes sets the "login_times" field.
func (m *AccountUserMutation) SetLoginTimes(i int32) {
	m.login_times = &i
	m.addlogin_times = nil
}

// LoginTimes returns the value of the "login_times" field in the mutation.
func (m *AccountUserMutation) LoginTimes() (r int32, exists bool) {
	v := m.login_times
	if v == nil {
		return
	}
	return *v, true
}

// OldLoginTimes returns the old "login_times" field's value of the AccountUser entity.
// If the AccountUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountUserMutation) OldLoginTimes(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLoginTimes is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLoginTimes requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLoginTimes: %w", err)
	}
	return oldValue.LoginTimes, nil
}

// AddLoginTimes adds i to the "login_times" field.
func (m *AccountUserMutation) AddLoginTimes(i int32) {
	if m.addlogin_times != nil {
		*m.addlogin_times += i
	} else {
		m.addlogin_times = &i
	}
}

// AddedLoginTimes returns the value that was added to the "login_times" field in this mutation.
func (m *AccountUserMutation) AddedLoginTimes() (r int32, exists bool) {
	v := m.addlogin_times
	if v == nil {
		return
	}
	return *v, true
}

// ResetLoginTimes resets all changes to the "login_times" field.
func (m *AccountUserMutation) ResetLoginTimes() {
	m.login_times = nil
	m.addlogin_times = nil
}

// SetStatus sets the "status" field.
func (m *AccountUserMutation) SetStatus(i int8) {
	m.status = &i
	m.addstatus = nil
}

// Status returns the value of the "status" field in the mutation.
func (m *AccountUserMutation) Status() (r int8, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the AccountUser entity.
// If the AccountUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountUserMutation) OldStatus(ctx context.Context) (v int8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// AddStatus adds i to the "status" field.
func (m *AccountUserMutation) AddStatus(i int8) {
	if m.addstatus != nil {
		*m.addstatus += i
	} else {
		m.addstatus = &i
	}
}

// AddedStatus returns the value that was added to the "status" field in this mutation.
func (m *AccountUserMutation) AddedStatus() (r int8, exists bool) {
	v := m.addstatus
	if v == nil {
		return
	}
	return *v, true
}

// ResetStatus resets all changes to the "status" field.
func (m *AccountUserMutation) ResetStatus() {
	m.status = nil
	m.addstatus = nil
}

// Where appends a list predicates to the AccountUserMutation builder.
func (m *AccountUserMutation) Where(ps ...predicate.AccountUser) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *AccountUserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (AccountUser).
func (m *AccountUserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AccountUserMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.email != nil {
		fields = append(fields, accountuser.FieldEmail)
	}
	if m.phone != nil {
		fields = append(fields, accountuser.FieldPhone)
	}
	if m.username != nil {
		fields = append(fields, accountuser.FieldUsername)
	}
	if m.password != nil {
		fields = append(fields, accountuser.FieldPassword)
	}
	if m.create_at != nil {
		fields = append(fields, accountuser.FieldCreateAt)
	}
	if m.create_ip_at != nil {
		fields = append(fields, accountuser.FieldCreateIPAt)
	}
	if m.last_login_at != nil {
		fields = append(fields, accountuser.FieldLastLoginAt)
	}
	if m.last_login_ip_at != nil {
		fields = append(fields, accountuser.FieldLastLoginIPAt)
	}
	if m.login_times != nil {
		fields = append(fields, accountuser.FieldLoginTimes)
	}
	if m.status != nil {
		fields = append(fields, accountuser.FieldStatus)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AccountUserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case accountuser.FieldEmail:
		return m.Email()
	case accountuser.FieldPhone:
		return m.Phone()
	case accountuser.FieldUsername:
		return m.Username()
	case accountuser.FieldPassword:
		return m.Password()
	case accountuser.FieldCreateAt:
		return m.CreateAt()
	case accountuser.FieldCreateIPAt:
		return m.CreateIPAt()
	case accountuser.FieldLastLoginAt:
		return m.LastLoginAt()
	case accountuser.FieldLastLoginIPAt:
		return m.LastLoginIPAt()
	case accountuser.FieldLoginTimes:
		return m.LoginTimes()
	case accountuser.FieldStatus:
		return m.Status()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AccountUserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case accountuser.FieldEmail:
		return m.OldEmail(ctx)
	case accountuser.FieldPhone:
		return m.OldPhone(ctx)
	case accountuser.FieldUsername:
		return m.OldUsername(ctx)
	case accountuser.FieldPassword:
		return m.OldPassword(ctx)
	case accountuser.FieldCreateAt:
		return m.OldCreateAt(ctx)
	case accountuser.FieldCreateIPAt:
		return m.OldCreateIPAt(ctx)
	case accountuser.FieldLastLoginAt:
		return m.OldLastLoginAt(ctx)
	case accountuser.FieldLastLoginIPAt:
		return m.OldLastLoginIPAt(ctx)
	case accountuser.FieldLoginTimes:
		return m.OldLoginTimes(ctx)
	case accountuser.FieldStatus:
		return m.OldStatus(ctx)
	}
	return nil, fmt.Errorf("unknown AccountUser field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountUserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case accountuser.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case accountuser.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	case accountuser.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case accountuser.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case accountuser.FieldCreateAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateAt(v)
		return nil
	case accountuser.FieldCreateIPAt:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateIPAt(v)
		return nil
	case accountuser.FieldLastLoginAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastLoginAt(v)
		return nil
	case accountuser.FieldLastLoginIPAt:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastLoginIPAt(v)
		return nil
	case accountuser.FieldLoginTimes:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLoginTimes(v)
		return nil
	case accountuser.FieldStatus:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	}
	return fmt.Errorf("unknown AccountUser field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AccountUserMutation) AddedFields() []string {
	var fields []string
	if m.addcreate_at != nil {
		fields = append(fields, accountuser.FieldCreateAt)
	}
	if m.addlast_login_at != nil {
		fields = append(fields, accountuser.FieldLastLoginAt)
	}
	if m.addlogin_times != nil {
		fields = append(fields, accountuser.FieldLoginTimes)
	}
	if m.addstatus != nil {
		fields = append(fields, accountuser.FieldStatus)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AccountUserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case accountuser.FieldCreateAt:
		return m.AddedCreateAt()
	case accountuser.FieldLastLoginAt:
		return m.AddedLastLoginAt()
	case accountuser.FieldLoginTimes:
		return m.AddedLoginTimes()
	case accountuser.FieldStatus:
		return m.AddedStatus()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountUserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case accountuser.FieldCreateAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreateAt(v)
		return nil
	case accountuser.FieldLastLoginAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLastLoginAt(v)
		return nil
	case accountuser.FieldLoginTimes:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLoginTimes(v)
		return nil
	case accountuser.FieldStatus:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStatus(v)
		return nil
	}
	return fmt.Errorf("unknown AccountUser numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AccountUserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AccountUserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AccountUserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown AccountUser nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AccountUserMutation) ResetField(name string) error {
	switch name {
	case accountuser.FieldEmail:
		m.ResetEmail()
		return nil
	case accountuser.FieldPhone:
		m.ResetPhone()
		return nil
	case accountuser.FieldUsername:
		m.ResetUsername()
		return nil
	case accountuser.FieldPassword:
		m.ResetPassword()
		return nil
	case accountuser.FieldCreateAt:
		m.ResetCreateAt()
		return nil
	case accountuser.FieldCreateIPAt:
		m.ResetCreateIPAt()
		return nil
	case accountuser.FieldLastLoginAt:
		m.ResetLastLoginAt()
		return nil
	case accountuser.FieldLastLoginIPAt:
		m.ResetLastLoginIPAt()
		return nil
	case accountuser.FieldLoginTimes:
		m.ResetLoginTimes()
		return nil
	case accountuser.FieldStatus:
		m.ResetStatus()
		return nil
	}
	return fmt.Errorf("unknown AccountUser field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AccountUserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AccountUserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AccountUserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AccountUserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AccountUserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AccountUserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AccountUserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown AccountUser unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AccountUserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown AccountUser edge %s", name)
}
