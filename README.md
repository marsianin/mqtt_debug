#### Run subscriber
run in command line:
```sh
go run sub/sub.go -t "<topic_id>" -i "<client_id>" -u "<user_id>" -p '<password>'
```

#### Run publisher
run in command line:
```sh
go run pub/pub.go -t "<topic_id>" -i "<client_id>" -u "<user_id>" -p '<password>'
```

#### Generate VerneMQ password
run in command line:
```sh
go run security/gen.go -p "<password>"
```

#### Generate Root CA
run in command line:
```sh
make gen-ca
```

#### Generate Client/Server cert
run in command line:
```sh
make gen-server-cert
```


### SSL debug

```sh
openssl s_client -connect <host>:<port> -CAfile <ca cert file> -nbio -debug -msg -state -cert <cert> -key <secret key>
```

SSL handshake:
```sh
SSL_connect:before/connect initialization
SSL_connect:unknown state
SSL_connect:error in unknown state
SSL_connect:SSLv3 read server hello A
depth=1 C = RU, ST = Denial, L = Test, O = Test
verify return:1
depth=0 C = RU, ST = Denial, L = Test, O = Test, CN = server.application
verify return:1
SSL_connect:SSLv3 read server certificate A
SSL_connect:SSLv3 read server key exchange A
SSL_connect:SSLv3 read server certificate request A
SSL_connect:SSLv3 read server done A
SSL_connect:SSLv3 write client certificate A
0230 - 70 65 6e 53 53 4c 20 47-65 6e 65 72 61 74 65 64   penSSL Generated
SSL_connect:SSLv3 write client key exchange A
SSL_connect:SSLv3 write certificate verify A
SSL_connect:SSLv3 write change cipher spec A
SSL_connect:SSLv3 write finished A
SSL_connect:SSLv3 flush data
SSL_connect:error in SSLv3 read finished A
SSL_connect:error in SSLv3 read finished A
SSL_connect:SSLv3 read finished A
SSL handshake has read 2484 bytes and written 2604 bytes
New, TLSv1/SSLv3, Cipher is ECDHE-RSA-AES256-GCM-SHA384
SSL-Session:
SSL3 alert read:warning:close notify
SSL3 alert write:warning:close notify
```

####Verify cert content
````sh
 openssl x509 -noout -text -in certs/client.crt
 ````
 Response
 ````sh
 Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 3 (0x3)
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=RU, ST=Denial, L=Test, O=Test
        Validity
            Not Before: Nov 21 17:18:43 2018 GMT
            Not After : Nov 20 17:18:43 2023 GMT
        Subject: C=RU, ST=Denial, L=Test, O=Test, CN=client.application
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:c5:07:33:c4:dd:6e:87:57:11:31:40:b2:a3:11:
                    09:a7:3d:1c:8d:42:9c:43:36:b4:d4:88:22:dc:83:
                    3c:95:97:df:b2:e6:9a:84:90:96:3f:9c:e8:8e:cb:
                    eb:4e:53:b1:6a:0e:09:db:2a:8a:07:db:d7:f9:54:
                    7a:04:de:12:76:89:58:a2:f6:80:2d:d6:78:fc:f0:
                    62:33:33:4b:53:a8:c5:fd:07:89:41:b1:a9:ba:0a:
                    78:40:34:1a:89:51:31:ea:a7:72:d7:f3:f6:ef:54:
                    03:85:c5:68:b4:bc:52:41:34:14:98:5a:34:d4:77:
...
                    0d:37
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints:
                CA:FALSE
            Netscape Comment:
                OpenSSL Generated Certificate
            X509v3 Subject Key Identifier:
                3C:A2:54:22...:5C:28:1B:2E
            X509v3 Authority Key Identifier:
                keyid:14:70:...:C5:23:CE

            X509v3 Subject Alternative Name:
                IP Address:127.0.0.1
    Signature Algorithm: sha1WithRSAEncryption
         29:9b:75:8d:1a:29:99:d8:7b:ca:ab:16:35:a3:d0:9a:f2:11:
         5f:2a:79:7b:f1:42:f9:06:5c:92:e6:fa:af:14:14:37:6d:a4:
         81:0d:26:3f:99:1d:d0:f3:9c:aa:62:be:89:ba:1b:89:88:15:
...

 ````

 #### Verify certificate request
 ````sh
openssl req -noout -text -in certs/client.csr
```
Result:
````sh
Certificate Request:
    Data:
        Version: 0 (0x0)
        Subject: C=RU, ST=Denial, L=Test, O=Test, CN=client.application
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:c5:07:33:c4:dd:6e:87:57:11:31:40:b2:a3:11:
                    09:a7:3d:1c:8d:42:9c:43:36:b4:d4:88:22:dc:83:
                    3c:95:97:df:b2:e6:9a:84:90:96:3f:9c:e8:8e:cb:
                    ...
                Exponent: 65537 (0x10001)
        Attributes:
            a0:00
    Signature Algorithm: sha256WithRSAEncryption
         83:65:25:78:77:14:1b:45:b9:49:79:91:0b:fc:62:6e:7a:50:
         d6:5a:ed:07:16:4e:f4:34:33:60:cf:05:d1:74:cd:30:ac:bd:
         a7:d8:3a:b8:40:7d:07:b2:6b:da:98:90:64:d1:88:bc:20:21:
...
```
