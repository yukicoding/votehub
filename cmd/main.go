package main

import (
	"fmt"
	"log"

	config "yukicoding/voteHub/configs"
	"yukicoding/voteHub/internal/routes"
	"yukicoding/voteHub/pkg/database"
	"yukicoding/voteHub/pkg/logger"
	"yukicoding/voteHub/pkg/redis"

	"go.uber.org/zap"
)

func main() { // http://localhost:3000/swagger/index.html
	// 加载配置并初始化服务
	cfg, err := loading()
	if err != nil {
		log.Fatalf("Failed to load and initialize: %v", err)
	}

	// 设置路由
	r := routes.SetupRoutes()

	// 启动服务器
	logger.Info("Starting server", zap.String("port", cfg.Service.HttpPort))
	if err := r.Run(cfg.Service.HttpPort); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
func loading() (*config.Config, error) {
	cfg, err := config.LoadConfig("D:/gopath/src/yukicoding/voteHub/configs/config.yaml")

	// 初始化日志
	logger.Init(cfg.Log.LogPath, cfg.Log.LogLevel)

	logger.Warn(cfg.String())
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	// 初始化数据库
	if err := database.Init(cfg); err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// 初始化redis
	if err := redis.Init(cfg.Redis); err != nil {
		return nil, fmt.Errorf("failed to initialize redis: %w", err)
	}

	return cfg, nil
}
