CREATE TABLE users (
    user_id         bigserial PRIMARY KEY,
	first_name 		varchar(30),
	last_name 	    varchar(30),
	email	        varchar(50),
    hashed_pass     varchar(255)
);