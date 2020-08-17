/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resourcelock

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/chinglinwen/pkg/etcdutil"

	// "go.etcd.io/etcd/clientv3"

	v1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

// TODO: This is almost a exact replica of Endpoints lock.
// going forwards as we self host more and more components
// and use ConfigMaps as the means to pass that configuration
// data we will likely move to deprecate the Endpoints lock.

type EtcdLock struct {
	// ConfigMapMeta should contain a Name and a Namespace of a
	// ConfigMapMeta object that the LeaderElector will attempt to lead.
	// ConfigMapMeta metav1.ObjectMeta
	// Client        corev1client.ConfigMapsGetter
	// LockConfig ResourceLockConfig
	// EtcdEndpoints []string

	// *clientv3.Client
	// usetls bool
	WhoAmI string
	*etcdutil.Client

	// value []byte
	event string
	// cm *v1.ConfigMap
}

// Get returns the election record from a ConfigMap Annotation
func (cml *EtcdLock) Get(ctx context.Context) (*LeaderElectionRecord, []byte, error) {
	var record LeaderElectionRecord
	var err error
	recordBytes, err := cml.Client.Get(ctx, LeaderElectionRecordAnnotationKey)
	if err != nil {
		return nil, nil, err
	}
	if err := json.Unmarshal([]byte(recordBytes), &record); err != nil {
		return nil, nil, err
	}
	return &record, []byte(recordBytes), nil
}

// Create attempts to create a LeaderElectionRecord annotation
func (cml *EtcdLock) Create(ctx context.Context, ler LeaderElectionRecord) error {
	recordBytes, err := json.Marshal(ler)
	if err != nil {
		return err
	}
	return cml.Client.Put(ctx, LeaderElectionRecordAnnotationKey, string(recordBytes))
}

// Update will update an existing annotation on a given resource.
func (cml *EtcdLock) Update(ctx context.Context, ler LeaderElectionRecord) error {
	return cml.Create(ctx, ler)
}

// RecordEvent in leader election while adding meta-data
func (cml *EtcdLock) RecordEvent(s string) {
	cml.event = fmt.Sprintf("identity: %v, event: %v, type: %v, %v", cml.WhoAmI, s, v1.EventTypeNormal, "LeaderElection")
	klog.Infof("event: %v", cml.event)
}

// Describe is used to convert details on current resource lock
// into a string
func (cml *EtcdLock) Describe() string {
	return "etcdlock/leader"
}

// Identity returns the Identity of the lock
func (cml *EtcdLock) Identity() string {
	return cml.WhoAmI
}
