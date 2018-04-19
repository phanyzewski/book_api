ALTER TABLE IF EXISTS books
ADD COLUMN author_id integer,
ADD FOREIGN KEY (author_id) REFERENCES authors(id);
