<template>
  <div>
    <div>
      <b-card no-body>
        <b-tabs align="center" card active-nav-item-class="font-weight-bold text-uppercase text-light bg-primary">
          <b-tab title="STEP 1. Create User" >
            <CreateUser style="width: 30%;" class="col m-auto"/>
          </b-tab>

          <b-tab no-body title="STEP 2. Picture 2">
            <b-card-img bottom src="https://picsum.photos/600/200/?image=25" alt="Image 25"></b-card-img>
            <b-card-footer>Picture 2 footer</b-card-footer>
          </b-tab>

          <b-tab title="STEP 3. Finish" class="text-center">
            <Finish/>
          </b-tab>
        </b-tabs>
      </b-card>
    </div>
  </div>
</template>

<script>
import CreateUser from "@/components/initialize/CreateUser";
import Finish from "@/components/initialize/Finish";

export default {
  name: "Nav",
  components: {Finish, CreateUser}
}
</script>

<style scoped>

</style>