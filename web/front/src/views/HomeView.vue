<template>
  <v-app app>
    <div class="loginLayout">
      <LoginView v-if="this.LoginSeen" @update:data="LoginValid"></LoginView>
    </div>
    <TopBar @update:data="handleUpdateData" @update:Login="LoginValid"></TopBar>
    <v-main style="background-color:#f4f4f4">
      <v-container>
        <v-row>
          <v-col cols="3">
            <NavWeb></NavWeb>
            <FriendNav class="mt-5"></FriendNav>
          </v-col>
          <v-col cols="9">
            <v-sheet max-width="60vw" min-height="80vh" rounded="lg">
              <router-view :receivedData="searchInput"></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <footer-bar></footer-bar>
  </v-app>
</template>

<script>
import TopBar from "@/components/TopBar.vue";
import FooterBar from "@/components/FooterBar.vue";
import NavWeb from '@/components/NavWeb.vue';
import FriendNav from "@/components/FriendNav.vue";
import LoginView from "@/components/LoginView.vue";
export default {
  components:{FriendNav, TopBar,FooterBar,NavWeb,LoginView},
  methods:{
    handleUpdateData(data){
      this.searchInput = data
    },
    LoginValid(data){
      this.LoginSeen = data
    }
  },
  data(){
    return{
      searchInput:'',
      LoginSeen:false,
    }
  }
}
</script>

<style scoped>
.loginLayout{
  width:100%;
  z-index:1000;
  position: fixed;
  top:0;
  left:0;
}
</style>
