# JS Upload Load.

Uploads the given file for a total of 400 times and 80 asynchronous uploads.

# Configure.

- open minio.json 

- Fill in Secret Key, AccessKey, IP's of the nodes and the path of the file to upload.

- Here is the sample minio.json.
  
  ```sh
  {
    "access_key": "Z7IXGOO6BZ0REAN1Q26I",
    "public_ips": [
          "localhost",
          "192.168.1.10"
    ],
    "secret_key": "+m4G6buANjXWX8B/6/KUQRzbAi/l47aX7M+BG2+4",
    "file":"/home/user/minio.json"
  }
  ```

# Run.
```sh
  $ npm install minio@3 async uuid 

  $ node pound-it.js
```
