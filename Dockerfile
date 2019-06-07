FROM centos:7 as build-tools

LABEL maintainer "Red Hat Storage Management <rh-storage-mgmt@redhat.com>"
LABEL author "Umanga Chapagain <uchapaga@redhat.com>"

ENV LANG=en_US.utf8
ENV GOPATH /tmp/go
ARG GO_PACKAGE_PATH=github.com/monstorak/monstorak-operator

RUN yum install epel-release -y \
    && yum install --enablerepo=centosplus install -y --quiet \
    findutils \
    git \
    golang \
    make \
    tar \
    wget \
    which \
    kubectl \
    yamllint \
    && yum clean all

# install dep
RUN mkdir -p $GOPATH/bin && chmod a+rwx $GOPATH \
    && curl -L -s https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 -o dep \
    && echo "7479cca72da0596bb3c23094d363ea32b7336daa5473fa785a2099be28ecd0e3  dep" > dep-linux-amd64.sha256 \
    && sha256sum -c dep-linux-amd64.sha256 \
    && rm dep-linux-amd64.sha256 \
    && chmod +x ./dep \
    && mv dep $GOPATH/bin/dep

ENV PATH=$PATH:$GOPATH/bin

# download, verify and install openshift client tools (oc and kubectl)
WORKDIR /tmp

# install operator-sdk (from git with no history and only the tag)
RUN mkdir -p $GOPATH/src/github.com/operator-framework \
    && cd $GOPATH/src/github.com/operator-framework \
    && git clone --depth 1 -b v0.6.0 https://github.com/operator-framework/operator-sdk \
    && cd operator-sdk \
    && make dep \
    && make install

RUN mkdir -p ${GOPATH}/src/${GO_PACKAGE_PATH}/

WORKDIR ${GOPATH}/src/${GO_PACKAGE_PATH}

ENTRYPOINT [ "/bin/bash" ]
#--------------------------------------------------------------------

FROM build-tools as builder
ARG VERBOSE=2
COPY . .
RUN make VERBOSE=${VERBOSE} build
# RUN make VERBOSE=${VERBOSE} test

#--------------------------------------------------------------------

FROM registry.access.redhat.com/ubi7-dev-preview/ubi-minimal:latest

LABEL maintainer "Red Hat Storage Management <rh-storage-mgmt@redhat.com>"
LABEL author "Umanga Chapagain <uchapaga@redhat.com>"

ENV LANG=en_US.utf8
ENV GOPATH=/tmp/go
ARG GO_PACKAGE_PATH=github.com/monstorak/monstorak-operator

COPY --from=builder ${GOPATH}/src/${GO_PACKAGE_PATH}/out/operator /usr/local/bin/monstorak-operator

USER 10001

ENTRYPOINT [ "/usr/local/bin/monstorak-operator" ]
