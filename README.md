# s_senthivel.challenge
Intel Coding Cryto Sign Challenge

Crypto Sign
===========

Crypto Sign is your personal submission to the code challenge prompt found at the
bottom of this document.

Installing
----------

This is intended to be used on Unix based file systems.

To install this tool you need to have [Go 1.14][go] installed.

[go]: https://golang.org/

Usage
-----

    crypto-sign-challenge MESSAGE

`MESSAGE` is the message you wish to sign with your private key.  The message
must be 250 characters or less.

The following is an example with output included.

```
$ crypto-sign-challenge 'Welcome to the Jungle'
{
    "message": "Welcome to the Jungle",
    "signature": "MIGIAkIBHEc8FETUYOPze9YxePzBfN2OjbstTYQxfViHu6vziSfDbM5iJ8jCmH3LkScgoTNCRBAMBY407jDC/fYq88iN22cCQgCmytbObfzxtHWHpcYFvOb3PHHDKlv+rtAZJ/+AdxBvihjY/xRDi1PH8GhyEgzW7xzJ1KF7BhqmeMwH9pXUCx6JiA==",
    "pubkey": "-----BEGIN PUBLIC KEY-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQAxMXE/k5LOn1ZeSNgILi/fsDyHwwW\nSugmEndN786laNFUJ0Ulzit1FumnY71Op7Gwuqrv+YoqrEwpHtpnV8mLgvEBr9sX\ncNatfZzPtjOLpHzkVfLSCX94E7uNUZx13eigwugCsR87rn94CLRU3GDbLnLO6W4f\n12FkAhynQpvqaWNKpn8=\n-----END PUBLIC KEY-----\n"
}
```

Storage
-------

This project will generate a new private key if it does not exist and will store
it in:

    $HOME/.local/share/signer

Code Challenge Prompt
---------------------

We want to see how you think, how you write code, and the effort you put into
your work.

The following is derived from one of our products, so we're not wasting your
time solving a theoretical computer science problem, you're actually doing
something directly relevant to what you'll be doing if you join our team.

Using a language of your choice, provide an application that meets the following
requirements:

  - Given a string input of up to 250 characters, return a JSON response
    compliant to the schema defined below.
    - You are responsible for generating a public/private RSA or ECDSA keypair
      and persisting the keypair on the filesystem
      - Subsequent invocations of your application should read from the same
        files
  - Document your code, at a minimum defining parameter types and return values
    for any public methods
  - Include Unit Test(s) with instructions on how a Continuous Integration
    system can execute your test(s)
  - You may only use first order libraries, you may not use any third party
    libraries or packages.  For example, you may use the OpenSSL library, but
    you may not use any libraries built on top of OpenSSL.

JSON Schema for your application response:

```json
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "title": "Signed Identifier",
    "description": "Schema for a signed identifier",
    "type": "object",
    "required": [ "message", "signature", "pubkey" ],
    "properties": {
        "message": {
            "type": "string",
            "description": "original string provided as the input to your app"
        },
        "signature": {
            "type": "string",
            "description": "RFC 4648 compliant Base64 encoded cryptographic signature of the input, calculated using the private key and the SHA256 digest of the input"
        },
        "pubkey": {
            "type": "string",
            "description": "Base64 encoded string (PEM format) of the public key generated from the private key used to create the digital signature"
        }
    }
}
```


EXAMPLE

```
>./your-awesome-app "theAnswerIs42"
```

Returns:

```json
{
    "message":"theAnswerIs42",
    "signature":"MGUCMCDwlFyVdD620p0hRLtABoJTR7UNgwj8g2r0ipNbWPi4Us57YfxtSQJ3dAkHslyBbwIxAKorQmpWl9QdlBUtACcZm4kEXfL37lJ+gZ/hANcTyuiTgmwcEC0FvEXY35u2bKFwhA==",
    "pubkey":"-----BEGIN PUBLIC KEY-----\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEI5/0zKsIzou9hL3ZdjkvBeVZFKpDwxTb\nfiDVjHpJdu3+qOuaKYgsLLiO9TFfupMYHLa20IqgbJSIv/wjxANH68aewV1q2Wn6\nvLA3yg2mOTa/OHAZEiEf7bVEbnAov+6D\n-----END PUBLIC KEY-----\n"
}
```
