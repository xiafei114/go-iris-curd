<template>

  <Form ref="entityForm" :model="entity" :rules="ruleValidate" label-position="right" :label-width="110">
    <h4>基本信息</h4>
    <FormItem prop="numCode" label="编号：">
      <Input type="text" v-model="entity.numCode" placeholder="Enter text" clearable
             style="width: auto"/>
    </FormItem>
    <FormItem prop="chnName" label="名称：">
      <Input type="text" v-model="entity.chnName" placeholder="Enter text" clearable
             style="width: auto"/>
    </FormItem>
    <FormItem>
      <Button type="success" icon="md-checkmark" @click="handleSubmit('entityForm')">保存</Button>
      <Button type="error" icon="md-return-left" style="margin-left: 8px" @click="handleReset('entityForm')">重置
      </Button>
      <Button type="warning" icon="md-close" style="margin-left: 8px" @click="handleCloseTag">关闭</Button>
    </FormItem>

  </Form>

</template>

<script>
  import {mapMutations} from 'vuex'
  import {Get, Post, Put} from '@/api/data'

  export default {
    data() {
      return {
        entityBaseUrl:'/productCate',
        entityFormName:'product_cate_form_page',
        isEdit: false,
        entity: {
          numCode: '',
          chnName: '',
        },
        ruleValidate: {
          numCode: [
            {required: true, trigger: 'blur', max: 20, min: 1}
          ],
          chnName: [
            {required: true, message: '名称不能为空', trigger: 'blur', max: 50, min: 1}
          ]
        }
      }
    },
    methods: {
      ...mapMutations([
        'closeTag'
      ]),
      async getData() {
        let id = this.$route.params.id;
        if (id === undefined) {
          return;
        }

        let url = this.entityBaseUrl + '/' + id

        console.log(url);
        let resp = await Get(url);

        console.log(resp);
        this.entity = resp.data.data;
        this.isEdit = true;

      },
      handleSubmit(name) {
        console.log(this.entity)
        this.$refs[name].validate((valid) => {
          if (valid) {
            if (this.isEdit) {
              let url = this.entityBaseUrl;
              Put(url, this.entity).then(resp => {
                this.$Message.success('Success!')
                this.handleCloseTag()
              })
            } else {
              let url = this.entityBaseUrl
              Post(url, this.entity).then(resp => {
                this.$Message.success('Success!')
                this.handleCloseTag()
              })
            }
          } else {
            this.$Message.error('Fail!')
          }
        })
      },
      handleReset(name) {
        if (this.isEdit) {
          this.getData();
        } else {
          this.$refs[name].resetFields()
        }
      },
      handleCloseTag() {
        if (this.isEdit) {
          this.closeTag({
            name: this.entityFormName,
            params: {id: this.$route.params.id}
          })
        } else {
          this.closeTag({
            name: this.entityFormName
          })
        }
      }
    },
    watch: {},
    computed: {},
    created() {
      this.getData();
    },
    mounted: function () {

    }
  }
</script>
<style scoped>
  h4 {
    font-weight: 500;
    color: #6379bb;
    font-size: 15px;
    border-bottom: 1px solid #ddd;
    margin: 8px 10px 25px 10px;
    padding-bottom: 5px;
  }
</style>
