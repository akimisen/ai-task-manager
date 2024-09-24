package models

type VideoTask struct {
    ID     uint   `json:"id"`
    Name   string `json:"name"`
    Status string `json:"status"`
    // 其他字段
}