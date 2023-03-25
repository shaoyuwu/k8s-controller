
#generate deploy yml
kubectl create deployment ingress-manager --image nginx:1.16 --dry-run=client -oyaml > deployment.yml

#generate sa yml
kubectl create sa ingress-manager --dry-run=client -oyaml > ingress-manager-sa.yaml

#generate role yml
kubectl create clusterrole ingress-manager-role --resource=ingress,service --verb list,watch,update,delete --dry-run=client -oyaml > ingress-manager-role.yml

#generate rolebinding yml
kubectl create clusterrolebinding ingress-manage-rb --role ingress-manager-role --serviceaccount default:ingress-manager-sa --dry-run=client  -oyaml > ingress-manager-rb.yml