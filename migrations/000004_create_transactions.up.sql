CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    sender_wallet_id BIGINT
        REFERENCES wallets(id),
    receiver_wallet_id BIGINT 
        REFERENCES wallets(id),
    transaction_type_id BIGINT NOT NULL
        REFERENCES transaction_types(id),
    amount BIGINT NOT NULL CHECK (amount > 0),
    status VARCHAR(20) NOT NULL DEFAULT 'SUCCESS',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);