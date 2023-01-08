CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name       VARCHAR(127)       NOT NULL UNIQUE,
    avatar     VARCHAR(255)       NOT NULL,
    created_at TIMESTAMPTZ        NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ        NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS snippets
(
    id           SERIAL PRIMARY KEY NOT NULL UNIQUE,
    link         UUID               NOT NULL UNIQUE DEFAULT gen_random_uuid(),
    title        VARCHAR(255)       NOT NULL,
    code         TEXT               NOT NULL,
    author_id    INTEGER            NOT NULL,
    is_anonymous BOOLEAN            NOT NULL        DEFAULT FALSE,
    created_at   TIMESTAMPTZ        NOT NULL        default now(),
    updated_at   TIMESTAMPTZ        NOT NULL        default now(),
    CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES users (id)
);
