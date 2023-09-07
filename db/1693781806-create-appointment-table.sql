CREATE TABLE IF NOT EXISTS appointment (
	id UUID NOT NULL PRIMARY KEY,  
	barber_id UUID REFERENCES barber(id), 
	customer_id UUID REFERENCES customer(id), 
	duration_in_min INTEGER NOT NULL, 
	starts_at TIMESTAMP NOT NULL, 
	ends_at TIMESTAMP NOT NULL, 
	created_at TIMESTAMP NOT NULL
);