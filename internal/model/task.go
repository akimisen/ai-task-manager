package model

import "time"

type TaskType string

const (
	TaskTypeTTS      TaskType = "tts"
	TaskTypeImageGen TaskType = "imagen"
)

type Task interface {
	GetID() string
	GetType() TaskType
	GetStatus() string
	SetStatus(status string)
	// 其他通用方法...
}

// 具体任务类型的实现
type TTSTask struct {
	ID        string
	Type      TaskType
	Status    string
	Text      string
	AudioData []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string // 添加用户关联
	Error     string // 用于存储错误信息
}

// type ImagenTask struct {
//     ID     string
//     Type   TaskType
//     Status string
//     Prompt string
//     // 其他图像生成特定字段...
// }

// 实现Task接口的方法
func (t *TTSTask) GetID() string           { return t.ID }
func (t *TTSTask) GetType() TaskType       { return t.Type }
func (t *TTSTask) GetStatus() string       { return t.Status }
func (t *TTSTask) SetStatus(status string) { t.Status = status }

// 为ImageGenTask实现类似的方法...
