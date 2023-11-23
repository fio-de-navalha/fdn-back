package app

import (
	"log"
	"time"

	"github.com/fio-de-navalha/fdn-back/pkg/utils"
)

type VerificationCodeService struct {
	cache       map[string]verificationEntry
	defaultTTL  time.Duration
	cleanupTime time.Duration
}

type verificationEntry struct {
	code     interface{}
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
	vc.cache[key+":"+"code"] = verificationEntry{
		code:     code,
		expireAt: time.Now().Add(vc.defaultTTL),
	}
	return code
}

// Generates a temporary token and stores it in the cache with the specified key
func (vc *VerificationCodeService) GenerateTemporaryToken(key string) string {
	log.Println(`[VerificationCodeService.GenerateTemporaryToken] - Generating temporary token for user:`, key)
	code, _ := utils.GenerateRandomString(40)
	vc.cache[key+":"+"token"] = verificationEntry{
		code:     code,
		expireAt: time.Now().Add(vc.defaultTTL),
	}
	return code
}

func (vc *VerificationCodeService) GetCacheRegistry(key string) (interface{}, bool) {
	log.Println(`[VerificationCodeService.GetCacheRegistry] - Getting cache registry for key:`, key)
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
		for key, entry := range vc.cache {
			if time.Now().After(entry.expireAt) {
				log.Println(`[VerificationCodeService.cleanupExpiredEntries] - Cleaning verification code for user:`, key)
				delete(vc.cache, key)
			}
		}
	}
}
