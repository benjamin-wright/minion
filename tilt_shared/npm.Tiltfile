def npm_init(refresh_path):
    root_dir = __find_root(os.getcwd())
    refresh_dir = os.path.abspath(refresh_path)

    __login(root_dir)
    __build_module(root_dir, refresh_dir, 'eslint-config-minion', [])
    __build_module(root_dir, refresh_dir, 'crd-lib', [ 'eslint-config-minion' ])
    __build_module(root_dir, refresh_dir, 'async', [ 'eslint-config-minion' ])

def __login(root_dir):
    local(
        "cd " + root_dir + "/modules; ./login.sh"
    )

def __build_module(root_dir, refresh_dir, module, deps):
    local_resource(
        'npm: ' + module,
        [
            '/bin/bash',
            '-c',
            """
            set -o errexit
            cd {root_dir}/modules/{module}
            echo "node_modules" > .dockerignore
            DOCKER_BUILDKIT=1 docker build --secret id=npm_token,src=$HOME/.npmrc -t npm-{module} --network host -f ../Dockerfile ./
            npm version patch

            cd {refresh_dir}
            npm update @minion-ci/{module} --strict-ssl false
            """.format(
                root_dir=root_dir,
                refresh_dir=refresh_dir,
                module=module
            )
        ],
        deps=[
            root_dir + '/modules/' + module + '/lib',
            root_dir + '/modules/Dockerfile'
        ],
        resource_deps=[ 'npm: ' + dep for dep in deps ]
    )

def __find_root(filepath):
    files = listdir(filepath)

    if filepath == '/':
        fail('Couldn\'t find shared tilt folder')

    for file in files:
        if os.path.basename(file) == '.gitignore':
            return filepath

    return __find_root(os.path.abspath(os.path.join(filepath, '/..')))
