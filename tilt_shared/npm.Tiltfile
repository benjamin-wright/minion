def npm_init():
    root_dir = __find_root(os.getcwd())

    __builder_image(root_dir)
    __build_module(root_dir, 'eslint-config-minion', [])
    __build_module(root_dir, 'crd-lib', [ 'eslint-config-minion' ])

def npm_module(module):
    return 'npm: ' + module

def __builder_image(root_dir):
    local_resource(
        'npm: build image',
        [
            '/bin/bash',
            '-c',
            """
            cd {root_dir}/modules/build-image
            docker build -t npm_builder .
            """.format(
                root_dir=root_dir
            )
        ]
    )

def __build_module(root_dir, module, deps):
    local_resource(
        'npm: ' + module,
        [
            '/bin/bash',
            '-c',
            """
            cd {root_dir}/modules/{module}
            docker run --rm -v $(pwd):/var/app/src --network host npm_builder 'npm install --no-audit && npm test --silent && npm run lint --silent && npm version patch && npm publish'
            """.format(
                root_dir=root_dir,
                module=module
            )
        ],
        deps=[root_dir + '/modules/' + module + '/lib'],
        resource_deps=['npm: build image'] + [ 'npm: ' + dep for dep in deps ]
    )

def __find_root(filepath):
    files = listdir(filepath)

    if filepath == '/':
        fail('Couldn\'t find shared tilt folder')

    for file in files:
        if os.path.basename(file) == '.gitignore':
            return filepath

    return __find_root(os.path.abspath(os.path.join(filepath, '/..')))
