const crds = require('@minion-ci/crd-lib');

const NAMESPACE = process.env.NAMESPACE;
const SIDECAR_IMAGE = process.env.SIDECAR_IMAGE;
const SERVICE_ACCOUNT = process.env.SERVICE_ACCOUNT;

module.exports = {
    clean,
    createJob,
    getStatus
};

async function clean() {
    const client = await crds.client.get();
    await client.apis.batch.v1.namespaces(NAMESPACE).jobs.delete({ qs: { labelSelector: 'minion.ponglehub.co.uk/integration=true' } });
    await client.api.v1.namespaces(NAMESPACE).pods.delete({ qs: { labelSelector: 'minion.ponglehub.co.uk/integration=true' } });

    const versions = await crds.versions.list(NAMESPACE);
    await Promise.all(
        versions.map(v => crds.versions.delete(NAMESPACE, v.metadata.name))
    );
}

async function createJob(name, version, resource) {
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
}

async function getStatus(name) {
    const client = await crds.client.get();
    const job = await client.apis.batch.v1.namespaces(NAMESPACE).job(name).get();

    return job.body.status;
}
