#!/bin/bash

set -e
# i.e.
# PS /home/michal/workspace/onho.io> ./scripts/build.sh cid v0.1
usage(){
        cat <<EOF
        Usage: $(basename "$0") <COMMAND>  <TAG>
        Commands:
            ci                run build process with new version and properly tag
            cd                deploy app to container registry, redeploy k8s, install certificates
            cicd               ci+cd

        Command arguments:
            ci
                <TAG> required   docker tag, if empty than :latest is used

            cd
                <TAG> required   docker tag, if empty than :latest is used

            cicd
                <TAG> required   docker tag, if empty than :latest is used
EOF
}


panic() {
  (>&2 echo "$@")
  exit 1
}

check_kube_cli(){
	KUBECTL=`which kubectl`||true

	if [[ -z "${KUBECTL}" ]]; then
 		panic "Kubectl is not installed"
		exit 1
	fi
}


ci(){
cat <<EOF
***************************************************************
    building docker image
***************************************************************
EOF

    docker build . -t acronhosbx.azurecr.io/mutating-webhook:"${tag}" --no-cache

    docker push acronhosbx.azurecr.io/mutating-webhook:"${tag}"

}



cd(){
cat <<EOF
***************************************************************
    deploying mutating-webhook
***************************************************************
EOF
    check_kube_cli

    caBundle="$(kubectl config view --raw --minify --flatten -o jsonpath='{.clusters[].cluster.certificate-authority-data}')"
    yaml="$(cat ./deployment/deployment.yaml )"
    caBundle="$(kubectl config view --raw --minify --flatten -o jsonpath='{.clusters[].cluster.certificate-authority-data}')"
    buildVersion="$tag"_$(date '+%y%m.%d.%H%M')
    bind="$(echo "$yaml" | sed -e "s|\${tag}|${tag}|g" | sed -e "s|\${caBundle}|${caBundle}|g" | sed -e "s|\${buildVersion}|${buildVersion}|g")"

cat <<EOF | kubectl apply -f -
`echo "$bind"`
EOF
}



if [[ "$#" -lt 2 ]]; then
  usage
  exit 1
fi
tag=${2}

case "$1" in
    "ci")
       ci
    ;;
    "cd")
      cd
    ;;
    "cicd")
        ci
        cd
    ;;
      *)
  usage
  exit 0
  ;;
esac





