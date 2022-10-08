endpoint=localhost:8080
curl -X POST $endpoint/drivelog \
    -d { \
      token=sample1, \
      date="0000-01-01:00:00", \
      speed=100.00, \
      acceleration=100.00, \
      latitude=100.00, \
      longtitude=100.00, \
    }
