def npm_init():
    root_dir = find_root(os.getcwd())
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

def npm_module(module):
    root_dir = find_root(os.getcwd())

    local_resource(
        'npm: ' + module,
        [
            '/bin/bash',
            '-c',
            """
            cd {root_dir}/modules/{module}
            docker run --rm -v $(pwd):/var/app/src --network host npm_builder 'npm install && npm test && npm run lint && npm version patch && npm publish'
            """.format(
                root_dir=root_dir,
                module=module
            )
        ],
        deps=[root_dir + '/modules/' + module + '/lib'],
        resource_deps=['npm: build image']
    )

    return 'npm: ' + module

def find_root(filepath):
    files = listdir(filepath)

    if filepath == '/':
        fail('Couldn\'t find shared tilt folder')

    for file in files:
        if os.path.basename(file) == '.gitignore':
            return filepath

    return find_root(os.path.abspath(os.path.join(filepath, '/..')))
