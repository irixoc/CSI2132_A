CREATE TRIGGER guest_in_rentals ON dbo.payments
AFTER INSERT
AS
IF EXISTS (SELECT rental_id
           FROM inserted i
           INNER JOIN rental_agreement r 
              ON r.rental_id = i.rental_id
           WHERE i.guest_id = f.guest_id
      )
BEGIN
    RAISERROR ('guest does not exist in rental agreement', 16, 1);
    ROLLBACK TRANSACTION;

    RETURN 
END;

GO