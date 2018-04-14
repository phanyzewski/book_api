CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title TEXT,
  published_date TEXT,
  rating NUMERIC,
  book_available TEXT
  -- publisher TEXT,
  -- author TEXT
);

CREATE TABLE authors (
  id SERIAL PRIMARY KEY,
  first_name TEXT,
  last_name TEXT,
  pen_name TEXT
);
CREATE TABLE publishers (
  id SERIAL PRIMARY KEY,
  name TEXT
);
