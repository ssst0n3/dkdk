const api_ = '/api'
const v1 = `${api_}/v1`
const repository = 'repository'
const auth = 'auth'
const initialize = 'initialize'
const task = 'task'
const directory = 'directory'
const file = 'file'
const node = 'node'
const parent = 'parent'

const api = {
    v1: {
        repository: `${v1}/${repository}`,
        auth: `${v1}/${auth}`,
        initialize: `${v1}/${initialize}`,
        task: `${v1}/${task}`,
        directory: `${v1}/${directory}`,
        file: `${v1}/${file}`,
        node: `${v1}/${node}`,
    },
}

export default {
    BaseUrl: process.env.NODE_ENV === 'development' ? 'http://127.0.0.1:14000' : '',
    api: {
        v1: {
            repository: {
                repository: api.v1.repository,
                list: `${api.v1.repository}/list`,
                list_item: (id) => `${api.v1.repository}/list/${id}`,
                download: (digest) => `${api.v1.repository}/download/${digest}`,
                download_by_id: (id, digest) => `${api.v1.repository}/download/${id}/${digest}`,
                upload: `${api.v1.repository}/upload`,
                upload_by_id: (id) => `${api.v1.repository}/upload/${id}`,
            },
            auth: api.v1.auth,
            initialize: {
                initialize: api.v1.initialize,
                create_user: `${api.v1.initialize}/create_user`,
                end: `${api.v1.initialize}/end`
            },
            task: {
                task: api.v1.task,
                action: (id) => `${api.v1.task}/action/${id}`,
                batchTaskCreate: `${api.v1.task}/batch/create`,
            },
            directory: {
                directory: api.v1.directory,
                node: `${api.v1.directory}/node`,
                listDirectoryUnderDir: (id) => `${api.v1.directory}/list/${id}`
            },
            file: {
                file: api.v1.file,
                listFileUnderDir: (id) => `${api.v1.file}/${id}`,
                UploadRepositoryFileToDirectory: (id) => `${api.v1.file}/repository/${id}`,
            },
            node: api.v1.node
        },
    },
    Resource: {
        Repository: repository,
    },
    Path: {
        initialize: `/${initialize}`,
        home: `/`,
        dir: {
            path: (id) => `${api.v1.directory}/path/${id}`
        }
    },
    Model: {
        node: {
            type: {
                directory: 0,
                file: 1,
            },
            parent: parent,
        },
        task: {
            type: {
                OfflineDownload: 0,
            }
        },
        create_directory_node: {
            filename: "",
            parent_node_id: 0,
        },
    },
    Cookie: {
        token: 'token',
        isAdmin: 'is_admin',
        userId: 'user_id'
    }
}
