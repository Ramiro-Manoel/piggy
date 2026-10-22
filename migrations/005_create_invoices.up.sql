CREATE TABLE invoices (
    id           TEXT PRIMARY KEY,
    card_id      TEXT REFERENCES cards(id) NOT NULL,
    close_date   DATE NOT NULL,
    pay_date     DATE NOT NULL,
    amount       BIGINT NOT NULL DEFAULT 0,
    paid_at      DATE
);

CREATE UNIQUE INDEX invoices_card_close_date_unique
	ON invoices (card_id, close_date);
