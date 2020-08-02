const crds = require('@minion-ci/crd-lib');

const NAMESPACE = process.env.NAMESPACE;
const SIDECAR_IMAGE = process.env.SIDECAR_IMAGE;
const SERVICE_ACCOUNT = process.env.SERVICE_ACCOUNT;

module.exports = {
    clean,
    createJob
};

async function clean() {
    const client = await crds.client.get();
    await client.apis.batch.v1.namespaces(NAMESPACE).jobs.delete({ qs: { labelSelector: 'minion.ponglehub.co.uk/integration=true' } });
    await client.api.v1.namespaces(NAMESPACE).pods.delete({ qs: { labelSelector: 'minion.ponglehub.co.uk/integration=true' } });
}

async function createJob(name, version, pipeline, resource) {
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
                metadata: {
                    labels: {
                        'minion.ponglehub.co.uk/integration': 'true'
                    }
                },
                spec: {
                    serviceAccount: SERVICE_ACCOUNT,
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
                            env: [
                                { name: 'PIPELINE', value: pipeline },
                                { name: 'RESOURCE', value: resource }
                            ],
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
