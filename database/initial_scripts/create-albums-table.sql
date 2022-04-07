CREATE TABLE authors (
    id SERIAL PRIMARY KEY, 
    title VARCHAR NOT NULL, 
    artist VARCHAR(200) NOT NULL,
    price INT NOT NULL,
    CONSTRAINT Title_Artist_Unique UNIQUE (title, artist)
)