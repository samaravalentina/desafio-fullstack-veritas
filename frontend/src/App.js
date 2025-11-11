import React, { useState, useEffect } from 'react';
import KanbanBoard from './components/KanbanBoard';
import TaskForm from './components/TaskForm';
import { getTasks, createTask, updateTask, deleteTask } from './services/api';
import './App.css';

function App() {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showForm, setShowForm] = useState(false);
  const [editingTask, setEditingTask] = useState(null);

  useEffect(() => {
    loadTasks();
  }, []);

  const loadTasks = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await getTasks();
      console.log(data)
      setTasks(data);
      console.log(tasks)
    } catch (err) {
      setError('Erro ao carregar tarefas. Verifique se o backend está rodando.');
      console.error('Erro ao carregar tarefas:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleCreateTask = async (taskData) => {
    try {
      const newTask = await createTask(taskData);
      setTasks([...tasks, newTask]);
      setShowForm(false);
    } catch (err) {
      setError('Erro ao criar tarefa.');
      console.error('Erro ao criar tarefa:', err);
    }
  };

  const handleUpdateTask = async (taskId, taskData) => {
    try {
      const updatedTask = await updateTask(taskId, taskData);
      setTasks(tasks.map(task => task.id === taskId ? updatedTask : task));
      setEditingTask(null);
      setShowForm(false);
    } catch (err) {
      setError('Erro ao atualizar tarefa.');
      console.error('Erro ao atualizar tarefa:', err);
    }
  };

  const handleDeleteTask = async (taskId) => {
    if (!window.confirm('Tem certeza que deseja excluir esta tarefa?')) {
      return;
    }

    try {
      await deleteTask(taskId);
      setTasks(tasks.filter(task => task.id !== taskId));
    } catch (err) {
      setError('Erro ao excluir tarefa.');
      console.error('Erro ao excluir tarefa:', err);
    }
  };

  const handleMoveTask = async (taskId, newStatus) => {
    const task = tasks.find(t => t.id === taskId);
    if (!task || task.status === newStatus) return;

    try {
      const updatedTask = await updateTask(taskId, {
        ...task,
        status: newStatus
      });
      setTasks(tasks.map(t => t.id === taskId ? updatedTask : t));
    } catch (err) {
      setError('Erro ao mover tarefa.');
      console.error('Erro ao mover tarefa:', err);
    }
  };

  const handleEditTask = (task) => {
    setEditingTask(task);
    setShowForm(true);
  };

  const handleCancelForm = () => {
    setShowForm(false);
    setEditingTask(null);
  };

  if (loading) {
    return (
      <div className="app">
        <div className="loading">Carregando tarefas...</div>
      </div>
    );
  }

  return (
    <div className="app">
      <header className="app-header">
        <h1>📋 Kanban de Tarefas</h1>
        {!showForm && (
          <button 
            className="btn btn-primary" 
            onClick={() => setShowForm(true)}
          >
            + Nova Tarefa
          </button>
        )}
      </header>

      {error && (
        <div className="error-message">
          {error}
          <button onClick={() => setError(null)}>✕</button>
        </div>
      )}

      {showForm && (
        <TaskForm
          task={editingTask}
          onSubmit={editingTask ? 
            (data) => handleUpdateTask(editingTask.id, data) : handleCreateTask
          }
          onCancel={handleCancelForm}
        />
      )}

      <KanbanBoard
        tasks={tasks}
        onMoveTask={handleMoveTask}
        onEditTask={handleEditTask}
        onDeleteTask={handleDeleteTask}
      />
    </div>
  );
}

export default App;

