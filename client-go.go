package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	dv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	restClient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesConfig struct {
	ClientSet *kubernetes.Clientset //ClientSet客户端
	Config    *restClient.Config
	Namespace string
}

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "./config/config")
	if err != nil {
		panic("config 文件加载失败")
	}
	kubernetesConfig := NewKubernetesConfig(config, "testyaml")
	ctx := context.Background()
	//_, err = kubernetesConfig.CreateNamespaces(ctx, "./yaml/namespace.yaml")
	//_, err = kubernetesConfig.CreateDeployment(ctx, "./yaml/deployment.yaml")
	//_, err = kubernetesConfig.CreateServices(ctx, "./yaml/services.yaml")
	//namespacesList, err := kubernetesConfig.QueryNamespaces(ctx)
	//deploymentList, err := kubernetesConfig.QueryDeployment(ctx)
	servicesList, err := kubernetesConfig.QueryServices(ctx)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(servicesList)
}

func NewKubernetesConfig(config *restClient.Config, namespace string) *KubernetesConfig {
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &KubernetesConfig{
		Config:    config,
		ClientSet: clientSet,
		Namespace: namespace,
	}

}

/**
 * @Function: CreateNamespaces
 * @Description: 创建全新的命名空间
 * @receiver k
 * @param ctx
 * @param fileName	yaml文件地址
 * @return *v1.Namespace
 * @return error
 */
func (k *KubernetesConfig) CreateNamespaces(ctx context.Context, fileName string) (*v1.Namespace, error) {
	namespace := &v1.Namespace{}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Error(err)
		return namespace, err
	}
	if data, err = yaml.ToJSON(data); err != nil {
		logrus.Error(err)
		return namespace, err
	}
	if err = json.Unmarshal(data, namespace); err != nil {
		logrus.Error(err)
		return namespace, err
	}
	namespace, err = k.ClientSet.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{})
	if err != nil {
		logrus.Error(err)
		return namespace, err
	}
	return namespace, nil
}

/**
 * @Function: queryNamespaces
 * @Description: 查询全新的命名空间
 * @receiver k
 * @param ctx
 * @return *v1.Namespace
 * @return error
 */
func (k *KubernetesConfig) QueryNamespaces(ctx context.Context) (*v1.NamespaceList, error) {
	namespaceList, err := k.ClientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		logrus.Error(err)
		return namespaceList, err
	}
	return namespaceList, nil
}

/**
 * @Function: CreateDeployment
 * @Description: 创建全新的Deployment
 * @receiver k
 * @param ctx
 * @param fileName	yaml文件地址
 * @return *v1.Namespace
 * @return error
 */
func (k *KubernetesConfig) CreateDeployment(ctx context.Context, fileName string) (*dv1.Deployment, error) {
	deployment := &dv1.Deployment{}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Error(err)
		return deployment, err
	}
	if data, err = yaml.ToJSON(data); err != nil {
		logrus.Error(err)
		return deployment, err
	}
	if err = json.Unmarshal(data, deployment); err != nil {
		logrus.Error(err)
		return deployment, err
	}
	deployment, err = k.ClientSet.AppsV1().Deployments(k.Namespace).Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		logrus.Error(err)
		return deployment, err
	}
	return deployment, nil
}

/**
 * @Function: QueryDeployment
 * @Description: 查询pod
 * @receiver k
 * @param ctx
 * @return *dv1.DeploymentList
 * @return error
 */
func (k *KubernetesConfig) QueryDeployment(ctx context.Context) (*dv1.DeploymentList, error) {
	deploymentList, err := k.ClientSet.AppsV1().Deployments(k.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		logrus.Error(err)
		return deploymentList, err
	}
	return deploymentList, nil
}

/**
 * @Function: CreateServices
 * @Description: 创建全新的Services
 * @receiver k
 * @param ctx
 * @param fileName	yaml文件地址
 * @return *v1.Namespace
 * @return error
 */
func (k *KubernetesConfig) CreateServices(ctx context.Context, fileName string) (*v1.Service, error) {
	service := &v1.Service{}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Error(err)
		return service, err
	}
	if data, err = yaml.ToJSON(data); err != nil {
		logrus.Error(err)
		return service, err
	}
	if err = json.Unmarshal(data, service); err != nil {
		logrus.Error(err)
		return service, err
	}
	service, err = k.ClientSet.CoreV1().Services(k.Namespace).Create(ctx, service, metav1.CreateOptions{})
	if err != nil {
		logrus.Error(err)
		return service, err
	}
	return service, nil
}

/**
 * @Function: QueryServices
 * @Description: 查询Services
 * @receiver k
 * @param ctx
 * @return *v1.ServiceList
 * @return error
 */
func (k *KubernetesConfig) QueryServices(ctx context.Context) (*v1.ServiceList, error) {
	serviceList, err := k.ClientSet.CoreV1().Services(k.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		logrus.Error(err)
		return serviceList, err
	}
	return serviceList, nil
}

/**
 * @Function: CreateIngress
 * @Description: 创建Ingress
 * @receiver k
 * @param ctx
 * @param fileName yaml文件地址
 * @return *v1.Service
 * @return error
 */
func (k *KubernetesConfig) CreateIngress(ctx context.Context, fileName string) (*v1.Service, error) {
	service := &v1.Service{}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Error(err)
		return service, err
	}
	if data, err = yaml.ToJSON(data); err != nil {
		logrus.Error(err)
		return service, err
	}
	if err = json.Unmarshal(data, service); err != nil {
		logrus.Error(err)
		return service, err
	}
	service, err = k.ClientSet.CoreV1().Services(k.Namespace).Create(ctx, service, metav1.CreateOptions{})
	if err != nil {
		logrus.Error(err)
		return service, err
	}
	return service, nil
}
