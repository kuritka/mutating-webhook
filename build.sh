#!/bin/bash

set -e
# i.e.
# PS /home/michal/workspace/onho.io> ./scripts/build.sh ci v0.1
usage(){
        cat <<EOF
        Usage: $(basename "$0") <COMMAND>  <TAG>
        Commands:
            ci                run build process with new version and properly tag
            cd                deploy app to container registry, redeploy k8s, install certificates
            cid               ci+cd

        Command arguments:
            ci
                <TAG> required   docker tag, if empty than :latest is used

            cd
                <TAG> required   docker tag, if empty than :latest is used

            cid
                <TAG> required   docker tag, if empty than :latest is used
EOF
}


panic() {
  (>&2 echo "$@")
  exit 1
}


dir_exists(){
	local path="$1"
    	if [[ ! -d "$path" ]]; then
  		panic "$path doesn't exists"
     	fi
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

    docker build . -t acronhosbx.azurecr.io/mutating-webhook:${tag} --no-cache

    docker push acronhosbx.azurecr.io/mutating-webhook:${tag}

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
    "cid")
        ci
        cd
    ;;
      *)
  usage
  exit 0
  ;;
esac





