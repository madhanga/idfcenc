# How to use

## Prerequisite
- Go 1.2
- Java 21.0.1 

## Run in go
Run in go which will give a encypted data for the given data in argument.
```
go run main.go "hello this is a sample data. வணக்கம்"
```

Output:
```
Encrypted: NXFcXVNsTn0xTDlAfTRXWH+PX/u1mmK910mpbXgNWjCwOuaw244vkv6mPGEyv+bmzmrvwVgXaFnPRa76+b6Y2zhkh0pApEBnzYr3rzYvlsM=
Decrypted: hello this is a sample data. வணக்கம்
```

## Run in Java
Take the encypted value emitted from Go program above and pass teh same to java
```
java Main.java "NXFcXVNsTn0xTDlAfTRXWH+PX/u1mmK910mpbXgNWjCwOuaw244vkv6mPGEyv+bmzmrvwVgXaFnPRa76+b6Y2zhkh0pApEBnzYr3rzYvlsM="
```

The output will the plain text you originally passed it to Go.

