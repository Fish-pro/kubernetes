## kubectl

kubectl controls the Kubernetes cluster manager

### Synopsis

kubectl controls the Kubernetes cluster manager.

 Find more information at: https://kubernetes.io/docs/reference/kubectl/

```
kubectl [flags]
```

### Options

```
      --as string                      Username to impersonate for the operation. User could be a regular user or a service account in a namespace.
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --as-uid string                  UID to impersonate for the operation.
      --cache-dir string               Default cache directory (default "/Users/york/.kube/cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --context string                 The name of the kubeconfig context to use
  -h, --help                           help for kubectl
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
      --match-server-version           Require server version to match client version
  -n, --namespace string               If present, the namespace scope for this CLI request
      --password string                Password for basic authentication to the API server
      --profile string                 Name of profile to capture. One of (none|cpu|heap|goroutine|threadcreate|block|mutex) (default "none")
      --profile-output string          Name of the file to write the profile to (default "profile.pprof")
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                  The address and port of the Kubernetes API server
      --tls-server-name string         Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used
      --token string                   Bearer token for authentication to the API server
      --user string                    The name of the kubeconfig user to use
      --username string                Username for basic authentication to the API server
      --warnings-as-errors             Treat warnings received from the server as errors and exit with a non-zero exit code
```

### SEE ALSO

* [kubectl alpha](kubectl_alpha.md)	 - Commands for features in alpha
* [kubectl annotate](kubectl_annotate.md)	 - Update the annotations on a resource
* [kubectl api-resources](kubectl_api-resources.md)	 - Print the supported API resources on the server
* [kubectl api-versions](kubectl_api-versions.md)	 - Print the supported API versions on the server, in the form of "group/version"
* [kubectl apply](kubectl_apply.md)	 - Apply a configuration to a resource by file name or stdin
* [kubectl attach](kubectl_attach.md)	 - Attach to a running container
* [kubectl auth](kubectl_auth.md)	 - Inspect authorization
* [kubectl autoscale](kubectl_autoscale.md)	 - Auto-scale a deployment, replica set, stateful set, or replication controller
* [kubectl certificate](kubectl_certificate.md)	 - Modify certificate resources.
* [kubectl cluster-info](kubectl_cluster-info.md)	 - Display cluster information
* [kubectl completion](kubectl_completion.md)	 - Output shell completion code for the specified shell (bash, zsh, fish, or powershell)
* [kubectl config](kubectl_config.md)	 - Modify kubeconfig files
* [kubectl cordon](kubectl_cordon.md)	 - Mark node as unschedulable
* [kubectl cp](kubectl_cp.md)	 - Copy files and directories to and from containers
* [kubectl create](kubectl_create.md)	 - Create a resource from a file or from stdin
* [kubectl debug](kubectl_debug.md)	 - Create debugging sessions for troubleshooting workloads and nodes
* [kubectl delete](kubectl_delete.md)	 - Delete resources by file names, stdin, resources and names, or by resources and label selector
* [kubectl describe](kubectl_describe.md)	 - Show details of a specific resource or group of resources
* [kubectl diff](kubectl_diff.md)	 - Diff the live version against a would-be applied version
* [kubectl drain](kubectl_drain.md)	 - Drain node in preparation for maintenance
* [kubectl edit](kubectl_edit.md)	 - Edit a resource on the server
* [kubectl exec](kubectl_exec.md)	 - Execute a command in a container
* [kubectl explain](kubectl_explain.md)	 - Get documentation for a resource
* [kubectl expose](kubectl_expose.md)	 - Take a replication controller, service, deployment or pod and expose it as a new Kubernetes service
* [kubectl get](kubectl_get.md)	 - Display one or many resources
* [kubectl kustomize](kubectl_kustomize.md)	 - Build a kustomization target from a directory or URL.
* [kubectl label](kubectl_label.md)	 - Update the labels on a resource
* [kubectl logs](kubectl_logs.md)	 - Print the logs for a container in a pod
* [kubectl options](kubectl_options.md)	 - Print the list of flags inherited by all commands
* [kubectl patch](kubectl_patch.md)	 - Update fields of a resource
* [kubectl plugin](kubectl_plugin.md)	 - Provides utilities for interacting with plugins
* [kubectl port-forward](kubectl_port-forward.md)	 - Forward one or more local ports to a pod
* [kubectl proxy](kubectl_proxy.md)	 - Run a proxy to the Kubernetes API server
* [kubectl replace](kubectl_replace.md)	 - Replace a resource by file name or stdin
* [kubectl rollout](kubectl_rollout.md)	 - Manage the rollout of a resource
* [kubectl run](kubectl_run.md)	 - Run a particular image on the cluster
* [kubectl scale](kubectl_scale.md)	 - Set a new size for a deployment, replica set, or replication controller
* [kubectl set](kubectl_set.md)	 - Set specific features on objects
* [kubectl taint](kubectl_taint.md)	 - Update the taints on one or more nodes
* [kubectl top](kubectl_top.md)	 - Display resource (CPU/memory) usage
* [kubectl uncordon](kubectl_uncordon.md)	 - Mark node as schedulable
* [kubectl version](kubectl_version.md)	 - Print the client and server version information
* [kubectl wait](kubectl_wait.md)	 - Experimental: Wait for a specific condition on one or many resources

###### Auto generated by spf13/cobra on 27-Jul-2022
