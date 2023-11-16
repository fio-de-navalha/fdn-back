package app

import (
	"log"
	"sync"
	"time"

	"github.com/fio-de-navalha/fdn-back/pkg/utils"
)

type VerificationCodeService struct {
	mu          sync.Mutex
	cache       map[string]verificationEntry
	defaultTTL  time.Duration
	cleanupTime time.Duration
}

type verificationEntry struct {
	code     int
	expireAt time.Time
}

func NewVerificationCodeService(defaultTTL, cleanupTime time.Duration) *VerificationCodeService {
	cache := make(map[string]verificationEntry)
	vc := &VerificationCodeService{
		cache:       cache,
		defaultTTL:  defaultTTL,
		cleanupTime: cleanupTime,
	}
	go vc.cleanupExpiredEntries()
	return vc
}

// Generates a random 6-digit verification code and stores it in the cache with the specified key
func (vc *VerificationCodeService) GenerateVerificationCode(key string) int {
	log.Println(`[VerificationCodeService.GenerateVerificationCode] - Generating verification code for user:`, key)
	code := utils.GenerateSixDigitCode()
	vc.mu.Lock()
	defer vc.mu.Unlock()
	vc.cache[key] = verificationEntry{
		code:     code,
		expireAt: time.Now().Add(vc.defaultTTL),
	}
	return code
}

// Retrieves the verification code from the cache for the specified key
func (vc *VerificationCodeService) GetVerificationCode(key string) (int, bool) {
	log.Println(`[VerificationCodeService.GetVerificationCode] - Getting verification code for user:`, key)
	vc.mu.Lock()
	defer vc.mu.Unlock()
	entry, exists := vc.cache[key]
	if exists && time.Now().Before(entry.expireAt) {
		return entry.code, true
	}
	return 0, false
}

// Periodically cleans up expired entries in the cache
func (vc *VerificationCodeService) cleanupExpiredEntries() {
	for {
		time.Sleep(vc.cleanupTime)
		vc.mu.Lock()
		for key, entry := range vc.cache {
			if time.Now().After(entry.expireAt) {
				log.Println(`[VerificationCodeService.cleanupExpiredEntries] - Cleaning verification code for user:`, key)
				delete(vc.cache, key)
			}
		}
		vc.mu.Unlock()
	}
}
