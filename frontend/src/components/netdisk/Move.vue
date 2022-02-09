<template>
  <div>
    <b-button v-b-modal.modal-move variant="outline-info" class="ml-3">
      <b-icon icon="arrows-move"/>
      <span class="ml-2">Move To</span>
    </b-button>
    <b-modal id="modal-move" :title="`Move ${file.filename} To`" @ok="move">
      <SelectDirectory :directory="directory" v-on:open_directory="open_directory"/>
    </b-modal>
  </div>
</template>

<script>
import SelectDirectory from "@/components/netdisk/SelectDirectory";
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";

export default {
  name: "Move",
  components: {SelectDirectory,},
  props: {
    file: Object,
  },
  data() {
    return {
      directory: 0,
    }
  },
  methods: {
    // avoid overwritten props
    open_directory(id) {
      this.directory = id
    },
    async move() {
      let data = {}
      data[consts.Model.node.parent] = this.directory
      await lightweightRestful.api.updateResource(consts.api.v1.node, this.file.node_id, data, {
        caller: this,
      })
      this.$emit("hide")
    }
  }
}
</script>

<style scoped>

</style>
