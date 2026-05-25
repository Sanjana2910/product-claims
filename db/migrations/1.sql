CREATE TABLE IF NOT EXISTS products(
    id TEXT PRIMARY KEY,
    mcr_id TEXT,
    title TEXT NOT NULL,
    brand TEXT,
    category_name TEXT,
    raw_json JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS workflows (
    id          TEXT PRIMARY KEY,
    product_id  TEXT NOT NULL REFERENCES products(id),
    status      TEXT NOT NULL DEFAULT 'IN_PROGRESS',
    error_msg   TEXT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS claims (
    id           TEXT PRIMARY KEY,
    workflow_id  TEXT NOT NULL REFERENCES workflows(id),
    product_id   TEXT NOT NULL REFERENCES products(id),
    claim_type   TEXT NOT NULL,
    claim_value  TEXT NOT NULL,
    status       TEXT NOT NULL DEFAULT 'IDENTIFIED',
    created_at   TIMESTAMP DEFAULT NOW()
);
