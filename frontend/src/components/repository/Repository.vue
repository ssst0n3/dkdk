<template>
  <div class="col-11" style="margin: auto">
    <!--    <div class="row">-->
    <union class="row" :busy="busy" v-on:fetch="fetch"
           v-on:update_repository_config_id="update_repository_config_id"
           v-on:update_repository_config="update_repository_config"
           ref="union"/>
    <!--    </div>-->
    <b-table class="mt-1" striped hover :items="files" :fields="fields" :busy="busy">
      <template #table-busy>
        <div class="text-center text-danger my-2">
          <b-spinner class="align-middle"></b-spinner>
          <strong>Loading...</strong>
        </div>
      </template>
      <template #cell(alias)="row">
        {{ row.item.filename_in_dkdk }}
      </template>
      <template #cell(action)="row">
        <Download :repositoryId="repository_id" :digest="row.item.digest"
                  :filename="row.item.filename" :password="row.item.archive_password"
                  :repository_config="repository_config"/>
        <!--        <b-button size="sm" @click="download(row.item.filename, row.item.digest)" class="mr-2">Download</b-button>-->
        <b-button @click="add_to_net_disk(row.item)" size="sm" class="mr-2"
                  :variant="row.item.in_net_disk?'success':''" :disabled="row.item.in_net_disk">
          NetDisk
        </b-button>
      </template>
    </b-table>
    <b-modal id="add_to_net_disk" :title="`Add ${item.filename} to NetDisk`" @ok="do_add_to_net_disk">
      <b-input-group prepend="filename">
        <b-form-input v-model="item.filename"/>
      </b-input-group>
      <SelectDirectory :directory="directory" v-on:open_directory="open_directory"/>
    </b-modal>
  </div>
</template>

<script>
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";
import Union from "@/components/repository/Union";
import Download from "@/components/repository/Download";
import SelectDirectory from "@/components/netdisk/SelectDirectory";

export default {
  name: "Repository",
  components: {SelectDirectory, Download, Union},
  data() {
    return {
      fields: ['filename', 'size', 'action'],
      files: [],
      busy: false,
      link: '',
      filename: '',
      item: {},
      directory: 0,
      loggedIn: false,
      repository_id: 0,
      repository_config: {}
    }
  },
  mounted() {
    this.loggedIn = this.$refs.union.loggedIn
    this.repository_config = this.$refs.union.$refs.repository_config.model
  },
  watch: {},
  methods: {
    update_repository_config_id(id) {
      this.repository_id = id
    },
    update_repository_config(config) {
      this.repository_config = config
    },
    // avoid overwritten props
    open_directory(id) {
      this.directory = id
    },
    fetch(type) {
      switch (type) {
        case 'config':
          this.fetch_from_repository_config()
          break
        case 'id': {
          this.fetch_from_repository_config_id()
          break
        }
      }
    },
    async fetch_from_repository_config_id() {
      this.busy = true
      let id = this.$refs.union.$refs.repository_config_id.selected
      let api = id > 0 ? consts.api.v1.repository.list_item(id) : consts.api.v1.repository.list
      this.files = await lightweightRestful.api.listResource(api, {
        caller: this,
        error_msg: "failed"
      })
      this.busy = false
    },
    async fetch_from_repository_config() {
      this.busy = true
      let data = this.$refs.union.$refs.repository_config.model
      this.files = await lightweightRestful.api.post(consts.api.v1.repository.list, null, data, {
        caller: this,
      })
      this.busy = false
    },
    add_to_net_disk(item) {
      this.item = item
      this.$bvModal.show("add_to_net_disk")
    },
    async do_add_to_net_disk() {
      console.log("do_add_to_net_disk")
      let data = this.item
      data.directory_id = this.directory
      await lightweightRestful.api.createResource(
          consts.api.v1.file.UploadRepositoryFileToDirectory(this.$refs.union.$refs.repository_config_id.selected),
          data,
          {
            caller: this,
          }
      )
      this.$bvModal.hide('add_to_net_disk')
    }
  }
}
</script>

<style scoped>

</style>
