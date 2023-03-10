```bash
docker build -t liveness:v2 .
docker run -p 8081:8081 liveness:v2
docker save liveness:v2 -o livev2.tar
docker rmi  liveness:v2
docker load -i livev2.tar
kubectl apply -f http-liveness.yaml
kubectl delete -f http-liveness.yaml
```

测试表明，不需要在 Dockerfile 暴露端口，k8s 的 readinessProbe、livenessProbe 也是可以正常工作的

测试表明，kind 中也能够进行存活探针检测（需要把 yaml 中的 nodeName 注销掉）

# 参考

https://github.com/kubernetes/kubernetes/blob/master/test/images/agnhost/liveness/server.go

https://kubernetes.io/zh-cn/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/

[如何编写dockerfile来部署go项目     小饶coding      2021年12月01](https://juejin.cn/post/7036670636978077703)

> ```
> Dockfile 里面的注释不对
> docker build . -t my_docker_test
> docker run -p 8888:8888 my_docker_test 
> ```

[golang http.ListenAndServe 阻塞导致if else不执行问题分析 2019-11-29 ](http://blog.yoqi.me/lyq/16889.html)

## Error

```bash
[mifen@hp go-docker]$ docker run -p 8888:8888 my_docker_test 
docker: Error response from daemon: driver failed programming external connectivity on endpoint confident_shtern (3f0d6c05853b0d6f5cf0985d9815bd5984e8cb19141db194b0404b4b764e3f44):  (COMMAND_FAILED: '/usr/bin/iptables -w10 -t nat -A DOCKER -p tcp -d 0/0 --dport 8888 -j DNAT --to-destination 172.17.0.2:8888 ! -i docker0' failed: Warning: Extension DNAT revision 0 not supported, missing kernel module?
iptables v1.8.9 (nf_tables):  RULE_APPEND failed (No such file or directory): rule in chain DOCKER
).
ERRO[0000] error waiting for container: context canceled 
```

原来是端口冲突，8888已经被占用了
