package cbnt

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"

	"github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest"
	"github.com/tjfoc/gmsm/sm2"
)

func validateSignature(s *manifest.Signature, unsigneddata []byte, key crypto.PublicKey) error {
	switch s.SigScheme {
	case manifest.AlgRSAPSS:
		data, err := s.SignatureData()
		if err != nil {
			return err
		}
		if err := data.Verify(key.(*rsa.PublicKey), unsigneddata); err != nil {
			return err
		}

	case manifest.AlgRSASSA:
		data, err := s.SignatureData()
		if err != nil {
			return err
		}
		if err := data.Verify(key.(*rsa.PublicKey), unsigneddata); err != nil {
			return err
		}
	case manifest.AlgECDSA:
		ecdsaSig, err := s.SignatureData()
		if err != nil {
			return err
		}
		if err := ecdsaSig.Verify(key.(*ecdsa.PublicKey), unsigneddata); err != nil {
			return fmt.Errorf("ECDSA signature is not valid")
		}
	case manifest.AlgSM2:
		sm2Sig, err := s.SignatureData()
		if err != nil {
			return err
		}
		if sm2Sig.Verify(key.(*sm2.PublicKey), unsigneddata); err != nil {
			return fmt.Errorf("SM2 signature is not valid")
		}
	default:
		return fmt.Errorf("Signature has invalid signature scheme: %v", s.SigScheme)
	}
	return nil
}
