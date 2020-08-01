const crds = require('@minion-ci/crd-lib');

const NAMESPACE = process.env.NAMESPACE;
const SIDECAR_IMAGE = process.env.SIDECAR_IMAGE;

module.exports = {
    clean,
    createJob
};

async function clean() {
    const client = await crds.client.get();
    await client.apis.batch.v1.namespaces(NAMESPACE).jobs.delete({ qs: { labelSelector: 'minion.ponglehub.co.uk/integration=true' } });
}

async function createJob(name, version) {
    const manifest = {
        apiVersion: 'batch/v1',
        kind: 'Job',
        metadata: {
            name,
            namespace: NAMESPACE,
            labels: {
                'minion.ponglehub.co.uk/integration': 'true'
            }
        },
        spec: {
            template: {
                spec: {
                    restartPolicy: 'Never',
                    initContainers: [
                        {
                            name: 'version',
                            image: 'docker.io/alpine',
                            args: ['/bin/sh', '-c', `echo "${version}" > /output/version.txt`],
                            volumeMounts: [
                                {
                                    name: 'versions',
                                    mountPath: '/output'
                                }
                            ]
                        }
                    ],
                    containers: [
                        {
                            name: 'sidecar',
                            image: SIDECAR_IMAGE,
                            volumeMounts: [
                                {
                                    name: 'versions',
                                    mountPath: '/input'
                                }
                            ]
                        }
                    ],
                    volumes: [
                        {
                            name: 'versions',
                            emptyDir: {}
                        }
                    ]
                }
            }
        }
    };

    const client = await crds.client.get();
    await client.apis.batch.v1.namespaces(NAMESPACE).job.post({ body: manifest });
};
