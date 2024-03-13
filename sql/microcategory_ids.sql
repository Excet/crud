CREATE TABLE IF NOT EXISTS mc_schema.category(
    id SERIAL,
    type VARCHAR(60),
    name VARCHAR(60),
    parent INT,
    PRIMARY KEY(id)
);

