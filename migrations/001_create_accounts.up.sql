  CREATE TABLE accounts (
      id            TEXT PRIMARY KEY,
      external_id   TEXT,
	  source 		TEXT,
      name          TEXT NOT NULL,
      number        TEXT NOT NULL,
      owner         TEXT NOT NULL,
      balance       BIGINT NOT NULL DEFAULT 0
  );

  CREATE UNIQUE INDEX accounts_external_unique
	ON accounts (external_id, source)
	WHERE external_id IS NOT NULL;