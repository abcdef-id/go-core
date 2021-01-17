package config

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"github.com/go-redis/redis"
	//driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	mocket "github.com/selvatico/go-mocket"
	"gopkg.in/mgo.v2"

	mongo "github.com/abcdef-id/go-core/dependency/mgo"
)

var (
	//DB DB
	DB *gorm.DB
	//Mgo Mgo
	Mgo mongo.Session
	//Redis Redis
	Redis *redis.Client
)

//Database Database
type Database struct {
	Host              string
	User              string
	Password          string
	DBName            string
	DBNumber          int
	Port              int
	API_URL           string
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
}

// LoadDBConfig load database configuration
func LoadDBConfig(name string) Database {
	db := viper.Sub("database." + name)
	conf := Database{
		Host:              db.GetString("host"),
		User:              db.GetString("user"),
		Password:          db.GetString("password"),
		DBName:            db.GetString("db_name"),
		DBNumber:          db.GetInt("db_number"),
		Port:              db.GetInt("port"),
		API_URL:           db.GetString("api_url"),
		ReconnectRetry:    db.GetInt("reconnect_retry"),
		ReconnectInterval: db.GetInt64("reconnect_interval"),
		DebugMode:         db.GetBool("debug"),
	}
	return conf
}

// MysqlConnect connect to mysql using config name. return *gorm.DB incstance
func MysqlConnect(configName string) *gorm.DB {
	mysql := LoadDBConfig(configName)
	connection, err := gorm.Open("mysql", mysql.User+":"+mysql.Password+"@tcp("+mysql.Host+":"+strconv.Itoa(mysql.Port)+")/"+mysql.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	if mysql.DebugMode {
		return connection.Debug()
	}

	return connection
}

//MysqlConnectTest connect to mysql using config name. return *gorm.DB incstance
func MysqlConnectTest(configName string) *gorm.DB {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	connection, err := gorm.Open(mocket.DriverName, "connection_string")

	if err != nil {
		panic(err)
	}

	return connection
}

//OpenDbPool OpenDbPool
func OpenDbPool() {
	if viper.Get("env") != "testing" {
		DB = MysqlConnect("mysql")
	} else {
		DB = MysqlConnectTest("mysql")
	}
	pool := viper.Sub("database.mysql.pool")
	DB.DB().SetMaxOpenConns(pool.GetInt("maxOpenConns"))
	DB.DB().SetMaxIdleConns(pool.GetInt("maxIdleConns"))
	DB.DB().SetConnMaxLifetime(pool.GetDuration("maxLifetime") * time.Second)
}

//RedisConnect RedisConnect
func RedisConnect() {
	if viper.Get("env") != "testing" {
		conf := LoadDBConfig("redis")
		client := redis.NewClient(&redis.Options{
			Addr:     conf.Host + ":" + strconv.Itoa(conf.Port),
			Password: conf.Password,
			DB:       conf.DBNumber,
		})

		Redis = client
	}
}

//MongoConnect MongoConnect
func MongoConnect() {
	if viper.Get("env") != "testing" {
		conf := LoadDBConfig("mongo")
		session, err := mgo.Dial(conf.Host)
		if err != nil {
			panic(err)
		}
		Mgo = mongo.MongoSession{Session: session}
	}
}
