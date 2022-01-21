// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/data/ent/accountuser"
)

// AccountUserCreate is the builder for creating a AccountUser entity.
type AccountUserCreate struct {
	config
	mutation *AccountUserMutation
	hooks    []Hook
}

// SetEmail sets the "email" field.
func (auc *AccountUserCreate) SetEmail(s string) *AccountUserCreate {
	auc.mutation.SetEmail(s)
	return auc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auc *AccountUserCreate) SetNillableEmail(s *string) *AccountUserCreate {
	if s != nil {
		auc.SetEmail(*s)
	}
	return auc
}

// SetPhone sets the "phone" field.
func (auc *AccountUserCreate) SetPhone(s string) *AccountUserCreate {
	auc.mutation.SetPhone(s)
	return auc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (auc *AccountUserCreate) SetNillablePhone(s *string) *AccountUserCreate {
	if s != nil {
		auc.SetPhone(*s)
	}
	return auc
}

// SetUsername sets the "username" field.
func (auc *AccountUserCreate) SetUsername(s string) *AccountUserCreate {
	auc.mutation.SetUsername(s)
	return auc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (auc *AccountUserCreate) SetNillableUsername(s *string) *AccountUserCreate {
	if s != nil {
		auc.SetUsername(*s)
	}
	return auc
}

// SetPassword sets the "password" field.
func (auc *AccountUserCreate) SetPassword(s string) *AccountUserCreate {
	auc.mutation.SetPassword(s)
	return auc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auc *AccountUserCreate) SetNillablePassword(s *string) *AccountUserCreate {
	if s != nil {
		auc.SetPassword(*s)
	}
	return auc
}

// SetCreateAt sets the "create_at" field.
func (auc *AccountUserCreate) SetCreateAt(i int32) *AccountUserCreate {
	auc.mutation.SetCreateAt(i)
	return auc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auc *AccountUserCreate) SetNillableCreateAt(i *int32) *AccountUserCreate {
	if i != nil {
		auc.SetCreateAt(*i)
	}
	return auc
}

// SetCreateIPAt sets the "create_ip_at" field.
func (auc *AccountUserCreate) SetCreateIPAt(s string) *AccountUserCreate {
	auc.mutation.SetCreateIPAt(s)
	return auc
}

// SetNillableCreateIPAt sets the "create_ip_at" field if the given value is not nil.
func (auc *AccountUserCreate) SetNillableCreateIPAt(s *string) *AccountUserCreate {
	if s != nil {
		auc.SetCreateIPAt(*s)
	}
	return auc
}

// SetLastLoginAt sets the "last_login_at" field.
func (auc *AccountUserCreate) SetLastLoginAt(i int32) *AccountUserCreate {
	auc.mutation.SetLastLoginAt(i)
	return auc
}

// SetNillableLastLoginAt sets the "last_login_at" field if the given value is not nil.
func (auc *AccountUserCreate) SetNillableLastLoginAt(i *int32) *AccountUserCreate {
	if i != nil {
		auc.SetLastLoginAt(*i)
	}
	return auc
}

// SetLastLoginIPAt sets the "last_login_ip_at" field.
func (auc *AccountUserCreate) SetLastLoginIPAt(s string) *AccountUserCreate {
	auc.mutation.SetLastLoginIPAt(s)
	return auc
}

// SetNillableLastLoginIPAt sets the "last_login_ip_at" field if the given value is not nil.
func (auc *AccountUserCreate) SetNillableLastLoginIPAt(s *string) *AccountUserCreate {
	if s != nil {
		auc.SetLastLoginIPAt(*s)
	}
	return auc
}

// SetLoginTimes sets the "login_times" field.
func (auc *AccountUserCreate) SetLoginTimes(i int32) *AccountUserCreate {
	auc.mutation.SetLoginTimes(i)
	return auc
}

// SetNillableLoginTimes sets the "login_times" field if the given value is not nil.
func (auc *AccountUserCreate) SetNillableLoginTimes(i *int32) *AccountUserCreate {
	if i != nil {
		auc.SetLoginTimes(*i)
	}
	return auc
}

// SetStatus sets the "status" field.
func (auc *AccountUserCreate) SetStatus(i int8) *AccountUserCreate {
	auc.mutation.SetStatus(i)
	return auc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auc *AccountUserCreate) SetNillableStatus(i *int8) *AccountUserCreate {
	if i != nil {
		auc.SetStatus(*i)
	}
	return auc
}

// SetID sets the "id" field.
func (auc *AccountUserCreate) SetID(i int64) *AccountUserCreate {
	auc.mutation.SetID(i)
	return auc
}

// Mutation returns the AccountUserMutation object of the builder.
func (auc *AccountUserCreate) Mutation() *AccountUserMutation {
	return auc.mutation
}

// Save creates the AccountUser in the database.
func (auc *AccountUserCreate) Save(ctx context.Context) (*AccountUser, error) {
	var (
		err  error
		node *AccountUser
	)
	auc.defaults()
	if len(auc.hooks) == 0 {
		if err = auc.check(); err != nil {
			return nil, err
		}
		node, err = auc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auc.check(); err != nil {
				return nil, err
			}
			auc.mutation = mutation
			if node, err = auc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(auc.hooks) - 1; i >= 0; i-- {
			if auc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (auc *AccountUserCreate) SaveX(ctx context.Context) *AccountUser {
	v, err := auc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (auc *AccountUserCreate) Exec(ctx context.Context) error {
	_, err := auc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auc *AccountUserCreate) ExecX(ctx context.Context) {
	if err := auc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auc *AccountUserCreate) defaults() {
	if _, ok := auc.mutation.Email(); !ok {
		v := accountuser.DefaultEmail
		auc.mutation.SetEmail(v)
	}
	if _, ok := auc.mutation.Phone(); !ok {
		v := accountuser.DefaultPhone
		auc.mutation.SetPhone(v)
	}
	if _, ok := auc.mutation.Username(); !ok {
		v := accountuser.DefaultUsername
		auc.mutation.SetUsername(v)
	}
	if _, ok := auc.mutation.Password(); !ok {
		v := accountuser.DefaultPassword
		auc.mutation.SetPassword(v)
	}
	if _, ok := auc.mutation.CreateAt(); !ok {
		v := accountuser.DefaultCreateAt
		auc.mutation.SetCreateAt(v)
	}
	if _, ok := auc.mutation.CreateIPAt(); !ok {
		v := accountuser.DefaultCreateIPAt
		auc.mutation.SetCreateIPAt(v)
	}
	if _, ok := auc.mutation.LastLoginAt(); !ok {
		v := accountuser.DefaultLastLoginAt
		auc.mutation.SetLastLoginAt(v)
	}
	if _, ok := auc.mutation.LastLoginIPAt(); !ok {
		v := accountuser.DefaultLastLoginIPAt
		auc.mutation.SetLastLoginIPAt(v)
	}
	if _, ok := auc.mutation.LoginTimes(); !ok {
		v := accountuser.DefaultLoginTimes
		auc.mutation.SetLoginTimes(v)
	}
	if _, ok := auc.mutation.Status(); !ok {
		v := accountuser.DefaultStatus
		auc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auc *AccountUserCreate) check() error {
	if _, ok := auc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "email"`)}
	}
	if _, ok := auc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "phone"`)}
	}
	if _, ok := auc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "username"`)}
	}
	if _, ok := auc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "password"`)}
	}
	if _, ok := auc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := auc.mutation.CreateIPAt(); !ok {
		return &ValidationError{Name: "create_ip_at", err: errors.New(`ent: missing required field "create_ip_at"`)}
	}
	if _, ok := auc.mutation.LastLoginAt(); !ok {
		return &ValidationError{Name: "last_login_at", err: errors.New(`ent: missing required field "last_login_at"`)}
	}
	if _, ok := auc.mutation.LastLoginIPAt(); !ok {
		return &ValidationError{Name: "last_login_ip_at", err: errors.New(`ent: missing required field "last_login_ip_at"`)}
	}
	if _, ok := auc.mutation.LoginTimes(); !ok {
		return &ValidationError{Name: "login_times", err: errors.New(`ent: missing required field "login_times"`)}
	}
	if _, ok := auc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	return nil
}

func (auc *AccountUserCreate) sqlSave(ctx context.Context) (*AccountUser, error) {
	_node, _spec := auc.createSpec()
	if err := sqlgraph.CreateNode(ctx, auc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (auc *AccountUserCreate) createSpec() (*AccountUser, *sqlgraph.CreateSpec) {
	var (
		_node = &AccountUser{config: auc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: accountuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: accountuser.FieldID,
			},
		}
	)
	if id, ok := auc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := auc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountuser.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := auc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountuser.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := auc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountuser.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := auc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountuser.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := auc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: accountuser.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := auc.mutation.CreateIPAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountuser.FieldCreateIPAt,
		})
		_node.CreateIPAt = value
	}
	if value, ok := auc.mutation.LastLoginAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: accountuser.FieldLastLoginAt,
		})
		_node.LastLoginAt = value
	}
	if value, ok := auc.mutation.LastLoginIPAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountuser.FieldLastLoginIPAt,
		})
		_node.LastLoginIPAt = value
	}
	if value, ok := auc.mutation.LoginTimes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: accountuser.FieldLoginTimes,
		})
		_node.LoginTimes = value
	}
	if value, ok := auc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: accountuser.FieldStatus,
		})
		_node.Status = value
	}
	return _node, _spec
}

// AccountUserCreateBulk is the builder for creating many AccountUser entities in bulk.
type AccountUserCreateBulk struct {
	config
	builders []*AccountUserCreate
}

// Save creates the AccountUser entities in the database.
func (aucb *AccountUserCreateBulk) Save(ctx context.Context) ([]*AccountUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aucb.builders))
	nodes := make([]*AccountUser, len(aucb.builders))
	mutators := make([]Mutator, len(aucb.builders))
	for i := range aucb.builders {
		func(i int, root context.Context) {
			builder := aucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aucb *AccountUserCreateBulk) SaveX(ctx context.Context) []*AccountUser {
	v, err := aucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aucb *AccountUserCreateBulk) Exec(ctx context.Context) error {
	_, err := aucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aucb *AccountUserCreateBulk) ExecX(ctx context.Context) {
	if err := aucb.Exec(ctx); err != nil {
		panic(err)
	}
}
