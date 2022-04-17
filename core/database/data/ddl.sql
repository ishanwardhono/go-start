-- articles definition

-- Drop table

-- DROP TABLE public.articles;

CREATE TABLE IF NOT EXISTS articles (
	id serial4 NOT NULL,
	title varchar(50) NOT NULL,
	content text NULL,
	author varchar(30) NOT NULL,
	created_time timestamp,
	created_by varchar(30),
	modified_time timestamp,
	modified_by timestamp,
	CONSTRAINT articles_pkey PRIMARY KEY (id)
);