CREATE TABLE card_transactions (
    id          	  	TEXT PRIMARY KEY,
	external_id	 		TEXT,
	source 		 		TEXT,
    description			TEXT NOT NULL,
	amount       		BIGINT NOT NULL DEFAULT 0,
	date         		TIMESTAMPTZ NOT NULL,
	category_id  	   	TEXT REFERENCES categories(id),
	invoice_id   	  	TEXT REFERENCES invoices(id) NOT NULL,
	installment_number	INT,
	total_installments	INT
);

CREATE UNIQUE INDEX card_transactions_external_unique
	ON card_transactions (external_id, source)
	WHERE external_id IS NOT NULL;