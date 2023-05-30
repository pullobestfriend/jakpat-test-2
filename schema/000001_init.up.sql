CREATE TABLE users
(
    id serial not null unique,
    name varchar not null unique,
    password varchar not null,
    role int not null,
    join_date timestamp DEFAULT NOW(),
    status int not null
);

CREATE TABLE items
(
    id serial not null unique,
    seller_id int not null,
    name varchar not null,
    stock int not null DEFAULT 0,
    status int not null
);

CREATE TABLE orders
(
    id serial not null unique,
    item_id int not null,
    buyer_id int not null,
    status int not null,
    expired_date timestamp not null,
    created_date timestamp DEFAULT NOW(),
    last_updated timestamp DEFAULT NOW()
);