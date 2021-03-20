# pods
Kubernetes course


Resources for kuernetes course on Windows with docker
KUBELET: Servicio que corre en cada pod


COMANDS

minikube start --vm-driver=none
minikube stop

minikube stop & REM stops the VM

minikube delete & REM deleted the VM

minikube start --driver=docker


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
  
  
  CREAR POD TEMPORAL INTERACTIVO: kubectl run --rm -ti --generator=run-pod/v1 podtest3 --image=nginx:alpine -- sh
  POD INTERACTIVO DE UN POD EXISTENTE: kubectl exec -ti deployment-test-7fb785464c-hbdlv -- sh
  

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

========DOCKER==============================================
docker rm -fv c7afec7e0b67d9434559640b56bf17bf0aff92e8928084cd6b44f661056c77fe eliminar imagen
docker run --rm -dti -v /$PWD/:/go --net host --name golang golang bash
docker run --rm -dti -v "%cd%":/go --net host --name golang golang bash crear imagen y compartir carpeta a go dentro
docker exec -ti 0944d4859381854d20439bf22c4d5f51917d38de956a2ed1af3ca84c8fc6cada bash entrar dentro de la imagen
go run main.go

docker build -t k8s-hands-on -f Dockerfile 
docker build -t k8s-hands-on:v2 -f Dockerfile .
docker run -d(daemon) -p(puerto) 9091(maquina local):9090(contenedor --name k8s-hands-on(nombre) k8s-hands-on(imagen) 
 					
====PROBLEMA DE IMAGEN LOCAL =================

The command minikube docker-env returns a set of Bash environment variable exports to configure your local environment to re-use the Docker daemon inside the Minikube instance.

Passing this output through eval causes bash to evaluate these exports and put them into effect.

You can review the specific commands which will be executed in your shell by omitting the evaluation step and running minikube docker-env directly. However, this will not perform the configuration – the output needs to be evaluated for that.

This is a workflow optimization intended to improve your experience with building and running Docker images which you can run inside the minikube environment. It is not mandatory that you re-use minikube's Docker daemon to use minikube effectively, but doing so will significantly improve the speed of your code-build-test cycle.

In a normal workflow, you would have a separate Docker registry on your host machine to that in minikube, which necessitates the following process to build and run a Docker image inside minikube:

Build the Docker image on the host machine.
Re-tag the built image in your local machine's image registry with a remote registry or that of the minikube instance.
Push the image to the remote registry or minikube.
(If using a remote registry) Configure minikube with the appropriate permissions to pull images from the registry.
Set up your deployment in minikube to use the image.
By re-using the Docker registry inside Minikube, this becomes:

Build the Docker image using Minikube's Docker instance. This pushes the image to Minikube's Docker registry.
Set up your deployment in minikube to use the image.

PS C:\k8s\pods\go-backend-service> minikube docker-env                                                                  $Env:DOCKER_TLS_VERIFY = "1"
PS C:\k8s\pods\go-backend-service> minikube -p minikube docker-env | Invoke-Expression                                  PS C:\k8s\pods\go-backend-service> cd .\src\                                                                            PS C:\k8s\pods\go-backend-service\src> docker build -t k8s-hands-on -f Dockerfile .                                     Sending build context to Docker daemon  3.584kB
PS C:\k8s\pods\go-backend-service\src> docker build -t k8s-hands-on -f Dockerfile . 
Successfully tagged k8s-hands-on:latest
SECURITY WARNING: You are building a Docker image from Windows against a non-Windows Docker host. All files and directories added to build context will have '-rwxr-xr-x' permissions. It is recommended to double check and reset permissions for sensitive files and directories.
PS C:\k8s\pods\go-backend-service\src> kubectl apply -f .\backend.yml                                                   deployment.apps/backend-dep created
service/backend-svc created
PS C:\k8s\pods\go-backend-service\src> kubectl get pods                                                                 NAME                           READY   STATUS    RESTARTS   AGE
backend-dep-66ccdbbd64-k54zf   1/1     Running   0          12s
backend-dep-66ccdbbd64-lzvrh   1/1     Running   0          12s
backend-dep-66ccdbbd64-zvwqc   1/1     Running   0          12s					

IMAGEN FRONT: docker build -t front-k8s-hands-on:v1 -f .\Dockerfile .  

VER CONTEXTO ACTUAL: kubectl config current-context 
CREAR CONTEXTO: kubectl config set-context ci-context --namespace=ci --cluster=minikube --user=minikube  
USAR CONTEXTO: kubectl config use-context arn:aws:eks:us-west-2:676699563051:cluster/upic-adsales-eks
DESCRIBIR NAMESPACE: kubectl describe ns dev
LIMIT RANGE Y RESOURCE QUOTA: kubectl get limitranges
							  kubectl get resourcequota
PROBE:
	LIVENESS:	se ejecuta en el contenedor cada n intervalo de tiempo
				puede ser comando, http o tcp. por ejemplo puede evaluar un status code-build-test
	READINESS:	nos ayuda a garantizar que el servicio dentro del pod est
	START UP:	pausa liveness y readiness si está definido

CREAR CONFIG MAP:  kubectl create configmap nginx-config --from-file=examples (from file toma todos los archivos de la carpeta)
GET CONFIG MAPS: kubectl get cm
