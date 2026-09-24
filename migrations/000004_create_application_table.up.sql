CREATE TABLE IF NOT EXISTS applications(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
job_id UUID NOT NULL REFERENCES jobs(id) ON DELETE CASCADE,
resume_id UUID REFERENCES resumes(id) ON DELETE CASCADE,
status VARCHAR(50) NOT NULL,
notes TEXT,
applied_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP ,
CHECK (status IN ('applied',
                'interview',
                'offer',
                'rejected',
                'withdrawn',
                'accepted')),
UNIQUE(user_id,job_id)
);