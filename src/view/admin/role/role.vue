<template>

  <div>
    <Card>
      <Button class="tool" type="primary" icon="md-person-add" @click="addMenu">新增</Button>
      <Button class="tool" type="warning" icon="md-create" @click="modifyMenu" :disabled="modifyMenuDisable">修改
      </Button>
      <Button class="tool" type="error" icon="md-close" @click="deleteMenus" :disabled="deleteMenusDisable">删除
      </Button>
      <Button class="tool" type="dashed" icon="ios-refresh" :loading="refreshMenuLoading" @click="pullMenuTable">刷新
      </Button>
      <tables ref="tables"
              searchable
              search-place="top"
              v-model="tableData"
              :columns="columns"
              :total="total"
              :highlightRow=true
              @change-page="pullMenuTable"
              @on-delete="handleDelete"
              @on-selection-change="selectChange"
      />
    </Card>
  </div>
</template>

<script>
  import Tables from '_c/tables'
  import {Get} from '@/api/data'

  export default {
    components: {
      Tables,
    },
    data() {
      return {
        modifyMenuDisable: true,
        deleteMenusDisable: true,

        total: 0,
        columns: [
          {title: '', key: 'handle', type: 'selection', width: 50, align: 'center'},
          {title: '名称', key: 'name', sortable: true, width: 90},
          {title: '路径', key: 'path', width: 100},
          {title: '模块', key: 'modular', width: 80},
          {title: '组件', key: 'component', width: 60},
          {title: '子菜单', key: 'isSub', width: 80},
          // {title: '父菜单ID', key: 'parentId', editable: true, width: 70},
          {title: '元信息', key: 'meta', width: 400},
          {title: '创建时间', key: 'createTime'},
          {title: '更新时间', key: 'updateTime'},
          // {
          //   title: 'Handle',
          //   key: 'handle',
          //   options: ['delete'],
          //   button: [
          //     (h, params, vm) => {
          //       return h('Poptip', {
          //         props: {
          //           confirm: true,
          //           title: '你确定要删除吗?'
          //         },
          //         on: {
          //           'on-ok': () => {
          //             vm.$emit('on-delete', params)
          //             vm.$emit('input', params.tableData.filter((item, index) => index !== params.row.initRowIndex))
          //           }
          //         }
          //       }, [
          //         h('Button', '自定义删除')
          //       ])
          //     }
          //   ]
          // }
        ],
        tableData: [],

        refreshMenuLoading: false
      }
    },
    methods: {
      addMenu() {
      },
      modifyMenu() {
      },
      deleteMenus() {
      },
      handleDelete() {
      },

      pullMenuTable({start, size}) {
        this.tableData = []
        this.refreshMenuLoading = true

        if (start === undefined) start = 1
        if (size === undefined) size = 10
        let url = '/admin/menu?start=' + start + '&size=' + size
        console.log('menu table url=', url)
        // return
        Get(url).then(resp => {
          this.total = resp.data.total
          resp.data.rows.forEach(d => {
            d.meta = JSON.stringify(d.meta)
            d.isSub = JSON.stringify(d.isSub)
            this.tableData.push(d)
          })
          // this.tableData = resp.data.rows
        })
        setTimeout(() => {
          this.refreshMenuLoading = false
        }, 1.5 * 1000)
      },
      selectChange(selection) {
        (selection instanceof Array && selection.length === 1)
          ? this.modifyMenuDisable = false
          : this.modifyMenuDisable = true;
        (selection instanceof Array && selection.length > 0)
          ? this.deleteMenusDisable = false
          : this.deleteMenusDisable = true
      },
    },
    mounted() {
      this.pullMenuTable({})
    },
  }

</script>
<style scoped>
  .tool {
    margin: 0 5px;
  }
</style>
