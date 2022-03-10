package ent

import (
	"database/sql"
	"log"

	entsql "entgo.io/ent/dialect/sql"
)

// DB provides sql.DB
func (c *Client) DB() *sql.DB {
	return c.driver.(*entsql.Driver).DB()
}

// EnableSQLSafeUpdates turns sql_safe_updates on
func (c *Client) EnableSQLSafeUpdates() {
	_, err := c.DB().Exec(`set SESSION sql_safe_updates=1;`)
	if err != nil {
		log.Fatalf("EnableSQLSafeUpdates failed: %v", err)
	}
}

// DisableSQLSafeUpdates turns sql_safe_updates off
func (c *Client) DisableSQLSafeUpdates() {
	_, err := c.DB().Exec(`set SESSION sql_safe_updates=0;`)
	if err != nil {
		log.Fatalf("DisableSQLSafeUpdates failed: %v", err)
	}
}

// EnableForeignKeyChecks turns FOREIGN_KEY_CHECKS on.
func (c *Client) EnableForeignKeyChecks() {
	_, err := c.DB().Exec(`SET SESSION foreign_key_checks = 1;`)
	if err != nil {
		log.Fatalf("EnableForeignKeyChecks failed: %v", err)
	}
}

// DisableForeignKeyChecks turns FOREIGN_KEY_CHECKS off.
func (c *Client) DisableForeignKeyChecks() {
	_, err := c.DB().Exec(`SET SESSION foreign_key_checks = 0;`)
	if err != nil {
		log.Fatalf("DisableForeignKeyChecks failed: %v", err)
	}
}
