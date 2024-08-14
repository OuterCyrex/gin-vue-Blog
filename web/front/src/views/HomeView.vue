<template>
  <v-app app>
    <TopBar @update:data="handleUpdateData"></TopBar>
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
export default {
  components:{FriendNav, TopBar,FooterBar,NavWeb},
  methods:{
    handleUpdateData(data){
      this.searchInput = data
    }
  },
  data(){
    return{
      searchInput:'',
    }
  }
}
</script>
