import os
import subprocess
import sys


def run(cmd):
    cmd = ('argocd ' + cmd).split()
    output = subprocess.check_output(cmd)
    return output.decode('utf-8')


def login():
    host = 'localhost:8080'
    password = os.environ['ARGOCD_PASSWORD']
    cmd = f'login {host} --insecure --username admin --password {password}'
    output = run(cmd)
    print(output)


def get_apps(namespace):
    """ """
    output = run(f'app list -o wide')
    keys = 'name project namespace path repo'.split()
    apps = []
    lines = output.split('\n')
    headers = [h.lower() for h in lines[0].split()]
    for line in lines[1:]:
        items = line.split()
        app = {k: v for k, v in zip(headers, items) if k in keys}
        if app:
            apps.append(app)
    return apps


def create_project(project, cluster, namespace, description, repo):
    """ """
    cmd = f'proj create {project} --description {description} -d {cluster},{namespace} -s {repo}'
    output = run(cmd)
    print(output)

    # Add access to resources
    cmd = f'proj allow-cluster-resource {project} "*" "*"'
    output = run(cmd)
    print(output)


def create_app(name, project, namespace, repo, path):
    """ """
    cmd = f"""app create {name}-staging --project {project} --dest-server https://kubernetes.default.svc --dest-namespace {namespace} --repo {repo} --path {path} --sync-policy automated"""
    output = run(cmd)
    print(output)


def copy_apps_from_default_to_staging(project_name, source_namespace):
    apps = get_apps(source_namespace)

    for a in apps:
        print("creating %s"%(a['name']))
        create_app(a['name'], project_name, 'staging', a['repo'], a['path'])


def main(project_name, source_namespace):
    login()
    copy_apps_from_default_to_staging(project_name, source_namespace)

    apps = get_apps('staging')
    for a in apps:
        print(a)


if __name__ == '__main__':
    project_name: str
    source_namespace: str

    if len(sys.argv) > 1:
        project_name = sys.argv[1]
    else:
        project_name= "staging"

    if len(sys.argv) > 2:
        source_namespace = sys.argv[2]
    else:
        source_namespace= "default"

    main(project_name, source_namespace)

