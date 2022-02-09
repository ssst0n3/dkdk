<template>
  <div id="app">
    <Initialize v-show="false"/>
    <Nav ref="nav" v-show="$route.meta.navShow!==false"/>
    <router-view/>
  </div>
</template>

<script>
import Nav from "@/components/Nav";
import Initialize from "@/components/initialize/Initialize";

export default {
  name: 'App',
  components: {
    Initialize,
    Nav,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
