FROM python:3-slim
ADD . /app

RUN pip3 install psycopg2-binary

CMD ["/app/main.py"]