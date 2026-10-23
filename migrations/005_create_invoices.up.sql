CREATE TABLE invoices (
    id           TEXT PRIMARY KEY,
    external_id  TEXT,
    source       TEXT,
    card_id      TEXT REFERENCES cards(id) NOT NULL,
    close_date   DATE NOT NULL,
    pay_date     DATE NOT NULL,
    amount       BIGINT NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX invoices_card_close_date_unique
	ON invoices (card_id, close_date);

CREATE UNIQUE INDEX invoices_external_unique
	ON invoices (external_id, source)
	WHERE external_id IS NOT NULL;
