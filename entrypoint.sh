make cover
curl -Os https://uploader.codecov.io/latest/linux/codecov
chmod 777 codecov

./codecov -t ${CODECOV}