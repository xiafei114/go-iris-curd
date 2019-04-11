<template>

  <Form ref="entityValidate" :model="entityValidate" :rules="ruleValidate" label-position="right" :label-width="110">
    <Row :gutter="16">
      <h4>基本信息</h4>
      <FormItem prop="product_Code" label="产品编号：">
        <Input type="text" v-model="entityValidate.product_Code" placeholder="Enter text" clearable
               style="width: auto"/>
      </FormItem>
      <FormItem prop="product_Name" label="产品名称：">
        <Input type="text" v-model="entityValidate.product_Name" placeholder="Enter text" clearable
               style="width: auto"/>
      </FormItem>
      <FormItem prop="price" label="单价：">
        <Input type="text" v-model="entityValidate.price" number placeholder="Enter text" clearable
               style="width: auto"/>
      </FormItem>
      <FormItem prop="number" label="数量：">
        <Input type="text" v-model="entityValidate.number" number placeholder="Enter text" clearable
               style="width: auto"/>
      </FormItem>

    </Row>


    <FormItem>
      <Button type="success" icon="md-checkmark" @click="handleSubmit('entityValidate')">保存</Button>
      <Button type="error" icon="md-return-left" style="margin-left: 8px" @click="handleReset('entityValidate')">重置
      </Button>
      <Button type="warning" icon="md-close" style="margin-left: 8px" @click="handleCloseTag">关闭</Button>
    </FormItem>

  </Form>

</template>

<script>
  import {mapMutations} from 'vuex'

  export default {
    data() {
      return {
        entityValidate: {enable: true},
        ruleValidate: {
          product_Code: [
            {required: true, trigger: 'blur',max: 20 ,min:5}
          ],
          product_Name: [
            {required: true, message: '名称不能为空', trigger: 'blur',max: 50 ,min:5}
          ]
        }
      }
    },
    methods: {
      ...mapMutations([
        'closeTag'
      ]),
      handleSubmit(name) {
        console.log(this.entityValidate)
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.$Message.success('Success!')
          } else {
            this.$Message.error('Fail!')
          }
        })
      },
      handleReset(name) {
        this.$refs[name].resetFields()
        this.userValidate = {enable: true}
      },
      handleCloseTag() {
        this.closeTag({
          name: 'add_user_page'
        })
      }
    },
    watch: {},
    computed: {},
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
