CREATE TABLE transactions (
    id           TEXT PRIMARY KEY,
	external_id	 TEXT ,
	source 		 TEXT NOT NULL,
    description  TEXT NOT NULL,
	amount       BIGINT NOT NULL DEFAULT 0,
	date         TIMESTAMPTZ NOT NULL,
	category_id  TEXT REFERENCES categories(id),
	account_id   TEXT REFERENCES accounts(id) NOT NULL
);

CREATE UNIQUE INDEX transactions_external_unique
	ON transactions (external_id, source)
	WHERE external_id IS NOT NULL;