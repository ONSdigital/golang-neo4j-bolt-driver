package golangNeo4jBoltDriver

import "net"

// Conn represents a connection to Neo4J
type Conn interface {
	Prepare(query string) (Stmt, error)
	Close() error
	Begin() (Tx, error)
}

type boltConn struct {
	connStr       string
	conn          net.Conn
	serverVersion []byte
	initialized   bool
}

// Prepare prepares a new statement for a query
func (c *boltConn) Prepare(query string) (Stmt, error) {
	return newStmt(query, c)
}

// Begin begins a new transaction with the Neo4J Database
func (c *boltConn) Begin() (Tx, error) {
	// TODO: Implement
	return nil, nil
}

// Close closes the connection
// Driver may allow for pooling in the future, keeping connections alive
func (c *boltConn) Close() error {
	// TODO: Connection Pooling?
	err := c.conn.Close()
	if err != nil {
		Logger.Print("An error occurred closing the connection", err)
		return err
	}
	return nil
}