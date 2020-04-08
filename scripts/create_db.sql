CREATE TABLE IF NOT EXISTS http_users (
	id SERIAL PRIMARY KEY,
	name text NOT NULL CHECK(length(name)>0),
	lastname text NOT NULL CHECK(length(lastname)>0),
	birthdate timestamp with time zone NOT NULL
);