CREATE TABLE sectors(
    sector_id SERIAL PRIMARY KEY,
    x_max DECIMAL(8, 2),
    x_min DECIMAL(8, 2),
    y_max DECIMAL(8, 2),
    y_min DECIMAL(8, 2)
);

CREATE TABLE billboards(
    billboard_id SERIAL PRIMARY KEY,
    sector_id INT NOT NULL,
    lat DECIMAL(8, 2) NOT NULL,
    lon DECIMAL(8, 2) NOT NULL,
    azimuth INT NOT NULL,

    FOREIGN KEY(sector_id) REFERENCES sectors(sector_id) ON DELETE CASCADE
);

CREATE TABLE requests(
    request_id SERIAL PRIMARY KEY,

    age_from INT NOT NULL,
    age_to INT NOT NULL,
    gender VARCHAR(255) NOT NULL,
    income_a BOOLEAN NOT NULL,
    income_b BOOLEAN NOT NULL,
    income_c BOOLEAN NOT NULL,
    name_billboard VARCHAR(255) NOT NULL
);

CREATE TABLE requests_billboards(
    requests_billboards_id SERIAL PRIMARY KEY,

    request_id INT NOT NULL,
    billboard_id INT NOT NULL,
    value VARCHAR(255),

    FOREIGN KEY(request_id) REFERENCES requests(request_id) ON DELETE CASCADE,
    FOREIGN KEY(billboard_id) REFERENCES billboards(billboard_id) ON DELETE CASCADE
);
