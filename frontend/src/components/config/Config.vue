<template>
  <div class="row" style="margin: 20px;">
    <b-col>
      <b-table class="mt-1" striped hover :items="repositories" :fields="fields" :busy="busy">
        <template #cell(action)="row">
          <div class="row">
            <b-button size="sm" @click.prevent="populateResourceToEdit(row.item)" class="mr-2">
              <b-icon icon="pen"/>
            </b-button>
            <b-button size="sm" @click="deleteResource(row.item.id)" class="mr-2">
              <b-icon icon="trash"/>
            </b-button>
          </div>
        </template>
      </b-table>
    </b-col>
    <b-col lg="3">
      <b-card :title="(model.id ? 'Edit Repository ID#' + model.id : 'New Repository')">
        <RepositoryConfig v-on:clean="clean" is-config :update="model.id !== undefined" ref="repository_config"/>
      </b-card>
    </b-col>
  </div>
</template>

<script>
import RepositoryConfig from "@/components/config/RepositoryConfig";
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";

export default {
  name: "Config",
  components: {RepositoryConfig},
  data() {
    return {
      busy: false,
      fields: ['service_address', 'name', 'reference', 'private', 'username', 'action'],
      repositories: [],
      model: {}
    }
  },
  created() {
    this.ListRepositoryConfig()
  },
  methods: {
    async ListRepositoryConfig() {
      this.repositories = await lightweightRestful.api.listResource(consts.api.v1.repository.repository, {
        caller: this,
      })
    },
    populateResourceToEdit(object) {
      this.model = object
      this.$refs.repository_config.address = `${object.service_address}/${object.name}:${object.reference}`
      this.$refs.repository_config.model = Object.assign({}, object)
    },
    clean() {
      this.model = {}
    },
    async deleteResource(id) {
      await lightweightRestful.api.deleteResource(consts.api.v1.repository.repository, id, {
        caller: this,
      })
    }
  }
}
</script>

<style scoped>

</style>
