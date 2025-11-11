package storage

import(
	"encoding/json"
	"log"
	"os"
	"sync" 
	"time"
)

type Task struct {
	ID          int 	  `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TaskStore struct {
	Mu     sync.RWMutex
	Tasks  []Task
	NextID int
}

var Store = &TaskStore{
	Tasks:  make([]Task, 0),
	NextID: 1,
}

// SaveTasksToFile salva as tarefas no arquivo JSON
func SaveTasksToFile(tasks []Task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		log.Printf("Erro ao criar arquivo: %v", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		log.Printf("Erro ao codificar JSON: %v", err)
		return err
	}
	return nil
}
