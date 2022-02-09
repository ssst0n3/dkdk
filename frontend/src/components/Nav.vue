<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-navbar-brand href="/#/">DkDk</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item href="/#/net_disk">NetDisk</b-nav-item>
          <b-nav-item href="/#/offline_download">OfflineDownload</b-nav-item>
          <b-nav-item href="/#/config">Config</b-nav-item>
          <b-nav-item href="#" disabled>Disabled</b-nav-item>
        </b-navbar-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
            <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
          </b-nav-form>

          <b-nav-item-dropdown text="Lang" right>
            <b-dropdown-item href="#">EN</b-dropdown-item>
            <b-dropdown-item href="#">ES</b-dropdown-item>
            <b-dropdown-item href="#">RU</b-dropdown-item>
            <b-dropdown-item href="#">FA</b-dropdown-item>
          </b-nav-item-dropdown>

          <b-nav-item-dropdown right>
            <!-- Using 'button-content' slot -->
            <template #button-content>
              <em>User</em>
            </template>
            <b-dropdown-item href="#">Profile</b-dropdown-item>
            <b-dropdown-item href="#">Sign Out</b-dropdown-item>
          </b-nav-item-dropdown>

          <Auth :base_url="base_url" :auth_uri="auth_uri"
                welcome_class="text-light" title="login to get persistent storage"
                button_size="sm" style="line-height: 2.5"
                :after_sign_in="function (){$router.go(0)}" :after_sign_out="function () {$router.go(0)}"
                ref="auth"
          />
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>

import consts from "@/consts";

export default {
  name: "Nav",
  data() {
    return {
      base_url: consts.BaseUrl,
      auth_uri: consts.api.v1.auth
    }
  }
}
</script>

<style scoped>

</style>
