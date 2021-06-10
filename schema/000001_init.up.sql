CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    email          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE comments
(
    id      serial not null unique,
    id_user int references users (id) on delete cascade not null,
    txt varchar(255)
);
