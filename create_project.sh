#!/bin/bash

# 创建主项目文件夹
mkdir -p ai-task-manager
cd ai-task-manager

# 创建文件夹结构
mkdir -p cmd/server
mkdir -p internal/api/handlers
mkdir -p internal/config
mkdir -p internal/models
mkdir -p internal/service
mkdir -p internal/queue
mkdir -p internal/clients
mkdir -p pkg/logger
mkdir -p pkg/utils
mkdir -p scripts
mkdir -p tests

# 创建主要文件
touch cmd/server/main.go
touch internal/api/routes.go
touch internal/api/handlers/task_handlers.go
touch internal/api/handlers/tts_handlers.go
touch internal/api/handlers/image_handlers.go
touch internal/api/handlers/video_handlers.go
touch internal/config/config.go
touch internal/models/task.go
touch internal/models/tts_task.go
touch internal/models/image_task.go
touch internal/models/video_task.go
touch internal/service/task_service.go
touch internal/service/tts_service.go
touch internal/service/image_service.go
touch internal/service/video_service.go
touch internal/queue/task_queue.go
touch internal/clients/tts_client.go
touch internal/clients/image_client.go
touch internal/clients/video_client.go
touch pkg/logger/logger.go
touch pkg/utils/utils.go
touch scripts/db_init.sql
touch go.mod
touch README.md

echo "Project structure created successfully!"