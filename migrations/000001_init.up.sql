CREATE TABLE billboards(
    billboard_id SERIAL PRIMARY KEY,
    lat VARCHAR(255) NOT NULL,
    lon VARCHAR(255) NOT NULL,
    azimuth VARCHAR(255) NOT NULL
);

CREATE TABLE users(
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE requests(
    requests_id SERIAL PRIMARY KEY,

    ageFrom INT NOT NULL,
    ageTo INT NOT NULL,
    gender VARCHAR(255) NOT NULL,
    income_a BOOLEAN NOT NULL,
    income_b BOOLEAN NOT NULL,
    income_c BOOLEAN NOT NULL,
    name_billboard VARCHAR(255) NOT NULL,

    value VARCHAR(255),

    user_id INT NOT NULL,
    billboard_id INT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (billboard_id) REFERENCES billboards(billboard_id) ON DELETE CASCADE
);
