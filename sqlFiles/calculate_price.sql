DROP FUNCTION calculate_price CASCADE;

CREATE FUNCTION calculate_price(property_id_x int, start_date DATE, end_date DATE) 
RETURNS NUMERIC(12,2)
AS $price$
DECLARE price NUMERIC(12,2);
DECLARE id_of_price int; DECLARE day_price NUMERIC(12,2); DECLARE week_price NUMERIC(12,2);
DECLARE rental_time int; DECLARE weeks int; DECLARE days int;
BEGIN 
SELECT pricing_id into id_of_price FROM properties WHERE property_id = property_id_x;
SELECT rate_per_day into day_price FROM pricing WHERE price_id = id_of_price;
SELECT rate_per_week into week_price FROM pricing WHERE price_id = id_of_price;
SELECT (end_date - start_date) INTO rental_time;
SELECT FLOOR(rental_time / 7) INTO weeks;
SELECT rental_time % 7 INTO days;
price = (weeks*week_price) + (days*day_price);
return price;
END
$price$ LANGUAGE 'plpgsql';

--ALTER TABLE rental_agreement ADD COLUMN price NUMERIC(12,2) DEFAULT calculate_price(property_id, start_date, end_date);

--SELECT concat(g.first_name, ' ', g.last_name) AS 'Guest Name', p.propert_type AS 'Rental Type', r.signing_date AS 