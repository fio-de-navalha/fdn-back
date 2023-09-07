CREATE TABLE IF NOT EXISTS appointment_product (
	id BIGSERIAL NOT NULL PRIMARY KEY,  
	appointment_id UUID REFERENCES appointment(id), 
	product_id UUID REFERENCES product(id)
);