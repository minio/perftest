# mc-cat serial test.

Uses `mc cat` to download the object and verify its sanctity using its MD5Sum. 

# Usage.
- Install [mc](https://github.com/minio/mc).
- Upload an object. 
- Set the object path. 

  ```sh
  # export OBJECT=<mc-alias>/<bucket>/Object 
  export OBJECT=myminio/bucket/file
  ```
- Set the expected MD5 of the object.

  ```sh
  export MD5=xxxxxxxxxxxxxxxxxxxxx
  ```

- Set the number of times the test has to be run.

  ```sh
  export COUNT=100
  ```

- Run the program.
  ```sh
  go run mc-cat.go
  ```
- Check output.log for the result.

