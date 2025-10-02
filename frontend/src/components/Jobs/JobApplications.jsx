import { useState, useEffect } from 'react';
import { applicationAPI } from '../../services/api';
import './Jobs.css';

function JobApplications({ jobId, jobTitle }) {
  const [applications, setApplications] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [message, setMessage] = useState('');

  useEffect(() => {
    if (jobId) {
      fetchApplications();
    }
  }, [jobId]);

  const fetchApplications = async () => {
    try {
      const response = await applicationAPI.getJobApplications(jobId);
      setApplications(response.data.applications || []);
      setLoading(false);
    } catch (err) {
      setError('Failed to load applications');
      setLoading(false);
    }
  };

  const handleStatusUpdate = async (applicationId, status) => {
    try {
      await applicationAPI.updateApplicationStatus(applicationId, status);
      setMessage(`Application ${status} successfully!`);
      // Refresh applications
      fetchApplications();
      setTimeout(() => setMessage(''), 3000);
    } catch (err) {
      setMessage('Failed to update application status');
    }
  };

  const getStatusBadge = (status) => {
    const statusColors = {
      pending: 'badge-pending',
      accepted: 'badge-accepted',
      rejected: 'badge-rejected',
    };
    return statusColors[status] || 'badge-pending';
  };

  if (loading) return <div>Loading applications...</div>;
  if (error) return <div className="error">{error}</div>;

  return (
    <div className="job-applications-container">
      <h2>Applications for: {jobTitle}</h2>
      <p>Total Applications: {applications.length}</p>

      {message && <div className="success">{message}</div>}

      {applications.length === 0 ? (
        <p>No applications received yet.</p>
      ) : (
        <div className="applications-list">
          {applications.map((app) => (
            <div key={app.id} className="application-card employer-view">
              <div className="application-header">
                <div>
                  <h3>{app.worker_name}</h3>
                  <p className="worker-contact">
                    📧 {app.worker_email}
                    {app.worker_phone && ` | 📞 ${app.worker_phone}`}
                  </p>
                  {app.worker_location && (
                    <p>📍 {app.worker_location}</p>
                  )}
                </div>
                <span className={`status-badge ${getStatusBadge(app.status)}`}>
                  {app.status.toUpperCase()}
                </span>
              </div>

              {app.cover_letter && (
                <div className="cover-letter-full">
                  <strong>Cover Letter:</strong>
                  <p>{app.cover_letter}</p>
                </div>
              )}

              <p className="applied-date">
                Applied on: {new Date(app.applied_at).toLocaleString()}
              </p>

              {app.status === 'pending' && (
                <div className="action-buttons">
                  <button
                    onClick={() => handleStatusUpdate(app.id, 'accepted')}
                    className="accept-btn"
                  >
                    Accept
                  </button>
                  <button
                    onClick={() => handleStatusUpdate(app.id, 'rejected')}
                    className="reject-btn"
                  >
                    Reject
                  </button>
                </div>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default JobApplications;
