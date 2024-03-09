CREATE TABLE IF NOT EXISTS graphs_schema.category(
    id SERIAL,
    name VARCHAR(60),
    children []INT,
    parent INT,
    PRIMARY KEY(id)
);
