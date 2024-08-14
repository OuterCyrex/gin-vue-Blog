<template>
  <v-card class="mx-auto" max-width="300">
    <v-img src="../assets/SummerPocketisland.png">
      <v-card-title>
        <v-col align="center">
          <v-avatar size="130" color="grey">
            <img :src="profileInfo.avatar">
          </v-avatar>
          <div class="ma-4 white--text">{{profileInfo.name}}</div>
          <div></div>
        </v-col>
      </v-card-title>
    </v-img>
    <v-divider></v-divider>
    <v-card-title>About Me:</v-card-title>
    <v-card-text>{{profileInfo.desc}}</v-card-text>

    <v-divider color="indigo"></v-divider>

    <v-list>
      <v-list-item>
        <v-list-item-icon>
        <v-icon color="blue">{{'mdi-qqchat'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profileInfo.qqchat}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon>
          <v-icon color="dark">{{'mdi-github'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profileInfo.github}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon>
          <v-icon color="blue">{{'mdi-youtube'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profileInfo.bili}}</v-list-item-content>
      </v-list-item>
    </v-list>
  </v-card>
</template>
<script>
export default {
  data(){
    return {
      profileInfo:{},
    }
  },
  created(){
    this.getProfileInfo()
  },
  methods:{
    async getProfileInfo(){
      const {data:res} = await this.$http.get("profile/1")
      this.profileInfo = res.data
    },
  },
}
</script>
