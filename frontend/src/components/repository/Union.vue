<template>
  <div>
    <RepositoryConfig v-if="!loggedIn" v-on:fetch="$emit('fetch', 'config')" ref="repository_config"
                      class="col-6 mt-2" style="border: darkcyan 1px dashed; padding: 20px"/>
    <RepositoryConfigSelect v-else ref="repository_config_id" v-on:fetch="$emit('fetch', 'id')"
                            v-on:update_selected="update_repository_config_id"
                            class="col-6 mt-2" style="border: darkcyan 1px dashed; padding: 20px"/>
    <div class="col"/>
    <uploader :repository="repository_config" v-show="!busy" class="col-5 mt-2" style="margin: auto"/>
  </div>
</template>

<script>
import RepositoryConfig from "@/components/config/RepositoryConfig";
import Uploader from "@/components/Uploader";
import RepositoryConfigSelect from "@/components/repository/RepositoryConfigSelect";

export default {
  name: "Union",
  components: {RepositoryConfigSelect, Uploader, RepositoryConfig},
  props: {
    busy: Boolean,
  },
  data() {
    return {
      repository_config: {},
      loggedIn: false,
      repository_config_id: 0,
    }
  },
  watch: {
    repository_config: function () {
      this.$emit('update_repository_config', this.repository_config)
    },
  },
  mounted() {
    this.repository_config = this.$refs.repository_config.model
    if (this.$root.$children[0].$refs.nav.$refs.auth.loggedIn) {
      this.loggedIn = true
    } else {
      if (this.$root.should_initialize === false) {
        this.$bvModal.show('sign')
      }
    }
  },
  methods: {
    update_repository_config_id(id) {
      this.repository_config_id = id
      this.$emit('update_repository_config_id', this.repository_config_id)
    }
  }
}
</script>

<style scoped>

</style>
