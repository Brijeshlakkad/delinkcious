FROM python:3.8-alpine

RUN apk add --no-cache \
  # required by grpc
  g++ \
  musl-dev \
  libffi-dev

COPY requirements.txt ./

RUN apk add build-base

RUN pip install --upgrade pip

RUN pip install --no-cache-dir -r requirements.txt \
  # Clean up
  && rm requirements.txt \
  && apk del \
  g++ \
  musl-dev \
  libffi-dev