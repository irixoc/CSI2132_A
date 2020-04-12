ALTER TABLE users DROP COLUMN user_name;
ALTER TABLE users ALTER COLUMN phone_number TYPE CHAR(10);

INSERT INTO users (user_address, first_name, last_name, phone_number)
	VALUES
	( ROW(213, 'Kingston', 'Place St.', 'Ontario', 'K1S4D4', 'Canada')::address, 'John', 'Bryant',6134325938);