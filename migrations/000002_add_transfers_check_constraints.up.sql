-- ALTER TABLE transfers ADD CONSTRAINT transfers_balance_check CHECK (balance >= 0);
ALTER TABLE users ADD CONSTRAINT users_balance_check CHECK (balance >= 0);