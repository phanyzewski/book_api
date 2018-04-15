
ALTER TABLE IF EXISTS books
ADD COLUMN publisher_id integer,
ADD FOREIGN KEY (publisher_id) REFERENCES publishers(id);
