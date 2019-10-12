CREATE TABLE posts.posts
(
    id serial NOT NULL,
    name "char" NOT NULL,
    content character varying
    [] NOT NULL,
    "createdAt" date NOT NULL,
    "updatedAt" date NOT NULL,
    "imagesLink" "char"[],
    PRIMARY KEY
    (id),
    CONSTRAINT name UNIQUE
    (name)

);

    ALTER TABLE posts.posts
    OWNER to postgres;