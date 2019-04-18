<template>

  <Form ref="entityForm" :model="entity" :rules="ruleValidate" label-position="right" :label-width="110">
    <h4>基本信息</h4>
    <FormItem prop="name" label="名称：">
      <Input type="text" v-model="entity.name" placeholder="Enter text" clearable
             style="width: auto"/>
    </FormItem>
    <FormItem prop="path" label="路径：">
      <Input type="text" v-model="entity.path" placeholder="Enter text" clearable
             style="width: auto"/>
    </FormItem>
    <FormItem prop="modular" label="模块：">
      <Input type="text" v-model="entity.modular" placeholder="Enter text" clearable
             style="width: auto"/>
    </FormItem>
    <FormItem prop="component" label="组件：">
      <Input type="text" v-model="entity.component" placeholder="Enter text" clearable
             style="width: auto"/>
    </FormItem>
    <FormItem prop="isSub" label="子菜单：">
      <Checkbox v-model="entity.isSub" style="width: auto">Checkbox</Checkbox>
    </FormItem>
    <!--<FormItem prop="meta" label="元信息：">-->
      <!--<Input type="textarea" v-model="entity.meta" :autosize="{maxRows: 5}" placeholder="Enter text"-->
             <!--/>-->
    <!--</FormItem>-->
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
        entityBaseUrl:'/admin/menu',
        entityFormName:'menu_form_page',
        isEdit: false,
        entity: {
          name: '',
          path: '',
          modular: '',
          component: '',
          isSub: true,
          // meta:'{ "hideInMenu": false, "showAlways": false, "notCache": false, "access": null, "icon": "ios-settings", "href": "", "title": "" }'
        },
        ruleValidate: {
          name: [
            {required: true, message: '名称不能为空', trigger: 'blur', max: 20, min: 1}
          ],
          path: [
            {required: true, message: '路径不能为空', trigger: 'blur', max: 50, min: 1}
          ],
          modular: [
            {required: true, message: '模块不能为空', trigger: 'blur', max: 50, min: 1}
          ],
          component: [
            {required: true, message: '组件不能为空', trigger: 'blur', max: 50, min: 1}
          ],
          meta: [
            {required: true, message: '元信息不能为空', trigger: 'blur', max: 500, min: 1}
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
