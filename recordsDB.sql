
CREATE TABLE albums ( albumID int NOT NULL, title varchar(64) NOT NULL, artist varchar(64) NOT NULL, price double(4,2), PRIMARY KEY (albumID) );

Insert into albums (albumID, title, artist, price) VALUES (1, 'Blue Train', 'John Coltrane', 56.99);
Insert into albums (albumID, title, artist, price) VALUES (2, 'Jeru', 'Gerry Mulligan', 17.99);
Insert into albums (albumID, title, artist, price) VALUES (3, 'Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 33.99);