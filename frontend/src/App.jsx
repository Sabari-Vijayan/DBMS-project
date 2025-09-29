import { useState } from 'react';
import Login from './components/Auth/Login';
import Register from './components/Auth/Register';
import './App.css';

function App() {
  const [showLogin, setShowLogin] = useState(true);

  return (
    <div className="App">
      <h1>Handyman Gig Platform</h1>
      
      <div className="auth-toggle">
        <button onClick={() => setShowLogin(true)}>Login</button>
        <button onClick={() => setShowLogin(false)}>Register</button>
      </div>

      {showLogin ? <Login /> : <Register />}
    </div>
  );
}

export default App;
