import { useState } from 'react';
import { jobAPI } from '../../services/api';
import { useAuth } from '../../context/AuthContext';
import './Jobs.css';

function CreateJob() {
  const { user } = useAuth();
  
  const [formData, setFormData] = useState({
    // Remove employer_id - comes from JWT now
    title: '',
    description: '',
    category_id: '',
    location: '',
    salary_min: '',
    salary_max: '',
    duration: '',
    requirements: '',
    contact_phone: '',
    contact_email: '',
    expiry_days: 3,
  });
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const categories = [
    { id: 1, name: 'Retail & Sales' },
    { id: 2, name: 'Food Service' },
    { id: 3, name: 'Delivery & Logistics' },
    { id: 4, name: 'Tutoring' },
    { id: 5, name: 'Data Entry' },
    { id: 6, name: 'Cleaning' },
    { id: 7, name: 'Event Help' },
    { id: 8, name: 'Tech Support' },
    { id: 9, name: 'Content Creation' },
    { id: 10, name: 'General Labor' },
  ];

  const handleChange = (e) => {
    const value = e.target.type === 'number' ? 
      (e.target.value === '' ? '' : Number(e.target.value)) : 
      e.target.value;
    
    setFormData({
      ...formData,
      [e.target.name]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    try {
      // Prepare data - JWT will provide employer_id
      const jobData = {
        ...formData,
        category_id: formData.category_id ? Number(formData.category_id) : null,
        salary_min: formData.salary_min ? Number(formData.salary_min) : null,
        salary_max: formData.salary_max ? Number(formData.salary_max) : null,
        duration: formData.duration || null,
        requirements: formData.requirements || null,
        contact_phone: formData.contact_phone || null,
        contact_email: formData.contact_email || null,
      };

      const response = await jobAPI.createJob(jobData);
      setSuccess('Job posted successfully!');
      console.log('Job created:', response.data);
      
      // Reset form
      setFormData({
        title: '',
        description: '',
        category_id: '',
        location: '',
        salary_min: '',
        salary_max: '',
        duration: '',
        requirements: '',
        contact_phone: '',
        contact_email: '',
        expiry_days: 3,
      });
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to create job');
      console.error('Error:', err.response?.data);
    }
  };

  if (!user || user.user_type !== 'employer') {
    return <div className="error">Only employers can post jobs</div>;
  }

  return (
    <div className="create-job-container">
      <h2>Post a New Job</h2>
      {error && <div className="error">{error}</div>}
      {success && <div className="success">{success}</div>}

      <form onSubmit={handleSubmit} className="job-form">
        {/* Rest of the form stays the same */}
        <div className="form-group">
          <label>Job Title *</label>
          <input
            type="text"
            name="title"
            value={formData.title}
            onChange={handleChange}
            placeholder="e.g., Part-time Retail Assistant"
            required
          />
        </div>

        <div className="form-group">
          <label>Description *</label>
          <textarea
            name="description"
            value={formData.description}
            onChange={handleChange}
            rows="4"
            placeholder="Describe the job duties and what you're looking for..."
            required
          />
        </div>

        <div className="form-row">
          <div className="form-group">
            <label>Category</label>
            <select
              name="category_id"
              value={formData.category_id}
              onChange={handleChange}
            >
              <option value="">Select Category</option>
              {categories.map((cat) => (
                <option key={cat.id} value={cat.id}>
                  {cat.name}
                </option>
              ))}
            </select>
          </div>

          <div className="form-group">
            <label>Location *</label>
            <input
              type="text"
              name="location"
              value={formData.location}
              onChange={handleChange}
              placeholder="e.g., Bangalore, MG Road"
              required
            />
          </div>
        </div>

        <div className="form-row">
          <div className="form-group">
            <label>Salary Min (₹/hour)</label>
            <input
              type="number"
              name="salary_min"
              value={formData.salary_min}
              onChange={handleChange}
              placeholder="200"
              min="0"
            />
          </div>

          <div className="form-group">
            <label>Salary Max (₹/hour)</label>
            <input
              type="number"
              name="salary_max"
              value={formData.salary_max}
              onChange={handleChange}
              placeholder="500"
              min="0"
            />
          </div>
        </div>

        <div className="form-group">
          <label>Duration</label>
          <input
            type="text"
            name="duration"
            value={formData.duration}
            onChange={handleChange}
            placeholder="e.g., 2 hours per day, 1 week, flexible"
          />
        </div>

        <div className="form-group">
          <label>Requirements</label>
          <textarea
            name="requirements"
            value={formData.requirements}
            onChange={handleChange}
            rows="3"
            placeholder="Skills, experience, or qualifications needed..."
          />
        </div>

        <div className="form-row">
          <div className="form-group">
            <label>Contact Phone</label>
            <input
              type="tel"
              name="contact_phone"
              value={formData.contact_phone}
              onChange={handleChange}
              placeholder="9876543210"
            />
          </div>

          <div className="form-group">
            <label>Contact Email</label>
            <input
              type="email"
              name="contact_email"
              value={formData.contact_email}
              onChange={handleChange}
              placeholder="contact@example.com"
            />
          </div>
        </div>

        <div className="form-group">
          <label>Job Expires In (Days) *</label>
          <select
            name="expiry_days"
            value={formData.expiry_days}
            onChange={handleChange}
            required
          >
            <option value="1">1 day</option>
            <option value="2">2 days</option>
            <option value="3">3 days</option>
            <option value="5">5 days</option>
            <option value="7">7 days</option>
          </select>
        </div>

        <button type="submit" className="submit-btn">Post Job</button>
      </form>
    </div>
  );
}

export default CreateJob;
