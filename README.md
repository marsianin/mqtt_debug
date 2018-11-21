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