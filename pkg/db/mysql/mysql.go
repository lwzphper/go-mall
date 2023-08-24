package mysql

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
)

type Logger logger.Interface
type LoggerConfig logger.Config

var (
	mysqlClients = make(map[string]*DB)
	StdLogger    stdLogger
)

type stdLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

func init() {
	StdLogger = log.New(os.Stdout, "[Gorm] ", log.LstdFlags|log.Lshortfile)
}

type DB struct {
	*gorm.DB
	ClientName string
	Username   string
	password   string
	Host       string
	DBName     string
}

type Config struct {
	MaxOpenConn        int
	MaxIdleConn        int
	ConnMaxLifeSecond  time.Duration
	MaxIdleTime        time.Duration
	PrepareStmt        bool
	LogName            string
	SlowLogMillisecond int64
	EnableSqlLog       bool
	Logger             Logger
}
type Option func(*Config)

const (
	DefaultMaxOpenConn        = 1000
	DefaultMaxIdleConn        = 100
	DefaultConnMaxLifeSecond  = 30 * time.Minute
	DefaultLogName            = "gorm"
	DefaultSlowLogMillisecond = 200 // 慢查日志
	DefaultClient             = "default-mysql-client"
	ReadClient                = "read-mysql"
	WriteClient               = "write-mysql"
	TxClient                  = "tx-mysql"
)

// Reset reset Default config
func (o *Config) Reset() {
	o.MaxOpenConn = 0
	o.MaxIdleConn = 0
	o.ConnMaxLifeSecond = 0
	o.LogName = DefaultLogName
	o.PrepareStmt = false
	o.SlowLogMillisecond = DefaultSlowLogMillisecond
}

func WithMaxOpenConn(maxOpenConn int) Option {
	return func(opt *Config) {
		opt.MaxOpenConn = maxOpenConn
	}
}

func WithMaxIdleConn(maxIdleConn int) Option {
	return func(opt *Config) {
		opt.MaxIdleConn = maxIdleConn
	}
}

func WithConnMaxLifeSecond(connMaxLifeTime time.Duration) Option {
	return func(opt *Config) {
		opt.ConnMaxLifeSecond = connMaxLifeTime
	}
}
func WithMaxIdleTime(MaxIdleTime time.Duration) Option {
	return func(opt *Config) {
		opt.MaxIdleTime = MaxIdleTime
	}
}

func WithLogName(logName string) Option {
	return func(opt *Config) {
		opt.LogName = logName
	}
}

func WithSlowLogMillisecond(slowLogMillisecond int64) Option {
	return func(opt *Config) {
		opt.SlowLogMillisecond = slowLogMillisecond
	}
}

func WithPrepareStmt(prepareStmt bool) Option {
	return func(opt *Config) {
		opt.PrepareStmt = prepareStmt
	}
}
func WithEnableSqlLog(enableSqlLog bool) Option {
	return func(opt *Config) {
		opt.EnableSqlLog = enableSqlLog
	}
}

func WithLogger(writer io.Writer, config LoggerConfig) Option {
	return func(opt *Config) {
		opt.Logger = logger.New(log.New(writer, "[Gorm] ", log.LstdFlags|log.Lshortfile), logger.Config(config))
	}
}

// NewDefaultOption 创建默认配置项
func NewDefaultOption() *Config {
	logConf := NewDefaultLoggerConf()
	return &Config{
		MaxOpenConn:       DefaultMaxOpenConn,
		MaxIdleConn:       DefaultMaxIdleConn,
		ConnMaxLifeSecond: DefaultConnMaxLifeSecond,
		PrepareStmt:       true,
		Logger:            logger.New(StdLogger, logger.Config(logConf)),
	}
}

// NewDefaultLoggerConf 创建日志默认配置
func NewDefaultLoggerConf() LoggerConfig {
	return LoggerConfig{
		SlowThreshold:             time.Duration(DefaultSlowLogMillisecond) * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}
}

func InitMysqlClient(clientName, username, password, host, port, dbName string, options ...Option) error {
	if len(clientName) == 0 {
		return errors.New("client name is empty")
	}
	if len(username) == 0 {
		return errors.New("username is empty")
	}
	opt := NewDefaultOption()
	for _, f := range options {
		if f != nil {
			f(opt)
		}
	}

	db, err := dbConnect(username, password, host, port, dbName, opt)
	if err != nil {
		return errors.Wrapf(err, "host : "+host)
	}
	mysqlClients[clientName] = &DB{
		DB:         db,
		ClientName: clientName,
		Username:   username,
		password:   password,
		Host:       host,
		DBName:     dbName,
	}
	return nil
}

func GetMysqlClient(clientName string) *DB {
	if client, ok := mysqlClients[clientName]; ok {
		return client
	}
	return nil
}

func CloseMysqlClient(clientName string) error {
	sqlDB, err := GetMysqlClient(clientName).DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
	/*err = sqlDB.Close()
	if err != nil {
		return err
	}
	delete(mysqlClients, clientName)
	return nil*/
}

func dbConnect(user, pass, host, port, dbName string, option *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		pass,
		host,
		port,
		dbName,
		true,
		"Local")
	if option.SlowLogMillisecond == 0 {
		option.SlowLogMillisecond = DefaultSlowLogMillisecond
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）
		//如果没有这方面的要求，可以设置SkipDefaultTransaction为true来禁用它。
		//SkipDefaultTransaction: true,
		//执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续执行的效率
		PrepareStmt: option.PrepareStmt,
		NamingStrategy: schema.NamingStrategy{
			//使用单数表名,默认为复数表名，即当model的结构体为User时，默认操作的表名为users
			//设置	SingularTable: true 后当model的结构体为User时，操作的表名为user
			SingularTable: true,
			//TablePrefix: "pre_", //表前缀
		},
		Logger: option.Logger, // 日志配置，默认值：logger.Default.LogMode(logger.Info)
	})

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] Database name: %s", dbName))
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	if option.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(option.MaxOpenConn)
	} else {
		sqlDB.SetMaxOpenConns(DefaultMaxOpenConn)
	}

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	if option.MaxIdleConn > 0 {
		sqlDB.SetMaxIdleConns(option.MaxIdleConn)
	}

	// 设置最大连接超时时间
	if option.ConnMaxLifeSecond > 0 {
		sqlDB.SetConnMaxLifetime(time.Second * option.ConnMaxLifeSecond)
	}

	if option.MaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(time.Second * option.MaxIdleTime)
	}

	//// 监听事件
	//err = db.Callback().Create().After("gorm:after_create").Register(DefaultLogName, afterLog)
	//if err != nil {
	//	StdLogger.Print("Register Create error", err)
	//}
	//err = db.Callback().Query().After("gorm:after_query").Register(DefaultLogName, afterLog)
	//if err != nil {
	//	StdLogger.Print("Register Query error", err)
	//}
	//err = db.Callback().Update().After("gorm:after_update").Register(DefaultLogName, afterLog)
	//if err != nil {
	//	StdLogger.Print("Register Update error", err)
	//}
	//err = db.Callback().Delete().After("gorm:after_delete").Register(DefaultLogName, afterLog)
	//if err != nil {
	//	StdLogger.Print("Register Delete error", err)
	//}
	return db, nil
}

func afterLog(db *gorm.DB) {
	err := db.Error
	if err != nil {
		ctx := db.Statement.Context
		db.Logger.Error(ctx, err.Error())
	} else {
		sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
		fmt.Println("[ SQL语句 ]", sql)
	}

}
