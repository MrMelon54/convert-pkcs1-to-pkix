package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	p, _ := pem.Decode([]byte(`-----BEGIN RSA PUBLIC KEY-----
MIICCgKCAgEA7jJ3w6FA+oMrCOsotCcCzEg7bKl/LQV6TBBG2yOvwjySXfmO1/zl
hP4JRIhbXbc9WpWV7fVmEBiiW4G/QyxyE/1XR/OhwBWfQVdIhujsglzO7eC0s5+A
C8eE2/gQz+lLMXl7wGDxIykuKlOzdprhWjZ2v02R7gVK9COl7Anc6CGzjwfusgU3
41Ll3znFwdlXVmXZRATHxW43EdCmJFGbqyYGalLjuK2dgrSXTJlGSIa0JopqoLN5
mZTEuT0uJ0Bd1+ycJCW4VJEXUioo+i16KV1ngKU19ycoJoPjyuV+pAlaJZ0J7N6o
2ZTvc6B5Uk/nLUm7ZgWPQPzGVhJNMHjfuFM4k7mKixGdBVu1iH6B81ZwvvHUYced
82JslP5bxL7078Y0/qxatlSMncm5f8+E10J+eTeTazdMnYA0wtzo9oVZ48Nv6rhE
Ky5ml9otZBNgpU4inIiCsXQjoSl3skmE+q3eAhw2VE6a5Ndo/jFqlXSL0JS3VUg2
BKufYrAsAh4vQliFpfNyOGOjQizBDXOSWuyeMpsyXXfcYu1r9n2+xYiwH6deeVDZ
sYKCNvFtEK7vTHs7aTZuME46217xUV8qV8shxXnpE54ej8AiVyn4h+MbVgmMPSif
Iy/o/olimVY4w9dlHtYDRe1QfOqgRBDT5363/7iVAECFqQZAnKI9sFECAwEAAQ==
-----END RSA PUBLIC KEY-----
`))
	if p == nil {
		panic("nil pem block")
	}
	key, err := x509.ParsePKCS1PublicKey(p.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		panic(err)
	}
	mem := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicKey})
	fmt.Printf("%s\n", mem)
}
