CREATE DATABASE dbz;
USE dbz;
---
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
--- 
CREATE TABLE theatre (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
);
---
CREATE TABLE seats (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    theatre_id INT,
    user_id INT,
    FOREIGN KEY (theatre_id) REFERENCES theatre(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);