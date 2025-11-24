CREATE TABLE dev_user.users (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name      VARCHAR(100)     NOT NULL,
    last_name       VARCHAR(100)     NOT NULL,
    date_of_birth   DATE             NOT NULL,
    age             INT             NOT NULL,
    address         TEXT             NOT NULL,
    created_at      TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ      NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_id ON dev_user.users (id);