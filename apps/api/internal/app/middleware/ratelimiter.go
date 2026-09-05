package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/ulule/limiter/v3"
	stdlibmw "github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimiter membuat middleware rate limit
// format limit string: "requests/period" contoh "5-M" = 5 requests per minute
func RateLimiter(next http.Handler) http.Handler {
	// Simpan data limit di memory
	store := memory.NewStore()

	// Aturan: 10 request per menit per IP
	rate, err := limiter.NewRateFromFormatted("200-M")
	if err != nil {
		log.Fatalf("failed to create rate limiter: %v", err)
	}

	// Buat limiter
	instance := limiter.New(store, rate, limiter.WithTrustForwardHeader(true))

	// Middleware bawaan ulule-limiter
	middleware := stdlibmw.NewMiddleware(instance)

	// Bungkus middleware ke handler berikutnya
	return middleware.Handler(next)
}

// RateLimiterCustom membuat limiter dengan jumlah request & periode custom
func RateLimiterCustom(requests int64, period time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		store := memory.NewStore()
		rate := limiter.Rate{
			Period: period,
			Limit:  requests,
		}
		instance := limiter.New(store, rate, limiter.WithTrustForwardHeader(true))
		middleware := stdlibmw.NewMiddleware(instance)
		return middleware.Handler(next)
	}
}
