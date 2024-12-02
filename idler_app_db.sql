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

INSERT INTO "machines" ("description", "generation", "id", "increment", "name", "price", "priority", "type") VALUES ('Generates +5 cash per second', '5.00', '2', '110.00', 'Broken Drill', '10.00', 2, 'generator');
INSERT INTO "machines" ("description", "generation", "id", "increment", "name", "price", "priority", "type") VALUES ('Generates +1 cash per second', '1.00', '1', '105.00', 'Crude Pickaxe', '1.00', 1, 'generator');
INSERT INTO "machines" ("description", "generation", "id", "increment", "name", "price", "priority", "type") VALUES ('Generates +250 cash per second', '250.00', '5', '125.00', 'Quarry', '5000.00', 5, 'generator');
INSERT INTO "machines" ("description", "generation", "id", "increment", "name", "price", "priority", "type") VALUES ('Generates +50 cash per second', '50.00', '4', '120.00', 'Iron Drill', '500.00', 4, 'generator');
INSERT INTO "machines" ("description", "generation", "id", "increment", "name", "price", "priority", "type") VALUES ('Generates +500 cash per second', '500.00', '7', '125.00', 'Factory', '25000.00', 7, 'generator');
INSERT INTO "machines" ("description", "generation", "id", "increment", "name", "price", "priority", "type") VALUES ('Generates +2000 cash per second', '2000.00', '8', '150.00', 'Industrial Plant', '50000.00', 8, 'generator');
INSERT INTO "machines" ("description", "generation", "id", "increment", "name", "price", "priority", "type") VALUES ('Increases all machine''s generation by 100%', '100.00', '9', '200.00', 'Perfect Amplifier', '100000.00', 9, 'amplifier');
INSERT INTO "machines" ("description", "generation", "id", "increment", "name", "price", "priority", "type") VALUES ('Increases all machine''s generation by 5%', '5.00', '3', '105.00', 'Small Amplifier', '25.00', 3, 'amplifier');
INSERT INTO "machines" ("description", "generation", "id", "increment", "name", "price", "priority", "type") VALUES ('Increases all machine''s generation by 20%', '20.00', '6', '150.00', 'Industrial Amplifier', '7500.00', 6, 'amplifier');