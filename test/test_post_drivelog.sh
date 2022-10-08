endpoint=localhost:8080
curl -X POST $endpoint/drivelog \
    -d token='sample1' \
    -d date="0000-01-01:00:00" \
    -d speed=100.00 \
    -d acceleration=100.00 \
    -d latitude=100.00 \
    -d longtitude=100.00
