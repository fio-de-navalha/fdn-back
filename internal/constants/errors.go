package constants

const (
	INVALID_CREDENTIAL_ERROR_CODE          = 101
	PERMISSION_DENIED_ERROR_CODE           = 102
	PROFESSIONAL_ALREADY_EXISTS_ERROR_CODE = 103
	CUSTOMER_ALREADY_EXISTS_ERROR_CODE     = 104

	PROFESSIONAL_NOT_FOUND_ERROR_CODE           = 201
	CUSTOMER_NOT_FOUND_ERROR_CODE               = 202
	SALON_NOT_FOUND_ERROR_CODE                  = 203
	SERVICE_NOT_FOUND_ERROR_CODE                = 204
	PRODUCT_NOT_FOUND_ERROR_CODE                = 205
	CONTACT_NOT_FOUND_ERROR_CODE                = 206
	ADDRESS_NOT_FOUND_ERROR_CODE                = 207
	PERIOD_NOT_FOUND_ERROR_CODE                 = 208
	APPOINTMENT_NOT_FOUND_ERROR_CODE            = 209
	SECURITY_QUESTION_NOT_FOUND_ERROR_CODE      = 210
	SECURITY_QUESTION_ANSWER_INVALID_ERROR_CODE = 211

	SALON_CLOSED_ERROR_CODE                         = 301
	APPOINTMENT_TIME_UNAVAILABLE_ERROR_CODE         = 302
	CANNOT_CANCEL_PAST_APPOINTMENT_ERROR_CODE       = 303
	PROFESSIONAL_ALREADY_EXISTS_IN_SALON_ERROR_CODE = 304

	CLOUDFLARE_UNAVAILABLE_ERROR_CODE = 901
)

const (
	INVALID_CREDENTIAL_ERROR_MESSAGE          = "Invalid credential"
	PERMISSION_DENIED_ERROR_MESSAGE           = "Permission denied"
	PROFESSIONAL_ALREADY_EXISTS_ERROR_MESSAGE = "professional alredy exists"
	CUSTOMER_ALREADY_EXISTS_ERROR_MESSAGE     = "customer alredy exists"

	PROFESSIONAL_NOT_FOUND_ERROR_MESSAGE           = "Professional not found"
	CUSTOMER_NOT_FOUND_ERROR_MESSAGE               = "Customer not found"
	SALON_NOT_FOUND_ERROR_MESSAGE                  = "Salon not found"
	SERVICE_NOT_FOUND_ERROR_MESSAGE                = "Service not found"
	PRODUCT_NOT_FOUND_ERROR_MESSAGE                = "Product not found"
	CONTACT_NOT_FOUND_ERROR_MESSAGE                = "Contact not found"
	ADDRESS_NOT_FOUND_ERROR_MESSAGE                = "Address not found"
	PERIOD_NOT_FOUND_ERROR_MESSAGE                 = "Period not found"
	APPOINTMENT_NOT_FOUND_ERROR_MESSAGE            = "Appointment not found"
	SECURITY_QUESTION_NOT_FOUND_ERROR_MESSAGE      = "Security question not found"
	SECURITY_QUESTION_ANSWER_INVALID_ERROR_MESSAGE = "security question or answer does not match"

	SALON_CLOSED_ERROR_MESSAGE                         = "Appointment time conflict. Salon is cloed"
	APPOINTMENT_TIME_UNAVAILABLE_ERROR_MESSAGE         = "Appointment time conflict. Time unavailable"
	CANNOT_CANCEL_PAST_APPOINTMENT_ERROR_MESSAGE       = "Cannot cancel past appointment"
	PROFESSIONAL_ALREADY_EXISTS_IN_SALON_ERROR_MESSAGE = "professional alredy exists in salon"

	CLOUDFLARE_UNAVAILABLE_ERROR_MESSAGE = "Cloudflare service unavailbale"
)
