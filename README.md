# client_go
## 简介

client-go是一个调用kubernetes集群资源对象API的客户端，即通过client-go实现对kubernetes集群中资源对象（包括deployment、service、ingress、replicaSet、pod、namespace、node等）的增删改查等操作。大部分对kubernetes进行前置API封装的二次开发都通过client-go这个第三方包来实现。Kubernetes官方从2016年8月份开始，将Kubernetes资源操作相关的核心源码抽取出来，独立出来一个项目Client-go，作为官方提供的Go client。client-go支持RESTClient、ClientSet、DynamicClient、DiscoveryClient四种客户端与Kubernetes Api Server进行交互。

### package包的功能说明

- kubernetes: 访问kubernetes API的一系列的clientset
- discovery: 通过Kubernetes API进行服务发现
- dynamic： 对任意Kubernetes对象执行通用操作的动态client
- transport：启动连接和鉴权auth
- tools/cache： controllers控制器

### client四种客户端

#### RESTClient

RESTClient是最基础的，相当于最底层的基础结构，可以直接通过RESTClient提供的RESTful方法如Get()、Put()、Post()、Delete()进行交互。

- 同时支持json和protobuf
- 支持所有原生资源和CRDs
- 但是，一般而言，为了更为优雅的处理，需要进一步封装，通过Clientset封装RESTClient，然后再对外提供接口和服务。

作为最基础的客户端，其他的客户端都是基于RESTClient实现的。RESTClient对HTTP Request进行了封装，实现了RESTFul风格的API，具有很高的灵活性，数据不依赖于方法和资源，因此RESTClient能够处理多种类型的调用，返回不同的数据格式。

#### Clientset

Clientset是调用Kubernetes资源对象最常用的client，可以操作所有的资源对象，包含RESTClient。需要制定Group、Version，然后根据Resource获取

- 优雅的姿势是利用一个controller对象，再加上informer

Clientset是在RESTClient的基础上封装了对Resource和Version的管理方法。每一个Resource可以理解为一个客户端，而Clientset是多个客户端的集合，每一个Resource和Version都以函数的方式暴露出来。 Clientset仅能访问Kubernetes自身内置的资源，不能直接访问CRD自定义资源。如果要想Clientset访问CRD自定义资源，可以通过client-gin代码生成器重新生成CRD自定义资源的Clientset。

#### DynamicClient

Dynamic client是一种动态的client，它能处理kubernetes所有的资源。不同于clientset，dynamic client返回的对象是一个map[string]interface{},如果一个controller中需要控制所有的API，可以使用dynamic client，目前它在garbage collector和namespace controller中被使用。

- 只支持json

DynamicClient是一个动态客户端，可以对任意Kubernetes资源进行RESTFul操作，包括CRD自定义资源。DynamicClient与ClientSet操作类型，同样是封装了RESTClient。DynamicClient与ClientSet最大的不同就是：ClientSet仅能访问Kubernetes自带的资源，不能直接访问CRD自定义资源。Clientset需要预先实现每种Resource和Version的操作，其内部的数据都是结构化的。而DynamicClient内部实现了Unstructured，用于处理非结构化数据结构，这是能够处理CRD自定义资源的关键。DynamicClient不是类型安全的，因此访问CRD自定义资源时需要特别注意。

#### DiscoveryClient

DiscoveryClient是发现客户端，主要用于发现kubernetes API Server所支持的资源组、资源版本、资源信息。除此之外，还可以将这些信息存储到本地，用户本地缓存，以减轻对Kubernetes API Server访问的压力。 kubectl的api-versions和api-resources命令输出也是通过DisconversyClient实现的。

DiscoveryClient是发现客户端，主要用于发现Kubernetes API Server所支持的资源组、资源版本、资源信息。除此之外，还可以将这些信息存储到本地，用户本地缓存，以减轻Api Server访问压力。

kubectl的api-versions和api-resources命令输出也是通过DiscoverysyClient实现的。

