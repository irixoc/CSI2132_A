/*
CREATE OR REPLACE FUNCTION public.get_branch(
	user_address address)
    RETURNS INT
    LANGUAGE 'plpgsql'

    COST 100
    VOLATILE 
    
AS $result$
 DECLARE result int;
BEGIN
 SELECT branch_id into result FROM branch WHERE country = (user_address).country;
return result;
END
$result$;
*/

CREATE OR REPLACE FUNCTION public.user_branch_trigger()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE NOT LEAKPROOF
AS $BODY$ 
BEGIN
  IF NEW.branch_id IS NULL THEN
    NEW.branch_id := get_branch(NEW.user_address);
  END IF;
  RETURN NEW;
END$BODY$;
/*
CREATE TRIGGER user_branch_trigger
    BEFORE INSERT
    ON public.users
    FOR EACH ROW
    EXECUTE PROCEDURE public.user_branch_trigger();
	*/