package dns

import (
	"bytes"
	"html/template"
	"net"
)

func dnsDefault(clusterSubnet string) (string, error) {
	ip, _, err := net.ParseCIDR(clusterSubnet)
	if err != nil {
		return "", err
	}

	data := map[string]string{
		"CLUSTER_DNS_IP": string(ip[3] + 10),
	}

	const dnsDefault = `apiVersion: v1
kind: Service
metadata:
  name: dns-default
  namespace: openshift-dns
spec:
  clusterIP: "{{.CLUSTER_DNS_IP}}"
  ports:
  - name: dns
    port: 53
    protocol: UDP
    targetPort: dns
  - name: dns-tcp
    port: 53
    protocol: TCP
    targetPort: dns-tcp
  - name: metrics
    port: 9154
    protocol: TCP
    targetPort: metrics
  selector:
    dns.operator.openshift.io/daemonset-dns: default
  sessionAffinity: None
  type: ClusterIP`

	tmpl, err := template.New("dnsDefault").Parse(dnsDefault)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func Manifests(OpenshiftVersion string) (map[string]string, error) {
	manifests := make(map[string]string)
	manifests["99_openshift-dns_ns.yaml"] = dnsNamespace
	dnfDefault, err := dnsDefault(OpenshiftVersion)
	if err != nil {
		return map[string]string{}, err
	}
	manifests["99_openshift-dns_default.yaml"] = dnfDefault
	return manifests, nil
}

const dnsNamespace = `kind: Namespace
apiVersion: v1
metadata:
  annotations:
    openshift.io/node-selector: ""
  name: openshift-dns
  labels:
    # set value to avoid depending on kube admission that depends on openshift apis
    openshift.io/run-level: "0"
    # allow openshift-monitoring to look for ServiceMonitor objects in this namespace
    openshift.io/cluster-monitoring: "true"`
