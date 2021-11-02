CREATE TABLE IF NOT EXISTS priority(
    id serial NOT NULL,
    name varchar(256) NOT NULL,
    weight integer NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO priority (id, name, weight)
SELECT 1, 'Low', 1
WHERE NOT EXISTS (SELECT id FROM priority WHERE name = 'Low');

INSERT INTO priority (id, name, weight)
SELECT 2, 'Normal', 2
WHERE NOT EXISTS (SELECT id FROM priority WHERE name = 'Normal');

INSERT INTO priority (id, name, weight)
SELECT 3, 'High', 3
WHERE NOT EXISTS (SELECT id FROM priority WHERE name = 'High');

ALTER TABLE todo
ADD COLUMN IF NOT EXISTS
    priority_id integer REFERENCES priority NOT NULL DEFAULT 2
