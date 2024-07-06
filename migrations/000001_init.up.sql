CREATE TABLE billboards(
    billboard_id SERIAL PRIMARY KEY,
    lat VARCHAR(255) NOT NULL,
    lon VARCHAR(255) NOT NULL,
    azimuth VARCHAR(255) NOT NULL
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
