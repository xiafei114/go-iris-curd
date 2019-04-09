<template>

  <div>
    <Card shadow>
      <Button class="tool" type="primary" icon="md-person-add" @click="showCreateDepModal">新增</Button>
      <Button class="tool" type="dashed" icon="ios-refresh" :loading="refreshLoading" @click="pullTableData">刷新</Button>
      <!--<Button class="tool" type="error" icon="md-close" @click="deleteDeps" :disabled="deleteDepsDisable">删除</Button>-->
      <zk-table
        ref="table"
        :show-header="true"
        :show-index="false"
        :tree-type="true"
        :is-fold="false"
        :expand-type="false"
        :selection-type="false"
        max-height="470"
        :columns="columns"
        :data="tableData"
      >
        <template slot="opts" scope="scope">
          <Button class="tool2" type="success" icon="md-create" @click="modifyDep(scope)" size="small">编辑</Button>
          <Button class="tool2" type="error" icon="ios-trash" @click="deleteOneDep(scope)" size="small">删除</Button>
        </template>
      </zk-table>

    </Card>

    <dep-modals ref="depModels"></dep-modals>
  </div>

</template>

<script>
import Tables from '_c/tables'
import DepModals from './dep_modals'
import { Get, Delete } from '@/api/data'
export default {
  components: {
    Tables,
    DepModals,
  },
  data () {
    return {
      // total: 102,
      tableData: [],
      columns: [
        {label: '部门名称', prop: 'depName', width:'150px'},
        {label: '领导名称', prop: 'leader', width:'95px'},
        {label: '手机号码', prop: 'phone', width:'110px'},
        {label: '邮箱', prop: 'email', width:'140px'},
        // {label: '可用否', prop: 'disabled', width:'70px'},
        {label: '部门描述', prop: 'depDesc', width:'380px'},
        {label: '创建时间', prop: 'createTime', width:'180px'},
        {label: '操作', prop: 'opts', minWidth: '170px', type: 'template', template: 'opts',},
      ],

      refreshLoading: false,
    }
  },
  methods: {
    // 显示新建部门modal
    showCreateDepModal () { this.$refs.depModels.changeCreateDepModal(true, false, null) },
    // 修改部门
    modifyDep (scope) {
      console.log(scope)
      let {id, depName, leader, phone, email, parentId, depDesc, disabled, superName, createTime} = scope.row
      let dep = {id: id, depName: depName, leader: leader, phone: phone, email: email, parentId: parentId, depDesc: depDesc, disabled: !disabled, createTime: createTime}
      let info = {dep: dep, superName: superName}
      console.log(info)
      this.$refs.depModels.changeCreateDepModal(true, true, info)
    },
    // 删除部门
    deleteOneDep (scope) {
      let {id, depName} = scope.row
      if (id === 1) {
        this.$Message.warning({content: '请不要删除根节点'})
        return
      }
      this.$Modal.confirm({
        title: '删除部门',
        content: '部门名称：<strong>' + depName + '</strong>',
        loading: true,
        closable: true,
        okText: '删除',
        onOk: () => {
          Delete('/admin/dep/' + id).then(resp => {
            this.$Message.success({content: '删除成功'})
            this.$Modal.remove()
            this.pullTableData()
          })
        },
        cancelText: '取消',
        onCancel: () => {this.$Message.info({content: '取消删除'})}
      })
    },

    // 拉去部门表格
    pullTableData() {
      this.refreshLoading = true
      Get("/admin/dep").then(resp => {
        let d = resp.data.data
        this.tableData = []
        this.tableData.push(d)
      })
      setTimeout(() => {this.refreshLoading = false}, 1 * 1000)
    },
  },
  watch: {
  },
  computed: {
  },
  mounted: function () {
    this.pullTableData()
  }
}
</script>
<style scoped>
.tool {
  margin: 0 5px 7px 5px;
}
.tool2 {
  margin: 0 2px;
}
</style>
