import { useState, useEffect } from 'react';
import { jobAPI } from '../../services/api';
import './Jobs.css';

function JobList() {
  const [jobs, setJobs] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [selectedJob, setSelectedJob] = useState(null);

  useEffect(() => {
    fetchJobs();
  }, []);

  const fetchJobs = async () => {
    try {
      const response = await jobAPI.getAllJobs();
      setJobs(response.data.jobs || []);
      setLoading(false);
    } catch (err) {
      setError('Failed to load jobs');
      setLoading(false);
    }
  };

  const viewJobDetails = (job) => {
    setSelectedJob(job);
  };

  const closeJobDetails = () => {
    setSelectedJob(null);
  };

  if (loading) return <div>Loading jobs...</div>;
  if (error) return <div className="error">{error}</div>;

  return (
    <div className="job-list-container">
      <h2>Available Jobs ({jobs.length})</h2>

      {jobs.length === 0 ? (
        <p>No jobs available at the moment.</p>
      ) : (
        <div className="jobs-grid">
          {jobs.map((job) => (
            <div key={job.id} className="job-card">
              <h3>{job.title}</h3>
              <p className="employer">Posted by: {job.employer_name}</p>
              <p className="location">📍 {job.location}</p>
              {job.category_name && <span className="category">{job.category_name}</span>}
              {job.salary_min && job.salary_max && (
                <p className="salary">
                  ₹{job.salary_min} - ₹{job.salary_max} per hour
                </p>
              )}
              <p className="description">{job.description.substring(0, 100)}...</p>
              <p className="expires">
                Expires: {new Date(job.expires_at).toLocaleDateString()}
              </p>
              <button onClick={() => viewJobDetails(job)} className="view-btn">
                View Details
              </button>
            </div>
          ))}
        </div>
      )}

      {/* Job Details Modal */}
      {selectedJob && (
        <div className="modal-overlay" onClick={closeJobDetails}>
          <div className="modal-content" onClick={(e) => e.stopPropagation()}>
            <button className="close-btn" onClick={closeJobDetails}>×</button>
            <h2>{selectedJob.title}</h2>
            <p><strong>Employer:</strong> {selectedJob.employer_name}</p>
            <p><strong>Location:</strong> {selectedJob.location}</p>
            {selectedJob.category_name && (
              <p><strong>Category:</strong> {selectedJob.category_name}</p>
            )}
            {selectedJob.salary_min && selectedJob.salary_max && (
              <p><strong>Salary:</strong> ₹{selectedJob.salary_min} - ₹{selectedJob.salary_max} per hour</p>
            )}
            {selectedJob.duration && (
              <p><strong>Duration:</strong> {selectedJob.duration}</p>
            )}
            <div className="description-full">
              <strong>Description:</strong>
              <p>{selectedJob.description}</p>
            </div>
            {selectedJob.requirements && (
              <div>
                <strong>Requirements:</strong>
                <p>{selectedJob.requirements}</p>
              </div>
            )}
            <div className="contact-info">
              <strong>Contact:</strong>
              {selectedJob.contact_phone && <p>📞 {selectedJob.contact_phone}</p>}
              {selectedJob.contact_email && <p>📧 {selectedJob.contact_email}</p>}
            </div>
            <p className="expires-detail">
              <strong>Expires:</strong> {new Date(selectedJob.expires_at).toLocaleString()}
            </p>
            <button className="apply-btn">Apply for this Job</button>
          </div>
        </div>
      )}
    </div>
  );
}

export default JobList;
