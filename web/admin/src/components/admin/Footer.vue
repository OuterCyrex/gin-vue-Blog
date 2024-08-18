<template>
  <div class="footer">
    <span>Copyright © - Outer Blog 2024</span>
  </div>
</template>

<script>
export default {
  data(){
    return{
      collapsed:false,
    }
  }
}
</script>

<style scoped>
.footer{
  background-color:black;
  text-align: center;
  width: 100vw;
  height:50px;
  padding-top:10px;
}
.footer span{
  font-size: 18px;
  color: white;
}
</style>
