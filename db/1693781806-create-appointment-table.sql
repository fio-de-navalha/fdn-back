CREATE TABLE IF NOT EXISTS appointments (
	id BIGSERIAL NOT NULL PRIMARY KEY,  
	barber_id UUID REFERENCES barbers(id), 
	customer_id UUID REFERENCES customers(id), 
	duration_in_min INTEGER NOT NULL, 
	starts_at TIMESTAMP NOT NULL, 
	ends_at TIMESTAMP NOT NULL, 
	created_at TIMESTAMP NOT NULL
);