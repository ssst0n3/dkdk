<template>
  <div>
    <b-icon class="mr-2" scale="1.2" style="color: sandybrown" icon="folder-fill"/>
    <b-badge href="#" @click="open_directory(dir.id)" v-for="dir in current_dir_path" :key="dir.id"
             class="align-self-center ml-1" variant="warning">
      {{ dir.filename }}
      <span v-if="dir.id !== 0">/</span>
    </b-badge>
  </div>
</template>

<script>
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";

export default {
  name: "CurrentDirectory",
  props: {
    current_dir: Number,
  },
  data() {
    return {
      current_dir_path: [{filename: "/", id: 0}],
    }
  },
  created() {
    this.getCurrentDirPath()
  },
  methods: {
    open_directory(id) {
      this.$emit("open_directory", id)
    },
    async getCurrentDirPath() {
      if (this.current_dir === 0) {
        this.current_dir_path = []
      } else {
        this.current_dir_path = await lightweightRestful.api.get(consts.Path.dir.path(this.current_dir))
      }
      this.current_dir_path = [{filename: "/", id: 0}].concat(this.current_dir_path)
    },
  },
  watch: {
    current_dir: function () {
      this.getCurrentDirPath()
    }
  }
}
</script>

<style scoped>

</style>
