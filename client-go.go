package main

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	dv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	storagev1 "k8s.io/api/storage/v1"
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
	//获取config配置
	config, err := clientcmd.BuildConfigFromFlags("", "./config/config")
	if err != nil {
		panic("config 文件加载失败")
	}
	kubernetesConfig := NewKubernetesConfig(config, "testyaml")
	ctx := context.Background()
	//Namespaces
	//_, err = kubernetesConfig.CreateNamespaces(ctx, "./yaml/namespace.yaml")
	//namespacesList, err := kubernetesConfig.QueryNamespaces(ctx)
	//Deployment
	//_, err = kubernetesConfig.CreateDeployment(ctx, "./yaml/deployment.yaml")
	//deploymentList, err := kubernetesConfig.QueryDeployment(ctx)
	//err = kubernetesConfig.DeleteDeployment(ctx,"dataservice")
	//services
	//_, err = kubernetesConfig.CreateServices(ctx, "./yaml/services.yaml")
	//servicesList, err := kubernetesConfig.QueryServices(ctx)
	//lngress
	_, _ = kubernetesConfig.Createlngress(ctx, "./yaml/lngress/lngress.yaml")
	//storage
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println(namespacesList)
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
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if data, err = yaml.ToJSON(data); err != nil {
		logrus.Error(err)
		return nil, err
	}
	namespace := &v1.Namespace{}
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
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if data, err = yaml.ToJSON(data); err != nil {
		logrus.Error(err)
		return nil, err
	}
	deployment := &dv1.Deployment{}
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
 * @Function: DeleteDeployment
 * @Description: 删除pod
 * @receiver k
 * @param ctx
 * @return *dv1.DeploymentList
 * @return error
 */
func (k *KubernetesConfig) DeleteDeployment(ctx context.Context, name string) error {
	err := k.ClientSet.AppsV1().Deployments(k.Namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
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
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if data, err = yaml.ToJSON(data); err != nil {
		logrus.Error(err)
		return nil, err
	}
	service := &v1.Service{}
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
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if data, err = yaml.ToJSON(data); err != nil {
		logrus.Error(err)
		return nil, err
	}
	service := &v1.Service{}
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
 * @Function: CreateStorageClass
 * @Description: 创建StorageClass
 * @receiver k
 * @param ctx
 * @param fileName yaml文件地址
 * @return *v1.Service
 * @return error
 */
func (k *KubernetesConfig) CreateStorageClass(ctx context.Context, fileName string) (*storagev1.StorageClass, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if data, err = yaml.ToJSON(data); err != nil {
		logrus.Error(err)
		return nil, err
	}
	storageClass := &storagev1.StorageClass{}
	if err = json.Unmarshal(data, storageClass); err != nil {
		logrus.Error(err)
		return storageClass, err
	}
	storageClass, err = k.ClientSet.StorageV1().StorageClasses().Create(ctx, storageClass, metav1.CreateOptions{})
	if err != nil {
		logrus.Error(err)
		return storageClass, err
	}
	return storageClass, nil
}

/**
 * @Function: Createlngress
 * @Description: 创建lngress
 * @receiver k
 * @param ctx
 * @param fileName yaml文件地址
 * @return *v1.Service
 * @return error
 */
func (k *KubernetesConfig) Createlngress(ctx context.Context, fileName string) (*networkingv1.IngressClass, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if data, err = yaml.ToJSON(data); err != nil {
		logrus.Error(err)
		return nil, err
	}
	ingressClass := &networkingv1.IngressClass{}
	if err = json.Unmarshal(data, ingressClass); err != nil {
		logrus.Error(err)
		return ingressClass, err
	}
	ingressClass, err = k.ClientSet.NetworkingV1().IngressClasses().Create(ctx, ingressClass, metav1.CreateOptions{})
	if err != nil {
		logrus.Error(err)
		return ingressClass, err
	}
	return ingressClass, nil
}
