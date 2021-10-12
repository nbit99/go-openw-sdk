package openwsdk

import (
	"encoding/hex"
	"fmt"
	"github.com/nbit99/go-owcrypt"
	"github.com/nbit99/openwallet/v2/hdkeystore"
)

//SignRawTransaction 签名交易单
func SignRawTransaction(rawTx *RawTransaction, key *hdkeystore.HDKey) error {

	for accountID, keySignatures := range rawTx.Signatures {
		//log.Infof("accountID: %s", accountID)
		if keySignatures != nil {
			for _, keySignature := range keySignatures {

				childKey, err := key.DerivedKeyWithPath(keySignature.DerivedPath, keySignature.EccType)
				keyBytes, err := childKey.GetPrivateKeyBytes()
				if err != nil {
					return err
				}
				//log.Debug("privateKey:", hex.EncodeToString(keyBytes))

				//privateKeys = append(privateKeys, keyBytes)
				txHash, err := hex.DecodeString(keySignature.Message)
				//transHash = append(transHash, txHash)

				//log.Infof("sign hash: %s", txHash)

				//签名交易
				/////////交易单哈希签名

				//signature, err := signatureSet.SignTxHash(rawTx.Coin.Symbol, txHash, keyBytes, keySignature.EccType)
				//if err != nil {
				//	return fmt.Errorf("transaction hash sign failed, unexpected error: %v", err)
				//}

				signature, v, sigErr := owcrypt.Signature(keyBytes, nil, txHash, keySignature.EccType)
				if sigErr != owcrypt.SUCCESS {
					return fmt.Errorf("transaction hash sign failed")
				}

				if keySignature.RSV {
					signature = append(signature, v)
				}

				//log.Debug("Signature:", txHash)

				keySignature.Signature = hex.EncodeToString(signature)
			}
		}
		rawTx.Signatures[accountID] = keySignatures
	}

	return nil
}


//SignTxHash 签名交易单Hash
func SignTxHash(signatures map[string][]*KeySignature, key *hdkeystore.HDKey) (map[string][]*KeySignature, error) {

	for accountID, keySignatures := range signatures {
		//log.Infof("accountID: %s", accountID)
		if keySignatures != nil {
			for _, keySignature := range keySignatures {

				childKey, err := key.DerivedKeyWithPath(keySignature.DerivedPath, keySignature.EccType)
				keyBytes, err := childKey.GetPrivateKeyBytes()
				if err != nil {
					return nil, err
				}
				//log.Debug("privateKey:", hex.EncodeToString(keyBytes))

				//privateKeys = append(privateKeys, keyBytes)
				txHash, err := hex.DecodeString(keySignature.Message)
				//transHash = append(transHash, txHash)

				//log.Infof("sign hash: %s", txHash)

				//签名交易
				/////////交易单哈希签名

				//signature, err := signatureSet.SignTxHash(rawTx.Coin.Symbol, txHash, keyBytes, keySignature.EccType)
				//if err != nil {
				//	return fmt.Errorf("transaction hash sign failed, unexpected error: %v", err)
				//}

				signature, v, sigErr := owcrypt.Signature(keyBytes, nil, txHash, keySignature.EccType)
				if sigErr != owcrypt.SUCCESS {
					return nil, fmt.Errorf("transaction hash sign failed")
				}

				if keySignature.RSV {
					signature = append(signature, v)
				}

				//log.Debug("Signature:", txHash)

				keySignature.Signature = hex.EncodeToString(signature)
			}
		}
		signatures[accountID] = keySignatures
	}

	return signatures, nil
}
