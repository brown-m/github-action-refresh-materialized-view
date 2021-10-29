FROM python:3-slim
ADD . /app
WORKDIR /app
RUN apt-get update \
    && apt-get -y install libpq-dev gcc \
    && pip install psycopg2
ENV PYTHONPATH /app
CMD ["/app/main.py"]