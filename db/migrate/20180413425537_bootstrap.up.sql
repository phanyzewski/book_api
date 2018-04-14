CREATE TABLE books (
  book_id SERIAL PRIMARY KEY,
  title TEXT,
  published_date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  rating NUMERIC,
  book_available TEXT
  -- publisher TEXT,
  -- author TEXT
);

CREATE TABLE authors (
  author_id SERIAL PRIMARY KEY,
  first_name TEXT,
  last_name TEXT,
  pen_name TEXT
);
CREATE TABLE publishers (
  publisher_id SERIAL PRIMARY KEY,
  name TEXT
);
