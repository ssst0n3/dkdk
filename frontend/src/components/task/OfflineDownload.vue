<template>
  <div class="row" style="margin: 20px;">
    <b-col>
      <b-pagination v-model="currentPage" :total-rows="totalRows" :per-page="perPage" align="fill" size="sm"/>
      <b-table class="mt-1" striped hover :fields="fields"
               :items="tasks"
               :current-page="currentPage" :per-page="perPage">
        <template #cell(action)="row">
          <b-button size="sm" @click="start_task(row.item.ID)" class="mr-2">
            <b-icon icon="caret-right-square-fill"/>
          </b-button>
        </template>
        <template #cell(origin_url)="row">
          <b-button @click="show_content('origin url', row.item.origin_url)" v-b-modal.modal-show_content>Show
          </b-button>
        </template>
        <template #cell(filename)="row">
          <div style="word-break: break-all;white-space: normal;">{{ row.item.filename }}</div>
        </template>
        <template #cell(digest)="row">
          <div style="word-break: break-all;white-space: normal;">{{ row.item.digest }}</div>
          <!--<b-button @click="show_content('digest', row.item.digest)" v-b-modal.modal-show_content>Show</b-button>-->
        </template>
      </b-table>
      <b-modal id="modal-show_content" :title="title" hide-footer>
        <div style="word-break: break-all;white-space: normal;">{{ content }}</div>
      </b-modal>
    </b-col>
    <b-col lg="3">
      <b-tabs>
        <b-tab title="manually" active title-link-class="text-info">
          <b-card :title="(model.ID ? 'Edit Task ID#' + model.ID : 'New Task')" class="border-0">
            <b-input-group prepend="repository">
              <b-form-select class="col-11" v-model="model.repository_id" :options="options_repository"></b-form-select>
              <!--          <b-form-input v-model="model.repository_id"/>-->
            </b-input-group>
            <b-input-group prepend="url" class="mt-2">
              <b-form-input v-model="model.origin_url"/>
            </b-input-group>
            <b-input-group prepend="filename" class="mt-2">
              <b-form-input v-model="model.filename"/>
            </b-input-group>
            <div class="row mt-3">
              <div class="col"/>
              <b-form-checkbox class="col-3" v-model="model.net_disk">net disk</b-form-checkbox>
            </div>
            <div v-if="model.net_disk">
              <b-input-group prepend="alias" class="mt-2">
                <b-form-input v-model="model.filename_in_dkdk"/>
              </b-input-group>
              <SelectDirectory v-on:open_directory="open_directory"/>
            </div>
            <b-btn class="mt-2 float-right" @click="submit">Submit</b-btn>
          </b-card>
        </b-tab>
        <b-tab title="batch" title-link-class="text-info">
          <Batch/>
        </b-tab>
      </b-tabs>
    </b-col>
  </div>
</template>

<script>
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";
import Batch from "@/components/task/Batch";
import SelectDirectory from "@/components/netdisk/SelectDirectory";

export default {
  name: "OfflineDownload",
  components: {SelectDirectory, Batch},
  data() {
    return {
      fields: ['repository_id', 'origin_url', 'status', 'archive_password', 'filename', 'size', 'digest', 'action'],
      model: {
        type: 0,
      },
      tasks: [],
      options_repository: [],
      title: '',
      content: '',
      currentPage: 1,
      totalRows: 3,
      perPage: 20,
    }
  },
  created() {
    this.ListRepositoryConfig()
    this.ListTask()
  },
  methods: {
    populateResourceToEdit(object) {
      this.model = Object.assign({}, object)
    },
    async ListTaskByPage() {
      let params = {
        page: this.currentPage,
        size: this.perPage,
      }
      let tasks = await lightweightRestful.api.get(consts.api.v1.task.task, params, {
        caller: this,
      })
      this.totalRows = tasks.length
      return tasks
    },
    async ListTask() {
      this.tasks = await lightweightRestful.api.listResource(consts.api.v1.task.task, {
        caller: this,
      })
      this.totalRows = this.tasks.length
    },
    async ListRepositoryConfig() {
      let repositories = await lightweightRestful.api.listResource(consts.api.v1.repository.repository, {
        caller: this,
      })
      let options = []
      repositories.forEach(function (repository) {
        options.push({
          value: repository.id,
          text: `${repository.service_address}/${repository.name}:${repository.reference}`,
        })
      })
      this.options_repository = options
    },
    submit() {
      if (this.model.ID) {
        // edit
      } else {
        // create
        console.log("model", this.model)
        lightweightRestful.api.createResource(consts.api.v1.task.task, this.model, {
          caller: this,
        })
      }
    },
    start_task(id) {
      lightweightRestful.api.post(consts.api.v1.task.action(id), {action: "start"}, null, {
        caller: this,
      })
    },
    show_content(title, content) {
      this.title = title
      this.content = content
      this.$bvModal.show("modal-show_content")
    },
    open_directory(directory) {
      this.model.directory_in_dkdk=directory
    }
  }
}
</script>

<style scoped>

</style>
