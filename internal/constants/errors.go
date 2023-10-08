package constants

const (
	PROFESSIONAL_NOT_FOUND_ERROR_CODE = 301
	CUSTOMER_NOT_FOUND_ERROR_CODE     = 302
	SALON_NOT_FOUND_ERROR_CODE        = 303
	SERVICE_NOT_FOUND_ERROR_CODE      = 304
	PRODUCT_NOT_FOUND_ERROR_CODE      = 305
	CONTACT_NOT_FOUND_ERROR_CODE      = 306
	ADDRESS_NOT_FOUND_ERROR_CODE      = 307
	APPOINTMENT_NOT_FOUND_ERROR_CODE  = 308

	PERMISSION_DENIED_ERROR_CODE              = 401
	SALON_CLOSED_ERROR_CODE                   = 402
	APPOINTMENT_TIME_UNAVAILABLE_ERROR_CODE   = 403
	CANNOT_CANCEL_PAST_APPOINTMENT_ERROR_CODE = 404

	CLOUDFLARE_UNAVAILABLE_ERROR_CODE = 901
)

const (
	PROFESSIONAL_NOT_FOUND_ERROR_MESSAGE = "Professional not found"
	CUSTOMER_NOT_FOUND_ERROR_MESSAGE     = "Customer not found"
	SALON_NOT_FOUND_ERROR_MESSAGE        = "Salon not found"
	SERVICE_NOT_FOUND_ERROR_MESSAGE      = "Service not found"
	PRODUCT_NOT_FOUND_ERROR_MESSAGE      = "Product not found"
	CONTACT_NOT_FOUND_ERROR_MESSAGE      = "Contact not found"
	ADDRESS_NOT_FOUND_ERROR_MESSAGE      = "Address not found"
	APPOINTMENT_NOT_FOUND_ERROR_MESSAGE  = "Appointment not found"

	PERMISSION_DENIED_ERROR_MESSAGE              = "Permission denied"
	SALON_CLOSED_ERROR_MESSAGE                   = "Appointment time conflict. Salon is cloed"
	APPOINTMENT_TIME_UNAVAILABLE_ERROR_MESSAGE   = "Appointment time conflict. Time unavailable"
	CANNOT_CANCEL_PAST_APPOINTMENT_ERROR_MESSAGE = "Cannot cancel past appointment"

	CLOUDFLARE_UNAVAILABLE_ERROR_MESSAGE = 901
)
