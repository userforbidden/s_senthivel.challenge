package RSASignature

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

func TestOverall(t *testing.T) {
	keyFromSystem := `-----BEGIN RSA PRIVATE KEY-----
	MIIEpwIBAAKCAQIAsmg2jqguDen/Hp4/J1c/fgo0iMM3MzfFXKB9krcOSdOM3i2W
	vaYFXXexECihP8yQrGkUCVIJBTrqaG5rVpjDzMFsCSC1QZzNOPTN00CshPlA/pSJ
	TKHxfCpHDiETu3bwnY4CuYUPO22kib0PI8PQztB0Q564PlpQTeFT5fhNw3ZWdrXP
	9siqIb2aDbfO2YHFqtkkzEJvIuR1zxdXNoGGghuEOoVwSlXDPew2U5csMMSq4boK
	rBmREXis+NO8N7pPiTLhUfcHniRPY6Rti5K9eWM5hSwbS5fVlF/5TZl/i+AOgNIF
	c5P+e9X37dCTco9O2OmIpL5nwRYyop80iGeaSm8CAwEAAQKCAQFhek9U4s9HCvM5
	53zaIfGvDCRoFUGz2mUblxUFdL951vzA6GfUd0G3MA6upYR1N3BUw5UPAqOrF+W7
	6Zj8V7ujxIsi8ZwXnXUECb3sDwo0HXxFnSLnnBQ1o2PvdGMzqmkNrOoQM4N/FQkx
	uPrsr7EHTx43VCGW8AGwi80reeADTDaScgX2ibBXPcbFT2CghjnsoSHzC4bQeRYX
	vbbOGYMO59WHbj/JvjRONu6QKP5pBjZJuKfUHIjcmPgjUBHi+1pfq3ZnncdSute/
	AwxGJXIKAYZkdScQvkPpG0mLIKcuyiTkf5e68c2yHCL0K1VJDHY85F6zlbgTzWmX
	aI7vyLR7AQKBgQySZRKJY9HJ9Fuafgxb7vtx604cVw/+JZWPW1/ZS4jGAOb9H6cU
	JxsWCXZJf6gIEqcEa8Bt3XLuWyfPTEzEQ/6KMT6FOs9lqqSotGOzRqo1PI9QwQ15
	orn3HdCibjY5yXHNqkJQnFycJ4fVqDFqto+lKle6tnEhnEC25VjWvV+adQKBgQ4w
	5M/lnjf26b6K3CXk7lcDbk5r4G8vukDmmVEwQRQYOQJ6dCWxvEhwU8nTeY6Nx/uD
	v2iUOSpPtqbRntkCWjIloLS8k56YFog6obdrgIVQmAU293Cb7aOu4UDffF7xpl8N
	8nwGwn6dRCKmz8d2GqOuo0/bPlQfx79TsG7Z/T6M0wKBgQdA+NyHNXc4pautKomL
	Cgs/IdwB8iIiLCFtg/V79V9HkAG3j2nh8OZqAq+jnOqhLdH8dQek3J+R9Dq2G3o3
	zFuz9NH0IYjFSn6TcPIN4PoYmYa/u4RA6KmQcj2NWzYihZ8KYfC9flrZJt3zY4aP
	4iJj1FimPNpt4DBe11o5HS9qeQKBgQTrpyr+7zFwh/Ptqg/ppUq+gcznUta6sdY3
	HwortSUGkR3u378CNcrYR/U/QnLPbnwjXocgfgHRyf9lx7JvQl8I9QQ2LulaytVJ
	nzE+MT4Ih+2m9X0yU9/hj3EzDKjPGWT0LRQcM+w7E5kARiNfLDV9JFtYJfB9ZIFO
	v2pwY4MAyQKBgQuua5LksbTQ7iLiHOyTIx7S+hiqHAO2IEFe8y9sXt/zYtcw/ErA
	qG14kmnd5KBTNGRbK8rt9hNpJfI1RWw+9wdXx4r9DhwpXJyZ/oCaibJmTBaYB25E
	0etPoLUlGZ4jACbvJdJaMC6v7BixTmZGPS1SbHB5e/R+Ae8hPPDtqHRyuA==
	-----END RSA PRIVATE KEY-----
	`
	block, _ := pem.Decode([]byte(keyFromSystem))
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil{
		fmt.Println("Failed here ")
		t.Fail()
	}
	
	message := "theAnswerIs42"
	actual := SignInput(message,privKey)
	expected := SignedMessageStructure{
		Message:   message,
		Signature: "kuoSwJ46gYKQzogQcG3Myje1TpHuo0WVKPqmlJ6Dy3pVm5Vwvz7sd/7Tl5h/145dmwV/OKemjxKvUezX6x8KKT3BUUTWuAfz6NsshPwYeJ5AC2mOAjUDeVKnURIYArXffSXi2+MBi+S5cgsql3eSPoLiFEv1z0uPuNUopuiKtzIIwrSwdzjsHsc8QGiaesXRIvLD+XAd18m7GWVGWFfp4eKxXmm+aZVS9UPNaBmhupZEHJm0ql4r9dMjdS6MJk4PTbxfu4qgIoQ3yxsOAS9y8yG/Zoob7r26xEMUXQTiC4eGT5o1OOLzVEgf1YMk7u0IneH3NQgMM5oxUYy5wMiDvZY=",
		Pubkey:    `-----BEGIN PUBLIC KEY-----
		MIIBCwKCAQIAsmg2jqguDen/Hp4/J1c/fgo0iMM3MzfFXKB9krcOSdOM3i2WvaYF
		XXexECihP8yQrGkUCVIJBTrqaG5rVpjDzMFsCSC1QZzNOPTN00CshPlA/pSJTKHx
		fCpHDiETu3bwnY4CuYUPO22kib0PI8PQztB0Q564PlpQTeFT5fhNw3ZWdrXP9siq
		Ib2aDbfO2YHFqtkkzEJvIuR1zxdXNoGGghuEOoVwSlXDPew2U5csMMSq4boKrBmR
		EXis+NO8N7pPiTLhUfcHniRPY6Rti5K9eWM5hSwbS5fVlF/5TZl/i+AOgNIFc5P+
		e9X37dCTco9O2OmIpL5nwRYyop80iGeaSm8CAwEAAQ==
		-----END PUBLIC KEY-----
		`,
	}
	if expected != actual {
		fmt.Println("Failed in comparison")
		t.Fail()
	}
}
