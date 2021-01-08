CREATE TABLE IF NOT EXISTS assignments (
    id SERIAL,

    investment INTEGER NOT NULL,
    success BOOLEAN NOT NULL,

    PRIMARY KEY (id)
);
