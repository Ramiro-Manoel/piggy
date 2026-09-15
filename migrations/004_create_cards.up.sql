CREATE TABLE cards (
    id              TEXT PRIMARY KEY,
    external_id     TEXT,
    source          TEXT,
    name            TEXT,
    brand           TEXT NOT NULL,
    credit_limit    BIGINT NOT NULL DEFAULT 0,
    available_limit BIGINT NOT NULL DEFAULT 0,
	closing_day     INT NOT NULL,
	due_day         INT NOT NULL,
    account_id      TEXT
    );

CREATE UNIQUE INDEX card_external_unique
	ON cards (external_id, source)
	WHERE external_id IS NOT NULL;