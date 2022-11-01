FROM openshift/golang-builder@sha256:4820580c3368f320581eb9e32cf97aeec179a86c5749753a14ed76410a293d83 AS builder
ENV __doozer=update BUILD_RELEASE=202202160843.p0.gf93da17.assembly.stream BUILD_VERSION=v4.10.0 OS_GIT_MAJOR=4 OS_GIT_MINOR=10 OS_GIT_PATCH=0 OS_GIT_TREE_STATE=clean OS_GIT_VERSION=4.10.0-202202160843.p0.gf93da17.assembly.stream SOURCE_GIT_TREE_STATE=clean 
ENV __doozer=merge OS_GIT_COMMIT=f93da17 OS_GIT_VERSION=4.10.0-202202160843.p0.gf93da17.assembly.stream-f93da17 SOURCE_DATE_EPOCH=1644962438 SOURCE_GIT_COMMIT=f93da179fe606775f8249fed96e0b9903d9188ed SOURCE_GIT_TAG=uccp-clients-4.6.0-202006250705.p0-704-gf93da179f SOURCE_GIT_URL=https://github.com/uccps-samples/oc 
WORKDIR /go/src/github.com/uccps-samples/oc
COPY . .
RUN make build --warn-undefined-variables

FROM openshift/ose-base:v4.10.0-202202160023.p0.g544601e.assembly.stream
ENV __doozer=update BUILD_RELEASE=202202160843.p0.gf93da17.assembly.stream BUILD_VERSION=v4.10.0 OS_GIT_MAJOR=4 OS_GIT_MINOR=10 OS_GIT_PATCH=0 OS_GIT_TREE_STATE=clean OS_GIT_VERSION=4.10.0-202202160843.p0.gf93da17.assembly.stream SOURCE_GIT_TREE_STATE=clean 
ENV __doozer=merge OS_GIT_COMMIT=f93da17 OS_GIT_VERSION=4.10.0-202202160843.p0.gf93da17.assembly.stream-f93da17 SOURCE_DATE_EPOCH=1644962438 SOURCE_GIT_COMMIT=f93da179fe606775f8249fed96e0b9903d9188ed SOURCE_GIT_TAG=uccp-clients-4.6.0-202006250705.p0-704-gf93da179f SOURCE_GIT_URL=https://github.com/uccps-samples/oc 
COPY --from=builder /go/src/github.com/uccps-samples/oc/oc /usr/bin/
RUN for i in kubectl uccp-deploy uccp-docker-build uccp-sti-build uccp-git-clone uccp-manage-dockerfile uccp-extract-image-content uccp-recycle; do ln -sf /usr/bin/oc /usr/bin/$i; done

LABEL \
        io.k8s.display-name="OpenShift Client" \
        io.k8s.description="OpenShift is a platform for developing, building, and deploying containerized applications." \
        io.uccp.tags="openshift,cli" \
        License="GPLv2+" \
        vendor="Red Hat" \
        name="utccp-components/cli" \
        com.redhat.component="uccp-enterprise-cli-container" \
        io.uccp.maintainer.product="OpenShift Container Platform" \
        io.uccp.maintainer.component="oc" \
        release="202202160843.p0.gf93da17.assembly.stream" \
        io.uccp.build.commit.id="f93da179fe606775f8249fed96e0b9903d9188ed" \
        io.uccp.build.source-location="https://github.com/uccps-samples/oc" \
        io.uccp.build.commit.url="https://github.com/uccps-samples/oc/commit/f93da179fe606775f8249fed96e0b9903d9188ed" \
        version="v4.10.0"

