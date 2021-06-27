package crypto

import (
  "testing"
)

func TestSHA512er(t *testing.T){
  var data string
  data = "sha512"

  var actual_hashstr string
  actual_hashstr = "1f9720f871674c18e5fecff61d92c1355cd4bfac25699fb7ddfe7717c9669b4d085193982402156122dfaa706885fd64741704649795c65b2a5bdec40347e28a"

  if hashstr := SHA512(data); actual_hashstr != hashstr {
    t.Error("SHA512 does not work properly")
  }
}
