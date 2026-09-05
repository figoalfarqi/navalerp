package middleware

import (
	"log"
	"net/http"
	"time"
)

// statusRecorder digunakan untuk menangkap status code response
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *statusRecorder) Flush() {
	if flusher, ok := rec.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

// Logger mencatat request dan response
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if strings.HasPrefix(r.URL.Path, "/sse/") {
		// 	next.ServeHTTP(w, r)
		// 	return
		// }

		start := time.Now()

		// Bungkus ResponseWriter supaya bisa ambil status code
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		// Lanjut ke handler berikutnya
		next.ServeHTTP(rec, r)

		duration := time.Since(start)
		// Logging format
		log.Printf("[%s] %s %d %s from %s",
			r.Method,
			r.URL.Path,
			rec.status,
			time.Since(start),
			getIP(r),
		)

		// FILTER ERROR & SLOW REQUEST
		if rec.status >= 400 {
			_ = WriteErrorLog(ErrorLog{
				Timestamp: time.Now().Format(time.RFC3339),
				Method:    r.Method,
				Path:      r.URL.Path,
				Status:    rec.status,
				Duration:  duration / time.Millisecond,
				IP:        getIP(r),
			})
		} else if duration > time.Second {
			_ = WriteSlowLog(ErrorLog{
				Timestamp: time.Now().Format(time.RFC3339),
				Method:    r.Method,
				Path:      r.URL.Path,
				Status:    rec.status,
				Duration:  duration / time.Millisecond,
				IP:        getIP(r),
			})
		}
	})
}

// getIP mengambil IP address asli user
func getIP(r *http.Request) string {
	// Jika lewat proxy
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	return r.RemoteAddr
}
