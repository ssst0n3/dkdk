<template>
  <div>
    <div class="row mt-2 ml-2">
      <span>Directory</span>
      <CurrentDirectory v-on:open_directory="open_directory" :current_dir="model" ref="current_directory"
                        class="ml-2"/>
    </div>
    <b-form-select v-show="directories.length>0" class="mt-2" v-model="model" :options="directories"/>
  </div>
</template>

<script>
import CurrentDirectory from "@/components/netdisk/CurrentDirectory";
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";

export default {
  name: "SelectDirectory",
  components: {CurrentDirectory},
  props: {
    directory: {
      type: Number,
      default: 0,
    },
  },
  data() {
    return {
      model: this.directory,
      directories: [],
    }
  },
  created() {
    this.list_directories()
  },
  watch: {
    directory: function () {
      this.model = this.directory
    },
    model: function () {
      this.open_directory(this.model)
      this.list_directories()
    }
  },
  methods: {
    open_directory(id) {
      this.model = id
      this.$emit("open_directory", id) // avoid overwritten props
    },
    async list_directories() {
      this.directories = []
      let directory = await lightweightRestful.api.listResource(consts.api.v1.directory.listDirectoryUnderDir(this.model))
      directory.forEach(e => this.directories.push({
        value: e.ID,
        text: e.filename,
      }))
    },
  }
}
</script>

<style scoped>

</style>
