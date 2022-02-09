<template>
  <div v-if="this.at_initialize">
    <h1 class="text-center">Initialize</h1>
    <Nav/>
  </div>
</template>

<script>
import api from "@/api";
import consts from "@/consts";
import Nav from "@/components/initialize/Nav";

export default {
  name: "Initialize",
  components: {Nav},
  data() {
    return {
    }
  },
  computed: {
    at_initialize() {
      return this.$route.path === consts.Path.initialize
    }
  },
  async created() {
    if (!this.at_initialize) {
      this.$root.should_initialize = await api.check_initialize()
      if (this.$root.should_initialize) {
        await this.$router.push({path: consts.Path.initialize})
      }
    }
  },
  methods: {}
}
</script>

<style scoped>

</style>
