<template>
  <v-app app>
    <div class="loginLayout">
      <LoginView v-if="this.LoginSeen" @update:data="LoginValid"  @update:isLoggedIn="changeAvatar"></LoginView>
    </div>
    <TopBar @update:data="handleUpdateData" @update:Login="LoginValid" :isLoggedIn="isLoggedIn"></TopBar>
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
    <div :class="{ 'slide-in-from-bottom': true,'visible':scrollVisible}" style="position:fixed;bottom:-7px;right:1vw;display:block;">
      <button id="scrollToTopButton" @click="scrollBack()" style="height:170px;width:170px;background-image:url(https://august-soft.com/cms/wp-content/themes/august/img/common/pagetop_01.png);background-size:cover;background-color:transparent;border:none;">
      </button>
      </div>
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
  created(){
    this.scrollVisible = false;
    this.changeAvatar()
  },
  mounted() {
    window.addEventListener('scroll', this.showScroll);
    window.addEventListener('resize', this.updateResize);
  },
  methods:{
    handleUpdateData(data){
      this.searchInput = data
    },
    LoginValid(data){
      this.LoginSeen = data
    },
    scrollBack() {
      window.scrollTo({top: 0, behavior: "smooth"});
    },
    changeAvatar(isLoggedIn){
      this.isLoggedIn = isLoggedIn
    },
    showScroll() {
      this.scrollVisible = window.scrollY > window.innerHeight;
    },
    updateResize() {
      this.viewportWidth = window.innerWidth;
      this.scrollVisible = this.viewportWidth > 1264;
    },
  },
  data(){
    return{
      searchInput:'',
      LoginSeen:false,
      isLoggedIn:false,
      scrollVisible:false,
      viewportWidth:window.innerWidth,
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
.slide-in-from-bottom {
  transform: translateY(100%);
  opacity: 0;
  transition: transform 0.3s ease-out, opacity 0.3s ease-out;
}
.slide-in-from-bottom.visible {
  transform: translateY(0);
  opacity: 1;
}
</style>
