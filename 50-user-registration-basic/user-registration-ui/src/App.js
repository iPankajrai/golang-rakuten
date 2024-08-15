// src/App.js
import React, { useState, useEffect } from 'react';
import UserForm from './components/UserForm';
// import './App.css';


function App() {
    const [users, setUsers] = useState([]);
    const [editingUser, setEditingUser] = useState(null);

    useEffect(() => {
        fetch('/users')
            .then(response => response.json())
            .then(data => setUsers(data));
    }, []);

    const handleCreateUser = (user) => {
        fetch('http://localhost:8080/users', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(user)
        })
            .then(response => response.json())
            .then(newUser => setUsers([...users, newUser]));
    };

    const handleUpdateUser = (user) => {
        fetch(`http://localhost:8080/users/${user.id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(user)
        })
            .then(response => response.json())
            .then(updatedUser => {
                setUsers(users.map(u => (u.id === updatedUser.id ? updatedUser : u)));
                setEditingUser(null);
            });
    };

    const handleDeleteUser = (id) => {
        fetch(`http://localhost:8080/users/${id}`, { method: 'DELETE' })
            .then(() => setUsers(users.filter(user => user.id !== id)));
    };

    return (
        <div>
            <h1>User Registration</h1>
            <UserForm onSubmit={editingUser ? handleUpdateUser : handleCreateUser} user={editingUser} />
            <ul>
                {users.map(user => (
                    <li key={user.id}>
                        {user.first_name} {user.last_name}
                        <button onClick={() => setEditingUser(user)}>Edit</button>
                        <button onClick={() => handleDeleteUser(user.id)}>Delete</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default App;

