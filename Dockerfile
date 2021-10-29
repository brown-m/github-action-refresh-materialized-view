FROM python:3-slim AS builder
ADD . /app
WORKDIR /app

RUN pip install --target=/app psycopg2

# A distroless container image with Python and some basics like SSL certificates
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/python3-debian10
COPY --from=builder /app /app
WORKDIR /app
ENV PYTHONPATH /app
CMD ["/cmd/main.py"]