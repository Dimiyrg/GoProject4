CREATE TABLE users
(
    id           serial       PRIMARY KEY,
    email        varchar(255) not null,
    username     varchar(255) not null,
    password     varchar(255) not null
);