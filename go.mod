module github.com/uccps-samples/oc

go 1.16

require (
	github.com/AaronO/go-git-http v0.0.0-20161214145340-1d9485b3a98f
	github.com/MakeNowJust/heredoc v0.0.0-20170808103936-bb23615498cd
	github.com/RangelReale/osincli v0.0.0-20160924135400-fababb0555f2
	github.com/alexbrainman/sspi v0.0.0-20180613141037-e580b900e9f5
	github.com/alicebob/sqlittle v1.4.0
	github.com/apcera/gssapi v0.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go v1.35.24
	github.com/blang/semver v3.5.1+incompatible
	github.com/containers/image/v5 v5.10.2
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v20.10.3+incompatible
	github.com/docker/go-units v0.4.0
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7
	github.com/elazarl/goproxy v0.0.0-20190911111923-ecfe977594f1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/fsouza/go-dockerclient v1.7.1
	github.com/ghodss/yaml v1.0.0
	github.com/golangplus/testing v1.0.0 // indirect
	github.com/gonum/graph v0.0.0-20170401004347-50b27dea7ebb
	github.com/gonum/mathext v0.0.0-20181121095525-8a4bf007ea55 // indirect
	github.com/gonum/stat v0.0.0-20181125101827-41a0da705a5b // indirect
	github.com/google/go-cmp v0.5.5
	github.com/joelanford/ignore v0.0.0-20210610194209-63d4919d8fb2
	github.com/mitchellh/go-wordwrap v1.0.0
	github.com/moby/buildkit v0.0.0-20181107081847-c3a857e3fca0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.2-0.20190823105129-775207bd45b6
	github.com/openshift/api v3.9.0+incompatible // indirect
	github.com/openshift/client-go v3.9.0+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/russross/blackfriday v1.5.2
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/uccps-samples/api v0.0.0-20221025085120-91bb4873437a
	github.com/uccps-samples/build-machinery-go v0.0.0-20221024073655-3930cec85ae1
	github.com/uccps-samples/client-go v0.0.0-20221025085836-81245c36050d
	github.com/uccps-samples/library-go v0.0.0-20221031055730-9c42ab8f039d
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
	gopkg.in/ldap.v2 v2.5.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/apiserver v0.22.1
	k8s.io/cli-runtime v0.21.1
	k8s.io/client-go v0.22.1
	k8s.io/component-base v0.22.1
	k8s.io/klog/v2 v2.9.0
	k8s.io/kubectl v0.21.0-beta.1
	k8s.io/utils v0.0.0-20210707171843-4b05e18ac7d9
	sigs.k8s.io/yaml v1.2.0
)

replace (
	github.com/Microsoft/hcsshim => github.com/Microsoft/hcsshim v0.8.7
	github.com/apcera/gssapi => github.com/openshift/gssapi v0.0.0-20161010215902-5fb4217df13b
	k8s.io/apimachinery => github.com/openshift/kubernetes-apimachinery v0.0.0-20210521074607-b6b98f7a1855
	k8s.io/cli-runtime => github.com/openshift/kubernetes-cli-runtime v0.0.0-20210521074950-112a61d2624f
	k8s.io/client-go => github.com/openshift/kubernetes-client-go v0.0.0-20210521075216-71b63307b5df
	k8s.io/kubectl => github.com/openshift/kubernetes-kubectl v0.0.0-20210521075729-633333dfccda
)
