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
	value    interface{}
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

// Generates a temporary token and stores it in the cache with the specified key
func (vc *VerificationCodeService) GenerateTemporaryToken(key string, identifier string) string {
	log.Println(`[VerificationCodeService.GenerateTemporaryToken] - Generating temporary token for user:`, key)
	token, _ := utils.GenerateRandomString(40)
	enc := utils.Base64Encode(identifier)
	value := token + "-" + enc
	vc.cache[key+":"+"token"] = verificationEntry{
		value:    value,
		expireAt: time.Now().Add(vc.defaultTTL),
	}
	return value
}

func (vc *VerificationCodeService) GetCacheRegistry(key string) (interface{}, bool) {
	log.Println(`[VerificationCodeService.GetCacheRegistry] - Getting cache registry for key:`, key)
	entry, exists := vc.cache[key]
	if exists && time.Now().Before(entry.expireAt) {
		return entry.value, true
	}
	return 0, false
}

// Periodically cleans up expired entries in the cache
func (vc *VerificationCodeService) cleanupExpiredEntries() {
	for {
		time.Sleep(vc.cleanupTime)
		for key, entry := range vc.cache {
			if time.Now().After(entry.expireAt) {
				log.Println(`[VerificationCodeService.cleanupExpiredEntries] - Cleaning value for key:`, key)
				delete(vc.cache, key)
			}
		}
	}
}
