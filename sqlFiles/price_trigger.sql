DROP FUNCTION price_trigger CASCADE;


CREATE FUNCTION price_trigger()
RETURNS trigger AS '
BEGIN
  IF NEW.price_of_stay IS NULL THEN
    NEW.price_of_stay := calculate_price(NEW.property_id, NEW.start_date, NEW.end_date);
  END IF;
  RETURN NEW;
END' LANGUAGE 'plpgsql';
 
CREATE TRIGGER price_trigger
BEFORE INSERT ON rental_agreement
FOR EACH ROW
EXECUTE PROCEDURE price_trigger();
