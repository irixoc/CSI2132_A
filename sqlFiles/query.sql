CREATE TYPE valid_property_type AS ENUM('apartment', 'bed and breakfast', 'unique home', 'vacation home');

CREATE TABLE client(
    client_id SERIAL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    middle_name VARCHAR(30),
    last_name VARCHAR(30) NOT NULL,
    email VARCHAR(30) UNIQUE NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    address_id VARCHAR(20) NOT NULL,
    host BOOLEAN,
    guest BOOLEAN
);

CREATE TABLE employee(
    employee_id SERIAL PRIMARY KEY,
    first_name VARCHAR(20) NOT NULL,
    middle_name VARCHAR(20),
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(20) UNIQUE NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    address VARCHAR(20) NOT NULL,
    salary INTEGER,
    position VARCHAR(20)
);

CREATE TABLE pricing(
	price_id int NOT NULL,
	rate_per_day DEC(12,2),
	rate_per_week DEC(12,2),
	PRIMARY KEY (price_id));
	
	
CREATE TABLE property(
	property_id int NOT NULL,
	rental_id int NOT NULL,
	address_id int NOT NULL,
	pricing_id int NOT NULL,
	property_type VALID_PROPERTY_TYPE NOT NULL,
	branch_id int NOT NULL,
	accommodates int,
	amenities varchar(800),
	bathrooms int,
	bedrooms int,
	PRIMARY KEY (property_id),
	FOREIGN KEY (address_id) REFERENCES addresses(address_id),
	FOREIGN KEY (pricing_id) REFERENCES pricing(price_id)
);	

CREATE TABLE rental_agreement(
	rental_id INTEGER NOT NULL,
	property_id INTEGER NOT NULL,
    employee_id INTEGER,
	guest_id INTEGER NOT NULL,
	signing_date TIMESTAMP DEFAULT NOW(),
	start_date DATE NOT NULL,
	end_date DATE NOT NULL,
	PRIMARY KEY (rental_id),
	FOREIGN KEY (property_id) REFERENCES property(property_id),
	FOREIGN KEY (guest_id) REFERENCES client(client_id),
    FOREIGN KEY(employee_id) REFERENCES employee(employee_id) 
); 



CREATE TABLE payment(
	payment_id int NOT NULL,
	guest_id int  NOT NULL,
	host_id int NOT NULL,
	rental_id int NOT NULL,
	payment_method varchar(20),
	payment_status boolean,
	PRIMARY KEY (payment_id),
	FOREIGN KEY (guest_id) REFERENCES guests(guest_id),
	FOREIGN KEY (host_id) REFERENCES hosts(host_id),
	FOREIGN KEY (rental_id) REFERENCES rental_agreement(rental_id));