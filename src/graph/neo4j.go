package graph

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Neo struct {
	conn bolt.Conn
}

func New(addr string) (*Neo, error) {
	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo(addr)
	if err != nil {
		return nil, err
	}
	return &Neo{
		conn: conn,
	}, nil
}

func (n *Neo) AddNode(username string) error {
	_, err := n.conn.ExecNeo("MERGE (a:Person {name: {username}}) RETURN a", map[string]interface{}{"username": username})
	return err
}

func (n *Neo) AddConnection(a, b string) error {
	_, err := n.conn.ExecNeo("MATCH (a:Person {name: {a}}), (b:Person {name: {b}}) MERGE (a)-[:FOLLOW]->(b)",
		map[string]interface{}{"a": a, "b": b},
	)
	return err
}
