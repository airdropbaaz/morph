ARG GO_VERSION=1.19
ARG RUST_VERSION=nightly-2022-12-10
ARG CARGO_CHEF_TAG=0.1.41

FROM ubuntu:20.04

RUN apt-get update && ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime

# Install basic packages
RUN apt-get install build-essential curl wget git pkg-config -y
# Install dev-packages
RUN apt-get install libclang-dev libssl-dev llvm software-properties-common -y
# Install golang
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt install golang-1.19-go -y
ENV PATH="/usr/lib/go-1.19/bin:${PATH}"

# Install Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"
ENV CARGO_HOME=/root/.cargo
# Add Toolchain
ARG RUST_VERSION
RUN rustup toolchain install ${RUST_VERSION}
ARG CARGO_CHEF_TAG
RUN cargo install cargo-chef --locked --version ${CARGO_CHEF_TAG} \
    && rm -rf $CARGO_HOME/registry/