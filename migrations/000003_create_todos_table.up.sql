CREATE TABLE IF NOT EXISTS todos (
    id bigserial PRIMARY KEY,
    task text NOT NULL,
    completed boolean DEFAULT false,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
);
