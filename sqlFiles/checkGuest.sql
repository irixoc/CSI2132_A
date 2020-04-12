DROP FUNCTION check_guest CASCADE;

CREATE FUNCTION check_guest(guest_id int)
 RETURNS BOOLEAN
 
 AS
 $result$
 DECLARE result BOOLEAN;
BEGIN
 SELECT guest into result FROM users WHERE user_id = guest_id;
return result;
END
$result$ LANGUAGE 'plpgsql';


ALTER TABLE rental_agreement ADD CONSTRAINT chk_guest CHECK( check_guest(guest_id) = TRUE);