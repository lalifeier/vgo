// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/lalifeier/vvgo-mall/app/auth/service/internal/data/ent/migrate"

	"github.com/lalifeier/vvgo-mall/app/auth/service/internal/data/ent/accountuser"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AccountUser is the client for interacting with the AccountUser builders.
	AccountUser *AccountUserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AccountUser = NewAccountUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		AccountUser: NewAccountUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:      cfg,
		AccountUser: NewAccountUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AccountUser.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.AccountUser.Use(hooks...)
}

// AccountUserClient is a client for the AccountUser schema.
type AccountUserClient struct {
	config
}

// NewAccountUserClient returns a client for the AccountUser from the given config.
func NewAccountUserClient(c config) *AccountUserClient {
	return &AccountUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `accountuser.Hooks(f(g(h())))`.
func (c *AccountUserClient) Use(hooks ...Hook) {
	c.hooks.AccountUser = append(c.hooks.AccountUser, hooks...)
}

// Create returns a create builder for AccountUser.
func (c *AccountUserClient) Create() *AccountUserCreate {
	mutation := newAccountUserMutation(c.config, OpCreate)
	return &AccountUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AccountUser entities.
func (c *AccountUserClient) CreateBulk(builders ...*AccountUserCreate) *AccountUserCreateBulk {
	return &AccountUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AccountUser.
func (c *AccountUserClient) Update() *AccountUserUpdate {
	mutation := newAccountUserMutation(c.config, OpUpdate)
	return &AccountUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountUserClient) UpdateOne(au *AccountUser) *AccountUserUpdateOne {
	mutation := newAccountUserMutation(c.config, OpUpdateOne, withAccountUser(au))
	return &AccountUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountUserClient) UpdateOneID(id int64) *AccountUserUpdateOne {
	mutation := newAccountUserMutation(c.config, OpUpdateOne, withAccountUserID(id))
	return &AccountUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AccountUser.
func (c *AccountUserClient) Delete() *AccountUserDelete {
	mutation := newAccountUserMutation(c.config, OpDelete)
	return &AccountUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AccountUserClient) DeleteOne(au *AccountUser) *AccountUserDeleteOne {
	return c.DeleteOneID(au.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AccountUserClient) DeleteOneID(id int64) *AccountUserDeleteOne {
	builder := c.Delete().Where(accountuser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountUserDeleteOne{builder}
}

// Query returns a query builder for AccountUser.
func (c *AccountUserClient) Query() *AccountUserQuery {
	return &AccountUserQuery{
		config: c.config,
	}
}

// Get returns a AccountUser entity by its id.
func (c *AccountUserClient) Get(ctx context.Context, id int64) (*AccountUser, error) {
	return c.Query().Where(accountuser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountUserClient) GetX(ctx context.Context, id int64) *AccountUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AccountUserClient) Hooks() []Hook {
	return c.hooks.AccountUser
}
