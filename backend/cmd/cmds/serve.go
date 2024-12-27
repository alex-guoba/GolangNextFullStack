/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmds

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/alex-guoba/gin-clean-template/internal/dao"
	"github.com/alex-guoba/gin-clean-template/pkg/logger"
	"github.com/alex-guoba/gin-clean-template/pkg/setting"
	"github.com/alex-guoba/gin-clean-template/pkg/signals"
	"github.com/alex-guoba/gin-clean-template/server"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"gorm.io/gorm"

	// for file migration source directory.
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfgFile string

func svrInit() error {
	if err := setting.Load(cfgFile); err != nil {
		return errors.Wrap(err, "loading config file failed")
	}
	gin.SetMode(setting.Conf.Server.RunMode)

	logger.SetupLogger(&setting.Conf.Log)

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "gin-clean-template",
	Short: "A clean architecture template for Golang Gin services",

	Run: func(_ *cobra.Command, _ []string) {
		if err := svrInit(); err != nil {
			log.Error("init server failed.", err)
			return
		}

		// init db
		engine, err := dbInit(&setting.Conf.Database)
		if err != nil {
			log.Error("init db failed.", err)
			return
		}

		// start http server
		svr := server.NewServer(setting.Conf, engine)
		if err := svr.Start(); err != nil {
			log.Error("init server failed.", err)
			return
		}

		// graceful shutdown
		stopCh := signals.SetupSignalHandler()
		sd, _ := signals.NewShutdown(setting.Conf.App.ServerShutdownTimeout)
		sd.Graceful(stopCh, svr, engine)
	},
}

func Execute() error {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd.Execute()
}

func dbInit(dbc *setting.DatabaseSettingS) (*gorm.DB, error) {
	// var err error
	engine, err := dao.NewDBEngine(dbc)
	if err != nil {
		return nil, err
	}

	// Run db migration
	dsn := fmt.Sprintf("%s://%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbc.DBType,
		dbc.UserName,
		dbc.Password,
		dbc.Host,
		dbc.DBName,
		dbc.Charset,
		dbc.ParseTime,
	)
	migration, err := migrate.New(dbc.MigrationURL, dsn)
	if err != nil {
		log.Error("migration init error.", err)
		return nil, err
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error("failed to run migrate up.", err)
		return engine, err
	}

	log.Info("db migrated successfully.")
	return engine, nil
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config.yml", "config file (default is ./config.yaml)")
}
