CREATE TABLE IF NOT EXISTS appointment_product (
	id BIGSERIAL NOT NULL PRIMARY KEY,  
	appointment_id BIGSERIAL REFERENCES appointments(id), 
	product_id BIGSERIAL REFERENCES products(id)
);