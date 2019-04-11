<template>

  <div>
    <Row :gutter="16">
      <Card>
        <Button class="tool" type="primary" icon="md-person-add" @click="addUser">新增</Button>
        <Button class="tool" type="warning" icon="md-create" @click="modifyUser" :disabled="modifyUserDisable">修改
        </Button>
        <Button class="tool" type="error" icon="md-close" @click="deleteUsers" :disabled="deleteUsersDisable">删除
        </Button>
        <Button class="tool" type="dashed" icon="ios-refresh" :loading="refreshUserLoading" @click="pullTableList">刷新
        </Button>
        <tables ref="tables"
                searchable
                search-place="top"
                v-model="tableData"
                :columns="columns"
                :total="total"
                :highlightRow=true
                @change-page="pullTableList"
                @change-size="pullTableList"
                @handle-Search="pullTableList"
                @on-delete="handleDelete"
                @on-selection-change="selectChange"
        />
      </Card>
    </Row>
  </div>

</template>

<script>
  import Tables from '_c/tables'
  import {Get,Delete} from '@/api/data'
  import Button from "iview/src/components/button/button";

  export default {
    components: {
      Button,
      Tables,
    },
    data() {
      return {

        modifyUserDisable: true,
        deleteUsersDisable: true,

        refreshUserLoading: false,
        total: 0,
        columns: [
          {title: '', key: 'handle', type: 'selection', width: 50, align: 'center'},
          {title: '编号', key: 'productCode', sortable: true},
          {title: '品名', key: 'productName', sortable: true},
          {title: '金额', key: 'price', sortable: true},
          {title: '数量', key: 'number', sortable: true},
          {title: '创建时间', key: 'handle'},
          {title: '更新时间', key: 'handle'},
          {
            title: '操作',
            key: 'handle',
            options: ['delete'],
            align: 'center',
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
                })
              }
            ]
          }
        ],
        tableData: []
      }
    },
    methods: {
      handleDelete(params) {  //删除
        console.log(params.row)

        let url = '/product/' + params.row.id
        Delete(url).then(resp => {
          this.total = this.total -1;
        })
      },
      addUser() { //添加产品
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

      //获得内容
      pullTableList({start, size,searchKey,searchValue}) {
        this.tableData = []
        this.refreshUserLoading = true

        if (start === undefined) start = 1
        if (size === undefined) size = 10
        let url = '/product?start=' + start + '&size=' + size

        if(! (searchValue === undefined)){
          url +=  '&searchKey=' + searchKey + '&searchValue=' + searchValue
        }
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
      handleClear(searchKey,searchValue){
        console.log(searchKey);
        console.log(searchValue);
      },
      pullData() {
        this.pullTableList({})
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
