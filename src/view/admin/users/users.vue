<template>

  <div>
    <Row :gutter="16">
      <Col span="5">
      <Collapse v-model="org">
        <Panel :hide-arrow="true" name='1'>
          组织机构
          <Button @click="pullDepTree" :loading="refreshDepTreeLoading" slot="content" icon="ios-refresh" type="dashed"
                  size="small" long>刷新部门
          </Button>
          <Tree slot="content" :data="depTree" @on-select-change="selectDep"></Tree>
        </Panel>
      </Collapse>
      </Col>

      <Col span="19">
      <Card>
        <Button class="tool" type="primary" icon="md-person-add" @click="addUser">新增</Button>
        <Button class="tool" type="warning" icon="md-create" @click="modifyUser" :disabled="modifyUserDisable">修改
        </Button>
        <Button class="tool" type="error" icon="md-close" @click="deleteUsers" :disabled="deleteUsersDisable">删除
        </Button>
        <Button class="tool" type="dashed" icon="ios-refresh" :loading="refreshUserLoading" @click="pullUserTable">刷新
        </Button>
        <tables ref="tables"
                searchable
                search-place="top"
                v-model="tableData"
                :columns="columns"
                :total="total"
                :highlightRow=true
                @change-page="pullUserTable"
                @on-delete="handleDelete"
                @on-selection-change="selectChange"
        />
      </Card>
      </Col>
    </Row>
  </div>

</template>

<script>
  import Tables from '_c/tables'
  import {Get} from '@/api/data'
  import Button from "iview/src/components/button/button";

  export default {
    components: {
      Button,
      Tables,
    },
    data() {
      return {
        org: '1',
        depTree: [],
        selectDepId: -1,
        refreshDepTreeLoading: false,

        modifyUserDisable: true,
        deleteUsersDisable: true,

        refreshUserLoading: false,
        total: 0,
        columns: [
          {title: '', key: 'handle', type: 'selection', width: 50, align: 'center'},
          {title: '账号', key: 'username', sortable: true},
          {title: '名字', key: 'name', sortable: true},
          {title: '可用', key: 'enable', sortable: true},
          {title: '性别', key: 'gender', sortable: true},
          {title: '电话', key: 'phone', editable: true},
          {title: '邮件', key: 'email', editable: true},
          {title: '头像', key: 'userface', editable: true},
          {title: '创建时间', key: 'createTime'},
          {title: '更新时间', key: 'updateTime'},
          {
            title: 'Handle',
            key: 'handle',
            options: ['delete'],
            button: [
              (h, params, vm) => {
                return h('Poptip', {
                  props: {
                    confirm: true,
                    title: '你确定要删除吗?'
                  },
                  on: {
                    'on-ok': () => {
                      vm.$emit('on-delete', params)
                      vm.$emit('input', params.tableData.filter((item, index) => index !== params.row.initRowIndex))
                    }
                  }
                }, [
                  h('Button', '自定义删除')
                ])
              }
            ]
          }
        ],
        tableData: [
          // {name: 'Thomas Jackson', email: '1', createTime: 's'},
          // {name: 'Thomas Jackson2', email: '2', createTime: 's'},
          // {name: 'Thomas Jackson3', email: '3', createTime: 's'},
        ]
      }
    },
    methods: {
      handleDelete(params) {
        console.log(params)
      },
      selectDep(dep) {
        if (dep instanceof Array && dep.length === 1) {
          let {id} = dep[0]
          console.log(id)
        }
      },
      addUser() {
        this.$router.push({
          name: 'add_user_page'
        })
      },
      modifyUser() {
        console.log('modifyUser ...')
      },
      deleteUsers() {
        this.$Modal.confirm({
          title: '确认删除选中的用户?',
          // content: '<p>Content of dialog</p><p>Content of dialog</p>',
          closable: true,
          onOk() {
            alert('删除了用户')
          },
          onCancel() {
          }
        })
      },

      // 监听选择的表格
      selectChange(selection) {
        (selection instanceof Array && selection.length === 1)
          ? this.modifyUserDisable = false
          : this.modifyUserDisable = true;
        (selection instanceof Array && selection.length > 0)
          ? this.deleteUsersDisable = false
          : this.deleteUsersDisable = true
      },

      // 拉取部门树
      pullDepTree() {
        this.depTree = []
        this.refreshDepTreeLoading = true
        Get('/admin/users/dep').then(resp => {
          this.depTree.push(resp.data.data)
        })
        setTimeout(() => {
          this.refreshDepTreeLoading = false
        }, 1 * 1000)
      },
      pullUserTable({start, size}) {
        this.tableData = []
        this.refreshUserLoading = true

        if (start === undefined) start = 1
        if (size === undefined) size = 10
        let url = '/admin/users?start=' + start + '&size=' + size + '&depId=' + this.selectDepId
        console.log('url=', url)
        // return
        Get(url).then(resp => {
          this.total = resp.data.total
          // resp.data.rows.forEach(d => this.tableData.push(d))
          this.tableData = resp.data.rows
        })
        setTimeout(() => {
          this.refreshUserLoading = false
        }, 1.5 * 1000)
      },
      pullData() {
        this.pullDepTree()
        this.pullUserTable({})
      }
    },
    computed: {},
    mounted() {
      this.pullData()
    }
  }
</script>
<style scoped>
  .tool {
    margin: 0 5px;
  }
</style>
