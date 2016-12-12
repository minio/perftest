
* Clone minio-java source by
```bash
git clone https://github.com/minio/minio-java.git
```

* Go into minio-java source directory
```bash
cd minio-java
```

* Run functional test using gradle
```bash
./gradlew -Pendpoint=<YOUR-ENDPOINT> -PaccessKey=<YOUR-ACCESS-KEY> -PsecretKey=<YOUR-SECRET-KEY> runFunctionalTest
```
