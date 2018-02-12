# extend_k8s_example

> 这个例子演示如何扩展k8s apiserver，这里我们设计一个NginxApp资源，代表一个nginx app

## 工具
生成符合k8s开发规范的代码，比如说client go对象要`runtime.Object`对象，要有deepcopy方法.

现在k8s社区提供的生产代码工具
- deepcopy-gen: creates a method func (t* T) DeepCopy() *T for each type T
- client-gen: creates typed clientsets for CustomResource APIGroups
- informer-gen: creates informers for CustomResources which offer an event based interface to react on changes of CustomResources on the server
- lister-gen: creates listers for CustomResources which offer a read-only caching layer for GET and LIST requests.

后面两个工具是编写Controller(或者说operator)的基础。使用这四个工具，我们就能编写成熟的和k8s风格一致的contoller。
这几个工具都是基于  [k8s.io/gengo](https://github.com/kubernetes/gengo)

除了这些工具[apiserver-builder](https://github.com/kubernetes-incubator/apiserver-builder)更进了一步，把生成controller的相关逻辑也自动化了。

具体的创建逻辑参考[这里](https://github.com/kubernetes-incubator/apiserver-builder/blob/master/docs/concepts/api_building_overview.md)本质还是CRD+operator/controller的逻辑.

可以说k8s社区提供的四个代码生成工具以及[kubernetes/code-generator](https://github.com/kubernetes/code-generator)和 [apiserver-builder](https://github.com/kubernetes-incubator/apiserver-builder)是帮你开发extend api server的三个层次的工具，前面的更底层，后面的更高层。


## 操作步骤

### 1. init 

这里我们使用[apiserver-builder](https://github.com/kubernetes-incubator/apiserver-builder/blob/master/docs/installing.md)作为模版生成工具.

```
apiserver-boot init repo --domain example.com


// 可以看到已经生成了大量模版代码，包括vendor库
➜  extend_k8s_example git:(master) ✗ tree -I vendor
.
├── BUILD.bazel
├── Gopkg.lock
├── Gopkg.toml
├── LICENSE
├── README.md
├── WORKSPACE
├── bin
├── boilerplate.go.txt
├── cmd
│   ├── apiserver
│   │   └── main.go
│   └── controller-manager
│       └── main.go
└── pkg
    ├── apis
    │   └── doc.go
    ├── controller
    │   ├── doc.go
    │   └── sharedinformers
    │       └── doc.go
    ├── doc.go
    └── openapi
        └── doc.go
```

### 2. 创建资源

创建NginxApp资源, （这里有个小bug，使用nginx作为资源名会出现，而使用nginxapp不会出现，原因和nginx复数名为nignxs还是nginxes内部出现了不一致有关）

```
apiserver-boot create group version resource --group simple --version v1alpha1  --kind NginxApp

// 观察资源，这个命令把NginxApp相关的资源模版都创建出来了, 比较重要的内容在
➜  extend_k8s_example git:(master) ✗ tree -I vendor
.
├── BUILD.bazel
├── Gopkg.lock
├── Gopkg.toml
├── LICENSE
├── README.md
├── WORKSPACE
├── bin
├── boilerplate.go.txt
├── cmd
│   ├── apiserver
│   │   └── main.go
│   └── controller-manager
│       └── main.go
├── docs
│   └── examples
│       └── nginxapp
│           └── nginxapp.yaml
├── pkg
│   ├── apis
│   │   ├── doc.go
│   │   └── simple
│   │       ├── doc.go
│   │       ├── install
│   │       │   └── doc.go
│   │       └── v1alpha1
│   │           ├── doc.go
│   │           ├── nginxapp_ptypes.go
│   │           ├── nginxapp_types_test.go
│   │           └── v1alpha1_suite_test.go
│   ├── controller
│   │   ├── doc.go
│   │   ├── nginxapp
│   │   │   ├── controller.go
│   │   │   ├── controller_test.go
│   │   │   └── nginxapp_suite_test.go
│   │   └── sharedinformers
│   │       ├── doc.go
│   │       └── informers.go
│   ├── doc.go
│   └── openapi
│       └── doc.go
└── sample
    └── nginxapp.yaml


// 运行这个命令实际会调用 apiregister-gen 等代码生成工具，生成对应的符合runtimeobject协议的代码
➜  extend_k8s_example git:(master) ✗ apiserver-boot build generated

```

### 3. 创建apiserver,controller manager和进行本地测试

```
// 这是使用bazel进行构建，当然也可以直接使用go build进行构建
// 注：bazel是google出品的构建工具，k8s都是采用bazel构建，对于大型项目，使用bazel能够显著加速构建速度
➜  extend_k8s_example git:(master) ✗ apiserver-boot build executables --bazel --gazelle --generate=false

// 不使用bazel进行构建的例子
➜  extend_k8s_example git:(master) ✗ apiserver-boot build executables  --generate=false

// 构建完成之后在bin目录下会发现构建好的二进制文件
// 尝试本地运行二进制，需要有etcd组件
➜  extend_k8s_example git:(master) ✗ apiserver-boot run local --build=false

// 这时候在根目录下面出现kubeconfig文件，使用这个kubeconfig运行这个例子
➜  extend_k8s_example git:(master) ✗ kubectl --kubeconfig kubeconfig api-versions
simple.example.com/v1alpha1

➜  extend_k8s_example git:(master) ✗ kubectl --kubeconfig kubeconfig create -f sample/nginxapp.yaml
nginxapp "nginxapp-example" created


➜  extend_k8s_example git:(master) ✗ kubectl --kubeconfig kubeconfig get nginxapp
NAME               AGE

```


### 4. 填充更多NginxSpec和controller的代码
添加了valid代码，nginx的副本数量不能大于10

```
// 修改yaml文件使得,replia 数量大于10
➜  extend_k8s_example git:(master) ✗ kubectl create -f sample/nginxapp.yaml
The NginxApp "nginxapp-example" is invalid: spec: Invalid value: 12: replicas must less than 10
```

添加了controller代码，当创建nginxapp资源的时候，创建一个deployment出来

### 5. 部署到cluster

下面的例子可以看出，我们已经实现了这样的功能，创建nginxapp之后, controller自动创建一个deloyment。下面我们还可以继续优化这个controller, 达到这样的目的：创建一个nginxapp, 我没创建好deloyment 和service，把这个app暴露出来.
```
➜  extend_k8s_example git:(master) ✗ apiserver-boot run in-cluster --name nginxcontroller --namespace default --image=hub.c.163.com/u2takey/tool:nginxextend2  --namespace=kube-system

➜  extend_k8s_example git:(master) ✗ kubectl create -f sample/nginxapp.yaml
nginxapp "nginxapp-example-2" created
➜  extend_k8s_example git:(master) ✗ kubectl get nginxapp
NAME                 AGE
nginxapp-example-2   10s



➜  extend_k8s_example git:(master) ✗ kubectl create -f sample/nginxapp.yaml
nginxapp "nginxapp-example" created
➜  extend_k8s_example git:(master) ✗ kubectl get nginxapp -n demo
NAME               AGE
nginxapp-example   26s
➜  extend_k8s_example git:(master) ✗ kubectl get deploy -n demo
NAME               DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
nginxapp-example   2         2         2            2           33s
➜  extend_k8s_example git:(master) ✗ kubectl delete nginxapp nginxapp-example -n demo
nginxapp "nginxapp-example" deleted
➜  extend_k8s_example git:(master) ✗ kubectl get deploy -n demo
No resources found.
```















## 参考:
- [Kubernetes Deep Dive: Code Generation for CustomResources](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/)
- [apiserver-builder](https://github.com/kubernetes-incubator/apiserver-builder)

