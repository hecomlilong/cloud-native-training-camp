#!/bin/bash
INGRESS_IP=$(sudo kubectl get svc istio-ingressgateway -n istio-system -o=jsonpath='{.spec.clusterIP}')
for((i=1;i<=100;i++));
do
    curl https://$INGRESS_IP/service1/hello -k;
done
