DELIMITER //

CREATE PROCEDURE InsertRandomUsers()
BEGIN
    DECLARE counter INT DEFAULT 0;
    DECLARE random_name VARCHAR(255);
    DECLARE characters CHAR(36) DEFAULT 'abcdefghijklmnopqrstuvwxyz0123456789';

    WHILE counter < 200 DO
        SET random_name = '';
        WHILE CHAR_LENGTH(random_name) < 10 DO
            SET random_name = CONCAT(random_name, SUBSTRING(characters, FLOOR(1 + RAND() * 36), 1));
        END WHILE;

        INSERT INTO users (name) VALUES (random_name);
        SET counter = counter + 1;
    END WHILE;
END //

DELIMITER ;

CALL InsertRandomUsers();
---
INSERT INTO theatre (name) VALUES ('PVR');
INSERT INTO theatre (name) VALUES ('INOX');
--- 

DELIMITER //

CREATE PROCEDURE InsertSeats()
BEGIN
    DECLARE row_num INT DEFAULT 1;
    DECLARE col_num CHAR(1);
    DECLARE seat_name VARCHAR(4);
    DECLARE max_rows INT DEFAULT 20;  -- Number of rows
    DECLARE max_cols CHAR(1) DEFAULT 'J';  -- Number of columns
    DECLARE current_col CHAR(1);

    -- Loop through rows
    WHILE row_num <= max_rows DO
        SET current_col = 'A';
        -- Loop through columns
        WHILE current_col <= max_cols DO
            SET seat_name = CONCAT(row_num, current_col);
            INSERT INTO seats (name) VALUES (seat_name);
            SET current_col = CHAR(ASCII(current_col) + 1);
        END WHILE;
        SET row_num = row_num + 1;
    END WHILE;
END //

DELIMITER ;

CALL InsertSeats();