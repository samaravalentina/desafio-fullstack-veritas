import React from 'react';
import TaskCard from './TaskCard';
import './KanbanBoard.css';

const KanbanBoard = ({ tasks, onMoveTask, onEditTask, onDeleteTask }) => {
  const columns = [
    { id: 'todo', title: 'A Fazer', status: 'todo' },
    { id: 'in_progress', title: 'Em Progresso', status: 'in_progress' },
    { id: 'done', title: 'Concluídas', status: 'done' }
  ];

  const getTasksByStatus = (status) => {
    return tasks.filter(task => task.status === status);
  };

  return (
    <div className="kanban-board">
      {columns.map(column => (
        <div key={column.id} className="kanban-column">
          <div className="column-header">
            <h2>{column.title}</h2>
            <span className="task-count">{getTasksByStatus(column.status).length}</span>
          </div>
          <div className="column-content">
            {getTasksByStatus(column.status).map(task => (
              <TaskCard
                key={task.id}
                task={task}
                onMoveTask={onMoveTask}
                onEditTask={onEditTask}
                onDeleteTask={onDeleteTask}
              />
            ))}
            {getTasksByStatus(column.status).length === 0 && (
              <div className="empty-column">Nenhuma tarefa</div>
            )}
          </div>
        </div>
      ))}
    </div>
  );
};

export default KanbanBoard;

