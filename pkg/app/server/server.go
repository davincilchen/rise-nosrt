package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	//"syscall"
	"github.com/sirupsen/logrus"
	gormlogger "gorm.io/gorm/logger"

	userUcase "rise-nostr/pkg/app/user/usecase"
	"rise-nostr/pkg/config"
	"rise-nostr/pkg/db"
)

type Server struct {
	Config *config.Config
}

func initLogger() {
	//log輸出為json格式
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//輸出設定為標準輸出(預設為stderr)
	logrus.SetOutput(os.Stdout)
	//設定要輸出的log等級
	logrus.SetLevel(logrus.DebugLevel)
}

func New(cfg *config.Config) *Server {

	return &Server{
		Config: cfg,
	}

}

func (t *Server) checkPort() {
	port, _ := strconv.Atoi(t.Config.Server.Port)

	for i := 0; i < 100; i++ {
		url := fmt.Sprintf("http://127.0.0.1:%d/info", port)
		_, err := http.Get(url)
		if err != nil {

			break
		}
		port++
	}

	t.Config.Server.Port = fmt.Sprintf("%d", port)

}

func (t *Server) Serve() {

	t.init()
	t.checkPort()
	addr := ":" + t.Config.Server.Port

	m := userUcase.GetUserManager()
	m.AddDefaultListener(config.GetRelayUrl(), "", "")

	logrus.Printf("======= Server start to listen (%s) and serve =======\n", addr)
	r := Router()
	r.Run(addr)

	logrus.Printf("======= Server Exit =======\n")
	//CloseLogger()
}

func (t *Server) init() {
	initLogger()
	t.initDB()
	db.Migration() //remove for production
}

func (t *Server) initDB() {

	newLogger := gormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormlogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			LogLevel:                  gormlogger.Info, // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,           // Disable color
		},
	)

	l := db.Logger{
		Logger: newLogger,
	}

	dbConn, err := db.GormOpen(&t.Config.DB, &l)
	if err != nil {
		logrus.Fatal(err)
	}
	db.SetMainDB(dbConn)

	sqlDB, err := dbConn.DB()
	if err != nil {
		logrus.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		logrus.Fatal(err)
	}
}
