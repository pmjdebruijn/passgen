
package main

import "fmt"
import "time"
import "strings"
import "crypto/rand"
import "crypto/hmac"
import "crypto/sha512"
import "encoding/base64"
import "encoding/binary"

func main () {

  // request a generous amount of raw entropy from operating system
  raw_entropy := make([]byte, 4096)

  n, err := rand.Read(raw_entropy[:])

  if n != len(raw_entropy) || err != nil {
    panic(err)
  }

  salt := make([]byte, 8)

  binary.BigEndian.PutUint64(salt, uint64((time.Now().UnixNano())))

  h := hmac.New(sha512.New, salt)
  h.Write(raw_entropy)

  hashed_entropy := h.Sum(nil)

  base64_entropy := base64.StdEncoding.EncodeToString(hashed_entropy)

  // avoid easy to confuse chars, ref DIN 1450, remove 10 chars from 64 total
  base64_entropy = strings.Replace(base64_entropy, "0", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "O", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "o", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "1", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "I", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "l", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "5", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "S", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "8", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "B", "", -1)

  // no special chars, remove 3 chars from 64 total
  base64_entropy = strings.Replace(base64_entropy, "+", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "/", "", -1)
  base64_entropy = strings.Replace(base64_entropy, "=", "", -1)

  // 128 / log2(64 - 10 - 3) = 22.5
  fmt.Println(base64_entropy[:24])

}

