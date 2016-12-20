"use strict";

var Minio = require('minio');
var settings = require('./minio.json');
var uuid = require('uuid');
var async = require('async');

var minioClients = [];
for (var i = 0; i < settings.public_ips.length; i++) {
    minioClients.push(new Minio.Client({
        endPoint: settings.public_ips[i],
        port: 9000,
        secure: false,
        accessKey: settings.access_key,
        secretKey: settings.secret_key
    }));
}
let file = settings.file;

minioClients[0].makeBucket('test', 'us-east-1', (err) => {
    if (err) {
        console.log("error creating the bucket", err);
    }
    async.map(minioClients, (client, callback) => {

        if (err) {
            console.log("error creating the bucket", err);
        }
        async.timesLimit(400, 80, (n, callback) => {
            var uuidStr = uuid.v4()
            client.fPutObject('test', uuidStr, file, 'application/octet-stream', (err) => {
                if (err) {
                    console.log(err);
                } else {
                    process.stdout.write('.');
                }
                var size = 0
                // Get a full object.
                client.getObject('test', uuidStr, function(e, dataStream) {
                    if (e) {
                        return console.log(e)
                    }
                    dataStream.on('data', function(chunk) {
                        size += chunk.length
                    })
                    dataStream.on('end', function() {
                        console.log("End. Total size = " + size)
                    })
                    dataStream.on('error', function(e) {
                        console.log(e)
                    })
                })
                callback();
            });
        }, (err) => {
            if (err) {
                console.log(err);
            }
            return callback();
        });

    }, (err) => {
        console.log(err);
    });
});
