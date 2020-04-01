# Prometheus
## Описание:

Prometheus - система мониторинга и time series db.
После установки Prometheus указываются селекоторы, в соответствии с которыми система будет знать от каких ресурсов получать метрики. В метадате ресурса так же должен быть указан этот селектор.
Prometheus работает через HTTP запросы к API куберенетиса.
Некоторые приложения (такие как redis) не умеют самостоятельно генерировать метрики. В этом случае в Pod добавляется sidecar-контейнер, который будет отвечать за генерацию метрик.

Для корректной работы нужно создать следующие сущности:
```
Secret
ConfigMaps
ServiceAccount
PVC - подразумевается использование StorageClass (в моём случае NFS-provisioner)
Deployment
ReplicaSets
DaemonSets
RoleBinding
```

## Установка

Далее описывается процесс установки Prometheus.
Нужен StorageClass, в этом репозиторие есть пример настройки NFS-provisioner.
В данной инструкции предполагается установка SC в качестве default-SC, для этого используется аннотация.

### Helm
Prometheus будет развёрнут в namespace: Prometheus на nodePort: 32322 и иметь два PersistentVolume.

#### Выбор чарта и настройка
```helm search repo prometheus```

Я выбрал stable/prometheus, теперь нужно настроить values, для этого их нужно загрузить на диск:
```
helm inspect values stable/prometheus > /tmp/prometheus.values
vim /tmp/prometheus.values
```
Нужно изменить type: ClusterIP на type: NodePort, и указать порт 32322.

Установка:
```
{
kubectl create namespace prometheus
helm install prometheus stable/prometheus --values /tmp/prometheus.values --namespace prometheus
helm ls -n prometheus
}
```
#### Подводные камни
- StorageClass при такой установке должен быть указан по default дял всего кластера, либо нужно настраивать values
- NFS client должен быть установлен на всех рабочих нодах (см. Dynamic volume provisioning)

## Конфигурирование
https://habr.com/ru/company/flant/blog/353410/

В Prometheus есть config и rule files.
Config имеет три секции:
- scrape_configs - настройка поиска целей, имя job'ы будет использовано как имя таблицы с метриками в TSDB
- rule_files - список директорий с правилами
- alerting - настройка алерт менеджеров

Пример секции scrape_config с комментариями:
```
scrape_configs:
  # Общие настройки
- job_name: kube-prometheus/custom/0    # просто название scrape job'а
                                        # показывается в разделе Service Discovery
  scrape_interval: 30s                  # как часто собирать данные
  scrape_timeout: 10s                   # таймаут на запрос
  metrics_path: /metrics                # path, который запрашивать
  scheme: http                          # http или https

  # Настройки Service Discovery
  kubernetes_sd_configs:                # означает, что targets мы получаем из Kubernetes
  - api_server: null                    # использовать адрес API-сервера из переменных
                                        # окружения (которые есть в каждом поде)
    role: endpoints                     # targets брать из endpoints
    namespaces:
      names:                            # искать endpoints только в этих namespaces
      - foo
      - baz

  # Настройки "фильтрации" (какие enpoints брать, какие — нет) и "релейблинга"
  # (какие лейблы добавить или удалить — для всех получаемых метрик)
  relabel_configs:
  # Фильтр по значению лейбла prometheus_custom_target,
  # полученного из service, связанного с endpoint
  - source_labels: [__meta_kubernetes_service_label_prometheus_custom_target]
    regex: .+                           # подходит любой НЕ пустой лейбл
    action: keep
  # Фильтр по имени порта
  - source_labels: [__meta_kubernetes_endpoint_port_name]
    regex: http-metrics                 # подходит, если порт называется http-metrics
    action: keep
  # Добавляем лейбл job, используем значение лейбла prometheus_custom_target
  # у service, к которому добавляем префикс "custom-"
  #
  # Лейбл job — служебный в Prometheus. Он определяет название группы,
  # в которой будет показываться target на странице targets, а также он будет
  # у каждой метрики, полученной у этих targets (чтобы можно было удобно
  # фильтровать в rules и dashboards)
  - source_labels: [__meta_kubernetes_service_label_prometheus_custom_target]
    regex: (.*)
    target_label: job
    replacement: custom-$1
    action: replace
  # Добавляем лейбл namespace
  - source_labels: [__meta_kubernetes_namespace]
    regex: (.*)
    target_label: namespace
    replacement: $1
    action: replace
  # Добавляем лейбл service
  - source_labels: [__meta_kubernetes_service_name]
    regex: (.*)
    target_label: service
    replacement: $1
    action: replace
  # Добавляем лейбл instance (в нём будет имя пода)
  - source_labels: [__meta_kubernetes_pod_name]
    regex: (.*)
    target_label: instance
    replacement: $1
    action: replace
```



# Grafana

#### Описание
Рисковалка графиков. Берёт метрики из различных источников и позволяет создавать графики.


## Установка

### Helm
#### Выбор чарта и настройка

Чарт графаны на момент написания один.
```helm search repo grafana```

```
helm inspect values stable/grafana > /tmp/grafana.values
vim /tmp/grafana.values
```
Для корректорной работы нужно изменить persistence:
```
persistence:
  type: pvc
  enabled: true
  # storageClassName: default
  accessModes:
    - ReadWriteOnce
  size: 10Gi
  # annotations: {}
  finalizers:
    - kubernetes.io/pvc-protection
  # subPath: ""
  # existingClaim:

```
При желании можно установить пароль админа. Если он не установлен, то можно будет достать его из лога
```kubectl get secret --namespace grafana grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo```

Далее установить чарт:
```
{
kubectl create namespace prometheus
helm install prometheus stable/prometheus --values /tmp/prometheus.values --namespace prometheus
helm ls -n prometheus
}
```
