DROP TABLE IF EXISTS reviews;

CREATE TABLE reviews
(
    isbn   INTEGER,
    reviewer   VARCHAR(50) NOT NULL,
    comment VARCHAR(200),
    rating INTEGER,
    CONSTRAINT reviews_isbn_fkey FOREIGN KEY (isbn) REFERENCES books(isbn)
)