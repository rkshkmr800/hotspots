FROM ubuntu:16.04

LABEL Author="Rakesh kumar Chanderki"

RUN apt-get update -y && \
    apt-get install -y python-pip python-dev

# # We copy just the requirements.txt first to leverage Docker cache
#COPY ./requirements.txt ./backend/ml-framework/requirements.txtt

WORKDIR /backend/ml-framework

COPY ./backend/ml-framework . 

RUN pip install -r requirements.txt

ENTRYPOINT [ "python" ]

CMD [ "app.py" ]