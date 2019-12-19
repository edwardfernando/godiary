CREATE TABLE entries (
    id UUID PRIMARY KEY,
    title TEXT,
    body TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
