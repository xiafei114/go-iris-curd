<template>

  <div>
    <Card shadow>
      系统操作
      <div class="text">
        <p class="send-tip">按 Ctrl + Enter 发送</p>
        <textarea v-model="content" @keyup="onKeyup"></textarea>
      </div>
    </Card>


  </div>

</template>

<script>

  import {mapMutations} from 'vuex';

  export default {
    components: {},
    data() {
      return {
        content: ''
      }
    },
    methods: {
      onKeyup(e) {
        if (e.ctrlKey && e.keyCode === 13 && this.content.length) {
          this.sendMessage(this.content);
          this.content = '';
        }
      },
      // (...) 对象展开运算符
      ...mapMutations([
        'sendMessage'
      ])
    },
    watch: {},
    computed: {},
    mounted: function () {

    }
  }
</script>
<style lang="less" scoped>
  .text {
    height: 160px;
    border-top: solid 1px #ddd;

    textarea {
      padding: 10px;
      height: 100%;
      width: 100%;
      border: none;
      outline: none;
      font-family: "Micrsofot Yahei";
      resize: none;
    }
  }

  .send-tip {
    height: 30px;
    line-height: 30px;
    margin: 0;
    padding: 0 5px;
  }
</style>
