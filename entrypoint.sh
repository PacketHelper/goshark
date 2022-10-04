make cover
curl -Os https://uploader.codecov.io/latest/linux/codecov
chmod a+x codecov

./codecov -t ${CODECOV}