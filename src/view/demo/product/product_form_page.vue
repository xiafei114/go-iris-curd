<template>

  <Form ref="entityForm" :model="entity" :rules="ruleValidate" label-position="right" :label-width="110">
    <Row :gutter="16">
      <h4>基本信息</h4>
      <FormItem prop="product_Code" label="产品编号：">
        <Input type="text" v-model="entity.product_Code" placeholder="Enter text" clearable
               style="width: auto"/>
      </FormItem>
      <FormItem prop="product_Name" label="产品名称：">
        <Input type="text" v-model="entity.product_Name" placeholder="Enter text" clearable
               style="width: auto"/>
      </FormItem>
      <FormItem prop="price" label="单价：">
        <Input type="text" v-model="entity.price" number placeholder="Enter text" clearable
               style="width: auto"/>
      </FormItem>
      <FormItem prop="number" label="数量：">
        <Input type="text" v-model="entity.number" number placeholder="Enter text" clearable
               style="width: auto"/>
      </FormItem>

    </Row>


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
  import {Get,Post,Put} from '@/api/data'

  export default {
    data() {
      return {
        isEdit:false,
        entity: {
          product_Code:'',
          product_Name:'',
          price:0,
          number:0,
        },
        ruleValidate: {
          product_Code: [
            {required: true, trigger: 'blur',max: 20 ,min:1}
          ],
          product_Name: [
            {required: true, message: '名称不能为空', trigger: 'blur',max: 50 ,min:1}
          ]
        }
      }
    },
    methods: {
      ...mapMutations([
        'closeTag'
      ]),
      async getData(){
        let id = this.$route.params.id;
        if(id === undefined){
          return;
        }

        let url = '/product/' + id

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
            if(this.isEdit){
              let url = '/product';
              Put(url,this.entity).then(resp => {
                this.$Message.success('Success!')
                this.handleCloseTag()
              })
            }else{
              let url = '/product';
              Post(url,this.entity).then(resp => {
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
        if(this.isEdit){
          this.getData();
        }else{
          this.$refs[name].resetFields()
        }
      },
      handleCloseTag() {
        if(this.isEdit){
          this.closeTag({
            name: 'product_form_page',
            params: {id: this.$route.params.id}
          })
        }else{
          this.closeTag({
            name: 'product_form_page'
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
