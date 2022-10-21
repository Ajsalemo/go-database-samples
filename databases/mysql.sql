CREATE DATABASE todos;

CREATE TABLE todos (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL
);

INSERT INTO todos (
    name,
    completed
) VALUES 
('sweep the floor', false),
('clean the car', false),
('mow the grass', true),
('vacuum the carpet', true);