-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	title varchar(50) NOT NULL,
	content text NULL,
	author varchar(30) NOT NULL,
	created_time timestamp,
	created_by varchar(30),
	modified_time timestamp,
	modified_by timestamp,
	CONSTRAINT users_name_key UNIQUE (name),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);