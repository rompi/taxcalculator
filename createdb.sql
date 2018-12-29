CREATE DATABASE "tax";

\c tax 

CREATE TABLE tax_object (
	id serial NOT NULL,
	"name" varchar NULL,
	tax_code int4 NULL,
	price int4 NULL,
	CONSTRAINT tax_object_pkey PRIMARY KEY (id)
);
