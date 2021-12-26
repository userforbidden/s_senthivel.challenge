package RSASignature

import (
	"testing"
)

func TestOverall(t *testing.T) {
	// keyFromFile := `-----BEGIN RSA PRIVATE KEY-----
	// MIIEpwIBAAKCAQIAvHmBQKiLppiJfzJ8gbFMhwWkGpTwoLiIjy+o9u+IduoDKujE
	// oPSoeN4po+HwXqbftG9lLrVIPYpiLcW06oAP17izo09aB0esKrpHAcmM5PLP+Eed
	// rKQB/g2YER6SctJxbfRwWJBuL2JqiXQ9+RXddeaLz4gR6mrTbKSmyGl3hXWWf0gf
	// KG08TTqIY9SVuV4ZJ4KdNjRDS6K59rEMUlrNUpTTqx/Z4hjjBWsWwpTovHhN3wt6
	// c4+WV/D6w0S0d28N/Tr7JH+xee2QsrIZS1gbCo1ZkuQ/py5dCDficX7qN3deAgty
	// nVJ3Rce+FxdUVFY3s/Rkv8Bldm+P0fM6kp6Mc/UCAwEAAQKCAQEVxeejo5T2evqv
	// YhK2HNhdZM5vb7NYrdjAyd24ITXfnuscwuUsPvdTdz7SOuC83oRLYO8MdU7SJTFw
	// Xj2nNYIfBqxh+mHJI4OFMB3R3AnVzSMvaH6TmiLMyZ8+owyCDUOjUzDDWnQQL0tl
	// c7NRGjCKVUbjg7GmJg8kpuTJlD2FgmvgImj1Brvm8YC98sx3cPq+NGhhBAItiFV7
	// XI57sPzJIoRtkLgHfLD1uZ0zJreHhURE2qG91PmUY/JVmpnb9MB8Xr56ywbUWTfj
	// QO56AVDDhd8fwQM8OxaF8lsWdCv8jt9UTGZmPGtuRj4qTYlUOgHDlirMoxp27pEg
	// /I05QC2ooQKBgQ8Dovwk6kjDDPVvQpu7o1YIuGGPfULnw3Hu3xp7NdvF0C4GSdZn
	// CxKkmCymZYJY7Lo0of6D/4Gv7sfjiz9h1Phi/6GMc2eGguFg1XkfEopWll2ef3qG
	// RVopBH0UuZr8cm8AtiXyhSizc11PaS/0jt4HS5ulqi+jdSmc5110xxfJKQKBgQyN
	// lxnnOGVDuzpza1BgNmCMaEpHLeqHSolVjjAvAk174pYnljjDbHnWCPNZWypCyX73
	// x7OTolwXXyB1/h3LKIZMQDC6uSemy1lRFxkOsTNZoM5jF/JcyFJ3kTn8h404JmDz
	// NVsvdRl2G0lUN9eO+HsyderaqAsVQEFdqrYRp7+R7QKBgQGJlswJH7CZwJ71aLW3
	// zi8GL5NV9Ta/suTc2B3HFinOJNZBsJfy71BWaHrSAz0IzBz5PMlqTOTD1kjDxUOV
	// ShCxWon5buvf+1EG8HU5uFVQLo+hpl4ul7V2083XLHZxeWpigUR7hCU50wtwEU5m
	// 1aZoytZCbZ35OAst1Qdd6PFHmQKBgQUEtduNLGaPZVJAwi4nyJVgjHDlc93GqG2u
	// 3mE06d7AIB2XRd77gWpTvtBKcL+8Y4F23UdVRhs63LTFdwnCJjlKUn8opszLSZAz
	// cuB5ly3ILxlTVEd4AD23vz9HTYmPYQhxMmt51X6QUOhH2us6JOxbh/iT4YpMvJtl
	// LtIKzN0xIQKBgQ25kVzkYf9eMeLMmltyVEC3UYwfhC7kYyi2Zh4WEVgGmbt6onJ3
	// rC9Jw2SLcyORBYPaCjqAAQznh0rSPfr44dEsSSikaHgXDyLTGk/4aQZuASIbne5b
	// Zir1gRJGJELnAoZKEf+m4dzI5/NdrFKau7jgyUF8BFZerUZZp6qQeGcSjQ==
	// -----END RSA PRIVATE KEY-----`
	// block, _ := pem.Decode(([]byte(keyFromFile)))
	// privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	// if err != nil{
	// 	t.Fail()
	// }

	message := "theAnswerIs42"
	actual := SignInput(message)
	expected := SignedMessageStructure{
		Message:   message,
		Signature: "TCCa+edvayZo9HHyQiURINrIc+urAUsPMmUUq+BRsJYWD+JhuB7Cp4Mbvb8qih5WibhjJWWZl7+xhvVfECJssMRDykippupk04f6jMaWkjl9yhLNZdAQtf0h1XB8Vq/K9bpDpP1zQRFv2lI/n8KDK6ce/SN7UD79n+V39qqWluF41si9ov8nMNYPyZ+2bKa/LY8oB/4K4RopuJmFEnrgHwBiqpKdfYs/+grYnRFpQF4z/fqOZhf+Kiu6A8rDkaFuzI03RiHOEvmXAFc6O1pMwNhkT5devtZuWqoAuyKnMTkeTjZxFdk0z8m9A6uhnAWPuBamFLhQ4QLnE67DiTUlCQI=",
		Pubkey:    "-----BEGIN PUBLIC KEY-----\nMIIBCwKCAQIAvHmBQKiLppiJfzJ8gbFMhwWkGpTwoLiIjy+o9u+IduoDKujEoPSo\neN4po+HwXqbftG9lLrVIPYpiLcW06oAP17izo09aB0esKrpHAcmM5PLP+EedrKQB\n/g2YER6SctJxbfRwWJBuL2JqiXQ9+RXddeaLz4gR6mrTbKSmyGl3hXWWf0gfKG08\nTTqIY9SVuV4ZJ4KdNjRDS6K59rEMUlrNUpTTqx/Z4hjjBWsWwpTovHhN3wt6c4+W\nV/D6w0S0d28N/Tr7JH+xee2QsrIZS1gbCo1ZkuQ/py5dCDficX7qN3deAgtynVJ3\nRce+FxdUVFY3s/Rkv8Bldm+P0fM6kp6Mc/UCAwEAAQ==\n-----END PUBLIC KEY-----\n",
	}
	if expected != actual {
		t.Fail()
	}
}
