#!/bin/bash

get_pod_name_by_label() {
    kubectl get po -l $1 -o custom-columns=NAME:.metadata.name | tail +2 | uniq
}

alias kpn='get_pod_name_by_label'
