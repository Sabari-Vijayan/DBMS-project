import { useState } from 'react';
import Login from './components/Auth/Login';
import Register from './components/Auth/Register';
import Profile from './components/Profile/Profile';
import './App.css';

function App() {
  const [showLogin, setShowLogin] = useState(true);
  const [loggedInUserId, setLoggedInUserId] = useState(1);

  return (
    <div className="App">
      <h1>Handyman Gig Platform</h1>
      
      <div className="auth-toggle">
        <button onClick={() => setShowLogin(true)}>Login</button>
        <button onClick={() => setShowLogin(false)}>Register</button>
      </div>

      {showLogin ? <Login /> : <Register />}

    <hr />

    <Profile userId={9}/>

    </div>
  );
}

export default App;
