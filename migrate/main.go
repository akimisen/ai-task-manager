package main

import (
	"ai-task-manager/internal/model"
	"flag"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 添加命令行参数
	dbPath := flag.String("db", "data/ai-task-manager.db", "Path to SQLite database file")
	flag.Parse()

	// 确保数据目录存在
	err := os.MkdirAll("data", 0755)
	if err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(*dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 执行迁移
	fmt.Println("Starting database migration...")
	err = db.AutoMigrate(&model.User{}, &model.TTSTask{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Migration completed successfully!")
}
