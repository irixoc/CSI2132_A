--1.
/*SELECT concat(u.first_name, ' ', u.last_name) AS guest_name, prop.property_type AS rental_type, 
r.price_of_stay AS rental_price, r.signing_date AS signing_date, (property_address).country AS branch, pay.payment_method, pay.payment_status
FROM users u, rental_agreement r, properties prop, payments pay
WHERE r.guest_id = u.user_id AND prop.property_id = r.property_id AND r.payment_id = pay.payment_id; 

--2.
CREATE VIEW GuestListView AS SELECT * FROM users WHERE guest = TRUE ORDER BY branch_id ASC, user_id ASC ;


--3.

(SELECT MIN(price_of_stay), signing_date, start_date, end_date, property_id
FROM rental_agreement AS R
       FULL OUTER JOIN
       payments AS PAY
       ON R.payment_id = PAY.payment_id 
	   WHERE PAY.payment_status = 'completed' 
 	OR PAY.payment_status = 'approved' GROUP BY  signing_date, start_date, end_date, property_id ) ; 
	
--4.
SELECT prop.*, rev.rating, rev.review_comments FROM properties prop, rental_agreement r, reviews rev
WHERE r.start_date <= NOW()::DATE AND r.end_date >= NOW()::DATE AND r.property_id = prop.property_id AND rev.review_id = r.review_id
	ORDER BY rev.rating ASC;
	
	
--5
SELECT * FROM properties WHERE property_id NOT IN(SELECT property_id FROM rental_agreement);

--6
SELECT prop.*, r.signing_date FROM properties prop, rental_agreement r WHERE DATE_PART('DAY', r.signing_date) = 10; 

--7
SELECT * FROM employees WHERE salary >= 15000 ORDER BY manager ASC, employee_id ASC;
*/

--8
SELECT prop.property_type, concat(u.first_name, ' ', u.last_name) AS host_name, prop.property_address, r.price_of_stay AS amount_paid,
	pay.payment_method FROM properties prop, users u, rental_agreement r, payments pay WHERE r.guest_id = 3 AND 
	r.property_id = prop.property_id AND u.user_id = r.guest_id AND r.payment_id = pay.payment_id;




