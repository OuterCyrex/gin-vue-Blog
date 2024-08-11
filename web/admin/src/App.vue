<template>
  <a-locale-provider :locale="locale">
  <div id="app">
    <router-view/>
  </div>
  </a-locale-provider>
</template>

<script>
import zh_CN from 'ant-design-vue/es/locale-provider/zh_CN'
export default {
  data() {
    return {
      locale: zh_CN,
    };
  },
};
</script>

