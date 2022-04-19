CREATE TABLE IF NOT EXISTS users
(
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    balance integer NOT NULL,
    version integer NOT NULL DEFAULT 1
);

alter table users owner to wallet;

create unique index users_id_uindex
    on users (id);

-- CREATE TABLE IF NOT EXISTS transfers
-- (
--     id serial not null
--             primary key,
--     user_id_from integer NOT NULL
--         references users,
--     user_id_to integer NOT NULL
--         references users,
--     amount integer NOT NULL,
--     balance integer NOT NULL
-- );

-- alter table transfers owner to wallet;

-- create unique index transfers_id_uindex
--     on transfers (id);

INSERT INTO public.users (id, name, balance) VALUES (1, 'John K.', 1000);
INSERT INTO public.users (id, name, balance) VALUES (2, 'Lana V.', 100);
INSERT INTO public.users (id, name, balance) VALUES (3, 'Jimmy J.', 10);
INSERT INTO public.users (id, name, balance) VALUES (4, 'Betty P.', 10000);