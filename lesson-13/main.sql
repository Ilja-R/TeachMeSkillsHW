CREATE TABLE IF NOT EXISTS users
(
    id    SERIAL,
    name  VARCHAR        NOT NULL DEFAULT 'Vasya',
    email VARCHAR UNIQUE NOT NULL,
    phone VARCHAR UNIQUE NOT NULL,
    age   INT            NOT NULL,
    PRIMARY KEY (id)
);

DROP TABLE users;

CREATE TABLE cars
(
    id      SERIAL,
    year    INT,
    model   VARCHAR,
    user_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE orders
(
    id SERIAL,
    user_id INT,
    product VARCHAR NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE DATABASE online_shop_db;

DROP TABLE users;

\c online_shop_db

CREATE TABLE IF NOT EXISTS employees
(
    id    SERIAL,
    name  VARCHAR        NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    age   INT            NOT NULL,
    PRIMARY KEY (id)
);

DROP TABLE employees;
