<template>

  <div>
    <Row :gutter="16">
      <Card>
        <Button class="tool" type="primary" icon="md-add" @click="addEntity">新增</Button>
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
        modifyEntityDisable: true,
        deleteEntitysDisable: true,
        refreshEntityLoading: false,

        total: 0,
        columns: [
          {title: '', key: 'handle', type: 'selection', width: 50, align: 'center'},
          {title: '编号', key: 'product_Code', sortable: true},
          {title: '品名', key: 'product_Name', sortable: true},
          {title: '金额', key: 'price', sortable: true},
          {title: '数量', key: 'number', sortable: true},
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
                      this.$emit('input', params.tableData.filter((item, index) => index !== params.row.initRowIndex))
                    }
                  }
                })
              }
            ]
          }
        ],
        tableData: [],  // 数据
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
      handleDelete(params) {  //删除
        console.log(params.row)
        let url = '/product/' + params.row.id
        Delete(url).then(resp => {
          this.total = this.total -1;
        })
      },
      addEntity() { //添加产品
        this.$router.push({
          name: 'add_product_page'
        })
      },
      modifyEntity() {
        console.log('modifyEntity ...')
      },
      deleteEntitys() {
        let vm = this;
        let newPage = this.page;
        function callBackOk(){
          console.log(vm.selection);

          let ids = "";
          vm.selection.forEach(item =>{
            ids += item.id+",";
          });

          let url = '/product/' + ids
          Delete(url).then(resp => {
            vm.pullTableList(newPage)
          })
        }
        this.$Modal.confirm({
          title: '确认删除选中的产品?',
          // content: '<p>Content o用f dialog</p><p>Content of dialog</p>',
          closable: true,
          onOk() {
            callBackOk();
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

      //获得内容
      pullTableList(page) {
        this.tableData = []
        this.refreshEntityLoading = true

        this.page = page;

        if (this.page.start === undefined) this.page.start = 1
        if (this.page.size === undefined) this.page.size = 10
        let url = '/product?start=' + this.page.start + '&size=' + this.page.size

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
