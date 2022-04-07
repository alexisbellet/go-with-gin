CREATE TABLE albums (
    id SERIAL PRIMARY KEY, 
    title VARCHAR NOT NULL, 
    artist VARCHAR(200) NOT NULL,
    price FLOAT NOT NULL,
    CONSTRAINT Title_Artist_Unique UNIQUE (title, artist)
);

INSERT INTO albums (title, artist, price)
VALUES 
    ('Blue Train','John Coltrane', 56.99),
    ('Jeru','Gerry Mulligan', 17.99),
    ('Sarah Vaughan and Clifford Brown','Sarah Vaughan', 39.99)