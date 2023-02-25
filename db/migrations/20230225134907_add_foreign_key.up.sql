SET statement_timeout = 0;

--bun:split

ALTER TABLE comments ADD FOREIGN KEY (post_id) REFERENCES posts(id);

