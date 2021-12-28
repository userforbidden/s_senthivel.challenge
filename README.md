# s_senthivel.challenge
Intel Coding Crypto Sign Challenge

Usage
-----

    s_senthivel.challenge MESSAGE

`MESSAGE` is the message you wish to sign with your private key.  The message
must be 250 characters or less.

The following is an example with output included.

```
$ s_senthivel.challenge 'Welcome to the Jungle'
{
    "message": "Welcome to the Jungle",
    "signature": "MIGIAkIBHEc8FETUYOPze9YxePzBfN2OjbstTYQxfViHu6vziSfDbM5iJ8jCmH3LkScgoTNCRBAMBY407jDC/fYq88iN22cCQgCmytbObfzxtHWHpcYFvOb3PHHDKlv+rtAZJ/+AdxBvihjY/xRDi1PH8GhyEgzW7xzJ1KF7BhqmeMwH9pXUCx6JiA==",
    "pubkey": "-----BEGIN PUBLIC KEY-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQAxMXE/k5LOn1ZeSNgILi/fsDyHwwW\nSugmEndN786laNFUJ0Ulzit1FumnY71Op7Gwuqrv+YoqrEwpHtpnV8mLgvEBr9sX\ncNatfZzPtjOLpHzkVfLSCX94E7uNUZx13eigwugCsR87rn94CLRU3GDbLnLO6W4f\n12FkAhynQpvqaWNKpn8=\n-----END PUBLIC KEY-----\n"
}
```

Code Challenge Requirements
---------------------
  - [X] Given a string input of up to 250 characters, return a JSON response
    compliant to the schema defined below.
    - You are responsible for generating a public/private RSA or ECDSA keypair
      and persisting the keypair on the filesystem
      - Subsequent invocations of your application should read from the same
        files
  - Document your code, at a minimum defining parameter types and return values
    for any public methods
  - Include Unit Test(s) with instructions on how a Continuous Integration
    system can execute your test(s)
  - You may only use first order libraries, you may not use any third party libraries or packages.  For example, you may use the OpenSSL library, but you may not use any libraries built on top of OpenSSL.
  - This project will generate a new private key if it does not exist and will store
it in:
    $HOME/.local/share/signer




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


To run tests
```
go test -v ./...
```

To run the program, run the following commands:
```
docker build -t <image name> .
docker run <image name> <input message>
```