# pods
Kubernetes course


Resources for kuernetes course on Windows with docker


COMANDS

minikube start --vm-driver=none
minikube stop

minikube stop & REM stops the VM

minikube delete & REM deleted the VM

minikube start --driver=docker


video 2 	18:32

POD KUBERNETES: Compartir namespaces entre contenedores
				Wrapper uno o mas contenedores que comparten namespaces entre si
					- namespace de red
					- comparte IPC entre contenedores (los contenedores pueden ver procesos entre si)
					- comparte UTS (hostname)
CREAR POD: kubectl run --generator=run-pod/v1 podtest --image=nginx:alpine

kubectl run test-nginx --image=nginx --port=80
kubectl run --generator=deployment/apps.v1 is DEPRECATED and will be removed in a future version. 
Use kubectl run --generator=run-pod/v1 or kubectl create instead.
deployment.apps/test-nginx created


kubectl run --generator=run-pod/v1 test --image=nginx --port=80 --env="DOMAIN=cluster"

  # Start a single instance of nginx.
  kubectl run nginx --image=nginx

  # Start a single instance of hazelcast and let the container expose port 5701 .
  kubectl run hazelcast --image=hazelcast --port=5701

  # Start a single instance of hazelcast and set environment variables "DNS_DOMAIN=cluster" and "POD_NAMESPACE=default" in the container.
  kubectl run hazelcast --image=hazelcast --env="DNS_DOMAIN=cluster" --env="POD_NAMESPACE=default"

  # Start a single instance of hazelcast and set labels "app=hazelcast" and "env=prod" in the container.
  kubectl run hazelcast --image=hazelcast --labels="app=hazelcast,env=prod"

  # Start a replicated instance of nginx.
  kubectl run nginx --image=nginx --replicas=5

  # Dry run. Print the corresponding API objects without creating them.
  kubectl run nginx --image=nginx --dry-run

  # Start a single instance of nginx, but overload the spec of the deployment with a partial set of values parsed from JSON.
  kubectl run nginx --image=nginx --overrides='{ "apiVersion": "v1", "spec": { ... } }'

  # Start a pod of busybox and keep it in the foreground, don't restart it if it exits.
  kubectl run -i -t busybox --image=busybox --restart=Never

  # Start the nginx container using the default command, but use custom arguments (arg1 .. argN) for that command.
  kubectl run nginx --image=nginx -- <arg1> <arg2> ... <argN>

  # Start the nginx container using a different command and custom arguments.
  kubectl run nginx --image=nginx --command -- <cmd> <arg1> ... <argN>

  # Start the perl container to compute π to 2000 places and print it out.
  kubectl run pi --image=perl --restart=OnFailure -- perl -Mbignum=bpi -wle 'print bpi(2000)'

  # Start the cron job to compute π to 2000 places and print it out every 5 minutes.
  kubectl run pi --schedule="0/5 * * * ?" --image=perl --restart=OnFailure -- perl -Mbignum=bpi -wle 'print bpi(2000)'
  
  CREAR POD TEMPORAL INTERACTIVO: kubectl run --rm -ti --generator=run-pod/v1 podtest3 --image=nginx:alpine -- sh
  

CREAR POD DESDE YAML: kubectl apply -f test.yml  
VER PODS: kubectl get pods o kube get pod NOMBRE
VER LOGS DE POD: kubectl describe pod podtest
ELIMINAR POD: kubectl delete pod podtest-fail O kubectl delete -f test.yml
OBTENER YAML DE DESDE UN POD:kubectl get pod podtest -o yaml
VER POD DESDE 127.0.0.1 APUNTANDO A OTRO PUERTO: kubectl port-forward NOMBRE 7000:80
												 luego 127.0.0.1:7000
												 
												 minikube service web --url
												 https://stackoverflow.com/questions/48857092/how-to-expose-nginx-on-public-ip-using-nodeport-service-in-kubernetes
												 https://stackoverflow.com/questions/48857092/how-to-expose-nginx-on-public-ip-using-nodeport-service-in-kubernetes
VER LOGS DE POD: kubectl logs NOMBRE -f
FILTRAR PODS POR LABEL APP: kubectl get pods -l app=front
PONER LABEL A POD YA CREADO: kubectl label pods podtest2 app=pod-label
MOSTRAR DEPLOYMENT: kubectl get deployment --show-labels
MOSTRAR ROLLOUT STATUS DEPLOYMENT: kubectl rollout status deployment NOMBRE
VER HISTORIAL DE ROLLOUT DE DEPLOYMENT:kubectl rollout history deployment deployment-test
CREAR/UPDATE DEPLOYMENT CON CHANGE CAUSE: kubectl apply -f deployment.yml --record
										  con kubernetes.io/change-cause: "Changes port to 110" en YML
VOLVER A UNA VERSION ANTERIOR DE DEPLOYMENT ROLLBACK:kubectl rollout undo deployment NOMBRE --to-revision=3
LEVANTAR DASHBOARD: kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.1/aio/deploy/recommended.yaml
					kubectl proxy
OBTENER TOKEN DESPUES DE CREAR SERVICE ACCOUNT/ROLE: kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | sls admin-user | ForEach-Object { $_ -Split '\s+' } | Select -First 1)										  




 					