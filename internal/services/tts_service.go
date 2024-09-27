package services

import (
    "ai-task-manager/internal/models"
    "ai-task-manager/internal/queue"
    "bytes"
    "encoding/json"
    "fmt"
    "io"
	"os"
    "path/filepath"
    "net/http"
    "github.com/spf13/viper"
)

type TTSService struct {
    queue queue.TaskQueue
    ttsURL string
    defaultRefAudio string
    defaultTextLang string
    defaultPromptLang string
}

func NewTTSService(q queue.TaskQueue, ttsURL string) *TTSService {
    return &TTSService{
        taskQueue: q,
		ttsURL: ttsURL,
    }
}

func (s *TTSService) CreateTask(text string) (*models.TTSTask, error) {
    task := &models.TTSTask{
        ID:     generateUniqueID(), // 实现一个生成唯一ID的函数
        Type:   models.TaskTypeTTS,
        Status: "pending",
        Text:   text,
    }

    err := s.queue.Push(task)
    if err != nil {
        return nil, err
    }

    // 异步调用TTS API
    go s.processTTSTask(task)

    return task, nil
}

func (s *TTSService) processTTSTask(task *models.TTSTask) {
    task.Status = "processing"
    s.queue.Push(task) // 更新任务状态

    // 准备请求数据
    requestBody, _ := json.Marshal(map[string]string{
        "text": task.Text,
		"text_lang": s.defaultTextLang,
        // "ref_audio_path": s.defaultRefAudio,
        "prompt_lang": s.defaultPromptLang,
        "streaming_mode": false,
    })

    // 调用第三方 TTS API
    resp, err := http.Post(s.ttsURL, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        task.Status = "failed"
        s.queue.Push(task)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        // 读取音频数据
        audioData, err := io.ReadAll(resp.Body)
        if err != nil {
            task.Status = "failed"
        } else {
            task.Status = "completed"
            task.AudioData = audioData // 假设 TTSTask 结构体有 AudioData 字段
        }
    } else {
        task.Status = "failed"
    }

    s.queue.Push(task) // 更新最终状态
}

func (s *TTSService) GetTask(id string) (*models.TTSTask, error) {
    tasks, err := s.queue.List()
    if err != nil {
        return nil, err
    }

    for _, t := range tasks {
        if ttsTask, ok := t.(*models.TTSTask); ok && ttsTask.ID == id {
            return ttsTask, nil
        }
    }

    return nil, nil // 任务未找到
}

func (s *TTSService) ListTasks() ([]*models.TTSTask, error) {
    tasks, err := s.taskQueue.List()
    if err != nil {
        return nil, err
    }
    ttsTasks := make([]*models.TTSTask, 0, len(tasks))
    for _, task := range tasks {
        if ttsTask, ok := task.(*models.TTSTask); ok {
            ttsTasks = append(ttsTasks, ttsTask)
        }
    }
    return ttsTasks, nil
}