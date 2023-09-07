CREATE TABLE IF NOT EXISTS appointment_services (
	id BIGSERIAL NOT NULL PRIMARY KEY,  
	appointment_id UUID REFERENCES appointments(id), 
	service_id UUID REFERENCES services(id)
);