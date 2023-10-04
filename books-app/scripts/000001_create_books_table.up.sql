DROP TABLE IF EXISTS books;

CREATE TABLE books
(
    isbn   INTEGER,
    name   VARCHAR(50) NOT NULL,
    publisher VARCHAR(50) NOT NULL,
    CONSTRAINT books_pkey PRIMARY KEY (isbn)
)