build:
	docker build -t mongweb1 .
run:
	docker run --rm -p 8228:8228 \
        --net=mongowebnetwork \
        --name=mongoweb1 \
        -e MONGODB_ADDR=mongodb://172.18.0.2:27017 \
        mongweb1



