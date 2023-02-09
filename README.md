```bash
docker build -t liveness:v1 .
docker run -p 8081:8081 liveness:v1
docker save liveness:v1 -o livev1.tar
docker load -i livev1.tar
kubectl apply -f http-liveness.yaml
```

# 参考

https://github.com/kubernetes/kubernetes/blob/master/test/images/agnhost/liveness/server.go

https://kubernetes.io/zh-cn/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/

[如何编写dockerfile来部署go项目     小饶coding      2021年12月01](https://juejin.cn/post/7036670636978077703)

> ```
> Dockfile 里面的注释不对
> docker build . -t my_docker_test
> docker run -p 8888:8888 my_docker_test 
> ```

## Error

```bash
[mifen@hp go-docker]$ docker run -p 8888:8888 my_docker_test 
docker: Error response from daemon: driver failed programming external connectivity on endpoint confident_shtern (3f0d6c05853b0d6f5cf0985d9815bd5984e8cb19141db194b0404b4b764e3f44):  (COMMAND_FAILED: '/usr/bin/iptables -w10 -t nat -A DOCKER -p tcp -d 0/0 --dport 8888 -j DNAT --to-destination 172.17.0.2:8888 ! -i docker0' failed: Warning: Extension DNAT revision 0 not supported, missing kernel module?
iptables v1.8.9 (nf_tables):  RULE_APPEND failed (No such file or directory): rule in chain DOCKER
).
ERRO[0000] error waiting for container: context canceled 
```

原来是端口冲突，8888已经被占用了
