-- Categories/Skills for jobs
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Jobs posted by employers
CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    employer_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    category_id INTEGER REFERENCES categories(id),
    location VARCHAR(255) NOT NULL,
    salary_min DECIMAL(10,2),
    salary_max DECIMAL(10,2),
    duration VARCHAR(50), -- e.g., "2 hours", "1 day", "1 week"
    requirements TEXT, -- Skills/requirements needed
    contact_phone VARCHAR(20),
    contact_email VARCHAR(255),
    expires_at TIMESTAMP NOT NULL, -- Job expires after 2-3 days
    status VARCHAR(20) DEFAULT 'open' CHECK (status IN ('open', 'closed', 'filled')),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Applications - when workers apply to jobs
CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    job_id INTEGER REFERENCES jobs(id) ON DELETE CASCADE,
    worker_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    cover_letter TEXT, -- Why they're a good fit
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'accepted', 'rejected', 'withdrawn')),
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(job_id, worker_id) -- Worker can only apply once per job
);

-- Worker Skills - multiple skills per worker
CREATE TABLE worker_skills (
    id SERIAL PRIMARY KEY,
    worker_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES categories(id) ON DELETE CASCADE,
    experience_years INTEGER DEFAULT 0,
    UNIQUE(worker_id, category_id)
);

-- Work Experience - past gigs completed
CREATE TABLE work_experience (
    id SERIAL PRIMARY KEY,
    worker_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    job_title VARCHAR(255) NOT NULL,
    employer_name VARCHAR(255), -- Can be from platform or external
    description TEXT,
    duration VARCHAR(50), -- e.g., "3 months", "1 week"
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default categories/skills
INSERT INTO categories (name, description) VALUES 
('Retail & Sales', 'Store assistance, customer service, cashier'),
('Food Service', 'Waiter, cook, food delivery, kitchen help'),
('Delivery & Logistics', 'Package delivery, courier services'),
('Tutoring', 'Teaching, homework help, subject tutoring'),
('Data Entry', 'Office work, typing, admin tasks'),
('Cleaning', 'House cleaning, office cleaning, janitorial'),
('Event Help', 'Event setup, coordination, assistance'),
('Tech Support', 'Computer repair, software help, IT assistance'),
('Content Creation', 'Writing, social media, graphic design'),
('General Labor', 'Moving, packing, physical work');

-- Indexes for better performance
CREATE INDEX idx_jobs_employer ON jobs(employer_id);
CREATE INDEX idx_jobs_category ON jobs(category_id);
CREATE INDEX idx_jobs_status ON jobs(status);
CREATE INDEX idx_jobs_expires ON jobs(expires_at);
CREATE INDEX idx_applications_job ON applications(job_id);
CREATE INDEX idx_applications_worker ON applications(worker_id);
CREATE INDEX idx_applications_status ON applications(status);
CREATE INDEX idx_worker_skills_worker ON worker_skills(worker_id);
CREATE INDEX idx_work_experience_worker ON work_experience(worker_id);
