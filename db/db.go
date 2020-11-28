package db


import (
	"context"
	"database/sql"
  "database/sqlx"
	"fmt"
	"net/url"
)


type DBModel interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Begin() (*sql.Tx, error)
}


//Setting Export
type Setting struct {
	User     string
	Pwd      string
	Host     string
	Port     int
	DBName   string
	Location string
	Charset  string

	MaxOpenConns int
	MaxIdleConns int
}



type MysqlDB struct {
	DB *sql.DB
	Setting
}

// String will return the MySQL connection string.
func (m *MysqlDB) ConnStr() string {
	if m.Charset == "" {
		m.Charset = "utf8"
	}

	if m.Location != "" {
		m.Location = url.QueryEscape(m.Location)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=%s&charset=%s&parseTime=true&rejectReadOnly=true",
		m.User,
		m.Pwd,
		m.Host,
		m.Port,
		m.DBName,
		m.Location,
		m.Charset,
	)
}

func (m *MysqlDB) GetDB() error {

	mysqldb := NewMySQLDB(m.ConnStr(), m.MaxOpenConns, m.MaxIdleConns)
	fmt.Println("[MysqlDB.GetDB] 开始建立连接")
	db, err := mysqldb.Connect()
	if err != nil {
		fmt.Printf("[InitMySQL] 建立 MySQL 连接错误, error: %v", err)
		return err
	}
	m.DB = db
	return nil
}

type MySQLDB struct {
	connStr                    string
	maxOpenConns, maxIdleConns int
	maxConnLifetimeSec         int

	db  *sql.DB
	dbx *sqlx.DB

	setUpPrometheusOnce sync.Once
	closeOnce           sync.Once
	close               chan struct{}
}

func NewMySQLDB(connStr string, maxOpenConns, maxIdleConns int) *MySQLDB {
	m := &MySQLDB{}
	m.connStr = fmt.Sprintf("%s&rejectReadOnly=true", connStr)
	m.maxIdleConns = maxIdleConns
	m.maxOpenConns = maxOpenConns
	m.maxConnLifetimeSec = defaultMaxConnLifetimeSec
	m.close = make(chan struct{})
	return m
}


func (m *MysqlDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.DB.Query(query, args...)
}

func (m *MysqlDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return m.DB.QueryRow(query, args...)
}

func (m *MysqlDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return m.DB.QueryContext(ctx, query, args...)
}

func (m *MysqlDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return m.DB.QueryRowContext(ctx, query, args...)
}

func (m *MysqlDB) Prepare(query string) (*sql.Stmt, error) {
	return m.DB.Prepare(query)
}

func (m *MysqlDB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return m.DB.PrepareContext(ctx, query)
}

func (m *MysqlDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.DB.Exec(query, args...)
}

func (m *MysqlDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.DB.ExecContext(ctx, query, args...)
}

func (m *MysqlDB) Begin() (*sql.Tx, error) {
	return m.DB.Begin()
}

func (m *MysqlDB) Close() error {
	return m.DB.Close()
}
