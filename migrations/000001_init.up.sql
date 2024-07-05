CREATE TABLE billboards(
    billboard_id SERIAL PRIMARY KEY,
    lat VARCHAR(255) NOT NULL,
    lon VARCHAR(255) NOT NULL,
    azimuth VARCHAR(255) NOT NULL
);
