CREATE DATABASE todos;

CREATE TABLE todos
(
	Id INT NOT NULL IDENTITY(1,1) PRIMARY KEY, 
    name VARCHAR(255) NOT NULL, 
    completed BIT NOT NULL 
);

INSERT INTO todos (
    name,
    completed
) VALUES 
('sweep the floor', 0),
('clean the car', 0),
('mow the grass', 1),
('vacuum the carpet', 1);