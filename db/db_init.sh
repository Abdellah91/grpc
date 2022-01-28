#!/bin/sh

curl -X GET http://admin:$1@127.0.0.1:5984/_all_dbs | grep registration || curl -X PUT http://admin:$1@127.0.0.1:5984/registration





