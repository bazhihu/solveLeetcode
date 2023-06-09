# K8S 主要组件
- kubelet - 节点node的代理，非容器运行；功能：创建和运行容器，向Master报告运行状态
- kube-proxy - 帮助service 转发请求到 Pod, 功能：转发TCP/UPD数据流至后端容器和负载均衡
- Pod网络 - 使Pod 相互通信

# 运行应用的Controller模式
- Deployment - 通常模式
- DaemonSet - only one 模式 (日志收集、监控、kube-flannel-ds、kube-proxy)
- Job - 工作类容器，一次性任务运行完成后退出 
```  
# k8s 默认没有enable CronJob, 需要在kube-apiserver加入功能
# 修改kube-apiserver配置文件 /etc/kubernetes/manifests/kubeapiserver.yaml
# 在启动参数中加上 --runtime-config=batch/v2alpha1=true 
```

# 查看节点
kubectl get node
# 查看节点 展示标签labels
kubectl get node --show-lables
# 给节点赋值 labels
kubectl label node minikube disktype=ssd
# 给节点删除 label   (- 即删除)
kubectl label node minikube disktype-

# 获取节点上 DaemonSet(最多只能运行一个副本) 
# --namespace=kube-system 代表系统组件, 不指定 返回默认namespace default中的资源
kubectl get daemonset --namespace=kube-system

# 查看所有Pod
kubectl get pod --all-namespaces -o wide

# 运行一个容器
kubectl run httpd-app --image=httpd --replicas=2

# 通过配置文件和kubectl apply创建
kubectl apply -f nginx.yml
# 删除创建的资源
kubectl delete deployment nginx-deployment
kubectl delete -f nginx.yml

# 查看配置
kubectl edit daemonset kube-proxy --namespace=kube-system

# 查看单次运行的Job 状态
kubectl get jobs
kubectl get CronJob

# 查看Pod的标准输出
kubectl logs httpd-app

# 重启kubelet 服务
systemctl restart kubelet.service

# 查看所有的namespace
kubectl get namespace


# 清理资源
kubectl delete service hello-node
kubectl delete deployment hello-node


# 内部访问Pod 方式：通过<SERVICE_NAME>.<NAMESPACE_NAME> 访问Service
# 外部访问Service 方式：<NodeIP>:<NodePort> 访问Service



cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-\$basearch

enabled=1
gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
exclude=kubelet kubeadm kubectl
EOF

sudo setenforce 0
sudo sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config

sudo yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes

sudo systemctl enable --now kubelet