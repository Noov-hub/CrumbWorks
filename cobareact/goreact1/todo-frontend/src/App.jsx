// src/App.jsx
import { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

// Tentukan URL base dari API Go Anda
const API_URL = 'http://localhost:8080';

function App() {
  const [todos, setTodos] = useState([]);
  const [inputValue, setInputValue] = useState('');

  // READ: Mengambil data saat komponen dimuat
  useEffect(() => {
    axios.get(`${API_URL}/todos`)
      .then(response => {
        setTodos(response.data || []); // Pastikan response.data bukan null
      })
      .catch(error => console.error("Error fetching todos:", error));
  }, []);

  // CREATE: Menambahkan todo baru
  const handleAddTodo = (e) => {
    e.preventDefault();
    if (inputValue.trim() === '') return;

    axios.post(`${API_URL}/todos`, { title: inputValue, completed: false })
      .then(response => {
        setTodos([...todos, response.data]);
        setInputValue('');
      })
      .catch(error => console.error("Error adding todo:", error));
  };

  // UPDATE: Mengubah status todo
  const handleToggleStatus = (id) => {
    axios.patch(`${API_URL}/todos/${id}`)
      .then(response => {
        setTodos(todos.map(todo =>
          todo.id === id ? response.data : todo
        ));
      })
      .catch(error => console.error("Error updating todo:", error));
  };
  
  // DELETE: Menghapus todo
  const handleDeleteTodo = (id) => {
    axios.delete(`${API_URL}/todos/${id}`)
      .then(() => {
        setTodos(todos.filter(todo => todo.id !== id));
      })
      .catch(error => console.error("Error deleting todo:", error));
  };


  return (
    <div className="app-container">
      <h1>Go + React To-Do List</h1>
      
      {/* Form untuk Create */}
      <form onSubmit={handleAddTodo} className="todo-form">
        <input
          type="text"
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          placeholder="Apa yang ingin Anda lakukan?"
        />
        <button type="submit">Tambah</button>
      </form>

      {/* List untuk Read */}
      <ul className="todo-list">
        {todos.map(todo => (
          <li
            key={todo.id}
            className={todo.completed ? 'completed' : ''}
          >
            <span onClick={() => handleToggleStatus(todo.id)}>
              {todo.title}
            </span>
            <button onClick={() => handleDeleteTodo(todo.id)} className="delete-btn">Hapus</button>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default App