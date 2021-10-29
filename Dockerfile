FROM python:3-slim
ADD . /app
WORKDIR /app
RUN pip install psycopg2
ENV PYTHONPATH /app
CMD ["main.py"]