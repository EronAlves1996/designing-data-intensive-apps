BEGIN;

CREATE TABLE branch (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL 
);

CREATE TABLE status (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE book (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  publication_year INTEGER NOT NULL
);

CREATE TABLE book_copy (
  id SERIAL PRIMARY KEY,
  book_id INTEGER NOT NULL REFERENCES book(id),
  branch_id INTEGER NOT NULL REFERENCES branch(id),
  status_id INTEGER NOT NULL REFERENCES status(id)
);

CREATE TABLE author (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE author_book (
  book_id INTEGER NOT NULL REFERENCES book (id),
  author_id INTEGER NOT NULL REFERENCES author(id),
  PRIMARY KEY (book_id, author_id)
);

CREATE TABLE tag (
  -- Particularly, the name should be the primary key for tag, 
  -- but gonna use the surrogate key for simplicity
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE
);

CREATE TABLE book_tag (
  book_id INTEGER REFERENCES book(id),
  tag_id INTEGER REFERENCES tag(id),
  PRIMARY KEY (book_id, tag_id)
);

COMMIT;
