# harbor-password
Manual Generation of Harbor-Encrypted Password for Database Password Reset

## Usage
- build the binary
```git clone https://github.com/snakeliwei/harbor-password.git && cd harbor-password && go mod tidy && go build```

- auto generate salt and password
```./harbor-password -p yourpass```

- Use yoursalt
```./harbor-password -p yourpass -salt yoursalt```

- Use sha1 or sha256 algorithm (default is sha256)
```./harbor-password -p yourpass -alg sha1```
