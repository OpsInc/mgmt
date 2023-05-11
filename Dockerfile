FROM golang:1.18-bullseye AS builder

RUN mkdir -p /build

WORKDIR /build

COPY . .

RUN go build -o /build/main cmd/mgmt/main.go

# Main Image
FROM public.ecr.aws/lambda/go:1

# COPY --from=builder mgmt ${LAMBDA_TASK_ROOT}
COPY --from=builder /build/main /var/task/

ENV GIN_MODE=release

CMD [ "main" ]
