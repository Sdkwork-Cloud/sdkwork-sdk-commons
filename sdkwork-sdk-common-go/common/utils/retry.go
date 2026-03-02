package common

import (
	"math"
	"time"
)

func WithRetry(fn func() error, config *RetryConfig) error {
	if config == nil {
		config = &DefaultRetryConfig
	}
	
	var lastErr error
	for attempt := 0; attempt <= config.MaxRetries; attempt++ {
		err := fn()
		if err == nil {
			return nil
		}
		lastErr = err
		if attempt >= config.MaxRetries {
			break
		}
		if config.RetryCondition != nil && !config.RetryCondition(err, attempt) {
			break
		}
		delay := CalculateDelay(attempt, *config)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	return lastErr
}

func CalculateDelay(retryCount int, config RetryConfig) float64 {
	switch config.RetryBackoff {
	case BackoffFixed:
		return config.RetryDelay
	case BackoffLinear:
		return config.RetryDelay * float64(retryCount+1)
	case BackoffExponential:
		delay := config.RetryDelay * math.Pow(2, float64(retryCount))
		return math.Min(delay, config.MaxRetryDelay)
	default:
		return config.RetryDelay
	}
}

func CreateRetryConfig(maxRetries int, retryDelay float64, retryBackoff RetryBackoff) *RetryConfig {
	return &RetryConfig{
		MaxRetries:    maxRetries,
		RetryDelay:    retryDelay,
		RetryBackoff:  retryBackoff,
		MaxRetryDelay: 30.0,
	}
}

func Sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
