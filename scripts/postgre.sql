CREATE SCHEMA param_tracker;

CREATE TABLE IF NOT EXISTS param_tracker.point (
    point_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    alias VARCHAR NOT NULL
);


CREATE TABLE IF NOT EXISTS param_tracker.gps (
    gps_id SERIAL PRIMARY KEY,
    point_id INT NOT NULL,
    point_gps_id INT NOT NULL,
    lat FLOAT NOT NULL,
    lon FLOAT NOT NULL,
    speed FLOAT NOT NULL,
    time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS param_tracker.param_str (
    string_id SERIAL PRIMARY KEY,
    string_point_id INT NOT NULL,
    point_id INT NOT NULL,
    channel INT NOT NULL,
    value TEXT NOT NULL,
    time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS param_tracker.param_int (
    int_id SERIAL PRIMARY KEY,
    int_point_id INT NOT NULL,
    point_id INT NOT NULL,
    channel INT NOT NULL,
    value INT NOT NULL,
    time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);