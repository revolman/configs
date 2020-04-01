# Dynamically provision NFS persistent volumes in Kubernetes
## Описание:

Для того, что бы не создавать Persistent Volume вручную в PVC указывается StorageClass, в соответствии с которыми PV будет создаваться автоматически.
У облачных провайдеров с этим проще, для создания SC нудно указать какой ресурс будет предоставлять дистовое пространство. Пример в AWS:
```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: standard
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
reclaimPolicy: Retain
allowVolumeExpansion: true
mountOptions:
  - debug
volumeBindingMode: Immediate
```

Если у тебя bare metal кластер, то нужно сначала создать privisioner.
NFS Client Provisioner - это специальный Pod, который будет монтировать удалённое NFS хранилище.

### NFS Server

В первую очередь нужно создать NFS server с достаточным дисковым пространством. Все данные будут отправляться на него.
####Мануал по установки NFS-Server
Centos 7:
https://wiki.it-kb.ru/unix-linux/centos/linux-how-to-setup-nfs-server-with-share-and-nfs-client-in-centos-7-2
Ubuntu:
https://vitux.com/install-nfs-server-and-client-on-ubuntu/

### NFS Provisioner

#### Установить:
```
ServiceAccount
Role
RoleBinding
ClusterRole
ClusterRoleBindings
StorageClass
Deployment
```
#### yamls/rbac.yaml установит:
ServiceAccount через который будет работать provisioner;
ClusterRole - роль с необходимыми правами;
ClusterRoleBinding - связка ServiceAccount и ClusterRole;
Role - роль для работы с эндпоинтами;
RoleBinding - связка Role и ClusterRole.

#### yamls/class.yaml установит:
StorageClass с именем managed-nfs-storage, в дальнейшем это имя будет указываться в директиве storageClass в PVC. Так же здесь указывается что нужно делать с данными после сметри PVC (удалить | архивировать).

#### yamls/deployment.yaml установит:
Pod nfs-client-provisioner, в котором указываются реквизиты подключения к NFS Server. Необходимо исправить на актуальные.

#### yamls/default-sc.yaml
Пример установки аннотации для указания SC по умолчанию.


### Пример использования

```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
	name: pvc1
spec:
	storageClassName: managed-nfs-storage
	accessModes:
		- ReadWriteMany
	resources:
		requests:
			storage: 500Mi
```