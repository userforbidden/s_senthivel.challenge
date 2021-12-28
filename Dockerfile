FROM golang:1.17

ENV PROJECT_URI=github.com/userforbidden/s_senthivel.challenge
ENV PROJECT_DIR=${GOPATH}/src/${PROJECT_URI}

RUN mkdir -p ${PROJECT_DIR}

WORKDIR ${PROJECT_DIR}

COPY . ${PROJECT_DIR}

RUN go install ${PROJECT_URI}/...


ENTRYPOINT ["sh","-c","$GOPATH/bin/s_senthivel.challenge $0"]

CMD [""]
