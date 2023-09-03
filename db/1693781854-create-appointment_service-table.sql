CREATE TABLE IF NOT EXISTS appointment_service (
	id BIGSERIAL NOT NULL PRIMARY KEY,  
	appointment_id BIGSERIAL REFERENCES appointments(id), 
	service_id BIGSERIAL REFERENCES services(id)
);