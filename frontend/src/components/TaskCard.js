import React from 'react';
import './TaskCard.css';

const TaskCard = ({ task, onMoveTask, onEditTask, onDeleteTask }) => {
  const statusOptions = [
    { value: 'todo', label: 'A Fazer' },
    { value: 'in_progress', label: 'Em Progresso' },
    { value: 'done', label: 'Concluída' }
  ];

  const handleStatusChange = (e) => {
    const newStatus = e.target.value;
    if (newStatus !== task.status) {
      onMoveTask(task.id, newStatus);
    }
  };

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('pt-BR', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  return (
    <div className="task-card">
      <div className="task-header">
        <h3 className="task-title">{task.title}</h3>
        <div className="task-actions">
          <button
            className="btn-icon"
            onClick={() => onEditTask(task)}
            title="Editar tarefa"
          >
            ✏️
          </button>
          <button
            className="btn-icon"
            onClick={() => onDeleteTask(task.id)}
            title="Excluir tarefa"
          >
            🗑️
          </button>
        </div>
      </div>

      {task.description && (
        <p className="task-description">{task.description}</p>
      )}

      <div className="task-footer">
        <label htmlFor="status" className="status-label">Alterar Status:</label>
        <select
          className="status-select"
          value={task.status}
          onChange={handleStatusChange}
        >
          {statusOptions.map(option => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>

        <div className="task-date">
          <small>Criada em: {formatDate(task.createdAt)}</small>
        </div>
      </div>
    </div>
  );
};

export default TaskCard;

