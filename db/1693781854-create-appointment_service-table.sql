CREATE TABLE IF NOT EXISTS appointment_service (
	id BIGSERIAL NOT NULL PRIMARY KEY,  
	appointment_id UUID REFERENCES appointment(id), 
	service_id UUID REFERENCES service(id)
);