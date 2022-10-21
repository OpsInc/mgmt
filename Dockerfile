FROM golang:1.18-bullseye AS builder

RUN mkdir -p /build

WORKDIR /build

COPY . .

RUN go build -o /mgmt .

# Main Image
FROM public.ecr.aws/lambda/go:1

COPY --from=builder mgmt ${LAMBDA_TASK_ROOT}

CMD [ "mgmt.main" ]
