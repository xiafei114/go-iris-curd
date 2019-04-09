<template>

  <div>
    <!--新增部门modal-->
    <Modal v-model="showCreateDepModal" @on-visible-change="createDepModalStatus" :z-index=1 draggable>
      <p slot="header">
        <Icon type="information-circled"></Icon>
        <span>{{text.titleText}}</span>
      </p>
      <Form ref="addDepValidate" :model="addDepValidate" :rules="addDepRule" label-position="right" :label-width="100">
        <Row :gutter="10">
          <Col span="24" push="1">
            <FormItem label="上级部门：">
              <Row>
                <Col span="12">
                  <Button type="dashed" long @click="showDepTree"><Icon type="md-hand" />{{selectSuperDep.superDepName}}</Button>
                </Col>
              </Row>
              <!--<div @click="aa"><Input disabled type="text" v-model="addDepValidate.parentId" placeholder="dep name" clearable style="width: auto" /></div>-->
            </FormItem>
            <FormItem prop="depName" label="部门名称：">
              <Input type="text" v-model="addDepValidate.depName" placeholder="dep name" clearable style="width: auto" />
            </FormItem>
            <FormItem label="领导名称：">
              <Input type="text" v-model="addDepValidate.leader" prefix="ios-contact" placeholder="leader" clearable style="width: auto" />
            </FormItem>
            <FormItem label="手机号码：">
              <Input type="text" v-model="addDepValidate.phone" prefix="ios-phone-portrait" placeholder="phone" clearable style="width: auto" />
            </FormItem>
            <FormItem label="邮箱：">
              <Input type="text" v-model="addDepValidate.email" prefix="ios-mail" placeholder="email" clearable style="width: auto" />
            </FormItem>
            <FormItem label="启用状态：">
              <i-switch size="large" v-model="addDepValidate.disabled">
                <span slot="open">ON</span>
                <span slot="close">OFF</span>
              </i-switch>
            </FormItem>
            <FormItem label="部门描述：">
              <Input type="textarea" :rows="2" v-model="addDepValidate.depDesc" prefix="ios-phone-portrait" placeholder="description" clearable style="width: auto" />
            </FormItem>
          </Col>
        </Row>
      </Form>
      <div slot="footer">
        <Button type="warning" @click="resetAddDep">重置</Button>
        <Button type="success" @click="handDep" :disabled="selectSuperDep.isSelected">{{text.typeText}}</Button>
      </div>
    </Modal>

    <!--部门树-->
    <Modal v-model="showModalDepTree" draggable :z-index=2>
      <p slot="header">
        <Icon type="information-circled"></Icon>
        <span>上级部门选择</span>
      </p>
      <Row><Col push="5"><Tree :data="allDeps" @on-select-change="selectDep"></Tree></Col></Row>
      <div slot="footer">
        <!--<Button @click="resetAddDep('addDepValidate')">取消</Button>-->
        <!--<Button type="success" @click="addDep('addDepValidate')" :disabled="selectSuperDep.isSelect">确定</Button>-->
      </div>
    </Modal>

  </div>

</template>

<script>
import { Get, Post, Put } from '@/api/data'
export default {
  name: 'DepModals',
  props: {
    // showCreateDepModal:{type: Boolean}
  },
  data () {
    return {
      text: {typeText: '创建', titleText: '新增部门',},
      // 新加部门
      showCreateDepModal: false,
      selectSuperDep: { superDepName: '请选择', isSelected: true },
      allDeps: [], // 拉取的菜单树
      constDep: 'addDepValidate',
      addDepValidate: { disabled: true },
      addDepRule: {
        depName: [
          { required: true, message: '部门名字不能为空', trigger: 'blur' }
        ]
      },
      showModalDepTree: false,
      isModify: false
    }
  },
  methods: {
    // 创建部门和修改部门的modal的状态操作
    changeCreateDepModal (show, isModify, info) {
      this.showCreateDepModal = show
      if (isModify) {
        this.isModify = true
        this.text = {typeText: '修改', titleText: '修改部门',}
        this.addDepValidate = info.dep
        this.selectSuperDep = {superDepName: info.superName, isSelected: false}
        // this.handDep(isModify)
      }
    },
    // 监听创建部门的modal状态，false=关闭状态
    createDepModalStatus (status) {
      if (!status) this.resetAddDep()
    },
    // 显示部门树结构
    showDepTree () {
      this.showModalDepTree = true
      this.allDeps = []
      Get('/admin/dep/tree').then(resp => {
        this.allDeps.push(resp.data.data)
      })
    },
    // 选择上级部门操作
    selectDep (dep) {
      if (dep instanceof Array && dep.length === 1) {
        let { id, title } = dep[0]
        this.selectSuperDep.superDepName = title
        this.selectSuperDep.isSelected = false
        this.addDepValidate.parentId = id
      } else {
        this.selectSuperDep = { superDepName: '', isSelected: true }
        this.addDepValidate.parentId = 1
      }
      console.log(this.addDepValidate.parentId)
    },
    // 部门请求操作
    handDep() {
      console.log('isModify:', this.isModify)
      if (this.isModify) {// 修改部门
        this.modifyDep()
      }else {
        this.createDep()
      }
    },
    // 修改部门
    modifyDep() {
      let m = this
      m.addDepValidate.disabled = !m.addDepValidate.disabled
      Put('/admin/dep', m.addDepValidate).then(resp => {
        this.$Message.success('modify dep Success!')
        m.showCreateDepModal = false
        m.resetAddDep()
      })
    },
    // 创建部门
    createDep () {
      let m = this
      m.addDepValidate.disabled = !m.addDepValidate.disabled
      this.$refs[m.constDep].validate((valid) => {
        if (valid) {
          Post('/admin/dep', m.addDepValidate).then(resp => {
            this.$Message.success('create dep Success!')
            m.resetAddDep()
          })
        }
      })
      // this.$router.push({
      //   name: 'add_dep_page'
      // })
    },
    // 重置部门信息
    resetAddDep () {
      this.$refs[this.constDep].resetFields()

      this.isModify = false
      this.showModalDepTree = false
      this.allDeps = []
      this.selectSuperDep = { superDepName: '', isSelected: true }
      this.addDepValidate = { disabled: true }
      this.text = {typeText: '创建', titleText: '新增部门',}
    }
  },
  watch: {
  },
  computed: {},
  mounted: function () {}
}
</script>
<style scoped>
</style>
