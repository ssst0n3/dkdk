<template>
  <div>
    <div class="row ml-5" style="font-size: 20px">
      <CurrentDirectory class="mt-2" :current_dir="current_dir"
                        v-on:open_directory="open_directory"/>
    </div>
    <div class="row mt-2 ml-3">
      <b-button v-b-modal.modal-upload variant="outline-info" class="ml-3">
        <b-icon icon="upload"/>
        <span class="ml-2">Upload</span>
      </b-button>
      <b-button @click="mkdir"
                variant="outline-info" class="ml-3">
        <!--        v-b-modal.modal-mkdir-->
        <b-icon icon="folder-plus"/>
        <span class="ml-2">Mkdir</span>
      </b-button>
      <b-button @click="delete_files" variant="outline-info" class="ml-3">
        <b-icon icon="trash"/>
        <span class="ml-2">Delete</span>
      </b-button>
      <b-button @click="refresh" variant="outline-info" class="ml-3">
        <b-icon icon="arrow-counterclockwise"/>
        <span class="ml-2">Refresh</span>
      </b-button>
      <b-button @click="rename" variant="outline-info" class="ml-3" v-if="select.length===1">
        <b-icon icon="pencil"/>
        <span class="ml-2">Rename</span>
      </b-button>
      <Move v-if="select.length===1" :file="select[0]" v-on:hide="hide"/>
      <!--      <b-button @click="root" variant="outline-info" class="ml-3">-->
      <!--        <b-icon icon="arrow-return-left"/>-->
      <!--        <span class="ml-2">return</span>-->
      <!--      </b-button>-->
      <!--      <b-button @click="root" variant="outline-info" class="ml-3">-->
      <!--        <b-icon icon="box-arrow-in-up"/>-->
      <!--        <span class="ml-2">root</span>-->
      <!--      </b-button>-->
    </div>
    <FileManager class="mt-2" ref="file_manager"
                 v-on:updateCurrentDir="updateCurrentDir"
                 v-on:updateCurrentDirPath="updateCurrentDirPath"
                 v-on:updateSelect="updateSelect"/>
    <b-modal id="modal-upload" title="Upload file" @ok="upload">
      <span class="my-4">Upload into {{ cwd }}</span>
      <b-form-file
          class="mt-3"
          v-model="file"
          :state="Boolean(file)"
          placeholder="Choose a file or drop it here..."
          drop-placeholder="Drop file here..."
      />
    </b-modal>

    <b-modal id="modal-mkdir" title="Mkdir" @ok="upload">
      <b-input-group prepend="dirname">
        <b-form-input></b-form-input>
      </b-input-group>
    </b-modal>
  </div>
</template>

<script>
import FileManager from "@/components/netdisk/FileManager";
import CurrentDirectory from "@/components/netdisk/CurrentDirectory";
import Move from "@/components/netdisk/Move";

export default {
  name: "NetDisk",
  components: {Move, CurrentDirectory, FileManager},
  data() {
    return {
      file: null,
      cwd: "/",
      current_dir_path: [{filename: "/", id: 0}],
      select_single_item: false,
      select: [],
      current_dir: 0,
    }
  },
  mounted() {
    this.updateCurrentDirPath()
  },
  methods: {
    updateCurrentDir() {
      this.current_dir = this.$refs.file_manager.current_dir
    },
    updateCurrentDirPath() {
      this.current_dir_path = [{filename: "/", id: 0}].concat(this.$refs.file_manager.current_dir_path)
    },
    refresh() {
      this.$refs.file_manager.refresh()
    },
    upload() {
      console.log(this.file)
      this.$bvToast.toast("upload " + this.file.name + " into " + this.cwd, {
        variant: "info",
        solid: true,
      })
    },
    delete_files() {
      this.$refs.file_manager.delete()
    },
    mkdir() {
      this.$refs.file_manager.mkdir()
    },
    root() {
      this.$refs.file_manager.root()
    },
    open_directory(id) {
      this.$refs.file_manager.open_directory(id)
    },
    rename() {
      this.$refs.file_manager.rename()
    },
    updateSelect(select) {
      this.select = select
    },
    hide() {
      this.$refs.file_manager.hide()
    }
  }
}
</script>

<style scoped>

</style>
