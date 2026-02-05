

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name TEXT 
    last_name TEXT 
    user_name TEXT NOT NULL UNIQUE 
    password TEXT NOT NULL 
    email    TEXT NOT NULL UNIQUE 
    role    TEXT NOT NULL 
    is_verfified BOOLEAN NOT NULL DEFAULT false 
    is_deleted  BOOLEAN  NOT NULL DEFAULT false 
    deleted_at  TIMESTAMP
    created_at TIMESTAMP NOT NULL DEFAULT now()
    updated_at TIMESTAMP NOT NULL DEFAULT now()
)