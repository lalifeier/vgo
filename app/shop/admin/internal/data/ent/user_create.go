// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/role"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetCreatedBy sets the "created_by" field.
func (uc *UserCreate) SetCreatedBy(i int64) *UserCreate {
	uc.mutation.SetCreatedBy(i)
	return uc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedBy(i *int64) *UserCreate {
	if i != nil {
		uc.SetCreatedBy(*i)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetUpdatedBy sets the "updated_by" field.
func (uc *UserCreate) SetUpdatedBy(i int64) *UserCreate {
	uc.mutation.SetUpdatedBy(i)
	return uc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedBy(i *int64) *UserCreate {
	if i != nil {
		uc.SetUpdatedBy(*i)
	}
	return uc
}

// SetUID sets the "uid" field.
func (uc *UserCreate) SetUID(i int64) *UserCreate {
	uc.mutation.SetUID(i)
	return uc
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (uc *UserCreate) SetNillableUID(i *int64) *UserCreate {
	if i != nil {
		uc.SetUID(*i)
	}
	return uc
}

// SetNickname sets the "nickname" field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uc *UserCreate) SetNillableNickname(s *string) *UserCreate {
	if s != nil {
		uc.SetNickname(*s)
	}
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// SetGender sets the "gender" field.
func (uc *UserCreate) SetGender(u user.Gender) *UserCreate {
	uc.mutation.SetGender(u)
	return uc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uc *UserCreate) SetNillableGender(u *user.Gender) *UserCreate {
	if u != nil {
		uc.SetGender(*u)
	}
	return uc
}

// SetRemark sets the "remark" field.
func (uc *UserCreate) SetRemark(s string) *UserCreate {
	uc.mutation.SetRemark(s)
	return uc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (uc *UserCreate) SetNillableRemark(s *string) *UserCreate {
	if s != nil {
		uc.SetRemark(*s)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int64) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uc *UserCreate) AddRoleIDs(ids ...int64) *UserCreate {
	uc.mutation.AddRoleIDs(ids...)
	return uc
}

// AddRoles adds the "roles" edges to the Role entity.
func (uc *UserCreate) AddRoles(r ...*Role) *UserCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRoleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if err := uc.defaults(); err != nil {
		return nil, err
	}
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		if user.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.CreatedBy(); !ok {
		v := user.DefaultCreatedBy
		uc.mutation.SetCreatedBy(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		if user.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedBy(); !ok {
		v := user.DefaultUpdatedBy
		uc.mutation.SetUpdatedBy(v)
	}
	if _, ok := uc.mutation.UID(); !ok {
		v := user.DefaultUID
		uc.mutation.SetUID(v)
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		v := user.DefaultNickname
		uc.mutation.SetNickname(v)
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		v := user.DefaultAvatar
		uc.mutation.SetAvatar(v)
	}
	if _, ok := uc.mutation.Gender(); !ok {
		v := user.DefaultGender
		uc.mutation.SetGender(v)
	}
	if _, ok := uc.mutation.Remark(); !ok {
		v := user.DefaultRemark
		uc.mutation.SetRemark(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "User.created_by"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "User.updated_by"`)}
	}
	if _, ok := uc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "User.uid"`)}
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`ent: missing required field "User.nickname"`)}
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "User.avatar"`)}
	}
	if _, ok := uc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "User.gender"`)}
	}
	if v, ok := uc.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "User.gender": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "User.remark"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
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

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := uc.mutation.UID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldUID,
		})
		_node.UID = value
	}
	if value, ok := uc.mutation.Nickname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
		_node.Nickname = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatar,
		})
		_node.Avatar = value
	}
	if value, ok := uc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldGender,
		})
		_node.Gender = value
	}
	if value, ok := uc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRemark,
		})
		_node.Remark = value
	}
	if nodes := uc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: []string{user.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}