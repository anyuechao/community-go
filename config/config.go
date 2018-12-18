package config

import (
	"io/ioutil"
	"os"
	"regexp"
	"encoding/json"
	"fmt"
	"community-go/logger"
	"community-go/utility"
	)

var jsonData map[string]interface{}

func initJSON() {
	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		logger.Error("ReadFile: ", err.Error())
		os.Exit(-1)
	}
	// 去除注释
	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		logger.Log("json parse fail", err.Error())
		os.Exit(-1)
	}
}

type dbConfig struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Charset      string
	Host         string
	Port         int
	URL          string
	MaxIdleConns int
	MaxOpenConns int
}

var DBConfig dbConfig

func initDB() {
	utility.SetObjectByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
	//url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		DBConfig.User,
		DBConfig.Password,
		DBConfig.Host,
		DBConfig.Port,
		DBConfig.Database,
		DBConfig.Charset)
	DBConfig.URL = url
	//"maylol_test:maylol_test@tcp(39.104.82.179:3306)/merlot_test?charset=utf8"
	//db, _ := sql.Open("mysql", "root:Anyuechao1@tcp(localhost:3306)/merlot_test?charset=utf8")
	//db, _ := sql.Open("mysql", "maylol_test:maylol_test@tcp(39.104.82.179:3306)/merlot_test?charset=utf8")
}

type serverConfig struct {
	Env string
	SessionID string
	Port int
	PageSize int
	MaxPageSize int
	MinPageSize int
	MinOrder int
	MaxOrder int
	MaxNameLength int
	MaxContentLength int
	MaxArticleCateCount int
	MaxCommentLength int
	PortStr string
}

var ServerConfig serverConfig

func initServer() {

	utility.SetObjectByJSON(&ServerConfig, jsonData["server"].(map[string]interface{}))
	portstr := fmt.Sprintf(":%d",ServerConfig.Port)
	ServerConfig.PortStr = portstr
}

type redisConfig struct {
	Host string
	Port int
	Password string
	DB int
	Address string
}
var RedisConfig redisConfig

func initRedis()  {
	utility.SetObjectByJSON(&RedisConfig, jsonData["redis"].(map[string]interface{}))
	addressstr := fmt.Sprintf("%s:%d",RedisConfig.Host,RedisConfig.Port)
	RedisConfig.Address = addressstr
}

func init() {
	initJSON()
	initServer()
	initDB()
	initRedis()
}