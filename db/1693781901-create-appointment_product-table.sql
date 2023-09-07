CREATE TABLE IF NOT EXISTS appointment_products (
	id BIGSERIAL NOT NULL PRIMARY KEY,  
	appointment_id UUID REFERENCES appointments(id), 
	product_id UUID REFERENCES products(id)
);