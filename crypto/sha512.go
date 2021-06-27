package crypto

import (
  goHash "hash"
  cryptoSHA512 "crypto/sha512"
  "encoding/hex"
)

func SHA512(data string) string {
  var hash goHash.Hash
  hash = cryptoSHA512.New()

  hash.Write([]byte(data))
  
  var hashstr string
  hashstr = hex.EncodeToString(hash.Sum(nil))
  
  return hashstr
}
