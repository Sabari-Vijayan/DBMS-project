import { useState } from 'react';
import Login from './components/Auth/Login';
import Register from './components/Auth/Register';
import Profile from './components/Profile/Profile';
import CreateJob from './components/Jobs/CreateJob';
import JobList from './components/Jobs/JobList';
import './App.css';

function App() {
  const [showLogin, setShowLogin] = useState(true);
  const [view, setView] = useState('auth'); // 'auth', 'profile', 'createJob', 'jobList'

  return (
    <div className="App">
      <h1>Job Board Platform</h1>
      
      <div className="nav-buttons">
        <button onClick={() => setView('auth')}>Auth</button>
        <button onClick={() => setView('profile')}>Profile</button>
        <button onClick={() => setView('createJob')}>Post Job</button>
        <button onClick={() => setView('jobList')}>Browse Jobs</button>
      </div>

      {view === 'auth' && (
        <>
          <div className="auth-toggle">
            <button onClick={() => setShowLogin(true)}>Login</button>
            <button onClick={() => setShowLogin(false)}>Register</button>
          </div>
          {showLogin ? <Login /> : <Register />}
        </>
      )}

      {view === 'profile' && <Profile userId={1} />}
      {view === 'createJob' && <CreateJob employerId={5} />} {/* Use actual employer ID */}
      {view === 'jobList' && <JobList />}
    </div>
  );
}

export default App;
