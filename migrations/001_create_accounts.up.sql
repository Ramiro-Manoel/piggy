  CREATE TABLE accounts (
      id      TEXT PRIMARY KEY,
      name    TEXT NOT NULL,
      number  TEXT NOT NULL,
      owner   TEXT NOT NULL,
      balance BIGINT NOT NULL DEFAULT 0
  );