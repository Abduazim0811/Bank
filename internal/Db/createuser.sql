CREATE TABLE if not exists users(
    id serial primary key, 
    firstname Varchar(50),
    lastname Varchar(50),
    fathers_name Varchar(50),
    email Varchar(64),
    Password Varchar(50),
    balans float64,
    status Varchar(50));