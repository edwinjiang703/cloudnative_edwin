FROM python:3.7
RUN pip install kopf kubernetes
COPY handlers.py /handlers.py
CMD kopf run --standalone /handlers.py
