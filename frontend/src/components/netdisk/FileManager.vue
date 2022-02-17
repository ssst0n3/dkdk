<template>
  <div>
    <b-table able-colgroup striped hover
             show-empty empty-text="no files under this directory"
             :items="items" :fields="fields" :busy="busy"
             ref="selectableTable" select-mode="multi" selectable @row-selected="onRowSelected"
             class="text-left">
      <template #table-busy>
        <div class="text-center text-danger my-2">
          <b-spinner class="align-middle"></b-spinner>
          <strong class="ml-2">Loading...</strong>
        </div>
      </template>
      <template #table-colgroup="scope">
        <col
            v-for="field in scope.fields"
            :key="field.key"
            :style="{ width: field.key === 'select' ? '5%' : '180px'}"
        >
      </template>
      <template #head()>

      </template>
      <template #head(select)>
        <b-checkbox @change="selectAllRows"/>
        <!--        select-->
      </template>

      <template #cell(select)="{ rowSelected }">
        <!--        <b-checkbox class="text-center"/>-->
        <template v-if="rowSelected">
          <span aria-hidden="true">&check;</span>
          <span class="sr-only">Selected</span>
        </template>
        <template v-else>
          <span aria-hidden="true">&nbsp;</span>
          <span class="sr-only">Not selected</span>
        </template>
      </template>
      <template #cell(filename)="row">
        <div class="text-left form-inline">
          <b-icon scale="1.6" style="color: sandybrown" icon="folder-fill"
                  v-if="row.item.type===consts.Model.node.type.directory"/>
          <b-icon scale="1.6" variant="secondary" icon="file-earmark-fill" v-else/>
          <div v-if="row.item.filename==='create_new_directory_'">
            <b-input class="ml-3" v-model="new_directory"/>
            <b-button class="ml-2" size="sm" variant="outline-dark" @click="create_directory">
              <b-icon icon="check"/>
            </b-button>
            <b-button class="ml-1" size="sm" variant="outline-dark" @click="items.shift()">
              <b-icon icon="x"/>
            </b-button>
          </div>
          <div v-else-if="row.item.flag==='rename_'">
            <b-input class="ml-3" v-model="rename_filename"/>
            <b-button class="ml-2" size="sm" variant="outline-dark" @click="do_rename">
              <b-icon icon="check"/>
            </b-button>
            <b-button class="ml-1" size="sm" variant="outline-dark" @click="items.shift()">
              <b-icon icon="x"/>
            </b-button>
          </div>
          <span v-else>
            <b-button @click="open_directory(row.item.ID)" v-if="row.item.type===consts.Model.node.type.directory"
                      variant="outline-info"
                      class="ml-2 text-decoration-none border-0">
              <span>{{ row.item.filename }}</span>
            </b-button>
            <b-link v-else class="ml-3 text-decoration-none" style="color: #17a2b8">
              <span class="ml-1">{{ row.item.filename }}</span>
            </b-link>
          </span>
        </div>
      </template>
      <template #cell(size)="row">
        <span v-if="row.item.size!==undefined">{{formatBytes(row.item.size)}}</span>
      </template>
      <template #cell(action)="row">
        <Download v-if="row.item.type===consts.Model.node.type.file"
                  :repositoryId="row.item.repository_id" :digest="row.item.digest"
                  :filename="row.item.filename" :password="row.item.archive_password"
                  :download_count="row.item.download_count" v-on:add_download_count="row.item.download_count+=1"/>
      </template>
    </b-table>
  </div>
</template>

<script>
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";
import Download from "@/components/repository/Download";

export default {
  name: "FileManager",
  components: {Download},
  data() {
    return {
      selected: [],
      busy: false,
      fields: ['select', 'filename', {key: 'size', tdClass: 'align-middle'}, 'action'],
      directories: [],
      files: [],
      items: [
        {type: 0, filename: 'Dickerson', size: 111},
        {type: 1, filename: 'Larsen', size: 222},
        {type: 0, filename: 'Geneva', size: 3333},
        {type: 1, filename: 'Jami', size: 444}
      ],
      current_dir: 0,
      current_dir_path: [],
      consts: consts,
      new_directory: "",
      rename_filename: "",
    }
  },
  created() {
    this.listNodes()
  },
  mounted() {
    this.getCurrentDirPath()
  },
  watch: {
    selected: function () {
      this.$emit("updateSelect", this.selected)
    },
    current_dir: function () {
      this.$emit("updateCurrentDir")
    }
  },
  methods: {
    async create_directory() {
      let data = consts.Model.create_directory_node
      data.filename = this.new_directory
      data.parent_node_id = this.current_dir
      this.items.splice(0, 1, await lightweightRestful.api.createResource(
          consts.api.v1.directory.node,
          data,
          {
            caller: this,
            success_msg: "create success",
          }
      ))
    },
    async getCurrentDirPath() {
      if (this.current_dir === 0) {
        this.current_dir_path = []
      } else {
        this.current_dir_path = await lightweightRestful.api.get(consts.Path.dir.path(this.current_dir))
      }
      this.$emit("updateCurrentDirPath")
    },
    async listNodes() {
      this.busy = true
      await this.listDirectoryUnderDir(this.current_dir)
      await this.listFileUnderDir(this.current_dir)
      this.items = this.directories.concat(this.files)
      this.busy = false
    },
    async listDirectoryUnderDir(dirId) {
      this.directories = await lightweightRestful.api.listResource(consts.api.v1.directory.listDirectoryUnderDir(dirId))
    },
    async listFileUnderDir(dirId) {
      this.files = await lightweightRestful.api.listResource(consts.api.v1.file.listFileUnderDir(dirId))
    },
    onRowSelected(items) {
      this.selected = items
    },
    selectAllRows() {
      if (this.selected.length < this.items.length) {
        this.$refs.selectableTable.selectAllRows()
      } else {
        this.$refs.selectableTable.clearSelected()
      }
    },
    open_directory(id) {
      this.current_dir = id
      this.listNodes()
      this.getCurrentDirPath()
    },
    refresh() {
      this.listNodes()
    },
    async delete() {
      console.log("delete:", this.selected)
      let dir_ids = []
      let file_ids = []
      this.selected.forEach(function (s) {
        if (s.type === consts.Model.node.type.directory) {
          dir_ids.push(s.node_id)
        }
        if (s.type === consts.Model.node.type.file) {
          file_ids.push(s.node_id)
        }
      })
      let ids = file_ids.concat(dir_ids)
      await lightweightRestful.api.exec('delete', consts.api.v1.node, null, ids, {
        caller: this,
      })
      await this.refresh()
    },
    mkdir() {
      let item = {type: consts.Model.node.type.directory, filename: 'create_new_directory_'}
      this.items.splice(0, 0, item,)
    },
    root() {
      this.current_dir = 0
      this.listNodes()
    },
    rename() {
      let select = this.selected[0]
      this.rename_filename = select.filename
      let index = this.$refs.selectableTable.selectedRows.indexOf(true)
      let item = select
      item.flag = "rename_"
      this.items.splice(index, 1, item)
    },
    do_rename() {
      console.log("// todo: rename")
    },
    hide() {
      this.items.splice(this.items.indexOf(this.selected[0]))
    },
    formatBytes(bytes, decimals = 6) {
      if (bytes === 0) return '0 Bytes';

      const k = 1024;
      const dm = decimals < 0 ? 0 : decimals;
      const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

      const i = Math.floor(Math.log(bytes) / Math.log(k));

      return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
    }
  },
}
</script>

<style scoped>

</style>
