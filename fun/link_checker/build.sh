docker build -t brijeshlakkad/link-checker:0.1 .
docker push brijeshlakkad/link-checker:0.1
kn service update link-checker \
--image brijeshlakkad/link-checker:0.1 \
--port 8080
