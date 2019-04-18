<template>

  <div>
    <Card>
      <Button class="tool" type="primary" icon="md-person-add" @click="addEntity">新增</Button>
      <Button class="tool" type="warning" icon="md-create" @click="modifyEntity" :disabled="modifyEntityDisable">修改
      </Button>
      <Button class="tool" type="error" icon="md-close" @click="deleteEntitys" :disabled="deleteEntitysDisable">删除
      </Button>
      <Button class="tool" type="dashed" icon="ios-refresh" :loading="refreshEntityLoading" @click="pullTableList">刷新
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
  </div>

</template>

<script>
  import Tables from '_c/tables'
  import {Get,Post,Delete} from '@/api/data'
  import Button from "iview/src/components/button/button";

  export default {
    components: {
      Button,
      Tables,
    },
    data() {
      return {
        modifyEntityDisable: true,
        deleteEntitysDisable: true,
        refreshEntityLoading: false,
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
          {
            title: '操作',
            key: 'handle',
            options: ['delete'],
            align: 'center',
            width: 100,
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
                    }
                  }
                })
              }
            ]
          }
        ],
        entityBaseUrl:'/admin/menu',
        entityFormName:'menu_form_page',
        tableData: [],
        selection:[],
        page: {
          start: 1,
          size: 10,
          searchKey :'',
          searchValue:''
        }
      }
    },
    methods: {
      handleDelete(params) {
        console.log(params)
        let url = this.entityBaseUrl + '/' + params.row.id
        Delete(url).then(resp => {
          this.total = this.total -1;
        })
      },
      addEntity() {
        this.$router.push({
          name: this.entityFormName
        })
      },
      modifyEntity() {
        console.log('modifyEntity ...')
        console.log(this.selection);
        if(this.selection.length === 1){
          this.$router.push({name: this.entityFormName, params: {id: this.selection[0].id}})
        }
      },
      deleteEntitys() {
        let vm = this;
        let newPage = this.page;
        function callBackOk(){
          console.log("callBackOk",vm.selection);
          let ids = "";
          vm.selection.forEach(item =>{
            ids += item.id+",";
          });

          // ids = ids.substr(0,ids.length-1);

          // let url = vm.entityBaseUrl;
          let url = vm.entityBaseUrl + '/del';
          // let url = String(vm.entityBaseUrl + '/'+ids) ;

          console.log(url);

          Delete(url,{ids:ids}).then(resp => {
            vm.pullTableList(newPage)
          })
        }
        this.$Modal.confirm({
          title: '确认删除选中的类别?',
          // content: '<p>Content of dialog</p><p>Content of dialog</p>',
          closable: true,
          onOk() {
            callBackOk()
          },
          onCancel() {
          }
        })
      },

      // 监听选择的表格
      selectChange(selection) {

        if(selection instanceof Array){
          selection.length === 1 ? this.modifyEntityDisable = false : this.modifyEntityDisable = true;
          selection.length > 0 ? this.deleteEntitysDisable = false : this.deleteEntitysDisable = true;
          this.selection = selection;
        }
      },

      pullTableList(page) {
        this.tableData = []
        this.refreshEntityLoading = true

        this.page = page;

        if (this.page.start === undefined) this.page.start = 1
        if (this.page.size === undefined) this.page.size = 10
        let url = this.entityBaseUrl + '?start=' + this.page.start + '&size=' + this.page.size ;

        if(! (this.page.searchValue === undefined)){
          url +=  '&searchKey=' + this.page.searchKey + '&searchValue=' + this.page.searchValue
        }

        console.log('url=', url)
        // return
        Get(url).then(resp => {
          this.total = resp.data.total
          // resp.data.rows.forEach(d => this.tableData.push(d))
          this.tableData = resp.data.rows
        })
        setTimeout(() => {
          this.refreshEntityLoading = false
        }, 1.5 * 1000)
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
