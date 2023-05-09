FROM golang:1.20 as builder

COPY . /trainbot
WORKDIR /trainbot/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN cd fuzz_trainbot_avg && go-fuzz-build -libfuzzer -o fuzz_trainbot_avg.a && \
    clang -fsanitize=fuzzer fuzz_trainbot_avg.a -o fuzz_trainbot_avg.libfuzzer

RUN cd fuzz_trainbot_imutil && go-fuzz-build -libfuzzer -o fuzz_trainbot_imutil.a && \
    clang -fsanitize=fuzzer fuzz_trainbot_imutil.a -o fuzz_trainbot_imutil.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /trainbot/mayhem/fuzz_trainbot_avg/fuzz_trainbot_avg.libfuzzer /
COPY --from=builder /trainbot/mayhem/fuzz_trainbot_imutil/fuzz_trainbot_imutil.libfuzzer /