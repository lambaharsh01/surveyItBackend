package middleware

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"

    "net/http"
    "sync"
    "time"
)

var ( //this setup allows the application to track and rate-limit clients by their IP addresses in a thread-safe way.
    limiters   = make(map[string]*rate.Limiter) //This is a map (dictionary) that stores rate limiters (*rate.Limiter) for each client IP address.
    limitersMu sync.Mutex //The mutex (mutual exclusion lock) is used to ensure that only one goroutine can access or modify the limiters map at a time. This prevents race conditions that could occur if multiple goroutines try to read from or write to the map simultaneously.
)

// getLimiter retrieves or creates a rate limiter for the given IP address
func getLimiter(ip string) *rate.Limiter {
    limitersMu.Lock()
    defer limitersMu.Unlock()

	//limiter: This will hold the *rate.Limiter associated with the IP if it exists.
	//exists: This is a boolean that will be true if a limiter is found for the IP, and false if there is no entry for that IP in the map.
	// If exists is true, it means there’s already a rate limiter for this IP in the limiters map, so return limiter returns the existing limiter.
	// If exists is false, it means there is no rate limiter for this IP, and the code will skip this if block (presumably to create a new limiter for the IP later in the function).
    if limiter, exists := limiters[ip]; exists {
        return limiter
    } 

    // NewLimiter(rate, burst)
    // rate: n requests per minuet 
	// burst: n requests at the same time
	limiter := rate.NewLimiter(10, 5) // per second request
	// limiter := rate.NewLimiter(rate.Every(time.Minute/5), 1) // per minuet request
    limiters[ip] = limiter
	
    //okay so if no limiter exists a new one is  creted and a set time out func is passed to 1 minuet tht will clear the exisitng limiter  
    // After one minute, this goroutine locks the limitersMu mutex to safely delete the limiter for the IP from the limiters map, ensuring thread safety.
    // Once deleted, this cleanup goroutine unlocks limitersMu.
    go func() {
        time.Sleep(time.Minute)
        limitersMu.Lock()
        delete(limiters, ip)
        limitersMu.Unlock()
    }()

    // after clearence a new limiter which is empty/ flest will be returned
    return limiter
}

// RateLimitMiddleware limits the request rate based on client IP address
func RateLimitHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        limiter := getLimiter(ip)

        if !limiter.Allow() {
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
                "success":false,
                "error": "Too many requests",
            })
            return
        }

        c.Next()
    }
}