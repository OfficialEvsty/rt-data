CREATE TABLE IF NOT EXISTS outbox (
    id UUID PRIMARY KEY,
    type INTEGER NOT NULL,
    payload JSONB NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    published_at TIMESTAMPTZ NULL,
    source UUID NULL
);

-- Вполне могут быть кейсы с поиском порождающего события --
CREATE INDEX idx_outbox_source
ON outbox (source);

CREATE TABLE IF NOT EXISTS requests (
    id UUID PRIMARY KEY,
    name TEXT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);