CREATE DATABASE book_development;
CREATE ROLE dev WITH superuser login;
GRANT ALL PRIVILEGES ON DATABASE book_development TO dev;

-- CREATE TABLE books (
--   id SERIAL PRIMARY KEY,
--   rating integer,
--   status integer,
--   title TEXT,
--   published_date TIMESTAMP WITHOUT TIME ZONE
-- );

-- CREATE TABLE authors (
--   id SERIAL PRIMARY KEY,
--   first_name TEXT,
--   last_name TEXT,
--   pen_name TEXT
-- );

-- CREATE TABLE publishers (
--   id SERIAL PRIMARY KEY,
--   name TEXT
-- );

-- ALTER TABLE IF EXISTS books
-- ADD COLUMN author_id integer,
-- ADD FOREIGN KEY (author_id) REFERENCES authors(id);

-- ALTER TABLE IF EXISTS books
-- ADD COLUMN publisher_id integer,
-- ADD FOREIGN KEY (publisher_id) REFERENCES publishers(id);
