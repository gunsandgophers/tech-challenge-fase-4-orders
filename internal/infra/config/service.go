package config

var (
	SERVICE_PRODUCT_URL = GetEnv("SERVICE_PRODUCTS_URL", "tech-challenge-fase-4-products")
	SERVICE_PAYMENT_URL = GetEnv("SERVICE_PAYMENT_URL", "tech-challenge-fase-4-payments")
)
