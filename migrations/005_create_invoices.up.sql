CREATE TABLE invoices (
    id           TEXT PRIMARY KEY,
    card_id      TEXT REFERENCES cards(id) NOT NULL
);