package main

import (
	"context"
	"fmt"
	"go-backend-clean-arch/api/httpserver"
	"go-backend-clean-arch/configs"
	"go-backend-clean-arch/internal/bootstrap"
	"go.uber.org/zap"
	log "log"
	"os"
	"os/signal"
)

func main() {
	// Bootstrap
	app := bootstrap.App(configs.Development)
	log.Printf("%#v", app)

	// Migrations
	//migrations(app)

	// Log
	app.Logger.Set().Named("main").Info("config", zap.Any("config", app.Config))

	// Start server
	server := httpserver.New(app)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)
	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, app.Config.Application.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	log.Println("received interrupt signal, shutting down gracefully..")
	// Close all db connection, etc
	app.CloseMysqlConnection()
	app.CloseRedisClientConnection()
	//app.ClosePostgresqlConnection() // Or etc..

	<-ctxWithTimeout.Done()
}

//func migrations(app *bootstrap.Application) {
// Mysql
//mysqlDir := "./internal/repository/migrations/mysqlmigration"
//migratorMysql := migratormysql.New(app.MysqlDB, mysqlDir)
//migratorMysql.Down()
//migratorMysql.Up()

// Pq
// pqDir := "./internal/repository/migrations/pqmigration"
// migratorPq := migratorpq.New(app.PostgresDB, pqDir)
// migratorPq.Down()
// migratorPq.Up()

// etc
//}
