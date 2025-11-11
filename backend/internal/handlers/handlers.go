package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"backend/internal/middleware"
	"backend/internal/storage"
)

// tasksHandler lida com GET /tasks e POST /tasks
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case "GET":
		getTasks(w, r)
	case "POST":
		createTask(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// taskByIDHandler lida com PUT /tasks/:id e DELETE /tasks/:id
func TaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Extrair ID da URL
	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	segments := strings.Split(path, "/")

	if len(segments) != 1 {
		http.Error(w, "URL inválida", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(segments[0])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "PUT":
		updateTask(w, r, id)
	case "DELETE":
		deleteTask(w, r, id)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// getTasks retorna todas as tarefas
func getTasks(w http.ResponseWriter, r *http.Request) {
	storage.Store.Mu.RLock()
	tasks := make([]storage.Task, len(storage.Store.Tasks))
	copy(tasks, storage.Store.Tasks)
	storage.Store.Mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
	}
}

const maxBodySize = 1 << 20 // 1MB

// createTask cria uma nova tarefa
func createTask(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	var task storage.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Validações
	if task.Title == "" {
		http.Error(w, "Título é obrigatório", http.StatusBadRequest)
		return
	}

	if task.Status == "" {
		task.Status = "todo"
	}

	if !isValidStatus(task.Status) {
		http.Error(w, "Status inválido. Use: todo, in_progress ou done", http.StatusBadRequest)
		return
	}

	// Criar tarefa
	storage.Store.Mu.Lock()
	task.ID = storage.Store.NextID
	storage.Store.NextID++
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	storage.Store.Tasks = append(storage.Store.Tasks, task)

	// Salvar tarefas no arquivo ainda sob proteção do mutex
	err := storage.SaveTasksToFile(storage.Store.Tasks)
    storage.Store.Mu.Unlock()

    if err != nil {
    	// Trate o erro fora da região crítica
    	log.Println("Erro ao salvar tarefas:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

// updateTask atualiza uma tarefa existente
func updateTask(w http.ResponseWriter, r *http.Request, id int) {
	var updatedTask storage.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Validações antes de travar o mutex
	if updatedTask.Title == "" {
		http.Error(w, "Título é obrigatório", http.StatusBadRequest)
		return
	}

	if !isValidStatus(updatedTask.Status) {
		http.Error(w, "Status inválido. Use: todo, in_progress ou done", http.StatusBadRequest)
		return
	}

	storage.Store.Mu.Lock()

	// Buscar tarefa
	found := false
	for i := range storage.Store.Tasks {
		if storage.Store.Tasks[i].ID == id {
			storage.Store.Tasks[i].Title = updatedTask.Title
			storage.Store.Tasks[i].Description = updatedTask.Description
			storage.Store.Tasks[i].Status = updatedTask.Status
			storage.Store.Tasks[i].UpdatedAt = time.Now()
			updatedTask = storage.Store.Tasks[i]
			found = true
			break
		}
	}
	storage.Store.Mu.Unlock()

	if !found {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	if err := storage.SaveTasksToFile(storage.Store.Tasks); err != nil {
		log.Println("Erro ao salvar tarefas:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

// deleteTask remove uma tarefa
func deleteTask(w http.ResponseWriter, r *http.Request, id int) {
	storage.Store.Mu.Lock()
	found := false
	for i, task := range storage.Store.Tasks {
		if task.ID == id {
			storage.Store.Tasks = append(storage.Store.Tasks[:i], storage.Store.Tasks[i+1:]...)
			found = true
			break
		}
	}
	storage.Store.Mu.Unlock()

	if !found {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	if err := storage.SaveTasksToFile(storage.Store.Tasks); err != nil {
		log.Println("Erro ao salvar tarefas:", err)
	}

	w.WriteHeader(http.StatusNoContent)
}

// isValidStatus verifica se o status é válido
func isValidStatus(status string) bool {
	validStatuses := []string{"todo", "in_progress", "done"}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}

