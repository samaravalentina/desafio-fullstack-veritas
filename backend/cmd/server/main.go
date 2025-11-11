package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"backend/internal/handlers"
	"backend/internal/storage"
)

func main() {
	// Carregar tarefas do arquivo JSON se existir
	loadTasksFromFile()

	// Rotas
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/tasks/", handlers.TaskByIDHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s", port)
	log.Printf("Acesse http://localhost:%s/tasks", port)
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}

// loadTasksFromFile carrega tarefas do arquivo JSON
func loadTasksFromFile() {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Arquivo tasks.json não encontrado. Iniciando com armazenamento em memória.")
			return
		}
		log.Printf("Erro ao abrir arquivo: %v", err)
		return
	}
	defer file.Close()

	var tasks []storage.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		return
	}

	storage.Store.Mu.Lock()
	defer storage.Store.Mu.Unlock()

	storage.Store.Tasks = tasks
	if len(tasks) > 0 {
		maxID := 0
		for _, task := range tasks {
			if task.ID > maxID {
				maxID = task.ID
			}
		}
		storage.Store.NextID = maxID + 1
	}

	log.Printf("Carregadas %d tarefas do arquivo tasks.json", len(tasks))
}
