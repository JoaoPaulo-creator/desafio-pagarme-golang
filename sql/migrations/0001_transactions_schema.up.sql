CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL ,
    transaction_value FLOAT,
    product_description TEXT,
    card_number BIGINT,
    name_in_card VARCHAR(50),
    card_expiration_date DATE,
    cvv INT
);
