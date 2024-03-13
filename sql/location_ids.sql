CREATE TABLE IF NOT EXISTS lc_schema.location(
    id SERIAL,
    type VARCHAR(60),
    name VARCHAR(60),
    parent INT,
    PRIMARY KEY(id)
);
