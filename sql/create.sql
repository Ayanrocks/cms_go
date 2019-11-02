
CREATE TABLE posts.posts
(
    id integer NOT NULL DEFAULT nextval('posts.posts_id_seq'
    ::regclass),
    name character varying
    (100) COLLATE pg_catalog."default" NOT NULL,
    "createdAt" timestamp
    (4) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" timestamp
    (4) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "imagesLink" "char"[],
    content character varying
    (1000) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT posts_pkey PRIMARY KEY
    (id),
    CONSTRAINT name UNIQUE
    (name)

)
    WITH
    (
    OIDS = FALSE
)
TABLESPACE pg_default;

    ALTER TABLE posts.posts
    OWNER to ayanrocks;

    GRANT ALL ON TABLE posts.posts TO ayanrocks;