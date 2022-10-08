endpoint=localhost:8080
curl -X GET $endpoint/drivelog \
    -d token='sample1' \
    -d date="0000-01-01:00:00"
