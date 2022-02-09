<template>
  <div style="border: darkcyan 1px dashed; padding: 20px">
    <div class="row">
      <b-form-select class="col-11" v-model="selected" :options="options"></b-form-select>
      <b-btn class="col-1" variant="light" v-b-modal.add_repository>
        <b-icon icon="plus" scale="2" variant="info"></b-icon>
      </b-btn>
    </div>
    <div class="mt-3" style="text-align: right">
      <b-btn class="ml-2" @click="fetch">fetch</b-btn>
    </div>
    <b-modal id="add_repository" hide-footer title="add repository">
      <RepositoryConfig is-config/>
    </b-modal>
  </div>
</template>

<script>
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";
import RepositoryConfig from "@/components/config/RepositoryConfig";

export default {
  name: "RepositoryConfigSelect",
  components: {RepositoryConfig},
  data() {
    return {
      selected: 0,
      options: [],
      repositories: []
    }
  },
  watch: {
    selected: function () {
      this.$emit('update_selected', this.selected)
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
      let options = []
      this.repositories.forEach(function (repository) {
        options.push({
          value: repository.id,
          text: `${repository.service_address}/${repository.name}:${repository.reference}`,
        })
      })
      this.options = options
      this.options.push({
        value: 0,
        text: 'all',
      })
    },
    fetch() {
      this.$emit('fetch')
    }
  }
}
</script>

<style scoped>

</style>
