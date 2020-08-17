module github.com/chinginwen/pkg

go 1.15

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace github.com/chinglinwen/pkg => ../github.com/chinglinwen/pkg

replace github.com/chinglinwen/pkg/etcdutil => ../github.com/chinglinwen/pkg/etcdutil

require (
	code.cloudfoundry.org/tlsconfig v0.0.0-20200131000646-bbe0f8da39b3
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/uuid v1.1.1 // indirect
	go.etcd.io/etcd v3.3.22+incompatible // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/grpc v1.26.0 // indirect
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20200815024018-e34d1aa459f9 // indirect
)

replace github.com/chinglinwen/pkg => ../github.com/chinglinwen/pkg

replace sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.1.12

replace sigs.k8s.io/kubebuilder => sigs.k8s.io/kubebuilder v1.0.8

replace github.com/markbates/inflect => github.com/markbates/inflect v1.0.4

replace github.com/kubernetes-incubator/reference-docs => github.com/kubernetes-sigs/reference-docs v0.0.0-20170929004150-fcf65347b256

replace sigs.k8s.io/structured-merge-diff => sigs.k8s.io/structured-merge-diff v1.0.1-0.20191108220359-b1b620dd3f06

replace k8s.io/api => k8s.io/api v0.17.3

replace k8s.io/apimachinery => k8s.io/apimachinery v0.17.3

replace k8s.io/apiserver => k8s.io/apiserver v0.17.3

replace k8s.io/client-go => k8s.io/client-go v0.17.3

replace k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a

replace k8s.io/gengo => k8s.io/gengo v0.0.0-20190822140433-26a664648505

replace k8s.io/utils => k8s.io/utils v0.0.0-20191114184206-e782cd3c129f

replace sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.5.1
