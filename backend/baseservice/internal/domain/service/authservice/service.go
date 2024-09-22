package authservice

import (
	"context"
	"errors"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/verificationcode"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	verificationCodeRedis repoiface.VerificationCodeRedisRepository
}

func New(verificationCodeRedis repoiface.VerificationCodeRedisRepository) *AuthService {
	return &AuthService{
		verificationCodeRedis: verificationCodeRedis,
	}
}

func (s *AuthService) CreateVerificationCode(ctx context.Context, bits, expireTime int64) (*verificationcode.VerificationCode, error) {
	codeString := utils.GenerateNumericString(bits)
	codeId := utils.GetSnowflakeId()
	code := verificationcode.New(codeId, codeString)
	if err := s.verificationCodeRedis.Insert(ctx, codeId, expireTime, codeString); err != nil {
		log.Context(ctx).Errorf("failed to create verification code: %v", err)
		return nil, err
	}

	log.Context(ctx).Infow(
		"msg", "create verification code successfully",
		"verification_code_id", codeId,
		"code", codeString,
	)
	return code, nil
}

func (s *AuthService) ValidateVerificationCode(ctx context.Context, code *verificationcode.VerificationCode) (bool, error) {
	codeString, err := s.verificationCodeRedis.Get(ctx, code.VerificationCodeId)
	if err != nil {
		log.Context(ctx).Errorf("failed to query verification code: %v", err)
		return false, errors.New("failed to query verification code")
	}

	anotherCode := verificationcode.New(code.VerificationCodeId, codeString)
	ok, err := code.Check(anotherCode)
	if err != nil {
		return false, err
	}

	_ = s.verificationCodeRedis.Remove(ctx, code.VerificationCodeId)
	return ok, nil
}
