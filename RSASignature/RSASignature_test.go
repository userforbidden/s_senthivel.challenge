package RSASignature

import (
	"testing"
)

func TestOverall(t *testing.T) {

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
