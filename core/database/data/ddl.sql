-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	name varchar(50) NULL,
	email varchar(30) NULL,
	CONSTRAINT users_name_key UNIQUE (name),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);