CREATE TABLE IF NOT EXISTS jobs(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    company_name VARCHAR(250) NOT NULL,
    position VARCHAR(250) NOT NULL,
    job_url VARCHAR(500),
    location VARCHAR(100),
    employment_type VARCHAR(100) NOT NULL,
    salary_min INTEGER CHECK(salary_min >=0),
    salary_max INTEGER CHECK(salary_max>=0),
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
CHECK(salary_min IS NULL OR salary_max IS NULL OR salary_max>=salary_min)


);