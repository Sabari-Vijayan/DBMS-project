import { useState, useEffect } from 'react';
import { applicationAPI } from '../../services/api';
import './Jobs.css';

function MyApplications({ workerId }) {
  const [applications, setApplications] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    if (workerId) {
      fetchApplications();
    }
  }, [workerId]);

  const fetchApplications = async () => {
    try {
      const response = await applicationAPI.getWorkerApplications(workerId);
      setApplications(response.data.applications || []);
      setLoading(false);
    } catch (err) {
      setError('Failed to load applications');
      setLoading(false);
    }
  };

  const getStatusBadge = (status) => {
    const statusColors = {
      pending: 'badge-pending',
      accepted: 'badge-accepted',
      rejected: 'badge-rejected',
      withdrawn: 'badge-withdrawn',
    };
    return statusColors[status] || 'badge-pending';
  };

  if (loading) return <div>Loading your applications...</div>;
  if (error) return <div className="error">{error}</div>;

  return (
    <div className="my-applications-container">
      <h2>My Applications ({applications.length})</h2>

      {applications.length === 0 ? (
        <p>You haven't applied to any jobs yet.</p>
      ) : (
        <div className="applications-list">
          {applications.map((app) => (
            <div key={app.id} className="application-card">
              <div className="application-header">
                <h3>{app.job_title}</h3>
                <span className={`status-badge ${getStatusBadge(app.status)}`}>
                  {app.status.toUpperCase()}
                </span>
              </div>
              <p><strong>Employer:</strong> {app.employer_name}</p>
              <p><strong>Location:</strong> {app.location}</p>
              {app.salary_min && app.salary_max && (
                <p><strong>Salary:</strong> ₹{app.salary_min} - ₹{app.salary_max} per hour</p>
              )}
              {app.cover_letter && (
                <div className="cover-letter-preview">
                  <strong>Your Cover Letter:</strong>
                  <p>{app.cover_letter}</p>
                </div>
              )}
              <p className="applied-date">
                Applied on: {new Date(app.applied_at).toLocaleString()}
              </p>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default MyApplications;
