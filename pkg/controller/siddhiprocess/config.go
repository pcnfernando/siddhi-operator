/*
 * Copyright (c) 2019 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package siddhiprocess

// Default configurations stored as constants. Further these constants used by the Configurations() function.
const (
	SiddhiHome           string = "/home/siddhi_user/siddhi-runner-0.1.0/"
	SiddhiRunnerImage    string = "siddhiio/siddhi-runner-alpine"
	SiddhiRunnerPath     string = "wso2/runner/"
	SiddhiRunnerImageTag string = "0.1.0"
	SiddhiCMExt          string = "-siddhi"
	SiddhiExt            string = ".siddhi"
	SiddhiFileRPath      string = "wso2/runner/deployment/siddhi-files/"
	ContainerName        string = "siddhi-runner-runtime"
	DepConfigName        string = "deploymentconfig"
	DepConfMountPath     string = "tmp/configs/"
	DepConfParameter     string = "-Dconfig="
	DepCMExt             string = "-deployment.yaml"
	Shell                string = "sh"
	RunnerRPath          string = "bin/runner.sh"
	HostName             string = "siddhi"
	OperatorName         string = "siddhi-operator"
	OperatorVersion      string = "0.1.1"
	CRDName              string = "SiddhiProcess"
	ReadWriteOnce        string = "ReadWriteOnce"
	ReadOnlyMany         string = "ReadOnlyMany"
	ReadWriteMany        string = "ReadWriteMany"
	PVCExt               string = "-pvc"
	FilePersistentPath   string = "wso2/runner/siddhi-app-persistence"
	ParserDomain         string = "http://siddhi-parser."
	ParserDefaultContext string = ".svc.cluster.local:9090/siddhi-parser/parse"
	ParserNATSContext    string = ".svc.cluster.local:9090/siddhi-parser/failover"
	PVCSize              string = "1Gi"
	NATSAPIVersion       string = "nats.io/v1alpha2"
	STANAPIVersion       string = "streaming.nats.io/v1alpha1"
	NATSKind             string = "NatsCluster"
	STANKind             string = "NatsStreamingCluster"
	NATSExt              string = "-nats"
	STANExt              string = "-stan"
	NATSClusterName      string = "siddhi-nats"
	STANClusterName      string = "siddhi-stan"
	NATSDefaultURL       string = "nats://siddhi-nats:4222"
	NATSTCPHost          string = "siddhi-nats:4222"
	NATSMSType           string = "nats"
	TCP                  string = "tcp"
	FExtOne              string = "-1"
	FExtTwo              string = "-2"
	NATSSize             int    = 1
	NATSTimeout          int    = 5
	DefaultRTime         int    = 1
	DeploymentSize       int32  = 1
)

// State persistence config is the different string constant used by the deployApp() function. This constant holds a YAML object
// which used to change the deployment.yaml file of the siddhi-runner image.
const (
	StatePersistenceConf string = `
state.persistence:
  enabled: true
  intervalInMin: 1
  revisionsToKeep: 2
  persistenceStore: io.siddhi.distribution.core.persistence.FileSystemPersistenceStore
  config:
    location: siddhi-app-persistence
`
)

// These are all other relevant constants that used by the operator. But these constants are not configuration varibles.
// That is why this has been seperated.
const (
	Push           string = "PUSH"
	Pull           string = "PULL"
	Failover       string = "failover"
	Default        string = "default"
	Distributed    string = "distributed"
	ProcessApp     string = "process"
	PassthroughApp string = "passthrough"
)

// Configs is the struct definition of the object which used to bundle the all default configurations.
type Configs struct {
	SiddhiHome           string
	SiddhiRunnerImage    string
	SiddhiRunnerImageTag string
	SiddhiCMExt          string
	SiddhiExt            string
	SiddhiFileRPath      string
	SiddhiRunnerPath     string
	ContainerName        string
	DepConfigName        string
	DepConfMountPath     string
	DepConfParameter     string
	DepCMExt             string
	Shell                string
	RunnerRPath          string
	HostName             string
	OperatorName         string
	OperatorVersion      string
	CRDName              string
	ReadWriteOnce        string
	ReadOnlyMany         string
	ReadWriteMany        string
	PVCExt               string
	FilePersistentPath   string
	ParserDomain         string
	ParserDefaultContext string
	ParserNATSContext    string
	PVCSize              string
	NATSAPIVersion       string
	STANAPIVersion       string
	NATSKind             string
	STANKind             string
	NATSExt              string
	STANExt              string
	NATSClusterName      string
	STANClusterName      string
	NATSDefaultURL       string
	NATSTCPHost          string
	NATSMSType           string
	TCP                  string
	FExtOne              string
	FExtTwo              string
	NATSSize             int
	NATSTimeout          int
	DefaultRTime         int
	DeploymentSize       int32
}

// Configurations function returns the default config object. Here all the configs used as constants and budle together into a
// object and then returns that object. This object used to differenciate default configs from other variables.
func Configurations() Configs {
	configs := Configs{
		SiddhiHome:           SiddhiHome,
		SiddhiRunnerImage:    SiddhiRunnerImage,
		SiddhiRunnerImageTag: SiddhiRunnerImageTag,
		SiddhiCMExt:          SiddhiCMExt,
		SiddhiExt:            SiddhiExt,
		SiddhiFileRPath:      SiddhiFileRPath,
		SiddhiRunnerPath:     SiddhiRunnerPath,
		ContainerName:        ContainerName,
		DepConfigName:        DepConfigName,
		DepConfMountPath:     DepConfMountPath,
		DepConfParameter:     DepConfParameter,
		DepCMExt:             DepCMExt,
		Shell:                Shell,
		RunnerRPath:          RunnerRPath,
		HostName:             HostName,
		OperatorName:         OperatorName,
		OperatorVersion:      OperatorVersion,
		CRDName:              CRDName,
		ReadWriteOnce:        ReadWriteOnce,
		ReadOnlyMany:         ReadOnlyMany,
		ReadWriteMany:        ReadWriteMany,
		PVCExt:               PVCExt,
		FilePersistentPath:   FilePersistentPath,
		ParserDomain:         ParserDomain,
		ParserDefaultContext: ParserDefaultContext,
		ParserNATSContext:    ParserNATSContext,
		PVCSize:              PVCSize,
		NATSAPIVersion:       NATSAPIVersion,
		STANAPIVersion:       STANAPIVersion,
		NATSKind:             NATSKind,
		STANKind:             STANKind,
		NATSExt:              NATSExt,
		STANExt:              STANExt,
		NATSClusterName:      NATSClusterName,
		STANClusterName:      STANClusterName,
		NATSDefaultURL:       NATSDefaultURL,
		NATSTCPHost:          NATSTCPHost,
		NATSMSType:           NATSMSType,
		TCP:                  TCP,
		FExtOne:              FExtOne,
		FExtTwo:              FExtTwo,
		NATSSize:             NATSSize,
		NATSTimeout:          NATSTimeout,
		DefaultRTime:         DefaultRTime,
		DeploymentSize:       DeploymentSize,
	}
	return configs
}
