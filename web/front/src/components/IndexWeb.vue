<template>
  <v-card>
    <v-col class="pl-5 pr-5">
    <v-img src="https://img.oulu.me/019%20Malibu%20Beach.png" max-height="300px" class="rounded">
      <v-card-title class="d-flex justify-center">
          <v-card-title class="d-flex flex-row justify-center mb-5" style="user-select: none; font-size: calc(20px + 2vw); padding-top:10%; font-family:'Arial Rounded MT Bold',sans-serif;color:white">
            OuterCyrex Blog
          </v-card-title>
        <v-card-text class="d-flex flex-row justify-center" style="color:white;user-select: none; font-size: calc(20px + 0.5vw)">欢迎访问我的博客</v-card-text>
      </v-card-title>
    </v-img>
      <v-row class="pt-5 pr-3" rows="12">
        <v-col cols="5">
          <v-card class="pa-0">
            <v-col class="pa-0">
              <v-date-picker width="100%" color="blue"></v-date-picker>
            </v-col>
          </v-card>
        </v-col>
        <v-col cols="7" align-self="center">
          <v-row rows="12" class="mb-2">
            <v-col cols="12" class="pa-0">
              <v-card>
                <v-card-title>现在是北京时间：</v-card-title>
                <v-card-text><v-icon class="mr-2">{{'mdi-clock'}}</v-icon>{{this.timeString}}</v-card-text>
              </v-card>
            </v-col>
          </v-row>
          <v-row rows="8">
            <v-col cols="12" class="pa-0">
              <v-card @click="$router.push(`/detail/${LastArticleInfo.ID}`)">
                <v-col cols="12">
                  <v-card-text color="blue">
                    <div style="font-size:2vw">最新博客文章</div>
                  </v-card-text>
                  <v-card-title class="mb-2">
                    <v-chip color="pink" label class="white--text mr-2">{{this.LastArticleInfo.Category.name}}</v-chip>
                    {{this.LastArticleInfo.title}}.
                  </v-card-title>
                  <v-card-subtitle>{{this.LastArticleInfo.desc}}</v-card-subtitle>
                  <v-divider></v-divider>
                  <v-card-text>
                    <v-icon class="mr-2">{{'mdi-calendar-month'}}</v-icon>
                    <span>上传时间：{{this.LastArticleInfo.CreatedAt.slice(0,10) + ' ' + this.LastArticleInfo.CreatedAt.slice(11,19)}}</span>
                  </v-card-text>
                  <v-card-text></v-card-text>
                </v-col>
              </v-card>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
      <v-row rows="12">
        <v-col cols="12">

        </v-col>
      </v-row>
    </v-col>
  </v-card>
</template>
<script>

export default {
  data(){
    return{
      LastArticleInfo:{},
      timeString:'00:00:00',
    }
  },
  created(){
    this.getArticleList()
    setInterval(this.updateClock,1000)
  },
  methods:{
    async getArticleList(){
      const {data:res} = await this.$http.get("articles",{
        params:{
          name:'',
          pagesize:1,
          pagenum:1
        }
      })
      this.LastArticleInfo = res.data[0];
    },
    updateClock(){
    const now = new Date();
    let hours = now.getHours();
    let minutes = now.getMinutes();
    let seconds = now.getSeconds();
    hours = hours.toString().padStart(2, '0');
    minutes = minutes.toString().padStart(2, '0');
    seconds = seconds.toString().padStart(2, '0');
    this.timeString = `${hours}:${minutes}:${seconds}`;
    },
  }
}
</script>
