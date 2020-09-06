package webserver

import (
	"fmt"
	"github.com/HordeGroup/horde/pkg/cache/session"
	"github.com/HordeGroup/horde/pkg/config"
	database "github.com/HordeGroup/horde/pkg/datebase"
	"github.com/HordeGroup/horde/pkg/helper"
	"github.com/HordeGroup/horde/pkg/webservice"
	"github.com/go-redis/redis"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"strconv"
	"time"
)

type Option struct {
	Verbose bool
	Host    string
	Port    int
	Service *webservice.Service
	Logger  zerolog.Logger
}

type Server struct {
	verbose bool
	host    string
	port    int
	service *webservice.Service
	logger  zerolog.Logger
}

func New(option Option) *Server {
	return &Server{
		verbose: option.Verbose,
		host:    option.Host,
		port:    option.Port,
		service: option.Service,
		logger:  option.Logger,
	}
}

func Run() error {
	var (
		err  error
		conf config.Config
		db   *gorm.DB
	)
	if err = configor.Load(&conf, "./pkg/config/config.yaml"); err != nil {
		return err
	}
	if db, err = database.New(database.Option{
		Name:     conf.Database.Name,
		Host:     conf.Database.Host,
		Port:     conf.Database.Port,
		User:     conf.Database.User,
		Password: conf.Database.Password,
		Debug:    conf.Debug,
	}); err != nil {
		return err
	}

	if err = database.GetMigration().Migrate(db); err != nil {
		return err
	}

	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		MasterName: "",
		Addrs:      []string{":" + strconv.FormatInt(conf.Redis.Port, 10)},
		Password:   conf.Redis.Password,
	})

	sessCache := session.NewCache(redisClient, time.Duration(conf.Session.TTL)*time.Second)

	svc := webservice.New(webservice.Option{
		Verbose:   false,
		DB:        db,
		SessCache: sessCache,
	})

	sv := New(Option{
		Host:    conf.Server.Host,
		Port:    conf.Server.Port,
		Service: svc,
		Logger:  zlog.Logger,
	})

	router := sv.BuildRouter()
	helper.RunHandler(fmt.Sprintf("0.0.0.0:%d", sv.port), router)
	return nil
}
