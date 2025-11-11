const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

const handleResponse = async (response) => {
  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || `Erro HTTP: ${response.status}`);
  }

  if (response.status === 204) {
    return null;
  }

  return response.json();
};

export const getTasks = async () => {
  const response = await fetch(`${API_URL}/tasks`);
  const data = await handleResponse(response);
  
  // Garante que o retorno seja sempre um array
  return Array.isArray(data) ? data : [];

};

export const createTask = async (taskData) => {
  const response = await fetch(`${API_URL}/tasks`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(taskData),
  });
  return handleResponse(response);
};

export const updateTask = async (taskId, taskData) => {
  const response = await fetch(`${API_URL}/tasks/${taskId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(taskData),
  });
  return handleResponse(response);
};

export const deleteTask = async (taskId) => {
  const response = await fetch(`${API_URL}/tasks/${taskId}`, {
    method: 'DELETE',
  });
  return handleResponse(response);
};

