CREATE TABLE users (
    id bigserial PRIMARY KEY,
    username text NOT NULL,
    password text NOT NULL,
    cash numeric(12, 2) DEFAULT 1,
    last_save timestamptz DEFAULT NOW()
);

CREATE TABLE machines (
    id bigserial PRIMARY KEY,
    priority integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    generation numeric(12, 2) NOT NULL,
    price numeric(12, 2) NOT NULL,
    increment numeric(12, 2) NOT NULL,
    type text NOT NULL
);

CREATE TABLE users_machines (
    user_id bigserial NOT NULL REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
    machine_id bigserial NOT NULL REFERENCES machines (id) ON UPDATE CASCADE ON DELETE CASCADE,
    level integer DEFAULT 0,
    CONSTRAINT pk_users_machines PRIMARY KEY (user_id, machine_id)
);